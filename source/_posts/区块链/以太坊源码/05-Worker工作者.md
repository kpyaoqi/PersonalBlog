---
title: 05-Worker工作者

date: 2022-11-03	

categories: 以太坊源码	

tags: [区块链,以太坊源码]
---	

# Mining：挖矿

## Block Reward：区块奖励

## How to Seal Block：组装区块

```go
// miner/worker.go
// Worker是负责向共识引擎提交新工作的主要对象和收集密封结果。
type worker struct {
	config      *Config
	chainConfig *params.ChainConfig
	engine      consensus.Engine
	eth         Backend				// eth的协议
	chain       *core.BlockChain	// 区块链

	// Feeds
	pendingLogsFeed event.Feed

	// Subscriptions
	mux          *event.TypeMux
	txsCh        chan core.NewTxsEvent
	txsSub       event.Subscription
	chainHeadCh  chan core.ChainHeadEvent
	chainHeadSub event.Subscription

	// Channels
	newWorkCh          chan *newWorkReq
	getWorkCh          chan *getWorkReq
	taskCh             chan *task
	resultCh           chan *types.Block
	startCh            chan struct{}
	exitCh             chan struct{}
	resubmitIntervalCh chan time.Duration
	resubmitAdjustCh   chan *intervalAdjust

	wg sync.WaitGroup

	current *environment 	// 当前运行循环的环境。

	mu       sync.RWMutex 	// 用于保护coinbase和额外字段的锁
	coinbase common.Address	// coinbase的地址
	extra    []byte

	pendingMu    sync.RWMutex
	pendingTasks map[common.Hash]*task

	snapshotMu       sync.RWMutex 	// 快照 RWMutex（快照读写锁）
	snapshotBlock    *types.Block
	snapshotReceipts types.Receipts
	snapshotState    *state.StateDB // 快照StateDB

	//原子状态计数器
	running atomic.Bool  // The indicator whether the consensus engine is running or not.
	newTxs  atomic.Int32 // New arrival transaction count since last sealing work submitting.
	syncing atomic.Bool  // The indicator whether the node is still syncing.

	// newpayloadTimeout是创建有效负载的最大超时。
	//默认值为2秒，但节点操作符可以将其设置为任意大的值。过大的超时允许可能导致Geth无法在指定的时间内创建非空负载，并且在txpool	  中存在一些计算开销较大的事务的情况下最终错过插槽。
	newpayloadTimeout time.Duration

	//重新提交是在权益证明阶段重新创建密封工作或重新构建有效载荷的时间间隔。
	recommit time.Duration

	// External functions
    //用于确定指定块是否被本地矿工挖掘的函数。
	isLocalBlock func(header *types.Header) bool 

	// Test hooks
	newTaskHook  func(*task)                        //接收到新的封装任务时调用的方法
	skipSealHook func(*task) bool                   //决定是否跳过密封的方法。
	fullTaskHook func()                             //在推送完全密封任务之前调用的方法。
	resubmitHook func(time.Duration, time.Duration) //更新重新提交间隔时调用的方法。
}
```

其中一个关键的函数是`miner/worker.go`中的`fillTransactions()`函数。

也就说如果我们希望修改Block中Transaction的打包顺序和从Transaction Pool选择Transactions的策略的话，我们可以通修改`fillTransactions()`函数。

```go
// miner/worker.go
// 从txpool中检索待处理的交易，并将它们填充到给定的密封块中，交易选择和排序策略可以在将来使用插件进行定制。
func (w *worker) fillTransactions(interrupt *atomic.Int32, env *environment) error {
    // 用所有可用的待定交易填充块。
	// 将待处理交易拆分为本地和远程交易
	pending := w.eth.TxPool().Pending(true)
	localTxs, remoteTxs := make(map[common.Address][]*types.Transaction), pending
	for _, account := range w.eth.TxPool().Locals() {
		if txs := remoteTxs[account]; len(txs) > 0 {
			delete(remoteTxs, account)
			localTxs[account] = txs
		}
	}
    //首先处理Local Pool中的交易
	if len(localTxs) > 0 {
        //按照GasPrice和Nonce的顺序进行排序形成新的txs并传递给commitTransactions()函数
		txs := types.NewTransactionsByPriceAndNonce(env.signer, localTxs, env.header.BaseFee)
		if err := w.commitTransactions(env, txs, interrupt); err != nil {
			return err
		}
	}
    //然后再处理从网络中接受到的远程交易。
	if len(remoteTxs) > 0 {
		txs := types.NewTransactionsByPriceAndNonce(env.signer, remoteTxs, env.header.BaseFee)
		if err := w.commitTransactions(env, txs, interrupt); err != nil {
			return err
		}
	}
	return nil
}
```

`commitTransactions()`函数的主体是一个for循环体。在这个for循环中，函数会从txs中不断拿出头部的tx进行调用`commitTransaction()`函数进行处理。在Transaction那一个Section我们提到的`commitTransaction()`函数会将成功执行的Transaction保存在`env.txs`中。

```go
func (w *worker) commitTransactions(env *environment, txs *types.TransactionsByPriceAndNonce, interrupt *atomic.Int32) error {
	.....
	for {
		.....
		logs, err := w.commitTransaction(env, tx)
		.....
	}
	.....
	return nil
}
```

```go
func (w *worker) commitTransaction(env *environment, tx *types.Transaction) ([]*types.Log, error) {
	.....
    //ApplyTransaction尝试将交易应用到给定的状态数据库，并使用其环境的输入参数。
    //它返回交易的收据、使用的gas，如果交易失败则返回一个错误，表明该区块无效。
	receipt, err := core.ApplyTransaction(w.chainConfig, w.chain, &env.coinbase, env.gasPool, env.state, env.header, tx, &env.header.GasUsed, *w.chain.GetVMConfig())
	.....
	return receipt.Logs, nil
}
```


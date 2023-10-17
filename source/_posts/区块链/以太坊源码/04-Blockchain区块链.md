---
title: 04-Blockchain区块链

date: 2022-10-30	

categories: 以太坊源码	

tags: [区块链,以太坊源码]
---	

# Blockchain

## Block区块数据结构

```go
// core/types/block.go
// Block代表以太坊区块链中的整个区块
type Block struct {
	header       *Header
	uncles       []*Header
	transactions Transactions
	// 缓存
	hash atomic.Value
	size atomic.Value
	//这些字段被包eth用来跟踪
	//点间块中继
	ReceivedAt   time.Time
	ReceivedFrom interface{}
}
```

```go
// Header表示以太坊区块链中的区块头
type Header struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"`
	Nonce       BlockNonce     `json:"nonce"`
	// BaseFee由EIP-1559添加，在遗留的报头中被忽略
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

}
```

## Blockchain区块链数据结构

```go
// core/blockchain.go
// 区块链还有助于从数据库中包含的任何链中返回块，以及代表规范链的块。值得注意的是，GetBlock可以返回任何块，而不需要包含在规范块	中，因为GetBlockByNumber总是代表规范链。
type BlockChain struct {
	chainConfig *params.ChainConfig // 链和网络配置
	cacheConfig *CacheConfig        // 缓存配置修剪
	db            ethdb.Database                   // 用于存储最终内容的底层持久数据库
	snaps         *snapshot.Tree                   // 快照树快速单词查找树的叶子的访问
	triegc        *prque.Prque[int64, common.Hash] // 优先级队列映射块编号到尝试gc
	gcproc        time.Duration                    // Trie转储的累积规范块处理
	lastWrite     uint64                           // 刷新状态时的最后一个块
	flushInterval atomic.Int64                     // 刷新状态的时间间隔(处理时间)
	triedb        *trie.Database                   // 用于维护trie节点的数据库处理程序。
	stateCache    state.Database                   // 要在导入之间重用的状态数据库（包含状态缓存）

	// txlookupllimit是从head中保留tx索引的最大块数:
    // * 0:表示没有限制，并重新生成任何缺失的索引
    // * N:表示N块限制[HEAD-N+1, HEAD]，并删除多余的索引
    // * nil:禁用tx索引器/删除器，但仍然索引新的块
	txLookupLimit uint64

	hc            *HeaderChain				// 只包含了区块头的区块链
	rmLogsFeed    event.Feed  				// 下面是很多消息通知的组件
	chainFeed     event.Feed
	chainSideFeed event.Feed
	chainHeadFeed event.Feed
	logsFeed      event.Feed
	blockProcFeed event.Feed
	scope         event.SubscriptionScope
	genesisBlock  *types.Block				// 创世区块

	//同步互斥锁链的写操作。
	//读者不需要拿，他们可以直接读取数据库。
	chainmu *syncx.ClosableMutex

	currentBlock      atomic.Pointer[types.Header] // 当前链头
	currentSnapBlock  atomic.Pointer[types.Header] // 当前快照同步头
	currentFinalBlock atomic.Pointer[types.Header] // 最新(共识)完成的区块
	currentSafeBlock  atomic.Pointer[types.Header] // 最新的(共识)安全的块

	bodyCache     *lru.Cache[common.Hash, *types.Body]
	bodyRLPCache  *lru.Cache[common.Hash, rlp.RawValue]
	receiptsCache *lru.Cache[common.Hash, []*types.Receipt]
	blockCache    *lru.Cache[common.Hash, *types.Block]
	txLookupCache *lru.Cache[common.Hash, *rawdb.LegacyTxLookupEntry]

	//未来块是为以后的处理添加的块
	futureBlocks *lru.Cache[common.Hash, *types.Block]

	wg            sync.WaitGroup 
	quit          chan struct{}  //关闭信号，在停止中关闭。
	stopping      atomic.Bool    //如果链正在运行则为false，停止时为true
	procInterrupt atomic.Bool    //块处理的中断信号

	engine     consensus.Engine // 一致性引擎
	validator  Validator 		// 块和状态验证者接口
	prefetcher Prefetcher
	processor  Processor 		// 块交易处理程序接口
	forker     *ForkChoice
	vmConfig   vm.Config
}
```

## 常见方法

NewBlockChain：返回一个使用数据库中可用信息的完全初始化的区块链。它初始化了默认的以太坊验证器和处理器。

loadLastState：从数据库中加载最后一个已知的链状态。此方法假定持有链管理器互斥锁。

Reset：清除整个区块链，将其恢复到初始状态。

SetHead：将本地链倒回到一个新的头。根据节点是快速同步还是完全同步以及处于哪种状态，该方法将尝试从磁盘中删除最小的数据，同时保持链的一致性。

InsertChain：尝试将给定批次的块插入到规范链中，否则，创建一个分叉。如果返回一个错误，它将返回失败块的索引号以及描述错误的错误。插入完成后，将触发所有累积的事件。

insertChain：是InsertChain的内部实现，它假设链是连续的，且 链互斥锁被持有。这个方法被分离出来，这样需要重新注入历史块的导入批可以在不释放锁的情况下完成，否则会导致不稳定的行为。如果侧链导入正在进行中，并且导入了历史状态，但是在实际侧链完成之前添加了新的canon-head，则可以再次修剪历史状态

writeblockandsehead将给定块和所有相关状态写入数据库，并应用该块作为新的链头。

reorg：获取两个区块，一个旧链和一个新链，并将重建区块并将它们插入到新的规范链中，并积累潜在的缺失交易并发布关于它们的事件。注意，这里不会处理新的头部块，调用者需要在外部处理它。

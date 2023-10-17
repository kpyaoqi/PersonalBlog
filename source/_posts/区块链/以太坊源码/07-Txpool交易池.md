---
title: 07-Txpool交易池

date: 2022-11-09	

categories: 以太坊源码	

tags: [区块链,以太坊源码]
---	

# 概述

txpool主要用来存放当前提交的等待写入区块的交易

Transaction Pool 里主要是Subpool结构，里面的 Pending 为当前可处理的交易

```go
// core/txpool/txpool.go
type TxPool struct {
	subpools []SubPool               // 用于专门交易处理的子池列表
	subs     event.SubscriptionScope // 关闭时取消所有订阅的订阅范围
	quit     chan chan error         // 退出频道以拆除head更新
}
```

```go
// SubPool 代表一个独立存在的专用事务池（例如 blob 池），由于与我们拥有多少个专用池无关，它们确实需要同步更新并组装成一个连贯的视图以进行块生产，因此该接口定义了允许主交易池管理子池的通用方法。
type SubPool interface {
	// Filter 是一个选择器，用于决定是否将交易添加到该特定子池中
	Filter(tx *types.Transaction) bool
	// Init 设置子池的基本参数
	Init(gasTip *big.Int, head *types.Header) error
	Close() error
	// 重置检索区块链的当前状态，并确保交易池的内容对于链状态而言是有效的。
	Reset(oldHead, newHead *types.Header)
	// 更新子池对新交易所需的最低价格，并将所有交易降低到低于此阈值。
	SetGasTip(tip *big.Int)
	// 返回子池是否具有使用给定哈希缓存的事务的指示符。
	Has(hash common.Hash) bool
	Get(hash common.Hash) *Transaction
	// 如果一批交易有效，则 Add 将其放入池中。 由于交易波动较大，add 可能会推迟将交易完全集成到稍后的时间点，以便将多个交易批量	   处理在一起。
	Add(txs []*Transaction, local bool, sync bool) []error
	// Pending 检索所有当前可处理的交易，按原始帐户分组并按随机数排序。
	Pending(enforceTips bool) map[common.Address][]*types.Transaction
	SubscribeTransactions(ch chan<- core.NewTxsEvent) event.Subscription
	// Nonce 返回帐户的下一个随机数，池中可执行的所有交易都已应用于顶部。
	Nonce(addr common.Address) uint64
	Stats() (int, int)
	// Content 检索交易池的数据内容，返回所有待处理和排队的交易，按帐户分组并按随机数排序。
	Content() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction)
	// ContentFrom 检索交易池的数据内容，返回该地址的待处理交易以及排队交易，并按随机数分组。
	ContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction)
	// Locals 检索当前被池视为本地的帐户.
	Locals() []common.Address
	// Status 返回由哈希值标识的事务的已知状态（未知/待处理/排队）
	Status(hash common.Hash) TxStatus
}
```


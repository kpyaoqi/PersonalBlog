---
title: 01-Accounts账户

date: 2022-10-17	

categories: 以太坊源码	

tags: [区块链,以太坊源码]
---	

# 账户

Accounts包实现了以太坊客户端的**钱包和账户管理**，以太坊的钱包**提供了keyStore模式和usb两种钱包**，在以太坊中，有两种类型的Account，**分别是外部账户(EOA)以及合约账户(Contract)**。在以太坊中，**State 对应的基本数据结构，称为 StateObject**。当StateObject 的值发生了变化时，我们称为**状态转移**。在 Ethereum 的运行模型中，**StateObject 所包含的数据会因为 Transaction 的执行引发数据更新/删除/创建，引发状态转移**，我们说：StateObject 的状态从当前的 State 转移到另一个 State。

## 账号

```go
// accounts/accounts.go
// 一个账号是20个字节的数据。 URL是可选的字段。
type Account struct {
	Address common.Address `json:"address"` // 从密钥导出的以太坊账户地址
	URL     URL            `json:"url"`     // 后端中的可选资源定位器
}
```

在实际代码中，这两种 Account 都是由`stateObject`这一数据结构定义的

```go
// core/state/state_object.go
// statobject代表一个正在被修改的以太坊账户
//使用模式如下:首先你需要获取一个状态对象-->通过对象访问和修改帐户值-->调用commit将修改后的存储树写入数据库
type stateObject struct {
	address  common.Address 	// 对应了一个20字节长的byte类型数组
	addrHash common.Hash 		// ethereum的散列地址的帐户,一个32字节长的byte类型数组
	data     types.StateAccount
	db       *StateDB
	// 写缓存
	trie Trie // 存储Tire，在第一次访问时变为非nil
	code Code // 合同的字节码,加载时设置代码
	 // 这里的Storage 是一个 map[common.Hash]common.Hash
	originStorage  Storage // 存储缓存原始条目dedup重写,重置为每笔交易
	pendingStorage Storage // 存储条目,需要刷新到磁盘,在整个区块结束后
	dirtyStorage   Storage // 存储条目,修改当前交易执行
	// 缓存标志.
	// 当一个对象标记为自杀,那么它将从单词查找树中删除
	// 在“更新”阶段的状态转换
	dirtyCode bool // 如果代码已更新，则为True
	suicided  bool
	deleted   bool
}
```

```go
// common/types.go
const (
	HashLength    = 32
	AddressLength = 20
)

type Address [AddressLength]byte
type Hash [HashLength]byte
```

```go
// core/types/state_account.go
type StateAccount struct {
  Nonce    uint64		// 该账户发送的交易序号
  Balance  *big.Int		// 该账户的余额
  Root     common.Hash  // 当前账户的下Storage层的MPT的Root,管理合约中持久化变量,对于EOA账户这个部分为空值
  CodeHash []byte		// 该账户的Contract代码的哈希值,对于 EOA账户这个部分为空值
}
```

DB:**db这个变量保存了一个 `StateDB` 类型的指针**。这是**为了方便调用 `StateDB` 相关的API对Account所对应的 `stateObject` 进行操作**。StateDB本质上是用于管理`stateObject`信息的而**抽象出来的内存数据库**。**所有的Account 数据的更新，检索都会使用 StateDB 提供的API。**

Cache:对于剩下的成员变量，它们的**主要用于内存Cache**。**trie用于保存和管理合约账户中的持久化变量存储的数据**，**code用于缓存合约中的代码段到内存中，它是一个byte数组**。剩下的四个Storage 字段主要**在执行 Transaction 的时候缓存合约修改的持久化数据**，比如dirtyStorage 就用于缓存在 Block 被 Finalize 之前，**Transaction所修改的合约中的持久化存储数据**。对于**外部账户，由于没有代码字段，所以对应 stateObject 对象中的code 字段，以及三个 Storage 类型的字段对应的变量的值都为空**

### 账户生成

EOA账户的**创建分为本地创建和链上注册**两个部分。当我们使用诸如 Metamask 等钱包工具创建账户的时候，在区块链上并没有同步注册账户信息。**链上账户的创建和管理都是通过`StateDB`模块来操作的**，因此我们将`geth`中账户管理部分的代码整合到`StateDB`模块章节来一起讲述。而**合约账户需要通过 EOA 账户构造特定的交易生成**的。

函数有一个string类型的passphrase参数，这个**参数仅用于加密本地保存私钥的Keystore文件**，与生成账户的私钥，地址的生成都无关

```go
// accounts/keystore/keystore.go
func (ks *KeyStore) NewAccount(passphrase string) (accounts.Account, error) {
    // 生成一个账户需要的私钥和公钥对(大量的椭圆曲线加密)
   _, account, err := storeNewKey(ks.storage, crand.Reader, passphrase)
   if err != nil {
      return accounts.Account{}, err
   }
   // 立即将帐户添加到缓存中
   // 然后等待文件系统通知来拾取它。
   ks.cache.add(account)
   ks.refreshWallets()
   return account, nil
}
```

## 智能合约Contract

在外部账户对应的 stateObject 结构体的实例中，有三个 **Storage 类型(`map[common.Hash]common.Hash`)**的变量是空值，这三个变量是为Contract类型的账户准备的。相比与外部账户，**合约账户额外保存了一个存储层(Storage)用于存储合约代码中持久化的变量的数据。**StateObject 结构体中的声明的三个 **Storage 类型的变量，作为 Contract Storage 层的内存缓存**。

**每个合约都维护了自己的独立的存储空间，用于保存合约中的持久化变量，我们称为 Storage 层。**Storage 层的**基本组成单元称为槽 (Slot)**，Contract 同样**使用 MPT 作为索引树来管理 Storage 层的Slot**。管理**合约账户中 Storage 层 Storage Trie 的根数据被保存在 StateAccount 结构体中的 Root 变量中**，**Storage 层的数据读取和修改是在执行相关 Transaction 的时候**，通过调用 EVM 中的两个专用的**指令*OpSload*和*OpSstore*来实际执行**的。关于这两个指令的具体实现原理，我们会在后续的 EVM 章节进行详细的解读。

# 钱包

- KeyStore
- Private Key
- 助记词

钱包应该是这里面最重要的一个接口了，具体的钱包也是实现了这个接口，钱包又有所谓的**分层确定性钱包和普通钱包**。

```go
// accounts/accounts.go
// Wallet 是指包含了一个或多个账户的软件钱包或者硬件钱包
type Wallet interface {
	// URL 用来获取这个钱包可以访问的规范路径。 它会被上层使用用来从所有的后端的钱包来排序。
	URL() URL
	// 用来返回一个文本值用来标识当前钱包的状态。 同时也会返回一个error用来标识钱包遇到的任何错误。
	Status() (string, error)
	// Open 初始化对钱包实例的访问。这个方法并不意味着解锁或者解密账户，而是简单地建立与硬件钱包的连接和/或访问衍生种子。.
	// passphrase参数可能在某些实现中并不需要。 没有提供一个无passphrase参数的Open方法的原因是为了提供一个统一的接口。 
	// 请注意，如果你open了一个钱包，你必须close它。不然有些资源可能没有释放。 特别是使用硬件钱包的时候需要特别注意。
	Open(passphrase string) error
	// Close 释放由Open方法占用的任何资源。
	Close() error
	// Accounts用来获取钱包发现了账户列表。 对于分层次的钱包， 这个列表不会详尽的列出所有的账号， 而是只包含在帐户派生期间明确固定的帐户。
	Accounts() []Account
	// Contains 返回一个账号是否属于本钱包。
	Contains(account Account) bool
	// Derive尝试在指定的派生路径上显式派生出分层确定性帐户。 如果pin为true，派生帐户将被添加到钱包的跟踪帐户列表中。
	Derive(path DerivationPath, pin bool) (Account, error)
	// SelfDerive设置一个基本帐户导出路径，从中钱包尝试发现非零帐户，并自动将其添加到跟踪帐户列表中。
	// 注意，SelfDerive将递增指定路径的最后一个组件，而不是下降到子路径，以允许从非零组件开始发现帐户。
	// 你可以通过传递一个nil的ChainStateReader来禁用自动账号发现。
	SelfDerive(base DerivationPath, chain ethereum.ChainStateReader)
	SignDataWithPassphrase(account Account, passphrase, mimeType string, data []byte) ([]byte, error)
	SignText(account Account, text []byte) ([]byte, error)
	// SignTextWithPassphrase Signtext相同,但是也需要一个密码
	SignTextWithPassphrase(account Account, passphrase string, hash []byte) ([]byte, error)
	// SignTx 请求钱包对指定的交易进行签名。
	SignTx(account Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
	// SignHashWithPassphrase请求钱包使用给定的passphrase来签名给定的transaction
	SignTxWithPassphrase(account Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
}
```

```go
// accounts/accounts.go
// Backend是一个钱包提供器。 可以包含一批账号。他们可以根据请求签署交易，这样做。
type Backend interface {
	// Wallets获取当前能够查找到的钱包
	// 返回的钱包默认是没有打开的。 
	//所产生的钱包列表将根据后端分配的内部URL按字母顺序排序。 由于钱包（特别是硬件钱包）可能会打开和关闭，所以在随后的检索过程中，相同的钱包可能会出现在列表中的不同位置。
	Wallets() []Wallet
	// 订阅创建异步订阅，以便在后端检测到钱包的到达或离开时接收通知。
	Subscribe(sink chan<- WalletEvent) event.Subscription
}
```

```go
// accounts/manager.go
// Manager是一个包含所有东西的账户管理工具。 可以和所有的Backends来通信来签署交易。
type Manager struct {
	config      *Config                    // 配置
	backends    map[reflect.Type][]Backend 	// 当前注册的后端索引
	updaters []event.Subscription      		 // 钱包更新订阅的所有Backend
	updates  chan WalletEvent           // Backend钱包更改的订阅接收器
    newBackends chan newBackendEvent	// 由manager跟踪的传入Backend
	wallets  []Wallet                   // 所有已经注册的Backends的钱包的缓存
	feed event.Feed 					// 钱包到达和离开的通知
	quit chan chan error	// 退出队列
    term chan struct{}		//通道在更新循环结束时关闭
	lock sync.RWMutex
}
```


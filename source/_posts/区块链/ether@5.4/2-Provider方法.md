---
title: 2-Provider方法

date: 2022-08-07	

categories: ether@5.4.js	

tags: [区块链,ether@5.4.js]
---	

# Providers

**提供者**(**Provider**)是以太坊网络连接的抽象,为标准以太坊节点功能提供简洁、一致的接口.

ethers.js 库 提供了几种选项应该涵盖了绝大多数用例,但如果需要个性化配置,也包括了子类化所需的函数和类.

大多数用户应该使用默认的Provider.

## 1.1 默认的Provider

默认的 provider 是在*Ethereum*上最简单、最安全的方法方式,而且它也具备足够强的鲁棒性,可以在生产环境中使用.

它创建一个连接到尽可能多的后端服务的FallbackProvider.当发出一个请求时,它会同时发送到多个后端.当从每个后端返回响应时,会检查它们是否同意. 一旦达到规定数量(即足够多的后端同意),你的应用程序就会得到相应.它确保后端不同步或者被破坏,那么会放弃这个响应,选择接受而支持大多数认同的响应.

***ethers*.getDefaultProvider( [ *network* , [ *options* ] ] ) ⇒ *Provider***

返回一个新的Provider,由多个服务支持连接到 *网络*.如果没有提供*网络*,那么将会使用 **homestead** (例如mainnet) .

*网络* 也可以作为 URL 来连接,比如 `http://localhost:8545` 或 `wss://example.com`.

***options*** 参数是一个对象,有以下几个属性:

| **Property** | **Description**                                              |      |
| ------------ | ------------------------------------------------------------ | ---- |
| *alchemy*    | Alchemy API Token                                            |      |
| *etherscan*  | Etherscan API Token                                          |      |
| *infura*     | INFURA Project ID 或 `{ projectId, projectSecret }`          |      |
| *pocket*     | Pocket Network Application ID 或 `{ applicationId, applicationSecretKey }` |      |
| *quorum*     | 必须满足规定的后端同意数量 *(默认: 2 for mainnet, 1 for testnets)* |      |

## 1.2 网络

任何接收Networkish 的API都可以传递一个通用名称(如 `"mainnet"` or `"ropsten"`) 或者 chain ID、自定义的参数来定义网络.

*ethers*.*providers*.**getNetwork**( *Networkish* ) ⇒ *Network.*

通过给定标准规范的Networkish,返回完整的 Network .

这对于希望接受Networkish作为输入参数的函数和类非常有用.

```js
// 通过链名
getNetwork("homestead");
// {
//   chainId: 1,
//   ensAddress: '0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e',
//   name: 'homestead'
// }

// 通过链ID
getNetwork(1);
// {
//   chainId: 1,
//   ensAddress: '0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e',
//   name: 'homestead'
// }
```

## 1.3 **自定义的 ENS 合约**

这是因为通常需要在为Network 指定自定义属性来覆盖root ENS registry,或者是在一个通用的网络中拦截ENS 方法,又或者是在一个开发环境的网络中（绝大多数开发环境中的网络中的ENS合约需要你手动部署)指定ENS registry.

```js
const network = {
    name: "dev",
    chainId: 1337,
    ensAddress: customEnsAddress
};
```

## 1.4 **Provider**

在ethers里面,**Provider** 是访问区块链数据的只读抽象.

> 如果你用过Web3.js,以下是与ethers.js最大不同点:ethers library 在 **Provider**和**Signer**可执行的操作之间创建了一个强大的划分,而Web3.js将二者结合在一起.这种划分的关注点在于Provider严格的操作子集允许更多的后端、更一致的API,并确保其他库可以在不依赖任何潜在假设的情况下运行.

# 账户的方法

## 2.1 **getBalance**

*provider*.**getBalance**( *address* [ , *blockTag* = *latest* ] ) ⇒ *Promise< 大数(**BigNumber**) >*

返回给定的*blockTag* 区块高度下**地址的余额**

## 2.2 **getCode**

*provider*.**getCode**( *address* [ , *blockTag* = *latest* ] ) ⇒ *Promise< string< **DataHexString** > >*

返回给定的*blockTag* 区块高度下的**合约源代码**,如果当前没有合约被部署, 将返回 `0x`.

## 2.3 **getStorageAt**

*provider*.**getStorageAt**( *addr* , *pos* [ , *blockTag* = *latest* ] ) ⇒ *Promise< string< **DataHexString** > >*

在当前的*blockTag*下,输入address参数addr和position(位置)参数pos,返回类型为`Bytes32`的值.

## 2.4 **getTransactionCount**

*provider*.**getTransactionCount**( *address* [ , *blockTag* = *latest* ] ) ⇒ *Promise< **number** >*

在当前的*blockTag*下,返回已经发送的交易的数量,这个值是用来提供给**发送**到网络的下一个交易的**nonce值**.

# 区块的方法

## 3.1 **getBlock**

*provider*.**getBlock**( *block* ) ⇒ *Promise< **Block** >*

从网络中得到某个**区块信息**, `result.transactions` 是一串交易集合的哈希值

```js
await provider.getBlock(99)
{
  _difficulty: { BigNumber: "3849295379889" },
  difficulty: 3849295379889,
  extraData: '0x476574682f76312e302e312d39383130306634372f6c696e75782f676f312e34',
  gasLimit: { BigNumber: "3141592" },
  gasUsed: { BigNumber: "21000" },
  hash: '0xf93283571ae16dcecbe1816adc126954a739350cd1523a1559eabeae155fbb63',
  miner: '0x909755D480A27911cB7EeeB5edB918fae50883c0',
  nonce: '0x1a455280001cc3f8',
  number: 100004,
  parentHash: '0x73d88d376f6b4d232d70dc950d9515fad3b5aa241937e362fdbfd74d1c901781',
  timestamp: 1439799168,
  transactions: [
    '0x6f12399cc2cb42bed5b267899b08a847552e8c42a64f5eb128c1bcbd1974fb0c',
    '0x7f12399cc2cb42bed5b267899b08a847552e8c42a64f5eb128c1bcbd1974fb0c',
    ....
  ]
}
```

## 3.2 **getBlockWithTransactions**

*provider*.**getBlockWithTransactions**( *block* ) ⇒ *Promise< **BlockWithTransactions** >*

从网络中得到某个**区块信息**, `result.transactions` 是 一个**TransactionResponse** 对象的数组集合.

```js
await provider.getBlock(99)
{
  ...(参考上面)
  transactions: [
    {
        accessList: null,
        blockHash: '0xf93283571ae16dcecbe1816adc126954a739350cd1523a1559eabeae155fbb63',
        blockNumber: 100004,
        chainId: 0,
        confirmations: 16283837,
        creates: null,
        data: '0x',
        from: '0xcf00A85f3826941e7A25BFcF9Aac575d40410852',
        gasLimit: { BigNumber: "90000" },
        gasPrice: { BigNumber: "54588778004" },
        hash: '0x6f12399cc2cb42bed5b267899b08a847552e8c42a64f5eb128c1bcbd1974fb0c',
        nonce: 25,
        r: '0xb23adc880d3735e4389698dddc953fb02f1fa9b57e84d3510a2a4b3597ac2486',
        s: '0x4e856f95c4e2828933246fb4765a5bfd2ca5959840643bef0e80b4e3a243d064',
        to: '0xD9666150A9dA92d9108198a4072970805a8B3428',
        transactionIndex: 0,
        type: 0,
        v: 27,
        value: { BigNumber: "5000000000000000000" },
        wait: [Function]
    },
    ....
  ]
}
```

# **以太坊域名服务 (ENS) 方法**

**以太坊域名服务 (ENS)** 允许一个简短且易于记忆的ENS名称附加到任何一组键和值.

最常见的用法之一是使用一个简单的命名来**引用以太坊地址**.

## 4.1 **lookupAddress**

*provider*.**lookupAddress**( *address* ) ⇒ *Promise< **string** >*

使用*反向注册器*对**ENS中的地址进行反向查找**.如果名称不存在,或者正向查找不匹配,则返回`null`.

ENS名称需要额外的配置来设置反向记录,它们不会自动设置.

```js
> await provider.lookupAddress("0x5555763613a12D8F3e73be831DFf8598089d3dCa");
> 'ricmoo.eth'
```

## 4.2 **resolveName**

*provider*.**resolveName**( *name* ) ⇒ *Promise< string< **Address** > >*

**查找一个名称的地址**,如果这个*名称*没有被拥有,或者没有配置一个解析器,或者解析器没有配置一个地址,则返回`null`.

```js
> await provider.resolveName("ricmoo.eth");
> '0x5555763613a12D8F3e73be831DFf8598089d3dCa'
```

## 4.3 **getResolver**

*provider*.**getResolver**( *name* ) ⇒ *Promise< **EnsResolver** >*

**返回一个EnsResolver实例,该实例可用于进一步查询ENS命名的特定实体.**

## 4.4 EnsResolver

**Resolver 的名称**:*resolver*.**name** ⇒ *string*

**Resolver 的地址**:*resolver*.**address** ⇒ *string< 地址(**Address**) >*

### 4.4.1 **getAddress**

*resolver*.**getAddress**( [ *cointType* = *60* ] ) ⇒ *Promise< **string** >*

返回一个解析的类型为EIP-2304 *多币地址*的 Promise,默认情况下,**返回一个以太坊的地址(Address)**(`coinType = 60`).

### 4.4.2 **getContentHash**

*resolver*.**getContentHash**( ) ⇒ *Promise< **string** >*

返回一个Promise,解析为任何**存储的EIP-1577内容哈希**

```js
> await resolver.getContentHash();
> 'ipfs://QmdTPkMMBWQvL8t7yXogo7jq5pAcWg8J7RkLrDsWZHT82y'
```

### 4.4.3 **getText**

*resolver*.**getText**( *key* ) ⇒ *Promise< **string** >*

返回一个Promise,解析为任何**存储的EIP-634作为key的文本实体**.

```js
> await resolver.getText("email");
> 'me@ricmoo.com'

> await resolver.getText("url");
> 'https://www.ricmoo.com/'

> await resolver.getText("com.twitter");
> '@ricmoo'
```

# Logs方法

*provider*.**getLogs**( *filter* ) ⇒ *Promise**< Array< Log >** >*

返回**匹配筛选器的 Log数组**.

> 许多后端会丢弃旧的事件,并且请求范围太广可能也会被丢弃,因为它们需要太多的资源来执行查询.

# **网络状态方法**

## 6.1 getNetwork

*provider*.**getNetwork**( ) ⇒ *Promise< **Network** >*

返回这个 Provider 所**连接的 Network**.

## 6.2 getBlockNumber

*provider*.**getBlockNumber**( ) ⇒ *Promise< **number** >*

返回**最近挖出的区块的序号(或高度**).

## 6.3 getGasPrice

*provider*.**getGasPrice**( ) ⇒ *Promise< 大数(**BigNumber**) >*

返回一个 关于这个交易中的 **gas price的*最准预测***.

## 6.4 getFeeData

*provider*.**getFeeData**( ) ⇒ *Promise< **FeeDatba** >*

返回在一笔交易中当前的**建议 FeeData**.

对于 EIP-1559 的交易, 应该使用 `maxFeePerGas` 和 `maxPriorityFeePerGas`.

对于不支持 EIP-1559 中被遗留的交易和网络,应该使用 `gasPrice`.

```js
// 燃料价格 (单位: wei)...
feeData = await provider.getFeeData()
{
   gasPrice: { BigNumber: "16674318809" },
   lastBaseFeePerGas: { BigNumber: "16492712329" },
   maxFeePerGas: { BigNumber: "34485424658" },
   maxPriorityFeePerGas: { BigNumber: "1500000000" }
}

// 通常来说燃料的价格用 gwei 会更好理解
utils.formatUnits(feeData.maxFeePerGas, "gwei")
'34.485424658'
```

## 6.5 ready

*provider*.**ready** ⇒ *Promise< **Network** >*

返回一个 Promise ,直到网络建立就失效,忽略由于目标节点还未激活而出现的错误. 

这可以用于测试或附上脚本,以等待节点启动并顺利运行.

# 交易方法

## 7.1 call

*provider*.**call**( *transaction* [ , *blockTag* = *latest* ] ) ⇒ *Promise< string< **DataHexString** > >*

使用*call*返回**执行交易的结果**,调用不需要任何的以太,但不能改变任何状态.这在合约上调用getter方法是非常有用的.

```js
> await provider.call({
  // ENS public resovler address
  to: "0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41",

  // `function addr(namehash("ricmoo.eth")) view returns (address)`
  data: "0x3b3b57debf074faa138b72c65adbdcfb329847e4f2c04bde7f7dd7fcad5a52d2f395a558"
});
> '0x0000000000000000000000005555763613a12d8f3e73be831dff8598089d3dca'
```

## 7.2 estimateGas

*provider*.**estimateGas**( *transaction* ) ⇒ *Promise< 大数(**BigNumber**) >*

返回向网络提交**交易所需的预估gas值**.

> 估计的gas值可能不准确,因为网络上可能有另一个交易没有被计算在内,但在被挖出来之后就会影响相关状态.

## 7.3 getTransaction

*provider*.**getTransaction**( *hash* ) ⇒ *Promise< **TransactionResponse** >*

返回**带有哈希值的交易**,如果交易未知,则返回null.

如果一个交易还没有被挖出,这个方法将搜索交易池.各种后端可能有更多的限制交易池访问(例如,燃料价格太低或交易最近才发送,还没有索引),在这种情况下,这个方法也可能返回null.

## 7.4 getTransactionReceipt

*provider*.**getTransactionReceipt**( *hash* ) ⇒ *Promise< **TransactionReceipt** >*

返回**交易收据的哈希值**,如果交易还没有被挖出则返回null.

如果需要等待交易被挖出,请考虑下面的`waitForTransaction`方法.

## 7.5 sendTransaction

*provider*.**sendTransaction**( *transaction* ) ⇒ *Promise< **TransactionResponse** >*

向在挖区块的网络提交交易,**交易必须经过签名**,并且需要合法(例如:nonce值需要正确并且账户要有足够的余额来进行该笔交易的支付）

## 7.6 waitForTransaction

*provider*.**waitForTransaction**( *hash* [ , *confirms* = *1* [ , *timeout* ] ] ) ⇒ *Promise< **TxReceipt** >*

返回一个Promise,直到transactionHash被挖出来,才能够被解析.

如果*confirms*参数为0,这个方法是非阻塞的,如果交易还没有被挖出,则返回null. 

否则,该方法是阻塞的,直到这个包含该交易的、被*confirms*标识的区块被挖出.

# 事件方法

## eventName参数

**Log Filter**:一个过滤器是一个对象.表示合约log Filter它具有可选的`address` (合约地址) 和 `topics` (一个要匹配的主题集).

如果参数`address`未被指定.则过滤器匹配任何合约地址.

**Topic-Set Filter**:**主题集过滤器**的值是一个主题集数组.

此事件与日志过滤器相同.但参数`address`可以不用填(匹配任何合约).

**Transaction Filter**:**交易过滤器**的值是交易的哈希值.

这个事件会在任何给定的交易挖出的链上区块中触发. 使用once方法比使用on方法要普遍地多.

**除了交易和过滤事件外.还有几个命名事件:**

| **EventName** | **Arguments**           | **Description**                                              |
| ------------- | ----------------------- | ------------------------------------------------------------ |
| `"block"`     | *blockNumber*           | 当一个区块被挖出时触发                                       |
| `"error"`     | *error*                 | 只要有错误就触发                                             |
| `"pending"`   | *pendingTransaction*    | 当一个新交易进入内存池时触发;只有特定的providers提供此事件.从而在运行在自己的节点上获得可靠的数据 |
| `"willPoll"`  | *pollId*                | 在一个polling loop开始之前触发;*(大多数开发者很少使用)*      |
| `"poll"`      | *pollId*, *blockNumber* | 在每个poll cycle中.`blockNumber`更新之后(如果改变了).以及与在poll loop中任何其他的事件(如果有)之前触发; *(大多数开发者很少使用)* |
| `"didPoll"`   | *pollId*                | 在polling loop中的所有事件被触发后触发;(*大多数开发者很少使用*) |
| `"debug"`     | provider dependent      | 每个Provider可以使用它来发出有用的调试信息.格式由开发者决定;*(大多数开发者很少使用)* *(very rarely used by most developers)* |

## 8.1 on

*provider*.**on**( *eventName* , *listener* ) ⇒ ***this***

为**每一个参数为*eventName*的 事件添加*监听器***.

## 8.2 once

*provider*.**once**( *eventName* , *listener* ) ⇒ ***this***

为参数**为*eventName*的 事件添加*监听器*,监听使用过后将会被移除**.

## 8.3 emit

*provider*.**emit**( *eventName* , ...*args* ) ⇒ ***boolean***

**通知所有的*eventName* event监听器,并把参数传递给它们,这通常只在内部使用.**

## 8.4 off

*provider*.**off**( *eventName* [ , *listener* ] ) ⇒ ***this***

**移除一个参数为*eventName*的事件*监听器***,如果没有提供*listener*参数,则移除所有关于*eventName*的监听器.

## **8.5** removeAllListeners

*provider*.**removeAllListeners**( [ *eventName* ] ) ⇒ ***this***

**移除所有参数为*eventName*的事件监听器**,如果没有提供*eventName*参数,则移除**所有**事件.

## 8.6 listenerCount

*provider*.**listenerCount**( [ *eventName* ] ) ⇒ ***number***

**返回所有参数为*eventName*事件的监听器数量**,如果没有提供*eventName*参数,返回所有监听器的数量.

## 8.7 listeners

*provider*.**listeners**( *eventName* ) ⇒ *Array< **Listener** >*

**返回参数为*eventName*事件监听器的list集合.**

```js
provider.on("block", (blockNumber) => {
    // 只有有一个区块改变就会触发
})


provider.once(txHash, (transaction) => {
    // 当含有交易的区块被挖出就会触发
})


// 这个过滤器也可以通过Contract或接口API生成。
// 如果不指定address，则默认匹配任何地址，
// 如果不指定topics，则默认匹配任何日志。
filter = {
    address: "dai.tokens.ethers.eth",
    topics: [
        utils.id("Transfer(address,address,uint256)")
    ]
}
provider.on(filter, (log, event) => {
    // 当 DAI token 发生转账时触发
})


// 注意，这是一个主题集数组，与使用没有address参数的过滤器相同(即匹配任何地址)
topicSets = [
    utils.id("Transfer(address,address,uint256)"),
    null,
    [
        hexZeroPad(myAddress, 32),
        hexZeroPad(myOtherAddress, 32)
    ]
]
provider.on(topicSets, (log, event) => {
    // 当任何一个token 发送到我的或者我的其他的地址时触发
})


provider.on("pending", (tx) => {
    // 当有一笔pending交易被捕获时触发
});


provider.on("error", (tx) => {
    // 只要有错误就会触发
});
```

# 校验方法

*Provider*.**isProvider**( *object* ) ⇒ ***boolean***

当且仅**当参数*object*是Provider返回 true .**

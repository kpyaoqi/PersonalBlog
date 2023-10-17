---
title: 4-Types

date: 2022-08-09	

categories: ether@5.4.js	

tags: [区块链,ether@5.4.js]
---	

# BlockTag

**BlockTag** 在区块链中指定了一个特定的区块位置

- **`"latest"`** - 最新挖出的区块
- **`"earliest"`** - 创世区块
- **`"pending"`** - 目前正在挖的区块;不是所有的操作和后端都支持这个BlockTag属性
- **`"number"`**- 区块高度
- **`"a negative number"`** - 几个区块前的区块
- **`"hex string"`** - 区块高度(十六进制表示)

# Networkish

**Networkish**可以是以下任意一种:

- 一个Network 对象

  *network*.**name** ⇒ *string*，可读的网络名称，如`homestead`。如果网络的name未指定，将会是`"unknown"`。

  *network*.**chainId** ⇒ *number*，网络的Chain ID。

  *network*.**ensAddress** ⇒ *string< 地址(Address) >*，部署在这个网络上的ENS注册表的地址。

- 字符串形式的常见网络 (如`"homestead"`)

- number类型的网络chain ID， 如果chain ID 是常见网络的一种,`name` 和 `ensAddress` 会自动填充, 否则 name 会是`"unknown"` 且不使用任何的 `ensAddress`。

# FeeData

**FeeData**对象根据最好且可用建议，去封装发送交易所需的费用数据。

***feeData*.gasPrice ⇒ *大数(BigNumber):***gasPrice用于那些不支持EIP-1559的遗留交易或网络。

***feeData*.maxFeePerGas ⇒ *大数(BigNumber)*:**`maxFeePerGas`用于交易。这基于最近产出区块的`baseFee`。

***feeData*.maxPriorityFeePerGas ⇒ *大数(BigNumber)*:**用于交易的“最大优先的Gas”

# Block

***block*.hash ⇒ *string< DataHexString< 32 > >*:**这个区块的哈希。

***block*.parentHash ⇒ *string< DataHexString< 32 > >*:**上个区块的哈希。

***block*.number ⇒ *number*:**当前区块的高度（数量）。

***block*.timestamp ⇒ numbe:**r当前区块的时间戳。

***block*.nonce ⇒ *string< DataHexString >*nonce:** 用在基于工作量证明（PoW）机制挖取区块

***block*.difficulty ⇒ *number:***该区块的矿工需要达到的难度指标。 开发人员通常对这个属性不感兴趣。

***block*.gasLimit ⇒ *大数(BigNumber):***这个区块允许使用的最大gas数量,该值可以由矿工投票决定要变大还是变小

***block*.gasUsed ⇒ *大数(BigNumber):***当前区块所有的交易已使用的gas总数量。

***block*.miner ⇒ *string:***当前区块的coinbase 地址，这个地址表示挖出这个区块的矿工想要接受这笔开采奖励的地址。

***block*.extraData ⇒ *string:***该数据表示的是当挖出一个区块时，这个矿工可以选择包含的额外数据。 

***block*.transactions ⇒ *Array< string< DataHexString< 32 > > >*:**一个当前区块包含的每一条交易哈希的列表。

***block*.transactions ⇒ *Array< TransactionResponse >:***当前区块包含的交易列表。

# Events and Logs

## **5.1 EventFilter**

***filter*.address ⇒ *string< 地址(Address) >*:**想要筛选的地址， `null` 表示匹配任何用户的地址。

***filter*.topics ⇒ *Array< string< Data< 32 > > | Array< string< Data< 32 > > > >*:**要筛选的主题，`null`示匹配任何用户的主题。

每一个 entry 表示一个**AND**匹配条件，entry 也可以是`null`来匹配所有的内容。 如果给定的 entry 是一个数组，那么这个 entry 是被视为 **OR**匹配条件来匹配entry里面的内容。 有关指定复杂过滤器的详细信息和示例，请参阅Filters。

## 5.2 Filter

***filter*.fromBlock ⇒ *BlockTag*:**用于搜索匹配过滤条件的日志的起始区块(包含在内)。

***filter*.toBlock ⇒ *BlockTag*:**用于搜索匹配过滤条件的日志的结束区块(包含在内)。

## 5.3 FilterByBlockHash

***filter*.blockHash ⇒ *string< DataHexString< 32 > >:***指定区块(按这个区块哈希值)搜索符合过滤条件的日志。

## 5.4 Log

***log*.blockNumber ⇒ *number:***包含该日志的交易的区块高度(区块数)。

***log*.blockHash ⇒ *string< DataHexString< 32 > >*:**包含该日志的交易的区块的哈希值。

***log*.removed ⇒ *boolean*:**在区块重组期间，如果一个交易是孤立的，它将被设置为true，表示Log entry已被删除; 在不久的将来，当使用包含该日志的交易被另一个区块挖出时，可能会再次触发此日志，但请记住，值可能会更改。

***log*.transactionLogIndex ⇒ *number*:**该日志在交易中的索引。

***log*.address ⇒ *string< 地址(Address) >*:**生成这条日志的合约地址。

***log*.data ⇒ *string< DataHexString >*:**日志的数据。

***log*.topics ⇒ *Array< string< DataHexString< 32 > > >*:**日志的主题(索引属性)列表。

***log*.transactionHash ⇒ *string< DataHexString< 32 > >*:**包含这条日志的交易的哈希值。

***log*.transactionIndex ⇒ *number*:**区块中包含这条日志的交易的索引。

***log*.logIndex ⇒ *number*:**在整个**区块中**所有的日志集合里，该条日志的索引。

# Transactions

## 6.1 TransactionRequest

描述了一笔将要被发送到网络或以其他方式处理的交易，所有字段都是可选的，并且可以是一个能被解析为所需类型的promise。

***transactionRequest*.to ⇒ *string | Promise< string >*:**接受这笔交易的地址(or ENS name)。

***transactionRequest*.from ⇒ *string< 地址(Address) > | Promise< string< 地址(Address) > >*:**发送这笔交易的地址。

***transactionRequest*.nonce ⇒ *number | Promise< number >:***交易的 nonce 值。 **发送**这笔交易要设置这个值，是number类型的。

***transactionRequest*.data ⇒ *DataHexString | Promise< DataHexString >*:**交易的数据。

***transactionRequest*.value ⇒ *大数(BigNumber) | Promise< 大数(BigNumber) >*:**这笔交易要发送的数量(以wei为单位)。

***transactionRequest*.gasLimit ⇒ *大数(BigNumber) | Promise< 大数(BigNumber) >*:**这笔交易允许使用的最大gas值。

如果未指定，ethers将使用`estimateGas`来确定要使用的gas值。对于无法预测的gas的交易，可能需要这样做显式地指定。

***transactionRequest*.gasPrice ⇒ *大数(BigNumber) | Promise< 大数(BigNumber) >*:**这笔交易将支付的每个gas的价格(以wei为单位)。

这个属性不适用于这笔交易的`type`设置为`1` 或者 `2`的情形，也不适用于交易已指定了`maxFeePerGas` 或者 `maxPriorityFeePerGas`的情形。

***transactionRequest*.maxFeePerGas ⇒ *大数(BigNumber) | Promise< 大数(BigNumber) >:***这笔交易将支付EIP-1559基础费用的每个gas的价格的最高价格(以wei为单位)。

大多数开发人员不应该指定该参数，而应该使用网络决定的默认值。 这个属性不适用于这笔交易的`type`设置为`0`的情形，也不适用于交易已指定了`gasPrice`的情形。

***transactionRequest*.maxPriorityFeePerGas ⇒ *大数(BigNumber) | Promise< 大数(BigNumber) >*:**这笔交易将支付EIP-1559优先费用的每个gas的价格的价格(以wei为单位)。 这包含在`maxFeePerGass`中，所以这**不会影响**`maxFeePerGas`设置的总最大价格。

大多数开发人员应该不指定该参数，而应该使用网络决定的默认值。 这个属性不适用于这笔交易的`type`设置为`0`的情形，也不适用于交易已指定了`gasPrice`的情形。

***transactionRequest*.chainId ⇒ *number | Promise< number >*:**已被授权交易的chain ID，由EIP-155指定。 如果chain ID为0将禁用EIP-155，交易将在任何网络上有效。这可能**很危险**，应该小心，因为它允许交易在可能不是你想指定的网络上复现。 在最新版本的Geth中，有意重现的交易在默认情况下是禁用的，需要配置才能启用。

***transactionRequest*.type ⇒ *null | number*:** 这笔交易envelope的EIP-2718类型，在网络上这个默认值是`null`。要强制使用没有envelope的lagacy交易，请使用type `0`。

***transactionRequest*.accessList ⇒ *AccessListish*:**要包含的AccessList;仅适用于EIP-2930和EIP-1559的交易。

## 6.2 TransactionResponse

**TransactionResponse**包括交易的所有属性，以及等区块被挖出来后就很有用的几个属性。

***transaction*.blockNumber ⇒ *number*:**这笔交易所在区块被挖出来时的高度(区块数量)如果区块还没有被挖出，这个值为`null`。

***transaction*.blockHash ⇒ *string< DataHexString< 32 > >:***交易所在区块挖出来时的哈希值,如果区块还没有被挖出，这个值为`null`。

***transaction*.timestamp ⇒ *number*:**这笔交易所在区块被挖出来时的时间戳。如果区块还没有被挖出，这个值为`null`。

***transaction*.confirmations ⇒ *number:***当这笔交易所在的区块被挖出来时，所有已挖出的区块数量(包括初始块)。

***transaction*.raw ⇒ *string< DataHexString >:***序列化交易。如果一些后端不重新填充它时可能是null。 如果你需要这个属性，可以使用this cookbook recipe从**TransactionResponse**对象中计算。

***transaction*.wait( [ *confirms* = *1* ] ) ⇒ *Promise< TransactionReceipt >*:**confirms*属性表示你愿意等待的挖出区块数量，一旦满足这个值将返回一个待解析的TransactionReceipt。 如果*confirms为0，并且交易还没有被挖出，则返回`null`。

如果交易执行失败(即接收状态为`0`)，将会生成一个**CALL_EXCEPTION**，其属性如下:

- `error.transaction` - 这笔交易
- `error.transactionHash` - 这笔交易的哈希值
  - `error.receipt` - the actual receipt, 状态是 `0`

如果这笔交易被另一个交易替换，将会生成一个**TRANSACTION_REPLACED**，其属性如下:

- `error.hash` - 被替换的交易的哈希值
- `error.reason` - string类型的错误原因; 是 `"repriced"`或 `"cancelled"`或 `"replaced"`其中一个
- `error.cancelled` - boolean类型; `"repriced"`表示不视为cancelled, 但`"cancelled"` 和`"replaced"`视为cancelled。
- `error.replacement` - the replacement transaction (a **TransactionResponse**)
- `error.receipt` - replacement transaction的收据 (a **TransactionReceipt**)

当用户在客户端相同的帐户中发送一笔新的交易，这笔交易与之前一笔交易的`nonce`相同时，之前的那笔交易将被替换。 这通常是为了加快交易或取消交易，新交易可以给矿工更多的费用，让他们更喜欢新交易，而不是原来的交易。

***transactionRequest*.type ⇒ *number*:**这笔交易的EIP-2718类型。如果事务是没有envelope的legacy交易，则其类型为`0`。

***transactionRequest*.accessList ⇒ *AccessList*:**包含的**AccessList**，对于不支持访问列表的交易类型时该值为空。

## 6.3 TransactionReceipt

***receipt*.to ⇒ *string< 地址(Address) >*:**接受这笔交易的地址。如果交易是用于部署合约，则为`null`。

***receipt*.from ⇒ *string< 地址(Address) >*:**发起这笔交易的地址。

***receipt*.contractAddress ⇒ *string< 地址(Address) >:***如果这笔交易是转去一个`null`的地址，它是表示一个**init transaction**用于部署合约。在这种情况下，这个属性就代表所创建的合约的地址。 要计算合约地址，getContractAddress的utility function还可以与TransactionResponse对象一起使用，该对象需要交易的nonce值和发送人的地址。

***receipt*.transactionIndex ⇒ *number*:**当前区块包含所有交易列表中，此交易的索引。

***receipt*.type ⇒ *number:***此交易的EIP-2718类型。如果事务是没有envelope的legacy事务，则其type 为`0`。

***receipt*.root ⇒ *string*:**收据的中间状态根

只有在Byzantium Hard Fork之前包含在区块中的交易才有这个属性，因为它被`status`属性取代了。

这些属性通常对开发者没什么用处。仅考虑单笔交易的防伪性的时候可以用于验证状态转换;如果没有它，则必须考虑整个区块。

***receipt*.gasUsed ⇒ *大数(BigNumber)*:**该交易实际使用的gas值。

***receipt*.logsBloom ⇒ *string< DataHexString >:***一个bloom-filter，包含此交易中所有日志中包含的全部地址和主题。

***receipt*.blockHash ⇒ *string< DataHexString< 32 > >:***包含该交易的区块的区块哈希值。

***receipt*.transactionHash ⇒ *string< DataHexString< 32 > >:***这笔交易的哈希值。

***receipt*.logs ⇒ *Array< Log >:***这笔交易触发的所有日志。

***receipt*.blockNumber ⇒ *number:***包含这笔交易的区块的高度（区块数量）。

***receipt*.confirmations ⇒ *number:***当这笔交易所在的区块被挖出来时，所有已挖出的区块数量(包括这笔刚挖出的区块)。

***receipt*.cumulativeGasUsed ⇒ 大数(BigNumber):**对包含该交易的块，这是截至(该交易的有序交易列表中每个交易使用的gas的总和

***receipt*.byzantium ⇒ *boolean*:**如果区块是在post-Byzantium Hard Fork区块，这个值是true。

***receipt*.status ⇒ *boolean:***如果交易成功，这个值为1；如果交易被reverted，这个值为为0。 只有post-Byzantium Hard Fork中包含的交易才有这个属性。



# Access Lists

Access List是可选的，它包含一个地址列表和地址的存储槽，这些地址应该是*warmed*或pre-fetched的，以便在这个交易的执行中使用。 一个*warmed*值有一个额外的预先支付访问的成本，但在整个读取和写入代码的执行过程中会被降低。

## 7.1 AccessListish

AccessList的一个更宽松的描述，它将在内部使用accessListify进行转换。

他可以是以下的任意一种:

- 任意的AccessList
- 一个包含两个元素的数组，第一个元素是地址，第二个是数组在storage中的key
- 一个对象, 所以的key表示地址，相应的value表示数组在storage中的key

当使用对象形式(上述最后一个选项)时，地址和存储槽将被排序。 如果需要access list的显式顺序，则必须使用方式。 大多数开发人员不需要明确的顺序。

```js
// Option 1:
// AccessList
// see below

// Option 2:
// Array< [ Address, Array<Bytes32> ] >
accessList = [
  [
    "0x0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
    [
      "0x0000000000000000000000000000000000000000000000000000000000000004",
      "0x0bcad17ecf260d6506c6b97768bdc2acfb6694445d27ffd3f9c1cfbee4a9bd6d"
    ]
  ],
  [
    "0x5FfC014343cd971B7eb70732021E26C35B744cc4",
    [
        "0x0000000000000000000000000000000000000000000000000000000000000001"
    ]
  ]
];

// Option 3:
// Record<Address, Array<Bytes32>>
accessList = {
  "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e": [
    "0x0000000000000000000000000000000000000000000000000000000000000004",
    "0x0bcad17ecf260d6506c6b97768bdc2acfb6694445d27ffd3f9c1cfbee4a9bd6d"
  ],
  "0x5FfC014343cd971B7eb70732021E26C35B744cc4": [
    "0x0000000000000000000000000000000000000000000000000000000000000001"
  ]
}equivalent to the AccessList example below

// Option 1:
// AccessList
// see below

// Option 2:
// Array< [ Address, Array<Bytes32> ] >
accessList = [
  [
    "0x0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
    [
      "0x0000000000000000000000000000000000000000000000000000000000000004",
      "0x0bcad17ecf260d6506c6b97768bdc2acfb6694445d27ffd3f9c1cfbee4a9bd6d"
    ]
  ],
  [
    "0x5FfC014343cd971B7eb70732021E26C35B744cc4",
    [
        "0x0000000000000000000000000000000000000000000000000000000000000001"
    ]
  ]
];

// Option 3:
// Record<Address, Array<Bytes32>>
accessList = {
  "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e": [
    "0x0000000000000000000000000000000000000000000000000000000000000004",
    "0x0bcad17ecf260d6506c6b97768bdc2acfb6694445d27ffd3f9c1cfbee4a9bd6d"
  ],
  "0x5FfC014343cd971B7eb70732021E26C35B744cc4": [
    "0x0000000000000000000000000000000000000000000000000000000000000001"
  ]
};
```

## 7.2 AccessList

一个EIP-2930交易允许一个可选的**AccessList**，它会导致交易*warm*（预缓存）另一个地址状态和指定的storage key。

这会使得交易的内在成本变大，但在整个事务执行过程中给存储和状态访问带来了好处。

```js
// Array of objects with the form:
// {
//   address: Address,
//   storageKey: Array< DataHexString< 32 > >
// }

accessList = [
  {
    address: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
    storageKeys: [
        "0x0000000000000000000000000000000000000000000000000000000000000004",
        "0x0bcad17ecf260d6506c6b97768bdc2acfb6694445d27ffd3f9c1cfbee4a9bd6d"
    ]
  },
  {
    address: "0x5FfC014343cd971B7eb70732021E26C35B744cc4",
    storageKeys: [
        "0x0000000000000000000000000000000000000000000000000000000000000001"
    ]
  }
];
```


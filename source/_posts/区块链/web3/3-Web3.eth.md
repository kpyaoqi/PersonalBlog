---
title: 3-Web3.eth

date: 2022-08-31	

categories: web3.js	

tags: [区块链,web3.js]
---	

**`web3-eth` 包用来与以太坊区块链和以太坊智能合约进行交互**

## 1.defaultAccount

**(默认地址)**

如果下面这些方法没有指定 `"from"` 属性，则使用该地址作为默认的 `"from"` 属性值。

- web3.eth.sendTransaction()
- web3.eth.call()
- new web3.eth.Contract() -> myContract.methods.myMethod().call()
- new web3.eth.Contract() -> myContract.methods.myMethod().send()

**属性**：`String` - 20 字节: 任意以太坊地址。 该地址对应的私钥应该保存在你的以太坊节点或 keystore 文件中， 默认值为 `undefined`

## 2.defaultBlock

**一些特定方法所使用的默认区块号( 默认值为 latest)**

- web3.eth.getBalance()
- web3.eth.getCode()
- web3.eth.getTransactionCount()

**属性**：`Number|BN|BigNumber`: 指定区块号、`"latest"` - `String`: 最新区块 (区块链的头号区块)、

​		   `"pending"` - `String`: 正要挖到的区块 (包括待处理交易)、`"earliest"` - `String`: 创世区块

## 3.defaultHardfork

**(使用的硬分叉属性)**

**属性(String)**:"chainstart"、"homestead"、"dao"、"tangerineWhistle"、"spuriousDragon"、"byzantium"、"constantinople"、

​					  "petersburg"、"istanbul"（默认值为 "petersburg"）

## 4.defaultChain

**(签名交易时所用的默认链)**

**属性(String):**"mainnet"、"goerli"、"kovan"、"rinkeby"、"ropsten"（默认值为 "mainnet"）

## 5.defaultCommon

**（签名交易时所使用的默认通用属性）**

**属性**（默认值为 undefined）：

- `customChain` - `Object`: 自定义链属性

  - `name` - `string`: (可选) 链名称

  - `networkId` - `number`: 自定义链的网络 ID

  - `chainId` - `number`: 自定义链的链 ID

- `baseChain` - `string`: (可选) `mainnet`, `goerli`, `kovan`, `rinkeby`, or `ropsten`

- `hardfork` - `string`: (可选) `chainstart`, `homestead`, `dao`, `tangerineWhistle`, `spuriousDragon`, `byzantium`, `constantinople`, `petersburg`, or `istanbul`

## 6.transactionBlockTimeout

**(直到第一次确认发生交易应该等待的区块数)**

**返回值**(number)：transactionBlockTimeout 当前值 (默认值: 50)

> 用在基于套接字的连接上

## 7.transactionConfirmationBlocks

**(一笔交易被认为已确认所需要的区块数)**

**返回值**(number):transactionConfirmationBlocks 当前值 (默认值: 24)

> 用在基于套接字的连接上

## 8.blockHeaderTimeout

**(返回轮询获取交易收据前等待“newBlockHeaders”事件的秒数)**

**返回值**(number):blockHeaderTimeout的当前值(默认值:10秒)

## 9.transactionPollingTimeout

**(Web3 等待网络挖出交易的确认收据的秒数)**

**返回值**(number):transactionPollingTimeout 当前值 (默认值: 750)

> 用在基于 HTTP 的连接上使用，当超时发生时，交易可能仍未完成

## 10.transactionPollingInterval

**(Web3调用接收确认网络挖掘交易的间隔秒数)**

**返回值**(number):transactionPollingInterval的当前值(默认值:1000毫秒)

> 用在基于 HTTP 的连接上使用

## 11.handleRevert

**(在调用下面这些方法时启用，将返回回退原因字符串)**

- web3.eth.call()
- web3.eth.sendTransaction()
- contract.methods.myMethod(…).send(…)
- contract.methods.myMethod(…).call(…)

**返回值**(boolean):handleRevert 当前值 (默认值: false)

> 回退原因字符串和签名会作为返回错误的属性存在。

## 12.maxListenersWarningThreshold

**(支持套接字订阅的提供程序的事件侦听器的数量)**

**返回值**(number)：maxListenersWarningThreshold的当前值(默认值:100)

## 13.getProtocolVersion

**(返回节点的以太坊协议版本)**

**返回值(Promise)**;返回协议版本

## 14.isSyncing

**(检测当前节点是否正在进行数据同步)**

**返回值(Promise)/"false"**:

- `startingBlock` - `Number`: 同步开启时的区块号。
- `currentBlock` - `Number`: 节点已经同步到的区块号。
- `highestBlock` - `Number`: 预计会同步到的区块号。
- `knownStates` - `Number`: 预计要下载的状态数据。
- `pulledStates` - `Number`: 已经下载的状态数据

## 15.getCoinbase

**(返回用来收取挖矿奖励的 coinbase 地址)**

**返回值(Promise)**：String - 20 个字节，在节点中设置的用来挖矿的 coinbase 地址。

## 16.isMining

**(检测节点是否正在挖矿)**

**返回值(Promise)**：Boolean,如果节点正在挖矿返回 true ，否则返回 false

## **17.getHashrate**

**返回值(Promise)**：Number,返回节点每秒所挖的哈希数。

## 18.getGasPrice

返回当前 gas 价格预言机， gas 价格由最后几个区块 gas 价格的中位数确定。

**返回值(Promise)**：String,当前以 wei 为单位的 gas 价格

## 19.getFeeHistory

返回所请求区块范围的基本gas费用和每笔交易有效优先级费用(如果可用)，对于前EIP-1559区块，天然气价格作为奖励返回，零作为每种天然气的基本费用返回。

**参数：**

- `String|Number|BN|BigNumber`-请求范围内的块数，在单个查询中可以请求1到1024个块。如果不是所有块都可用，则可能会返回少于请求的数量。
- `String|Number|BN|BigNumber`-所请求范围的最高数字块。
- `Array of numbers`-从每个区块每个gas的有效优先费用中取样的百分位值的单调递增列表，按使用的天然气进行加权。示例:[“0”, “25”, “50”, “75”, “100”]或者[“0”, “0.5”, “1”, “1.5”, “3”, “80”]

**返回值(Promise)：**Object,返回的区块范围的费用历史对象:

- `Number`oldestBlock -返回范围中编号最小的块。
- `Array of strings`baseFeePerGas——每个区块gas的批量基本费用的数组。这包括返回范围中最新块之后的下一个块，因为该值可以从最新块中导出。对于EIP1559年以前的块，返回零。
- `Array of numbers`gasUsedRatio -区块gas使用比率的数组，这些被计算为gas使用和gas限制的比率。
- `Array of string arrays`奖励-来自单个区块的每个gas数据点的有效优先费用的阵列。如果块为空，则返回所有零。

## 20.getAccounts

**返回值(Promise)：**Array,返回节点所控制的账户列表

## 21.getBlockNumber

**返回值(Promise)：**Number,返回当前区块号

## 22.getBalance

**参数：**

1. `String` - 要获取余额的地址。
2. `Number|String|BN|BigNumber` - (可选) 如果传入该参数，则会覆盖通过 web3.eth.defaultBlock 设置的默认区块。 也可以使用像 `"latest"`, `"earliest"`, `"pending"`, and `"genesis"` 这种预先设置的区块号。
2. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**string,返回给定地址的当前账户余额，以 wei 为单位

## 23.getStorageAt

**参数：**

1. `String` - 用来获取存储值的地址。
2. `Number|String|BN|BigNumber` - 存储的索引位置。
3. `Number|String|BN|BigNumber` - (可选) 如果传入值则会覆盖通过 web3.eth.defaultBlock 设置的默认区块号。预定义的区块号可以使用 `"latest"`, `"earliest"` `"pending"`, 和 `"genesis"` 等值。
3. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**string,返回地址在某个特定位置的存储值

## 24.getCode

**参数：**

1. `String` - 获取代码所需要的地址。
2. `Number|String|BN|BigNumber` - (可选) 如果传入值则会覆盖通过 web3.eth.defaultBlock 设置的默认区块号。预定义的区块号 `"latest"`, `"earliest"` `"pending"`, 和 `"genesis"` 等也可以使用。
2. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**string,返回给定地址的代码数据

## 25.getBlock

**参数：**

1. `String|Number|BN|BigNumber` - 区块号或区块哈希。 或者是 `"genesis"`、 `"latest"`、 `"earliest"` 、 `"pending"` 等在 默认区块参数 中定义的字符串。
2. `Boolean` - (可选, 默认值 `false`) 如果设定为 `true`，返回的区块会包含完整的交易对象。 默认值为 `false` ，没必要特别设置为 false 了。 如果设置为 false，则返回的区块中仅包含交易哈希值。
2. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Object,区块对象:

- `number` - `Number`: 区块号。 打包中的区块其值为 `null`。
- `hash` 32 字节 - `String`: Hash of the block. 打包中的区块其值为 `null`。
- `parentHash` 32 字节 - `String`: 父区块哈希值。
- `baseFeePerGas` - `Number`: 在网络上发送交易的最低收费.
- `nonce` 8 Bytes - `String`: 所生成的工作量证明哈希值。打包中的区块其值为 `null`。
- `sha3Uncles` 32 字节 - `String`: 区块中叔块数据的 SHA3 哈希值。
- `logsBloom` 256 Bytes - `String`: 区块日志的布隆过滤器。打包中的区块其值为 `null`。
- `transactionsRoot` 32 字节 - `String`: 区块中交易 trie 树的根哈希。
- `stateRoot` 32 字节 - `String`: 区块中状态 trie 树的根哈希。
- `miner` - `String`: 获得挖矿奖励的受益人地址。
- `difficulty` - `String`: 该区块难度值。
- `totalDifficulty` - `String`: 到此区块为止链的总难度值。
- `extraData` - `String`: 区块补充数据字段。
- `size` - `Number`: 该区块的字节数大小。
- `gasLimit` - `Number`: 该区块允许的最大 gas 消耗量。
- `gasUsed` - `Number`: 该区块所有交易所消耗的 gas 总量。
- `timestamp` - `Number`: 区块生成时的时间戳。
- `transactions` - `Array`: 交易对象数组, 或者基于 `returnTransactionObjects` 参数的 32 字节交易哈希。
- `uncles` - `Array`: 叔块哈希数组。

## 26.getBlockTransactionCount

**参数：**

- `String|Number|BN|BigNumber` - 区块号或区块哈希。 或者像 `"genesis"`, `"latest"`, `"earliest"`, or `"pending"` 这些在 默认区块参数中指定的字符串。

- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Number, 返回区块中包含的交易数量。

## 27.getBlockUncleCount

**参数：**

- `String|Number|BN|BigNumber` - 区块号或区块哈希。 或者像 `"genesis"`, `"latest"`, `"earliest"`, or `"pending"` 这些在 默认区块参数中指定的字符串。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Number, 返回区块中包含的叔块数量。

## 28.getUncle

**参数：**

- `String|Number|BN|BigNumber` - 区块号或区块哈希。 或者像 `"genesis"`, `"latest"`, `"earliest"`, or `"pending"` 这些在 默认区块参数 中指定的字符串。
- `Number` - 叔块位置索引。
- `Boolean` - (可选, 默认值 `false`) 如果设定为 `true`，返回的区块会包含完整的交易对象。 默认值为 `false` ，没必要特别设置为 false 了。 如果设置为 false，则返回的区块中仅包含交易哈希值。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Object, 返回的叔块对象。 具体返回值可查看 25.web3.eth.getBlock()。。

> 叔块不包含任何交易。

## 29.getTransaction

**参数:**

- `String` - 交易哈希。

- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Object, 返回的交易对象：

- `hash` 32 字节 - `String`: 交易哈希。
- `nonce` - `Number`: 发送人在此之前进行的交易次数。
- `blockHash` 32 字节 - `String`: 该交易所在区块的区块哈希。 打包中的区块其值为 `null`。
- `blockNumber` - `Number`: 该交易所在区块的区块号。 打包中的区块其值为 `null`。
- `transactionIndex` - `Number`: 该交易在区块中的位置索引。 打包中的区块其值为 `null`。
- `from` - `String`: 交易发送人地址。
- `to` - `String`: 交易接收人地址。 对于合约创建交易其值为``null``
- `value` - `String`: 转账金额。以 wei 为单位。
- `gasPrice` - `String`: 由发送人指定的 gas 价格。以 wei 为单位。
- `gas` - `Number`: 由发送人指定的 gas 数量。
- `input` - `String`: 伴随交易发送的数据。

## 30.getPendingTransactions

**参数：**`Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise<Object[]>)：** 等待打包的交易数组:

- `hash` 32 字节 - `String`: 交易哈希。
- `nonce` - `Number`: 发送人在此之前进行的交易次数。
- `blockHash` 32 字节 - `String`: 该交易所在区块的区块哈希。 打包中的区块其值为 `null`。
- `blockNumber` - `Number`: 该交易所在区块的区块号。 打包中的区块其值为 `null`。
- `transactionIndex` - `Number`: 该交易在区块中的位置索引。 打包中的区块其值为 `null`。
- `from` - `String`: 交易发送人地址。
- `to` - `String`: 交易接收人地址。 对于合约创建交易其值为``null``
- `value` - `String`: 转账金额。以 wei 为单位。
- `gasPrice` - `String`: 由发送人指定的 gas 价格。以 wei 为单位。
- `gas` - `Number`: 由发送人指定的 gas 数量。
- `input` - `String`: 伴随交易发送的数据。

## 31.getTransactionFromBlock

**参数:**

1. `String|Number|BN|BigNumber` - 区块号或区块哈希。 或者像 `"genesis"`, `"latest"`, `"earliest"`, or `"pending"` 这些在 默认区块参数 中指定的字符串。
2. `Number` - 交易位置索引。
2. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise):**Object,交易对象, 详情请查看 29.web3.eth.getTransaction()

## 32.getTransactionReceipt

**参数:**

- `String` - 交易哈希

- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Object, 返回的交易收据对象,未能找到交易收据时返回 `null` :

- `status` - `Boolean`: 若交易成功其值为 `TRUE` ， 交易被 EVM 回退其值为 `FALSE`。
- `blockHash` 32 字节 - `String`: 该交易所在区块的区块哈希。
- `blockNumber` - `Number`: 该交易所在区块的区块号。
- `transactionHash` 32 字节 - `String`: 交易哈希。
- `transactionIndex`- `Number`: 该交易在区块中的位置索引。
- `from` - `String`: 交易发送人地址。
- `to` - `String`: 交易接收人地址。 对于合约创建交易其值为``null``
- `contractAddress` - `String`: 若指定交易为合约创建交易，其值为新创合约地址，否则其值为 `null`。
- `cumulativeGasUsed` - `Number`: 该交易执行时其所在区块已经累计消耗的 gas 量。
- `gasUsed`- `Number`: 该交易本身所消耗的 gas 量。
- `logs` - `Array`: 该交易所生成的日志对象。

- `effectiveGasPrice` - `Number`(或者`hex String`):从发送人账户中扣除的gas的实际价值。

## 33.getTransactionCount

**参数：**

- `String` - 要获取发送交易量的地址。
- `Number|String|BN|BigNumber` - (可选) 如果传入值则会覆盖通过 web3.eth.defaultBlock 设置的默认区块号。预定义的区块号 `"latest"`, `"earliest"` `"pending"`, 和 `"genesis"` 等也可以使用。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise):**Object,返回从给定地址发出的交易数量

## 34.sendTransaction

**参数：**1.Object，要发送的交易对象：

- `from` - `String|Number`: 发送账户地址。 如未指定，则使用 web3.eth.defaultAccount 属性值。 或者在 web3.eth.accounts.wallet 中的本地钱包地址。
- `to` - `String`: (可选) 交易消息的目标地址，对于合约创建交易来说其值为空。
- `value` - `Number|String|BN|BigNumber`: (可选) 交易转账金额，以 wei 为单位，也是合约创建交易的初始转账。
- `gas` - `Number`: (可选, 默认值: 待定) 交易可用的 gas 量 (未使用的 gas 会被退回)。
- `gasPrice` - `Number|String|BN|BigNumber`: (可选) 交易可用的 gas 价格，以 wei 为单位, 默认值通过 web3.eth.gasPrice 获得。
- `type` - `Number|String|BN|BigNumber`:(可选)介于0和0x7f之间的8位无符号正数，表示交易的类型。
- `maxFeePerGas` - `Number|String|BN`:(可选，默认为`(2 * block.baseFeePerGas) + maxPriorityFeePerGas`)交易愿意总共支付的gas的最大费用
- `data` - `String`: (可选) 包含合约函数调用数据的 ABI 字节字符串 ，对合约创建交易来说，其值为合约初始化代码。
- `accessList` - `List of hexstrings`(可选)交易计划访问的地址和存储键的列表
- `nonce` - `Number`: (可选) 使用同样的 nonce 值可以覆盖处理中的交易。
- `chain` - `String`: (可选) 默认值为 `mainnet`.
- `hardfork` - `String`: (可选) 默认值为 `london`.
- `common` - `Object`: (可选) 通用对象
  - `customChain` - `Object`: 自定义链属性
    - `name` - `string`:  (可选) 链名称。
    - `networkId` - `number`: 自定义链的网络 ID。
    - `chainId` - `number`: 自定义链的链 ID
  - `baseChain` - `string`: (可选) `mainnet`, `goerli`, `kovan`, `rinkeby`, 或 `ropsten`
  - `hardfork` - `string`: (可选) `chainstart`, `homestead`, `dao`, `tangerineWhistle`, `spuriousDragon`, `byzantium`, `constantinople`, `petersburg`, `istanbul`, `berlin`, 或 `london`

2.`Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(PromiEvent)：** 一个 整合事件发生器的 promise 对象. 将在收到交易收据 receipt 后得到解析, 此外有以下事件可用：

- `"transactionHash"` 返回 `String`: 在交易发出并得到有效的交易哈希后立刻触发。
- `"receipt"` 返回 `Object`: 当收到交易收据后立刻触发。
- `"confirmation"` 返回 `Number`, `Object`: 每次确认后立刻触发，最多 12 次确认。确认编号为第一个参数，收据 receipt 为第二个参数。从交易所在区块被挖到的 0 号确认开始触发。

`"error"` 返回 `Error` and `Object|undefined`: 在发送交易的过程中如果出现错误则立刻触发。如果交易被网络拒绝时附带有交易收据，则第二个参数为交易收据。

> `from` 属性可以是个地址或者 web3.eth.accounts.wallet 的索引值。 随后可以在本地通过该账户对应的私钥签名并通过 web3.eth.sendSignedTransaction() 发送交易。 如果 `chain`、 `hardfork` 或 `common` 这些属性值都没有设置, Web3 会尝试通过网络查询相应的网络 ID 和链 ID 并设置它们。

```js
//例子
// 使用回调
web3.eth.sendTransaction({
    from: '0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe',
    data: code // deploying a contracrt
}, function(error, hash){
    ...
});

// 使用 promise
web3.eth.sendTransaction({
    from: '0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe',
    to: '0x11f4d0A3c12e86B4b5F39B213F7E19D048276DAe',
    value: '1000000000000000'
})
.then(function(receipt){
    ...
});


// 使用事件发生器
web3.eth.sendTransaction({
    from: '0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe',
    to: '0x11f4d0A3c12e86B4b5F39B213F7E19D048276DAe',
    value: '1000000000000000'
})
.on('transactionHash', function(hash){
    ...
})
.on('receipt', function(receipt){
    ...
})
.on('confirmation', function(confirmationNumber, receipt){ ... })
.on('error', console.error); // 如果是 out of gas 错误, 第二个参数为交易收据
```

## 35.sendSignedTransaction

发送已签名的交易，交易签名可以通过 37.web3.eth(.accounts).signTransaction 生成。

**参数：**

- `String` - 16 进制格式的签名交易数据。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值：**34.web3.eth.sendTransaction()

## 36.sign

使用指定账户对数据进行签名，该账户必须先解锁。

**参数：**

- `String` - 待签名的数据。 对于字符串要先用 web3.utils.utf8ToHex 将其转换为 16 进制数据。
- `String|Number` - 用来签名的账户地址。或者本地钱包 web3.eth.accounts.wallet 中的地址或其索引。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**string,返回签名

## 37.signTransaction

签名交易，用来签名的账户需要首先解锁。

**参数：**

- `Object` - 要签名的交易数据，更多信息请看 web3.eth.sendTransaction() 。
- `String` - 签名交易所用的账户地址。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Object,RLP 编码的交易。`raw` 属性可以用来通过 web3.eth.sendSignedTransaction 来发送交易。

## 38.call

执行一个消息调用交易，消息调用交易直接在节点旳 VM 中而不需要通过区块链挖矿来执行。

**参数：**

- `Object` - 交易对象，相关信息可以查看 web3.eth.sendTransaction, 消息调用交易和一般交易的区别是 `from` 属性也是可选的。
- `Number|String|BN|BigNumber` - (可选) 如果传入值则会覆盖通过 web3.eth.defaultBlock 设置的默认区块号。预定义的区块号 `"latest"`, `"earliest"` `"pending"`, 和 `"genesis"` 等也可以使用。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**String,消息调用的返回数据, 比如合约函数的返回值。

## 39.estimateGas

通过执行一个消息调用来得到交易的 gas 用量。

**参数：**

- `Object` - 交易对象，相关信息可以查看 web3.eth.sendTransaction, 消息调用交易和一般交易的区别是 `from` 属性也是可选的。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**String,模拟消息或交易调用的 gas 用量

## 40.getPastLogs

获取匹配给定条件的历史日志。

**参数：**`Object` - 过滤器对象，包含如下字段

- `fromBlock` - `Number|String`: 起始区块 (`"latest"` 是指最近已挖出区块， `"pending"` 指待挖出的区块). 默认为 `"latest"`。
- `toBlock` - `Number|String`: 终止区块 (`"latest"` 是指最近已挖出区块 and `"pending"` 指待挖出的区块). 默认为 `"latest"`。
- `address` - `String|Array`: 需要获取日志的一个或多个地址。
- `topics` - `Array`: 必须出现在日志项中的主题值数组。这里的顺序是很重要的， 你可以使用 `null` 值来忽略某个主题, e.g. `[null, '0x12...']`. 你也可以通过一个数组来指定该主题关联的一系列属性 e.g. `[null, ['option1', 'option2']]`

**返回值(Promise)：**Array,日志对象数组，每个对象的属性:

- `address` - `String`: 事件发生源地址。
- `data` - `String`: 包含未索引的日志参数。
- `topics` - `Array`: 包含最多 4 个 32 字节主题的数组，主题 1-3 包含日志的索引参数。
- `logIndex` - `Number`: 事件在块中的索引位置。
- `transactionIndex` - `Number`: 创建事件的交易的索引位置。
- `transactionHash` 32 字节 - `String`: 创建事件的交易的哈希值。
- `blockHash` 32 字节 - `String`: 创建事件的块的哈希值，若处于 pending 状态，其值为 `null`。
- `blockNumber` - `Number`: 创建事件的块编号，处于 pending 状态时其值为 `null`。

## 41.getWork

获取矿工要满足的挖矿条件。返回当前区块的哈希，种子哈希以及要满足的边界条件（”目标值”）。

**参数**：`Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Array，挖矿条件:

- `String` 32 字节 - 于 **索引位置 0**: 当前区块头工作量证明哈希。
- `String` 32 字节 - 于 **索引位置 1**: 用于 DAG 的种子哈希。
- `String` 32 字节 - 于 **索引位置 2**: 边界条件 (“目标值”), 2^256 / 难度.

## 42.submitWork

用来提交一个工作量证明方案。

**参数：**

- `String` 8 Bytes: 找到的 nonce 值 (64 位)
- `String` 32 字节: 区块头的工作量证明哈希 (256 位)
- `String` 32 字节: The mix digest (256 位)
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**String,如果提交的方案有效返回 `TRUE`，否则返回 false

## 43.requestAccounts

**参数：**`Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

该方法将从当前运行环境（Metamask，Status 或 Mist）中请求/启用帐户，如果你使用默认的 Web3.js provider (WebsocketProvider, HttpProvidder and IpcProvider) 连接节点，则该方法不起作用。 该方法只在你使用像 Status, Mist or Metamask 这些应用的嵌入式 provider 时才有效。

**返回值(Promise)：Array,**返回所启用账户数组

## 44.getChainId

**返回值(Promise)：**Number,如果提交的方案有效返回 `TRUE`，否则返回 false

## 45.getNodeInfo

**返回值(Promise)：**String,当前客户端版本

## 46.getProof

返回账户及相关存储数据，包括 [EIP-1186](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1186.md) 所描述的默克尔证明

**参数：**

- `String` 20 字节: 账户或合约地址。
- `Number[] | BigNumber[] | BN[] | String[]` 32 字节: 应该被证明和包含的存储键值数组， 更多信息请看 web3.eth.getStorageAt。
- `Number | String | BN | BigNumber`: 区块号。 也可以使用预定义区块号 `"latest"`, `"earliest"`, and `"genesis"` 。
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为函数运行结果。

**返回值(Promise)：**Object,一个账户对象：

- `address` - `String`: 账户地址。
- `balance` - `String`: 账户余额。 详情请查看 web3.eth.getBalance.
- `codeHash` - `String`: 账户关联代码哈希值。 对于没有代码的一般账户会返回 “0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470”
- `nonce` - `String`: 账户 nonce 值。
- `storageHash` - `String`: 存储树的根哈希. 所有存储将由此根哈希开始提供默克尔证明。
- `accountProof` - `String[]`: rlp 编码序列化的默克尔节点数组，从状态根节点开始，接着是以 SHA3 (地址) 为键值的路径。
  - `storageProof` - `Object[]` 请求的存储条目数组。
  - `key` - `String` 请求的存储键值。`value` - `String` 存储数据。


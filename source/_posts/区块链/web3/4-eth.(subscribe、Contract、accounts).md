---
title: 4-eth.(subscribe、Contract、accounts)

date: 2022-09-03	

categories: web3.js	

tags: [区块链,web3.js]
---	

# subscribe(订阅区块链中的指定事件)

## 1.1 clearSubscriptions

重置订阅

**参数：**`Boolean`: 值为 `true` 则表示保持 `同步` 订阅

**返回值（Boolean）**:是否同步订阅

## 1.2 subscribe

订阅区块链中的指定事件

**参数：**

- `String` - 订阅类型
- `Mixed` - (可选) 依赖于订阅类型的可选额外参数
- `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为结果。该函数在每次订阅事件发生时都会被调用，订阅实例会作为第三个参数传递进来

**返回值(EventEmitter)：**一个订阅实例

- `subscription.id`: 订阅 id 编号，用于标识一个订阅以及进行后续的取消订阅操作
- `subscription.subscribe([callback])`: 可用于使用相同的参数进行再次订阅
- `subscription.unsubscribe([callback])`: 取消订阅，如果成功取消的话，在回调函数中返回 TRUE
- `subscription.arguments`: 订阅参数，在重新订阅时使用
- `on("data")` 返回 `Object`: 每次有新的日志时都触发该事件，参数为日志对象
- `on("changed")` 返回 `Object`: 每次有日志从区块链上移除时触发该事件，被移除的日志对象将添加额外的属性： `"removed: true"`
- `on("error")` 返回 `Object`: 当订阅发生错误时，触发此事件
- `on("connected")` 返回 `String`: 一旦订阅成功就会触发该事件，返回订阅 id

**通知返回值：**`Mixed` - 取决于订阅具体是什么，可以参考后面提到的不同订阅获取更多相关信息

```js
//例子
var subscription = web3.eth.subscribe('logs', {
    address: '0x123456..',
    topics: ['0x12345...']
}, function(error, result){
    if (!error)
        console.log(result);
});

// 取消订阅
subscription.unsubscribe(function(error, success){
    if(success)
        console.log('Successfully unsubscribed!');
});
```

### 1.2.1 subscribe(“pendingTransactions”)

订阅 pending 状态的交易。

**参数：**

1. `String` - `"pendingTransactions"`, 订阅类型
2. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为结果。

**返回值(EventEmitter)：**一个订阅实例

- `"data"` 返回 `String`: 当接收到 pending 状态的交易时触发并返回交易哈希
- `"error"` 返回 `Object`: 当订阅中发生错误时触发

**通知返回值：**

1. `Object|Null` - 如果订阅失败第一个参数为错误对象
2. `String` - 第二个参数为交易哈希

### 1.2.2 subscribe(“newBlockHeaders”)

订阅新的区块头生成事件， 可用作检查链上变化的计时器。

**参数：**

1. `String` - `"newBlockHeaders"`, 订阅类型
2. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为结果。

**返回值(EventEmitter)：**一个订阅实例

- `"data"` 返回 `Object`: 当收到新的区块头时触发
  - `number` - `Number`: 区块编号，对于 pending 状态的块该值为``null``。
  - `hash` 32 Bytes - `String`: 块的哈希值，对于 pending 状态的块该值为``null``。
  - `parentHash` 32 Bytes - `String`: 父区块哈希值。
  - `nonce` 8 Bytes - `String`: 用来生成工作量证明的 nonce 值. 对于 pending 状态的块该值为``null``。
  - `sha3Uncles` 32 Bytes - `String`: 区块中叔块数据的 sha3哈希值。
  - `logsBloom` 256 Bytes - `String`: 区块日志数据的布隆过滤器。对于 pending 状态的块该值为``null``。
  - `transactionsRoot` 32 Bytes - `String`: 区块中交易 trie 树的根节点。
  - `stateRoot` 32 Bytes - `String`: 区块中状态 trie 树的根节点。
  - `receiptsRoot` 32 Bytes - `String`: 收据根节点。
  - `miner` - `String`: 接收挖矿奖励的矿工地址。
  - `extraData` - `String`: 区块的额外数据字段。
  - `gasLimit` - `Number`: 该块允许的最大 gas 用量。
  - `gasUsed` - `Number`: 该块中所有交易使用的 gas 总量。
  - `timestamp` - `Number`: 区块时间戳。
- `"error"` 返回 `Object`: 当订阅中出现错误时触发
- `"connected"` 返回 `Number`: 订阅成功后触发，返回订阅 id

**通知返回值：**

1. `Object|Null` - 如果订阅失败第一个参数为错误对象
2. `Object` - 像上面那样的区块头对象

### 1.2.3 subscribe(“syncing”)

订阅同步事件。当节点正在同步数据时将返回一个同步对象，同步结束后返回 `FALSE`。

**参数：**

1. `String` - `"syncing"`, 订阅类型
2. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为结果。

**返回值(EventEmitter)：**一个订阅实例

- `"data"` 返回 `Object`: 收到同步对象时触发(参考14.web3.eth.isSyncing()返回值)
- `"changed"` 返回 `Object`: 当同步状态由 `true` 变为 `false` 时触发
- `"error"` 返回 `Object`: 当订阅中出现错误时触发

**通知返回值：**

1. `Object|Null` - 如果订阅失败第一个参数为错误对象
2. `Object|Boolean` - 同步对象, 同步开始后返回 `true`，同步结束后返回 false

### 1.2.4 subscribe(“logs”)

订阅日志，并按指定条件进行过滤。 如果一个有效的 `fromBlock` 属性被指定，Web3 会提取从这个点开始的日志，必要时回填返回值。**参数：**

- `"logs"` - `String`, 订阅类型
- `Object` - 订阅选项
  - `fromBlock` - `Number`: 最早区块编号， 默认为 `null`。
  - `address` - `String|Array`: 地址或地址列表，仅订阅来自这些指定账户地址的日志。
  - `topics` - `Array`: 一个主题数组，数组中每个元素都应出现在日志项中。 数组中主题的顺序是很重要的, 如果你不想监听某个主题可以用 `null` 值, 比如 `[null, '0x00...']`. 你也可以为每个主题传入另外一个数组来表示主题选项， 比如 `[null, ['option1', 'option2']]`
- `callback` - `Function`: (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为结果。每当有订阅事件发生时都会被调用。

**返回值(EventEmitter)：**一个订阅实例

- `"data"` 返回 `Object`: 接收到新日志时触发，参数为日志对象(参考40. web3.eth.getPastEvents()返回值)
- `"changed"` 返回 `Object`: 日志从链上移除时触发，该日志同时带有附加属性 `"removed: true"`
- `"error"` 返回 `Object`: 当订阅中出现错误时触发
- `"connected"` 返回 `Number`: 订阅成功后触发，返回订阅 id

**通知返回值：**

1. `Object|Null` - 如果订阅失败第一个参数为错误对象
2. `Object` - 日志对象，参考 40.web3.eth.getPastEvents()

# Contract(与区块链上的智能合约交互)

## 2.1 new contract

创建新的合约实例，并在其 json interface 对象中定义所有的方法和事件。

**参数：**

- `jsonInterface` - `Object`: 所要实例化合约的 json 接口
- `address` - `String` （可选）: 要调用的智能合约地址
- `options` - `Object` （可选）: 合约配置选项。 其中某些选项被用作合约调用和交易的回调:
  - `from` - `String`: 交易发送方地址
  - `gasPrice` - `String`: 为交易指定的 gas 价格，以 wei 为单位
  - `gas` - `Number`: 交易可用的最大 gas 量（gas limit）。
  - `data` - `String`: 合约字节码。当合约被 :ref:部署 <contract-deploy>时需要使用。

**返回值(Object)：**带有所有方法和事件的合约实例

```js
//例子
var myContract = new web3.eth.Contract([...], '0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe', {
    from: '0x1234567890123456789012345678901234567891', // 默认交易发送地址
    gasPrice: '20000000000' // 以 wei 为单位的默认 gas 价格，当前价格为 20 gwei
});
```

## 2.2 options(选项属性)

- `address` - `String`: 合约的部署地址。 见 options.address.
- `jsonInterface` - `Array`: 合同的 json 接口。见 options.jsonInterface.
- `data` - `String`: 合约字节码。 部署 合约时使用。
- `from` - `String`: 交易发起方地址。
- `gasPrice` - `String`: 用于交易的 gas 价格。以 wei 为单位。
- `gas` - `Number`: 可用于该交易的 gas 用量上限 (gas limit)。
- `handleRevert` - `Boolean`: 如果这里不设置，将使用 Eth 模块提供的默认值。 见 11.handleRevert.
- `transactionBlockTimeout` - `Number`: 如果这里不设置，将使用 Eth 模块提供的默认值。 见6. transactionBlockTimeout.
- `transactionConfirmationBlocks` - `Number`: 如果这里不设置，将使用 Eth 模块提供的默认值。 见 7.transactionConfirmationBlocks.
- `transactionPollingTimeout` - `Number`: 如果这里不设置，将使用 Eth 模块提供的默认值。 见 9.transactionPollingTimeout.
- `chain` - `Number`: 如果这里不设置，将使用 Eth 模块提供的默认值。 见 4.defaultChain.
- `hardfork` - `Number`: 如果这里不设置，将使用 Eth 模块提供的默认值。 见3. defaultHardfork.
- `common` - `Number`: 如果这里不设置，将使用 Eth 模块提供的默认值。 见 5.defaultCommon.

```js
//例子
myContract.options;
> {
    address: '0x1234567890123456789012345678901234567891',
    jsonInterface: [...],
    from: '0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe',
    gasPrice: '10000000000000',
    gas: 1000000
}

myContract.options.from = '0x1234567890123456789012345678901234567891'; // 默认交易发送方地址
myContract.options.gasPrice = '20000000000000'; // 默认 gas 价格，以 wei 为单位
myContract.options.gas = 5000000; // 5M gas 作为备用值
```

### 2.2.1 options.address

用于本合约实例的地址。 所有通过 web3.js 从这个合约生成的交易都将包含这个地址作为 “to” 地址（也就是交易接收方地址）

**属性：**`address` - `String|null`: 本合约的地址, 如果未设置其值为 `null`。

```js
//例子 
myContract.options.address;
> '0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae'

// 设置新地址
myContract.options.address = '0x1234FFDD...';
```

### 2.2.2 options.jsonInterface

从合约的 ABI 派生出来的 json 接口 对象。

**属性：**`jsonInterface` - `Array`: 当前合约的 json 接口 。重设该接口会导致合约实例方法和事件的重新生成。

```js
//例子 
myContract.options.jsonInterface;
> [{
    "type":"function",
    "name":"foo",
    "inputs": [{"name":"a","type":"uint256"}],
    "outputs": [{"name":"b","type":"address"}]
},{
    "type":"event",
    "name":"Event",
    "inputs": [{"name":"a","type":"uint256","indexed":true},{"name":"b","type":"bytes32","indexed":false}],
}]

// 设置新的接口
myContract.options.jsonInterface = [...];
```

## 2.3 Methods(方法)

### 2.3.1 clone

克隆当前合约实例

**返回值(Object)：**新合约实例

### 2.3.2 deploy

调用此函数将合约部署到区块链上。 成功部署后 promise 对象会被解析为新的合约实例。

**参数：**`options` **-** `Object`: 用于合约部署的选项

- `data` - `String`: 合约字节码
- `arguments` - `Array` （可选）: 在部署合约时传递给构造函数的参数。

**返回值：**`Object`交易对象:

- `Array` - 参数: 之前传递给方法的参数。它们是可以被改变的。
- `Function` - send: 用来部署合约。promise 会被解析为合约实例而不是交易收据。
- `Function` - estimateGas: 估算部署合约所需要的 gas 用量。
- `Function` - encodeABI: 编码由合约数据和构造函数参数构成的合约部署 ABI。

```js
//例子
myContract.deploy({
    data: '0x12345...',
    arguments: [123, 'My String']
})
.send({
    from: '0x1234567890123456789012345678901234567891',
    gas: 1500000,
    gasPrice: '30000000000000'
}, function(error, transactionHash){ ... })
.on('error', function(error){ ... })
.on('transactionHash', function(transactionHash){ ... })
.on('receipt', function(receipt){
   console.log(receipt.contractAddress) // 包含新合约地址
})
.on('confirmation', function(confirmationNumber, receipt){ ... })
.then(function(newContractInstance){
    console.log(newContractInstance.options.address) // 带有新合约地址的合约实例
});


//数据已经设置为合约本身的选项
myContract.options.data = '0x12345...';

myContract.deploy({
    arguments: [123, 'My String']
})
.send({
    from: '0x1234567890123456789012345678901234567891',
    gas: 1500000,
    gasPrice: '30000000000000'
})
.then(function(newContractInstance){
    console.log(newContractInstance.options.address) // 具有新合同地址的合约实例
});


// 只是编码
myContract.deploy({
    data: '0x12345...',
    arguments: [123, 'My String']
})
.encodeABI();
> '0x12345...0000012345678765432'


// gas 估算
myContract.deploy({
    data: '0x12345...',
    arguments: [123, 'My String']
})
.estimateGas(function(err, gas){
    console.log(gas);
});
```

### 2.3.3 methods

#### 2.3.3.1 methods.myMethod.call

将在不发送交易的情况下调用该“常量”方法并在 EVM 中执行其智能合约方法。注意此种调用方式无法改变智能合约状态。

**参数 ：**

1. `options` - `Object` （可选）: 用于发起调用的选项。
   - `from` - `String` （可选）: 调用“交易”的发起方地址。
   - `gasPrice` - `String` （可选）: 用于该调用“交易”的 gas 价格。以 wei 为计量单位。
   - `gas` - `Number` （可选）: 用于该调用“交易”的 gas 用量上限 (gas limit)。
2. `callback` - `Function` （可选）: 回调函数，回调时将以智能合约方法执行结果作为第二个参数，以错误对象作为第一个参数。

**返回值(Promise)：**Mixed,智能合约方法返回值。 如果只返回一个值，则按原样返回。如果有多个返回值，则作为一个带有属性和索引的对象返回。

#### 2.3.3.2  methods.myMethod.send

向合约发送交易来执行其方法。注意这会改变合约状态。

**参数 ：**

1. `options` - `Object` （可选）: 用于发起调用的选项。
   - `from` - `String` （可选）: 调用“交易”的发起方地址。
   - `gasPrice` - `String` （可选）: 用于该调用“交易”的 gas 价格。以 wei 为计量单位。
   - `gas` - `Number` （可选）: 用于该调用“交易”的 gas 用量上限 (gas limit)。
   - `value` - `Number|String|BN|BigNumber`（可选）: 交易转账金额，以 wei 为单位。
2. `callback` - `Function` （可选）: 回调函数，回调时将以智能合约方法执行结果作为第二个参数，以错误对象作为第一个参数。

**返回值(PromiEvent)：**一个 整合了事件发生器 <promiEvent> promise. 当收到交易*收据*时会被解析, 当该``send()``调用是从``someContract.deploy()``发出时，promise 会解析为*新合约地址*。重设该接口会导致合约实例方法和事件的重新生成。 此外也存在下面这些事件:

- `"transactionHash"` 返回 `String`: 发送交易且得到交易哈希值后立即触发。
- `"receipt"` 返回 `Object`: 当收到交易*收据*时触发。合约收据带有的不是``logs``，而是以事件名称为健，以事件本身为属性值的``events``。 关于返回事件对象的详情，参见 getPastEvents 返回值 。
- `"confirmation"` 返回 `Number`, `Object`: 从区块被挖到的第一个区块确认开始，每次确认都会触发，直到第 24 次确认。触发时第一个参数为收到的确认数，第二个参数为收到交易收据。
- `"error"` 返回 `Error` 和 `Object|undefined`: 交易发送过程中出错时触发。如果交易被网络拒绝且带有交易收据，第二个参数就是该交易收据。

#### 2.3.3.3 methods.myMethod.estimateGas

通过在 EVM 中执行方法来估算链上执行是需要的 gas 用量。 由于彼时合约状态的不同，当前估算的 gas 用量和随后通过真实交易所得到的实际 gas 用量可能会有所出入。

**参数 ：**

1. `options` - `Object` （可选）: 用于发起调用的选项。
   - `from` - `String` （可选）: 调用“交易”的发起方地址。
   - `gas` - `Number` （可选）: 用于该调用“交易”的 gas 用量上限 (gas limit)。
   - `value` - `Number|String|BN|BigNumber` （可选）: 交易转账金额，以 wei 为单位。
2. `callback` - `Function` （可选）: 回调函数，回调时将以智能合约方法执行结果作为第二个参数，以错误对象作为第一个参数。

**返回值(Promise)：**Number,估算的 gas 用量

#### 2.3.3.4 methods.myMethod.encodeABI

为指定的合约方法进行 ABI 编码，可用于发送交易、调用方法或向另一个合约方法传递参数。

**返回值(String)：**编码后的 ABI 字节码，可用于交易发送或方法调用

#### 2.3.3.5 methods.myMethod.createaccesslist

创建访问列表并将用call调用，方法执行将在EVM中执行时访问该列表。注:目前eth_createAccessList好像只有Geth支持

**参数 ：**

1. `options` - `Object` （可选）: 用于发起调用的选项。
   - `from` - `String` （可选）: 调用“交易”的发起方地址。
   - `gas` - `Number` （可选）: 用于该调用“交易”的 gas 用量上限 (gas limit)。
2. `block` - `String|Number|BN|BigNumber`(可选):块号或哈希。或者`"earliest"`, `"latest"` , `"pending"` , `"safe"`或者`"finalized"`作为 默认块参数.
3. `callback` - `Function` （可选）: 回调函数，回调时将以智能合约方法执行结果作为第二个参数，以错误对象作为第一个参数。

**返回值(Promise)：**Object,为交易生成的访问列表。

## 2.4 Events(事件)

### 2.4.1 once

订阅一个事件并在第一次事件触发或错误发生后立即取消订阅。一个事件仅触发一次。

**参数 ：**

1. `event` - `String`: 要订阅的合约事件名, 或者用 `"allEvents"` 来订阅所有事件。
2. `options` - `Object` （可选）: 用于部署的选项。
   1. `filter` - `Object` （可选）: 按索引参数过滤事件, 例如 `{filter: {myNumber: [12,13]}}` 表示 “myNumber” 为 12 或 13 的所有事件。
   2. `topics` - `Array` （可选）: 手动设置事件过滤器的主题。如果提供了过滤器属性和事件签名，则不会自动设置（topic [0]）
3. `callback` - `Function`: 回调函数，触发时把*事件*对象作为第二个参数，错误作为第一个参数。 关于详细事件结构，参见 getPastEvents 返回值

### 2.4.2 events

订阅指定的合约事件

**参数 ：**

1. `options` - `Object` （可选）: 用于部署的选项。
   - `filter` - `Object` （可选）: 按索引参数过滤事件, 例如 `{filter: {myNumber: [12,13]}}` 表示 “myNumber” 为 12 或 13 的所有事件。
   - `fromBlock` - `Number|String|BN|BigNumber` （可选）: 读取从该编号开始的区块中的事件（大于或等于该区块号）。 也可以使用预先定义的区块号，比如 `"latest"`, `"earliest"`, `"pending"`, `"genesis"` 等。
   - `topics` - `Array` （可选）: 手动设置事件过滤器的主题。如果提供了过滤器属性和事件签名，则不会自动设置（topic [0]）。
2. `callback` - `Function` （可选）: 回调函数，回调时将以智能合约方法执行结果作为第二个参数，以错误对象作为第一个参数。

**返回值(EventEmitter)：**该事件发生器有以下事件：

- `"data"` 返回 `Object`: 接收到新传入的事件时触发，参数为事件对象。
- `"changed"` 返回 `Object`: 当事件从区块链上移除时触发。 该事件带有额外属性 `"removed: true"`。
- `"error"` 返回 `Object`: 当订阅中出现错误时触发。
- `"connected"` 返回 `String`: 当订阅成功连接时触发一次。返回订阅 id。

返回的事件 “对象” 结构如下：

- `event` - `String`: 事件名称。
- `signature` - `String|Null`: 事件签名，如果是匿名事件，其值为 `null`。
- `address` - `String`: 该事件的发源地地址。
- `returnValues` - `Object`: 事件返回值， 比如 `{myVar: 1, myVar2: '0x234...'}`.
- `logIndex` - `Number`: 事件在区块中的索引位置。
- `transactionIndex` - `Number`: 事件所在交易在区块中的索引位置。
- `transactionHash` 32 Bytes - `String`: 事件所在交易的哈希值。
- `blockHash` 32 Bytes - `String`: 事件所在区块链的哈希值。区块处于 pending 状态时其值为 `null` 。
- `blockNumber` - `Number`: 事件所在区块的区块号。 区块处于 pending 状态时其值为 `null` 。
- `raw.data` - `String`: 包含未索引的日志参数。
- `raw.topics` - `Array`: 最大可保存 4 个 32 字节长的主题数组，主题 1-3 包含事件的索引参数。

```js
//例子 
myContract.events.MyEvent({
    filter: {myIndexedParam: [20,23], myOtherIndexedParam: '0x123456789...'}, // 使用数组表示 或：如 20 或 23。
    fromBlock: 0
}, function(error, event){ console.log(event); })
.on("connected", function(subscriptionId){
    console.log(subscriptionId);
})
.on('data', function(event){
    console.log(event); // 与上述可选的回调结果相同
})
.on('changed', function(event){
    // 从本地数据库中删除事件
})
.on('error', function(error, receipt) { // 如果交易被网络拒绝并带有交易收据，第二个参数将是交易收据。
    ...
});

// 事件输出例子
> {
    returnValues: {
        myIndexedParam: 20,
        myOtherIndexedParam: '0x123456789...',
        myNonIndexParam: 'My String'
    },
    raw: {
        data: '0x7f9fade1c0d57a7af66ab4ead79fade1c0d57a7af66ab4ead7c2c2eb7b11a91385',
        topics: ['0xfd43ade1c09fade1c0d57a7af66ab4ead7c2c2eb7b11a91ffdd57a7af66ab4ead7', '0x7f9fade1c0d57a7af66ab4ead79fade1c0d57a7af66ab4ead7c2c2eb7b11a91385']
    },
    event: 'MyEvent',
    signature: '0xfd43ade1c09fade1c0d57a7af66ab4ead7c2c2eb7b11a91ffdd57a7af66ab4ead7',
    logIndex: 0,
    transactionIndex: 0,
    transactionHash: '0x7f9fade1c0d57a7af66ab4ead79fade1c0d57a7af66ab4ead7c2c2eb7b11a91385',
    blockHash: '0xfd43ade1c09fade1c0d57a7af66ab4ead7c2c2eb7b11a91ffdd57a7af66ab4ead7',
    blockNumber: 1234,
    address: '0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe'
}
```

## 2.5 getPastEvents

读取合约历史事件。

**参数 ：**

- `event` - `String`: 合约事件名，或者用 `"allEvents"` 读取所有事件。
- `options` - `Object` （可选）: 用于部署的选项。
  1. `filter` - `Object` （可选）: 按索引参数过滤事件, 例如 `{filter: {myNumber: [12,13]}}` 表示 “myNumber” 为 12 或 13 的所有事件。
  2. `fromBlock` - `Number|String|BN|BigNumber` （可选）: 读取从该编号开始的区块中的历史事件（大于或等于该区块号）。 也可以使用预先定义的区块编号，比如 `"latest"`, `"earliest"`, `"pending"`, `"genesis"` 。
  3. `toBlock` - `Number|String|BN|BigNumber` （可选）: 读取截止到该编号的区块中的历史事件（小于或等于该区块号）（默认值为 “latest”）。 也可以使用预先定义的区块编号，比如 `"latest"`, `"earliest"`, `"pending"`, `"genesis"` 。
  4. `topics` - `Array` （可选）: 用来手动设置事件过滤器的主题。如果设置了 filter 属性和事件签名，则不会自动设置（topic [0]）。
- `callback` - `Function` （可选）: 回调函数，回调时将以智能合约方法执行结果作为第二个参数，以错误对象作为第一个参数。

**返回值(Promise)：**Array，满足给定事件或过滤条件的历史事件对象数组

# accounts(生成账户和用来签名交易与数据的函数)

> 该程序包尚未经过安全评审，可能存在安全隐患。在用于生产环境之前，请采取必要的措施来合理的清理内存，安全的存储私钥并完整测试交易的接收和发送功能！

## 3.1 create

生成具有公私钥的账户对象

**参数 ：**`entropy` - `String` (可选): 增加混淆度的随机字符串，至少 32 字符长。如果未设定将使用 randomhex 生成一个随机字符串。

**返回值(Object)：** 账户对象，具有如下结构：

- `address` - `string`: 账户地址。
- `privateKey` - `string`: 帐户私钥。绝对不要将其共享或以未加密的形式存储在 localstorage 中，并确保使用后将内存清空 。
- `signTransaction(tx [, callback])` - `Function`: 签名交易的函数. 更多信息可参考 web3.eth.accounts.signTransaction() 。
- `sign(data)` - `Function`: 用来签名数据的函数. 更多信息可参考 web3.eth.accounts.sign() 。

```js
//例子 
{
    address: "0xF2CD2AA0c7926743B1D4310b2BC984a0a453c3d4",
    privateKey: "0xd7325de5c2c1cf0009fac77d3d04a9c004b038883446b065871bc3e831dcd098",
    signTransaction: function(tx){...},
    sign: function(data){...},
    encrypt: function(password){...}
}
```

## 3.2 privateKeyToAccount

通过私钥来创建账户对象

**参数 ：**`privateKey` - `String`: 用来创建账户的私钥。

**返回值(Object)：** 具体见web3.eth.accounts.create()

## 3.3 signTransaction

使用给定的私钥签名以太坊交易

**参数 ：**

- `tx` **-** `Object`:交易对象，具有如下结构：
  - `nonce` - `String`: (可选) 签名交易时所用的 nonce 值。默认使用33. web3.eth.getTransactionCount().
  - `chainId` - `String`: (可选) 签名交易时所用的 chain id。 默认使用 web3.eth.net.getId().
  - `to` - `String`: (可选) 交易接受者，部署合约时其值为空。
  - `data` - `String`: (可选) 交易调用数据，对简单转账交易来说其值为空。
  - `value` - `String`: (可选) 以 wei 为单位的以太币转账数量。
  - `gasPrice` - `String`: (可选) 交易所用的燃料价格，如果传入为空值则会使用 web3.eth.gasPrice()
  - `gas` - `String`: 交易可用的燃料上限。
  - `chain` - `String`: (可选) 默认值为 `mainnet`.
  - `hardfork` - `String`: (可选) 默认值为 `petersburg`.
  - `common` - `Object`: (可选) common 对象：
    - `customChain` - `Object`: 自定义链属性
      - `name` - `string`: (可选) 链名称
      - `networkId` - `number`: 自定义链的网络 id
      - `chainId` - `number`: 自定义链的 chain id
    - `baseChain` - `string`: (可选) `mainnet`, `goerli`, `kovan`, `rinkeby`, 或 `ropsten`
    - `hardfork` - `string`: (可选) `chainstart`, `homestead`, `dao`, `tangerineWhistle`, `spuriousDragon`, `byzantium`, `constantinople`, `petersburg`, 或 `istanbul`
- `privateKey` - `String`: 签名交易所用的私钥。
- `callback` - `Function`: (可选) 可选的回调函数, 其第一个返回值为错误对象，第二个返回值为调用结果。

**返回值(Promise)：** **Object**，**RLP 编码的交易对象：**

- `messageHash` - `String`: 给定消息的哈希值。
- `r` - `String`: 签名的头 32 个字节。
- `s` - `String`: 签名接下来的 32 个字节。
- `v` - `String`: 恢复值 + 27。
- `rawTransaction` - `String`: RLP 编码的交易, 可使用 35.web3.eth.sendSignedTransaction 直接发送。
- `transactionHash` - `String`: RLP 编码交易的交易哈希。

## 3.4 recoverTransaction

恢复用于签名给定 RLP 编码交易的以太坊地址。

**参数 ：**`signature` - `String`: RLP 编码的交易。

**返回值(String)：** 用于签名该交易的以太坊地址。

## 3.5 hashMessage

计算给定消息的哈希值，以便于用于 web3.eth.accounts.recover() 函数。 数据采用 UTF-8 十六进制解码封装: `"\x19Ethereum Signed Message:\n" + message.length + message`，并使用 keccak256 进行哈希运算。

**参数 ：**`message` - `String`: 要进行哈希运算的消息，如果是 16 进制字符串，首先对其进行 UTF-8 解码。

**返回值(String)：** 哈希运算后的消息

## 3.6 sign

签名任意数据。

**参数 ：**

1. `data` - `String`: 要签名的数据。
2. `privateKey` - `String`: 用以签名的私钥

**返回值(**Object**)：** 签名对象

- `message` - `String`: 给定消息。
- `messageHash` - `String`: 给定消息的哈希值。
- `r` - `String`: 签名的头 32 个字节。
- `s` - `String`: 签名接下来的 32 个字节。
- `v` - `String`: 恢复值(Recovery) + 27。

## 3.7 recover

恢复用于签名给定数据的以太坊地址。

**参数 ：**

- `message|signatureObject` - `String|Object`: 要么是签名消息或哈希，要么是带有下面属性的签名对象：

  - `messageHash` - `String`: 给定消息的哈希值，该消息已加前缀 `"\x19Ethereum Signed Message:\n" + message.length + message`。
- `r` - `String`: 签名的头 32 个字节
  - `s` - `String`: 签名接下来的 32 个字节。
- `v` - `String`: 恢复值 + 27。
- `signature` - `String`: 原始 RLP 编码签名或者通过第 2-4 号参数指定 v，r，s 值。
- `preFixed` - `Boolean` (可选, 默认值: `false`): 如果其值为 `true`, 给定的消息不会自动添加前缀 `"\x19Ethereum Signed Message:\n" + message.length + message`, 假定给定消息已经带了该前缀。

**返回值(String)：** 哈希运算后的消息

```js
//例子
web3.eth.accounts.recover({
    messageHash: '0x1da44b586eb0729ff70a73c326926f6ed5a25f5b056e7f47fbc6e58d86871655',
    v: '0x1c',
    r: '0xb91467e570a6466aa9e9876cbcd013baba02900b8979d43fe208a4a4f339f5fd',
    s: '0x6007e74cd82e037b800186422fc2da167c747ef045e5d18a5f5d4300f8e1a029'
})
> "0x2c7536E3605D9C16a7a3D7b1898e529396a65c23"

// 签名消息, 签名
web3.eth.accounts.recover('Some data', '0xb91467e570a6466aa9e9876cbcd013baba02900b8979d43fe208a4a4f339f5fd6007e74cd82e037b800186422fc2da167c747ef045e5d18a5f5d4300f8e1a0291c');
> "0x2c7536E3605D9C16a7a3D7b1898e529396a65c23"
```

## 3.8 encrypt

将私钥加密变换为 keystore v3 标准格式

**参数 ：**

1. `privateKey` - `String`: 要加密的私钥。
2. `password` - `String`: 用于加密的密码。

**返回值(**Object**)：** 加密后的 keystore v3 JSON。

```js
//例子
web3.eth.accounts.encrypt('0x4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318', 'test!')
> {
    version: 3,
    id: '04e9bcbb-96fa-497b-94d1-14df4cd20af6',
    address: '2c7536e3605d9c16a7a3d7b1898e529396a65c23',
    crypto: {
        ciphertext: 'a1c25da3ecde4e6a24f3697251dd15d6208520efc84ad97397e906e6df24d251',
        cipherparams: { iv: '2885df2b63f7ef247d753c82fa20038a' },
        cipher: 'aes-128-ctr',
        kdf: 'scrypt',
        kdfparams: {
            dklen: 32,
            salt: '4531b3c174cc3ff32a6a7a85d6761b410db674807b2d216d022318ceee50be10',
            n: 262144,
            r: 8,
            p: 1
        },
        mac: 'b8b010fff37f9ae5559a352a185e86f9b9c1d7f7a9f1bd4e82a5dd35468fc7f6'
    }
}
```

## 3.9 decrypt

解密 keystore v3 JSON，并创建账户

**参数 ：**

1. `encryptedPrivateKey` - `String`: 要进行解密的加密私钥。
2. `password` - `String`: 用于加密的密码。

**返回值(**Object**)：** 解密的帐户对象。

## 3.10 wallet

一个多账户内存钱包，这些账户可以用于 web3.eth.sendTransaction()。

```js
//例子
web3.eth.accounts.wallet;
> Wallet {
    0: {...}, // 账户索引
    "0xF0109fC8DF283027b6285cc889F5aA624EaC1F55": {...},  // 账户地址
    "0xf0109fc8df283027b6285cc889f5aa624eac1f55": {...},  // 全小写账户地址
    1: {...},
    "0xD0122fC8DF283027b6285cc889F5aA624EaC1d23": {...},
    "0xd0122fc8df283027b6285cc889f5aa624eac1d23": {...},

    add: function(){},
    remove: function(){},
    save: function(){},
    load: function(){},
    clear: function(){},

    length: 2,
}
```

### 3.10.1 wallet.create

在钱包中创建一个或多个账户。不会覆盖已经存在的钱包。

**参数：**

1. `numberOfAccounts` - `Number`: 要创建的账户数量。设为空值时创建空钱包。
2. `entropy` - `String` (可选): 创建账户时为增加随机性而使用的随机字符串，至少 32 个字符长。

**返回值(Object)：**钱包对象

### 3.10.2 wallet.add

使用私钥或账户对象向钱包中添加账户

**参数：**`account` - `String|Object`: 私钥或者通过 web3.eth.accounts.create() 创建的账户对象。

**返回值(Object)：**所添加的账户

### 3.10.3 wallet.remove

从钱包中移除账户。

**参数：**`account` - `String|Number`: 账户地址，或者账户在钱包中的索引。

**返回值(Boolean)：**移除成功则返回 `true` ，否则返回 `false` 。

### 3.10.4 wallet.clear

安全地清空钱包并移除全部账户。

**返回值(Object)：**钱包对象。

### 3.10.5 wallet.encrypt

加密所有的钱包账户为 keystore v3 对象。

**参数：**`password` - `String`: 用于加密的密钥。

**返回值(Array)：**已加密的 keystore v3 对象。

### 3.10.6 wallet.decrypt

解密 keystore v3 对象。

**参数：**

1. `keystoreArray` - `Array`: 要解密的加密 keystore v3 对象。
2. `password` - `String`: 用来加密的密钥。

**返回值(Object)：**钱包对象

### 3.10.7 wallet.save

将钱包加密并在 local storage 中保存为字符串。

**参数：**

1. `password` - `String`: 用来加密钱包的密钥。
2. `keyName` - `String`: (可选) 用来在 local storage 中寻址的健值, 默认值为 `"web3js_wallet"`。

**返回值(Boolean)**：是否成功

### 3.10.8 wallet.load

从 local storage 中载入钱包并解密。

**参数：**

1. `password` - `String`: 用来加密钱包的密钥。
2. `keyName` - `String`: (可选) 用来在 local storage 中寻址的健值, 默认值为 `"web3js_wallet"`。

**返回值(Object)**：钱包对象

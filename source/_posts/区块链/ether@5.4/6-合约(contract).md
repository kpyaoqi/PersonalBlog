---
title: 6-合约(contract)

date: 2022-08-20	

categories: ether@5.4.js	

tags: [区块链,ether@5.4.js]
---	

# 创建实例

**new *ethers*.Contract( *address* , *abi* , *signerOrProvider* )：**如果给定了Provider，那么合约只有**只读**访问权限，而Signer则提供了**对状态操作方法**的访问权限。

## 1.1 attach

*contract*.**attach**( *addressOrName* ) ⇒ *合约（**Contract**）*

返回一个新地址的**Contract**新实例,如果网络上有多个类似的合约副本,并且您希望与它们中的每一个进行交互，那么这是非常有用的。

## 1.2 connect

*contract*.**connect**( *providerOrSigner* ) ⇒ *合约（**Contract**）*

返回合约的一个新实例，但需要连接到*providerOrSigner*。

通过传入一个Provider，这将返回一个低级的合约实例，它只有只读访问的权限(即常量调用)。

通过传入一个Signer，这将返回一个代表该签名人（signer）的合约实例。

# 属性

***contract*.address ⇒ *string< 地址(Address) >*：**这是构建合约时使用的地址(或ENS名称)。

***contract*.resolvedAddress ⇒ *string< 地址(Address) >*：**这是一个合约对象将解析合约地址的promise。 如果一个地址(Address)被提供	给构造函数，它将解析成这个地址; 如果提供的是ENS名称，这将解析成对应的地址。

***contract*.deployTransaction ⇒ *TransactionResponse*：**如果**Contract**是ContractFactory部署后返回的对象， deployTransaction返	回的数据就是部署这个合约的交易信息（transaction）。

***contract*.interface ⇒ *Interface*：**这是作为Interface接口的ABI。

***contract*.provider ⇒ *Provider*：**如果构造函数使用的是provider生成的contract合约对象，那么这个结果就是这个provider，如果使用	的是具有Provider的signer，那么这个结果就是provider。

***contract*.signer ⇒ *Signer*：**如果构造函数使用的是signer生成的contract合约对象，那么这个结果就是这个signer。

# 方法

*contract*.**deployed**( ) ⇒ *Promise< 合约（Contract） >*

*Contract*.**isIndexed**( *value* ) ⇒ *boolean*

# 事件

***contract*.queryFilter( *event* [ , *fromBlockOrBlockHash* [ , *toBlock* ] ) ⇒ *Promise< Array< Event > >*：**返回与*event*匹配的事件

***contract*.listenerCount( [ *event* ] ) ⇒ *number*：**返回订阅该*event*的监听器数量。如果没有提供*event*，则返回所有事件的总数。

***contract*.listeners( *event* ) ⇒ *Array< Listener >*：**返回订阅该*event*的监听器列表。

***contract*.off( *event* , *listener* ) ⇒ *this*：**监听器取消订阅*event*事件。

***contract*.on( *event* , *listener* ) ⇒ *this*：**监听*event*事件，当事件发生时，会调用*listener*函数。

***contract*.once( *event* , *listener* ) ⇒ *this：***监听*event*事件，当事件发生时，仅调用一次*listener*函数。

***contract*.removeAllListeners( [ *event* ] ) ⇒ *this：***取消所有订阅*event*事件监听器，如果未提供*event*事件，则取消订阅所有事件的监听

# 元类(Meta-Class)

元类是在运行时确定其所有属性的类。**Contract**对象使用合约的ABI来确定可使用方法（methods）， 因此下面的部分将描述用一些属性，来与在合约构造函数期间交互的通用方法。

只读的方法 (常量; 如 view 或 pure)

常量方法是只读的，针对当前区块链状态计算少量EVM代码，并可以通过请求单个节点来计算，该节点会返回一个结果。 因此，它是免费的，不需要任何以太币，但不能更改区块链状态。

## 5.1 METHOD_NAME

*contract*.**METHOD_NAME**( ...*args* [ , *overrides* ] ) ⇒ *Promise< any >*

结果的类型取决于ABI。如果方法返回单个值，则将直接返回该值，否则将返回一个Result对象，其中包含每个位置可用的参数， 如果参数被命名，那么就按照命名后的值去执行。

如果返回的值匹配JavaScript中的类型值，如strings字符串类型 和 booleans布尔类型，那么返回的值类型就直接是JavaScript中的类型值。

但对于numbers类型，如果类型在JavaScript的安全范围内(即小于53位，如`int24`或`uint48`)，则使用标准的JavaScript number类型。 否则返回大数(BigNumber)类型。

对于字节bytes类型(包括固定长度和动态)，将返回一个DataHexString。

如果call reverts(或runs out of gas)，将抛出一个CALL_EXCEPTION，其中包括:

- `error.address` - 合约地址
- `error.args` - 合约方法中传入的参数
- `error.transaction` - 交易

只读方法的*overrides*对象可以包括以下任何一个:

- `overrides.from` - 代码执行期间使用的`msg.sender` (或 `CALLER`)
- `overrides.value` - 代码执行期间使用的`msg.value` (或 `CALLVALUE`)
- `overrides.gasPrice` - 每个gas的价格(理论上);因为没有交易，所以不会收取任何费用，但EVM仍然需要向`tx.gasprice`(或`GASPRICE`)传递value值; *大多数开发人员不需要这样做*

- `overrides.gasLimit` - 在执行代码期间允许节点使用的gas数量(理论上);因为没有交易，所以不会有任何费用，但EVM仍然估计gas总数量，因此会向`gasleft` (或 `GAS`) 传递这些值。

- `overrides.blockTag` - 一个用来模拟在哪里执行的块标签，可以用于假设性历史分析;注意，许多后端不支持这一点，或者可能需要付费访问，因为节点数据库存储和处理需求费用要高得多

## 5.2 *functions*.METHOD_NAME

*contract*.*functions*.**METHOD_NAME**( ...*args* [ , *overrides* ] ) ⇒ *Promise< Result >*

结果将始终是Result，即使只有一个返回值类型。

这简化了期望使用合约（Contract）对象的框架，因为它们不需要检查返回类型来展开简化函数。

此方法的另一个用途是用于错误恢复。例如，如果一个函数的结果是一个无效的UTF-8字符串，使用上述元类函数的普通调用将抛出一个异常。 这允许使用Result access error来访问低级字节以及错误的原因，从而允许使用另一种UTF-8错误策略。

大多数开发人员不需要这样做。

*overrides*与上面的只读操作相同。

### Write Methods (non-constant)

非常量方法要求签名一笔交易，并需要向矿工支付费用。这个交易将由整个网络上的所有节点以及矿工进行验证， 矿工会根据当前状态执行，接着在区块链计算新状态。

它不能返回结果。如果需要一个结果，那么应该使用Solidity事件(或EVM日志)对其进行记录，然后可以从交易收据中查询该事件。

## 5.3 METHOD_NAME

*contract*.**METHOD_NAME**( ...*args* [ , *overrides* ] ) ⇒ *Promise< TransactionResponse >*

交易被发送到网络后，返回交易的TransactionResponse。 **Contract**合约对象需要signer。

重写方法的*overrides*对象可以是以下任何一个:

- `overrides.gasPrice` - 每个gas的价格
- `overrides.gasLimit` - 该笔交易允许使用的gas的最大数量，未被使用的gas按照gasPrice退还
- `overrides.value` - 调用中转账的ether (wei格式)数量
- `overrides.nonce` - Signer使用的nonce值

如果返回的TransactionResponse使用了`wait()`方法，那么在收据上将会有额外的属性:

- `receipt.events` - 带有附加属性的日志数组(如果ABI包含事件的描述)
- `receipt.events[n].args` - 解析后的参数
- `receipt.events[n].decode` - 可以用来解析日志主题（topics）和数据的方法(用于计算`args`)
- `receipt.events[n].event` - 事件的名称
- `receipt.events[n].eventSignature` - 这个事件完整的签名
- `receipt.removeListener()` - 用于移除触发此事件的监听器的方法
- `receipt.getBlock()` - 返回发生触发此事件的Block
- `receipt.getTransaction()` - 返回发生触发此事件的Transaction
- `receipt.getTransactionReceipt()` - 返回发生触发此事件的Transaction Receipt

## Write Methods Analysis

在不实际执行的情况下，有几个选项可以分析write method的属性和结果。

## 5.4 estimateGas.METHOD_NAME

*contract*.*estimateGas*.**METHOD_NAME**( ...*args* [ , *overrides* ] ) ⇒ *Promise< 大数(BigNumber) >*

返回使用 *args*和*overrides*执行*METHOD_NAME*所需的gas的估计单位。

*overrides*与上面针对只读或写方法的overrides相同，具体取决于 *METHOD_NAME*调用的类型。

## 5.5 populateTransaction.METHOD_NAME

*contract*.*populateTransaction*.**METHOD_NAME**( ...*args* [ , *overrides* ] ) ⇒ *Promise< UnsignedTx >*

返回一个未签名交易(UnsignedTransaction)，它表示需要签名并提交给网络的交易，以执行带有*args*和*overrides*的*METHOD_NAME*。

*overrides*与上面针对只读或写方法的overrides相同，具体取决于 *METHOD_NAME*调用的类型。

## 5.6 callStatic.METHOD_NAME

*contract*.*callStatic*.**METHOD_NAME**( ...*args* [ , *overrides* ] ) ⇒ *Promise< any >*

比起执行交易状态更改，更可能是会要求节点*尝试*调用不进行状态的更改并返回结果。

这实际上并不改变任何状态，而是免费的。在某些情况下，这可用于确定交易是失败还是成功。

otherwise函数与Read-Only Method相同。

*overrides*与上面的只读操作相同。

## Event Filters

事件过滤器由主题（topics）组成，这些主题是Bloom Filter中记录的值，允许对匹配过滤器的条目进行有效搜索。

## 5.7 filters.EVENT_NAME

*contract*.*filters*.**EVENT_NAME**( ...*args* ) ⇒ *Filter*

返回*EVENT_NAME*的过滤器，可以通过增加其他约束进行过滤。

只有`indexed`索引的事件参数可以被过滤。如果参数为空(或未提供)，则该字段中的任何值都匹配。

# Example:ERC-20 Contract

## 6.1 部署合约

**new *ethers*.ContractFactory( *abi* , *bytecode* , *signer* )**：创建一个新的ContractFactory，它可以将合约部署到区块链。

```js
const bytecode = "0x60806040523480156100105760008......";

// 一个人类可读（Human-Readable）的ABI;我们只需要指定相关的部分内容，在这里我们用到了构造函数
const abi = [
    "constructor(uint totalSupply)"
];

const factory = new ethers.ContractFactory(abi, bytecode, signer)

// 部署，设置token的total supply为100，分配给部署人（deployer）
const contract = await factory.deploy(parseUnits("100"));

//链上现在虽然我们还没有这个合约，但是这个合约地址我们是预先经过计算后得到的
contract.address
// '0xa22aB6748282B3125dC26dAFb79e38B7eb24EAcC'

// 在交易被挖出来之前(即合约被部署)等待（Wait）
//  - 部署后，交易会返回receipt
await contract.deployTransaction.wait();
// {
//   blockHash: '0xe9b244958d066490c46a826a9733c1a43210316777185353b3ecbc2ec362ea87',
//   blockNumber: 22,
//   byzantium: true,
//   confirmations: 1,
//   contractAddress: '0xa22aB6748282B3125dC26dAFb79e38B7eb24EAcC',
	 ......
// }
```

## 6.2 **Meta-Class Methods**

> 因为合约是元类，这里面可用的方法取决于传入**合约**的ABI

***erc20*.decimals( [ *overrides* ] ) ⇒ *Promise< number >*：**返回此ERC-20 token所使用的小数位数。在前端界面，当从用户获取输入时，	可以使用parseUnits转化后传入合约； 从合约获得token后可以通过[formatUnits](utils-formatunits]转化后再显示给用户。

```js
await erc20.decimals(); 
// 18
```

***erc20*.balanceOf( *owner* [ , *overrides* ] ) ⇒ *Promise< 大数(BigNumber) >*：**返回持有这个ERC-20 token的*持有者(owner)*的余额。

```js
await erc20.balanceOf(signer.getAddress()) 
// { BigNumber: "100000000000000000000" }
```

***erc20*.symbol( [ *overrides* ] ) ⇒ *Promise< string >*：**返回token的symbol。

```js
await erc20.symbol(); 
// 'MyToken'
```

***erc20_rw*.transfer( *target* , *amount* [ , *overrides* ] ) ⇒ *Promise< TransactionResponse >*：**从当前的signer将数量为*amount*的tokens	转给接收者*target*。 在交易处于写入操作时，返回值(布尔类型)是得不到的。 如果需要这个值，则需要其他方法(如事件)。链上合约调	用`transfer`函数可以得到这个结果。

```js
// 在这之前先转化格式
formatUnits(await erc20_rw.balanceOf(signer.getAddress()));
// '100.0'

// 转 1.23 tokens 到 ENS name 为 "ricmoo.eth" 的地址
tx = await erc20_rw.transfer("ricmoo.eth", parseUnits("1.23"));
//TransactionRequest

// 等待交易所在的区块被挖出打包
await tx.wait();
//TransactionResponse 

// 成功后转化格式
formatUnits(await erc20_rw.balanceOf(signer.getAddress()));
// '98.77'

formatUnits(await erc20_rw.balanceOf("ricmoo.eth"));
// '1.23
```

***erc20*.*callStatic*.transfer( *target* , *amount* [ , *overrides* ] ) ⇒ *Promise< boolean >*:**执行一次从当前signer向*target*地址转移*amount*数量的token的演练，而不实际签名或发送交易,这可以用于检查真实的转账前，交易是否能成功。

***erc20*.*estimateGas*.transfer( *target* , *amount* [ , *overrides* ] ) ⇒ *Promise< 大数(BigNumber) >*:**返回估计的“将*amount*数量的tokens发送*target*地址”所需的多少gas值。

***erc20*.*populateTransaction*.transfer( *target* , *amount* [ , *overrides* ] ) ⇒ *Promise< UnsignedTx >*:**返回一个未签名交易，它可以被签名并提交给网络，达成“将*amount*数量的tokens发送*target*地址”的目的。

## 6.3 Meta-Class Fileters

***erc20*.*filters*.Transfer( [ *fromAddress* [ , *toAddress* ] ] ) ⇒ *Filter*：**返回一个新的Filter用于查询或者去subscribe/unsubscribe to 	events，如果*fromAddress* 是空的或者没有填写，将匹配任何发送人的地址。 如果 *toAddress*是空的或者没有填写，将匹配任何接受人的地址。

```js
filterFrom = erc20.filters.Transfer(signer.address);
// {
//   address: '0xa22aB6748282B3125dC26dAFb79e38B7eb24EAcC',
//   topics: [
//     '0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef',
//     '0x000000000000000000000000894ed91b666facce5a4d2ff8261924b4754a5759'
//   ]
// }

// 在最新的10个区块内查询我发送的交易
logsFrom = await erc20.queryFilter(filterFrom, -10, "latest");
//Array<TransactionResponse>
```

```js
// 监听signer的发送的event事件:
erc20.on(filterFrom, (from, to, amount, event) => {
  // The `from` will always be the signer address
});

// 监听signer的接受转账的event事件:
erc20.on(filterTo, (from, to, amount, event) => {
  // The `to` will always be the signer address
});

// 监听所有的转账交易event事件:
erc20.on("Transfer", (from, to, amount, event) => {
  // ...
});
```


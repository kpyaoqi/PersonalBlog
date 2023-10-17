---
title: 5-Singers

date: 2022-08-14	

categories: ether@5.4.js	

tags: [区块链,ether@5.4.js]
---	

# Singer

**Signer**类是抽象的，**不能直接实例化，而是应该使用一个具体的子类**，如 Wallet, VoidSigner 或 JsonRpcSigner。

## 1.1 属性

***signer*.connect( *provider* ) ⇒ *Signer***:子类**必须**实现这个，如果更改后的providers是不被支持的话，它们可能仅仅只抛出一个错误。

***signer*.getAddress( ) ⇒ *Promise< string< 地址(Address) > >***:返回一个解析为帐户地址的Promise。

​	子类**必须**实现这个,这是一个Promise，因此一个**Signer**可以围绕一个异步源进行设计，如硬钱包。

***Signer*.isSigner( *object* ) ⇒ *boolean***:当且仅当 *object* 是一个**Signer**时返回true。

## 1.2 **Blockchain** 方法

### 1.2.1 getBalance

*signer*.**getBalance**( [ *blockTag* = "*latest*" ] ) ⇒ *Promise< 大数(**BigNumber**) >*

在指定的*blockTag*下返回这个钱包的余额。

### 1.2.2  getChainId

*signer*.**getChainId**( ) ⇒ *Promise< **number** >*

返回这个钱包连接的链 ID。

### 1.2.3 getGasPrice

*signer*.**getGasPrice**( ) ⇒ *Promise< 大数(**BigNumber**) >*

返回当前的gas price。

### 1.2.4 getTransactionCount

*signer*.**getTransactionCount**( [ *blockTag* = "*latest*" ] ) ⇒ *Promise< **number** >*

返回此帐户曾经发送的交易数量，交易中的中nonce依赖这个值。

### 1.2.5 call

*signer*.**call**( *transactionRequest* ) ⇒ *Promise< string< **DataHexString** > >*

返回*transactionRequest*调用的结果，此帐户地址用作`from` 字段。

### 1.2.6 estimateGas

*signer*.**estimateGas**( *transactionRequest* ) ⇒ *Promise< 大数(**BigNumber**) >*

返回发送*transactionRequest*的估算费用，在使用这个方法之前先用账户地址填充`from`字段。

### 1.2.7 resolveName

*signer*.**resolveName**( *ensName* ) ⇒ *Promise< string< 地址(**Address**) > >*

返回与*ensName*关联的地址。

## 1.3 **Signing**方法

### 1.3.1 signMessage

*signer*.**signMessage**( *message* ) ⇒ *Promise< string< **RawSignature** > >*

这将返回一个解析为消息的Raw Signature Promise。

由于使用的是hashMessage方法，因此它是EIP-191兼容的。 如果在Solidity中恢复地址，则需要这个前缀来创建一个匹配的哈希。

子类**必须**实现这个方法。 如果不支持签名消息可能会抛出错误，比如在基于合约的钱包或基于元交易的钱包中使用时。

### 1.3.2 signTransaction

*signer*.**signTransaction**( *transactionRequest* ) ⇒ *Promise< string< **DataHexString** > >*

返回一个解析为*transactionRequest*中**已签名的交易**的Promise。 此方法不填充任何缺少的字段。

子类**必须**实现这个，如果不支持签名消息可能会抛出错误，出于安全这在许多客户端中是常见的。

### 1.3.3 sendTransaction

*signer*.**sendTransaction**( *transactionRequest* ) ⇒ *Promise< **TransactionResponse** >*

该方法使用populate Transaction填充缺少字段的transactionRequest，并返回一个解析成交易的Promise。

### 1.3.4 _signTypedData

*signer*.**_signTypedData**( *domain* , *types* , *value* ) ⇒ *Promise< string< **RawSignature** > >*

> 实验中的功能。如果使用它，请**指定**您正在使用的ethers的确切版本(例如指定`"5.0.18"`，而不是`"^5.0.18"`)， 因为方法名将从`_signTypedData`重命名为`signTypedData`。

## 1.4 **Sub-Classes**

**Signer**的所有重要属性都是不可变的，这一点非常重要。由于以太坊是异步的，并处理关键数据(如ether和其他潜在有价值的加密资产)， 整个Signer的生命周期中保持*provider*和 *address*等属性是静态的有助于防止严重的问题的出现， 而且许多其他类和库也是认定*provider*和 *address*等属性是静态的。

子类**必须**扩展Signer，并且**必须**调用`super()`。

### 1.4.1 **checkTransaction**

*signer*.**checkTransaction**( *transactionRequest* ) ⇒ ***TransactionRequest***

这应该返回一个*transactionRequest*的副本，包含`call`、`estimateGas`和`populateTransaction` (sendTransaction使用的)所需的任何属性。 如果指定了任何未知的key，它会抛出一个错误。

默认的实现只**验证有效的TransactionRequest属性是否存在**，如果不存在，则将`from`添加到交易中。

如果存在`from`字段，则**必须**验证它与 Signer的地址是否相等。

### 1.4.2 **populateTransaction**

*signer*.**populateTransaction**( *transactionRequest* ) ⇒ *Promise< **TransactionRequest** >*

这应该返回一个*transactionRequest*的副本，遵循与`checkTransaction`相同的过程， 并填写发送交易所需的任何属性。

返回的结果都是promises，可以使用**resolvePropertiesutility**函数来解析。

默认实现调用`checkTransaction`，如果它是一个ENS name就会解析它，并根据Signer上的相关操作添加`gasPrice`, `nonce`, `gasLimit`和`chainId`

# Wallet

Wallet类继承了Signer，可以使用私钥作为外部拥有帐户(EOA)的标准对交易和消息进行签名

**new *ethers*.Wallet( *privateKey* [ , *provider* ] )**

为*privateKey*创建一个新的钱包实例，并可选地连接到*provider*。

## 2.1 方法

### 2.1.1 createRandom

*ethers*.*Wallet*.**createRandom**( [ *options* = {} ] ) ⇒ ***Wallet***

返回一个带有随机私钥的新钱包，由加密安全的熵源生成。如果当前环境没有安全的熵源，则会抛出错误。

使用此方法创建的钱包将具有助记词。

### 2.1.2 fromEncryptedJson

*ethers*.*Wallet*.**fromEncryptedJson**( *json* , *password* [ , *progress* ] ) ⇒ *Promise< **Wallet** >*

从加密的JSON钱包创建一个实例。

如果提供了*进度*，它将在解密期间被调用，其值介于0到1之间，表示一个完成进度。

### 2.1.3 fromEncryptedJsonSync

*ethers*.*Wallet*.**fromEncryptedJsonSync**( *json* , *password* ) ⇒ ***Wallet***

从加密的JSON钱包创建一个实例。

此操作将同步操作，从而锁定用户界面一段时间。 大多数应用程序应该使用异步的`fromEncryptedJson`。

### 2.1.4 fromMnemonic

*ethers*.*Wallet*.**fromMnemonic**( *mnemonic* [ , *path* , [ *wordlist* ] ] ) ⇒ ***Wallet***

从助记短语中创建实例如果没有指定path，则使用的默认path路径(`m/44'/60'/0'/0/0`),如果不指定wordlist，则使用English Wordlist

### 2.1.5 encrypt

*wallet*.**encrypt**( *password* , [ *options* = {} , [ *progress* ] ] ) ⇒ *Promise< **string** >*

加密钱包，使用*password*返回一个解析为JSON钱包的Promise。

如果提供了*进度*，它将在解密期间被调用，其值介于0到1之间，表示一个完成进度。

## 2.2 属性

***wallet*.address ⇒ *string< 地址(Address) >*：**此钱包表示的帐户的地址。

***wallet*.provider ⇒ *Provider*：**这个钱包所连接的provider，它将用于任何Blockchain Methods的方法，它也可以是null。

> 一个**钱包**实例是不可变的，因此如果您希望更改Provider，您可以使用connect方法创建一个连接到所需的provider的新实例。

***wallet*.publicKey ⇒ *string< DataHexString< 65 > >***：此钱包的未压缩的公钥

```js
// 通过助记词创建一个钱包实例
mnemonic = "announce room limb pattern dry unit scale effort smooth jazz weasel alcohol"
walletMnemonic = Wallet.fromMnemonic(mnemonic)

// 或者通过私钥创建一个钱包实例
walletPrivateKey = new Wallet(walletMnemonic.privateKey)

walletMnemonic.address === walletPrivateKey.address
// true

// 有着每个Signer API的地址，是一个Promise对象
await walletMnemonic.getAddress()
// '0x71CB05EE1b1F506fF321Da3dac38f25c0c9ce6E1'

// 一个钱包地址，也可用同步的方式获得
walletMnemonic.address
// '0x71CB05EE1b1F506fF321Da3dac38f25c0c9ce6E1'

// 内部加密组件
walletMnemonic.privateKey
// '0x1da6847600b0ee25e9ad9a52abbd786dd2502fa4005dd5af9310b7cc7a3b25db'
walletMnemonic.publicKey
// '0x04b9e72dfd423bcf95b3801ac93f4392be5ff22143f9980eb78b3a860c4843bfd04829ae61cdba4b3b1978ac5fc64f5cc2f4350e35a108a9c9a92a81200a60cd64'

// 钱包助记词
walletMnemonic.mnemonic
// {
//   locale: 'en',
//   path: "m/44'/60'/0'/0/0",
//   phrase: 'announce room limb pattern dry unit scale effort smooth jazz weasel alcohol'
// }

// Note: 通过私钥创建的钱包实例没有助记词，因为从数学上无法推导
walletPrivateKey.mnemonic
// null

// 签名消息
await walletMnemonic.signMessage("Hello World")
// '0x14280e5885a19f60e536de50097e96e3738c7acae4e9e62d67272d794b8127d31c03d9cd59781d4ee31fb4e1b893bd9b020ec67dfa65cfb51e2bdadbb1de26d91c'

tx = {
  to: "0x8ba1f109551bD432803012645Ac136ddd64DBA72",
  value: utils.parseEther("1.0")
}

// 签名交易
await walletMnemonic.signTransaction(tx)
// '0xf865808080948ba1f109551bd432803012645ac136ddd64dba72880de0b6b3a7640000801ca0918e294306d177ab7bd664f5e141436563854ebe0a3e523b9690b4922bbb52b8a01181612cec9c431c4257a79b8c9f0c980a2c49bb5a0e6ac52949163eeb565dfc'

// connect方法返回一个连接到provider的新钱包实例
wallet = walletMnemonic.connect(provider)

// 从网络中查询
await wallet.getBalance();
// { BigNumber: "6846" }
await wallet.getTransactionCount();
// 3

// 发送 ether
await wallet.sendTransaction(tx)
// {
//   accessList: [],
//   chainId: 31337,
//   confirmations: 0,
//   data: '0x',
//   from: '0x894ed91B666FacCe5a4D2FF8261924b4754A5759',
//   gasLimit: { BigNumber: "21001" },
//   gasPrice: null,
//   hash: '0x1b4a95e9d23bd96d5429c535148eb3eaed326d89118b118e44cc05c26703224e',
//   maxFeePerGas: { BigNumber: "1572346688" },
//   maxPriorityFeePerGas: { BigNumber: "1500000000" },
//   nonce: 5,
//   r: '0x5e32c2741700a9120cfa186bc74e88b3d9488393be796a5124fddf08ffdbfdc6',
//   s: '0x2431f3e3274cd22d6de63ed6e23a5c6839c1fabeb97b6683fb15584b9bf1f29d',
//   to: '0x8ba1f109551bD432803012645Ac136ddd64DBA72',
//   type: 2,
//   v: 0,
//   value: { BigNumber: "1000000000000000000" },
//   wait: [Function]
// }
```

# VoidSinger

一个**VoidSigner**是一个简单的Signer，它不能签名。

当API需要signer作为参数时，它作为只读的signer是有用的，但它只能携带只读的操作。

比如，在`call`函数调用期间会自动传递所提供的地址。

**new *ethers*.VoidSigner( *address* [ , *provider* ] ) ⇒ *VoidSigner*：**为一个地址创建**VoidSigner**实例。

***voidSigner*.address ⇒ *string< 地址(Address) >***：**VoidSigner**的地址。

# **ExternallyOwnedAccount**

这个接口包含外部拥有帐户（EOA）所需的最小属性集，可以执行某些操作，比如将其编码为JSON钱包。

***eoa*.address ⇒ *string< 地址(Address) >：***地址(Address)EOA的地址

***eoa*.privateKey ⇒ *string< DataHexString< 32 > >：***EOA的私钥

***eoa*.mnemonic ⇒ *助记词*：**帐户HD的助记词，如果有的话可以打印出来。EOA账户源不编码助记符，如HD extended keys。
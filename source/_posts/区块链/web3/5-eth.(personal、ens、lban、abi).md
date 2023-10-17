---
title: 5-eth.(personal、ens、lban、abi)

date: 2022-09-05	

categories: web3.js	

tags: [区块链,web3.js]
---	

# personal(同节点上的账户进行交互)

## 1.1 newAccount

创建一个新的账户。

> 永远不要通过不安全的 Websocket 或 HTTP 服务提供器来调用这些函数，因为你的密码是明文发送的！

**参数：**`password` - `String`: 用来加密账户的密码。

**返回值(Promise<string>)：**新创建账户地址

## 1.2 sign

该方法通过下面的方式计算一个以太坊特定签名：

```js
sign(keccak256("\x19Ethereum Signed Message:\n" + dataToSign.length + dataToSign)))
```

在消息前加个前缀使得算出的签名可以被识别为以太坊特定签名。

如果你同时有原始消息和签名消息，就可以使用 web3.eth.personal.ecRecover 来恢复签名账户地址

> 通过不安全的 HTTP RPC 连接发送帐户密码非常危险。

**参数：**

1. `String` - 要签名的数据。 如果是字符串会使用 web3.utils.utf8ToHex 将其转换为 16 进制。
2. `String` - 用来签名的账户地址。
3. `String` - 用来签名的账户密码。
4. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为签名结果。

**返回值(Promise<string>)：**签名字符串

## 1.3 ecRecover

恢复数据签名帐户

**参数:**

1. `String` - 被签名的数据。 如果是字符串会使用 web3.utils.utf8ToHex 将其转换为 16 进制。
2. `String` - 签名。
3. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为签名结果。

**返回值:**`Promise<string>` - 签名账户。

## 1.4 signTransaction

对交易进行签名，账户必须先解锁。

> 通过不安全的 HTTP RPC 连接发送帐户密码非常危险。

**参数：**

1. `Object` - 要签名的交易数据，更多详情请看 web3.eth.sendTransaction() 。
2. `String` - 用来签名交易的 `from` 账户密码。
3. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为签名结果。

**返回值(Promise<Object>)：**RLP 编码的交易对象，其 `raw` 属性可以用来通过 web3.eth.sendSignedTransaction 来发送交易。

## 1.5 sendTransaction

该方法用来通过账户管理 API 来发送交易。

> 通过不安全的 HTTP RPC 连接发送帐户密码非常危险。

**参数：**

1. `Object` - 交易对象属性
2. `String` - 当前帐户的密码
3. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为签名结果。

**返回值(Promise<string>)：**交易哈希

## 1.6 unlockAccount

解锁账户

> 通过不安全的 HTTP RPC 连接发送帐户密码非常危险。

**参数：**

1. `address` - `String`: 要解锁的账户地址。
2. `password` - `String` - 账户密码。
3. `unlockDuration` - `Number` - 将帐户保持在解锁状态的持续时间。

## 1.7 lockAccount

锁定给定帐户。

> 通过不安全的 HTTP RPC 连接发送帐户密码非常危险。

**参数：**

1. `address` - `String`: 要锁的账户地址。
2. `Function` - (可选) 可选的回调函数，其第一个参数为错误对象，第二个参数为签名结果。

**返回值(Promise<boolean>)**

## 1.8 getAccounts

通过使用服务提供器并调用 RPC 方法 `personal_listAccounts` 返回节点控制的账户列表。使用 web3.eth.accounts.create() 创建的账户不会被添加到这个列表中。这方面的更多信息可以查看 web3.eth.personal.newAccount()。

结果和 web3.eth.getAccounts() 是一样的，只是它用的 RPC 方法是 `eth_accounts`。

**返回值(Promise<Array>)**:节点控制地址数组

## 1.9 importRawKey

将给定的私钥导入密钥存储区，并使用密码对其进行加密。返回和导入私钥对应的新账户地址。

> 通过不安全的 HTTP RPC 连接发送帐户密码非常危险。

**参数：**

1. `privateKey` - `String` - 为加密的私钥 (16 进制字符串)。
2. `password` - `String` - 账户密码。

**返回值(Promise<string>)**:账户地址

# ens(与 ENS 进行交互)

# lban(将以太坊地址和 IBAN/BBAN 地址之间相互转换)

# abi(解码及编码为 ABI用于EVM进行函数调用)


---
title: 7-private私有变量

date: 2023-05-06	

categories: 合约安全	

tags: [区块链,合约安全]
---	

# private私有变量

私有变量在区块链上仍然是可见的，所以敏感信息不应该被存储在那里。如果它们不能被访问，验证者如何能够处理取决于其值的交易？私有变量不能从外部的Solidity 合约中读取，但它们可以使用以太坊客户端在链外读取

> 标注为 `private` 区域的数据并不是不能访问，它们存储在一个又一个的 `slot`存储槽里

### 例子：

要读取一个变量，你需要知道它的存储槽。在下面的例子中，myPrivateVar的存储槽是0。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract PrivateVarExample {
    uint256 private myPrivateVar;

    constructor(uint256 _initialValue) {
        myPrivateVar = _initialValue;
    }
}
```

下面是读取已部署的智能合约的私有变量的javascript代码

```js
async function readPrivateVar() {
  ...
  const privateVarValue = await web3.eth.getStorageAt(
    PrivateVarExample_ADDRESS,
    0
  );
  ...
  console.log("Value of private variable 'myPrivateVar':",
  web3.utils.hexToNumberString(privateVarValue));
}
```
---
title: Truffle

date: 2023-03-03	

categories: 合约框架	

tags: [区块链,区块链知识点总结,合约框架]
---	

# Truffle 命令

### 初始化项目

- `truffle init`：在当前目录下创建新的 Truffle 项目。
- `truffle unbox <box-name>`：下载指定的 Truffle box 模板，以快速创建一个新项目。

### 编译和部署智能合约

- `truffle compile`：编译智能合约。
- `truffle migrate`：部署智能合约。
- `truffle migrate --reset`：重新部署智能合约。

```js
var MyContract = artifacts.require("XlbContract");
module.exports = function(deployer) {
  // 部署步骤
  deployer.deploy(MyContract);
};
```

### 交互式控制台

- `truffle console`：启动交互式控制台，可在其中与智能合约进行交互。
- `truffle console --network <network-name>`：在指定的网络上启动交互式控制台。

```js
truffle(develop)> let instance = await MetaCoin.deployed()
truffle(develop)> let accounts = await web3.eth.getAccounts()
truffle(develop)> let result = instance.sendCoin(accounts[1], 10, {from: accounts[0]})
//result.tx (string) - 交易哈希 hash
//result.logs (array) - 解码过的事件 (日志)
//result.receipt (object) - 交易收据 receipt（包括使用的gas）
```

### 测试

- `truffle test`：运行测试脚本。

### 编写外部脚本

- `truffle exec <path/to/file.js>`：运行外部脚本

### 其他命令

- `truffle version`：查看 Truffle 版本号。
- `truffle networks`：查看已配置的网络列表。
- `truffle help`：查看 Truffle 帮助文档。
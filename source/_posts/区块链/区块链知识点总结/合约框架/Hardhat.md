---
title: Hardhat

date: 2023-03-02	

categories: 合约框架	

tags: [区块链,区块链知识点总结,合约框架]
---	

# Hardhat 命令

### 初始化项目

- `npx hardhat`：在当前目录下创建新的 Hardhat 项目。
- `npx hardhat example`：在当前目录下创建新的 Hardhat 示例项目。
- `npx hardhat upgrade --network <network-name>`：升级当前项目到 Hardhat 最新版本。

### 编译和部署智能合约

- `npx hardhat compile`：编译智能合约。
- `npx hardhat run scripts/<script-name>.js --network <network-name>`：部署智能合约。
- `npx hardhat node`：在本地启动一个节点，可用于本地测试。

### 交互式控制台

- `npx hardhat console`：启动交互式控制台，可在其中与智能合约进行交互。
- `npx hardhat console --network <network-name>`：在指定的网络上启动交互式控制台。

### 测试

- `npx hardhat test`：运行测试脚本。

### 其他命令

- `npx hardhat clean`：删除构建文件和缓存文件。
- `npx hardhat accounts`：列出钱包地址。
- `npx hardhat config`：查看 Hardhat 配置信息。
- `npx hardhat help`：查看 Hardhat 帮助文档。


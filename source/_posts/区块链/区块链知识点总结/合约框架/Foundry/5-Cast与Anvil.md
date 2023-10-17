---
title: 5-Cast与Anvil

date: 2023-02-21	

categories: Foundry	

tags: [区块链,区块链知识点总结,合约框架,Foundry]
---	

# Cast 概述

Cast 是 Foundry 用于执行以太坊 RPC 调用的命令行工具。 您可以进行智能合约调用、发送交易或检索任何类型的链数据——所有这些都来自您的命令行！

## 如何使用 Cast

要使用 Cast，请运行 cast 命令，然后运行子命令：

```bash
$ cast <subcommand>
```

### 例子

让我们使用 cast 来检索 DAI 代币的总供应量：

```bash
$ cast call 0x6b175474e89094c44da98b954eedeac495271d0f "totalSupply()(uint256)" --rpc-url https://eth-mainnet.alchemyapi.io/v2/Lc7oIGYeL_QvInzI0Wiu_pOZZDEKBrdf
8603853182003814300330472690
```

cast 还提供了许多方便的子命令，例如用于解码 calldata：

```bash
$ cast 4byte-decode 0x1F1F897F676d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003e7
1) "fulfillRandomness(bytes32,uint256)"
0x676d000000000000000000000000000000000000000000000000000000000000
999
```

您还可以使用 cast 发送任意消息。 下面是在两个 Anvil 帐户之间发送消息的示例。

```bash
$ cast send --private-key <Your Private Key> 0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc $(cast --from-utf8 "hello world") --rpc-url http://127.0.0.1:8545/
```

## 命令

### General 命令

cast help   获取 Cast 命令的帮助。

cast completions   生成 shell 自动补全脚本。

### Chain 命令

cast chain-id   获取 Ethereum 的链 ID。

cast chain   获取当前链的名称。

cast client   获取当前客户端的版本。

### Transaction 命令

cast publish   向网络发布一个原始交易。

cast receipt   获取一个交易的交易收据。

cast send   签署并发布一项交易。

cast call   在不发布交易的情况下对一个账户进行调用。

cast rpc   执行一个原始的 JSON-RPC 请求 [aliases: rp]

cast tx   获得有关交易的信息。

cast run   在本地环境中运行一个已发布的交易，并打印出跟踪。

cast estimate   估算交易的 Gas 成本。

cast access-list   为一个交易创建一个访问列表。

### Block 命令

cast find-block   获取与提供的时间戳最接近的区块编号。

cast gas-price   获取当前 Gas 价格。

cast block-number   获取最新的区块号。

cast basefee   获取一个区块的基础费用。

cast block   获取一个区块的信息。

cast age   获取一个区块的时间戳。

### Account 命令

cast balance   获取一个账户的余额，单位为 Wei。

cast storage   获取一个合约的存储槽的原始值。

cast proof   为一个给定的存储槽生成一个存储证明。

cast nonce   获取一个账户的 nonce。

cast code   获取一个合约的字节码。

### ENS 命令

cast lookup-address   进行 ENS 反向查询。

cast resolve-name   进行 ENS 查询。

cast namehash   计算一个名字的 ENS namehash。

### Etherscan 命令

cast etherscan-source   从 Etherscan 获取合约的源代码。

### ABI 命令

cast abi-encode   对给定的函数参数进行 ABI 编码，不包括选择器。

cast 4byte   从 https://sig.eth.samczsun.com 中获取指定选择器的函数签名。

cast 4byte-decode   使用 https://sig.eth.samczsun.com 对 ABI 编码的 calldata 进行解码。

cast 4byte-event   从 https://sig.eth.samczsun.com 中获取 topic 0 的事件签名。

cast calldata   ABI 编码一个带参数的函数。

cast pretty-calldata   漂亮地打印 Calldata。

cast --abi-decode   解码 ABI 编码的输入或输出数据。

cast --calldata-decode   解码 ABI 编码的输入数据。

cast upload-signature   将指定的签名上传到 https://sig.eth.samczsun.com.

### Conversion 命令

cast --format-bytes32-string   将一个字符串转换成 bytes32 编码。

cast --from-bin   将二进制数据转换为十六进制数据。

cast --from-fix   将一个定点数转换成一个整数。

cast --from-utf8   将 UTF8 文本转换为十六进制。

cast --parse-bytes32-string   从 bytes32 编码中解析出一个字符串。

cast --to-ascii   将十六进制数据转换为ASCII字符串。

cast --to-base   将一个进制底数转换为其它底数。

cast --to-bytes32   右移十六进制数据至 32 字节。

cast --to-fix   将一个整数转换成一个定点数。

cast --to-hexdata   将输入规范化为小写，0x- 前缀的十六进制。

cast --to-int256   将一个数字转换为十六进制编码的 int256。

cast --to-uint256   将一个数字转换成十六进制编码的 uint256。

cast --to-unit   将一个 eth 单位转换为另一个单位。 (ether, gwei, wei).

cast --to-wei   将 eth 金额转换为 wei 单位。

cast shl   进行左移操作。

cast shr   进行右移操作。

### Utility Commands

cast sig   获取一个函数的选择器。

cast keccak   使用 keccak-256 对任意数据进行哈希。

cast compute-address   从给定的 nonce 和部署者地址计算合约地址。

cast interface   从一个给定的 ABI 生成一个 Solidity 接口。

cast index   计算集合中条目的存储插槽位置。

cast --concat-hex   串接十六进制字符串。

cast --max-int   获取 int256 最大值。

cast --min-int   获取 int256 最小值。

cast --max-uint   获取 uint256 最大值。

cast --to-checksum-address   将一个地址转换为校验过的格式 (EIP-55).

### Wallet Commands

cast wallet   钱包管理实用工具。

cast wallet new   创建一个新的随机密钥对。

cast wallet address   将一个私钥转换为一个地址。

cast wallet sign   签署消息。

cast wallet vanity   生成一个虚构的地址。

cast wallet verify   验证一个信息的签名。

# Anvil 概述

**Anvil 是 Foundry 附带的本地测试网节点。** 您可以使用它从前端测试您的合约或通过 RPC 进行交互。

Anvil 是 Foundry 套件的一部分，与 forge 和 cast 一起安装。

> 注意：如果您安装了旧版本的 Foundry，则需要重新安装 foundryup 才能下载 Anvil。

## 如何使用 Anvil

要使用 Anvil，只需输入 anvil。 您应该会看到可用的帐户和私钥列表，以及节点正在侦听的地址和端口。

Anvil 是高度可配置的。 您可以运行 anvil -h 查看所有配置选项。

一些基本选项是：

```bash
# 要生成和配置的用户的帐户数[default: 10]
anvil -a, --accounts <ACCOUNTS>
# 要使用的EVM硬分叉[default: latest]
anvil --hardfork <HARDFORK>
# 要侦听的端口号[default: 8545]
anvil  -p, --port <PORT>
```


---
title: 4-Gas追踪

date: 2023-02-20	

categories: Foundry	

tags: [区块链,区块链知识点总结,合约框架,Foundry]
---	

# Gas 追踪

Forge 可以帮助您估算您的合约将消耗多少 gas。

目前，Forge 为这项工作提供了两种不同的工具，但它们可能会在未来合并：

- Gas reports：**Gas 报告让您大致了解到 Forge 认为 你的合约中的各个函数消耗 gas 的概况**。
- Gas snapshots：**Gas 快照让您大致了解 每次测试消耗 Gas 的概况。**

Gas 报告和 Gas 快照在某些方面有所不同：

- **Gas 报告使用跟踪来计算单个合约调用的 Gas 成本**。 这以速度为代价提供了更精细的洞察力。
- **Gas 快照具有更多内置工具，例如将差异和结果导出到文件。** 快照**不像 Gas 报告那样精细，但生成速度更快。**

## Gas 报告

Forge 可以为您的合约生成 Gas 报告。 您可以**通过 foundry.toml 中的 gas_reports 字段配置哪些合约输出 Gas 报告**。

为特定合约生成报告：

```toml
gas_reports = ["MyContract", "MyContractFactory"]
```

为所有合约生成报告：

```toml
gas_reports = ["*"]
```

要生成 Gas 报告，请**运行 forge test --gas-report**。

您还可以**将它与其他子命令结合使用，例如 forge test --match-test testBurn --gas-report，以仅生成与此测试相关的 Gas 报告。**

示例输出：

<img src="/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/区块链知识点总结/合约框架/Foundry/img/image-20230621205206454.png" alt="image-20230621205206454" style="zoom:80%;" /> 

您还可以通过 foundry.toml 中的 gas_reports_ignore 字段忽略合约：

```toml
gas_reports_ignore = ["Example"]
```

这会将输出更改为：

<img src="/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/区块链知识点总结/合约框架/Foundry/img/image-20230621205342754.png" alt="image-20230621205342754" style="zoom:80%;" /> 

## Gas快照

Forge 可以为您的所有测试函数生成 Gas 快照。 这个可以有助于了解您的合约将消耗多少 Gas， 或者比较各种优化前后的 Gas 使用情况。

**要生成 Gas 快照，请运行 forge snapshot。**

默认情况下，这将**生成一个名为 .gas-snapshot** 的文件，其中包含你所有的测试及其各自的 Gas 使用情况。

```ignore
$ forge snapshot
$ cat .gas-snapshot

ERC20Test:testApprove() (gas: 31162)
ERC20Test:testBurn() (gas: 59875)
ERC20Test:testFailTransferFromInsufficientAllowance() (gas: 81034)
ERC20Test:testFailTransferFromInsufficientBalance() (gas: 81662)
ERC20Test:testFailTransferInsufficientBalance() (gas: 52882)
ERC20Test:testInfiniteApproveTransferFrom() (gas: 90167)
ERC20Test:testMetadata() (gas: 14606)
ERC20Test:testMint() (gas: 53830)
ERC20Test:testTransfer() (gas: 60473)
ERC20Test:testTransferFrom() (gas: 84152)
```

### 过滤

如果您想**指定一个不同的输出文件，请运行 forge snapshot --snap <FILE_NAME>。**

您还可以按 Gas 使用量对结果进行排序。 使用 **--asc 选项将结果按升序 排序和 --desc 将结果按降序排序。**

最后，您还可以为所有测试**指定最小/最大 Gas 阈值**。 要仅包含高于阈值的结果，您可以使用 --min <VALUE> 选项。 以同样的方式，只包括阈值以下的结果， 您可以使用 --max <VALUE> 选项。

您还可以将其与 forge test 的过滤器结合使用，例如 forge snapshot --match-path contracts/test/ERC721.t.sol 以生成与此测试合约相关的 Gas 快照。

### 比较Gas用量

如果您想**将当前快照文件与您的最新更改进行比较**，您可以使用 --diff 或 --check 选项。

**--diff 将与快照进行比较并显示快照的更改。**

它还可以**选择使用文件名（--diff <FILE_NAME>），默认是 .gas-snapshot。**

例如：

```ignore
$ forge snapshot --diff .gas-snapshot2

Running 10 tests for src/test/ERC20.t.sol:ERC20Test
[PASS] testApprove() (gas: 31162)
[PASS] testBurn() (gas: 59875)
[PASS] testFailTransferFromInsufficientAllowance() (gas: 81034)
[PASS] testFailTransferFromInsufficientBalance() (gas: 81662)
[PASS] testFailTransferInsufficientBalance() (gas: 52882)
[PASS] testInfiniteApproveTransferFrom() (gas: 90167)
[PASS] testMetadata() (gas: 14606)
[PASS] testMint() (gas: 53830)
[PASS] testTransfer() (gas: 60473)
[PASS] testTransferFrom() (gas: 84152)
Test result: ok. 10 passed; 0 failed; finished in 2.86ms
testBurn() (gas: 0 (0.000%))
testFailTransferFromInsufficientAllowance() (gas: 0 (0.000%))
testFailTransferFromInsufficientBalance() (gas: 0 (0.000%))
testFailTransferInsufficientBalance() (gas: 0 (0.000%))
testInfiniteApproveTransferFrom() (gas: 0 (0.000%))
testMetadata() (gas: 0 (0.000%))
testMint() (gas: 0 (0.000%))
testTransfer() (gas: 0 (0.000%))
testTransferFrom() (gas: 0 (0.000%))
testApprove() (gas: -8 (-0.000%))
Overall gas change: -8 (-0.000%)d
```

--check 将快照与现有快照文件进行比较并显示所有差异，如果有的话。 您可以通过提供不同的文件名来更改要比较的文件：--check <FILE_NAME>。

```ignore
$ forge snapshot --check .gas-snapshot2

Running 10 tests for src/test/ERC20.t.sol:ERC20Test
[PASS] testApprove() (gas: 31162)
[PASS] testBurn() (gas: 59875)
[PASS] testFailTransferFromInsufficientAllowance() (gas: 81034)
[PASS] testFailTransferFromInsufficientBalance() (gas: 81662)
[PASS] testFailTransferInsufficientBalance() (gas: 52882)
[PASS] testInfiniteApproveTransferFrom() (gas: 90167)
[PASS] testMetadata() (gas: 14606)
[PASS] testMint() (gas: 53830)
[PASS] testTransfer() (gas: 60473)
[PASS] testTransferFrom() (gas: 84152)
Test result: ok. 10 passed; 0 failed; finished in 2.47ms
Diff in "ERC20Test::testApprove()": consumed "(gas: 31162)" gas, expected "(gas: 31170)" 
```
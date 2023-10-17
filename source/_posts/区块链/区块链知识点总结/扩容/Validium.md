---
title: Validium

date: 2023-03-07	

categories: 扩容	

tags: [区块链,区块链知识点总结,扩容]
---	

# VALIDIUM

## 什么是 VALIDIUM？

Validium 是使用链下数据可用性和计算的扩展解决方案，旨在通过在以太坊主网外处理交易来提高吞吐量。 与零知识卷叠（ZK 卷叠）一样，Validium 发布[零知识证明](https://ethereum.org/zh/glossary/#zk-proof)以便在以太坊上验证链下交易。 这样可以防止无效的状态转换并增强 Validium 链的安全保障。

属于 Validium 用户的资金由以太坊上的智能合约控制。 恰如零知识卷叠一样，Validium 几乎可以提供即时提款；在主网上验证提款请求的有效性证明后，用户可以通过提供[默克尔证明](https://ethereum.org/zh/developers/tutorials/merkle-proofs-for-offline-data-integrity/)提取资金。 默克尔证明验证用户的提款交易是否包含在经过验证的交易批次中，从而允许链上合约处理提款。

但是，Validium 用户可以冻结他们的资金并限制提款。 如果 Validium 链上的数据可用性管理器不给用户提供链下状态数据，就会发生这种情况。 如果无法访问交易数据，用户将无法计算证明资金所有权和执行提款所需的默克尔证明。

这**是 Validium 和零知识卷叠之间的主要区别，它们在数据可用性范围内的位置不同**。 两种解决方案处理数据存储的方式不同，这会对安全性和去信任产生影响。
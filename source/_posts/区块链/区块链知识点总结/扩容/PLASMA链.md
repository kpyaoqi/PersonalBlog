---
title: PLASMA链

date: 2023-03-07	

categories: 扩容	

tags: [区块链,区块链知识点总结,扩容]
---	

# PLASMA链

## 什么是 PLASMA？

Plasma 是一个用于改善以太坊这类公共区块链的可扩展性的框架。 正如原 [Plasma 白皮书(opens in a new tab)↗](http://plasma.io/plasma.pdf)中所述，Plasma 链是在另一个区块链之上构建的，该区块链被称为“根链”。 每个“子链”都从根链延伸而来，通常由部署在母链上的智能合约进行管理。

Plasma 合约除了其他功能之外，还有一项功能是作为[链梁](https://ethereum.org/zh/developers/docs/bridges/)，让用户可以在以太坊主网和 plasma 链之间转移资产。 虽然这使它们**类似于[侧链](https://ethereum.org/zh/developers/docs/scaling/sidechains/)，但 plasma 链至少在某种程度上受益于以太坊主网的安全性。 这一点不同于单独负责其安全性的侧链。**

## PLASMA 如何工作？

Plasma 框架的基本组成部分包括：

### 链下计算

Plasma 假设以太坊主网不需要验证所有交易。 相反，我们可以在主网外处理交易，使节点不必验证每笔交易。

链下计算是必要的，因为 Plasma 链可以优化速度和成本。 例如，一个 Plasma 链可能，而且大多数情况下都使用单个“运营商”来管理交易的排序和执行。 由于只有一个实体验证交易，plasma 链上的处理速度比以太坊主网更快。

### 状态承诺

**虽然 Plasma 在链下执行交易，但它们是在以太坊主执行层上结算的，否则，Plasma 链无法从以太坊的安全保证中受益。** **但是在不知道 Plasma 链状态的情况下完成链下交易会破坏安全模型并让无效交易扩散。 这就是为什么运营商，即负责在 Plasma 链上生产区块的实体，需要定期在以太坊上发布“状态承诺”**。

[承诺方案(opens in a new tab)](https://en.wikipedia.org/wiki/Commitment_scheme)是一种加密技术，用于承诺价值或声明而不向另一方透露。 承诺是“有约束力的”，因为一旦你承诺了，就不能改变价值或声明。 **Plasma 中的状态承诺采用“Merkle 根”的形式**，运营商每隔一段时间将其发送到以太坊链上的 Plasma 合约。

Merkle 根是能够压缩大量信息的密码原语。 Merkle 根（在此情况下也称为“区块根”）可以代表区块中的所有交易。 Merkle 根还可以更容易地验证一小部分数据是否是较大数据集的一部分。 

Merkle 根对于向以太坊提供有关链下状态的信息非常重要。 你可以将 Merkle 根视为“保存点”：运营商表示，“这是 Plasma 链在 x 时间点的状态，这是 Merkle 根作为证明。” 运营商使用 Merkle 根对 Plasma 链的*当前状态*进行承诺，这就是为什么它被称为“状态承诺”。

### 入口和出口

为了让以太坊用户利用 Plasma，需要有一种机制在主网和 Plasma 链之间转移资金。 但是，我们不能随意将以太币发送到 Plasma 链上的地址 — 这些链是不兼容的，因此交易要么失败，要么导致资金损失。

Plasma 使用在以太坊上运行的主合约来处理用户的入口和出口。 该主合约还负责跟踪状态承诺（前面已解释）并通过欺诈证明惩罚不诚实行为（稍后将详细介绍）。

#### 进入 plasma 链

要进入 Plasma 链，Alice（用户）必须在 Plasma 合约中存入以太币或任何 ERC-20 代币。 监视合约存款的 Plasma 运营商重新创建与 Alice 的初始存款相等的金额，并将其释放到她在 Plasma 链上的地址。 Alice 需要证明在子链上收到资金，然后才能使用这些资金进行交易。

#### 退出 plasma 链

由于几个原因，退出 plasma 链比进入它更复杂。 最大的问题是，虽然以太坊有关于 Plasma 链状态的信息，但它无法验证信息是否真实。 恶意用户可能会做出不正确的断言（“我有 1000 个以太币”）并提供虚假证据来支持该声明而侥幸逃脱。

为防止恶意取款，引入了“挑战期”的概念。 在挑战期内（通常为一周），任何人都可以使用欺诈证明来挑战取款请求。 如果挑战成功，则取款请求被拒绝。

但是，通常情况下，用户是诚实的，并对他们拥有的资金做出正确的声明。 在这种情况下，Alice 将通过向 Plasma 合约提交交易，在根链（以太坊）上发起取款请求。

她还必须提供 Merkle 证明，验证在 Plasma 链上创建她的资金的交易是否包含在区块中。 这对于 Plasma 的迭代是必要的，例如[最小可行 Plasma(opens in a new tab)](https://www.learnplasma.org/en/learn/mvp.html) 使用[未花费的交易输出 (UTXO)(opens in a new tab)](https://en.wikipedia.org/wiki/Unspent_transaction_output) 模型。

其他的，如 [Plasma Cash(opens in a new tab)↗](https://www.learnplasma.org/en/learn/cash.html)，将资金表示为[非同质化代币](https://ethereum.org/zh/developers/docs/standards/tokens/erc-721/)，而不是未花费的交易输出。 在这种情况下，取款需要证明 Plasma 链上代币的所有权。 这是通过提交涉及代币的两个最新交易并提供 Merkle 证明来验证这些交易是否包含在区块中来完成的。

用户还必须在取款请求中添加保证金，作为诚实行为的保证。 如果挑战者证明 Alice 的取款请求无效，她的保证金将被罚没，其中一部分作为奖励交给挑战者。

如果在没有任何人提供欺诈证明的情况下经过挑战期，Alice 的取款请求被认为是有效的，允许她从以太坊上的 Plasma 合约中取回存款。
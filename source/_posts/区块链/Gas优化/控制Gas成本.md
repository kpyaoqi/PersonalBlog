---
title: 控制Gas成本

date: 2022-06-06	

categories: Gas优化	

tags: [区块链,Gas优化]
---	

# 优化合约

Solidity 编译器内置了一些优化器来改进合约的执行效率和生成更高效的字节码：

使用`--optimize`选项来启用优化器,使用`--optimize-runs`选项来指定优化器运行的次数

**solc --optimize (--optimize-runs 200) MyContract.sol** 

truffle、hardhat、Remix中都有配置这个选项

# 与永久性存储交互

> 以太坊上有三种数据存储位置： 内存（memory）、（永久性）存储（storage）以及调用数据calldata

查看[*以太坊黄皮书*](https://ethereum.github.io/yellowpaper/paper.pdf) 附录G 全面了解EVM操作码成本。

永久性存储操作码(`SSTORE`)非常昂贵。首次写插槽时，每个32个字节的当前成本是为20,000 Gas(在10 Gwei gas价格下为5美分，每ETH为250美元)，而后续每次修改则为5,000 Gas。尽管从理论上讲复杂度成本是`恒定的`，但它却是算术或内存运算成本的一千倍以上，而算术或内存运算的成本通常不到10 Gas。目前整个区块(截至2020年10月)的Gas限制为〜12,000,000 Gas实，开发人员应设计其智能合约以最大程度地减少所需的存储插槽数量。


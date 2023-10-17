---
title: UniswapV2Factory

date: 2023-01-14	

categories: UniswapV2	

tags: [区块链,区块链知识点总结,UniswapV2]
---	

# UniswapV2Factory

在构造函数中传入一个设置feeTo的权限者地址，主要用于创建两种token的交易对，并为其部署一个UniswapV2Pair合约用于管理这个交易对，UniswapV2Factory还包含一些手续费的一些设置.

## 合约当中具有的方法:

- function feeTo() external view returns (address)：返回收取手续费地址
- function feeToSetter() external view returns (address)：设置手续费收取地址的权限地址
- function getPair(address tokenA, address tokenB) external view returns (address pair)：获取两个token的交易对地址
- function allPairs(uint) external view returns (address pair)：返回指定位置的交易对地址
- function allPairsLength() external view returns (uint)：返回所有交易对的长度
- function createPair(address tokenA, address tokenB) external returns (address pair)：创建两个token的交易对地址
- function setFeeTo(address) external：更改收取手续费地址
- function setFeeToSetter(address) external：更改设置手续费收取地址的权限地址

在 Uniswap 协议中，`feeTo` 是一个变量，用于指定手续费收取地址。当用户在 Uniswap 上进行交易时，一定比例的交易手续费会被收取，并根据协议的设定进行分配。这个手续费分配的过程包括将一部分手续费发送给流动性提供者，同时还有一部分手续费发送到 `feeTo` 地址。

#### createPair:(address tokenA, address tokenB) returns (address pair)

```solidity
bytes32 salt = keccak256(abi.encodePacked(token0, token1));
bytes memory bytecode = type(UniswapV2Pair).creationCode;
assembly {
	//add(bytecode, 32)：opcode操作码的add方法,将bytecode偏移后32位字节处,因为前32位字节存的是bytecode长度
	//mload(bytecode)：opcode操作码的方法,获得bytecode长度
	pair := create2(0, add(bytecode, 32), mload(bytecode), salt)
}
```

pair地址是通过**内联汇编assembly做create2**方法创建的，其中**salt**盐值是通过两个两个代币的地址计算

> 内联汇编：在 Solidity 源程序中嵌入汇编代码，对 EVM 有更细粒度的控制


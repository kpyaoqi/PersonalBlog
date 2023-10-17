---
title: 9-Forge标准库参考

date: 2023-02-26	

categories: Foundry	

tags: [区块链,区块链知识点总结,合约框架,Foundry]
---	

# Forge 标准库

包含的内容：

- 标准库

  - Std Logs: 在 DSTest 库的日志事件的基础上进行扩展。
  - Std Assertions: 在 DSTest 库的断言功能的基础上进行扩展。
  - Std Cheats: 围绕 Forge 作弊代码的包装，以提高安全性和开发人员体验。
  - Std Errors: 围绕常见的内部 Solidity 错误和 reverts 的包装。
  - Std Storage: 用于存储操作的工具类。
  - Std Math: 实用的数学函数。
  - Script Utils: 可以在测试和脚本中访问的工具类。
  - Console Logging: 控制台日志功能。

- 一个作弊代码实例 vm，你可以通过它调用 Forge 作弊代码

  ```solidity
  vm.startPrank(alice);
  ```

- 所有 Hardhat console 功能的记录

  ```solidity
  console.log(alice.balance); // or `console2`
  ```

- 所有用于断言和日志的 Dappsys 测试函数

  ```solidity
  assertEq(dai.balanceOf(alice), 10000e18);
  ```

- 工具类函数也包括在 Script.sol 中

  ```solidity
  // 对于给定的部署者地址和 nonce，计算合约将被部署到的地址
  address futureContract = computeCreateAddress(alice, 1);
  ```

## Std Logs

Std Logs 在 DSTest 库的日志事件基础上进行扩展。

```solidity
event log_array(uint256[] val);
event log_array(int256[] val);
event log_array(address[] val);
event log_named_array(string key, uint256[] val);
event log_named_array(string key, int256[] val);
event log_named_array(string key, address[] val);
```

## Std Assertions

1. **fail**(string memory err)：如果命中某个分支或执行点，则测试失败，并发出消息。
2. **assertFalse**(bool data)、**assertFalse**(bool data, string memory err)：断言条件为假。
3. **assertEq**：断言 `a` 等于 `b`，对 `bool`、`bytes`、`int256` 和 `uint256` 数组起作用。
4. **assertApproxEqAbs**：uint256，断言 `a` 近似等于 `b`，delta 为百分比
5. **assertApproxEqRel**：uint256，断言 `a` 近似等于 `b`，delta 为百分比，其中 `1e18` 为100%。

## Std Cheats

1. **skip**(uint256 time)：将 `block.timestamp` 向前跳过指定的秒数

2. **rewind**(uint256 time)：将 `block.timestamp` 重新调整为指定的秒数

3. **hoax**(address who,(uint256 give),(address origin)):从一个有一些 ether 的地址设置一个 prank，如果没有指定余额，它将被设置为 2^128 wei

4. **startHoax**(address who,(uint256 give),(address origin)):从一个有一些 ether 的地址开始一个永久的prank，如果没有指定余额，它将被设置为 2^128 wei

5. **deal**((address token,)address to, uint256 give(, bool adjust))：围绕 deal 作弊码的包装，也适用于大多数 ERC-20 代币,如果使用 `deal` 的替代签名，在设置余额后调整代币的总供应量。

   ```solidity
   deal(address(dai), alice, 10000e18);
   assertEq(dai.balanceOf(alice), 10000e18);
   ```

6. **deployCode**(string memory what(,bytes memory args)(,uint256 val)) returns (address):通过从 artifacts 目录中获取合约字节码来部署合约,calldata 参数可以是 `ContractFile.sol`（如果文件名和合约名称相同），`ContractFile.sol:ContractName`，或者是相对于项目根目录的 artifact 路径。值也可以通过使用 `val` 参数来传递。如果你需要在构造函数上发送 ETH，这是必要的。

   ```solidity
   address deployment = deployCode("MyContract.sol", abi.encode(arg1, arg2));
   ```

7. **bound**(uint256 x, uint256 min, uint256 max) returns (uint256 result)：一个数学函数，用于将模糊测试的输入包裹在一定范围内

   ```solidity
   input = bound(input, 99, 101);
   返回99输入0。
   返回100输入1。
   返回101输入2。
   返回99输入3。
   ```

8. **changePrank**(address who)：

9. **makeAddr**(string memory name) returns(address addr)：创建一个由所提供的 `name` 衍生的地址

10. **makeAddrAndKey**(string memory name) returns(address addr, uint256 privateKey)：创建一个由提供的 `name` 衍生的地址和私钥。

11. **noGasMetering**：一个函数修改器，在函数的生命周期内关闭 Gas 计量

## Std Errors

1. **assertionError**：当断言失败时，Solidity 内部的错误
2. **arithmeticError**：当算术运算失败时，Solidity 内部的错误，例如下溢和溢出
3. **divisionError**：当除法失败时的 Solidity 内部错误，例如除以零
4. **enumConversionError**：当试图将一个数字转换为一个枚举的变体时，如果该数字大于该枚举的变体数量（从0开始计算），则 Solidity 内部会出现错误
5. **encodeStorageError**：当试图访问存储中被破坏的数据时的 Solidity 内部错误，除非使用了程序集，否则数据不会损坏。
6. **popError**：当试图从一个空数组上弹出一个元素时，Solidity 内部的错误
7. **indexOOBError**：当试图访问一个超出边界的数组元素时，Solidity 内部的错误
8. **memOverflowError**：当试图分配一个超过 2^64-1 项的动态内存数组，Solidity 内部出现的错误
9. **zeroVarError**：当试图通过一个尚未初始化的函数指针调用一个函数时，Solidity 的内部错误

> 为了捕捉的错误，我们在预计会导致下溢的函数调用之前插入`vm.expectRevert(stdError.arithmeticError)`。

## Std Storage

Std Storage 是一个库，使操纵存储变得容易，要使用 Std Storage，在你的测试合约中添加以下一行：

```solidity
using stdStorage for StdStorage;
```

### 功能

查询功能：

- target: 设置目标合约的地址
- sig: 将函数的 4 字节选择器设置为静态调用
- with_key: 传递一个参数给函数，可以多次使用以传递多个参数。顺序很重要。
- depth: 设置值在 tuple 中的位置（例如，在 struct 中）。

终端功能：

- find: 在给定的 target、sig、with_key(s) 和 depth 中找到一个任意的储存槽
- checked_write: 设置要写到存储槽的数据
- read_<type>:从存储槽中读取 `bytes32`，`bool`，`address`，`uint256`，或`int256` 的值。

```solidity
// MetaRPG.sol
struct Character {
    string name;
    uint256 level;
}
// playerToCharacter跟踪玩家的角色信息。
mapping (address => Character) public playerToCharacter;
```

```solidity
// MetaRPG.t.sol
// 比方说，我们想把我们的角色的等级设置为 120。
stdstore
    .target(address(metaRpg))
    .sig("playerToCharacter(address)")
    .with_key(address(this))
    .depth(1)
    .checked_write(120);
```

## Std Math

1. **abs**(int256 a) returns (uint256):返回一个数字的绝对值。
2. **delta**(uint256 a, uint256 b) returns (uint256):返回两个数字的绝对值之差
3. **percentDelta**(uint256 a, uint256 b) returns (uint256)：返回两个数字之间的差值，以百分比表示，其中 `1e18` 是 100%

## **Script Utils**

1. **computeCreateAddress**(address deployer, uint256 nonce) returns (address)：对于给定的部署者地址和 nonce，计算合约将被部署到的地址。对预先计算合约将被部署的地址很有用
2. **deriveRememberKey**(string memory mnemonic, uint32 index) returns (address who, uint256 privateKey)：从助记符中导出私钥，同时将其存储在 forge 的本地钱包中。返回地址和私钥。
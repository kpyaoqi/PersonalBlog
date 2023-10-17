---
title: 16-Assembly内联汇编

date: 2022-07-22	

categories: Solidity0.8.17	

tags: [区块链,Solidity0.8.17]
---	

使用内联汇编，可以在 Solidity 源程序中嵌入汇编代码，对 EVM 有更细粒度的控制。内联汇编主要用在编写库函数时很有用，一般用于写工具函数，比如椭圆签名解析等。在项目中用汇编编主要是 opensea 的 [seaport](https://github.com/ProjectOpenSea/seaport) 合约.

在合约的内部使用汇编，是在合约内部包含 `assembly` 关键字进行编写的，在 Solidity `inline assembly`(内联汇编) 中的语言被称为 Yul。

[Yul](https://docs.soliditylang.org/zh/latest/yul.html) 除了在 Solidity 之中作为 inline assembly 的一部分，也能当作独立的直译语言能够被编译成 bytecode 给不同的后端。

注意：内联汇编是一种在底层访问以太坊虚拟机的语言，由于编译器无法对汇编语句进行检查，所以 Solidity 提供的很多重要安全特性都没办法作用于汇编。写汇编代码相对比较困难，很多时候只有在处理一些相对复杂的问题时才需要使用它，并且开发者需要明确知道自己要做什么。

# 基本格式

通过 `assembly {}` 包裹代码。并且内部每一行语句不需要使用`;`显示的标注结束。Assembly 也支持注释，可以使用 `//` 和 `/* */` 来进行注释。

⚠️ 注意： Inline Assembly 中，代码块之间是不能彼此沟通的，里面声明的变量都是本地变量。

### 例子: 不同代码块无法互相访问

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public view returns (uint256) {
        assembly {
            let x := 2
        }
        assembly {
            let y := x // DeclarationError: Identifier "x" not found.
        }
    }
}
```

let 指令执行如下任务：

- 创建一个新的堆栈槽位
- 为变量保留该槽位
- 当到达代码块结束时自动销毁该槽位

因此，使用 let 指令在汇编代码块中定义的变量，在代码块外部是无法访问的。但是内部代码块可以访问外部代码块的内容。

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure {
        assembly {
            let x := 3

            {
                let y := x // success
            } // 到此处会销毁y
        }
    }
}
```

### 例子: 简单的加法

下面是一个计算 `_x + _y` 的两种写法对比，汇编的语法节省了 `1.76%` 的 gas。 assembly 核心是更细粒度的控制，省 gas 只是它的外在表现。

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    // 输入 1,2 ; 输出 22307 gas
    function addSolidity(uint256 _x, uint256 _y) public pure returns (uint256) {
        return (_x + _y);
    }

    // 输入 1,2 ; 输出 21915 gas
    function addAssembly(uint256 _x, uint256 _y) public pure returns (uint256) {
        assembly {
            // let result 是声明一个变量 result
            // add(_x, _y) 是计算 x + y 的结果
            // := 是将 x + y 的结果赋值给变量 result
            let result := add(_x, _y)

            // mstore(0x0, result) 在内存 `0x0` 的位置储存 `result`
            mstore(0x0, result)

            // 从内存索引 0x0 位置返回32字节
            return(0x0, 32)
        }
    }
}
```

# 语言基础

Yul 提供了高级结构，如 `for` 循环、`if` 语句 `switch` 和函数调用等等，下面按照分类进行介绍。

在 Inline Assembly 中，以下几个点很重要：

- 赋值: 使用的是`:=`，而不是`=`。
- 声明变量: 使用 `let` 声明；（不是正常带有指定类型的强类型方式）

### 声明与赋值

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (uint256, uint256) {
        assembly {
            let x := 2  // 声明 x，赋值为2
            let y       // 声明 y，初始化为 0
            y := 5      // 赋值 y 为5

            mstore(0x0, x) // 内存中储存 x
            mstore(add(0x0, 32), y) // 内存中移动32位，再储存y

            // 返回内存中 0 - 64 的数据
            return(0x0, 64)
        }
    }
}
```

结果就是 `2,5`；

在 Solidity 汇编中字面量的写法与 Solidity 一致。但是 字符串字面量 最多可以包含 32 个字符。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure {
        assembly {
            let a := 0x123 // 16进制
            let b := 42 // 10进制
            let c := "hello world" // 字符串
            let d := "very long string more than 32 bytes" // 长度 35 的 字符串，错误！
        }
    }
}
```

### 汇编只能读取局部变量

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    uint256 a = 2;
    function demoAssembly() public pure {
        uint256 b = 5;
        assembly {
            // 可以读取 x 和 y
            let x := add(2, 3)
            let y := 10
            let z := add(x, y)
        }
        assembly {
            // 可以读取 x 和 b
            let x := add(2, 3)
            let y := mul(x, b)
        }
        assembly {
            let x := add(2, 3)

            // ❌ TypeError: Only local variables are supported.
            // To access storage variables, use the ".slot" and ".offset" suffixes.
            let y := mul(x, a)
        }
    }
}
```

# 条件判断

- if
- switch

### if

特点如下

- 只有 if ，没有 else
- if 语句强制要求代码块使用大括号，`{}`不允许省略

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure {
        uint256 x;
        assembly {
            // success
            if iszero(x) {
                x := sub(1, x)
            }

            // fail: 没有使用 {} 包裹代码
            // if iszero(x) revert(0, 0)
        }
    }
}
```

如果需要在 Solidity 内联汇编中检查多种条件，可以考虑使用 switch 语句。

### switch

switch 语句支持 一个默认分支 default，当表达式的值不匹配任何其他分支条件时，将 执行默认分支的代码。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 x) public pure returns (uint256 result) {
        assembly {
            switch x
            case 0 {
                result := 0
            }
            case 1 {
                result := 1
            }
            default {
                result := mul(x, x)
            }
        }
    }
}
```

# for 循环

for 循环也包含 3 个元素

- 初始化：比如`let i := 0`
- 执行条件：比如`lt(i, n)` ，必须是函数风格表达式
- 迭代后续步骤：比如`add(i, 1)`

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (uint256 result) {
        uint256 leng = 10;

        assembly {
            for
            { let i := 0 }
            lt(i, leng)
            { i := add(i, 1) }
            {
                result := add(result, i)
            }

            // 下面可以省略
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

for 循环的**初始化部分**和**迭代后续步骤**可以留空 , 改写为下面的格式

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (uint256 result) {
        uint256 leng = 10;
        assembly {
            let i := 0 // 初始条件写在这
            for {} lt(i, leng) {} {
                // 核心部分
                result := add(result, i)

                // 迭代后续步骤写在这
                i := add(i, 1)
            }

            // 可以省略
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

备注: `continue` or `break` 语句只能在 `for` 循环体内使用

# 函数的定义和使用

函数的运行机制如下：

- 从堆栈提取参数
- 将结果压入堆栈
- 和 Solidity 函数不同，不需要指定汇编函数的可见性
  - 例如 public 或 private， 因为汇编函数仅在定义所在的汇编代码块内有效。

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (uint256 free_memory_pointer) {
        assembly {
            // 函数定义
            function allocate(length) -> pos {
                pos := mload(0x40)
                mstore(0x40, add(pos, length))
            }

            // 函数使用
            free_memory_pointer := allocate(64)
        }
    }
}
```

# EVM 内置函数/内置操作码

- 算数操作
  - `add`: 加法
  - `mul`:
- 比较操作
  - `lt`
  - `gt`
- 位操作
  - `not`:
  - `and`:
- 密码学操作，目前仅包含 keccak256
- 环境操作，主要指与区块链相关的全局信息，例如 blockhash 或 coinbase 收款账号
- 存储、内存和栈操作
- 交易与合约调用操作
- 停机操作
- 日志操作

下面是详细的列表说明，标记为`-`的操作不返回结果，其他所有操作码只返回一个值。

标有 F、H、B、C 、I 和 L 分别自出现的时间，对应的如下

- `F`: Frontier
- `H`: Homestead
- `B`: Byzantium
- `C`: Constantinople
- `I`: Istanbul
- `L`: London

常见的常量值是 `0x20` / `0x40` , 代表十进制的 32 和 64。

### 数学计算

| 操作符号        | 返回值 | 版本 | 解释说明                                                |
| --------------- | ------ | ---- | ------------------------------------------------------- |
| add(x, y)       |        | F    | `x + y`                                                 |
| sub(x, y)       |        | F    | `x - y`                                                 |
| mul(x, y)       |        | F    | `x * y`                                                 |
| div(x, y)       |        | F    | `x / y` (如果 y 为 0，则结果为 0)                       |
| mod(x, y)       |        | F    | `x % y` (如果 y 为 0，则结果为 0)                       |
| exp(x, y)       |        | F    | `x` 的 `y` 次方                                         |
| addmod(x, y, m) |        | F    | `(x + y) % m` 任意精度算术，如果 m == 0 则为 0          |
| mulmod(x, y, m) |        | F    | `(x * y) % m` 任意精度算术，如果 m == 0 则为 0          |
| sdiv(x, y)      |        | F    | `x / y`, 以二进制补码作为符号 (如果 y 为 0，则结果为 0) |
| smod(x, y)      |        | F    | `x % y`, 以二进制补码作为符号 (如果 y 为 0，则结果为 0) |

#### add: 加法

```solidity
function demoAssembly(uint256 _x, uint256 _y)
    public
    pure
    returns (uint256)
{
    assembly {
        let result := add(_x, _y)
        mstore(0x0, result)
        return(0x0, 32)
    }
}
```

上面合约函数，传入参数:`1,2`，返回`3`。

这里需要返回`uint256`类型，assembly 内部返回是，从什么位置开发，返回多少个数据。需要返回两个数据。比如我把 uint256 改为 uint8，代码如下

```solidity
function demoAssembly(uint8 _x, uint8 _y)
    public
    pure
    returns (uint8)
{
    assembly {
        let result := add(_x, _y)
        mstore(0x0, result)
        return(0x0, 2)
    }
}
```

相同的参数会报错: `error:Failed to decode output: Error: data out-of-bounds (length=2, offset=32, code=BUFFER_OVERRUN, version=abi/5.5.0)`

#### sub: 减法

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x, uint256 _y)
        public
        pure
        returns (uint256)
    {
        assembly {
            let result := sub(_x, _y)
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

传入参数:`2,1`，返回`1`。

注意：这时候如果传参 `1,2`，会溢出返回，得到的结果不会报错，反而是:`115792089237316195423570985008687907853269984665640564039457584007913129639935`，因为 assembly 绕过了 solidity 的安全检查。当我们使用 assembly 编码时候，安全问题需要自己控制，不要错误的认为 solidity 的默认机制会保护代码。

#### mul: 乘法

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x, uint256 _y)
        public
        pure
        returns (uint256)
    {
        assembly {
            let result := mul(_x, _y)
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

传入参数:`2,3`，返回`6`。

#### div: 除法

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x, uint256 _y)
        public
        pure
        returns (uint256)
    {
        assembly {
            let result := div(_x, _y)
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

- 传入参数:`3,2`，返回`1`。
- 传入参数:`3,1`，返回`3`。
- 传入参数:`3,0`，返回`0`。

#### mod: 求模

```so
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x, uint256 _y)
        public
        pure
        returns (uint256)
    {
        assembly {
            let result := mod(_x, _y)
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

- 传入参数:`3,2`，返回`1`
- 传入参数:`3,1`，返回`0`
- 传入参数:`3,0`,返回`0`
- 传入参数:`3,30`,返回`3`

#### exp: 次方

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x, uint256 _y)
        public
        pure
        returns (uint256)
    {
        assembly {
            let result := exp(_x, _y)
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

- 传入参数:`10,2`，返回`100`。
- 传入参数:`10,3`，返回`1000`。

#### addmod: 先求和再求模

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(
        uint256 _x,
        uint256 _y,
        uint256 _m
    ) public pure returns (uint256) {
        assembly {
            let result := addmod(_x, _y, _m)
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

- 传入参数:`2,3,3`，返回`2`。

#### mulmod: 先相乘再求模

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(
        uint256 _x,
        uint256 _y,
        uint256 _m
    ) public pure returns (uint256) {
        assembly {
            let result := mulmod(_x, _y, _m)
            mstore(0x0, result)
            return(0x0, 32)
        }
    }
}
```

- 传入参数:`2,3,3`，返回`0`。

#### 二进制补码

下面两个方法，用法基本和 div / mod 差不多

```
let result := sdiv(_x, _y)
let result := smod(_x, _y)
```

### 比较关系

| 操作符号  | 返回值 | 版本 | 解释说明                                          |
| --------- | ------ | ---- | ------------------------------------------------- |
| gt(x, y)  |        | F    | 如果 `x > y` 等于 1, 否则 0                       |
| lt(x, y)  |        | F    | 如果 `x < y` 等于 1, 否则 0                       |
| eq(x, y)  |        | F    | 如果 `x == y` 等于 1, 否则 0                      |
| iszero(x) |        | F    | 如果 `x == 0` 等于 1, 否则 0                      |
| slt(x, y) |        | F    | 如果 `x < y` 等于 1, 否则 0, 以二进制补码作为符号 |
| sgt(x, y) |        | F    | 如果 `x > y` 等于 1, 否则 0, 以二进制补码作为符号 |

#### gt: 大于

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x, uint256 _y)
        public
        pure
        returns (bool result)
    {
        assembly {
            result := gt(_x, _y)
        }
    }
}
```

- 传入参数:`1,2`，返回 `false`
- 传入参数:`10,3`，返回 `true`

#### lt: 小于

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x, uint256 _y)
        public
        pure
        returns (bool result)
    {
        assembly {
            result := lt(_x, _y)
        }
    }
}
```

- 传入参数:`1,2`，返回 `true`
- 传入参数:`10,3`，返回 `false`

#### eq: 等于

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x, uint256 _y)
        public
        pure
        returns (bool result)
    {
        assembly {
            result := eq(_x, _y)
        }
    }
}
```

- 传入参数:`1,2`，返回 `false`
- 传入参数:`2,2`，返回 `true`

#### iszero: 等于零

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(uint256 _x)
        public
        pure
        returns (bool result)
    {
        assembly {
            result := iszero(_x)
        }
    }
}
```

注意该参数只接收一个参数

- 传入参数:`1`，返回 `false`
- 传入参数:`0`，返回 `true`

### 按位 & 移位

| 操作符号  | 返回值 | 版本 | 解释说明                            |
| --------- | ------ | ---- | ----------------------------------- |
| not(x)    |        | F    | 对 x 按位取反,类似`~x`;`x` 的按位非 |
| and(x, y) |        | F    | x 和 y 的按位与                     |
| or(x, y)  |        | F    | x 和 y 的按位或                     |
| xor(x, y) |        | F    | x 和 y 的按位异或                   |
| shl(x, y) |        | C    | y 逻辑左移 x 位                     |
| shr(x, y) |        | C    | y 逻辑右移 x 位                     |
| sar(x, y) |        | C    | 将 y 算术右移 x 位                  |

#### not: 按位非

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly(int256 _x) public pure returns (int256 result) {
        assembly {
            result := not(_x)
        }
    }
}
```

- 传入参数:`0`，返回 `-1`
- 传入参数:`1`，返回 `-2`
- 传入参数:`-1`，返回 `0`
- 传入参数:`-11`，返回 `10`

#### and: 按位与

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (int256 result) {
        int256 _x = 2;
        int256 _y = 3;
        assembly {
            result := and(_x, _y)
        }
    }
}
```

结果是 2

#### or: 按位或

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (int256 result) {
        int256 _x = 2;
        int256 _y = 3;
        assembly {
            result := or(_x, _y)
        }
    }
}
```

结果是 3

#### xor: 按位异或

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (int256 result) {
        int256 _x = 2;
        int256 _y = 3;
        assembly {
            result := xor(_x, _y)
        }
    }
}
```

结果 1

#### shl: 逻辑左移

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (uint256 result) {
        uint256 A = 2;
        assembly {
            result := shl(A, 1) // 4
        }
    }
}
```

#### shr: 逻辑右移

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (uint256 result) {
        uint256 A = 2;
        assembly {
            result := shr(A, 1) // 0
        }
    }
}
```

#### sar: 算术右移

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public pure returns (uint256 result) {
        uint256 A = 2;
        assembly {
            result := sar(A, 1) // 0
        }
    }
}
```

### EVM 区块交易相关

| 操作符号       | 返回值 | 版本 | 解释说明                                                     |
| -------------- | ------ | ---- | ------------------------------------------------------------ |
| address()      |        | F    | 当前合约地址 / execution context                             |
| balance(a)     |        | F    | 地址 a 的 wei 余额                                           |
| selfbalance()  |        | I    | 相当于 `balance(address())`，但更便宜                        |
| extcodehash(a) |        | C    | 地址 a 的代码哈希                                            |
| **msg 相关**   |        |      |                                                              |
| caller()       |        | F    | call sender ( 类似`msg.sender`？) (excluding `delegatecall`) |
| callvalue()    |        | F    | wei sent together with the current call（类似`msg.value`？） |
| **block 相关** |        |      |                                                              |
| chainid()      |        | I    | 当前网络的链 ID (EIP-1344)                                   |
| basefee()      |        | L    | 当前区块的基本费用 (EIP-3198 and EIP-1559)                   |
| timestamp()    |        | F    | 当前块的时间戳，自纪元以来的秒数                             |
| coinbase()     |        | F    | 当前采矿受益人                                               |
| number()       |        | F    | 当前区块号                                                   |
| difficulty()   |        | F    | 当前区块的难度                                               |
| gaslimit()     |        | F    | 当前区块的区块 gas limit                                     |
| **tx 相关**    |        |      |                                                              |
| origin()       |        | F    | 交易发送方                                                   |
| gasprice()     |        | F    | 交易的 gas 价格                                              |
| **其它**       |        |      |                                                              |
| gas()          |        | F    | 剩余 gas                                                     |
| blockhash(b)   |        | F    | 指定 block 的 hash - 仅适用于最后 256 个块，不包括当前块     |

#### address()

相当于 `address(this)`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public view returns (address ads1, address ads2) {
        assembly {
            ads1 := address()
        }

        ads2 = address(this);
    }
}
```

返回

- `0:address: ads1 0x3c725134d74D5c45B4E4ABd2e5e2a109b5541288`
- `1:address: ads2 0x3c725134d74D5c45B4E4ABd2e5e2a109b5541288`

#### balance(a)

相当于 `address.balance`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        address sender = msg.sender;
        assembly {
            result1 := balance(sender)
        }

        result2 = address(sender).balance;
    }
}
```

返回

- `0:uint256: result1 99999999999992173039`
- `1:uint256: result2 99999999999992173039`

#### selfbalance()

相当于 `balance(address())`，但更便宜

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    constructor() payable {}

    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := balance(address())
        }

        result2 = address(this).balance;
    }
}
```

#### extcodehash(a)

相当于 `address.codehash`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (bytes32 result1, bytes32 result2)
    {
        address sender = msg.sender;
        assembly {
            result1 := extcodehash(address())
        }
        result2 = address(this).codehash;
    }
}
```

- `0:bytes32: result1 0xcbefd958c5e7814e7e635b599c5859eb893c410292a7f9f82088c3e84ee3c0e9`
- `1:bytes32: result2 0xcbefd958c5e7814e7e635b599c5859eb893c410292a7f9f82088c3e84ee3c0e9`

#### caller()

相当于 `msg.sender`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly() public view returns (address ads1, address ads2) {
        assembly {
            ads1 := caller()
        }

        ads2 = msg.sender;
    }
}
```

#### callvalue()

相当于 `msg.value`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        payable
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := callvalue()
        }
        result2 = msg.value;
    }
}
```

#### chainid()

相当于 `block.chainid`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := chainid() // 1
        }
        result2 = block.chainid; // 1
    }
}
```

#### basefee()

相当于 `block.basefee`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := basefee()
        }
        result2 = block.basefee;
    }
}
```

#### timestamp()

相当于 `block.timestamp`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := timestamp()
        }
        result2 = block.timestamp;
    }
}
```

#### coinbase()

相当于 `block.coinbase`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (address result1, address result2)
    {
        assembly {
            result1 := coinbase()
        }
        result2 = block.coinbase;
    }
}
```

#### number()

相当于 `block.number`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := number()
        }
        result2 = block.number;
    }
}
```

#### difficulty()

相当于 `block.difficulty`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := difficulty()
        }
        result2 = block.difficulty;
    }
}
```

#### gaslimit()

相当于 `block.gaslimit`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := gaslimit()
        }
        result2 = block.gaslimit;
    }
}
```

#### origin()

相当于 `tx.origin`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (address result1, address result2)
    {
        assembly {
            result1 := origin()
        }
        result2 = tx.origin;
    }
}
```

#### gasprice()

相当于 `tx.gasprice`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := gasprice()
        }
        result2 = tx.gasprice;
    }
}
```

#### gas()

相当于 `gasleft()`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (uint256 result1, uint256 result2)
    {
        assembly {
            result1 := gas() // 2978815
        }

        // 与 assembly 之间的顺序不改变最后的值
        // 所以 assembly 优先执行？
        result2 = gasleft(); // 2978808
    }
}
```

#### blockhash(b)

相当于 `blockhash(number)`

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function demoAssembly()
        public
        view
        returns (bytes32 result1, bytes32 result2)
    {
        assembly {
            result1 := blockhash(1)
        }
        result2 = blockhash(1);
    }
}
```

### 常见方法

| 操作符号            | 返回值 | 版本 | 解释说明                              |
| ------------------- | ------ | ---- | ------------------------------------- |
| sload(p)            |        | F    | `storage[p]`                          |
| mload(p)            |        | F    | `mem[p…(p+32))`                       |
| sstore(p, v)        | -      | F    | `storage[p] := v`                     |
| mstore(p, v)        | -      | F    | `mem[p…(p+32)) := v`                  |
| mstore8(p, v)       | -      | F    | `mem[p] := v` & 0xff (只修改单个字节) |
| keccak256(p, n)     |        | F    | `keccak(mem[p…(p+n)))`                |
| create(v, p, n)     |        | F    | create 创建合约                       |
| create2(v, p, n, s) |        | C    | create2 创建合约                      |

小例子:

- `mload(p)`: 分配数据
- `mstore(offset, value)`: 在 `offset` 的位置储存 `value`

#### sload(p)

sload 是 storage load，`sload(key)` 是从 storage 的哪个 slot 来 load，详细原理可以在后面介绍的 **状态变量在存储中的布局** 了解更多。

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    uint256 public a = 123;
    uint256 public b = 456;

    function demoAssembly() public view returns (uint256) {
        assembly {
            // v 是长度是 32 bytes
            // 从 slot #0 读数据 => 读到的是 123
            // 从 slot #1 读数据 => 读到的是 456
            let v := sload(0)

            // 在内存位置 0x80 处储存变量 v 后面的数据
            mstore(0x80, v)

            // 返回值:从 0x80 位置，返回 32个字节
            return(0x80, 32)
        }
    }
}
```

上面例子中，`slot #0` 是 123，`slot #1` 是 456。

注意: `slot #0` 可能是多个状态变量公用的。比如把状态变量改为如下类型,读 `slot #0` ，该位置储存了`a+b`;

```
uint128 public a = 1;
uint128 public b = 2;
uint256 public c = 456;
```

#### mload(p)

mload 是 memory load，`mload(key)` 是从 memory 的哪个 slot 来 load，类似 `sload`。

- 问题：内存数据 mload 时为什么从第 32 位开始?
  - 答案：前 32 个字节存储的是数据的长度;
  - 参考: https://www.cnblogs.com/wanghui-garcia/p/9592807.html

#### sstore(p, v)

#### mstore(p, v)

下面 name1 和 name2 都返回 “Anbang” 的字符串

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    // gas 23471
    bytes6 public name1 = "Anbang"; // 0x416e62616e67

    // gas 21229
    function name2() public pure returns (string memory) {
        assembly {
            // 在 0x20 处 储存值 0x20
            mstore(0x20, 0x20)

            // name1 length = 0x06
            // 参数1: 0x40 + length = 0x40 + 0x06 => 0x46
            // 参数2: length + name1  = 0x46 + 0x416e62616e67 => 0x06416e62616e67
            mstore(0x46, 0x06416e62616e67)

            // 返回 memory 从 0x20处之后的 0x60 长度的数据
            return(0x20, 0x60)
        }
    }
}
```

上面是从`0x20`处开始写数据，这个位置不是强制的，使用`0x00`也可以的。

#### mstore8(p, v)

#### keccak256(p, n)

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    // 0x01
    function solidityKeccak(bytes memory _input) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_input));
    }

    // 0x01
    function assemblyKeccak(bytes memory _input)
        public
        pure
        returns (bytes32 x)
    {
        assembly {
            x := keccak256(add(_input, 0x20), mload(_input))
        }
    }
}
```

#### create(v, p, n)

```
 assembly {
    // create(v,p,n);
    // v 是 发送的ETH值
    // p 是 内存中机器码开始的位置
    // n 是 内存中机器码的大小
    // msg.value 不能使用，需要用 callvalue()
    adds := create(callvalue(), add(_code, 0x20), mload(_code))
}
```

#### create2(v, p, n, s)

下面是 [UniswapV2Factory](https://github.com/Uniswap/v2-core/blob/master/contracts/UniswapV2Factory.sol) 中创建 pair 核心逻辑

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Z {
    bytes6 public name1 = "Anbang";
}

contract Demo {
    function addr() public returns (address pair) {
        bytes memory bytecode = type(Z).creationCode;
        // bytes32 salt = keccak256(abi.encodePacked(address(0), address(1)));
        uint256 salt = block.number;
        assembly {
            pair := create2(0, add(bytecode, 32), mload(bytecode), salt)
        }
    }
}
```

### 操作数据/大小

| 操作符号                | 返回值 | 版本 | 解释说明                                                     |
| ----------------------- | ------ | ---- | ------------------------------------------------------------ |
| msize()                 |        | F    | 内存大小，即最大访问内存索引                                 |
| pc()                    |        | F    | 当前在代码中的位置                                           |
| codesize()              |        | F    | 当前合约的代码大小 / execution context                       |
| codecopy(t, f, s)       | -      | F    | 从位置 f 的代码复制 s 个字节到位置 t 的内存                  |
| extcodesize(a)          |        | F    | 获取地址 a 的代码大小                                        |
| extcodecopy(a, t, f, s) | -      | F    | 像 codecopy(t, f, s) 但在地址 a 处获取代码                   |
| signextend(i, x)        |        | F    | sign extend from `(i*8+7)`th bit counting from least significant |
| byte(n, x)              |        | F    | x 的第 n 个字节，这个索引是从 0 开始的                       |
| pop(x)                  | -      | F    | 丢弃值 x                                                     |

#### msize()

内存大小，即最大访问内存索引

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    function test() public pure returns (int256) {
        int8 v0 = 1;
        assembly {
            v0 := msize()
        }
        return int256(v0);
    }
}
```

#### pc()

#### codesize()

#### codecopy(t, f, s)

#### extcodesize(a)

#### extcodecopy(a, t, f, s)

#### signextend(i, x)

#### byte(n, x)

#### pop(x)

### call 相关

| 操作符号                                     | 返回值 | 版本 | 解释说明                                                     |
| -------------------------------------------- | ------ | ---- | ------------------------------------------------------------ |
| calldataload(p)                              |        | F    | 从位置 p 开始调用数据 (32 bytes)                             |
| calldatasize()                               |        | F    | 调用数据的大小（以字节为单位）                               |
| calldatacopy(t, f, s)                        | -      | F    | 从位置 f 的 calldata 复制 s 个字节到位置 t 的内存            |
| call(g, a, v, in, insize, out, outsize)      |        | F    | 在地址 a 调用合约 [See more](https://www.axihe.com/#yul-call-return-area) |
| callcode(g, a, v, in, insize, out, outsize)  |        | F    | 与 `call` 相同，但仅使用 a 中的代码，否则留在当前合约的上下文中 [See more](https://www.axihe.com/#yul-call-return-area) |
| delegatecall(g, a, in, insize, out, outsize) |        | H    | 与 `callcode` 相同，但也保留 `caller` 和 `callvalue` [See more](https://www.axihe.com/#yul-call-return-area) |
| staticcall(g, a, in, insize, out, outsize)   |        | B    | 与 `call(g, a, 0, in, insize, out, outsize)` 相同，但不允许状态修改 ons [See more](https://www.axihe.com/#yul-call-return-area) |

### 结束执行

| 操作符号                | 返回值 | 版本 | 解释说明                                                 |
| ----------------------- | ------ | ---- | -------------------------------------------------------- |
| return(p, s)            | -      | F    | 结束执行, return data mem[p…(p+s))                       |
| stop()                  | -      | F    | 结束执行, 类似 `return(0, 0)`                            |
| revert(p, s)            | -      | B    | 结束执行, revert state changes, return data mem[p…(p+s)) |
| selfdestruct(a)         | -      | F    | 结束执行, destroy current contract and send funds to a   |
| invalid()               | -      | F    | 结束执行 with invalid instruction                        |
| returndatasize()        |        | B    | 最后返回数据的大小                                       |
| returndatacopy(t, f, s) | -      | B    | 将 s 个字节从位置 f 的 returndata 复制到位置 t 的 mem    |

### log 信息

| 操作符号                   | 返回值 | 版本 | 解释说明                                             |
| -------------------------- | ------ | ---- | ---------------------------------------------------- |
| log0(p, s)                 | -      | F    | log without topics and data mem[p…(p+s))             |
| log1(p, s, t1)             | -      | F    | log with topic t1 and data mem[p…(p+s))              |
| log2(p, s, t1, t2)         | -      | F    | log with topics t1, t2 and data mem[p…(p+s))         |
| log3(p, s, t1, t2, t3)     | -      | F    | log with topics t1, t2, t3 and data mem[p…(p+s))     |
| log4(p, s, t1, t2, t3, t4) | -      | F    | log with topics t1, t2, t3, t4 and data mem[p…(p+s)) |

# 实战应用

### 例子 1:一个演示

```
library VectorSum {
    // 此函数效率较低，因为优化器当前无法删除数组访问中的边界检查。
    function sumSolidity(uint256[] memory data)
        public
        pure
        returns (uint256 sum)
    {
        for (uint256 i = 0; i < data.length; ++i) sum += data[i];
    }

    // We know that we only access the array in bounds, so we can avoid the check.
    // 0x20 needs to be added to an array because the first slot contains the
    // array length.
    // 我们知道我们只在边界内访问数组，
    // 所以我们可以避免检查。0x20 需要添加到数组，因为第一个槽包含数组长度。
    function sumAsm(uint256[] memory data) public pure returns (uint256 sum) {
        for (uint256 i = 0; i < data.length; ++i) {
            assembly {
                sum := add(sum, mload(add(add(data, 0x20), mul(i, 0x20))))
            }
        }
    }

    // Same as above, but accomplish the entire code within inline assembly.
    function sumPureAsm(uint256[] memory data)
        public
        pure
        returns (uint256 sum)
    {
        assembly {
            // Load the length (first 32 bytes)
            let len := mload(data)

            // Skip over the length field.
            //
            // Keep temporary variable so it can be incremented in place.
            //
            // NOTE: incrementing data would result in an unusable
            //       data variable after this assembly block
            let dataElementLocation := add(data, 0x20)

            // Iterate until the bound is not met.
            for {
                let end := add(dataElementLocation, mul(len, 0x20))
            } lt(dataElementLocation, end) {
                dataElementLocation := add(dataElementLocation, 0x20)
            } {
                sum := add(sum, mload(dataElementLocation))
            }
        }
    }
}
```

### 例子 2:获取合约代码

gas 相差无几，重点看一下 code 的背后原理

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract Demo {
    // 24310 gas
    function codeBySolidity(address _addr)
        public
        view
        returns (bytes memory o_code)
    {
        return _addr.code;
    }

    // 24286 gas
    function codeByAssembly(address _addr)
        public
        view
        returns (bytes memory o_code)
    {
        assembly {
            // 1.使用 extcodesize 获取合约内的代码大小
            let size := extcodesize(_addr)

            // 2.使用 mload 分配输出字节数组
            // 类似 o_code = new bytes（size）
            o_code := mload(0x40)

            // 在 0x40 的位置存入数据
            //      add(size, 0x20) :
            //              size 加 0x20
            //      add(add(size, 0x20), 0x1f)
            //              size 加 0x20,再加 0x1f
            //      not(0x1f)
            //              0x1f 的按位非
            //      and(add(add(size, 0x20), 0x1f), not(0x1f))
            //          "size 加 0x20,再加 0x1f" 和 "0x1f的按位非" 的按位与
            mstore(
                0x40,
                add(o_code, and(add(add(size, 0x20), 0x1f), not(0x1f)))
            )

            // 把长度保存到内存中
            mstore(o_code, size)

            // 实际获取代码，这需要汇编语言
            extcodecopy(_addr, add(o_code, 0x20), 0, size)
        }
    }
}
```

### 例子 3:计算数值数组的和

```
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

library VectorSum {
    // 因为目前的优化器在访问数组时无法移除边界检查，
    // 所以这个函数的执行效率比较低。
    function sumSolidity(uint256[] memory _data)
        public
        pure
        returns (uint256 o_sum)
    {
        for (uint256 i = 0; i < _data.length; ++i) o_sum += _data[i];
    }

    // 我们只能在数组范围内访问数组元素，所以我们可以在内联汇编中不做边界检查。
    // 由于 ABI 编码中数组数据的第一个字（32 字节）的位置保存的是数组长度，
    // 所以我们在访问数组元素时需要加入 0x20 作为偏移量。
    function sumAsm(uint256[] memory _data)
        public
        pure
        returns (uint256 o_sum)
    {
        for (uint256 i = 0; i < _data.length; ++i) {
            assembly {
                o_sum := add(o_sum, mload(add(add(_data, 0x20), mul(i, 0x20))))
            }
        }
    }

    // 和上面一样，但在内联汇编内完成整个代码。
    function sumPureAsm(uint256[] memory _data)
        public
        pure
        returns (uint256 o_sum)
    {
        assembly {
            // 取得数组长度（取前 32 字节）
            let len := mload(_data)

            // 略过长度字段。
            //
            // 保持临时变量以便它可以在原地增加。
            //
            // 注意：对 _data 数值的增加将导致 _data 在这个汇编语句块之后不再可用。
            //      因为无法再基于 _data 来解析后续的数组数据。
            let data := add(_data, 0x20)

            // 迭代到数组数据结束
            for {
                let end := add(data, mul(len, 0x20))
            } lt(data, end) {
                data := add(data, 0x20)
            } {
                o_sum := add(o_sum, mload(data))
            }
        }
    }
}
```


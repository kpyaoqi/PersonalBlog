---
title: 14-Library库

date: 2022-07-20	

categories: Solidity0.8.17	

tags: [区块链,Solidity0.8.17]
---	

library 是智能合约的精简版，像智能合约一样，位于区块链上，包含可以被其他合约使用的代码。库合约有两种使用方法，直接调用和 `using...for...` 使用。

# 库合约与普通智能合约区别

库与合约类似，库的目的是只需要在特定的地址部署一次，而它们的代码可以通过EVM 的 `DELEGATECALL` 特性进行重用。这意味着如果库函数被调用，它的代码在调用合约的上下文中执行，即 `this` 指向调用合约，特别注意，他访问的是调用合约存储的状态。因为每个库都是一段独立的代码，所以它仅能访问调用合约明确提供的状态变量（否则它就无法通过名字访问这些变量）。

- **库不能有任何状态变量**

- **它们也不能继承其他合约**。

- 库合约函数的可视范围通常为internal，可变性为pure，也就是对所有使用它的合约可见。

  - 定义成 `external` 毫无意义，因为库合约函数只在内部使用，不独立运行。
  - 同样定义成 `private` 也不行，因为其它合约无法使用。
  
- 禁止使用 `fallback` / `receive` 函数，所以导致也不能接收以太币

- Library 被销毁后，则所有方法恢复为初始值，功能失效。

- 使用库 library 的合约，可以将库合约视为隐式的父合约，当然它们不会显式的出现在继承关系中。也就是不用写 `is` 来继承，直接可以在合约中使用。

- 按照规范，库合约的名字需要首字母大写（大驼峰命名方式）

可以通过类型转换, 将库类型更改为 `address` 类型, 例如: 使用`address(LibraryName)`,由于编译器无法知道库的部署位置，编译器会生成`__$30bbc0abd4d6364515865950d3e0d10953$__`形式的占位符，该占位符是完整的库名称的 keccak256 哈希的十六进制编码的 34 个字符的前缀，例如：如果该库存储在 libraries 目录中名为 bigint.sol 的文件中，则完整的库名称为`libraries/bigint.sol:BigInt`。

此类字节码不完整的合约，不应该部署。占位符需要替换为实际地址。你可以通过在编译库时将它们传递给编译器或使用链接器更新已编译的二进制文件来实现。有关如何使用命令行编译器进行链接的信息，请参见 [`library-linking`](https://learnblockchain.cn/docs/solidity/using-the-compiler.html#/library-linking) 。

# 直接调用库合约方法

### 案例 1

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

library Math {
    function max(uint256 _x, uint256 _y) internal pure returns (uint256) {
        return _x > _y ? _x : _y;
    }
}

contract Test {
    function testMax(uint256 _x, uint256 _y) external pure returns (uint256) {
        return Math.max(_x, _y);
    }
}
```

### 案例 2

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

library ArrayLib {
    function find(uint256[] storage _arr, uint256 _value)
        internal
        view
        returns (uint256)
    {
        for (uint256 index = 0; index < _arr.length; index++) {
            if (_arr[index] == _value) {
                return index;
            }
        }
        revert("Not Found");
    }
}

contract Test {
    uint256[] public arr = [10, 11, 12, 13, 14, 15, 16, 17, 18, 19];

    // 会成功
    function test1() external view returns (uint256) {
        return ArrayLib.find(arr, 15);
    }

    // 会失败
    function test2() external view returns (uint256) {
        return ArrayLib.find(arr, 99);
    }
}
```

# `using...for...` 使用库合约

使用库合约还有更方便的方法，那就是 `using for` 指令。

例如：`using A for B` 用来将 A 库里定义的函数附着到类型 B。**这些函数将会默认接收调用函数对象的实例作为第一个参数**。
using For 可在文件或合约内部及合约级都是有效的。

核心如下:

```
using ArrayLib for uint256[];
uint256[] public arr = [10, 11, 12, 13, 14,...];
...
arr.find(15); // 直接使用
...
```

### 例子 1

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

library ArrayLib {
    function find(uint256[] storage _arr, uint256 _value)
        internal
        view
        returns (uint256)
    {
        for (uint256 index = 0; index < _arr.length; index++) {
            if (_arr[index] == _value) {
                return index;
            }
        }
        revert("Not Found");
    }
}

contract Test {
    // using for 可以让所有 uint256[] 数据，都具有 ArrayLib 内的方法
    using ArrayLib for uint256[];

    uint256[] public arr = [10, 11, 12, 13, 14, 15, 16, 17, 18, 19];

    function test1() external view returns (uint256) {
        // return ArrayLib.find(arr, 15);

        // 可以直接使用 arr.find，而不需要额外修改 ArrayLib 内的代码
        return arr.find(15);
    }

    function test2() external view returns (uint256) {
        // return ArrayLib.find(arr, 99);
        return arr.find(99);
    }
}
```

### using for 其他用法

第一部分 `A` 可以是以下之一：

- 一些库或文件级的函数列表(`using {f, g, h, L.t} for uint;`)，仅是那些函数被附加到类型。
- 库名称 (`using L for uint;`) ，库里所有的函数(包括 public 和 internal 函数) 被附加到类型上。

在文件级，第二部分 `B` 必须是一个显式类型（不用指定数据位置）

在合约内，你可以使用 `using L for *;`， 表示库 `L`中的函数被附加在所有类型上。

如果你指定一个库，库内所有函数都会被加载，即使它们的第一个参数类型与对象的类型不匹配。类型检查会在函数调用和重载解析时执行。

如果你使用函数列表 (`using {f, g, h, L.t} for uint;`)， 那么类型
(`uint`) 会隐式的转换为这些函数的第一个参数。
即便这些函数中没有一个被调用，这个检查也会进行。

`using A for B;` 指令仅在当前作用域有效（要么是合约中，或当前模块、或源码单元），包括在作用域内的所有函数，在合约或模块之外则无效。

当 `using for` 指令在文件级别使用，并应用于一个用户定义类型（在用一个文件定义的文件级别的用户类型），`global` 关键字可以添加到末尾。产生的效果是，这些函数被附加到使用该类型的任何地方（包括其他文件），而不仅仅是声明处所在的作用域。

在下面的例子中，我们将使用库：

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

library Search {
    function indexOf(uint[] storage self, uint value)
        public
        view
        returns (uint)
    {
        for (uint i = 0; i < self.length; i++)
            if (self[i] == value) return i;
        return type(uint).max;
    }
}

using Search for uint[];

contract C {
    using Search for uint[];
    uint[] data;

    function append(uint value) public {
        data.push(value);
    }

    function replace(uint from, uint to) public {
        // 执行库函数调用
        uint index = data.indexOf(from);
        if (index == type(uint).max)
            data.push(to);
        else
            data[index] = to;
    }
}
```

注意，所有 external 库调用都是实际的 EVM 函数调用。这意味着如果传递内存或值类型，都将产生一个副本，即使是 `self` 变量。 引用存储变量或者 internal 库调用 是唯一不会发生拷贝的情况。

# 直接调用 和 using for 对比

- using for 更符合语义化
- 库合约使用 using for 比直接使用更省 gas

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

library Sum {
    function sum(uint256[] memory _data) public pure returns (uint256 temp) {
        for (uint256 i = 0; i < _data.length; ++i) {
            temp += _data[i];
        }
    }
}

contract Test {
    using Sum for uint256[];
    uint256[] data;

    constructor() {
        data.push(1);
        data.push(2);
        data.push(3);
        data.push(4);
        data.push(5);
    }

    // 43874 gas
    function sumA1() external view returns (uint256) {
        return Sum.sum(data);
    }

    // 43531 gas
    function sumA2() external view returns (uint256) {
        return data.sum();
    }
}
```

# 销毁合约库

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

library Math {
    function max(uint256 _x, uint256 _y) internal pure returns (uint256) {
        return _x > _y ? _x : _y;
    }

    function kill() internal {
        selfdestruct(payable(msg.sender));
    }

    // Libraries cannot have fallback functions.
    // fallback() external {}

    // Libraries cannot have receive ether functions.
    // receive() external payable {}
}

contract Test {
    function testMax(uint256 _x, uint256 _y) external pure returns (uint256) {
        return Math.max(_x, _y);
    }

    function testKill() external {
        return Math.kill();
    }
}
```

- 执行 `Test.testMax`
- 执行 `Test.testKill`
- 再次执行 `Test.testMax`，发现结果是默认值

# 扩展:库的调用保护

如果库的代码是通过 `CALL` 来执行，而不是 `DELEGATECALL` 那么执行的结果会被回退，除非是对 `view` 或者 `pure` 函数的调用。EVM 没有为合约提供检测是否使用 `CALL` 的直接方式，但是合约可以使用`ADDRESS` 操作码找出正在运行的"位置"。生成的代码通过比较这个地址和构造时的地址来确定调用模式。

更具体地说，库的运行时代码总是从一个 push 指令开始，它在编译时是 20 字节的零。当运行部署代码时，这个常数被内存中的当前地址替换，修改后的代码存储在合约中。在运行时，部署时地址就成为了第一个被 push 到堆栈上的常数， 对于任何 non-view 和 non-pure 函数，调度器代码都将对比当前地址与这个常数是否一致。这意味着库在链上存储的实际代码与编译器输出的 `deployedBytecode` 的编码是不同。




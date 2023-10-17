---
title: 节省Gas的合约

date: 2022-06-06	

categories: Gas优化	

tags: [区块链,Gas优化]
---	

# 使用短路模式排序Solidity操作

短路（short-circuiting）是一种使用或/与逻辑来排序不同成本操作的solidity合约 开发模式，它**将低gas成本的操作放在前面**，高gas成本的操作放在后面，这样如果前面的低成本操作可行，就可以跳过（短路）后面的高成本以太坊虚拟机操作了。

```cpp
// f(x) 是低gas成本的操作
// g(y) 是高gas成本的操作

// 按如下排序不同gas成本的操作
f(x) || g(y)
f(x) && g(y)
```

# 删减不必要的Solidity库

在开发Solidity智能合约时，我们引入的库**通常只需要用到其中的部分功能**，这意味着其中可能会包含大量对于你的智能合约而言其实是冗余的solidity代码。如果可以在你自己的合约里安全有效地实现所依赖的库功能，那么就能够达到优化solidity合约的gas利用的目的。

例如，在下面的solidity代码中，我们的以太坊合约只是用到了SafeMath库的`add`方法：

```solidity
import './SafeMath.sol' as SafeMath;

contract SafeAddition {
 function safeAdd(uint a, uint b) public pure returns(uint) {
 return SafeMath.add(a, b);
 }
}
```

通过参考SafeMath的这部分代码的实现，可以把对这个solidity库的依赖剔除掉：

```solidity
contract SafeAddition {
 function safeAdd(uint a, uint b) public pure returns(uint) {
 uint c = a + b;
 require(c >= a, "Addition overflow");
 return c;
 }
}
```

# 精确声明Solidity合约函数的可见性

在Solidity合约开发中，显式声明函数的可见性不仅可以提高智能合约的安全性， 同时也有利于优化合约执行的gas成本。例如，通过显式地标记函数为外部函数（External），可以强制将函数参数的存储位置设置为`calldata`，这会节约每次函数执行时所需的以太坊gas成本。

> External 可见性比 public 消耗gas 少

# 使用适合的数据类型

在Solidity中，有些数据类型要比另外一些数据类型的gas成本高。有必要 了解可用数据类型的gas利用情况，以便根据你的需求选择效率最高的那种。 下面是关于solidity数据类型gas消耗情况的一些规则：

- 在任何可以使用`uint`类型的情况下，不要使用`string`类型
- 存储uint256要比存储uint8的gas成本低，为什么？点击这里查看[原文](https://ethereum.stackexchange.com/questions/3067/why-does-uint8-cost-more-gas-than-uint256)
- 当可以使用`bytes`类型时，不要在solidity合约种使用`byte[]`类型
- 如果`bytes`的长度有可以预计的上限，那么尽可能改用bytes1~bytes32这些具有固定长度的solidity类型
- bytes32所需的gas成本要低于string类型

# 避免Solidity智能合约中的死代码

死代码（Dead code）是指那些永远也不会执行的Solidity代码，例如那些执行条件永远也不可能满足的代码，就像下面的两个自相矛盾的条件判断里的Solidity代码块，消耗了以太坊gas资源但没有任何作用：

```solidity
function deadCode(uint x) public pure {
 if(x < 1) {
    if(x > 2) {
        return x;
    }
 }
}
```

# 避免使用不必要的条件判断

有些条件断言的结果不需要Solidity代码的执行就可以知道结果，那么这样的条件判断就可以精简掉。例如下面的Solidity合约代码中的两级判断条件，内层的判断是在浪费宝贵的以太坊gas资源：

```solidity
function opaquePredicate(uint x) public pure {
 if(x < 1) {
    if(x < 0 ) {  // uint 不可能小于0
    return x;
    }
 }
}
```

# 避免在循环中执行gas成本高的操作

由于`SLOAD`和`SSTORE`操作码的成本高昂，因此管理storage变量的gas成本 要远远高于内存变量，所以要避免在循环中操作storage变量。例如下面的 solidity 代码中，`num`变量是一个storage变量，那么未知循环次数的若干次操作，很可能会造成solidity开发者意料之外的以太坊gas消耗黑洞：

```solidity
uint num = 0;

function expensiveLoop(uint x) public {
  for(uint i = 0; i < x; i++) {
    num += 1;
  }
}
```

解决上述以太坊合约代码问题的方法，是创建一个solidity临时变量 来代替上述全局变量参与循环，然后在循环结束后重新将临时变量的值赋给全局状态变量：

```solidity
uint num = 0;

function lessExpensiveLoop(uint x) public {
  uint temp = num;
  for(uint i = 0; i < x; i++) {
    temp += 1;
  }
  num = temp;
}
```

# 避免使用常量结果的循环

如果一个循环计算的结果是无需编译执行Solidity代码就可以预测的，那么 就不要使用循环，这可以可观地节省gas。例如下面的以太坊合约代码就可以 直接设置num变量的值：

```solidity
function constantOutcome() public pure returns(uint) {
  uint num = 0;
  for(uint i = 0; i < 100; i++) {
    num += 1;
  }
  return num;
}
```

# 合并循环

有时候在Solidity智能合约中，你会发现两个循环的判断条件一致，那么在这种情况下就没有理由不合并它们。例如下面的以太坊合约代码：

```solidity
function loopFusion(uint x, uint y) public pure returns(uint) {
  for(uint i = 0; i < 100; i++) {
    x += 1;
  }
  for(uint i = 0; i < 100; i++) {
    y += 1;
  }
  return x + y;
}
```

# 避免循环中的重复计算

如果循环中的某个Solidity表达式在每次迭代都产生同样的结果，那么就可以将其 移出循环先行计算，从而节省掉循环中额外的gas成本。如果表达式中使用的变量是storage变量， 这就更重要了。例如下面的智能合约代码中表达式`a*b`的值，并不需要每次迭代重新计算：

```solidity
uint a = 4;
uint b = 5;
function repeatedComputations(uint x) public returns(uint) {
  uint sum = 0;
  for(uint i = 0; i <= x; i++) {
    sum = sum + a * b;
  }
}
```

# 去除循环中的比较运算

如果在循环的每个迭代中执行比较运算，但每次的比较结果都相同，则应将其从循环中删除。

```solidity
function unilateralOutcome(uint x) public returns(uint) {
  uint sum = 0;
  for(uint i = 0; i <= 100; i++) {
    if(x > 1) {
      sum += 1;
    }
  }
  return sum;
} 
```

# 不要存储不必要的数据

这听起来似乎很明显，但是非常值得一提。编写智能合约时，你应该只存储交易验证所需的内容。与合约逻辑无关的交易记录或详细说明之类的数据可能不需要保存在合约存储中。

# 将多个小变量打包到单个字中

> 译者注：标题中的"字", 也称为字长，表示每个指令操作的数据长度。

EVM在32字节字长存储模型下运行。可以将小于32个字节的多个变量打包到一个存储槽中，以最大程度地减少`SSTORE`操作码的数量。尽管Solidity [自动尝试将小的基本类型打包到同一插槽中](https://learnblockchain.cn/docs/solidity/internals/layout_in_storage.html)，但是糟糕的结构体成员排序可能会阻止编译器执行此操作。考虑下面的`Good`和`Bad`结构体。

![Image for post](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Gas优化/img/MGf38d9w.png)

第一个`doBad()`函数调用执行消耗约60,000 Gas，而`doGood()`仅消耗约40,000 Gas。注意是一个字长存储的差异(20,000 Gas)，因为`Good`结构将两个uint128打包为一个字。

# 在合约的字节码中存储值

一种相对便宜的存储和读取信息的方法是，将信息部署在区块链上时，直接将其包含在智能合约的字节码中。不利之处是此值以后不能更改。但是，用于加载和存储数据的 gas 消耗将大大减少。有两种可能的实现方法：

1. 将变量声明为 *constant* 常量 (声明为 [immutable](https://learnblockchain.cn/article/1059) 同样也可以降低 gas)
2. 在你要使用的任何地方对其进行硬编码。

```js
uint256 public v1;
uint256 public constant v2;

function calculate() returns (uint256 result) {
    return v1 * v2 * 10000
}
```

变量*v1* 是合约状态的一部分，而*v2*和*1000*是合约字节码的一部分。

# 带有`_`数字的变量，读取时候花费 gas 更少

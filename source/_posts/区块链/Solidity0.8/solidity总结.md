---
title: solidity总结

date: 2022-07-28	

categories: Solidity0.8.17	

tags: [区块链,Solidity0.8.17]
---	

# 初识

### 1.receive 和 fallback 共存的调用

```solidity
/**
    调用时发送了ETH
            |
判断 msg.data 是否为空
          /     \
        是       否
是否存在 receive   fallbak()
      /   \
    存在   不存在
    /        \
receive()   fallbak()
 */
```

---

# 数据

### 2.即便一个合约的代码中没有显式地调用 `selfdestruct`，它仍然有可能通过 `delegatecall` 或 `callcode` 执行自毁操作。

---

### 3.合约进行`selfdestruct`后，还可以调用状态变量和函数么？

> 可以调用，但是返回默认值。如果想调用，也可以在存档节点里指定未删除的高度进行调用。

---

### 4.三种 call 的总结(call/delegatecall/staticcall)

为了与不知道 ABI 的合约进行交互，Solidity 提供了函数 `call`/`delegatecall`/`staticcall` 直接控制编码。它们都带有一个 `bytes memory` 参数和返回执行**成功状态**（bool）和**数据**（bytes memory）。

函数 `abi.encode`，`abi.encodePacked`，`abi.encodeWithSelector` 和 `abi.encodeWithSignature` 可用于编码结构化数据。

**它们可以接受任意类型，任意数量的参数**。这些参数会被打包到以 32 字节为单位的连续区域中存放。其中一个**例外是当第一个参数被编码成正好 4 个字节的情况**。 在这种情况下，这个参数后边不会填充后续参数编码，以允许使用函数签名。

下面具体的介绍三种 call。

#### call()

用给定的有效载荷（payload）发出低级 `CALL` 调用，并**返回交易成功状态和返回数据**（调用合约的方法并转账）, 格式如下：

```
<address>.call(bytes memory) returns (bool, bytes memory)
```

1. 低级CALL调用：不需要 payable address, 普通地址即可

   1. 注意: 调用 `call` 的时候，地址可以不具备 payable 属性

2. 返回两个参数，一个 `bool` 值代表成功或者失败，另外一个是可能存在的 `data`

3. 发送所有可用 gas，也可以自己调节 gas。

   1. 如果 `fallback` 和 `receive` 内的代码相对复杂也可以，但是如果是恶意代码，需要考虑消耗的 gas 是否值得执行。
   2. `_ads.call{value: msg.value,gas:2300}(data)`

4. 当合约调用合约时，不知道对方源码和 ABI 时候，可以使用 call 调用对方合约

5. 推荐使用 call 转账 ETH，但是不推荐使用 call 来调用其他合约。

   1. 原因是: call 调用的时候，将合约控制权交给对方，如果碰到恶意代码，或者不安全的代码就很容易凉凉。

6. 当调用不存在的合约方法时候，会触发对方合约内的fallback或receive

   1. 我们的合约也可以在 `fallback` / `receive` 这两个方法内抛出事件，查看是否有人对其做了什么操作。

7. 三种方法都提供 `gas` 选项，而 `value` 选项仅 `call` 支持 。三种 call 里只有 `call` 可以进行 ETH 转账，其他两种不可以进行转账。

**例子 （重要）：调用其他合约方法**

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Test1 {
    string public name;
    uint256 public age;
    address public owner;

    function setNameAndAge(string memory name_, uint256 age_)
        external
        payable
        returns (string memory __name, uint256 __age)
    {
        name = name_;
        age = age_;
        owner = msg.sender;
        return (name_, age_);
    }
}

contract CallTest {
    // 需要一个网页，动态的解析 _bys
    bytes public bys;

    function call_Test1_setNameAndAge(
        address ads_,
        string memory name_,
        uint256 age_
    ) external payable {
        bytes memory data = abi.encodeWithSignature(
            "setNameAndAge(string,uint256)",
            name_,
            age_
        );
        (bool success, bytes memory _bys) = ads_.call{value: msg.value}(data);
        require(success, "Call Failed");
        bys = _bys;
    }
}
```

简单说下这个例子的原理

```solidity
/**
    普通调用:用户A 调用 callB 合约, 发送 100 wei ; 
    callB 调用 Test1, 发送 50 wei此时在 Test1 合约内部
        msg.sender = B
        msg.value = 50
        Test1 内部如果有状态变量修改，则会被修改
        发送到 Test1 内的ETH主币也会被留在Test1内
 */
```

#### delegatecall() 委托调用

发出低级函数 `DELEGATECALL`，失败时返回 false，发送所有可用 gas，也可以自己调节 gas。

```
<address>.delegatecall(bytes memory) returns (bool, bytes memory)
```

`delegatecall` 使用方法和 `call` 完全一样。区别在于，`delegatecall` 只调用给定地址的代码（函数），其他状态属性如（存储，余额 …）都来自当前合约。`delegatecall` 的目的是使用另一个合约中的库代码。

委托调用是：**委托对方调用自己数据的**。类似授权转账，比如我部署一个 Bank 合约， 授权 ContractA 使用 Bank 地址内的资金，ContractA 只拥有控制权，但是没有拥有权。

- 委托调用后，所有变量修改都是发生在委托合约内部，并不会保存在被委托合约中。
  - 利用这个特性，可以通过更换被委托合约，来升级委托合约。
- 委托调用合约内部，需要和被委托合约的内部参数完全一样，否则容易导致数据混乱
  - 可以通过顺序来避免这个问题，但是推荐完全一样

**例子 1（重要）**

代码如下:

- `DelegateCall` 是委托合约
- `TestVersion1` 是第 1 次被委托合约
- `TestVersion2` 是第 2 次被委托合约

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

// 合约版本 V.1
contract TestVersion1 {
    address public sender;
    uint256 public value;
    uint256 public num;

    function set(uint256 num_) external payable {
        sender = msg.sender;
        value = msg.value;
        num = num_;
    }
}

// 合约版本 V.2
contract TestVersion2 {
    address public sender;
    uint256 public value;
    uint256 public num;

    function set(uint256 num_) external payable {
        sender = msg.sender;
        value = msg.value;
        num = num_ * 2;
    }
}

// 委托调用测试
contract DelegateCall {
    address public sender;
    uint256 public value;
    uint256 public num;

    function set(address _ads, uint256 num_) external payable {
        sender = msg.sender;
        value = msg.value;
        num = num_;
        // 第1种 encode
        // 不需知道合约名字，函数完全自定义
        bytes memory data1 = abi.encodeWithSignature("set(uint256)", num_);
        // 第2种 encode
        // 需要合约名字，可以避免函数和参数写错
        bytes memory data2 = abi.encodeWithSelector(TestVersion1.set.selector, num_);

        (bool success, bytes memory _data) = _ads.delegatecall(data2);

        require(success, "DelegateCall set failed");
    }
}
```

简单说下这个例子的原理

```solidity
/**
    委托调用
    用户A 调用 DelegateCall 合约, 发送 100 wei ; DelegateCall 委托调用 Test1
    此时在 Test1 合约内部
        msg.sender = A
        msg.value = 100
        Test1 内部如果有状态变量修改，也不会被修改，会在DelegateCallB 内改变
        发送到 Test1 内的ETH主币，会被留在 DelegateCallB 内，不会在Test1 内
 */
```

#### staticcall() 静态调用

用给定的有效载荷（payload）发出低级 `STATICCALL` 调用，并**返回交易成功状态和返回数据**（调用合约的方法并转账）

```
<address>.staticcall(bytes memory) returns (bool, bytes memory)
```

它与 call 基本相同，发送所有可用 gas，也可以自己调节 gas，**但如果被调用的函数以任何方式修改状态变量，都将回退**。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

// 被调用的合约
contract Hello1 {
    function echo() external pure returns (string memory) {
        return "Hello World!";
    }
}

contract Hello2 {
    uint8 public a;
    function echo() external returns (string memory) {
        a = 1;
        return "Hello World!";
    }
}

// 调用者合约
contract SoldityTest {
    function callHello(address ads_) external view returns (string memory) {
        // 编码被调用者的方法签名
        bytes4 methodId = bytes4(keccak256("echo()"));

        // 调用合约
        (bool success, bytes memory data) = ads_.staticcall(
            abi.encodeWithSelector(methodId)
        );
        if (success) {
            return abi.decode(data, (string));
        } else {
            return "error";
        }
    }
}
```

1. `call` ， `delegatecall` 和 `staticcall` 是非常低级的函数，应该把它们当作最后一招来使用，破坏了 Solidity 的类型安全性。
2. 三种方法都提供 `gas` 选项，而 `value` 选项仅 `call` 支持 。三种 call 里只有 `call` 可以进行 ETH 转账，其他两种不可以进行转账
2. 如果在通过低级函数 `delegatecall` 发起调用时需要访问存储中的变量，那么这两个合约的存储布局需要一致，以便被调用的合约代码可以正确地通过变量名访问合约的存储变量。 这不是指在库函数调用（高级的调用方式）时所传递的存储变量指针需要满足那样情况。

---

### 5.调用合约时，不知道对方源码和 ABI 时候，可以使用 call 调用对方合约

---

### 6.transfer / send / call 三种转账的总结

低级CALL调用**不需要`payable address`**,ransfer 和 send **只能是 `payable address`**

transfer()失败时抛出异常,send()失败时仅会返回false,不会终止执行（合约地址转账）

call的 gas 可以动态调整,transfer 和 send 只能是固定制 `2300`

call除了可以转账外，可以还可以调用不知道 ABI 的方法，还可以调用的时候转账

- 当调用不存在的合约方法时候，会触发对方合约内的 `fallback` 或者 `receive`。
- 如果使用 `_to.call{value: 100}(data)`，那么`data`中被调用的方法必须添加 `payable` 修饰符，否则转账失败！
- 因为可以调用方法，所以 call 有两个参数，除了一个 `bool` 值代表成功或者失败，另外一个是可能存在的 `data`，比如创建合约时候得到部署的地址，调用函数时候得到的函数放回值

---

### 7.所有的引用类型，都有数据位置这个额外的注解来指定存储在哪里，所以一定要掌握好

### **7.比较字符串和字节**

- `keccak256(abi.encodePacked(s1)) == keccak256(abi.encodePacked(s2))`
- `keccak256(bytes(s1)) == keccak256(bytes(s2))`:更省 gas

---

### 8.堆栈是由 EVM (Ethereum 虚拟机)维护的非持久性数据。EVM 使用堆栈数据位置在执行期间加载变量。堆栈位置最多有 1024 个级别的限制。

---

### 9.数据位置总结

- storage: 存储区: 状态变量总是储存在**存储区**上
- memory: 内存区: 局部变量使用，只在内存中生效。
  - 值类型的局部变量，存储在**内存**中。
  - **引用类型局部变量，需要显式地指定数据位置**。
  - 函数的输入参数如果是数组或者 string，必须是 `memory` 或 `calldata`
  - 内存中的数组必须是定长数组（不能使用 push 赋值），动态数组只能储存在状态变量中。
- calldata
  - 和 memory 类似，但是 calldata 只能用在函数的输入参数中。
  - 相比使用 memory ,合约输入参数如果使用 calldata, 可以节约 gas

按照函数参数:

- 内部函数参数: (包括返回参数)都存储在**memory（内存）**中。
- 外部函数参数: (不包括返回参数)存储在 `calldata` 中

---

### 10.为什么映射不能像哈希表一样遍历？

映射与哈希表不同的地方：**在映射中,并不存储 key，而是存储它的 `keccak256` 哈希值，从而便于查询实际的值**。正因为如此，映射是没有长度的，也没有 `key 的集合`或 `value 的集合`的概念。映射只能是存储的数据位置，因此只允许作为状态变量或作为函数内的存储引用 或 作为库函数的参数。

---

### 11.遍历所有 Mapping 内的数据，（Mapping 配合 array ）

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Demo {
    mapping(address => uint256) public balances;
    // 用于检查:地址是否已经存在于 balancesKey
    mapping(address => bool) public balancesInserted;
    address[] public balancesKey; // 所有地址

    // 设置
    function set(address ads_,uint256 amount_) external{
        balances[ads_] = amount_;
        // 1.检查
        if(!balancesInserted[ads_]){
            // 2.修改检查条件
            balancesInserted[ads_] = true;
            // 3.正在的操作
            balancesKey.push(ads_);
        }
    }
    // 获取
    function get(uint256 index_) external view returns(uint256){
        require(index_<balancesKey.length,"index_ error");
        return balances[balancesKey[index_]];
    }
    // 获取所有
    function totalAddress() external view returns(uint256){
        return balancesKey.length;
    }

    // 获取第一个值
    function first() external view returns(uint256){
        return balances[balancesKey[0]];
    }
    // 最后一个值
    function latest() external view returns(uint256){
        return balances[balancesKey[balancesKey.length-1]];
    }
}
```

- 更完善的实现: https://github.com/ethereum/dapp-bin/blob/master/library/iterable_mapping.sol
- 更新的实现: https://learnblockchain.cn/docs/solidity/types.html#iterable-mappings

---

### 12.聊一聊 `checked` 和 `unchecked`

0.8.0 开始，算术运算有两种计算模式：一种是`checked`（检查）模式，另一种是 `unchecked`（不检查）模式。 默认情况下，算术运算在 `checked` 模式下，即都会进行溢出检查，如果结果落在取值范围之外，调用会通过 失败异常 回退。 你也可以通过 `unchecked{ ... }` 切换到 “unchecked”模式，更多可参考 `unchecked` .

---

# 变量

### 13.constant常量 和 immutable 不可变量总结

####  值的确定时机不同

状态变量声明为 constant (常量)或者 immutable （不可变量），在这两种情况下，合约一旦部署之后，变量将不在修改。

- 对于 constant 常量, 他的值在编译器确定.
- 对于 immutable, 它的值在部署时确定。

####  gas 不同

与常规状态变量相比，常量和不可变量的 gas 成本要低得多。

- 对于常量，赋值给它的表达式将复制到所有访问该常量的位置，并且每次都会对其进行重新求值。这样可以进行本地优化。
- 不可变变量在构造时进行一次求值，并将其值复制到代码中访问它们的所有位置。 对于这些值，将保留 32 个字节，即使它们适合较少的字节也是如此。很多时候常量的 gas 更低。

如果可以使用常量的时候，推荐使用常量。

####  支持的数据不同

不是所有类型的状态变量都支持用 `constant` 或 `immutable` 来修饰

- 当前`constant`仅支持值类型和引用类型中的 string 和 bytes
- `immutable`仅支持值类型

---

### 14.`private` / `external` / `internal` / `public`

- 函数有四可见型，分别是 `private` / `external` / `internal` / `public`
- 状态变量可以有三种可见型，分别是 `private` / `internal` / `public`

1. **private**: 私有，仅在当前合约中可以访问，在继承的合约内不可访问
2. **internal(默认)**: 能在它们所定义的合约和派生合同中访问,它们不能被外部访问,需要注意的是不能加前缀 this，前缀 this 是表示通过外部方式访问
3. **external:** 不能声明在状态变量上，只能标识在函数上,只能从智能合约外部调用。 如果要从智能合约中调用它，则必须使用this
4. **public**: 公开可视(合约内部，被继承的，外部都可以调用)

---

### 15.变量作用域的规则

- 变量将会从它们被声明之后可见，直到一对 `{}` 块的结束。
- 对于参数形式的变量（例如：函数参数、修饰器参数、catch 参数等等）在其后接着的代码块内有效。
  - 这些代码块是函数的实现，catch 语句块等。
  - 有一个例外，在 for 循环语句中初始化的变量，其可见性仅维持到 `for` 循环的结束。

---

# 函数

### 16.extcodesize 操作码来检查要调用的合约是否确实存在?

由于 EVM 认为可以调用不存在的合约的调用，因此在 Solidity 语言层面里会使用 extcodesize 操作码来检查要调用的合约是否确实存在（包含代码），如果不存在该合约，则抛出异常。如果返回数据在调用后被解码，则跳过这个检查，因此 ABI 解码器将捕捉到不存在的合约的情况。请注意，这个检查在 **低级 call** 时不被执行，这些调用是对地址而不是合约实例进行操作。

---

### 17.mutability状态可变性(Pure、view、payable)

1. **pure: 既不读取也不修改状态变量,这种函数被称为纯函数**

   - 读取状态变量。
     - 这也意味着读取 `immutable` 变量也不是一个 `pure` 操作。
   - 访问 `address(this).balance` 或 `<address>.balance`
   - 访问 `block`，`tx`， `msg` 中任意成员 （除 `msg.sig` 和 `msg.data` 之外）。
   - 调用任何未标记为 `pure` 的函数。
   - 使用包含特定操作码的内联汇编。
     - `TODO:` 这个不了解，需要用例子加深印象。
   - 使用操作码 `STATICCALL` , 这并不保证状态未被读取, 但至少不被修改。

2. **view: 读取状态变量，但是不修改状态变量,这种函数被称为视图函数**

   - 修改状态变量。
   - 触发事件。
   - 创建其它合约。
   - 使用 `selfdestruct`。
   - 通过调用发送以太币。
   - 调用任何没有标记为 view 或者 pure 的函数。
   - 使用底层调用
     - (TODO:这里是 call 操作么？)
   - 使用包含某些操作码的内联程序集。

   - 状态变量的 Getter 方法默认是 view 函数。

3. **payable：用 payable 声明的函数可以接受发送给合约的以太币.**

   - 如果未指定，该函数将自动拒绝所有发送给它的以太币

---

### 18.函数的构造函数有什么特点？

它仅能在智能合约部署的时候调用一次，创建之后就不能再次被调用。

构造函数是可选的，只允许有一个构造函数，这意味着不支持重载。（普通函数支持重载）

在合约创建的过程中，它的代码还是空的，所以直到构造函数执行结束，我们都不应该在其中调用合约自己的函数。(可以直接写函数名调用，但是不推荐调用，不可以通过 this 来调用函数，因为此时真实的合约实例还没有被创建。)

---

# 运算操作符

### 19. `A && B`,如果 A 为 false，B 就不执行了,`A || B`,如果 A 为 true，B 就不执行了合理的使用短路操作，可以省一些 gas 费。

---

### 20.三元运算符

三元运算符的结果类型是由两个操作数的类型决定的，方法与上面一样，如果需要的话，首先转换为它们的最小可容纳类型。

因此， `255 + (true ? 1 : 0)` 将由于算术溢出而被回退。 原因是 `(true ? 1 : 0)` 是 uint8 类型，这迫使加法也要在 `uint8` 中执行。 而 `256` 超出了这个类型所允许的范围。

另一个结果是，像 `1.5 + 1.5` 这样的表达式是有效的，但 `1.5 + (true ? 1.5 : 2.5)` 则无效。 这是因为前者是以无限精度来进行有理表达式运算，只有它的最终结果值才是重要的。 后者涉及到将小数有理数转换为整数，这在目前是不允许的。

---

# 错误处理

### 21.require、assert、revert对比

`require(false)` 会退还所有剩余的 gas，同时可以返回一个自定义的报错信息。

`assert(false)` 会消耗掉所有剩余的 gas，并恢复所有的操作。

 `revert()` 会触发一个没有任何错误数据的回退

以下三个语句的功能完全相同：

```solidity
// revert
if(msg.sender != owner) {
   revert error();
 }
 
// require
require(msg.sender == owner,"error");

// assert
assert(msg.sender == owner);
```

---

### 22.自定义error

**`error` 只能通过 `revert` 触发**

错误产生的数据，会通过 revert 操作传递给调用者，可以交由链外组件处理或在 try/catch 语句 中捕获它。

注意，只有外部调用的错误才能被捕获。发生在内部调用或同一函数内的 revert 不能被捕获。

---

### 23.try catch

Solidity 如果遇到异常错误，是通过回退状态的方式来进行处理。发生异常时，会撤消当前调用和所有子调用改变的状态变量，同时给调用者返回一个错误标识。

调用者调用某个函数方法，要么成功修改了所有状态变量，要么遇到异常不修改任何状态变量，不存在成功修改部分变量的情况，

Solidity 提供了 **require** 、**assert** 和 **revert** 来处理异常。同时可以使用 `error` 关键字来实现错误。

跟用错误字符串相比， error 更便宜并且允许你编码额外的数据，还可以用 `NatSpec` 为用户去描述错误。

Solidity 使用状态恢复异常来处理错误。这种异常将撤消对当前调用（及其所有子调用）中的状态所做的所有更改，并且还向调用者标记错误。

如果异常在子调用发生，那么异常会自动冒泡到顶层（例如：异常会重新抛出），除非他们在 `try/catch` 语句中捕获了错误。 但是如果是在 `send` 和 低级 `call`, `delegatecall` 和 `staticcall` 的调用里发生异常时， 他们会返回 `false` （第一个返回值） 而不是冒泡异常。

警告注意：根据 EVM 的设计，如果被调用的地址不存在，低级别函数 `call`, `delegatecall` 和 `staticcall` 第一个返回值同样是 `true`。 如果需要，请在调用之前检查账号的存在性。

异常可以包含错误数据，以 error 示例 的形式传回给调用者。 内置的错误 `Error(string)` 和 `Panic(uint256)` 被作为特殊函数使用，下面将解释。 `Error` 用于 “常规” 错误条件，而 `Panic` 用于在（无 bug）代码中不应该出现的错误。

函数 assert 和 require 可用于检查条件并在条件不满足时抛出异常。

---

# 事件

### 24. 带 indexed 参数名的 event

这种事件也被称为**索引事件**

语法:`event EventName(TypeName indexed varibleName....);`

事件中 indexed 标记过的参数，可以在链外进行搜索查询:

主要用在链下服务，可以通过 RPC 获取，比如 web3 的以下方法:

- ```js
  myContract.once //订阅一个事件并在第一次事件触发或错误发生后立即取消订阅
  ```

- ```solidity
  myContract.events.MyEvent //订阅指定的合约事件
  ```

- ```solidity
  myContract.getPastEvents //读取合约历史事件
  ```

一个事件中 indexed 标记过的参数最多有 3 个。

---

# 合约继承

### 25.注意：

- 父合约必须写在子合约的前面，否则会报错
- 子类可以访问父类的权限修饰符只有：`public/internal`，不能是 `external/private`
- 多重继承中不允许出现相同的**函数名**、**事件名**、**修改器名**以及**状态变量名**等

### 26.`virtual`, `override` , `abstract`

父合约标记为 `virtual`函数可以在继承合约里重写(overridden)以更改他们的行为。重写的函数需要使用关键字 `override` 修饰。

基础合约中可以包含没有实现代码的函数，也就是纯虚函数，那么基础合约必须声明为 `abstract`。

# 接口Interface

### 27.限制

- 无法实现任何功能，没有函数体。

- 无法定义构造函数。

- 无法定义状态变量。

- 无法定义结构（`struct`）（`0.5.0` 版本开始接口里可以支持声明 `enum` 类型）。

- 不可以声明修改器。

- 所有声明的函数必须是external的，尽管在合约里可以是 public


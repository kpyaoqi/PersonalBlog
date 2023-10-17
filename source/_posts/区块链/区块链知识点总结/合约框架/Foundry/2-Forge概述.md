---
title: 2-Forge概述

date: 2023-02-14	

categories: Foundry	

tags: [区块链,区块链知识点总结,合约框架,Foundry]
---	

# Forge

Forge 是 `Foundry` 附带的命令行工具。 Forge 可用来测试、构建和部署您的智能合约。

## Forge概述

### 测试

**使用 forge test命令运行测试（用例）。 所有测试都是用 Solidity 编写的**

任何**具有以`test`开头的函数的合约都被认为是一个测试**。 通常，测试将按照约定**放在 `src/test` 中，并以 `.t.sol` 结尾**。

您还可以通过传递过滤器来运行特定测试：

```sh
$ forge test --match-contract ComplicatedContractTest --match-test testDeposit
```

这将在**名称中带有 `testDeposit` 的 `ComplicatedContractTest` 测试合约中运行测试**。 这些标志的反向版本也存在（`--no-match-contract` 和 `--no-match-test`）。

您可以**使用 `--match-path` 与 glob 模式匹配的文件名中运行测试**，`--match-path` 标志的反面是 `--no-match-path`。

```sh
$ forge test --match-path test/ContractB.t.sol
```

### 日志和跟踪

`forge test` 的默认行为是只显示通过和失败测试的摘要。 您可以通过增加详细程度（使用`-v`标志）来控制此行为。 每个详细级别都会添加更多信息：

- **级别 2 (`-vv`)**：还会显示测试期间发出的日志。 这包括来自测试的断言错误，显示诸如预期与实际结果等之类的信息。
- **级别 3 (`-vvv`)**：还显示失败测试的堆栈跟踪。
- **级别 4 (`-vvvv`)**：显示所有测试的堆栈跟踪，并显示失败测试的设置（setup）跟踪。
- **级别 5 (`-vvvvv`)**：始终显示堆栈跟踪和设置（setup）跟踪。

### Watch模式

当您使用`forge test --watch`对文件进行更改时，Forge 可以重新运行您的测试。

默认情况下，仅重新运行更改的测试文件。 如果你想重新运行更改的所有测试，你可以使用 `forge test --watch --run-all`。

## 测试

### 编写测试

使用Forge 标准库的 `Test` 合约，这是使用 Forge 编写测试的首选方式。

```solidity
import "forge-std/Test.sol";
```

Forge 在测试中使用以下关键字：

- `setUp`：在**每个测试用例运行之前调用的可选函数**

```solidity
    function setUp() public {
        testNumber = 42;
    }
```

- `test`：**以 `test` 为前缀的函数作为测试用例运行**

```solidity
    function testNumberIs42() public {
        assertEq(testNumber, 42);
    }
```

- `testFail`: `test` 前缀的测试的反面 - **如果函数没有 revert，则测试失败**

```solidity
    function testFailSubtract43() public {
        testNumber -= 43;
    }
```

> **测试函数必须具有`external`或`public`可见性**。 声明为`internal`或 `private` 不会被 Forge 选中，即使它们以 `test` 为前缀。

### 共享设置

通过创建辅助抽象合约并在测试合约中继承它们来使用共享设置

```solidity
abstract contract HelperContract {
    address constant IMPORTANT_ADDRESS = 0x543d...;
    SomeContract someContract;
    constructor() {...}
}

contract MyContractTest is Test, HelperContract {
    function setUp() public {
        someContract = new SomeContract(0, IMPORTANT_ADDRESS);
        ...
    }
}

contract MyOtherContractTest is Test, HelperContract {
    function setUp() public {
        someContract = new SomeContract(1000, IMPORTANT_ADDRESS);
        ...
    }
}
```

### 作弊码(Cheatcodes)

作弊码允许您更改区块号、您的身份（地址）等。 它们是通过**在特别指定的地址上调用特定函数来调用的**。

通过 Forge 标准库的 `Test` 合约中**提供的 `vm` 实例轻松访问作弊码**

使用`prank` 作弊码将我们的身份更改为零地址再进行下一次调用 (`upOnly.increment()`)：

```solidity
vm.prank(address(0));
upOnly.increment();
//测试失败
```

使用`expectRevert`作弊码来验证我们不是所有者的情况：

```solidity
vm.expectRevert(Unauthorized.selector);
vm.prank(address(0));
upOnly.increment();
//测试通过，因为获得了期望的Unauthorized的错误
```

另一个可能不那么直观的作弊码是 `expectEmit` 函数：

```solidity
event Transfer(address indexed from, address indexed to, uint256 amount);
vm.expectEmit(true, true, false, true);
// The event we expect
emit Transfer(address(this), address(1337), 1337);
// The event we get
emitter.t();
//在emitter.t()事件中期望是from是address(this)，而to是address(1337)

event Transfer(address indexed from, address indexed to, uint256 amount);
function t() public {
	emit Transfer(msg.sender, address(1337), 1337);
}
```

`expectEmit` 中的第三个参数设置为 `false`，因为不需要检查 `Transfer` 事件中的第三个主题，因为只有两个主题。 即使我们设置为 `true` 也没关系，**`expectEmit` 中的第 4 个参数设置为 `true`**，这**意味着我们要检查 "non-indexed topics（非索引主题）"，**也称为数据。

### 理解跟踪

Forge 可以为失败的测试（`-vvv`）或所有测试（`-vvvv`）生成跟踪(Traces)。

跟踪遵循相同的通用格式：

```ignore
[<Gas Usage>] <Contract>::<Function>(<Parameters>)
  ├─ [<Gas Usage>] <Contract>::<Function>(<Parameters>)
  │   └─ ← <Return Value>
  └─ ← <Return Value>
```

每个跟踪可以有更多的子跟踪（subtraces），每个 subtraces 表示对合约的调用和返回值。

如果您的终端支持颜色，跟踪也会有多种颜色：

- **绿色**：对于不会 revert 的调用
- **红色**：用于有 evert 的调用
- **蓝色**：用于调用作弊码
- **青色**：用于触发日志
- **黄色**：用于合约部署

### 分叉测试(分叉模式/分叉作弊码)

#### 分叉模式

要在分叉环境（例如分叉的以太坊主网）中运行所有测试，请通过 --fork-url 标志传递 RPC URL：

```bash
$ forge test --fork-url <your_rpc_url>
```

以下值会更改以反映分叉时链的值：

block_number、chain_id、gas_limit、gas_price、block_base_fee_per_gas、block_coinbase、block_timestamp、block_difficulty

可以使用 `--fork-block-number` 指定要从中分叉的区块高度：

```bash
$ forge test --fork-url <your_rpc_url> --fork-block-number 1
```

当您需要与现有合约进行交互时，分叉特别有用。 您可以选择以这种方式进行集成测试，就好像您在实际网络上一样。

##### 缓存（Caching）

如果同时指定了 `--fork-url` 和 `--fork-block-number`，那么该块的数据将被缓存以供将来的测试运行。

数据缓存在 `~/.foundry/cache/rpc/<chain name>/<block number>` 中。 

**要清除缓存，只需删除目录或运行 forge clean**（删除所有构建工件和缓存目录）。

也可以**通过传递 --no-storage-caching 或通过配置 no_storage_caching 和 foundry.toml 完全忽略缓存 rpc_storage_caching。**

##### 已改进的跟踪 traces

Forge 支持使用 Etherscan 在分叉环境中识别合约。

要使用此功能，请**通过 `--etherscan-api-key` 标志传递 `Etherscan API `密钥：**

```bash
$ forge test --fork-url <your_rpc_url> --etherscan-api-key <your_etherscan_api_key>
```

或者，您可以设置 `ETHERSCAN_API_KEY` 环境变量。

#### 分叉作弊码

在逐个测试的基础上使用分叉模式，并在测试中使用多个分叉，每个分叉都通过其自己唯一的 `uint256` 标识符进行识别。

##### 用法

用 **createFork** 创建两个分叉，没有在初始选择一个，**每个 fork 有一个唯一标识符 (uint256 forkId)**，该标识符在首次创建时分配。

**通过将该 forkId 传递给 selectFork 来启用特定的分叉。**createSelectFork 是 createFork 加上 selectFork 的单行代码。**一次只能有一个活动分叉，**当前活动分叉的标识符**可以通过 activeFork 检索**。类似于 roll，您可以**使用 rollFork 设置分叉的 block.number。**

```solidity
contract ForkTest is Test {
    uint256 mainnetFork;
    uint256 optimismFork;
    
    function setUp() public {
        mainnetFork = vm.createFork(MAINNET_RPC_URL);
        optimismFork = vm.createFork(OPTIMISM_RPC_URL);
    }

    function testForkIdDiffer() public {
        assert(mainnetFork != optimismFork);
    }

    function testCanSelectFork() public {
        vm.selectFork(mainnetFork);
        assertEq(vm.activeFork(), mainnetFork);
    }

    function testCanSwitchForks() public {
        vm.selectFork(mainnetFork);
        assertEq(vm.activeFork(), mainnetFork);
        vm.selectFork(optimismFork);
        assertEq(vm.activeFork(), optimismFork);
    }

    function testCanCreateAndSelectForkInOneStep() public {
        uint256 anotherFork = vm.createSelectFork(MAINNET_RPC_URL);
        assertEq(vm.activeFork(), anotherFork);
    }

    function testCanSetForkBlockNumber() public {
        vm.selectFork(mainnetFork);
        vm.rollFork(1_337_000);
        assertEq(block.number, 1_337_000);
    }
}
```

选择分叉时，只有 msg.sender 和测试合约（ForkTest）的账户是持久的。 但是**任何帐户都可以变成持久帐户**：**makePersistent**，

persistent 帐户是唯一的， 即它存在于所有分叉上

>  vm.makePersistent(address(simple));

# 高级测试

## 模糊测试

vm.assume(amount > 0.1 ether);

**"vm.assume()"**方法通常用于在智能合约中进行条件断言（assertions）。它用于在代码执行过程中检查某个条件是否满足，如果条件不满足，则会引发异常或终止合约执行。

## 差异测试

**ffi** 允许您执行任意 shell 命令并捕获输出。 这是一个模拟示例：

```solidity
import "forge-std/Test.sol";

contract TestContract is Test {

    function testMyFFI () public {
        string[] memory cmds = new string[](2);
        cmds[0] = "cat";
        cmds[1] = "address.txt"; // assume contains abi-encoded address.
        bytes memory result = vm.ffi(cmds);
        address loadedAddress = abi.decode(result, (address));
        // Do something with the address
        // ...
    }
}
```

一个地址之前已经写入了`address.txt`，我们使用 FFI 作弊码读取了它。 现在可以在整个测试合约中使用此数据。
---
title: 13-interface接口

date: 2022-07-17	

categories: Solidity0.8.17	

tags: [区块链,Solidity0.8.17]
---	

很多时候，我们需要调用已经部署在链上的已经合约，这时候可以通过接口合约实现部分调用的逻辑，我们只需要写一个与之对应的接口合约，就可以调用了。

在 solidity 语言中，只要某个合约有和接口种相同的函数声明，就可以被此合约所接受。接口就是起到一个桥接的作用；类似手机的接口，只要匹配，可以进行充电，也可以进行听歌。

interface 类似于[抽象合约](https://www.axihe.com/source/10.inheritance.html#abstract) ，但它们不能实现任何功能。还有其他限制。 `interface`内的函数被隐式标记为virtual

# 限制

- 无法实现任何功能，没有函数体。

- 无法定义构造函数。

- 无法定义状态变量。

- 无法定义结构（`struct`）（`0.5.0` 版本开始接口里可以支持声明 `enum` 类型）。

- 不可以声明修改器。

- 所有声明的函数必须是external的，尽管在合约里可以是 public

  - 文档说：将来可能会解除这里的某些限制。

⚠️ 注意： interface 可以基于别的 interface，可以继承其他合约。比如 `interface IERC20Metadata is IERC20{}`,定义 `IERC20Metadata` 基于 `IERC20` 接口。

# 定义和使用

接口需要有 interface 关键字，并且内部只需要有函数的声明，不用实现。只要某合约中有和词接口相同的函数声明，就可以被此合约所接受。语法如下

```
interface 接口名{
    函数声明;
}
```

接口基本上仅限于合约 ABI 可以表示的内容，并且 ABI 和接口之间的转换不应该丢失任何信息。

在下面的例子中，定义了 cat 合约以及 dog 合约。他们都有 eat 方法.以此他们都可以被上面的 animalEat 接口所接收。

### 使用例子

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Cat {
    uint256 public age;

    function eat() public returns (string memory) {
        age++;
        return "cat eat fish";
    }

    function sleep1() public pure returns (string memory) {
        return "sleep1";
    }
}

contract Dog {
    uint256 public age;

    function eat() public returns (string memory) {
        age += 2;
        return "dog miss you";
    }

    function sleep2() public pure returns (string memory) {
        return "sleep2";
    }
}

interface AnimalEat {
    function eat() external returns (string memory);
}

contract Animal {
    function test(address _addr) external returns (string memory) {
        AnimalEat general = AnimalEat(_addr);
        return general.eat();
    }
}
```

测试流程:

1. 部署 Cat 合约

2. 部署 Dog 合约

3. 部署 Animal 合约

4. 调用Animal.test,参数是 Cat 合约地址

   1. 返回 `"string: cat eat fish"`
   2. 在 Cat 合约内查看 `age` 返回的数字
   
5. 调用Animal.test,参数是 Dog 合约地址

   1. 返回 `"string: dog miss you"`
2. 在 Dog 合约内查看 `age` 返回的数字

在合约 Animal 中，调用函数 test，如果传递的是部署的 Cat 的合约地址，那么我们在调用接口的 eat 方法时，实则调用了 Cat 合约的 eat 方法。 同理，如果传递的是部署的 Dog 的合约地址，那么我们在调用接口的 eat 方法时，实则调用了 dog 合约的 eat 方法。

### 隐式的标记为`virtual`

就像继承其他合约一样，合约可以继承接口。接口中的函数都会隐式的标记为`virtual` ，意味着他们会被重写并不需要 `override` 关键字。
但是不表示重写（overriding）函数可以再次重写，仅仅当重写的函数标记为`virtual` 才可以再次重写。

接口可以继承其他的接口，遵循同样继承规则。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

interface ParentA {
    function test() external returns (uint256);
}

interface ParentB {
    function test() external returns (uint256);
}

interface SubInterface is ParentA, ParentB {
    // 必须重新定义 test 函数，以表示兼容父合约含义
    function test() external override(ParentA, ParentB) returns (uint256);
}
```

# 全局属性 `type(I).interfaceId`

返回接口`I` 的 bytes4 类型的接口 ID，接口 ID 参考： EIP-165 定义的， 接口 ID 被定义为 XOR （异或） 接口内所有的函数的函数选择器（除继承的函数。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

interface ParentA {
    function test() external returns (uint256);
}

contract Demo {
    function interfaceId() public pure returns (bytes4) {
    	//0xf8a8fd6d
        return type(ParentA).interfaceId;
    }
}
```

# ERC20 标准

### 标准

**问题: 如何判断一个 Token 合约是否为标准的 ERC20 合约？**

- 只要含有 ERC20 接口规定的所有内容，就算标准的 ERC20 合约。
  - 至于方法内的逻辑是如何实现的，是不做判断。

详情参考: https://eips.ethereum.org/EIPS/eip-20

### 标准 ERC20 接口

- 3 个查询

  - `balanceOf`: 查询指定地址的 Token 数量
  - `totalSupply`: 查询当前合约的 Token 总量
  - `allowance`: 查询指定地址对另外一个地址的剩余授权额度
- 2 个交易

  - `transfer`: 从当前调用者地址发送指定数量的 Token 到指定地址,这是一个写入方法，所以还会抛出一个 `Transfer` 事件。
  - `transferFrom`: 当向另外一个合约地址存款时，对方合约必须调用 transferFrom 才可以把 Token 拿到它自己的合约中。
- 2 个事件

  - `Transfer`
  - `Approval`
- 1 个授权

  - `approve`: 授权指定地址可以操作调用者的最大 Token 数量。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

interface IERC20 {
    // 1个授权
    function approve(address spender, uint256 amount) external returns (bool);

    // 2个事件
    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 amount
    );

    // 2个交易
    function transfer(address recipient, uint256 amount)
        external
        returns (bool);

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

    // 3个查询
    function totalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);
}
```

### ERC20 标准合约实现

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

interface IERC20 {
    // 1个授权
    function approve(address spender, uint256 amount) external returns (bool);

    // 2个事件
    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 amount
    );

    // 2个交易
    function transfer(address recipient, uint256 amount)
        external
        returns (bool);

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

    // 3个查询
    function totalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);
}

contract ERC20 is IERC20 {
    // 状态变量
    string public name;
    string public symbol;
    uint8 public immutable decimals;

    address public immutable owner;

    // uint256 public immutable totalSupply; // 不增加总量
    uint256 public totalSupply; // 总价总量
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    // 函数修改器
    modifier onlyOwner() {
        require(msg.sender == owner, "not owner");
        _;
    }

    // 构造函数
    constructor(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _totalSupply
    ) {
        owner = msg.sender;
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        totalSupply = _totalSupply;
        balanceOf[msg.sender] = _totalSupply;
        emit Transfer(address(0), msg.sender, _totalSupply);
    }

    // 1个授权
    function approve(address spender, uint256 amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    // 2个交易
    function transfer(address recipient, uint256 amount)
        external
        returns (bool)
    {
        balanceOf[msg.sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(msg.sender, recipient, amount);
        return true;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool) {
        // msg.sender 也就是当前调用者，是被批准者
        allowance[sender][msg.sender] -= amount;
        balanceOf[sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(sender, recipient, amount);
        return true;
    }

    // 1个铸币 - 非必须
    function mint(uint256 amount) external onlyOwner returns (bool) {
        totalSupply += amount;
        balanceOf[msg.sender] += amount;
        emit Transfer(address(0), msg.sender, amount);
        return true;
    }

    // 1个销毁 - 非必须
    function burn(uint256 amount) external returns (bool) {
        totalSupply -= amount;
        balanceOf[msg.sender] -= amount;
        emit Transfer(msg.sender, address(0), amount);
        return true;
    }

    // 转移 owner 权限等其他一些操作均是看各自业务，非必需的
}
```

# ERC721 标准

参考资料:

- https://eips.ethereum.org/EIPS/eip-721
- https://ethereum.org/zh/developers/docs/standards/tokens/erc-721/

### 场景说明

非同质化代币（NFT）用于以唯一的方式标识某人或者某物。 此类型的代币可以被完美地用于出售下列物品的平台：收藏品、密钥、彩票、音乐会座位编号、体育比赛等。 这种类型的代币有着惊人的潜力，因此它需要一个适当的标准。ERC-721 就是为解决这个问题而来！

所有 NFTs 都有一个 `uint256` 变量，名为 `tokenId`，所以对于任何 ERC-721 合约，这对值 `contract address, tokenId` 必须是全局唯一的。 也就是说，去中心化应用程序可以有一个“转换器”， 使用 tokenId 作为输入并输出一些很酷的事物图像，例如僵尸、武器、技能或神奇的小猫咪！

### 合约代码

```solidity
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/**
 * @dev ERC165 标准的接口 https://eips.ethereum.org/EIPS/eip-165
 * https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified
 */
interface IERC165 {
    /// @notice 查询合约是否实现接口
    /// @param interfaceID ERC-165 中指定的接口标识符
    /// @dev 接口标识在 ERC-165 中指定。此功能需要低于 30,000 gas。
    /// @return 如果合约实现了 interfaceID 且 interfaceID 不是 0xffffffff，则为 true，否则为 false
    function supportsInterface(bytes4 interfaceID) external view returns (bool);
}

/// @title ERC-721 Non-Fungible Token Standard
/// @dev See https://eips.ethereum.org/EIPS/eip-721
///  Note: the ERC-165 identifier for this interface is 0x80ac58cd.
interface IERC721 is IERC165 {
    /**
     @dev 当任何 NFT 的所有权通过任何形式发生变化时，需要触发该事件。
     当 NFT 创建（`from` == 0）和销毁（`to` == 0）时会触发此事件。
     例外情况：在合约创建期间，可以创建和分配任意数量的 NFT，而不会发出 Transfer。
     在任何形式的资产转移时，该 NFT如果有批准地址将重置为无。
    */
    event Transfer(
        address indexed _from,
        address indexed _to,
        uint256 indexed _tokenId
    );

    /**
     * 当 NFT 的批准地址被更改或重新确认时，它会发出。
     * 零地址表示没有批准的地址。
     * 当 Transfer 事件发出时，这也表明该 NFT 如果有批准地址被重置为无。
     */
    event Approval(
        address indexed _owner,
        address indexed _approved,
        uint256 indexed _tokenId
    );

    /// @dev 当为所有者启用或禁用操作员时，它会发出。 运营者可以管理所有者的所有 NFT。
    event ApprovalForAll(
        address indexed _owner,
        address indexed _operator,
        bool _approved
    );

    /// @notice 所有者的 NFT 数量
    /// @dev 分配给零地址的 NFT 被认为是无效的，并且该函数抛出有关零地址的查询。
    /// @param _owner 查询余额的地址
    /// @return `_owner` 拥有的 NFT 数量，可能为零
    function balanceOf(address _owner) external view returns (uint256);

    /// @notice 找到 NFT 的所有者
    /// @dev 分配给零地址的 NFT 被认为是无效的，并且对它们的查询确实会抛出异常。
    /// @param _tokenId NFT 的标识符
    /// @return NFT所有者的地址
    function ownerOf(uint256 _tokenId) external view returns (address);

    /// @notice 将 NFT 的所有权从一个地址转移到另一个地址
    /// @dev Throws unless `msg.sender` is the current owner, an authorized
    ///  operator, or the approved address for this NFT. Throws if `_from` is
    ///  not the current owner. Throws if `_to` is the zero address. Throws if
    ///  `_tokenId` is not a valid NFT. When transfer is complete, this function
    ///  checks if `_to` is a smart contract (code size > 0). If so, it calls
    ///  `onERC721Received` on `_to` and throws if the return value is not
    ///  `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`.
    /// @param _from NFT的当前所有者
    /// @param _to 新 owner
    /// @param _tokenId 转移的 NFT
    /// @param data 没有指定格式的附加数据，在调用 _to 时发送
    function safeTransferFrom(
        address _from,
        address _to,
        uint256 _tokenId,
        bytes calldata data
    ) external payable;

    /// @notice 将 NFT 的所有权从一个地址转移到另一个地址
    /// @dev 这与具有额外数据参数的其他函数的工作方式相同，只是此函数只是将数据设置为“”。
    /// @param _from NFT的当前所有者
    /// @param _to 新 owner
    /// @param _tokenId 转移的 NFT
    function safeTransferFrom(
        address _from,
        address _to,
        uint256 _tokenId
    ) external payable;

    /// @notice 转移 NFT 的所有权——调用者有责任确认 `_to` 能够接收 NFT，否则它们可能会永久丢失
    /// @dev Throws unless `msg.sender` is the current owner, an authorized
    ///  operator, or the approved address for this NFT. Throws if `_from` is
    ///  not the current owner. Throws if `_to` is the zero address. Throws if
    ///  `_tokenId` is not a valid NFT.
    /// @param _from NFT的当前所有者
    /// @param _to 新 owner
    /// @param _tokenId 转移的 NFT
    function transferFrom(
        address _from,
        address _to,
        uint256 _tokenId
    ) external payable;

    /// @notice 更改或重申 NFT 的批准地址
    /// @dev The zero address indicates there is no approved address.
    ///  Throws unless `msg.sender` is the current NFT owner, or an authorized
    ///  operator of the current owner.
    /// @param _approved 新批准的 NFT 控制器
    /// @param _tokenId NFT 批准
    function approve(address _approved, uint256 _tokenId) external payable;

    /// @notice 启用或禁用对第三方（“操作员”）的批准以管理所有 `msg.sender` 的资产
    /// @dev 发出 ApprovalForAll 事件。 合同必须允许每个所有者有多个操作员。
    /// @param _operator 添加到授权运营商集中的地址
    /// @param _approved 如果运营商获得批准，则为 True，如果撤消批准，则为 false
    function setApprovalForAll(address _operator, bool _approved) external;

    /// @notice 获取单个 NFT 的认可地址
    /// @dev 如果 _tokenId 不是有效的 NFT，则抛出。
    /// @param _tokenId NFT寻找批准的地址
    /// @return 此 NFT 的批准地址，如果没有则为零地址
    function getApproved(uint256 _tokenId) external view returns (address);

    /// @notice 查询一个地址是否是另一个地址的授权操作员
    /// @param _owner 拥有 NFT 的地址
    /// @param _operator 代表所有者的地址
    /// @return 如果 _operator 是 _owner 的批准运算符，则为真，否则为假
    function isApprovedForAll(address _owner, address _operator)
        external
        view
        returns (bool);
}
```

# ERC1155 标准

参考资料:

- https://eips.ethereum.org/EIPS/eip-1155
- https://ethereum.org/zh/developers/docs/standards/tokens/erc-1155/

### 场景说明

用于多种代币管理的合约标准接口。单个部署的合约可以包括同质化代币、非同质化代币或其他配置（如半同质化代币）的任何组合。

它的目的很单纯，就是创建一个智能合约接口，可以代表和控制任何数量的同质化和非同质化代币类型。 这样一来，ERC-1155 代币就具有与 ERC-20 和 ERC-721 代币相同的功能，甚至可以同时使用这两者的功能。 而最重要的是，它能改善这两种标准的功能，使其更有效率，并纠正 ERC-20 和 ERC-721 标准上明显的实施错误。

### 合约代码

```solidity
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

/**
 * @dev ERC165 标准的接口 https://eips.ethereum.org/EIPS/eip-165
 * https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified
 */
interface IERC165 {
    /// @notice 查询合约是否实现接口
    /// @param interfaceID ERC-165 中指定的接口标识符
    /// @dev 接口标识在 ERC-165 中指定。此功能需要低于 30,000 gas。
    /// @return 如果合约实现了 interfaceID 且 interfaceID 不是 0xffffffff，则为 true，否则为 false
    function supportsInterface(bytes4 interfaceID) external view returns (bool);
}

/**
    @title ERC-1155 Multi Token Standard
    @dev See https://eips.ethereum.org/EIPS/eip-1155
    Note: The ERC-165 identifier for this interface is 0xd9b67a26.
 */
interface IERC1155 is IERC165 {
    /**
        @dev Either `TransferSingle` or `TransferBatch` MUST emit when tokens are transferred,
        including zero value transfers as well as minting or burning (see "Safe Transfer Rules" section of the standard).
        The `_operator` argument MUST be the address of an account/contract that is approved to make the transfer (SHOULD be msg.sender).
        The `_from` argument MUST be the address of the holder whose balance is decreased.
        The `_to` argument MUST be the address of the recipient whose balance is increased.
        The `_id` argument MUST be the token type being transferred.
        The `_value` argument MUST be the number of tokens the holder balance is decreased by and match what the recipient balance is increased by.
        When minting/creating tokens, the `_from` argument MUST be set to `0x0` (i.e. zero address).
        When burning/destroying tokens, the `_to` argument MUST be set to `0x0` (i.e. zero address).
    */
    event TransferSingle(
        address indexed _operator,
        address indexed _from,
        address indexed _to,
        uint256 _id,
        uint256 _value
    );

    /**
        @dev Either `TransferSingle` or `TransferBatch` MUST emit when tokens are transferred,
        including zero value transfers as well as minting or burning (see "Safe Transfer Rules" section of the standard).
        The `_operator` argument MUST be the address of an account/contract that is approved to make the transfer (SHOULD be msg.sender).
        The `_from` argument MUST be the address of the holder whose balance is decreased.
        The `_to` argument MUST be the address of the recipient whose balance is increased.
        The `_ids` argument MUST be the list of tokens being transferred.
        The `_values` argument MUST be the list of number of tokens (matching the list and order of tokens specified in _ids)
        the holder balance is decreased by and match what the recipient balance is increased by.
        When minting/creating tokens, the `_from` argument MUST be set to `0x0` (i.e. zero address).
        When burning/destroying tokens, the `_to` argument MUST be set to `0x0` (i.e. zero address).
    */
    event TransferBatch(
        address indexed _operator,
        address indexed _from,
        address indexed _to,
        uint256[] _ids,
        uint256[] _values
    );

    /**
        @dev 必须在批准第二方/运营商地址管理所有者地址的所有令牌时启用或禁用（没有事件假定禁用）
    */
    event ApprovalForAll(
        address indexed _owner,
        address indexed _operator,
        bool _approved
    );

    /**
        @dev 必须在为令牌 ID 更新 URI 时发出。
        URI 在 RFC 3986 中定义。
        URI 必须指向符合“ERC-1155 元数据 URI JSON 模式”的 JSON 文件。
    */
    event URI(string _value, uint256 indexed _id);

    /**
        @notice Transfers `_value` amount of an `_id` from the `_from` address
                to the `_to` address specified (with safety call).
        @dev Caller must be approved to manage the tokens being transferred
        out of the `_from` account (see "Approval" section of the standard).
        MUST revert if `_to` is the zero address.
        MUST revert if balance of holder for token `_id` is lower than the `_value` sent.
        MUST revert on any other error.
        MUST emit the `TransferSingle` event to reflect the balance change (see "Safe Transfer Rules" section of the standard).
        After the above conditions are met, this function MUST check if `_to` is a smart contract (e.g. code size > 0). If so,
        it MUST call `onERC1155Received` on `_to` and act appropriately (see "Safe Transfer Rules" section of the standard).
        @param _from    Source address
        @param _to      Target address
        @param _id      ID of the token type
        @param _value   Transfer amount
        @param _data    Additional data with no specified format, MUST be sent unaltered in call to `onERC1155Received` on `_to`
    */
    function safeTransferFrom(
        address _from,
        address _to,
        uint256 _id,
        uint256 _value,
        bytes calldata _data
    ) external;

    /**
        @notice 将 `_ids` 的 `_values` 数量从 `_from` 地址转移到指定的 `_to` 地址（使用安全调用）。
        @dev Caller must be approved to manage the tokens being transferred out of the `_from` account (see "Approval" section of the standard).
        MUST revert if `_to` is the zero address.
        MUST revert if length of `_ids` is not the same as length of `_values`.
        MUST revert if any of the balance(s) of the holder(s) for token(s) in `_ids` is lower than the respective amount(s) in `_values` sent to the recipient.
        MUST revert on any other error.
        MUST emit `TransferSingle` or `TransferBatch` event(s) such that all the balance changes are reflected (see "Safe Transfer Rules" section of the standard).
        Balance changes and events MUST follow the ordering of the arrays (_ids[0]/_values[0] before _ids[1]/_values[1], etc).
        After the above conditions for the transfer(s) in the batch are met, this function MUST check if `_to` is a smart contract (e.g. code size > 0). If so,
        it MUST call the relevant `ERC1155TokenReceiver` hook(s) on `_to` and act appropriately (see "Safe Transfer Rules" section of the standard).
        @param _from    Source address
        @param _to      Target address
        @param _ids     每个令牌类型的 ID（顺序和长度必须匹配 _values 数组）
        @param _values  每种代币类型的转账金额（顺序和长度必须匹配 _ids 数组）
        @param _data    没有指定格式的额外数据，必须在调用 _to 上的 `ERC1155TokenReceiver` 钩子时原封不动地发送
    */
    function safeBatchTransferFrom(
        address _from,
        address _to,
        uint256[] calldata _ids,
        uint256[] calldata _values,
        bytes calldata _data
    ) external;

    /**
        @notice 获取帐户令牌的余额。
        @param _owner  令牌持有者的地址
        @param _id     ID of the token
        @return        请求的代币类型的所有者余额
     */
    function balanceOf(address _owner, uint256 _id)
        external
        view
        returns (uint256);

    /**
        @notice 获取多个账户/代币对的余额
        @param _owners 代币持有者的地址
        @param _ids    ID of the tokens
        @return        请求的令牌类型的 _owner 余额（即每个 (owner, id) 对的余额）
     */
    function balanceOfBatch(address[] calldata _owners, uint256[] calldata _ids)
        external
        view
        returns (uint256[] memory);

    /**
        @notice 启用或禁用对第三方（“操作员”）的批准以管理所有调用者的令牌。
        @dev 必须在成功时发出 ApprovalForAll 事件。
        @param _operator  添加到授权运营商集中的地址
        @param _approved  如果运营商获得批准，则为 True，如果撤消批准，则为 false
    */
    function setApprovalForAll(address _operator, bool _approved) external;

    /**
        @notice 查询给定所有者的操作员的批准状态。
        @param _owner     The owner of the tokens
        @param _operator  授权操作员的地址
        @return           如果操作员被批准则为真，否则为假
    */
    function isApprovedForAll(address _owner, address _operator)
        external
        view
        returns (bool);
}
```

# ERC3525 标准

每个符合 EIP-3525 的合约都必须实现 EIP-3525、EIP-721 和 EIP-165 接口

参考资料:

- https://eips.ethereum.org/EIPS/eip-3525

### 场景说明

描述一组具有相同类型，但是有轻微不同的东西。比如相同的 100 元人民币，一共 100 张，每一张都是价值 100 的纸币，大部分的防伪等等都不同，但是每一张都编号都不同。

### 合约代码

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title EIP-3525 Semi-Fungible Token Standard
 * Note: the EIP-165 identifier for this interface is 0xd5358140.
 */

interface IERC3525  /* is IERC165, IERC721 */ {
    /**
     * @dev MUST emit when value of a token is transferred to another token with the same slot,
     *  including zero value transfers (_value == 0) as well as transfers when tokens are created
     *  (`_fromTokenId` == 0) or destroyed (`_toTokenId` == 0).
     * @param _fromTokenId The token id to transfer value from
     * @param _toTokenId The token id to transfer value to
     * @param _value The transferred value
     */
    event TransferValue(
        uint256 indexed _fromTokenId,
        uint256 indexed _toTokenId,
        uint256 _value
    );

    /**
     * @dev MUST emit when the approval value of a token is set or changed.
     * @param _tokenId The token to approve
     * @param _operator The operator to approve for
     * @param _value The maximum value that `_operator` is allowed to manage
     */
    event ApprovalValue(
        uint256 indexed _tokenId,
        address indexed _operator,
        uint256 _value
    );

    /**
     * @dev MUST emit when the slot of a token is set or changed.
     * @param _tokenId The token of which slot is set or changed
     * @param _oldSlot The previous slot of the token
     * @param _newSlot The updated slot of the token
     */
    event SlotChanged(
        uint256 indexed _tokenId,
        uint256 indexed _oldSlot,
        uint256 indexed _newSlot
    );

    /**
     * @notice Get the number of decimals the token uses for value - e.g. 6, means the user
     *  representation of the value of a token can be calculated by dividing it by 1,000,000.
     *  Considering the compatibility with third-party wallets, this function is defined as
     *  `valueDecimals()` instead of `decimals()` to avoid conflict with EIP-20 tokens.
     * @return The number of decimals for value
     */
    function valueDecimals() external view returns (uint8);

    /**
     * @notice Get the value of a token.
     * @param _tokenId The token for which to query the balance
     * @return The value of `_tokenId`
     */
    function balanceOf(uint256 _tokenId) external view returns (uint256);

    /**
     * @notice Get the slot of a token.
     * @param _tokenId The identifier for a token
     * @return The slot of the token
     */
    function slotOf(uint256 _tokenId) external view returns (uint256);

    /**
     * @notice Allow an operator to manage the value of a token, up to the `_value`.
     * @dev MUST revert unless caller is the current owner, an authorized operator, or the approved
     *  address for `_tokenId`.
     *  MUST emit the ApprovalValue event.
     * @param _tokenId The token to approve
     * @param _operator The operator to be approved
     * @param _value The maximum value of `_toTokenId` that `_operator` is allowed to manage
     */
    function approve(
        uint256 _tokenId,
        address _operator,
        uint256 _value
    ) external payable;

    /**
     * @notice Get the maximum value of a token that an operator is allowed to manage.
     * @param _tokenId The token for which to query the allowance
     * @param _operator The address of an operator
     * @return The current approval value of `_tokenId` that `_operator` is allowed to manage
     */
    function allowance(uint256 _tokenId, address _operator)
        external
        view
        returns (uint256);

    /**
     * @notice Transfer value from a specified token to another specified token with the same slot.
     * @dev Caller MUST be the current owner, an authorized operator or an operator who has been
     *  approved the whole `_fromTokenId` or part of it.
     *  MUST revert if `_fromTokenId` or `_toTokenId` is zero token id or does not exist.
     *  MUST revert if slots of `_fromTokenId` and `_toTokenId` do not match.
     *  MUST revert if `_value` exceeds the balance of `_fromTokenId` or its allowance to the
     *  operator.
     *  MUST emit `TransferValue` event.
     * @param _fromTokenId The token to transfer value from
     * @param _toTokenId The token to transfer value to
     * @param _value The transferred value
     */
    function transferFrom(
        uint256 _fromTokenId,
        uint256 _toTokenId,
        uint256 _value
    ) external payable;

    /**
     * @notice 转移将指定数量的代币到新地址。调用者应确认 _to 能够接收 EIP-3525 资产。
     * @dev This function MUST create a new EIP-3525 token with the same slot for `_to`,
     *  or find an existing token with the same slot owned by `_to`, to receive the transferred value.
     *  MUST revert if `_fromTokenId` is zero token id or does not exist.
     *  MUST revert if `_to` is zero address.
     *  MUST revert if `_value` exceeds the balance of `_fromTokenId` or its allowance to the
     *  operator.
     *  MUST emit `Transfer` and `TransferValue` events.
     * @param _fromTokenId The token to transfer value from
     * @param _to The address to transfer value to
     * @param _value The transferred value
     * @return ID of the token which receives the transferred value
     */
    function transferFrom(
        uint256 _fromTokenId,
        address _to,
        uint256 _value
    ) external payable returns (uint256);
}
```

### 扩展

更多关于 3525 协议的内容，参考 https://cloud.tencent.com/developer/article/2155201




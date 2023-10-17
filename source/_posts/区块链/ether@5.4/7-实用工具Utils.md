---
title: 7-实用工具Utils

date: 2022-08-24	

categories: ether@5.4.js	

tags: [区块链,ether@5.4.js]
---	

# 应用程序二进制接口

## 1.1 AbiCoder

**AbiCoder**是编码器的集合，可用于在EVM和更高级别库之间通过二进制数据格式进行编码和解码的操作。

大多数开发人员不需要直接使用这个类，因为Interface类极大地简化了这些操作。

## 1.2 ABI格式

### 1.2.1 HUman-REadable ABI

**Human-Readable ABI** 是早期被 ethers 所描述提出的， 它允许使用一个Solidity式签名(Solidity signature)来描述每个方法(method)、事件(event)和错误(error)。

### 1.2.2 **Solidity JSON **ABI

**Solidity JSON ABI**是许多工具导出的标准格式，包括Solidity编译器.

### 1.2.3 **Solidity Object** ABI

使用JSON.parse解析Solidity JSON ABI的结果完全与Interface类兼容， 且该对象的每个方法、事件和错误都与Fragment类兼容。

一些开发人员可能更喜欢这种方式，因为它允许将ABI属性作为普通的JavaScript对象访问，并且Solidity ABI与JSON ABI非常匹配。

## 1.3 Fragments

对一个ABI进行描述解释

### 1.3.0 Formats

每个Fragment和ParamType可以使用其`format`方法输出。

***ethers*.*utils*.*FormatTypes*.full ⇒ *string*：**这是一个完整的人类可读(human-readable)的字符串，包括所有的参数名，任何可选的修饰符	(例如，`indexed`，`public`等)和空格，以提高代码可读性。

***ethers*.*utils*.*FormatTypes*.minimal ⇒ *string*：**这类似于`full`，除了没有不必要的空白或参数名。这对于存储最小的字符串非常有用，	该字符串仍然可以使用Fragment . from， 从完全重构原始的Fragment。

***ethers*.*utils*.*FormatTypes*.json ⇒ *string*：**这将返回一个JavaScript对象，安全地调用`JSON.stringify`创建一个JSON字符串。

***ethers*.*utils*.*FormatTypes*.sighash ⇒ *string*：**一个最小的输出格式，Solidity在计算签名哈希或事件主题topic的哈希时使用它。

### 1.3.1 **Fragment**

一个ABI是**Fragments**的集合，每个fragment指定:一个错误、一个事件、一个函数、一个构造函数

**属性：**

***fragment*.name ⇒ *string*：**事件或函数的name。如果是ConstructorFragment时为null。

***fragment*.type ⇒ *string：***这是一个表示Fragment类型的字符串。如:`constructor`、`event`、`function`

***fragment*.inputs ⇒ *Array< ParamType >*：**构造函数、事件等输入参数的ParamType的数组。

**方法：**

***fragment*.format( [ *format* = *sighash* ] ) ⇒ *string*：**使用可用的输出格式创建Fragment的字符串描述。

***ethers*.*utils*.______Fragment.from( *objectOrString* ) ⇒ ____*Fragment*：**从任何兼容的object或String创建一个新的**__Fragment**子类。

***ethers*.*utils*.*Fragment*.is__Fragment( *object* ) ⇒ *boolean*：**如果*object* 是一个**__Fragment**返回true。

### 1.3.2 ConstructorFragment

***fragment*.gas ⇒ *大数(BigNumber)*：**表示部署期间应该使用的gas limit，它可以是null。

***fragment*.payable ⇒ *boolean*：**表示构造函数在部署期间是否可以接收ether(例如：msg.value != 0)。

***fragment*.stateMutability ⇒ *string*：**构造函数的state mutability。它可以是:`nonpayable`、`payable`

### 1.3.3 ErrorFragment

### 1.3.4 EventFragment

***fragment*.anonymous ⇒ *boolean*：**表示事件(event)是否匿名。匿名事件在创建日志时不会将其topic哈希值注入到topic0中

### 1.3.5 FunctionFragment

***fragment*.constant ⇒ *boolean*：**表示函数是否为常量(即不改变状态)。如果设为true表示状态可变性是`pure` 或 `view`。

***fragment*.stateMutability ⇒ *string*：**构造器的状态可变性。它可以是:`nonpayable`、`payable`、`pure`、`view`

***fragment*.outputs ⇒ *Array< ParamType >*：**函数输出参数的列表。

### 1.3.6 ParamType

表示Solidity中的参数:

**属性：**

***paramType*.name ⇒ *string*：**本地参数名。对于未命名的参数，这个值为null。例如，参数字符串`string foobar`会输出`foobar`。

***paramType*.type ⇒ *string：***参数的完整类型，包括元组和数组符号。对于未命名的参数，这个值可能为null。

***paramType*.baseType ⇒ *string：***参数的基类型(base type)。对于原始类型(例如`address`, `uint256`等)，这等同于type。 对于数组，	它将是string `array`，对于元组，它将是string `tuple`。

***paramType*.indexed ⇒ *boolean：***参数是否被标记为索引。这只适用于参数是EventFragment的部分。

***paramType*.arrayChildren ⇒ *ParamType：***数组的children类型。对于任何非数组的参数，这将是null。

***paramType*.arrayLength ⇒ *number：***数组的长度，或动态数组的长度`-1`。对于不是数组的参数，这个值为null。

***paramType*.components ⇒ *Array< ParamType >*：**元组的组成部分。这对于非元组参数是null。

**方法：**

***paramType*.format( [ *outputType* = *sighash* ] )：**使用可用的output formats创建Fragment的字符串描述。

***ethers*.*utils*.*ParamType*.from( *objectOrString* ) ⇒ *ParamType*：**从任何兼容的*object或String*创建一个新的**ParamType**。

***ethers*.*utils*.*ParamType*.isParamType( *object* ) ⇒ *boolean*：**如果*object* 是一个**ParamType**返回true。

## 1.4 Interface

**接口(Interface)**类是以太坊网络上的与合约交互所需的编码和解码的一种抽象,EVM本身并不理解ABI是什么。

它只是一组商定的格式，用于编码合约所需的各种类型的数据，以便它们可以相互交互。

**（创建实例)new *ethers*.*utils*.Interface( *abi* ):**从表示*abi*的JSON字符串或对象创建一个新的**接口**。

### 1.4.1 属性

***interface*.fragments ⇒ *Array< Fragment >*：**接口中所有的Fragments。

***interface*.errors ⇒ *Array< ErrorFragment >：***接口中所有的Error Fragments。

***interface*.events ⇒ *Array< EventFragment >：***接口中所有的Event Fragments。

***interface*.functions ⇒ *Array< FunctionFragment >：***接口中所有的Function Fragments。

***interface*.deploy ⇒ *ConstructorFragment“：***接口中所有的Constructor Fragments。

### 1.4.2 格式化

***interface*.format( [ *format* ] ) ⇒ *string | Array< string >*：**返回格式化的接口。如果格式类型是`json`，则返回一个字符串，否则返回	一个由human-readable（人类可读）字符串组成的数组。

```js
const FormatTypes = ethers.utils.FormatTypes;

iface.format(FormatTypes.json)
// '[{"type":"constructor","payable":false,"inputs":[{"type":"string","name":"symbol"},{"type":"string","name":"name"}]},{"type":"function","name":"transferFrom","constant":false,"payable":false,"inputs":[{"type":"address","name":"from"},{"type":"address","name":"to"},{"type":"uint256","name":"amount"}],"outputs":[]},{"type":"function","name":"mint","constant":false,"stateMutability":"payable","payable":true,"inputs":[{"type":"uint256","name":"amount"}],"outputs":[]},{"type":"function","name":"balanceOf","constant":true,"stateMutability":"view","payable":false,"inputs":[{"type":"address","name":"owner"}],"outputs":[{"type":"uint256"}]},{"type":"event","anonymous":false,"name":"Transfer","inputs":[{"type":"address","name":"from","indexed":true},{"type":"address","name":"to","indexed":true},{"type":"uint256","name":"amount"}]},{"type":"error","name":"AccountLocked","inputs":[{"type":"address","name":"owner"},{"type":"uint256","name":"balance"}]},{"type":"function","name":"addUser","constant":false,"payable":false,"inputs":[{"type":"tuple","name":"user","components":[{"type":"string","name":"name"},{"type":"address","name":"addr"}]}],"outputs":[{"type":"uint256","name":"id"}]},{"type":"function","name":"addUsers","constant":false,"payable":false,"inputs":[{"type":"tuple[]","name":"user","components":[{"type":"string","name":"name"},{"type":"address","name":"addr"}]}],"outputs":[{"type":"uint256[]","name":"id"}]},{"type":"function","name":"getUser","constant":true,"stateMutability":"view","payable":false,"inputs":[{"type":"uint256","name":"id"}],"outputs":[{"type":"tuple","name":"user","components":[{"type":"string","name":"name"},{"type":"address","name":"addr"}]}]}]'

iface.format(FormatTypes.full)
// [
//   'constructor(string symbol, string name)',
//   'function transferFrom(address from, address to, uint256 amount)',
//   'function mint(uint256 amount) payable',
//   'function balanceOf(address owner) view returns (uint256)',
//   'event Transfer(address indexed from, address indexed to, uint256 amount)',
//   'error AccountLocked(address owner, uint256 balance)',
//   'function addUser(tuple(string name, address addr) user) returns (uint256 id)',
//   'function addUsers(tuple(string name, address addr)[] user) returns (uint256[] id)',
//   'function getUser(uint256 id) view returns (tuple(string name, address addr) user)'
// ]

iface.format(FormatTypes.minimal)
// [
//   'constructor(string,string)',
//   'function transferFrom(address,address,uint256)',
//   'function mint(uint256) payable',
//   'function balanceOf(address) view returns (uint256)',
//   'event Transfer(address indexed,address indexed,uint256)',
//   'error AccountLocked(address,uint256)',
//   'function addUser(tuple(string,address)) returns (uint256)',
//   'function addUsers(tuple(string,address)[]) returns (uint256[])',
//   'function getUser(uint256) view returns (tuple(string,address))'
// ]
```

### 1.4.3 **Fragment Access**

***interface*.getFunction( *fragment* ) ⇒ FunctionFragment：**返回*fragment*的FunctionFragment

***interface*.getError( *fragment* ) ⇒ ErrorFragment：**返回*fragment*的ErrorFragment

***interface*.getEvent( *fragment* ) ⇒ EventFragment：**返回*fragment*的EventFragment

```js
// 通过方法(method)的签名，这是经过标准化后的，因此空格和多余的属性会被舍去

iface.getFunction("transferFrom(address, address, uint256)");

// 通过name的方式;这只在方法名称是唯一确定的情况下才有效
iface.getFunction("transferFrom");

// 通过函数选择器
iface.getFunction("0x23b872dd");

// 如果方法不存在将抛出异常
iface.getFunction("doesNotExist()");
// [Error: no matching function] {
//   argument: 'signature',
//   code: 'INVALID_ARGUMENT',
//   reason: 'no matching function',
//   value: 'doesNotExist()'
// }
```

### 1.4.4 签名和主题的哈希

***interface*.getSighash( *fragment* ) ⇒ *string< DataHexString< 4 > >*：**返回*fragment*的签名哈希(sighash)或函数选择器(Function Selector)

```js
iface.getSighash("balanceOf");
iface.getSighash("balanceOf(address)");
const fragment = iface.getFunction("balanceOf")
iface.getSighash(fragment);
// '0x70a08231'
```

***interface*.getEventTopic( *fragment* ) ⇒ string< DataHexString< 32 > >：**返回*fragment*的主题哈希(topic hash)

```js
iface.getEventTopic("Transfer");
iface.getEventTopic("Transfer(address, address, uint)");
const fragment = iface.getEvent("Transfer")
iface.getEventTopic(fragment);
// '0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef'
```

### 1.4.5 编码数据

### 1.4.6 解码数据

### 1.4.7 解析

# 地址

***ethers*.*utils*.getAddress( *address* ) ⇒ *string< 地址(Address) >*：**返回一个校验和*地址*。

***ethers*.*utils*.getIcapAddress( *address* ) ⇒ *string< IcapAddress >*：**返回一个ICAP address地址，与getAddress具有相同的限制条件。

***ethers*.*utils*.isAddress( *address* ) ⇒ boolean：**如果地址有效(任何支持的格式)则返回true。

***ethers*.*utils*.computeAddress( *publicOrPrivateKey* ) ⇒ *string< 地址(Address) >*：**返回*publicOrPrivateKey*的地址。公钥可以压缩或不压缩，私钥将自动转换为派生的公钥。

***ethers*.*utils*.recoverAddress( *digest* , *signature* ) ⇒ *string< 地址(Address) >*：**使用ECDSA Public Key Recovery来确定摘要(digest)生成签名的公钥地址，

***ethers*.*utils*.getContractAddress( *transaction* ) ⇒ *string< 地址(Address) >*：**一个交易用于部署合约则返回部署合约后的合约地址。

***ethers*.*utils*.getCreate2Address( *from* , *salt* , *initCodeHash* ) ⇒ *string< 地址(Address) >*：**返回给定CREATE2调用后的合约地址。

# 大数(BigNumber)

**BigNumber**是一个可以安全地对任意大小的数字进行数学运算的对象。

大多数需要返回值的操作将返回一个**BigNumber**，接受值的参数通常会接收它们。

## 3.1 创建实例

BigNumber的构造函数不能被直接调用。相反，使用静态`BigNumber.from`。

***ethers*.*BigNumber*.from( *aBigNumberish* ) ⇒ *大数(BigNumber)*：**为*aBigNumberish*返回一个**BigNumber**的实例

## 3.2 方法

### 3.2.1 数学运算

***BigNumber*.add( *otherValue* ) ⇒ *大数(BigNumber)*：**返回值为*BigNumber* **+** *otherValue*的BigNumber。

***BigNumber*.sub( *otherValue* ) ⇒ *大数(BigNumber)*：**返回值为 *BigNumber* **-** *otherValue*的BigNumber。

***BigNumber*.mul( *otherValue* ) ⇒ *大数(BigNumber)*：**返回值为 *BigNumber* **×** *otherValue*的BigNumber。

***BigNumber*.div( *divisor* ) ⇒ *大数(BigNumber)：***返回值为 *BigNumber* **÷** *divisor*的BigNumber。

***BigNumber*.mod( *divisor* ) ⇒ *大数(BigNumber)：***返回值为 *BigNumber* ÷ *divisor*余数的BigNumber。

***BigNumber*.pow( *exponent* ) ⇒ *大数(BigNumber)：***返回值为 *BigNumber* 指数的幂为*exponent*的BigNumber。

***BigNumber*.abs( ) ⇒ *大数(BigNumber):***返回值为绝对值的BigNumber。

***BigNumber*.mask( *bitcount* ) ⇒ *大数(BigNumber):***返回一个BigNumber，其BigNumber的值超出*bitcount*最低有效位的位则设为0。

#### Two's Complement

Two's Complement是用于编码和解码固定宽度的有符号值，同时有效地保留数学运算。大多数用户不需要与它们交互。

***BigNumber*.fromTwos( *bitwidth* ) ⇒ *大数(BigNumber)*:**返回一个BigNumber, 值由带位宽（bitwidth）的二进制补码转换而来。

***BigNumber*.toTwos( *bitwidth* ) ⇒ *大数(BigNumber)*：**返回一个BigNumber, BigNumber的值转换为带位宽的二进制补码。

### 3.2.2 比较和相等

***BigNumber*.eq( *otherValue* ) ⇒ *boolean*：**当且仅当*BigNumber*的值等于*otherValue*时返回true。

***BigNumber*.lt( *otherValue* ) ⇒ *boolean：***当且仅当*BigNumber*的值**<**otherValue时返回true。

***BigNumber*.lte( *otherValue* ) ⇒ *boolean：***当且仅当*BigNumber*的值**≤**otherValue时返回true。

***BigNumber*.gt( *otherValue* ) ⇒ *boolean*：**当且仅当*BigNumber*的值**>***otherValue*时返回true。

***BigNumber*.gte( *otherValue* ) ⇒ *boolean：***当且仅当*BigNumber*的值**≥***otherValue*时返回true。

***BigNumber*.isZero( ) ⇒ *boolean：***当且仅当*BigNumber*的值为0时返回true。

### 3.2.3 转换

***BigNumber*.toBigInt( ) ⇒ *bigint*：**在支持BigInt的平台上以JavScript BigInt值返回BigNumber的值。

***BigNumber*.toNumber( ) ⇒ *number*：**将BigNumber的值转换JavaScript值。

> 如果该值大于*Number.MAX_SAFE_INTEGER*或小于等于*Number.MIN_SAFE_INTEGER*， 则会**抛出一个错误**。

***BigNumber*.toString( ) ⇒ *string*：**以十进制字符串的形式返回BigNumber的值。

***BigNumber*.toHexString( ) ⇒ *string< DataHexString >*：**返回BigNumber的值为十六进制的值，`0x`是前缀DataHexString.。

### 3.2.4 检查

***ethers*.*BigNumber*.isBigNumber( *object* ) ⇒ *boolean*：**当且仅当对象是BigNumber是*对象*时返回true。

# 字节处理

## 类型

**Bytes：Bytes**是一个Array或TypedArray对象， 每个值都在有效字节范围内(例如0到255之间)，或者是一个具有`length`属性的对象，每	个索引的属性都在有效字节范围内。

**BytesLike：BytesLike**可以是Bytes或DataHexString。

**DataHexString：DataHexstring**与HexString是相同的，除了它有偶数个nibbles，因此二进制数据作为字符串是有效的。

**HexString：Hexstring**是一个字符串，有一个0x前缀，后面跟着nibbles number类型(例如，不区分大小写的十六进制字符`0-9`和`a-f`)

**Signature：**

- **r** and **s** --- The x co-ordinate of **r** and the **s** value of the signature
- **v** --- The parity of the y co-ordinate of **r**
- **_vs** --- The compact representation of the **s** and **v**
- **recoveryParam** --- The normalized (i.e. 0 or 1) value of **v**

**Raw Signature：原始签名(Raw Signature)**是一种常见的签名格式，r, s和v被连接成一个65字节(130 nibble)的DataHexString。

**SignatureLike：SignatureLike**类似于一个Signature，除了多余的属性可以被省略或者它也可以是一个 Raw Signature。

​	例如，如果指定了**_vs**，则**s**和**v**可以省略。同样，如果提供了**recoveryParam**， 则可以省略**v**(在这种情况下，可以计算出缺失的值)。

## 4.1 检查

***ethers*.*utils*.isBytes( *object* ) ⇒ *boolean*：**当且仅当*object*为有效Bytes时返回true。

***ethers*.*utils*.isBytesLike( *object* ) ⇒ *boolean*：**当且仅当*object*是Bytes或DataHexString时返回true。

***ethers*.*utils*.isHexString( *object* , [ *length* ] ) ⇒ *boolean*：**当且仅当*object*是一个有效的十六进制字符串时返回true。 如果指定了*length*并且*object* 不是一个有效的长度字节的DataHexString，则抛出一个InvalidArgument错误。

## 4.2 数组和十六进制字符串之间的转换

***ethers*.*utils*.arrayify( *DataHexStringOrArrayish* [ , *options* ] ) ⇒ *Uint8Array*：**将*DataHexStringOrArrayish*转换为Uint8Array。

***ethers*.*utils*.hexlify( *hexstringOrArrayish* ) ⇒ *string< DataHexString >*：**将*hexstringOrArrayish*转换为DataHexString。

***ethers*.*utils*.hexValue( *aBigNumberish* ) ⇒ *string< HexString >*：**将*aBigNumberish*转换为HexString，没有前导零。

## 4.3 数组处理

***ethers*.*utils*.concat( *arrayOfBytesLike* ) ⇒ *Uint8Array*：**将*arrayOfBytesLike*中的所有[BytesLike](https://learnblockchain.cn/ethers_v5/api/utils/bytes/#BytesLike)连接到一个单独的Uint8Array中。

***ethers*.*utils*.stripZeros( *aBytesLike* ) ⇒ *Uint8Array*：**返回一个无前导零*aBtyesLike*字节的Uint8Array

***ethers*.*utils*.zeroPad( *aBytesLike* , *length* ) ⇒ *Uint8Array*：**

返回一个以aBytesLike为单位的Uint8Array，前面有0字节的length bytes long。

如果*aBytesLike*的长度已经超过长度字节，则会抛出InvalidArgument错误。

## 4.4 十六进制字符串处理

***ethers*.*utils*.hexConcat( *arrayOfBytesLike* ) ⇒ *string< DataHexString >*：**将*arrayOfBytesLike*中的所有BytesLike连接成一个单一的DataHexString。

***ethers*.*utils*.hexDataLength( *aBytesLike* ) ⇒ *string< DataHexString >*：**返回*aBytesLike*的长度(以字节为单位)。

***ethers*.*utils*.hexDataSlice( *aBytesLike* , *offset* [ , *endOffset* ] ) ⇒ *string< DataHexString >***：返回一个*aBytesLike*切片的DataHexString表示，从*offset*(以字节为单位)到*endOffset*(以字节为单位)。 如果省略了*endOffset*，则使用*aBytesLike*的长度。

***ethers*.*utils*.hexStripZeros( *aBytesLike* ) ⇒ *string< HexString >*：**返回*aBytesLike*的HexString表示形式，去掉所有前导零。

***ethers*.*utils*.hexZeroPad( *aBytesLike* , *length* ) ⇒ *string< DataHexString >*：**返回被填充到*length*字节的*aBytesLike*的DataHexString表示。如果*aBytesLike*的长度已经超过*length*字节，则会抛出InvalidArgument错误。

## 4.5 签名转换

***ethers*.*utils*.joinSignature( *aSignatureLike* ) ⇒ *string< RawSignature >*：**返回*aSignaturelike*的原始格式，它是65字节(130个nibbles)长，连接签名的**r**, **s**和(标准化后))**v**。

***ethers*.*utils*.splitSignature( *aSignatureLikeOrBytesLike* ) ⇒ *Signature*：**返回*aSignaturelike*的完整扩展格式或原始格式DataHexString，将自动计算所有缺失的属性。

## 4.6 随机字节

***ethers*.*utils*.randomBytes( *length* ) ⇒ *Uint8Array*：**返回一个新的Uint8Array 长度为*length*的随机字节。

***ethers*.*utils*.shuffled( *array* ) ⇒ *Array< any >*：**返回一个使用Fisher-Yates Shuffle打乱后的数组副本。

# 以太币格式化与转换

**utils  . etherSymbol** 

以太坊符号(希腊字母  *Xi*  ) 

**utils  . parseEther ( etherString )   =>   BigNumber** 

将代表 ether 单位数的  *etherString*  解析为 wei 单位数的 BitNumber 实例。 （译者注：ether、gwei、wei 等是以太坊的货币单位名称，下同。） 

**utils  . formatEther ( wei )   =>   string** 

将代表 wei 单位数的  *wei*  格式化为代表 ether 单位数的十进制字符串。 输出值总是包含至少一个整数和一个小数位，否则将剪除前导和尾随的 0。 

**utils  . parseUnits ( valueString , decimalsOrUnitName )   =>   BigNumber** 

将代表某单位数的  *valueString*  解析为一个代表 wei 单位数的 BigNumber 实例。 参数  *decimalsOrUnitsName*  可以是 3 到 18 之间（3 的倍数）的小数位数， 或者是以太币单位名称，如： gwei 。 

**utils  . formatUnits ( wei , decimalsOrUnitName )   =>   string** 

将  *wei*  单位数格式化为一个代表某单位数的十进制字符串。 输出值总是包含至少一个整数和一个小数位，否则将修剪前导和尾随的 0。 参数  *decimalsOrUnitsName*  可以是 3 到 18 之间（3 的倍数）的小数位数， 或者是单位名称，如： gwei 。 

**utils  . commify ( numberOrString )   =>   string** 

返回含有千分符的*numberOrString*。如果  *numberOrString*  包含小数点， 则输出值将至少具有一位整数和小数。如果不包含小数点，则输出值不会包含小数。 
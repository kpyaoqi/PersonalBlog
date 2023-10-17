---
title: 6-Web3.(_.net、bzz、shh、utils)

date: 2022-09-09	

categories: web3.js	

tags: [区块链,web3.js]
---	

# web3.*.net(获取网络属性)

## 1.1 getId

获取当前的网络 ID.

**返回值(Promise):**  Number,网络 ID.

## 1.2 isListening

查看当前节点是否正在连接其它对等节点。

**返回值(Promise):**  Boolean

## 1.3 getPeerCount

查看当前节点是否正在连接其它对等节点。

**返回值(Promise):**  Number

# bzz(与swarm交互)

# shh(与 whisper 协议的交互进行消息广播)

# utils(工具性函数)

## 4.1 randomHex

来根据指定字节大小生成密码学强度的伪随机 16 进制字符串.

**参数:**`size` - `Number`: 16 进制字符串的字节大小, 比如 `32` 生成32 个字节大小的 16 进制字符串，以 “0x” 为前缀的 64 个字符来表示。

**返回值:(String): **生成的 16 进制随机字符串.**

## 4.2 BN

为了在不同类型之间进行安全的类型转换, 包括 BigNumber.js 可以使用 utils.toBN

**参数:**`mixed` - `String|Number`: 要转换为 BN 对象的一个数字字符串或 16 进制字符串

**返回值(Object):** BN.js实例

```js
var BN = web3.utils.BN;

new BN(1234).toString();
> "1234"

new BN('1234').add(new BN('1')).toString();
> "1235"

new BN('0xea').toString();
> "234"
```

## 4.3 isBN

检测给定值是否为 BN.js 实例

**参数:**`bn` - `Object`: BN.js 实例.

**返回值(Boolean)**

## 4.4 isBigNumber

检测给定值是否为 BigNumber.js 实例

**参数:**`bignumber ` - `Object`: BigNumber.js 实例.

**返回值(Boolean)**

## 4.5 sha3

将计算输入参数的 sha3 值

> web3.utils.keccak256(string) // 别名

**参数:**`string` - `String`: 要进行哈希运算的字符串.

**返回值(String)**：计算所得哈希值.

> web3.utils.sha3Raw(string)
>
> 将计算输入字符串的 sha3 哈希值，如果传入的是空字符串，将返回空字符串的哈希值而不是 `null`

## 4.6 soliditySha3

使用和 solidity 同样的方式对输入参数进行 sha3 哈希运算。这意味着对这些参数在进行哈希运算之前先进行 ABI 转换和紧凑打包编码。

**参数:**`paramX` - `Mixed`: 任意类型，或者具有 `{type: 'uint', value: '123456'}` 或 `{t: 'bytes', v: '0xfff456'}` 的对象. 基本类型将按下面的规则进行自动检测:

- `String` 非数字 UTF-8 字符串会被解析为 `string`.
- `String|Number|BN|HEX` 正数会被解析为 `uint256`.
- `String|Number|BN` 负数会被解析为 `int256`.
- `Boolean` 作为 `bool`.
- `String` 以 `0x` 开头的 16 进制字符串会被解析为 `bytes`.
- `HEX` 16 进制表示的数字会被解析为 `uint256`.

**返回值(String)：**哈希值

> web3.utils.soliditySha3Raw(param1 [, param2, ...])
>
>  和 `soliditySha3` 不同的是，如果传入的是空字符串，将返回空字符串的哈希值而不是 `null`

## 4.7 isHex

判断给定的字符串是否是个 16 进制字符串

**参数:**`hex` - `String|HEX`: 给定的 16 进制字符串.

**返回值(Boolean)**

> web3.utils.isHexStrict(hex)
>
>  和 `web3.utils.isHex()` 不同的是它要求 16 进制字符串必须以 `0x` 开头.

## 4.8 isAddress

判断给定的地址是否是一个有效的以太坊地址。大小写混合的地址还会检测校验和。

**参数:**`address` - `String`: 地址字符串.

**返回值(Boolean)**

## 4.9 toChecksumAddress

将一个只有大写或小写字符的以太坊地址转换为一个校验和地址

**参数:**`address` - `String`: 地址字符串.

**返回值(String)：**校验和地址

## 4.10 checkAddressChecksum

检测给定地址的校验和. 非校验和地址同样会返回 false.

**参数:**`address` - `String`: 地址字符串.

**返回值(Boolean)：**地址的校验和有效时返回 `true` , 地址不是一个校验和地址或者校验和无效时返回 `false`.

## 4.11 toHex

自动将给定值转换为 16 进制。Number 字符串会被解析为数字。Text 字符串会被解析为 UTF8 字符串

**参数:**`mixed` - `String|Number|BN|BigNumber`: 要进行 16 进制转化的输入.

**返回值(String)：**得到的 16 进制字符串.

> web3.utils.stripHexPrefix(string)
>
> 返回`String`:没有0x前缀的输入字符串。

## 4.12 toBN

将任何给定值安全转换为 (BigNumber.js 实例) 一个 BN.js 实例, 以便于在 JavaScript 中处理大数.

> 如果只是 BN.js 类，可以直接使用 web3.utils.BN

**参数:**`number` - `String|Number|HEX`: 要转换为大数类型的数字.

**返回值(Object)：**BN.js实例

```js
web3.utils.toBN(1234).toString();
> "1234"

web3.utils.toBN('1234').add(web3.utils.toBN('1')).toString();
> "1235"

web3.utils.toBN('0xea').toString();
> "234"
```

## 4.13 hexToNumberString

返回 16 进制字符串的数字表示

**参数:**`hexString` - `String|HEX`: 16 进制字符串

**返回值(String)：**数字字符串

```js
web3.utils.hexToNumberString('0xea');
> "234"
```

## 4.14 hexToNumber

返回 16 进制字符串的数字表示

> 这个对大数不适用, 大数应该使用 web3.utils.toBN.

**参数:**`hexString` - `String|HEX`: 16 进制字符串

**返回值(Number)**

## 4.15 numberToHex

返回给定数字的 16 进制表示

**参数:**`number` - `String|Number|BN|BigNumber`: 数字或数字字符串.

**返回值(String)：**给定数字的 16 进制表示.

```js
web3.utils.numberToHex('234');
> '0xea'
```

## 4.16 hexToUtf8

返回给定 16 进制数字的 UTF-8 字符串表示

**参数:**`hex` - `String`: 要转换为 UTF-8 字符串的 16 进制字符串.

**返回值(String)：**UTF-8 字符串.

```js
web3.utils.hexToUtf8('0x49206861766520313030e282ac');
> "I have 100€"
```

## 4.17 hexToAscii

返回给定 16 进制数字的 ASCII 码字符串表示

**参数:**`hex` - `String`: 要转换为 ASCII 码字符串的 16 进制字符串.

**返回值(String)：**ASCII 码字符串.

```js
web3.utils.hexToAscii('0x49206861766520313030e282ac');
> "I have 100€"
```

## 4.18 utf8ToHex

返回给定 UTF-8 字符串的 16 进制表示。

**参数:**`string` - `String`: 要转换为 16 进制字符串的 UTF-8 字符串.

**返回值(String)：**16 进制字符串

## 4.19 asciiToHex

返回给定 ASCII 码字符串的 16 进制表示

**参数:**`string` - `String`: 要转换为 16 进制字符串的 ASCII 码字符串.

**返回值(String)：**16 进制字符串

## 4.20 hexToBytes

返回给定 16 进制字符串的字节数组表示

**参数:**`hex` - `String|HEX`: 要进行转换的 16 进制串.

**返回值(Array)：**字节数组

## 4.21 bytesToHex

返回一个字节数组的 16 进制字符串表示

**参数:**`byteArray` - `Array`: 要进行转换的字节数组.

**返回值(String)：**16进制字符串

## 4.22 toWei

将任意 ether 值转换为 wei.

**参数:**

- `number` - `String|BN`: 要转换的数字.
- `unit` - `String` (optional, defaults to `"ether"`): 要转换的以太币单位. 支持的单位包括:
  - `noether`: ‘0’
  - `wei`: ‘1’
  - `kwei`: ‘1000’
  - `Kwei`: ‘1000’
  - `babbage`: ‘1000’
  - `femtoether`: ‘1000’
  - `mwei`: ‘1000000’
  - `Mwei`: ‘1000000’
  - `lovelace`: ‘1000000’
  - `picoether`: ‘1000000’
  - `gwei`: ‘1000000000’
  - `Gwei`: ‘1000000000’
  - `shannon`: ‘1000000000’
  - `nanoether`: ‘1000000000’
  - `nano`: ‘1000000000’
  - `szabo`: ‘1000000000000’
  - `microether`: ‘1000000000000’
  - `micro`: ‘1000000000000’
  - `finney`: ‘1000000000000000’
  - `milliether`: ‘1000000000000000’
  - `milli`: ‘1000000000000000’
  - `ether`: ‘1000000000000000000’
  - `kether`: ‘1000000000000000000000’
  - `grand`: ‘1000000000000000000000’
  - `mether`: ‘1000000000000000000000000’
  - `gether`: ‘1000000000000000000000000000’
  - `tether`: ‘1000000000000000000000000000000’

**返回值(String|BN)：**给一个数字字符串就返回一个字符串, 否则返回一个 BN.js 实例.

```js
web3.utils.toWei('1', 'ether');
> "1000000000000000000"
```

## 4.23 fromWei

将任意数量的 wei 转换为 ether.

**参数:**参考**web3.utils.toWei()**

**返回值(String)：**数字字符串

```js
web3.utils.fromWei('1', 'ether');
> "0.000000000000000001"
```

## 4.24 unitMap

显示所有的 ether 单位 和它们对应的 wei 数量.

**参数:**参考**web3.utils.toWei()**

**返回值(Object)：**web3.utils.toWei()参数

## 4.25 padLeft/padRight

对一个字符串进行左/右补齐, 在需要对 16 进制字符串进行填充是很有用.

**参数:**

1. `string` - `String`: 需要左/右补齐的字符串.
2. `characterAmount` - `Number`: 整个字符串所含有的字符数量.
3. `sign` - `String` (可选): 要使用的字符符号, 默认为 `"0"`.

**返回值(String)：**补齐后的字符串.

## 4.26 toTwosComplement

将一个负数转换为补码表示。

**参数:**`number` - `Number|String|BigNumber`: 需要进行转换的数字.

**返回值(String)：**转换后的 16 进制字符串.
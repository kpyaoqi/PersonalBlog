---
title: 17-metadata元数据

date: 2022-07-25	

categories: Solidity0.8.17	

tags: [区块链,Solidity0.8.17]
---	

Solidity 编译器在编译的时候自动生成`xx_metadata.json`的 JSON 文件，中文叫合约的元数据，其中包含了当前合约的相关信息。

# metadata 包含信息

元数据文件具有以下格式。 下面的例子将以人类可读的方式呈现。正确格式化的元数据应正确使用引号，将空白减少到最小，并对所有对象的键值进行排序以得到唯一的格式。（下面的代码注释是不允许的，这里仅用于解释目的。）

```json
{
  // 必选：元数据格式的版本(注意和Solidity版本不是同一个version )
  "version": "1",

  // 必选：源代码的编程语言，一般会选择规范的“子版本”
  "language": "Solidity",

  // 必选：编译器的细节，内容视语言而定。
  "compiler": {
    // 对 Solidity 来说是必须的：编译器的版本
    "version": "0.8.7+commit.e28d00a7",
    // 可选： 生成此输出的编译器二进制文件的哈希值
    "keccak256": "0x123..."
  },

  // 必选：合约的生成信息
  "output": {
    // 必选：合约的 ABI 定义
    "abi": [ /*...*/],
    // 必选：合约的 NatSpec 用户文档
    "userdoc": [ /*...*/],
    // 必选：合约的 NatSpec 开发者文档
    "devdoc": [ /*...*/],
  },

  // 必选：编译器的设置
  "settings": {
    "compilationTarget": {
      "a.sol": "Sum"
    },
    "evmVersion": "london",
    "libraries": {},
    "metadata": {
      // Reflects the setting used in the input json, defaults to false
      "useLiteralContent": true,
      // Reflects the setting used in the input json, defaults to "ipfs"
      "bytecodeHash": "ipfs"
    },
    // 可选： 优化器的设置（ enabled 默认设为 false ）
    "optimizer": {
      "enabled": true,
      "runs": 200
    },
    // 对 Solidity 来说是必须的： 已排序的重定向列表
    "remappings": [":g/dir"],
  },

  // 必选：编译的源文件／源单位，键值为文件名
  "sources": {
    "myFile.sol": {
      // 必选：源文件的 keccak256 哈希值
      "keccak256": "0x123...",

      // Optional: 在源文件中定义的 SPDX license 标识
      "license": "MIT",
      // 必选（除非定义了 content，详见下文）：
      // 已排序的源文件的URL，URL的协议可以是任意的，但建议使用 Swarm 的URL
      "urls": [
        "bzz-raw://fd33d...",
        "dweb:/ipfs/Qme8Vrt"
      ]
    },
    // 可选
    "mortal": {
      // 必选：源文件的 keccak256 哈希值
      "keccak256": "0x234...",
      // 必选（除非定义了“urls”）： 源文件的字面内容
      "content": "contract mortal is owned { function kill() {
        if (msg.sender == owner) selfdestruct(owner); } }"
    }
  },
}
```

### metadata 的作用

metadata 主要是为了更安全地与合约进行交互并验证其源代码。

- 查询编译器版本
- 所使用的源代码
- ABI
- natspec 文档

编译器会将元数据文件的 Swarm 哈希值附加到每个合约的字节码末尾（详情请参阅下文），以便你可以以认证的方式获取该文件，而不必求助于中心化的数据提供者。当然，你必须将元数据文件发布到 Swarm（或其他服务），以便其他人可以访问它。 该文件可以通过使用`solc --metadata` 来生成，并被命名为 `ContractName_meta.json` 。 它将包含源代码的在 Swarm 上的引用，因此你必须上传所有源文件和元数据文件。

------

⚠️ 警告: 由于生成的合约的字节码包含元数据的哈希值，因此对元数据的任何更改都会导致字节码的更改。
此外，由于元数据包含所有使用的源代码的哈希值，所以任何源代码中的，哪怕是一个空格的变化都将导致不同的元数据，并随后产生不同的字节代码。

⚠️ 警告: 需注意，上面的 ABI 没有固定的顺序，随编译器的版本而不同。

# 字节码中元数据哈希的编码

由于在将来可能会支持其他方式来获取元数据文件， 类似`{"bzzr0"：<Swarm hash>}` 的键值对，将会以
CBOR (https://tools.ietf.org/html/rfc7049) 编码来存储。由于这种编码的起始位不容易找到，因此添加两个字节来表述其长度，以大端方式编码。所以，当前版本的 Solidity 编译器，将以下内容添加到部署的字节码的末尾

```
    0xa2
    0x64 'i' 'p' 'f' 's' 0x58 0x22 <34 bytes IPFS hash>
    0x64 's' 'o' 'l' 'c' 0x43 <3 byte version encoding>
    0x00 0x33
```

因此，为了检索数据，可以检查已部署字节码的末尾以匹配该模式，并使用 IPFS 哈希来检索文件。

solc 的发布版本使用如上所示的版本的 3 字节编码（major, minor and patch version number 版本号各一个字节），而预发布版本将使用完整的版本字符串，包括提交哈希和构建日期。

CBOR 映射还可以包含其他密钥，因此最好完全解码数据而不是依赖以 `0xa264` 开头的数据。 例如，如果使用任何影响代码生成的实验性功能，则映射也将包含 `"experimental"：true`。

编译器目前默认使用元数据的 IPFS 哈希，但将来也可能使用 bzzr1 哈希或其他一些哈希，因此不要依赖此序列以 `0xa2 0x64 'i' 'p' 'f' 's'` 开头 的。 我们可能还会向此 CBOR 结构中添加其他数据，因此最好的选择是使用适当的 CBOR 解析器。

# 自动化接口生成和 natspec 使用

元数据以下列方式使用：通过钱包想要与合约交互时检索合约代码，然后检索文件的 IPFS/Swarm 哈希。该文件被 JSON 解码为上面的结构。

组件可以使用 ABI 自动为合约生成一个基本的用户界面。

此外，钱包可以使用 NatSpec 用户文档，在用户与合约交互，授权请求签名时候做辅助工作。

# 源代码如何验证？

为了验证编译，可以通过元数据文件中的链接从 IPFS/Swarm 中获取源代码。获取到的源码，会根据元数据中指定的设置，被正确版本的编译器所处理。处理得到的字节码会与创建交易的数据或者 `CREATE` 操作码使用的数据进行比较。这会自动验证元数据，因为它的哈希值是字节码的一部分。而额外的数据，则是与基于接口进行编码并展示给用户的构造输入数据相符的。

在 [sourcify](https://github.com/ethereum/sourcify) 库([npm package](https://www.npmjs.com/package/source-verify))可以看到如何使用该特性的示例代码。

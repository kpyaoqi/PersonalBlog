---
title: 02-核心 Core

date: 2021-11-08	

categories: EIP笔记	

tags: [EIP笔记]
---	

# Core 核心

## 最终的

| 数字                                            | 标题                                                  | 内容                                                         |
| ----------------------------------------------- | ----------------------------------------------------- | ------------------------------------------------------------ |
| [2](https://eips.ethereum.org/EIPS/eip-2)       | Homestead 硬分叉变化                                  | 增加了创建合约的Gas成本(create不受影响)、调整难度算法        |
| [5](https://eips.ethereum.org/EIPS/eip-5)       | “RETURN”和“CALL”的 Gas 使用情况                       | 只对CALL返回时实际写入的内存收取gas费用                      |
| [7](https://eips.ethereum.org/EIPS/eip-7)       | DELEGATECALL                                          | 添加一个新的操作码`DELEGATECALL`                             |
| [100](https://eips.ethereum.org/EIPS/eip-100)   | 将难度调整更改为目标平均区块时间（包括叔块）          | 调整难度算法                                                 |
| [140](https://eips.ethereum.org/EIPS/eip-140)   | REVERT指令                                            | 增加了REVERT指令                                             |
| [141](https://eips.ethereum.org/EIPS/eip-141)   | 指定无效的EVM指令                                     | 无效指令可以用作中止执行的明显原因(从未使用)                 |
| [145](https://eips.ethereum.org/EIPS/eip-145)   | EVM 中的按位移位指令                                  | 提供本机按位移位，其成本与其他算术运算相当                   |
| [150](https://eips.ethereum.org/EIPS/eip-150)   | IO 密集型操作的 Gas 成本变化                          | 修改了一些指令的Gas成本                                      |
| [152](https://eips.ethereum.org/EIPS/eip-152)   | 添加 BLAKE2 压缩函数 `F` 预编译                       | 引入了一个新的预编译合约，该合约实现了BLAKE2加密哈希算法中使用的压缩功能`F` |
| [155](https://eips.ethereum.org/EIPS/eip-155)   | 简单的重放攻击保护                                    | 当出于签名目的计算交易的散列时，您**应该**散列九个 rlp 编码元素，而不是仅散列六个 rlp 编码元素 |
| [158](https://eips.ethereum.org/EIPS/eip-158)   | 状态清算                                              | 消除了大量空帐户，缩小负担                                   |
| [160](https://eips.ethereum.org/EIPS/eip-160)   | EXP成本增加                                           | EXP 的 Gas 成本从指数中每字节 10 + 10 增加到指数中每字节 10 + 50 |
| [161](https://eips.ethereum.org/EIPS/eip-161)   | 状态特里树清理（不变性保留替代方案）                  |                                                              |
| [170](https://eips.ethereum.org/EIPS/eip-170)   | 合约代码大小限制                                      | 如果合约创建初始化返回长度**超过** `MAX_CODE_SIZE`字节的数据，合约创建将失败并出现耗尽气体错误 |
| [196](https://eips.ethereum.org/EIPS/eip-196)   | 椭圆曲线上的加法和标量乘法的预编译合约 alt_bn128      | 为了在区块 Gas 限制内执行 zkSNARK 验证，需要用于椭圆曲线操作的预编译合约 |
| [197](https://eips.ethereum.org/EIPS/eip-197)   | 用于椭圆曲线上最佳 ate 配对检查的预编译合约 alt_bn128 | 为了在区块 Gas 限制内执行 zkSNARK 验证，需要用于椭圆曲线配对操作的预编译合约。 |
| [198](https://eips.ethereum.org/EIPS/eip-198)   | 大整数模幂                                            | 允许在 EVM 内部进行高效的 RSA 验证，以及其他形式的基于数论的密码学。 |
| [211](https://eips.ethereum.org/EIPS/eip-211)   | 新操作码：RETURNDATASIZE 和 RETURNDATACOPY            |                                                              |
| [214](https://eips.ethereum.org/EIPS/eip-214)   | 新操作码 STATICCALL                                   |                                                              |
| [225](https://eips.ethereum.org/EIPS/eip-225)   | Clique 权威证明共识协议                               | Clique 是一种权威证明共识协议。它模仿了以太坊主网的设计，因此可以以最小的努力将其添加到任何客户端 |
| [649](https://eips.ethereum.org/EIPS/eip-649)   | 大都会难度炸弹延迟和区块奖励减少                      |                                                              |
| [658](https://eips.ethereum.org/EIPS/eip-658)   | 在收据中嵌入交易状态代码                              |                                                              |
| [684](https://eips.ethereum.org/EIPS/eip-684)   | 发生碰撞时恢复创建                                    |                                                              |
| [1014](https://eips.ethereum.org/EIPS/eip-1014) | 瘦CREATE2                                             |                                                              |
| [1052](https://eips.ethereum.org/EIPS/eip-1052) | EXTCODEHASH 操作码                                    |                                                              |
| [1108](https://eips.ethereum.org/EIPS/eip-1108) | 减少 alt_bn128 预编译 Gas 成本                        |                                                              |
| [1234](https://eips.ethereum.org/EIPS/eip-1234) | 君士坦丁堡难度炸弹延迟和区块奖励调整                  |                                                              |
| [1283](https://eips.ethereum.org/EIPS/eip-1283) | SSTORE 的净燃气计量，无需脏地图                       |                                                              |
| [1344](https://eips.ethereum.org/EIPS/eip-1344) | ChainID操作码                                         |                                                              |
| [1559](https://eips.ethereum.org/EIPS/eip-1559) | ETH 1.0 链的费用市场变化                              |                                                              |
| [1884](https://eips.ethereum.org/EIPS/eip-1884) | 依赖特里结构大小的操作码的重新定价                    |                                                              |
| [2028](https://eips.ethereum.org/EIPS/eip-2028) | 交易数据gas成本降低                                   |                                                              |
| [2200](https://eips.ethereum.org/EIPS/eip-2200) | 净燃气计量的结构化定义                                |                                                              |
| [2384](https://eips.ethereum.org/EIPS/eip-2384) | 缪尔冰川难度炸弹延迟                                  |                                                              |
| [2565](https://eips.ethereum.org/EIPS/eip-2565) | ModExp 天然气成本                                     |                                                              |
| [2681](https://eips.ethereum.org/EIPS/eip-2681) | 将帐户随机数限制为 2^64-1                             |                                                              |
| [2718](https://eips.ethereum.org/EIPS/eip-2718) | 打字交易信封                                          |                                                              |
| [2929](https://eips.ethereum.org/EIPS/eip-2929) | 状态访问操作码的 Gas 成本增加                         |                                                              |
| [2930](https://eips.ethereum.org/EIPS/eip-2930) | 可选访问列表                                          |                                                              |
| [3198](https://eips.ethereum.org/EIPS/eip-3198) | 基本费操作码                                          |                                                              |
| [3529](https://eips.ethereum.org/EIPS/eip-3529) | 减少退款                                              |                                                              |
| [3541](https://eips.ethereum.org/EIPS/eip-3541) | 拒绝以 0xEF 字节开头的新合约代码                      |                                                              |
| [3554](https://eips.ethereum.org/EIPS/eip-3554) | 难度炸弹延迟至 2021  12 月                            |                                                              |
| [3607](https://eips.ethereum.org/EIPS/eip-3607) | 使用已部署的代码拒绝来自发件人的交易                  |                                                              |
| [3651](https://eips.ethereum.org/EIPS/eip-3651) | 温暖的COINBASE                                        |                                                              |
| [3675](https://eips.ethereum.org/EIPS/eip-3675) | 将共识升级为权益证明                                  |                                                              |
| [3855](https://eips.ethereum.org/EIPS/eip-3855) | PUSH0指令                                             |                                                              |
| [3860](https://eips.ethereum.org/EIPS/eip-3860) | 限制和计量初始化代码                                  |                                                              |
| [4345](https://eips.ethereum.org/EIPS/eip-4345) | 难度炸弹延迟至 2022  6 月                             |                                                              |
| [4399](https://eips.ethereum.org/EIPS/eip-4399) | 用 PREVRANDAO 替换 DIFFICULTY 操作码                  |                                                              |
| [4895](https://eips.ethereum.org/EIPS/eip-4895) | 信标链将提款作为操作                                  |                                                              |
| [5133](https://eips.ethereum.org/EIPS/eip-5133) | 将难度炸弹推迟到 2022  9 月中旬                       |                                                              |

## 最后一次call

| 数字                                            | 审核结束   | 标题           |
| ----------------------------------------------- | ---------- | -------------- |
| [1153](https://eips.ethereum.org/EIPS/eip-1153) | 2022-12-08 | 瞬时存储操作码 |

## 审查

| 数字                                            | 标题                      |
| ----------------------------------------------- | ------------------------- |
| [663](https://eips.ethereum.org/EIPS/eip-663)   | 无限的 SWAP 和 DUP 指令   |
| [2294](https://eips.ethereum.org/EIPS/eip-2294) | 显式绑定到 Chain ID 大小  |
| [3074](https://eips.ethereum.org/EIPS/eip-3074) | AUTH 和 AUTHCALL 操作码   |
| [3540](https://eips.ethereum.org/EIPS/eip-3540) | EOF - EVM 对象格式 v1     |
| [3670](https://eips.ethereum.org/EIPS/eip-3670) | EOF - 代码验证            |
| [4200](https://eips.ethereum.org/EIPS/eip-4200) | EOF - 静态相对跳转        |
| [4750](https://eips.ethereum.org/EIPS/eip-4750) | EOF——函数                 |
| [4758](https://eips.ethereum.org/EIPS/eip-4758) | 停用自毁功能              |
| [4844](https://eips.ethereum.org/EIPS/eip-4844) | 分片 Blob 交易            |
| [5450](https://eips.ethereum.org/EIPS/eip-5450) | EOF - 堆栈验证            |
| [5656](https://eips.ethereum.org/EIPS/eip-5656) | MCOPY——内存复制指令       |
| [5920](https://eips.ethereum.org/EIPS/eip-5920) | 支付操作码                |
| [6188](https://eips.ethereum.org/EIPS/eip-6188) | 随机数上限                |
| [6189](https://eips.ethereum.org/EIPS/eip-6189) | 别名合约                  |
| [6190](https://eips.ethereum.org/EIPS/eip-6190) | Verkle 兼容的自毁         |
| [7044](https://eips.ethereum.org/EIPS/eip-7044) | 永久有效的签名自愿退出    |
| [7045](https://eips.ethereum.org/EIPS/eip-7045) | 增加最大证明包含槽        |
| [7212](https://eips.ethereum.org/EIPS/eip-7212) | 预编译 secp256r1 曲线支持 |

## 草稿

| 数字                                            | 标题                              |
| ----------------------------------------------- | --------------------------------- |
| [1418](https://eips.ethereum.org/EIPS/eip-1418) | 区块链存储租金支付                |
| [3102](https://eips.ethereum.org/EIPS/eip-3102) | 二元特里结构                      |
| [3455](https://eips.ethereum.org/EIPS/eip-3455) | SUDO 操作码                       |
| [4788](https://eips.ethereum.org/EIPS/eip-4788) | EVM 中的信标块根                  |
| [5000](https://eips.ethereum.org/EIPS/eip-5000) | MULDIV指令                        |
| [5027](https://eips.ethereum.org/EIPS/eip-5027) | 取消合约代码大小限制              |
| [5081](https://eips.ethereum.org/EIPS/eip-5081) | 过期火车动作                      |
| [5806](https://eips.ethereum.org/EIPS/eip-5806) | 委托交易                          |
| [5988](https://eips.ethereum.org/EIPS/eip-5988) | 添加Poseidon哈希函数预编译        |
| [6046](https://eips.ethereum.org/EIPS/eip-6046) | 将 SELFDESTRUCT 替换为 DEACTIVATE |
| [6110](https://eips.ethereum.org/EIPS/eip-6110) | 在链上提供验证者存款              |
| [6206](https://eips.ethereum.org/EIPS/eip-6206) | EOF-JUMPF指令                     |
| [6404](https://eips.ethereum.org/EIPS/eip-6404) | SSZ 交易根                        |
| [6465](https://eips.ethereum.org/EIPS/eip-6465) | SSZ提款根                         |
| [6466](https://eips.ethereum.org/EIPS/eip-6466) | SSZ 收据根                        |
| [6475](https://eips.ethereum.org/EIPS/eip-6475) | SSZ 可选                          |
| [6493](https://eips.ethereum.org/EIPS/eip-6493) | SSZ交易签名方案                   |
| [6780](https://eips.ethereum.org/EIPS/eip-6780) | 仅在同一事务中自毁                |
| [6810](https://eips.ethereum.org/EIPS/eip-6810) | 事后级联恢复                      |
| [6811](https://eips.ethereum.org/EIPS/eip-6811) | 前往月球——10 分钟街区             |
| [6888](https://eips.ethereum.org/EIPS/eip-6888) | EVM 中的数学检查                  |
| [6913](https://eips.ethereum.org/EIPS/eip-6913) | 设置代码指令                      |
| [6968](https://eips.ethereum.org/EIPS/eip-6968) | 基于 EVM 的 L2 上的合同担保收入   |
| [6988](https://eips.ethereum.org/EIPS/eip-6988) | 当选的区块提议者尚未被削减        |
| [7002](https://eips.ethereum.org/EIPS/eip-7002) | 执行层可触发退出                  |
| [7069](https://eips.ethereum.org/EIPS/eip-7069) | 修改了 CALL 指令                  |
| [7266](https://eips.ethereum.org/EIPS/eip-7266) | 删除 BLAKE2 压缩预编译            |
| [7377](https://eips.ethereum.org/EIPS/eip-7377) | 移民交易                          |

## 停滞不前

| 数字                                            | 标题                                                         |
| ----------------------------------------------- | ------------------------------------------------------------ |
| [86](https://eips.ethereum.org/EIPS/eip-86)     | 交易来源和签名的抽象                                         |
| [101](https://eips.ethereum.org/EIPS/eip-101)   | 宁静货币和加密抽象                                           |
| [210](https://eips.ethereum.org/EIPS/eip-210)   | 区块哈希重构                                                 |
| [615](https://eips.ethereum.org/EIPS/eip-615)   | EVM 的子例程和静态跳转                                       |
| [616](https://eips.ethereum.org/EIPS/eip-616)   | EVM 的 SIMD 操作                                             |
| [665](https://eips.ethereum.org/EIPS/eip-665)   | 添加Ed25519签名验证的预编译合约                              |
| [689](https://eips.ethereum.org/EIPS/eip-689)   | 合约地址冲突导致异常停顿                                     |
| [698](https://eips.ethereum.org/EIPS/eip-698)   | 操作码 0x46 区块奖励                                         |
| [858](https://eips.ethereum.org/EIPS/eip-858)   | 减少区块奖励并延迟难度炸弹                                   |
| [969](https://eips.ethereum.org/EIPS/eip-969)   | 修改 ethash 以使现有的专用硬件实现失效                       |
| [1010](https://eips.ethereum.org/EIPS/eip-1010) | 0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B 和 0x15E55EF43efA8348dDaeAa455F16C43B64917e3c 之间的一致性 |
| [1011](https://eips.ethereum.org/EIPS/eip-1011) | 混合 Casper FFG                                              |
| [1015](https://eips.ethereum.org/EIPS/eip-1015) | 可配置的链上发行                                             |
| [1051](https://eips.ethereum.org/EIPS/eip-1051) | EVM 的溢出检查                                               |
| [1057](https://eips.ethereum.org/EIPS/eip-1057) | ProgPoW，一种程序化的工作量证明                              |
| [1087](https://eips.ethereum.org/EIPS/eip-1087) | SSTORE 运营的净燃气计量                                      |
| [1109](https://eips.ethereum.org/EIPS/eip-1109) | PRECOMPILEDCALL 操作码（删除预编译合约的 CALL 成本）         |
| [1227](https://eips.ethereum.org/EIPS/eip-1227) | 拆除难度炸弹并重置区块奖励                                   |
| [1276](https://eips.ethereum.org/EIPS/eip-1276) | 消除难度炸弹并调整君士坦丁堡转移的区块奖励                   |
| [1285](https://eips.ethereum.org/EIPS/eip-1285) | 增加 CALL 操作码中的 Gcallstipend 气体                       |
| [1295](https://eips.ethereum.org/EIPS/eip-1295) | 修改以太坊PoW激励结构并延迟难度炸弹                          |
| [1352](https://eips.ethereum.org/EIPS/eip-1352) | 指定预编译/系统合约的限制地址范围                            |
| [1380](https://eips.ethereum.org/EIPS/eip-1380) | 降低自我调用的 Gas 成本                                      |
| [1482](https://eips.ethereum.org/EIPS/eip-1482) | 定义最大区块时间戳漂移                                       |
| [1485](https://eips.ethereum.org/EIPS/eip-1485) | TEthashV1                                                    |
| [1681](https://eips.ethereum.org/EIPS/eip-1681) | 临时重放保护                                                 |
| [1702](https://eips.ethereum.org/EIPS/eip-1702) | 通用帐户版本控制方案                                         |
| [1829](https://eips.ethereum.org/EIPS/eip-1829) | 椭圆曲线线性组合的预编译                                     |
| [1895](https://eips.ethereum.org/EIPS/eip-1895) | 支持椭圆曲线循环                                             |
| [1930](https://eips.ethereum.org/EIPS/eip-1930) | 具有严格气体语义的 CALL。如果没有足够的可用气体，请恢复。    |
| [1959](https://eips.ethereum.org/EIPS/eip-1959) | 用于检查 chainID 是否是 chainID 历史记录的一部分的新操作码   |
| [1962](https://eips.ethereum.org/EIPS/eip-1962) | EC 算术和与运行时定义的配对                                  |
| [1965](https://eips.ethereum.org/EIPS/eip-1965) | 检查 chainID 在特定区块号上是否有效的方法                    |
| [1985](https://eips.ethereum.org/EIPS/eip-1985) | 某些 EVM 参数的合理限制                                      |
| [2014](https://eips.ethereum.org/EIPS/eip-2014) | 扩展状态预言机                                               |
| [2026](https://eips.ethereum.org/EIPS/eip-2026) | 状态租金 H - 账户固定预付款                                  |
| [2027](https://eips.ethereum.org/EIPS/eip-2027) | State Ren C - 净合同规模核算                                 |
| [2029](https://eips.ethereum.org/EIPS/eip-2029) | State Rent A - 国家柜台合同                                  |
| [2031](https://eips.ethereum.org/EIPS/eip-2031) | 国租B - 净交易柜台                                           |
| [2035](https://eips.ethereum.org/EIPS/eip-2035) | 无状态客户端 - 重新定价 SLOAD 和 SSTORE 以支付区块证明       |
| [2045](https://eips.ethereum.org/EIPS/eip-2045) | EVM 操作码的粒子气体成本                                     |
| [2046](https://eips.ethereum.org/EIPS/eip-2046) | 减少预编译静态调用的 Gas 成本                                |
| [2242](https://eips.ethereum.org/EIPS/eip-2242) | 交易后数据                                                   |
| [2327](https://eips.ethereum.org/EIPS/eip-2327) | BEGINDATA 操作码                                             |
| [2330](https://eips.ethereum.org/EIPS/eip-2330) | EXTSLOAD 操作码                                              |
| [2474](https://eips.ethereum.org/EIPS/eip-2474) | Coinbase 电话                                                |
| [2488](https://eips.ethereum.org/EIPS/eip-2488) | 弃用 CALLCODE 操作码                                         |
| [2515](https://eips.ethereum.org/EIPS/eip-2515) | 实施难度冻结                                                 |
| [2537](https://eips.ethereum.org/EIPS/eip-2537) | BLS12-381 曲线操作的预编译                                   |
| [2539](https://eips.ethereum.org/EIPS/eip-2539) | BLS12-377 曲线运算                                           |
| [2542](https://eips.ethereum.org/EIPS/eip-2542) | 新操作码 TXGASLIMIT 和 CALLGASLIMIT                          |
| [2583](https://eips.ethereum.org/EIPS/eip-2583) | 对帐户 trie 未命中的惩罚                                     |
| [2584](https://eips.ethereum.org/EIPS/eip-2584) | 具有覆盖树的 Trie 格式转换                                   |
| [2593](https://eips.ethereum.org/EIPS/eip-2593) | ETH 1.0 链的自动扶梯费用市场变化                             |
| [2666](https://eips.ethereum.org/EIPS/eip-2666) | 预编译和 Keccak256 函数的重新定价                            |
| [2803](https://eips.ethereum.org/EIPS/eip-2803) | 丰富的交易                                                   |
| [2926](https://eips.ethereum.org/EIPS/eip-2926) | 基于块的代码默克尔化                                         |
| [2935](https://eips.ethereum.org/EIPS/eip-2935) | 在状态中保存历史区块哈希                                     |
| [2936](https://eips.ethereum.org/EIPS/eip-2936) | 用于自毁合约的 EXTCLEAR 操作码                               |
| [2937](https://eips.ethereum.org/EIPS/eip-2937) | SET_INDESTRUCTIBLE 操作码                                    |
| [2938](https://eips.ethereum.org/EIPS/eip-2938) | 账户抽象                                                     |
| [2970](https://eips.ethereum.org/EIPS/eip-2970) | IS_STATIC 操作码                                             |
| [2997](https://eips.ethereum.org/EIPS/eip-2997) | IMPERSONATECALL 操作码                                       |
| [3026](https://eips.ethereum.org/EIPS/eip-3026) | BW6-761曲线运算                                              |
| [3068](https://eips.ethereum.org/EIPS/eip-3068) | BN256 HashToCurve 算法的预编译                               |
| [3143](https://eips.ethereum.org/EIPS/eip-3143) | 将区块奖励增加至 5 ETH                                       |
| [3220](https://eips.ethereum.org/EIPS/eip-3220) | 跨链标识符规范                                               |
| [3238](https://eips.ethereum.org/EIPS/eip-3238) | 难度炸弹延迟至 2022 二季度                                   |
| [3267](https://eips.ethereum.org/EIPS/eip-3267) | 将以太坊费用支付给未来工资                                   |
| [3298](https://eips.ethereum.org/EIPS/eip-3298) | 取消退款                                                     |
| [3300](https://eips.ethereum.org/EIPS/eip-3300) | 逐步取消退款                                                 |
| [3322](https://eips.ethereum.org/EIPS/eip-3322) | 账户储气操作码                                               |
| [3336](https://eips.ethereum.org/EIPS/eip-3336) | EVM 的分页内存分配                                           |
| [3337](https://eips.ethereum.org/EIPS/eip-3337) | 内存加载和存储操作的帧指针支持                               |
| [3368](https://eips.ethereum.org/EIPS/eip-3368) | 将区块奖励增加至 3 ETH，计划 2 衰减至 1 ETH                  |
| [3372](https://eips.ethereum.org/EIPS/eip-3372) | ethash 的 5 个 FNV 素数                                      |
| [3403](https://eips.ethereum.org/EIPS/eip-3403) | 部分取消退款                                                 |
| [3416](https://eips.ethereum.org/EIPS/eip-3416) | 天然气溢价中位数                                             |
| [3436](https://eips.ethereum.org/EIPS/eip-3436) | 扩展派块选择规则                                             |
| [3508](https://eips.ethereum.org/EIPS/eip-3508) | 交易数据操作码                                               |
| [3520](https://eips.ethereum.org/EIPS/eip-3520) | 交易目标操作码                                               |
| [3521](https://eips.ethereum.org/EIPS/eip-3521) | 降低访问列表成本                                             |
| [3534](https://eips.ethereum.org/EIPS/eip-3534) | 受限链上下文类型交易                                         |
| [3584](https://eips.ethereum.org/EIPS/eip-3584) | 阻止访问列表                                                 |
| [3690](https://eips.ethereum.org/EIPS/eip-3690) | EOF - JUMPDEST 表                                            |
| [3756](https://eips.ethereum.org/EIPS/eip-3756) | 气体限制上限                                                 |
| [3788](https://eips.ethereum.org/EIPS/eip-3788) | 严格执行chainId                                              |
| [3978](https://eips.ethereum.org/EIPS/eip-3978) | 恢复时 Gas 退款                                              |
| [4396](https://eips.ethereum.org/EIPS/eip-4396) | 时间感知的基本费用计算                                       |
| [4488](https://eips.ethereum.org/EIPS/eip-4488) | 通过总调用数据限制，交易调用数据 Gas 成本降低                |
| [4520](https://eips.ethereum.org/EIPS/eip-4520) | 以 EB 和 EC 为前缀的多字节操作码。                           |
| [4573](https://eips.ethereum.org/EIPS/eip-4573) | EVM 的程序                                                   |
| [4747](https://eips.ethereum.org/EIPS/eip-4747) | 简化 EIP-161                                                 |
| [4760](https://eips.ethereum.org/EIPS/eip-4760) | 自毁炸弹                                                     |
| [4762](https://eips.ethereum.org/EIPS/eip-4762) | 无状态 Gas 成本变化                                          |
| [4803](https://eips.ethereum.org/EIPS/eip-4803) | 将交易 Gas 限制为最大 2^63-1                                 |
| [4863](https://eips.ethereum.org/EIPS/eip-4863) | 信标链推送提现                                               |
| [5003](https://eips.ethereum.org/EIPS/eip-5003) | 使用 AUTHUSURP 将代码插入 EOA                                |
| [5022](https://eips.ethereum.org/EIPS/eip-5022) | 将 SSTORE 的价格从零提高到非零至 40k Gas                     |
| [5065](https://eips.ethereum.org/EIPS/eip-5065) | 以太币转账说明                                               |
| [5283](https://eips.ethereum.org/EIPS/eip-5283) | 用于重入保护的信号量                                         |
| [5478](https://eips.ethereum.org/EIPS/eip-5478) | CREATE2COPY 操作码                                           |

## 取消

| 数字                                            | 标题                                                   |
| ----------------------------------------------- | ------------------------------------------------------ |
| [3](https://eips.ethereum.org/EIPS/eip-3)       | 添加 CALLDEPTH 操作码                                  |
| [908](https://eips.ethereum.org/EIPS/eip-908)   | 奖励客户建立可持续的网络                               |
| [999](https://eips.ethereum.org/EIPS/eip-999)   | 恢复合约代码0x863DF6BFa4469f3ead0bE8f9F2AAE51c91A907b4 |
| [1240](https://eips.ethereum.org/EIPS/eip-1240) | 移除难度炸弹                                           |
| [1355](https://eips.ethereum.org/EIPS/eip-1355) | 以太坊1a                                               |
| [1682](https://eips.ethereum.org/EIPS/eip-1682) | 仓储租金                                               |
| [1706](https://eips.ethereum.org/EIPS/eip-1706) | 禁用 SSTORE，gasleft 低于call津贴                      |
| [1890](https://eips.ethereum.org/EIPS/eip-1890) | 对可持续生态系统资金的承诺                             |
| [2025](https://eips.ethereum.org/EIPS/eip-2025) | 资助 Eth1.x 的区块奖励提案                             |
| [2315](https://eips.ethereum.org/EIPS/eip-2315) | EVM 的简单子例程                                       |
| [2677](https://eips.ethereum.org/EIPS/eip-2677) | 限制“initcode”的大小                                   |
| [2711](https://eips.ethereum.org/EIPS/eip-2711) | 赞助、到期和批量交易。                                 |
| [2733](https://eips.ethereum.org/EIPS/eip-2733) | 交易套餐                                               |
| [2780](https://eips.ethereum.org/EIPS/eip-2780) | 减少内在交易气体                                       |
| [2972](https://eips.ethereum.org/EIPS/eip-2972) | 打包的旧交易                                           |
| [3332](https://eips.ethereum.org/EIPS/eip-3332) | MEDGAASPRICE 操作码                                    |
| [3338](https://eips.ethereum.org/EIPS/eip-3338) | 将帐户随机数限制为 2^52                                |
| [3374](https://eips.ethereum.org/EIPS/eip-3374) | 可预测的工作量证明（POW）日落                          |
| [3382](https://eips.ethereum.org/EIPS/eip-3382) | 硬编码的区块 Gas 限制                                  |
| [3779](https://eips.ethereum.org/EIPS/eip-3779) | EVM 更安全的控制流程                                   |

|      |      |
| ---- | ---- |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
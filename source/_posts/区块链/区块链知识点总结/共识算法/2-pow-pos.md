---
title: 2-pow-pos

date: 2023-01-23	

categories: 共识算法	

tags: [区块链,区块链知识点总结,共识算法]
---	

## POW **工作量证明**

**提交一份用来确认你做过一定量的工作的证明。**监测工作的整个过程通常是极为低效的，而**通过对工作的结果进行认证来证明完成了相应的工作量是一种非常高效的方式。**

工作量证明最常用的技术原理是哈希函数。由于输入哈希函数h的任意值n，会对应到一个h(n)结果，而n只要变动一个位元，就会引起雪崩效应，所以几乎无法从h(n)反推回n，因此借由指定查找h(n)的特征，**让用户进行大量的穷举运算，就可以达成工作量证明**。

**特点**：稳定可靠(攻击者必须投入超过总体一半的运算量51%攻击，难于实现)

​			非常浪费能源，技术人员开发出了由ASIC组成的特制计算设备（矿机），垄断算力

## POS **权益证明**

Pow带来的问题最主要的是能耗问题，通过挖矿比拼的是设备的数量，其实也就是拼钱。因此为什么不直接拼钱来分成奖励呢？省下了不必要的挖矿过程。因此基于这个理论，提出权益证明。

权益证明的原理是指**通过抵押代币，并计算抵押的代币数量和抵押时间的乘积，也就是币龄。每次出块时，币龄最大的矿工获得出块权。产生区块后，该矿工获得出块奖励，同时币龄清零并重新开始计算，如此类推。**

**特点**:硬件要求低，不需要消耗巨大的能源，出块速度更快。缺点是：去中心化程度不高

## DPOW **延迟工作量证明**

DPow的原理是**允许一个区块链利用第二个区块链的哈希算力所提供的安全。**该机制是通过一组公证员节点实现的。公证员节点实现将第一个区块链的数据添加到第二个区块链中。进而，第二个区块链请求在两个区块链间达成妥协，弱化第一个区块链的安全。

**DPOW系统中有两种类型的节点：公证人节点和正常节点**

64 个公证人节点是由 DPoW 区块链的权益持有者（stakeholder）选举产生的，它们可从 DPoW 区块链向所附加的 PoW 区块链添加经公证确认的块。**一旦添加了一个块，该块的哈希值将被添加到由 33 个公证人节点签署的 Bitcoin 交易中，并创建一个哈希到 Bitcoin 区块链的 dPow 块记录**。该记录已被网络中的大多数公证人节点公证。

为避免公证人节点间在挖矿上产生战争，进而降低网络的效率，Komodo 设计采用轮询机制的挖矿方法，该方法具有两种运行模式。

**在“无公证人”（No Notary）模式下，支持所有网络节点参与挖矿，这类似于传统 PoW 共识机制。而在“公证人激活”（Notaries Active）模式下，网络公证人使用一种显著降低的网络难度率挖矿**。“公证人激活”模式下，允许每位公证人使用其当前的难度挖掘一个区块，而其它公证人节点必须采用 10 倍难度挖矿，所有正常节点使用公证人节点难度的 100 倍挖矿。

**特点：**只有使用PoW或PoS的区块链，才能采用这种共识算法；在“公证员激活”模式下，必须校准不同节点的哈希率

## DPOS **委托权益证明**

先**由代币持有者投票选出若干个见证人**，**又称为超级节点**，**再由这些见证人轮流出块**。这种做法是在运行效率和去中心化两者中获得平衡。见证人类似于股份制公司里的董事会成员。**普通的代币持有者只有进行投票的权利，持有的代币越多，他能投的票数也越多。获得投票数最高的若干候选人将当选见证人。**见证人有任期，一般是一周。一周过后重新选举新的见证人。**每个区块如果能获得一定比例（EOS为大于2/3）的所有见证人的同意，这个区块就是有效的。**区块链上的所有的升级和提议，都必须经过委员会（由所有见证人组成）的同意才能执行。

**特点：**不需要消耗巨大的能源，运行效率更高，出块速度更快，不容易产生分叉。缺点是：去中心化程度不高，容易出现贿选问题

## **PoB（Proof of Burn）焚烧证明机制**

**是一种通过焚烧自己手中的代币来表决谁拥有对网络的领导地位的承诺**

在基于DPoW的区块链中，矿工挖矿所获得的不再是奖励的代币，而是可以焚烧的“wood”——燃木。矿工使用自己的算力，通过哈希算法，最终证明自己的工作量之后，获取对应的wood，wood不可交易。当wood积攒到一定量之后，可以前往燃烧场地燃烧wood。**通过一组算法计算后，燃烧较多wood的人或者BP或者一组BP可以获取下个事件段出块的权利，成功出块后获取奖励（代币）**。
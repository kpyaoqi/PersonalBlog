---
title: 06-Sync同步

date: 2022-11-07	

categories: 以太坊源码	

tags: [区块链,以太坊源码]
---	

# 概述

 在前面的章节中，我们已经讨论了在以太坊中 Transactions 是从 Transaction Pool 中，被 Validator/Miner 们验证打包，最终被保存在区块链中。那么，接下来的问题是，**Transaction 是怎么被进入到 Transaction Pool 中的呢？**基于同样的思考方式，那么一个刚刚在某个节点被打包好的 Block，它又将**怎么传输到区块链网络中的其他节点**那里，并**最终实现 Blockchain 长度是一致**的呢？在本章中，我们就来探索一下，节点是如何发送和接收 Transaction 和 Block 的。

## syncs Transactions：同步交易状态

## syncs Blocks：同步区块状态
---
title: 08-Bloom Filter布隆过滤器

date: 2022-11-12	

categories: 以太坊源码	

tags: [区块链,以太坊源码]
---	

# Bloom Filter

**Bloom Filter 是一种可以快速检索的工具**。Bloom Filter 本身由**是一个长度为m的bit array，k个不相同的hash函数和源dataset组成**。具体的说，Bloom Filter是由k个不同的hash function将源dataset hash到m位的bit array构成。**通过Bloom Filter，我们可以快速检测出一个data是不是在源dataset中**（O(k) time）。

**Bloom Filter不保证完全的正确性**：如果被检索的data得到了false的反馈那他一定不在源data之中，true不一定在

在文件的起始位置，定义了两个常量BloomByteLength 和 BloomBitLength

```go
// core/types/bloom9.go”
const (
	BloomByteLength = 256
	BloomBitLength = 8 * BloomByteLength
)
```

在Ethereum中的Bloom Filter是一个长度为256的byte数组组成的。

```go
// Bloom represents a 2048 bit bloom filter
type Bloom [BloomByteLength]byte
```

Ethereum 中Bloom Filter使用的SHA Hash Function.

基本的思想是，使用三个value的值来判断log是否存在。
首先对data使用SHA function进行求值。选择hash后的
这三个value的选择[0,1],[2,3],[4,5]的值对2048取模得到目标位置的下标，并把这几个位置设为1.

对待判断的log进行相同的操作。
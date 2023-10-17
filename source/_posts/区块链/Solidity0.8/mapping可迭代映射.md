---
title: mapping可迭代映射

date: 2022-07-27	

categories: Solidity0.8.17	

tags: [区块链,Solidity0.8.17]
---	

# 学校和学生

我们想创建一个“学校”智能合约来收集学生地址。合约必须具有 3 个主要功能：

1. 从合同中添加或删除学生。
2. 询问给定的学生地址是否属于学校。
3. 获取所有学生的列表。

我们的`School`智能合约将如下所示：

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1lY-XnsfdjPztGwztIEBIvQ.png) 

## 简单解决方案 1：使用`mapping (address => bool)`

我们使用映射来存储每个学生的存在。如果映射到给定地址的值为`true`，则表示该地址是我们的学生之一。虽然解决方案很简单，但它的局限性在于它不能支持获取所有学生。与大多数其他语言不同，在 Solidity 中，不支持映射迭代。

## 简单解决方案 2：使用`address[] students`

我们使用地址数组而不是映射。现在很明显，我们解决了第三个需求（能够返回所有学生的列表）。但是，查找和删除现有学生变得更加困难。我们必须循环遍历数组中的每个元素以查找地址、检查地址是否存在或删除学生。

## 更好的解决方案：使用`mapping(address ⇒ address)`

激动人心的部分来了！这种数据结构的基础是[链表](https://en.wikipedia.org/wiki/Linked_list)。我们将下一个学生的地址（即指向下一个学生的指针）存储为映射值而不是普通布尔值。听起来很困惑吧？这张图会帮助你理解。

<img src="/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1ybjLvYv-CdGfOinFvFq4vA.png" alt="img" style="zoom:80%;" /> 

顶部：链表数据结构。每个节点指向它的下一个节点，最后一个节点指向 GUARD。底部：使用键值映射的顶部图像的具体表示。

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1eXK0rf5Ec4VQ8a_eSyHtnA.png) 

数据结构的初始化是通过将 GUARD 设置为指向 GUARD 来完成的，这意味着列表为空

现在让我们来看看每个功能的实现。

### 检查学生是否在学校：`isStudent`

`mapping`我们使用这样一个事实，即学校中特定学生的结构中的值始终指向下一个学生的地址。因此，我们可以通过检查给定地址映射到的值来轻松验证给定地址是否在学校内。如果它指向某个非零地址，则表示该学生地址在学校。

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1Leo2GwdsqhJBpnLxowweWw.png) 

### 向学校添加新学生：`addStudent`

我们可以在（表示列表的 HEAD 指针）之后添加一个新地址，`GUARD`方法是将守卫的指针更改为这个新地址，并将这个新地址（New Student）的指针设置为先前的前面地址（Front Student）。

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1h3xft5gEAseZGZEFdkbDHA.png) 

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1gd__pHCftvpWsmPbVCQ3Nw.png) 

### 从学校删除学生：`removeStudent`

这个函数比上面的两个函数更棘手。我们知道地址是否在列表中，但我们无法轻易推导出任何给定学生的先前地址（除非我们使用[双重链接列表](https://en.wikipedia.org/wiki/Doubly_linked_list)，但就存储成本而言，这要昂贵得多）。要删除一个地址，我们需要让它的前一个学生指向删除地址的下一个地址，并将删除地址的指针设置为零。

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1CAAcxCXqJ3zXdNfol-bBCg.png) 

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/15ziw2ePv1CTuqyp164ZYAQ.png) 

请注意，要实现`removeStudent`，我们还必须引入`getPrevStudent`有助于在任何给定学生之前找到以前学生地址的功能。

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/12oIfywFrb5VoQW1x2UD07w.png) 

### 获取所有学生的列表：`getStudents`

这很简单。我们从 GUARD 地址开始循环遍历映射，并将当前指针设置为下一个指针，直到它再次指向 GUARD，即迭代完成。

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1Z5KdTBzFq4NFwcq3cRs-Hg.png) 

### `removeStudent`进一步优化

请注意，`removeStudent`我们实现的函数消耗的 gas 与学校的学生人数成正比，因为我们需要遍历整个列表一次才能找到要删除的地址的前一个地址。我们可以通过使用链下计算将先前的地址发送给函数来优化此函数。因此，智能合约只需要验证之前的地址确实指向我们要删除的地址。

![img](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/Solidity0.8/img/1hFzsKxzbmwYiDWqM7USzNA.png) 

# 结论

在本文中，我们探索了*Iterable Maps*的一种实现，这是一种数据结构，它不仅支持**O(1) 的**添加、删除和查找，类似于传统的`mapping`，而且还支持集合迭代。
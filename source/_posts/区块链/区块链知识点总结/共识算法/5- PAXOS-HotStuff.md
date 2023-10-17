---
title: 5- PAXOS-HotStuff

date: 2023-02-04	

categories: 共识算法	

tags: [区块链,区块链知识点总结,共识算法]
---	

# PAXOS

Paxos算法运行在允许宕机故障的异步系统中，不要求可靠的消息传递，可容忍消息丢失、延迟、乱序以及重复。它利用大多数 (Majority) 机制**保证了2F+1的容错能力，即2F+1个节点的系统最多允许F个节点同时出现故障。**

一个或多个提议进程 (Proposer) 可以发起提案 (Proposal)，Paxos算法使所有提案中的某一个提案，在所有进程中达成一致。系统中的多数派同时认可该提案，即达成了一致。最多只针对一个确定的提案达成一致。

Paxos将系统中的角色分为**提议者 (Proposer)，决策者 (Acceptor)，和最终决策学习者 (Learner):**

- **Proposer**: 提出提案 (Proposal)。Proposal信息包括提案编号 (Proposal ID) 和提议的值 (Value)。
- **Acceptor**：参与决策，回应Proposers的提案。收到Proposal后可以接受提案，若Proposal获得多数Acceptors的接受，则称该Proposal被批准。
- **Learner**：不参与决策，从Proposers/Acceptors学习最新达成一致的提案（Value）。

在多副本状态机中，每个副本同时具有Proposer、Acceptor、Learner三种角色。

 ![image-20230321093843333](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/区块链/区块链知识点总结/共识算法/img/image-20230321093843333.png)

1. **第一阶段：Prepare阶段。Proposer向Acceptors发出Prepare请求，Acceptors针对收到的Prepare请求进行Promise承诺。**
2. **第二阶段：Accept阶段。Proposer收到多数Acceptors承诺的Promise后，向Acceptors发出Propose请求，Acceptors针对收到的Propose请求进行Accept处理。**
3. **第三阶段：Learn阶段。Proposer在收到多数Acceptors的Accept之后，标志着本次Accept成功，决议形成，将形成的决议发送给所有Learners。**

# HotStuff

是一种**优化后的三阶段bft算法**，在拜占庭节点数小于总数1/3时，保证系统的安全运行，同时提供更加高效的运行效率

以下是HotStuff算法的基本原理和过程：

1. 角色定义：
   - Leader（领导者）：负责提出新区块的候选人，并驱动共识过程。
   - Validator（验证者）：参与共识的节点，验证和存储区块链数据。
2. 阶段一：投票阶段（Voting Phase）：
   - **领导者提出一个新的候选区块，并将其广播给所有验证者。**
   - **验证者收到候选区块后，对其进行投票**，表示接受或拒绝该候选区块。
   - **当一个验证者收到多数其他验证者的接受票时，它会将自己的投票广播给其他节点。**
3. 阶段二：证明阶段（Certify Phase）：
   - **当一个验证者收到多数其他验证者的接受票时，它可以将该候选区块标记为已被证明。**
   - **该验证者会将证明消息广播给其他节点，以通知它们该候选区块已被共识接受。**
4. 阶段三：提交阶段（Commit Phase）：
   - **一旦一个验证者收到多数其他验证者的证明消息，它可以将该候选区块提交到区块链中。**
   - **验证者将该候选区块添加到自己的本地区块链中，并广播提交消息给其他节点。**

HotStuff算法具有以下特点和优势：

- 高性能：HotStuff采用了基于投票和多数决策的方式，使得共识过程具有较低的通信和计算开销。
- 简化流程：相比于传统的拜占庭容错共识算法，HotStuff的共识流程更为简化，减少了复杂性。
- 快速最终性：一旦一个候选区块被多数验证者接受，它就会被最终性地提交到区块链中，不再受到更长链的覆盖。
- 安全性：HotStuff算法基于拜占庭容错模型，可以容忍一部分节点的恶意行为或故障，保证共识结果的正确性和安全性。
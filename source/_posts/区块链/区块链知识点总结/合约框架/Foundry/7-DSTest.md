---
title: 7-DSTest

date: 2023-02-23	

categories: Foundry	

tags: [区块链,区块链知识点总结,合约框架,Foundry]
---	

# DSTest Reference

Dappsys Test（简称DSTest）提供基本的日志和断言功能

## 日志记录

这是对所有可用的日志事件的完整概述。关于详细的描述和使用实例，见下文。

```solidity
event log                    (string);
event logs                   (bytes);

event log_address            (address);
event log_bytes32            (bytes32);
event log_int                (int);
event log_uint               (uint);
event log_bytes              (bytes);
event log_string             (string);

event log_named_address      (string key, address val);
event log_named_bytes32      (string key, bytes32 val);
event log_named_decimal_int  (string key, int val, uint decimals);
event log_named_decimal_uint (string key, uint val, uint decimals);
event log_named_int          (string key, int val);
event log_named_uint         (string key, uint val);
event log_named_bytes        (string key, bytes val);
event log_named_string       (string key, string val);
```

## 断言

这是对所有可用断言功能的完整概述。关于详细的描述和使用实例，见下文。

```solidity
// 断言 `条件` 为真
function assertTrue(bool condition) internal;
function assertTrue(bool condition, string memory err) internal;

// 断言 `a` 等于 `b`
function assertEq(address a, address b) internal;
function assertEq(address a, address b, string memory err) internal;
function assertEq(bytes32 a, bytes32 b) internal;
function assertEq(bytes32 a, bytes32 b, string memory err) internal;
function assertEq(int a, int b) internal;
function assertEq(int a, int b, string memory err) internal;
function assertEq(uint a, uint b) internal;
function assertEq(uint a, uint b, string memory err) internal;
function assertEqDecimal(int a, int b, uint decimals) internal;
function assertEqDecimal(int a, int b, uint decimals, string memory err) internal;
function assertEqDecimal(uint a, uint b, uint decimals) internal;
function assertEqDecimal(uint a, uint b, uint decimals, string memory err) internal;
function assertEq(string memory a, string memory b) internal;
function assertEq(string memory a, string memory b, string memory err) internal;
function assertEq32(bytes32 a, bytes32 b) internal;
function assertEq32(bytes32 a, bytes32 b, string memory err) internal;
function assertEq0(bytes memory a, bytes memory b) internal;
function assertEq0(bytes memory a, bytes memory b, string memory err) internal;

// 断言  `a` 大于 `b`
function assertGt(uint a, uint b) internal;
function assertGt(uint a, uint b, string memory err) internal;
function assertGt(int a, int b) internal;
function assertGt(int a, int b, string memory err) internal;
function assertGtDecimal(int a, int b, uint decimals) internal;
function assertGtDecimal(int a, int b, uint decimals, string memory err) internal;
function assertGtDecimal(uint a, uint b, uint decimals) internal;
function assertGtDecimal(uint a, uint b, uint decimals, string memory err) internal;

// 断言  `a` 大于等于 `b`
function assertGe(uint a, uint b) internal;
function assertGe(uint a, uint b, string memory err) internal;
function assertGe(int a, int b) internal;
function assertGe(int a, int b, string memory err) internal;
function assertGeDecimal(int a, int b, uint decimals) internal;
function assertGeDecimal(int a, int b, uint decimals, string memory err) internal;
function assertGeDecimal(uint a, uint b, uint decimals) internal;
function assertGeDecimal(uint a, uint b, uint decimals, string memory err) internal;

// 断言  `a` 小于 `b`
function assertLt(uint a, uint b) internal;
function assertLt(uint a, uint b, string memory err) internal;
function assertLt(int a, int b) internal;
function assertLt(int a, int b, string memory err) internal;
function assertLtDecimal(int a, int b, uint decimals) internal;
function assertLtDecimal(int a, int b, uint decimals, string memory err) internal;
function assertLtDecimal(uint a, uint b, uint decimals) internal;
function assertLtDecimal(uint a, uint b, uint decimals, string memory err) internal;

// 断言  `a` 小于等于 `b`
function assertLe(uint a, uint b) internal;
function assertLe(uint a, uint b, string memory err) internal;
function assertLe(int a, int b) internal;
function assertLe(int a, int b, string memory err) internal;
function assertLeDecimal(int a, int b, uint decimals) internal;
function assertLeDecimal(int a, int b, uint decimals, string memory err) internal;
function assertLeDecimal(uint a, uint b, uint decimals) internal;
function assertLeDecimal(uint a, uint b, uint decimals, string memory err) internal;
```


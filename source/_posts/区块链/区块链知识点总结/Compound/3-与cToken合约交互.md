---
title: 3-与cToken合约交互

date: 2022-12-31	

categories: Compound	

tags: [区块链,区块链知识点总结,Compound]
---	

## 介绍

Compound 协议所支持的每项资产都通过 **cToken 合约**进行整合，cToken 合约是提供给协议的余额的 [EIP-20](https://eips.ethereum.org/EIPS/eip-20) 合格表示。通过铸造 cToken，用户（1）通过 cToken 的汇率赚取利息，该汇率相对于标的资产的价值增加，以及（2）获得试用 **cToken 作为抵押品的能力**。

cToken 是与 Compound 协议交互的主要手段；当用户在使用 cToken 合约时，会**使用 cToken  合约来进行铸币(mint)、赎回(redeem)、借款(borrow)、偿还借款(repay)、清算借款(liquidate)或者划转  cToken (transfer)**。

目前有两种类型的 cToken ： CErc20 和 CEther。尽管两者都暴露了 EIP-20 接口，但 CErc20 封装了  ERC-20 标的资产，而 CEther 只是封装了 Ether  本身。因此，根据类型的不同，涉及将资产传输到协议中的核心功能的接口也略有不同，每种类型的接口如下所示。

## 铸币(Mint)

mint 方法将一个资产转移到协议中，协议开始根据该资产的当前[供应率](https://learnblockchain.cn/article/3168)开始累积利息。用户收到的 cToken 数量等于标的代币数量除以当前[汇率](https://learnblockchain.cn/article/3168)。

**CErc20**

```solidity
1 function mint(uint mintAmount) returns (uint)
```

- msg.sender : 供应资产并拥有已铸 cToken 的账户。
- mintAmount : 供应的资产金额，单位为标的资产的单位。
- RETURN : 成功时为0，否则是[错误代码](https://learnblockchain.cn/article/3168)。

供应资产前，用户必须先 [批准(approve)](https://eips.ethereum.org/EIPS/eip-20#approve)， cToken 才能访问其代币余额。

**CEther**

```javascript
1 function mint() payable
```

- msg.value :  供应的 ETH 数量，单位 wei。
- msg.sender : 供应 ETH 并拥有已铸 cToken 的账户。
- RETURN : 无返回值，错误时还原。

**Solidity**

```solidity
1 Erc20 underlying = Erc20(0xToken...);     // 获取标的资产合约句柄 
2 CErc20 cToken = CErc20(0x3FDA...);        // 获取相应 cToken 合约的句柄 
3 underlying.approve(address(cToken), 100); // 批准划转 
4 assert(cToken.mint(100) == 0);            // 铸 cToken 币，并断言是否有错误
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...); 
2 await cToken.methods.mint().send({from: myAccount, value: 50});
```

## 赎回(Redeem)

redeem 方法将指定数量的 cToken 转换为标的资产，并将其返还给用户。收到的标的数量等于赎回的 cToken 数量乘以当前[汇率](https://learnblockchain.cn/article/3168)。赎回额必须小于用户的[账户流动性](https://compound.finance/docs/comptroller#account-liquidity)和市场可用的流动性。

**CErc20 / CEther**

```solidity
function redeem(uint redeemTokens) returns (uint)
```

- msg.sender : 应将赎回的标的资产转入的账户。
- redeemTokens : 将被赎回的 cToken 数量。
- RETURN : 成功时为0，否则是[错误代码](https://learnblockchain.cn/article/3168)。

**Solidity**

```solidity
1 CEther cToken = CEther(0x3FDB...);
2 require(cToken.redeem(7) == 0, "something went wrong");
```

## 赎回标的(Redeem Underlying)

redeem underlying 方法将 cToken兑换成指定数量的标的资产，并返回给用户。赎回的 cToken的数量等于收到的标的数量除以当前[汇率](https://learnblockchain.cn/article/3168)。赎回额必须小于用户的[账户流动性](https://compound.finance/docs/comptroller#account-liquidity)和市场可用的流动性。

**CErc20 / CEther**

```solidity
1 function redeemUnderlying(uint redeemAmount) returns (uint)
```

- msg.sender : 应将赎回的标的资产转入的账户。
- redeemAmount : 将被赎回的标的资产数量。
- RETURN : 成功时为0，否则是[错误代码](https://learnblockchain.cn/article/3168)。

**Solidity**

```solidity
1 CEther cToken = CEther(0x3FDB...);
2 require(cToken.redeemUnderlying(50) == 0, "something went wrong");
```

**Web3 1.0**

```solidity
1 const cToken = CErc20.at(0x3FDA...);
2 cToken.methods.redeemUnderlying(10).send({from: ...});
```

## 借款(Borrow)

borrow 方法**将协议中的资产转移给用户**，并创建一个借款余额，根据该资产的[借款利率](https://learnblockchain.cn/article/3168)开始累积利息。借款额必须小于用户的[账户流动性](https://compound.finance/docs/comptroller#account-liquidity)和市场可用的流动性。

要想借入以太，借款人必须是 'payable' （Solidity）。

**CErc20 / CEther**

```solidity
1 function borrow(uint borrowAmount) returns (uint)
```

- msg.sender : 应将借款资金转入的账户。
- borrowAmount :  标的资产借款数量。
- ETURN : 成功时为0，否则是[错误代码](https://learnblockchain.cn/article/3168)。

**Solidity**

```solidity
1 CErc20 cToken = CErc20(0x3FDA...);
2 require(cToken.borrow(100) == 0, "got collateral?");
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 await cToken.methods.borrow(50).send({from: 0xMyAccount});
```

## 偿还借款(Repay Borrow)

repay 方法将资产转移到协议中，并减少用户的借款余额。

**CErc20**

```solidity
1 function repayBorrow(uint repayAmount) returns (uint)
```

- msg.sender : 借入资产的账户，应偿还借款。
- repayAmount : 拟偿还的标的资产借款金额。值为-1（即2256-1）表示偿还全部借款额。
- RETURN : 成功时为0，否则是[错误代码](https://learnblockchain.cn/article/3168)。

偿还资产前，用户必须先 [批准(approve)](https://eips.ethereum.org/EIPS/eip-20#approve)， cToken 才能访问其代币余额。

**CEther**

```javascript
1 function repayBorrow() payable
```

- msg.value : 用于偿还的 ETH 数量，单位 wei。
- msg.sender : 借入资产的账户，应偿还借款。
- RETURN : 无返回值，错误时还原。

**Solidity**

```solidity
1 CEther cToken = CEther(0x3FDB...);

2 require(cToken.repayBorrow.value(100)() == 0, "transfer approved?");
```

**Web3 1.0**

```solidity
1 const cToken = CErc20.at(0x3FDA...);

2 cToken.methods.repayBorrow(10000).send({from: ...});
```

## 代偿还借款(Repay Borrow Behalf)

repay 方法将资产转移到协议中，并减少目标用户的借款余额。

**CErc20**

```solidity
1 function repayBorrowBehalf(address borrower, uint repayAmount) returns (uint)
```

- msg.sender : 偿还借款的账户。
- borrower : 借入资产要偿还的账户。
- repayAmount : 拟偿还的标的资产借款金额。值为-1（即2256-1）表示偿还全部借款额。
- RETURN : 成功时为0，否则是[错误代码](https://learnblockchain.cn/article/3168)。

偿还资产前，用户必须先 [批准(approve)](https://eips.ethereum.org/EIPS/eip-20#approve)， cToken 才能访问其代币余额。

**CEther**

```solidity
1 function repayBorrowBehalf(address borrower) payable
```

- msg.value : 用于偿还的 ETH 数量，单位 wei。
- msg.sender : 偿还借款的账户。
- borrower : 借入资产要偿还的账户。
- RETURN : 无返回值，错误时还原。

**Solidity**

```solidity
1 CEther cToken = CEther(0x3FDB...);
2 require(cToken.repayBorrowBehalf.value(100)(0xBorrower) == 0, "transfer approved?");
```

**Web3 1.0**

```solidity
1 const cToken = CErc20.at(0x3FDA...);
2 await cToken.methods.repayBorrowBehalf(0xBorrower, 10000).send({from: 0xPayer});
```

## 清算借款(Liquidate Borrow)

[账户流动性](https://compound.finance/docs/comptroller#account-liquidity)为负值的用户，由协议的其他用户进行[清算](https://learnblockchain.cn/article/3168)，使其账户流动性恢复到正值（即高于抵押品要求）。当发生清算时，清算人(liquidator)可以代表借款人偿还部分或全部未偿还的借款，作为回报，可以获得借款人持有的抵押品的折价；这种折价被定义为清算奖励。

清算人可将水下账户的任何一个未偿还借款以一定的固定比例（即结算系数）进行结算。与 V1 中不同的是，清算者必须与每一个 cToken  合约中偿还借款并扣押另一项资产作为抵押品的 cToken 进行交互。当抵押品被扣押时，清算人会被转移 cToken，他们可以赎回这些  cToken，就像他们自己提供资产一样。用户必须在调用清算前批准每个 cToken  合约（即在他们要偿还的借款资产上），因为他们正在将资金转移到合约中。

**CErc20**

```solidity
1 function liquidateBorrow(address borrower, uint amount, address collateral) returns (uint)
```

- msg.sender : 应通过偿还借款人的债务和扣押抵押品来结算借款人的账户。
- borrower : 应予清算的负[账户流动性](https://compound.finance/docs/comptroller#account-liquidity)账户。
- repayAmount : 借款资产的偿还和转换为抵押品的金额，以标的抵押资产的单位为单位。
- cTokenCollateral : 清算人应扣押借款人目前作为抵押品持有的 cToken 地址。
- RETURN : 成功时为0，否则是[错误代码](https://learnblockchain.cn/article/3168)。

偿还资产前，用户必须先 [批准(approve)](https://eips.ethereum.org/EIPS/eip-20#approve)， cToken 才能访问其代币余额。

**CEther**

```solidity
1 function liquidateBorrow(address borrower, address cTokenCollateral) payable
```

- msg.value : 转换为抵押品的 ETH 金额，单位 wei 。
- msg.sender : 应通过偿还借款人的债务和扣押抵押品来结算借款人的账户。
- borrower : 应予清算的负[账户流动性](https://compound.finance/docs/comptroller#account-liquidity)账户。
- cTokenCollateral : 清算人应扣押借款人目前作为抵押品持有的 cToken 地址。
- RETURN : 成功时为0，否则是[错误代码](https://learnblockchain.cn/article/3168)。

**Solidity**

```solidity
1 CEther cToken = CEther(0x3FDB...); 
2 CErc20 cTokenCollateral = CErc20(0x3FDA...);
3  require(cToken.liquidateBorrow.value(100)(0xBorrower, cTokenCollateral) == 0, "borrower underwater??");
```

**Web3 1.0**

```solidity
1 const cToken = CErc20.at(0x3FDA...); 
2 const cTokenCollateral = CEther.at(0x3FDB...);
3  await cToken.methods.liquidateBorrow(0xBorrower, 33, cTokenCollateral).send({from: 0xLiquidator});
```

## 主要事件(Key Events)

| Event                                                        | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| Mint(address minter, uint mintAmount, uint mintTokens)       | [铸币](https://compound.devdapp.cn/kai-fa-zhe-wen-dang/ctoken#zhu-bi-mint)成功后发出 |
| Redeem(address redeemer, uint redeemAmount, uint redeemTokens) | [赎回](https://learnblockchain.cn/article/3168)成功后发出    |
| Borrow(address borrower, uint borrowAmount, uint accountBorrows, uint totalBorrows) | [借款](https://compound.devdapp.cn/kai-fa-zhe-wen-dang/ctoken#jie-kuan-borrow)成功后发出 |
| RepayBorrow(address payer, address borrower, uint repayAmount, uint accountBorrows, uint totalBorrows) | [偿还借款](https://learnblockchain.cn/article/3168)成功后发出 |
| LiquidateBorrow(address liquidator, address borrower, uint repayAmount, address cTokenCollateral, uint seizeTokens) | [借款清算](https://learnblockchain.cn/article/3168)成功后发出 |

## 错误代码(Error Codes)

| Code | Name                           | Description                                                  |
| ---- | ------------------------------ | ------------------------------------------------------------ |
| 0    | NO_ERROR                       | 不是失败。                                                   |
| 1    | UNAUTHORIZED                   | 发件人无权实施这一操作。                                     |
| 2    | BAD_INPUT                      | 调用者提供了一个无效参数。                                   |
| 3    | COMPTROLLER_REJECTION          | 该操作将违反审计官政策。                                     |
| 4    | COMPTROLLER_CALCULATION_ERROR  | 审计内部计算失败 。                                          |
| 5    | INTEREST_RATE_MODEL_ERROR      | 利率模型返回了一个无效值。                                   |
| 6    | INVALID_ACCOUNT_PAIR           | 指定的账户组合是无效的。                                     |
| 7    | INVALID_CLOSE_AMOUNT_REQUESTED | 清算金额无效。                                               |
| 8    | INVALID_COLLATERAL_FACTOR      | 抵押因子无效。                                               |
| 9    | MATH_ERROR                     | 出现了数学计算错误。                                         |
| 10   | MARKET_NOT_FRESH               | 利息没有正确产生。                                           |
| 11   | MARKET_NOT_LISTED              | 目前，市场没有由审计员列出。                                 |
| 12   | TOKEN_INSUFFICIENT_ALLOWANCE   | ERC-20 合约必须**允许(allow)**货币市场合约调用`transferForm`。当前 allow 额为0或小于请求的供给、偿还借款或清算金额。 |
| 13   | TOKEN_INSUFFICIENT_BALANCE     | 调用者在 ERC-20 合约中没有足够的余额来完成所需的操作。       |
| 14   | TOKEN_INSUFFICIENT_CASH        | 市场上没有足够的现金金额来完成交易。您可以稍后再尝试此项交易。 |
| 15   | TOKEN_TRANSFER_IN_FAILED       | 在 ERC-20 代币转入市场时失败。                               |
| 16   | TOKEN_TRANSFER_OUT_FAILED      | 在 ERC-20 代币转出市场时失败。                               |

## 失败信息(Failure Info)

| Code | Name                                                       |
| ---- | ---------------------------------------------------------- |
| 0    | ACCEPT_ADMIN_PENDING_ADMIN_CHECK                           |
| 1    | ACCRUE_INTEREST_ACCUMULATED_INTEREST_CALCULATION_FAILED    |
| 2    | ACCRUE_INTEREST_BORROW_RATE_CALCULATION_FAILED             |
| 3    | ACCRUE_INTEREST_NEW_BORROW_INDEX_CALCULATION_FAILED        |
| 4    | ACCRUE_INTEREST_NEW_TOTAL_BORROWS_CALCULATION_FAILED       |
| 5    | ACCRUE_INTEREST_NEW_TOTAL_RESERVES_CALCULATION_FAILED      |
| 6    | ACCRUE_INTEREST_SIMPLE_INTEREST_FACTOR_CALCULATION_FAILED  |
| 7    | BORROW_ACCUMULATED_BALANCE_CALCULATION_FAILED              |
| 8    | BORROW_ACCRUE_INTEREST_FAILED                              |
| 9    | BORROW_CASH_NOT_AVAILABLE                                  |
| 10   | BORROW_FRESHNESS_CHECK                                     |
| 11   | BORROW_NEW_TOTAL_BALANCE_CALCULATION_FAILED                |
| 12   | BORROW_NEW_ACCOUNT_BORROW_BALANCE_CALCULATION_FAILED       |
| 13   | BORROW_MARKET_NOT_LISTED                                   |
| 14   | BORROW_COMPTROLLER_REJECTION                               |
| 15   | LIQUIDATE_ACCRUE_BORROW_INTEREST_FAILED                    |
| 16   | LIQUIDATE_ACCRUE_COLLATERAL_INTEREST_FAILED                |
| 17   | LIQUIDATE_COLLATERAL_FRESHNESS_CHECK                       |
| 18   | LIQUIDATE_COMPTROLLER_REJECTION                            |
| 19   | LIQUIDATE_COMPTROLLER_CALCULATE_AMOUNT_SEIZE_FAILED        |
| 20   | LIQUIDATE_CLOSE_AMOUNT_IS_UINT_MAX                         |
| 21   | LIQUIDATE_CLOSE_AMOUNT_IS_ZERO                             |
| 22   | LIQUIDATE_FRESHNESS_CHECK                                  |
| 23   | LIQUIDATE_LIQUIDATOR_IS_BORROWER                           |
| 24   | LIQUIDATE_REPAY_BORROW_FRESH_FAILED                        |
| 25   | LIQUIDATE_SEIZE_BALANCE_INCREMENT_FAILED                   |
| 26   | LIQUIDATE_SEIZE_BALANCE_DECREMENT_FAILED                   |
| 27   | LIQUIDATE_SEIZE_COMPTROLLER_REJECTION                      |
| 28   | LIQUIDATE_SEIZE_LIQUIDATOR_IS_BORROWER                     |
| 29   | LIQUIDATE_SEIZE_TOO_MUCH                                   |
| 30   | MINT_ACCRUE_INTEREST_FAILED                                |
| 31   | MINT_COMPTROLLER_REJECTION                                 |
| 32   | MINT_EXCHANGE_CALCULATION_FAILED                           |
| 33   | MINT_EXCHANGE_RATE_READ_FAILED                             |
| 34   | MINT_FRESHNESS_CHECK                                       |
| 35   | MINT_NEW_ACCOUNT_BALANCE_CALCULATION_FAILED                |
| 36   | MINT_NEW_TOTAL_SUPPLY_CALCULATION_FAILED                   |
| 37   | MINT_TRANSFER_IN_FAILED                                    |
| 38   | MINT_TRANSFER_IN_NOT_POSSIBLE                              |
| 39   | REDEEM_ACCRUE_INTEREST_FAILED                              |
| 40   | REDEEM_COMPTROLLER_REJECTION                               |
| 41   | REDEEM_EXCHANGE_TOKENS_CALCULATION_FAILED                  |
| 42   | REDEEM_EXCHANGE_AMOUNT_CALCULATION_FAILED                  |
| 43   | REDEEM_EXCHANGE_RATE_READ_FAILED                           |
| 44   | REDEEM_FRESHNESS_CHECK                                     |
| 45   | REDEEM_NEW_ACCOUNT_BALANCE_CALCULATION_FAILED              |
| 46   | REDEEM_NEW_TOTAL_SUPPLY_CALCULATION_FAILED                 |
| 47   | REDEEM_TRANSFER_OUT_NOT_POSSIBLE                           |
| 48   | REDUCE_RESERVES_ACCRUE_INTEREST_FAILED                     |
| 49   | REDUCE_RESERVES_ADMIN_CHECK                                |
| 50   | REDUCE_RESERVES_CASH_NOT_AVAILABLE                         |
| 51   | REDUCE_RESERVES_FRESH_CHECK                                |
| 52   | REDUCE_RESERVES_VALIDATION                                 |
| 53   | REPAY_BEHALF_ACCRUE_INTEREST_FAILED                        |
| 54   | REPAY_BORROW_ACCRUE_INTEREST_FAILED                        |
| 55   | REPAY_BORROW_ACCUMULATED_BALANCE_CALCULATION_FAILED        |
| 56   | REPAY_BORROW_COMPTROLLER_REJECTION                         |
| 57   | REPAY_BORROW_FRESHNESS_CHECK                               |
| 58   | REPAY_BORROW_NEW_ACCOUNT_BORROW_BALANCE_CALCULATION_FAILED |
| 59   | REPAY_BORROW_NEW_TOTAL_BALANCE_CALCULATION_FAILED          |
| 60   | REPAY_BORROW_TRANSFER_IN_NOT_POSSIBLE                      |
| 61   | SET_COLLATERAL_FACTOR_OWNER_CHECK                          |
| 62   | SET_COLLATERAL_FACTOR_VALIDATION                           |
| 63   | SET_COMPTROLLER_OWNER_CHECK                                |
| 64   | SET_INTEREST_RATE_MODEL_ACCRUE_INTEREST_FAILED             |
| 65   | SET_INTEREST_RATE_MODEL_FRESH_CHECK                        |
| 66   | SET_INTEREST_RATE_MODEL_OWNER_CHECK                        |
| 67   | SET_MAX_ASSETS_OWNER_CHECK                                 |
| 68   | SET_ORACLE_MARKET_NOT_LISTED                               |
| 69   | SET_PENDING_ADMIN_OWNER_CHECK                              |
| 70   | SET_RESERVE_FACTOR_ACCRUE_INTEREST_FAILED                  |
| 71   | SET_RESERVE_FACTOR_ADMIN_CHECK                             |
| 72   | SET_RESERVE_FACTOR_FRESH_CHECK                             |
| 73   | SET_RESERVE_FACTOR_BOUNDS_CHECK                            |
| 74   | TRANSFER_COMPTROLLER_REJECTION                             |
| 75   | TRANSFER_NOT_ALLOWED                                       |
| 76   | TRANSFER_NOT_ENOUGH                                        |
| 77   | TRANSFER_TOO_MUCH                                          |

## 汇率(Exchange Rate)

每个cToken 都可以兑换成数量不断增加的标的资产，因为在市场上会产生利息。cToken 与相关资产之间的汇率等于：

```undefined
1 exchangeRate = (getCash() + totalBorrows() - totalReserves()) / totalSupply()
```

**CErc20 / CEther**

```solidity
1 function exchangeRateCurrent() returns (uint)
```

- 返回值 : 当前的汇率，为无符号整数，按 1e18 缩放；

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2 uint exchangeRateMantissa = cToken.exchangeRateCurrent();
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...); 
2 const exchangeRate = (await cToken.methods.exchangeRateCurrent().call()) / 1e18;
```

> 请注意，对比 send 使用 `call` 是在链外调用方法，而不需要导致 Gas 成本。

## 获取现金(Get Cash)

现金是指该 cToken 合约所拥有的标的资产余额。人民可以查询到目前这个市场上的现金总量。

**CErc20 / CEther**

1 function getCash() returns (uint)

- 返回值 : 该 cToken 合约所拥有的标的资产数量。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2 uint cash = cToken.getCash();
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 const cash = (await cToken.methods.getCash().call());
```

## 总借款(Total Borrow)

总借贷额是指目前市场上借出的标的额，以及累计向市场上的供应商支付的利息。

**CErc20 / CEther**

```solidity
1 function totalBorrowsCurrent() returns (uint)
```

- 返回值：借出的标的总额与利息。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2 uint borrows = cToken.totalBorrowsCurrent();
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 const borrows = (await cToken.methods.totalBorrowsCurrent().call());
```

## 借款余额(Borrow Balance)

从协议中借入资产的用户要根据借款利率进行累计利息。利息是每块都要累计的，集成商可以通过这个功能来获得用户的借款余额的当前价值与利息。

**CErc20 / CEther**

```solidity
1 function borrowBalanceCurrent(address account) returns (uint)
```

- account : 借入资产的账户。
- 返回值 : 用户的当前借款余额（与利息），单位是标的资产的单位。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2 uint borrows = cToken.borrowBalanceCurrent(msg.caller);
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 const borrows = await cToken.methods.borrowBalanceCurrent(account).call();
```

## 借款利率(Borrow Rate)

在任何时候，任何人都可以通过查询合约来获取当前每个区块的借款利率。

**CErc20 / CEther**

```solidity
1 function borrowRatePerBlock() returns (uint)
```

- 返回值 : 当前借款利率，为一个无符号整数，按 1e18 缩放。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2 uint borrowRateMantissa = cToken.borrowRatePerBlock();
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 const borrowRate = (await cToken.methods.borrowRatePerBlock().call()) / 1e18;
```

## 总供应量(Total Supply)

总供应量是指目前在 cToken 市场上流通的代币数量。它是 cToken 合约的 EIP-20 接口的一部分 。

**CErc20 / CEther**

```solidity
1 function totalSupply() returns (uint)
```

- 返回值 : 在市场上流通的代币总数量。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2  uint tokens = cToken.totalSupply();
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 const tokens = (await cToken.methods.totalSupply().call());
```

## 标的余额 Balance

用户的标的月，代币他们在协议中的资产，等于用户的 cToken 余额乘以[汇率](https://learnblockchain.cn/article/3168)。

**CErc20 / CEther**

```solidity
1 function balanceOfUnderlying(address account) returns (uint)
```

- account : 获取标的余额的账户。
- 返回值 : 该账户当前拥有的标的数量。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2 uint tokens = cToken.balanceOfUnderlying(msg.caller);
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 const tokens = await cToken.methods.balanceOfUnderlying(account).call();
```

## 供给率(Supply Rate)

在任何时候，人们都可以通过查询合约来获取当前每块的供给率。供给率是由[借款利率](https://learnblockchain.cn/article/3168)、准备金系数([reserve factor](https://compound.finance/docs/ctokens#reserve-factor))和[总借款额](https://learnblockchain.cn/article/3168)得出的。

**CErc20 / CEther**

```solidity
1 function supplyRatePerBlock() returns (uint)
```

- 返回值 : 当前供给率，为一个无符号整数，按 1e18 缩放。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2 uint supplyRateMantissa = cToken.supplyRatePerBlock();
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 const supplyRate = (await cToken.methods.supplyRatePerBlock().call()) / 1e18;
```

## 总储备金(Total Reserves)

储备金是每个 cToken 合约中的会计分录，代币历史利息的一部分，作为现金预留，可以通过协议的治理来提取或转移。借款人利息的一小部分应计入协议中，由准备金系数 [reserve factor](https://compound.finance/docs/ctokens#reserve-factor)决定。

**CErc20 / CEther**

```solidity
1 function totalReserves() returns (uint)
```

- 返回值 : 市场上持有的储备金总额。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2 uint reserves = cToken.totalReserves();
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...);
2 const reserves = (await cToken.methods.totalReserves().call());
```

## 储备金系数(Reserve Factor)

准备金系数确定了借款人利息中转化为准备金的部分。

**CErc20 / CEther**

```solidity
1 function reserveFactorMantissa() returns (uint)
```

- 返回值 : 当前市场准备金系数，为一个无符号整数，按 1e18 缩放。

**Solidity**

```solidity
1 CErc20 cToken = CToken(0x3FDA...);
2uint reserveFactorMantissa = cToken.reserveFactorMantissa();
```

**Web3 1.0**

```solidity
1 const cToken = CEther.at(0x3FDB...); 
2 const reserveFactor = (await cToken.methods.reserveFactorMantissa().call()) / 1e18;
```


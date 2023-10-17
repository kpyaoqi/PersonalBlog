---
title: UniswapV2Pair

date: 2023-01-17	

categories: UniswapV2	

tags: [区块链,区块链知识点总结,UniswapV2]
---	

# UniswapV2Pair

在构造函数中factory地址为msg.sender，因为该合约是由UniswapV2Factory进行部署，该合约继承了UniswapV2ERC20，主要用于管理以及操作交易对，托管两种token

## 合约当中比较重要的方法有(lptoken即UniswapV2ERC20):

  function permit(address owner, address spender, uint value, uint deadline, uint8 v, bytes32 r, bytes32 s) external：判断签名的有效性

  function mint(address to) external returns (uint liquidity)：铸造lptoken

  function burn(address to) external returns (uint amount0, uint amount1)：销毁lptoken退出流动性

  function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external：根据tokenA的数量在交易池中进行兑换tokenB

  function skim(address to) external：使两个token的余额与储备相等

  function sync() external：使两个token的储备与余额相匹配

  function initialize(address, address) external：设置pair地址交易对的两种token

#### permit:(address owner, address spender, uint value, uint deadline, uint8 v, bytes32 r, bytes32 s)

```solidity
// ecrecover 函数可以返回与签名者对应的公钥地址
address recoveredAddress = ecrecover(digest, v, r, s);
// 判断签名者对应的公钥地址与授权地址是否一致
require(recoveredAddress != address(0) && recoveredAddress == owner, 'UniswapV2: INVALID_SIGNATURE');
_approve(owner, spender, value);
```

用户现在签名一笔交易，用该方法可有判断该签名的有效性,如果通过判断，则进行授权

#### mint：(address to) lock returns (uint liquidity)

```solidity
// 1.获取进行添加流动性的两个token的数量
(uint112 _reserve0, uint112 _reserve1, ) = getReserves();
uint balance0 = IERC20(token0).balanceOf(address(this));
uint balance1 = IERC20(token1).balanceOf(address(this));
uint amount0 = balance0.sub(_reserve0);
uint amount1 = balance1.sub(_reserve1);
// 2.调用_mintFee方法
// 3.添加流动性所获得的lptoken数量(进行添加流动性的两种token的数量*目前lptoken的数量/当前token的储备量-->取较小值)
liquidity = Math.min(amount0.mul(_totalSupply) / _reserve0, amount1.mul(_totalSupply) / _reserve1);
// 4.铸造lptoken函数和更新储备函数
_mint(to, liquidity);
_update(balance0, balance1, _reserve0, _reserve1);
```

主要是根据两个token在交易对的增量计算出应该铸造lptoken的数量，然后将lptoken铸造给to地址，具有防重入锁lock

#### burn：(address to) lock returns (uint amount0, uint amount1)

```solidity
// 1.为什么用addres(this)?-->因为获取退出lptoken数量时，是在Route合约中先将lptoken转到当前合约，然后直接获得当前合约lptoken的数量
uint liquidity = balanceOf[address(this)];
// 2.调用_mintFee方法
// 3.使用余额确保按比例分配-->(持有lptoken/总lptoken)*合约中持有token的数量
amount0 = liquidity.mul(balance0) / _totalSupply; 
amount1 = liquidity.mul(balance1) / _totalSupply; 
// 4.转账两种token并更新储备量
_safeTransfer(_token0, to, amount0);
_safeTransfer(_token1, to, amount1);
balance0 = IERC20(_token0).balanceOf(address(this));
balance1 = IERC20(_token1).balanceOf(address(this));
_update(balance0, balance1, _reserve0, _reserve1);
```

根据lptokne的比例计算出两种token的各自的数量，然后销毁lptoken并将转账两种token给to地址

#### swap：(uint amount0Out, uint amount1Out, address to, bytes calldata data)

```solidity
// 1.转移代币
if (amount0Out > 0) _safeTransfer(_token0, to, amount0Out); 
if (amount1Out > 0) _safeTransfer(_token1, to, amount1Out); 
// 2.用于回调合约来实现一些特定的业务逻辑或其他自定义功能(例如：闪电贷....)
if (data.length > 0) IUniswapV2Callee(to).uniswapV2Call(msg.sender, amount0Out, amount1Out, data);
// 3.确保在交易完成后，资金池的储备量满足 Uniswap V2 中的 K 恒定公式，即 K = _reserve0 * _reserve1
// 4.更新储备
```

1.在Route合约用户已经将需要兑换的tokenA转入pair合约中，在Route合约中传入需要输出的tokenB的数量和一个data，转移tokenB后判断data长度是否大于零去进行回调合约

2.直接调用swap方法进行回调合约获得套利，只要套利后满足后续条件即可

## 合约内部调用的方法：

#### _update：(uint balance0, uint balance1, uint112 _reserve0, uint112 _reserve1) private

```solidity
// 1.更新priceCumulativeLast，永远不会溢出，+ overflow是理想的
price0CumulativeLast += uint(UQ112x112.encode(_reserve1).uqdiv(_reserve0)) * timeElapsed;
price1CumulativeLast += uint(UQ112x112.encode(_reserve0).uqdiv(_reserve1)) * timeElapsed;
// 2.更新储备量
reserve0 = uint112(balance0);
reserve1 = uint112(balance1);
```

更新储备方法：四个参数前两个为更新后两个token的储备量，后两个为更新前两个token的储备量

#### _mintFee：(uint112 _reserve0, uint112 _reserve1) private returns (bool feeOn)

```solidity
// 1.获取收取手续费的地址如果不是零地址并且kLast!=0则继续下面部分
// 2.获取上一次交易后和目前交易对中的K值
uint rootK = Math.sqrt(uint(_reserve0).mul(_reserve1));
uint rootKLast = Math.sqrt(_kLast);
// 3.如果rootK>rootKLast
uint numerator = totalSupply.mul(rootK.sub(rootKLast));
uint denominator = rootK.mul(5).add(rootKLast);
uint liquidity = numerator / denominator;
// 4.如果liquidity大于零为收取手续费地址铸造lptoken 
if (liquidity > 0) _mint(feeTo, liquidity);
```

收取手续费方法，参数为当前两个token的储备量

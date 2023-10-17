---
title: 2-Web3

date: 2022-08-30	

categories: web3.js	

tags: [区块链,web3.js]
---	

## Web3

Web3是web.js的主类，可以通过`Web3.modules`返回所有子模块类的对象

**子模块列表**：

- Eth：与以太坊网络交互的 Eth 模块
- Net：与网络属性进行交互的 Net 模块
- Personal：与以太坊账户交互的 Personal 模块
- Shh：与whisper协议交互的 Shh 模块
- Bzz：与swarm网络交互的 Bzz 模块
- Utils：工具性函数

### 1.**version**

**（ 返回当前 web3.js 库的软件包版本）**

**返回值**(String)：当前版本

### 2.**setProvider**

**（改变相应模块的 provider）**

**参数**：Object(myProvider)：一个有效的provider

**返回值(String)**：Boolean

> 通过web3调用它的时候，web3.eth、web3.net等子模块的provider也会被设置，除了web3.bzz，它总需要单独设置

### 3.providers

**（包含当前可用的 providers）(子模块也有)**

**返回值(Object)：**

-  `HttpProvider`: 因为不能用于订阅，HTTP provider 已经**不推荐使用**。
-  `WebsocketProvider`: Websocket provider 是用于传统的浏览器中的标准方法.
- `IpcProvider`: 当运行一个本地节点时，IPC provider 用于 node.js 下的DApp 环境，提供最为安全的连接。

### 4.giveProvider

**(返回浏览器设置的原生 provider ，否则返回 null)(子模块也有)**

**返回值：**浏览器设置好的 provider 或者 `null`

### 5.currentProvider

**(当前在用的 provider)(子模块也有)**

**返回值**：当前在用的 provider 或者 `null`

### 6.BatchRequest

**（创建并执行批量请求的类）(子模块也有)**

**返回值(Object)：**具有如下方法

- `add(request)`: 添加请求对象到批量调用中。
- `execute()`: 执行批量请求。

```js
//例子
var contract = new web3.eth.Contract(abi, address);

var batch = new web3.BatchRequest();
batch.add(web3.eth.getBalance.request('0x0000000000000000000000000000000000000000', 'latest', callback));
batch.add(contract.methods.balance(address).call.request({from: '0x0000000000000000000000000000000000000000'}, callback2));
batch.execute();
```

### 7.extend

**（扩展 web3 模块）(子模块也有)**

**参数**：

- `methods` - `Object`: 扩展对象，带有一组如下所示的方法描述对象:
  - `property` - `String`: (可选) 要添加到模块上的属性名称。如果没有设置属性，则直接添加到模块上。
  - `methods` - `Array`: 方法描述对象数组：
    - `name` - `String`: 要添加的方法名称。
    - `call` - `String`: RPC 方法名称。
    - `params` - `Number`: (可选) 方法的参数个数，默认值为 0。
    - `inputFormatter` - `Array`: (可选) 输入格式化函数数组，每个成员对应一个函数参数，或者使用 null 来对应不需要进行格式化处理的参数。
    - `outputFormatter - ``Function`: (可选) 用来格式化方法输出。

**返回值(Object)**: 扩展模块.

```js
//l
web3.extend({
    property: 'myModule',
    methods: [{
        name: 'getBalance',
        call: 'eth_getBalance',
        params: 2,
        inputFormatter: [web3.extend.formatters.inputAddressFormatter, web3.extend.formatters.inputDefaultBlockNumberFormatter],
        outputFormatter: web3.utils.hexToNumberString
    },{
        name: 'getGasPriceSuperFunction',
        call: 'eth_gasPriceSuper',
        params: 2,
        inputFormatter: [null, web3.utils.numberToHex]
    }]
});

web3.extend({
    methods: [{
        name: 'directCall',
        call: 'eth_callForFun',
    }]
});

console.log(web3);
> Web3 {
    myModule: {
        getBalance: function(){},
        getGasPriceSuperFunction: function(){}
    },
    directCall: function(){},
    eth: Eth {...},
    bzz: Bzz {...},
    ...
}
```
---
title: 3-Providers

date: 2022-08-08	

categories: ether@5.4.js	

tags: [区块链,ether@5.4.js]
---	

# JsonRpcProvider

JSON-RPC API是与以太坊交互比较流行的方法，在主要的以太坊节点实现(如Geth 和 Parity)以及许多第三方web服务中都可以用。

**new *ethers*.*providers*.JsonRpcProvider( [ *urlOrConnectionInfo* [ ，*networkish* ] ] )**

使用 URL 或者 urlOrConnectionInfo 作为一个一个 JSON-RPC HTTP API 连接到 *networkish* 网络。

如果 *urlOrConnectionInfo* 没有被指定， 默认值是`http://localhost:8545`。 如果网络没有被指定， 它将会自动使用 `eth_chaindId`查询网络节点并返回 `eth_networkId`从而确定网络。

## 1.1 getSigner

*jsonRpcProvider*.**getSigner**( [ *addressOrIndex* ] ) ⇒ ***JsonRpcSigner***

返回由该以太坊节点管理的JsonRpcSigner，地址为*addressOrIndex*。如果没有提供*addressOrIndex*，则使用第一个帐户(account #0)。

## 1.2 listAccounts

*jsonRpcProvider*.**listAccounts**( ) ⇒ *Promise< Array< **string** > >*

返回此provider管理的所有帐户地址的列表。

## 1.3 send

*jsonRpcProvider*.**send**( *method*，*params* ) ⇒ *Promise< **any** >*

允许向provider发送原始消息,这可以用于后端特定的调用，比如调试或特定的帐户管理。

获取不常见的或特定于某些以太坊节点(例如Parity 与 Geth)才有的

## 1.4 **getUncheckedSigner**

*jsonRpcProvider*.**getUncheckedSigner**( [ *addressOrIndex* ] ) ⇒ ***JsonRpcUncheckedSigner***

**JsonRpcUncheckedSigner**:JSON-RPC API仅在发送交易时提供交易哈希作为响应，但ether Provider要求在返回交易之前填充交易的所有细节。 例如，燃料价格（gas price）和燃料限制（gas limit）可能由节点或自动包含在内的nonce调整。

为了解决这个问题，**JsonRpcSigner**立即查询provider的详细信息，用返回的交易哈希来填充**TransactionResponse**对象。

**UncheckedSigner**不填充任何附加信息，并将一个模拟的类似**TransactionResponse**的对象作为立即返回的结果， 其中大部分属性设置为null，但如果只需要这些，则可以快速获得交易哈希。

## 1.5 StaticJsonRpcProvider

ethers Provider将频繁执行`getNetwork`调用，以确保网络调用和网络通信是一致的(**MetaMask**)

在某些情况下网络不能改变，如当连接到一个INFURA端点，在这种情况下，可以使用**StaticJsonRpcProvider**来将一直缓存链ID，可以减少网络流量和往返查询chain ID的次数。

# JsonRpcSinger

**JsonRpcSigner**是一个简单的Signer，它由一个连接的JsonRpcProvider支持。

## 2.1 provider

*signer*.**provider** ⇒ ***JsonRpcProvider***

创建这个signer的provider

## 2.2 connectUnchecked

*signer*.**connectUnchecked**( ) ⇒ ***JsonRpcUncheckedSigner***

返回一个新的Signer对象，该对象在发送交易时不执行额外的检查。详情请参阅getUncheckedSigner。

## 2.3 sendUncheckedTransaction

*signer*.**sendUncheckedTransaction**( *transaction* ) ⇒ *Promise< string< **DataHexString**< 32 > > >*

发送*交易*并返回一个解析为不透明交易哈希的Promise。

## 2.4 unlock

*signer*.**unlock**( *password* ) ⇒ *Promise< **boolean** >*

使用*密码*请求节点解除锁定（如果被锁了）。

# API Proviers

很多服务提供了访问以太坊区块链的web API,这些提供商允许连接到它们，因为你不需要运行你自己的实例或以太坊节点集群，这就简化了开发。这种对第三方服务的依赖会降低弹性和安全性，并增加所需的信任度。为了缓解这些问题，建议您使用 **Default** **Provider**。

## 3.1 EtherscanProvider

**new** *ethers*.*providers*.**EtherscanProvider**( [ *network* = "*homestead*" , [ *apiKey* ] ] )

使用可选的*apiKey*创建一个新的**EtherscanProvider**连接到*网络*。

这个*网络*可以被指定为一个**字符串**类型的网络名称、或**number**类型的链ID，或[网络对象]provider-(network)。

如果没有提供*apiKey*,将使用一个共享的API key,可能导致性能降低和请求受限,推荐在生产环境中使用您在Etherscan上注册的API密钥。

```js
// 连接主网 (homestead)
provider = new EtherscanProvider();

// 连接 rinkeby 测试网(以下这两种方式是等价的)
provider = new EtherscanProvider("goerli");
provider = new EtherscanProvider(5);

provider = new EtherscanProvider(network);

// 通过API key 连接主网(homestead)
provider = new EtherscanProvider(null, apiKey);
provider = new EtherscanProvider("homestead", apiKey);
```

## 3.2 InfuraProvider

**new *ethers*.*providers*.InfuraProvider( [ *network* = "*homestead*" , [ *apiKey* ] ] )**

*apiKey*可以是一个**string**类型的Project ID，也可以是一个带有`projectId`和`projectSecret`属性**对象**， 用于指定一个可以在非公共源(如服务器)上使用的Project Secret， 以进一步保护你的API access 和 quotas。



***InfuraProvider*.getWebSocketProvider( [ *network* [ , *apiKey* ] ] ) ⇒ *WebSocketProvider***

Create a new WebSocketProvider 使用 INFURA web-socket 端点创建一个新的WebSocketProvider，并使用可选的*apiKey*连接到网络。 网络和apiKey的设定与构造函数相同。

```js
// 使用一个Project ID连接到主网(以下这两种方式是等价的)
provider = new InfuraProvider(null, projectId);
provider = new InfuraProvider("homestead", projectId);

// 使用Project ID 和 Project Secret 连接到主网
provider = new InfuraProvider("homestead", {
    projectId: projectId,
    projectSecret: projectSecret
});

// 使用一个 WebSocketProvider 连接到 INFURA WebSocket 端点
provider = InfuraProvider.getWebSocketProvider()
```

## 3.3 AlchemyProvider

## 3.4 CloundfalreProvier

# 其他的Providers

## 4.1 **FallbackProvider**

**FallbackProvider** 是ethers中可用的且最高级的Provider。

它使用一个quorum并连接到多个Providers作为后端，每个Providers都配置了一个*优先级*和*权重*。

当发出一个请求时，请求被分配给多个随机选择的后端(优先级较低的后端总是被先选中)， 并将每个请求的结果与其他请求进行比较。只有达到quorum规定的数量后，该结果才会被接受并返回给调用方。

**new** *ethers*.*providers*.**FallbackProvider**( *providers* [ , *quorum* ] )

创建一个连接到*providers*的FallbackProvider的新实例。 如果quorum未指定，则默认为provider权重总和的一半。

*providers*可以是**Provider** 或 **FallbackProviderConfig**的数组。 如果提供了Provider，默认的优先级为1，权重为1。

描述后端的Provider配置列表:**providerConfigs ⇒ *Array< FallbackProviderConfig >***

在得到结果之前，后端响应的quorum必须达成一致:***provider*.quorum ⇒ *number***

### FallbackProviderConfig

该配置的provider:***fallbackProviderConfig*.provider ⇒ *Provider***

表示provider使用的优先级:***fallbackProviderConfig*.priority ⇒ *number***

超时后(以ms为单位)将会尝试其他的Provider:***fallbackProviderConfig*.stallTimeout ⇒ *number***

表示这个provider响应的权重:**fallbackProviderConfig.weight ⇒ *number***

## 4.2 IpcProvider

**IpcProvider** 允许JSON-RPC API在文件系统的本地文件名上使用。 Geth, Parity和其他节点都会开放了这个功能。

这只能在*node.js*中使用，因为它需要访问文件系统，并且由于文件权限可能有增加额外的复杂性。(请参阅相关节点实现的文档)

这个Provider要连接的路径:**ipcProvider.path ⇒ *string***

## 4.3 **UrlJsonRpcProvider**

这个类打算作为子类而不是直接使用。只需要额外生成一个JSON-RPC URL后，就能通过JsonRpcProvider创建一个Provider。

**new** *ethers*.*providers*.**UrlJsonRpcProvider**( [ *network* [ , *apiKey* ] ] )

子类通常不需要重写这一部分。相反应该重写静态的方法 `getUrl` 和可选的`getApiKey`。

***urlJsonRpcProvider*.apiKey ⇒ *any***

从`InheritedClass.getApiKey`返回的apiKey的值。

***InheritingClass*.getApiKey( *apiKey* ) ⇒ *any***

这个函数应该检查*apiKey*以确保它是有效的，并返回一个(可能修改过的)值在`getUrl`函数中使用。

***InheritingClass*.getUrl( *network* , *apiKey* ) ⇒ *string***

这个URL在JsonRpcProvider 实例中使用。

## 4.4 **Web3Provider**

## 4.5 **ExternalProvider**

## 4.6 WebSocketProvider

**WebSocketProvider** 连接到一个JSON-RPC websocket兼容的后端，它允许持久连接、多路复用请求和发布-子事件，以实现更即时的事件调度。WebSocket API是较新的，如果运行自己的基础设施，请注意WebSockets对服务器资源的占用会很大， 因为它们必须管理和维护每个客户端的状态。由于这个原因，许多服务也可能会为使用他们的WebSocket端点而收取额外的费用。

**new *ethers*.*providers*.WebSocketProvider( [ *url* [ , *network* ] ] )**

通过一个*url*连接到*网络*，返回一个新的WebSocketProvider。

如果*url*未指定，默认值是`"ws://localhost:8546"`。 如果*network*未被指定，它将会从网络中查询。

---
title: 1-Geth启动命令

date: 2021-11-30	

categories: Geth	

tags: [Geth]
---	

# 命令

account                账户管理

attach                启动一个交互式JavaScript环境(连接到节点)

console                启动一个交互式JavaScript环境

db    数据库低级操作

dump    从存储中转储指定块

dumpconfig   以TOML格式导出配置值

dumpgenesi   s将创世块JSON配置转储到stdout

export                 导出区块链到文件

export-preimages       将预映像数据库导出为RLP流

import                 导入区块链文件

import-preimages       从RLP流导入预映像数据库

init                   Bootstrap并初始化一个新的创世块

js                     (DEPRECATED)执行指定的JavaScript文件

license	显示license信息

removedb	删除区块链和状态数据库

Show -deprecated-flags	显示已弃用的标志

snapshot	基于快照的命令集

verkle	一组实验性的verkle树管理命令

version	打印版本号

version-check	检查(在线)已知的Geth安全漏洞

wallet                 管理以太坊预售钱包

help, h	显示命令列表或命令的帮助信息

# 全局选项

## 账户

​    --allow-insecure-unlock        (default: false)
​          当http协议暴露了与帐号相关的rpc时，允许不安全的帐号解锁

​    --keystore value
​          密钥存储库的目录 (default = inside the datadir)

​    --lightkdf                     (default: false)
​          以牺牲KDF强度为代价减少密钥派生RAM和CPU的使用

​    --password value
​          用于非交互式密码输入的密码文件

​    --pcscdpath value
​         智能卡守护进程(pcscd)套接字文件的路径

​    --signer value
​          外部签名者(url或路径到ipc文件)

​    --unlock value
​          要解锁的帐户列表，以逗号分隔

​    --usb                          (default: false)
​         启用USB硬件钱包的监控和管理

## 别名(弃用)

​    --nousb                        (default: false)
​        禁止监控和管理USB硬件钱包(已弃用)

--whitelist value
      逗号分隔的块编号到哈希值的映射强制执行(<number>=<hash>)

​	(已弃用，支持——eth.requiredblocks)

## API和控制台

​    --authrpc.addr value           (default: "localhost")
​         已认证api的监听地址

​    --authrpc.jwtsecret value
​         用于经过身份验证的RPC端点的JWT秘密的路径

​    --authrpc.port value           (default: 8551)
​         认证api的监听端口

​    --authrpc.vhosts value         (default: "localhost")
​          接受请求的虚拟主机名列表(服务器强制)，以逗号分隔。接受'*'通配符。

​    --exec value
​          执行JavaScript语句

--graphql                      (default: false)
      在HTTP-RPC服务器上启用GraphQL。注意，GraphQL只能在以下情况下启动：同时也启动了一个HTTP服务器。

​    --graphql.corsdomain value
​          接受跨域请求的域列表，以逗号分隔
​          (browser enforced)

--graphql.vhosts value         (default: "localhost")
    接受请求的虚拟主机名列表(以逗号分隔)。接受'*'通配符。

--header value, -H value
      当使用--remotedb或geth attach时，将自定义报头传递给RPC服务。这个标志可以多次出现。

​    --http                         (default: false)
​        启用HTTP-RPC服务器

​    --http.addr value              (default: "localhost")
​        HTTP-RPC服务器监听接口

​    --http.api value
​        通过HTTP-RPC接口提供的API

​    --http.corsdomain value
​         接受跨域请求的域列表，以逗号分隔
​          (browser enforced)

​    --http.port value              (default: 8545)
​        HTTP-RPC服务器监听端口

​    --http.rpcprefix value
​          HTTP路径JSON-RPC服务的路径前缀。使用'/'服务于所有路径

--http.vhosts value            (default: "localhost")
      接受请求的虚拟主机名列表(以逗号分隔)接受'*'通配符。

​    --ipcdisable                   (default: false)
​          关闭IPC-RPC服务器

​    --ipcpath value
​        数据目录中IPC套接字/管道的文件名(显式路径转义它)

​    --jspath value                 (default: .)
​         loadScript的根路径

​    --preload value
​         要预加载到控制台的JavaScript文件列表，以逗号分隔

​    --rpc.allow-unprotected-txs    (default: false)
​        允许通过RPC提交未受保护(非EIP155签名)的事务

​    --rpc.enabledeprecatedpersonal (default: false)
​          启用(已弃用的)个人名称空间

​    --rpc.evmtimeout value         (default: 5s)
​          设置eth_call (0=infinite)的超时时间

​    --rpc.gascap value             (default: 50000000)
​          设置可以在eth_call/estimateGas中使用的gas上限(0=infinite)

​    --rpc.txfeecap value           (default: 1)
​          设置可以通过RPC api(0 =)发送的交易费用上限(以以太为单位)(0 =no cap)

​    --ws                           (default: false)
​          启用WS-RPC服务器

​    --ws.addr value                (default: "localhost")
​          WS-RPC服务器监听接口

​    --ws.api value
​         通过WS-RPC接口提供的API

​    --ws.origins value
​        接受websockets请求的来源

​    --ws.port value                (default: 8546)
​          WS-RPC服务器侦听端口

​    --ws.rpcprefix value
​        为JSON-RPC提供服务的HTTP路径前缀。使用'/'服务于所有路径。

## 开发人员链

​    --dev                          (default: false)
​          Ephemeral proof-of-authority network with a pre-funded developer account, mining
​          enabled

​    --dev.gaslimit value           (default: 11500000)
​          Initial block gas limit

​    --dev.period value             (default: 0)
​          Block period to use in developer mode (0 = mine only if transaction pending)

##    以太坊

​    --bloomfilter.size value       (default: 2048)
​          Megabytes of memory allocated to bloom-filter for pruning

​    --config value
​          TOML configuration file

​    --datadir value                (default: /Users/fjl/Library/Ethereum)
​          Data directory for the databases and keystore

​    --datadir.ancient value
​          Root directory for ancient data (default = inside chaindata)

​    --datadir.minfreedisk value
​          Minimum free disk space in MB, once reached triggers auto shut down (default =
​          --cache.gc converted to MB, 0 = disabled)

​    --db.engine value              (default: "leveldb")
​          Backing database implementation to use ('leveldb' or 'pebble')

​    --eth.requiredblocks value
​          Comma separated block number-to-hash mappings to require for peering
​          (<number>=<hash>)

​    --exitwhensynced               (default: false)
​          Exits after block synchronisation completes

​    --gcmode value                 (default: "full")
​          Blockchain garbage collection mode ("full", "archive")

​    --goerli                       (default: false)
​          Görli network: pre-configured proof-of-authority test network

​    --mainnet                      (default: false)
​          Ethereum mainnet

​    --networkid value              (default: 1)
​          Explicitly set network id (integer)(For testnets: use --rinkeby, --goerli,
​          --sepolia instead)

​    --override.shanghai value      (default: 0)
​          Manually specify the Shanghai fork timestamp, overriding the bundled setting

​    --rinkeby                      (default: false)
​          Rinkeby network: pre-configured proof-of-authority test network

​    --sepolia                      (default: false)
​          Sepolia network: pre-configured proof-of-work test network

​    --snapshot                     (default: true)
​          Enables snapshot-database mode (default = enable)

​    --syncmode value               (default: snap)
​          Blockchain sync mode ("snap", "full" or "light")

​    --txlookuplimit value          (default: 2350000)
​          Number of recent blocks to maintain transactions index for (default = about one
​          year, 0 = entire chain)

## 天然气价格预测

​    --gpo.blocks value             (default: 20)
​          Number of recent blocks to check for gas prices

​    --gpo.ignoreprice value        (default: 2)
​          Gas price below which gpo will ignore transactions

​    --gpo.maxprice value           (default: 500000000000)
​          Maximum transaction priority fee (or gasprice before London fork) to be
​          recommended by gpo

​    --gpo.percentile value         (default: 60)
​          Suggested gas price is the given percentile of a set of recent transaction gas
​          prices

## 轻客户端

​    --light.egress value           (default: 0)
​          Outgoing bandwidth limit for serving light clients (kilobytes/sec, 0 =
​          unlimited)

​    --light.ingress value          (default: 0)
​          Incoming bandwidth limit for serving light clients (kilobytes/sec, 0 =
​          unlimited)

​    --light.maxpeers value         (default: 100)
​          Maximum number of light clients to serve, or light servers to attach to

​    --light.nopruning              (default: false)
​          Disable ancient light chain data pruning

​    --light.nosyncserve            (default: false)
​          Enables serving light clients before syncing

​    --light.serve value            (default: 0)
​          Maximum percentage of time allowed for serving LES requests (multi-threaded
​          processing allows values over 100)

​    --ulc.fraction value           (default: 75)
​          Minimum % of trusted ultra-light servers required to announce a new head

​    --ulc.onlyannounce             (default: false)
​          Ultra light server sends announcements only

​    --ulc.servers value
​          List of trusted ultra-light servers

##    日志记录和调试

​    --fakepow                      (default: false)
​          Disables proof-of-work verification

​    --log.backtrace value
​          Request a stack trace at a specific logging statement (e.g. "block.go:271")

​    --log.debug                    (default: false)
​          Prepends log messages with call-site location (file and line number)

​    --log.file value
​          Write logs to a file

​    --log.json                     (default: false)
​          Format logs with JSON

​    --nocompaction                 (default: false)
​          Disables db compaction after import

​    --pprof                        (default: false)
​          Enable the pprof HTTP server

​    --pprof.addr value             (default: "127.0.0.1")
​          pprof HTTP server listening interface

​    --pprof.blockprofilerate value (default: 0)
​          Turn on block profiling with the given rate

​    --pprof.cpuprofile value
​          Write CPU profile to the given file

​    --pprof.memprofilerate value   (default: 524288)
​          Turn on memory profiling with the given rate

​    --pprof.port value             (default: 6060)
​          pprof HTTP server listening port

​    --remotedb value
​          URL for remote database

​    --trace value
​          Write execution trace to the given file

​    --verbosity value              (default: 3)
​          Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail

​    --vmodule value
​          Per-module verbosity: comma-separated list of <pattern>=<level> (e.g.
​          eth/*=5,p2p=4)

##    指标和统计数据

​    --ethstats value
​          Reporting URL of a ethstats service (nodename:secret@host:port)

​    --metrics                      (default: false)
​          Enable metrics collection and reporting

​    --metrics.addr value
​          Enable stand-alone metrics HTTP server listening interface.

​    --metrics.expensive            (default: false)
​          Enable expensive metrics collection and reporting

​    --metrics.influxdb             (default: false)
​          Enable metrics export/push to an external InfluxDB database

​    --metrics.influxdb.bucket value (default: "geth")
​          InfluxDB bucket name to push reported metrics to (v2 only)

​    --metrics.influxdb.database value (default: "geth")
​          InfluxDB database name to push reported metrics to

​    --metrics.influxdb.endpoint value (default: "http://localhost:8086")
​          InfluxDB API endpoint to report metrics to

​    --metrics.influxdb.organization value (default: "geth")
​          InfluxDB organization name (v2 only)

​    --metrics.influxdb.password value (default: "test")
​          Password to authorize access to the database

​    --metrics.influxdb.tags value  (default: "host=localhost")
​          Comma-separated InfluxDB tags (key/values) attached to all measurements

​    --metrics.influxdb.token value (default: "test")
​          Token to authorize access to the database (v2 only)

​    --metrics.influxdb.username value (default: "test")
​          Username to authorize access to the database

​    --metrics.influxdbv2           (default: false)
​          Enable metrics export/push to an external InfluxDB v2 database

​    --metrics.port value           (default: 6060)
​          Metrics HTTP server listening port.
​          Please note that --metrics.addr must be set
​          to start the server.

## 矿工

​    --mine                         (default: false)
​          Enable mining

​    --miner.etherbase value
​          0x prefixed public address for block mining rewards

​    --miner.extradata value
​          Block extra data set by the miner (default = client version)

​    --miner.gaslimit value         (default: 30000000)
​          Target gas ceiling for mined blocks

​    --miner.gasprice value         (default: 0)
​          Minimum gas price for mining a transaction

​    --miner.newpayload-timeout value (default: 2s)
​          Specify the maximum time allowance for creating a new payload


​    --miner.recommit value         (default: 2s)
​          Time interval to recreate the block being mined

##    MISC

​    --help, -h                     (default: false)
​          show help

​    --synctarget value
​          File for containing the hex-encoded block-rlp as sync target(dev feature)

​    --version, -v                  (default: false)
​          print the version

##    网络

​    --bootnodes value
​          Comma separated enode URLs for P2P discovery bootstrap

​    --discovery.dns value
​          Sets DNS discovery entry points (use "" to disable DNS)

​    --discovery.port value         (default: 30303)
​          Use a custom UDP port for P2P discovery

​    --identity value
​          Custom node name

​    --maxpeers value               (default: 50)
​          Maximum number of network peers (network disabled if set to 0)

​    --maxpendpeers value           (default: 0)
​          Maximum number of pending connection attempts (defaults used if set to 0)

​    --nat value                    (default: "any")
​          NAT port mapping mechanism (any|none|upnp|pmp|pmp:<IP>|extip:<IP>)

​    --netrestrict value
​          Restricts network communication to the given IP networks (CIDR masks)

​    --nodekey value
​          P2P node key file

​    --nodekeyhex value
​          P2P node key as hex (for testing)

​    --nodiscover                   (default: false)
​          Disables the peer discovery mechanism (manual peer addition)

​    --port value                   (default: 30303)
​          Network listening port

​    --v5disc                       (default: false)
​          Enables the experimental RLPx V5 (Topic Discovery) mechanism

## 性能调优

​    --cache value                  (default: 1024)
​          Megabytes of memory allocated to internal caching (default = 4096 mainnet full
​          node, 128 light mode)

​    --cache.blocklogs value        (default: 32)
​          Size (in number of blocks) of the log cache for filtering

​    --cache.database value         (default: 50)
​          Percentage of cache memory allowance to use for database io

​    --cache.gc value               (default: 25)
​          Percentage of cache memory allowance to use for trie pruning (default = 25% full
​          mode, 0% archive mode)

​    --cache.noprefetch             (default: false)
​          Disable heuristic state prefetch during block import (less CPU and disk IO, more
​          time waiting for data)

​    --cache.preimages              (default: false)
​          Enable recording the SHA3/keccak preimages of trie keys

​    --cache.snapshot value         (default: 10)
​          Percentage of cache memory allowance to use for snapshot caching (default = 10%
​          full mode, 20% archive mode)

​    --cache.trie value             (default: 15)
​          Percentage of cache memory allowance to use for trie caching (default = 15% full
​          mode, 30% archive mode)

​    --cache.trie.journal value     (default: "triecache")
​          Disk journal directory for trie cache to survive node restarts

​    --cache.trie.rejournal value   (default: 1h0m0s)
​          Time interval to regenerate the trie cache journal

​    --fdlimit value                (default: 0)
​          Raise the open file descriptor resource limit (default = system fd limit)

## 交易池

​    --txpool.accountqueue value    (default: 64)
​          Maximum number of non-executable transaction slots permitted per account

​    --txpool.accountslots value    (default: 16)
​          Minimum number of executable transaction slots guaranteed per account

​    --txpool.globalqueue value     (default: 1024)
​          Maximum number of non-executable transaction slots for all accounts

​    --txpool.globalslots value     (default: 5120)
​          Maximum number of executable transaction slots for all accounts

​    --txpool.journal value         (default: "transactions.rlp")
​          Disk journal for local transaction to survive node restarts

​    --txpool.lifetime value        (default: 3h0m0s)
​          Maximum amount of time non-executable transaction are queued

​    --txpool.locals value
​          Comma separated accounts to treat as locals (no flush, priority inclusion)

​    --txpool.nolocals              (default: false)
​          Disables price exemptions for locally submitted transactions

​    --txpool.pricebump value       (default: 10)
​          Price bump percentage to replace an already existing transaction

​    --txpool.pricelimit value      (default: 1)
​          Minimum gas price limit to enforce for acceptance into the pool

​    --txpool.rejournal value       (default: 1h0m0s)
​          Time interval to regenerate the local transaction journal

## 虚拟机

​    --vmdebug                      (default: false)
​          Record information useful for VM and contract debugging
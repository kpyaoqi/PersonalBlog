---
title: 命令模式与go-redis设计

date: 2022-02-13	

categories: Go其他知识点	

tags: [Go知识点,Go其他知识点]
---	

# [命令模式与go-redis command设计](https://www.cnblogs.com/wangxinwen/p/14311991.html)



目录

- [一、什么是命令（Command）模式](https://www.cnblogs.com/wangxinwen/p/14311991.html#一什么是命令command模式)
- [二、go-redis command相关代码](https://www.cnblogs.com/wangxinwen/p/14311991.html#二go-redis-command相关代码)
- [三、总结](https://www.cnblogs.com/wangxinwen/p/14311991.html#三总结)



#### 一、什么是命令（Command）模式[#](https://www.cnblogs.com/wangxinwen/p/14311991.html#一什么是命令command模式)

命令模式是行为型设计模式的一种，其目的是将一个请求封装为一个对象，从而使你可以用不同的请求对客户进行参数化。与另一种将每种命令与调用命令的对象结合形成一个专有类的方式相比，命令模式的优点有将调用操作的对象与知道如何实现该操作的对象解耦，增加新的命令不需要修改现有的类。
命令模式的结构如下：
[![command](https://img2020.cnblogs.com/blog/1447810/202101/1447810-20210122095641252-470394572.png)](https://img2020.cnblogs.com/blog/1447810/202101/1447810-20210122095641252-470394572.png)
参与者有：
1.Invoker请求者
要求该命令执行这个请求，即命令的调用者
2.Command接口
3.ConcreteCommand具体接口
4.Receiver接收者
命令的相关操作的实际实施者
5.Client
协作过程：
1.Client创建一个ConcreteCommand对象并指定它的Receiver对象
2.某Invoker对象存储该ConcreteCommand对象
3.该Invoker通过调用Command对象的Excute操作来提交一个请求。若该命令是可撤消的，ConcreteCommand就在执行Excute操作之前存储当前状态以用于取消该命令
4.ConcreteCommand对象对调用它的Receiver的一些操作以执行该请求

#### 二、go-redis command相关代码[#](https://www.cnblogs.com/wangxinwen/p/14311991.html#二go-redis-command相关代码)

```haskell
// commands.go

// Invoker请求者接口
type Cmdable interface {
      Pipeline() Pipeliner
      Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)
      TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)
      TxPipeline() Pipeliner
      Command(ctx context.Context) *CommandsInfoCmd
      ClientGetName(ctx context.Context) *StringCmd
      // ...
      // 和所有Redis命令的相关方法
}

// cmdable实现了Cmdable接口
type cmdable func(ctx context.Context, cmd Cmder) error  
func (c cmdable) Echo(ctx context.Context, message interface{}) *StringCmd {
	cmd := NewStringCmd(ctx, "echo", message)
	_ = c(ctx, cmd)
	return cmd
}
```

这里值得一提的是cmdable是一个函数类型，func(ctx context.Context, cmd Cmder) error
并且每个cmdable方法里都会有_ = c(ctx, cmd)，也就是如何去调用cmd在这里还没有明确写出
再回头看redis.go，会发现这样一段代码

```go
type Client struct {
      *baseClient
      cmdable
      hooks
      ctx context.Context
}

func NewClient(opt *Options) *Client {
      opt.init()

      c := Client{
            baseClient: newBaseClient(opt, newConnPool(opt)),
            ctx:        context.Background(),
      }
      c.cmdable = c.Process //划线

      return &c
}
```

c.cmdable = c.Process这行指定了请求如何调用Command的
在ctrl+左键追踪几层后，会在redis.go里找到调用的具体过程

```go
// redis.go 
func (c *baseClient) process(ctx context.Context, cmd Cmder) error {
      ......
      err := cn.WithWriter(ctx, c.opt.WriteTimeout, func(wr *proto.Writer) error {
            eturn writeCmd(wr, cmd)
      })
				
      err = cn.WithReader(ctx, c.cmdTimeout(cmd), cmd.readReply)
      ......			
```

然后再去找Command，这边就比较清晰了，都在command.go中

```scss
// command.go

// Command接口
type Cmder interface {
      Name() string
      FullName() string
      Args() []interface{}
      String() string
      stringArg(int) string
      firstKeyPos() int8
      setFirstKeyPos(int8)

      readTimeout() *time.Duration
      readReply(rd *proto.Reader) error
      
      SetErr(error)
      Err() error
}

// 还有许多Cmder的具体实现，其中一个实现的部分代码如下
type XPendingExtCmd struct {
      baseCmd
      val []XPendingExt
}
func (cmd *XPendingExtCmd) Val() []XPendingExt {
      return cmd.val
}
```

在这里没有看到Receiver，是因为每个Cmder实现都自己实现了所有功能，根本不需要额外的接收者对象。

#### 三、总结[#](https://www.cnblogs.com/wangxinwen/p/14311991.html#三总结)

有时必须向某对象提交请求，但并不知道关于被请求的操作或请求的接受者的任何信息。这个时候可以用到命令模式，通过将请求本身变成一个对象来使工具箱对象可向未指定的应用对象提出请求。
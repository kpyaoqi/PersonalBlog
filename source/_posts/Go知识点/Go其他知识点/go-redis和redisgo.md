---
title: go-redis和redisgo

date: 2022-02-08	

categories: Go其他知识点	

tags: [Go知识点,Go其他知识点]
---	

# 总结

[go-redis](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Fgithub.com%2Fgo-redis%2Fredis)和[redigo](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Fgithub.com%2Fgomodule%2Fredigo)底层是通过调用的万能 Do 方法实现, 但是

[redigo](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Fgithub.com%2Fgomodule%2Fredigo):

- 由于输入是万能类型所以必须记住每个命令的参数和返回值情况, 使用起来非常的不友好，
- 参数类型是万能类型导致在编译阶段无法检查参数类型,
- 每个命令都需要花时间记录使用方法，参数个数等，使用成本高；

[go-redis](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Fgithub.com%2Fgo-redis%2Fredis):

- 细化了每个redis每个命令的功能, 我们只需记住命令，具体的用法直接查看接口的申请就可以了，使用成本低；
- 其次它对数据类型按照redis底层的类型进行统一，编译时就可以帮助检查参数类型
- 并且它的响应统一采用 Result 的接口返回，确保了返回参数类型的正确性，对用户更加友好；

# Redigo库

## 介绍

[redigo](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Fgithub.com%2Fgomodule%2Fredigo) 是[Redis](https://cloud.tencent.com/developer/tools/blog-entry?target=http%3A%2F%2Fredis.io%2F)[数据库](https://cloud.tencent.com/solution/database?from_column=20065&from=20065)的[Go](https://cloud.tencent.com/developer/tools/blog-entry?target=http%3A%2F%2Fgolang.org%2F)客户端, 操作Redis基本和commands一样. Redigo命令基本都是通过Do方法来实现的.

```javascript
Do(ctx context.Context, cmd string, args ...interface{}) (interface{}, error)
```

复制

虽然调用`Do`函数万能参数可以实现所有的功能，但是使用起来非常的不友好，参数类型是万能类型，所以在编译阶段无法检查参数类型, 其次每个命令都需要花时间记录使用方法，参数个数等，使用成本高；

## 演示

演示基本的**连接池建立, ping, string操作, hash操作, list操作, expire**等操作

```javascript
package main

import (
   "fmt"
   "github.com/gomodule/redigo/redis"
)
func main() {
   // 新建一个连接池
   var pool *redis.Pool
   pool = &amp;redis.Pool{
      MaxIdle:     10,  //最初的连接数量
      MaxActive:   0,   //连接池最大连接数量,（0表示自动定义），按需分配
      IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
      Dial: func() (redis.Conn, error) { //要连接的redis数据库
         return redis.Dial("tcp", "localhost:6379")
      },
   }
   conn := pool.Get() //从连接池，取一个链接
   defer conn.Close()

   // 0. ping正常返回pong， 异常res is nil, err not nil
   res, err := conn.Do("ping")
   fmt.Printf("ping res=%v\n", res)
   if err != nil {
      fmt.Printf("ping err=%v\n", err.Error())
   }

   // string操作
   // set
   res, err = conn.Do("set", "name", "测试001")
   fmt.Printf("set res=%v\n", res)
   if err != nil {
      fmt.Printf("set err=%v\n", err.Error())
   }

   // get
   res, err = redis.String(conn.Do("get", "name"))
   fmt.Printf("get res=%v\n", res)
   if err != nil {
      fmt.Printf("get err=%v\n", err.Error())
   }

   // MSet   MGet
   res, err = conn.Do("MSet", "name", "测试001", "age", 18)
   fmt.Printf("MSet res=%v\n", res)
   if err != nil {
      fmt.Printf("MSet err=%v\n", err.Error())
   }

   r, err := redis.Strings(conn.Do("MGet", "name", "age"))
   fmt.Printf("MGet res=%v\n", r)
   if err != nil {
      fmt.Printf("MGet err=%v\n", err.Error())
   }

   // expire
   res, err = conn.Do("expire", "name", 5)
   fmt.Printf("expire res=%v\n", r)
   if err != nil {
      fmt.Printf("expire err=%v\n", err.Error())
   }

   // list操作
   // lpush lpop
   res, err = conn.Do("lpush", "hobby", "篮球", "足球", "乒乓球")
   fmt.Printf("lpush res=%v\n", r)
   if err != nil {
      fmt.Printf("lpush err=%v\n", err.Error())
   }

   // lpop
   rs, er := conn.Do("lpop", "hobby")
   fmt.Printf("lpop res=%v\n", rs)
   if er != nil {
      fmt.Printf("lpop err=%v\n", er.Error())
   }

   // hash 操作
   // hset
   res, err = conn.Do("HSet", "userinfo", "name", "lqz")
   fmt.Printf("HSet res=%v\n", r)
   if err != nil {
      fmt.Printf("HSet err=%v\n", err.Error())
   }

   // hget
   r4, er4 := conn.Do("HGet", "userinfo", "name")
   fmt.Printf("HGet res=%v\n", r4)
   if er4 != nil {
      fmt.Printf("HGet err=%v\n", er4.Error())
   }

}
```

# go-redis组件介绍和使用

## 介绍

[go-redis](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Fgithub.com%2Fgo-redis%2Fredis)提供了三种对应服务端的客户端模式，集群，哨兵，和单机模式，三种模式在连接池这一块都是公用的, 同时还提供了灵活的Hook机制, 其底层实际也是调用的万能 Do 方法.

![img](https://ask.qcloudimg.com/developer-images/article/7380026/ooeifbs9by.png)

但go-redis细化了每个redis每个命令的功能, 我们只需记住命令，具体的用法直接查看接口的申请就可以了，使用成本低；其次它对数据类型按照redis底层的类型进行统一，编译时就可以帮助检查参数类型, 并且它的响应统一采用 Result 的接口返回，确保了返回参数类型的正确性，对用户更加友好；

## 演示

演示基本的**连接池建立, ping, string操作, hash操作, list操作, expire**等操作

```javascript
func main() {
   var rdb = redis2.NewClient(
      &amp;redis2.Options{
         Addr:     "localhost:6379",
         Password: "", DB: 1,
         MinIdleConns: 1,
         PoolSize:     1000,
      })

   ctx := context.Background()
   res, err = rdb.Ping(ctx).Result()
   fmt.Printf("ping res=%v\n", res)
   if err != nil {
      fmt.Printf("ping err=%v\n", err.Error())
   }

   // string操作
   // set
   res, err = rdb.Set(ctx, "name", "测试001", 0).Result()
   fmt.Printf("set res=%v\n", res)
   if err != nil {
      fmt.Printf("set err=%v\n", err.Error())
   }

   // get
   res, err = rdb.Get(ctx, "name").Result()
   fmt.Printf("get res=%v\n", res)
   if err != nil {
      fmt.Printf("get err=%v\n", err.Error())
   }

   // MSet   MGet
   res, err = rdb.MSet(ctx, "name", "测试001", "age", "18").Result()
   fmt.Printf("MSet res=%v\n", res)
   if err != nil {
      fmt.Printf("MSet err=%v\n", err.Error())
   }

   var ret []interface{}
   ret, err = rdb.MGet(ctx, "name", "age").Result()
   fmt.Printf("MGet res=%v\n", ret)
   if err != nil {
      fmt.Printf("MGet err=%v\n", err.Error())
   }

   // expire
   res, err = rdb.Expire(ctx, "name", time.Second).Result()
   fmt.Printf("expire res=%v\n", res)
   if err != nil {
      fmt.Printf("expire err=%v\n", err.Error())
   }

   // list操作
   // lpush lpop
   res, err = rdb.LPush(ctx, "hobby", "篮球", "足球", "乒乓球").Result()
   fmt.Printf("lpush res=%v\n", res)
   if err != nil {
      fmt.Printf("lpush err=%v\n", err.Error())
   }

   // lpop
   rs, err = rdb.LPop(ctx, "hobby").Result()
   fmt.Printf("lpop res=%v\n", rs)
   if er != nil {
      fmt.Printf("lpop err=%v\n", er.Error())
   }

   // hash 操作
   // hset
   res, err = rdb.HSet(ctx, "userinfo", "name", "lqz").Result()
   fmt.Printf("HSet res=%v\n", r)
   if err != nil {
      fmt.Printf("HSet err=%v\n", err.Error())
   }

   // hget
   r4, er4 = rdb.HGet(ctx, "userinfo", "name").Result()
   fmt.Printf("HGet res=%v\n", r4)
   if er4 != nil {
      fmt.Printf("HGet err=%v\n", er4.Error())
   }

}
```
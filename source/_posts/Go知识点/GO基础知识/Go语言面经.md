---
title: Go语言面经

date: 2022-01-14	

categories: GO基础知识	

tags: [Go知识点,GO基础知识]
---	

# 进程、线程、协程的区别

**进程：**就是二进制可执行文件在计算机内存中的运行实例，**进程是操作系统最小的资源分配单位**，拥有独立的内存空间和系统资源，为了方便管理，每个进程都有自己的描述符，我们称之为**进程控制块**，即PCB，进程分类：

- **用户进程：位于用户空间中，是程序执行的实例**
- **内核进程：位于内核空间中，可以访问硬件**

进程在创建后，在执行过程中，其状态一直在变化。分别是：**初始态、就绪态、运行态、挂起态（阻塞）、终止态**

**线程：线程是操作系统最小的调度执行单位**，线程不能独立于进程而存在，其生命周期不可能逾越其所属的进程生命周期，**线程与进程一样拥有独立的PCB，但是没有独立的地址空间，即线程之间共享了地址空间，直接就能通信！！**进程的**大多数资源会被其内部的线程所共享。**

**协程：**程与进程、线程相比并不是一个维度的概念，协程不是被操作系统内核所管理的，而是**完全由程序所控制，也就是在用户态执行**。由程序员显式地定义和控制，可以在执行过程中暂停和恢复。这样带来的好处是性能大幅度的提升，因为不会像线程切换那样消耗资源。   

优点：

- 占用小：**协程更加轻量，创建成本更小，降低了内存消耗**，协程一般只占据极小的内存（2~5KB），而线程是1MB左右。虽然线程和协程都是独有栈，但是线程栈是固定的，比如在Java中，基本是2M，假如一个栈只有一个打印方法，还要为此开辟一个2M的栈，就太浪费了。而**Go的的协程具备动态收缩功能**，初始化为2KB，最大可达1GB
- **运行效率高：**线程切换需要从用户态->内核态->用户态，而协程切换是在用户态上，即用户态->用户态->用户态，其切换过程由语言层面的调度器（coroutine）或者语言引擎（goroutine）实现。
- 减少了同步锁：协程最终还是运行在线程上，本质上还是单线程运行，没有临界区域的话自然不需要锁的机制。**多协程自然没有竞争关系**。但是，如果存在临界区域，依然需要使用锁，**协程可以减少以往必须使用锁的场景**

缺点：

- 无法利用多核资源：协程运行在线程上，单线程应用无法很好的利用多核，只能以多进程方式启动。
- **协程不能有阻塞操作**：线程是抢占式，线程在遇见IO操作时候，线程从运行态→阻塞态，释放cpu使用权。这是由操作系统调度。**协程是非抢占式，如果遇见IO操作时候，协程是主动释放执行权限的，如果无法主动释放，程序将阻塞，无法往下执行，随之而来是整个线程被阻塞。**
- CPU密集型不是长处：假设这个线程中有一个协程是 CPU 密集型的他没有 IO 操作，也就是自己不会主动触发调度器调度的过程，那么就会出现其他协程得不到执行的情况，所以这种情况下需要程序员自己避免。

# Golang协程间如何通信

Go推荐使用通道（channel）的方式解决数据传递问题，在多个goroutine之间，channel负责传递数据，还能保证整个过程的并发安全性。  

# GMP模型(Go协程调度模型)

每个G的执行需要P和M的支持，M与P关联后才会形成一个有效的G运行环境，即 `工作线程+上下文环境`。  

![image-20230530155013217](/noteimg/C:/Users/zhuba/Desktop/PersonalBlog/source/_posts/Go知识点/GO基础知识/img/image-20230530155013217.png) 

- G：goroutine，一个G代表一个Go协程
- M：machine，一个M代表一个工作线程，runtime/debug 中的 **SetMaxThreads** 函数，设置 M 的最大数量
- P：processor，一个P代表执行一个Go代码段需要的上下文环境，由**启动时环境变量 $GOMAXPROCS 或者是由 runtime 的方法 GOMAXPROCS()** 决定

> M 与 P 的数量没有绝对关系，一个 M 阻塞，P 就会去创建或者切换另一个 M，所以，即使 P 的默认数量是 1，也有可能会创建很多个 M 出来。

**go func () 调度流程**：

 1、我们通过 go func () 来创建一个 goroutine；

 2、有两个存储 G 的队列，一个是局部调度器 P 的本地队列、一个是全局 G 队列。新创建的 G 会先保存在 P 的本地队列中，如果 P 的本地队列已经满了就会保存在全局的队列中；（所有的 M 都可以从全局队列中拉取 G 来执行）

 3、G 只能运行在 M 中，一个 M 必须持有一个 P，M 和 P 存在一一绑定的关系。M 会从 P 的本地队列弹出一个可执行状态的 G 来执行，如果 P 的本地队列为空，则  M 从全局队列中拉取；如果全局队列也为空，则从其他的 P 中拉取 G

 4、一个 M 调度 G 执行的过程是一个循环机制；

 5、当 M 执行某一个 G 时候如果发生了 syscall 或则其余阻塞操作，M 会阻塞，如果当前有一些 G 在执行，runtime  会把这个线程 M 从 P 中摘除 (detach)，然后再创建一个新的操作系统的线程 (如果有空闲的线程可用就复用空闲线程) 来服务于这个 P；

 6、当 M 系统调用结束时候，这个 G 会尝试获取一个空闲的 P 执行，并放入到这个 P 的本地队列。如果获取不到 P，那么这个线程 M 变成休眠状态， 加入到空闲线程中，然后这个 G 会被放入全局队列中。

# Golang Map底层

Map的底层实现是由哈希表（Hash Table）实现的。

哈希表是一种基于哈希函数的数据结构，它将键映射到存储位置，从而实现快速的查找。在Go语言中，Map的底层结构是一个哈希表数组，每个元素称为桶（Bucket）。每个桶存储着一个或多个键值对。

下面是Map底层的几个关键点：

1. 哈希函数：Map使用哈希函数将键转换为一个索引，用于确定键值对在哈希表数组中的存储位置。哈希函数应该具有良好的散列性，以减少冲突的发生。
2. 桶（Bucket）：Map底层的哈希表数组由多个桶组成，每个桶存储着一个或多个键值对。当多个键映射到同一个索引位置时，它们会被存储在同一个桶中，形成一个链表或其他数据结构。
3. 冲突解决：由于哈希函数的有限性，可能会发生不同的键映射到相同的索引位置，这就是哈希冲突。Map使用链表或其他冲突解决方法来处理冲突，确保所有键值对都能正确存储和访问。
4. 动态调整大小：Map具有自动扩容和收缩的能力，以适应键值对的增加或减少。当Map中的键值对数量达到一定阈值时，Map会自动进行扩容，重新分配更大的哈希表数组。反之，如果键值对数量减少到一定程度，Map会自动进行收缩，释放不必要的内存。

总结来说，Go语言中的Map底层使用哈希表实现，通过哈希函数将键映射到存储位置，解决键值对的查找、插入和删除操作。哈希冲突通过链表等冲突解决方法来处理。同时，Map具有动态调整大小的能力，以适应键值对的变化。这种底层实现使得Map在大多数情况下能够提供高效的性能和常数时间复杂度的操作。

# 如何实现Map的有序查找

一种常见的方法是使用有序的数据结构，如切片（Slice）或平衡二叉树（Balanced Binary Tree）来维护键值对的有序性。

使用切片（Slice）：可以将Map的键值对复制到一个切片中，并按照键的顺序排序切片。然后，可以使用二分查找算法在有序的切片中进行查找操作

使用平衡二叉树（如红黑树）：将Map的键作为树节点的关键字，值作为节点的值，构建一颗平衡二叉树。可以使用平衡二叉树的查找操作来实现有序查找。

# sync包(sync包提供了基本的同步基元)

#### 互斥锁 sync.Mutex

**互斥锁**是传统并发程序进行共享资源访问控制的主要方法。Go中由结构体`sync.Mutex`表示互斥锁，**保证同时只有一个 goroutine 可以访问共享资源**。

```go
var mutex sync.Mutex

func (m *Mutex) Lock()
func (m *Mutex) Unlock()
```

#### 读写锁 sync.RWMutex

在开发场景中，经常遇到多处并发读取，一次并发写入的情况，Go为了方便这些操作，**在互斥锁基础上，提供了读写锁操作**。  读写锁即针对读写操作的互斥锁，简单来说，就是将数据设定为 写模式（只写）或者读模式（只读）。使用读写锁可以分别针对读操作和写操作进行锁定和解锁操作。  

```go
var rwm sync.RWMutex
// 设定为写模式：与互斥锁使用方式一致，一路只写
func (*RWMutex) Lock()				// 锁定写
func (*RWMutex) Unlock()			// 解锁写
// 设定为读模式：对读执行加锁解锁，即多路只读
func (*RWMutex) RLock()
func (*RWMutex) RUnlock()
//返回值Locker是实现了接口`sync.Lokcer`的值，该接口同样被 `*sync.Mutex`和`*sync.RWMutex`实现，包含方法：`Lock`和`Unlock`。
func (rw *RWMutex) RLocker() Locker
```

#### 等待组 sync.WaitGroup

`sync.WaitGroup`类型的值也是并发安全的，该类型结构体中**内部拥有一个计数器**，计数器的值可以通过方法调用实现计数器的增加和减少 。  

当我们**添加了 N 个并发任务进行工作时，就将等待组的计数器值增加 N。每个任务完成时，这个值减1**。 同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时， 表示所有任务己经完成。  

等待组常用方法：

```go
//Add方法向内部计数加上delta，delta可以是负数；如果内部计数器变为0，切记不能减少为负数，会引发崩溃
func (wg *WaitGroup) Add(delta int)
//Done方法减少WaitGroup计数器的值，应在线程的最后执行
func (wg *WaitGroup) Done()
//Wait方法阻塞直到WaitGroup计数器减为0
func (wg *WaitGroup) Wait()
```

####  条件变量 sync.Cond

条件变量通常与锁配合使用： 

```go
func NewCond(l locker) *Cond        // 条件变量必须传入一个锁，二者需要配合使用
```

`*sync.Cond`类型有三个方法：

```go
//Broadcast唤醒所有等待c的线程
func (c *Cond) Broadcast()
//Signal唤醒等待c的一个线程（如果存在）
func (c *Cond) Signal()
//该方法会阻塞等待条件变量满足条件。也会对锁进行解锁，一旦收到通知则唤醒，并立即锁定该锁
//Wait自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L。
func (c *Cond) Wait()
```

#### Once 只执行一次

`sync.Once`，负责只执行一次，也即全局唯一操作。  使用方式如下：

```go
var once sync.Once
once.Do(func(){})           // Do方法的有效调用次数永远是1
```

`sync.Once`的典型应用场景是只执行一次的任务，如果这样的任务不适合在init函数中执行，该结构体类就会派上用场。  

sync.Once内部使用了“卫述语句、双重检查锁定、共享标记的原子操作”来实现`Once`功能。

#### 对象池 sync.Pool

`sync.Pool`可以作为临时值的容器，该容器具备自动伸缩、高效特性，同时也是并发安全的，其方法有：

```go
func (p *Pool) Get() interface{}
func (p *Pool) Put(x interface{})
```

- 如果池子从未Put过，其New字段也没有被赋值一个非nil值，那么Get方法返回结果一定是nil。  
- Get获取的值不一定存在于池中，如果Get到的值存在于池中，则该值Get后会被删除

#### 原子操作 sync/atomic

这些函数可以对一些数据类型进行原子操作这些函数提供的原子操作有5种：增、减、比较并交换、载入、存储、交换。  

##### 原子运算：增/减

增加函数的函数名前缀都是Add开头

```go
// 原子性的把一个int32类型变量 i32 增大3 ，下列函数返回值必定是已经被修改的值
newi32 := atomic.AddInt32(&i32, 3)      // 传入指针类型因为该函数需要获得数据的内存位置，以施加特殊的CPU指令
```

常见的增/减原子操作函数：

```go
func AddInt32(addr *int32, delta int32) (new int32)
func AddInt64(addr *int64, delta int64) (new int64)
func AddUint32(addr *uint32, delta uint32) (new uint32)
func AddUint64(addr *uint64, delta uint64) (new uint64)
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
```

注意:

- 如果需要执行减少操作，可以这样书写 atomic.AddInt32(&i32, -3)
- 对uint32执行增加NN（代表负整数，增加NN也可以理解为减少-NN）：atomic.AddUint32(&ui32, ^uint32(-NN-1)) 
- 不存在atomic.AddPointer的函数，因为unsafe.Poniter类型的值无法被增减

#####  原子运算：比较与替换

比较并替换即“Compare And Swap”，简称CAS。该类原子操作名称都以`CompareAndSwap`为前缀。   

```go
// 参数一：被操作数 参数二和参数三代表被操作数的旧值和新值
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
```

CAS的一些特点：

- CAS与锁相比，明显不同是它总是假设操作值未被改变，一旦确认这个假设为真，立即进行替换。所以锁的做法趋于悲观，CAS的做法趋于乐观。  
- CAS的优势：可以在不创建互斥量和不形成临界区的情况下，完成并发安全的值替换操作，可以大大减少性能损耗。  
- CAS的劣势：在被操作之被频繁变更的情况下，CAS操作容易失败，有时候需要for循环判断返回结构的bool来进行多次尝试 
- CAS操作不会阻塞协程，但是仍可能使流程执行暂时停滞（这种停滞极短）

应用场景：并发安全的更新一些类型的值，可以优先选择CAS操作。  

##### 原子读取：载入

为了原子的读取数值，Go提供了一系列载入函数，名称以`Load`为前缀。   

```go
func LoadInt32(addr *int32) (val int32)
func LoadInt64(addr *int64) (val int64)
func LoadUint32(addr *uint32) (val uint32)
func LoadUint64(addr *uint64) (val uint64)
func LoadUintptr(addr *uintptr) (val uintptr)
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
```

CAS与载入的配合示例：

```go
// value 增加 num
func addValue(value,num int32) {
    for {
        v := atomic.LoadInt32(&value)
        if atomic.ComapreAndSwapInt32(&value, v, (v + num)) {
            break
        }
    }
}
```

##### 原子写入：存储

在原子存储时，任何CPU都不会进行针对同一个值的读写操作，此时不会出现并发时候，别人读取到了修改了一半的值。

Go的`sync/atomic`包提供的存储函数都是以`Store`为前缀。  

```go
func StoreInt32(addr *int32, val int32)
func StoreInt64(addr *int64, val int64)
func StoreUint32(addr *uint32, val uint32)
func StoreUint64(addr *uint64, val uint64)
func StoreUintptr(addr *uintptr, val uintptr)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
```

```go
// 参数一为被操作数据的指针 参数二是要存储的新值
atomic.StoreInt32(i *int3, v int32)     
```

Go原子存储的特点：存储操作总会成功，因为不关心被操作值的旧值是什么，这与CAS有明显区别。  

#####  交换

交换与CAS操作相似，但是交换不关心被操作数据的旧值，而是直接设置新值，不过会返回被操作值的旧值。交换操作比CAS操作的约束更少，且比载入操作功能更强。 在Go中，交换操作都以`Swap`为前缀，示例：

```go
func SwapInt32(addr *int32, new int32) (old int32)
func SwapInt64(addr *int64, new int64) (old int64)
func SwapUint32(addr *uint32, new uint32) (old uint32)
func SwapUint64(addr *uint64, new uint64) (old uint64)
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
```

```go
// 参数一是被操作值指针  参数二是新值  返回值为旧值
atomic.SwapInt32(i *int32, v int32)         
```

# 怎么实现Map的并发安全(sync.Map，底层实际上用了一个Map缓存)

`sync.Map`类型提供了一种并发安全的键值对映射，它可以在多个goroutine之间安全地读取和写入数据

```go
// 创建一个并发安全的Map
var m sync.Map
// 添加键值对
m.Store("key1", "value1")
m.Store("key2", "value2")
// 从Map中获取值
value1, ok1 := m.Load("key1")
value2, ok2 := m.Load("key2")
// 删除键值对
m.Delete("key2")
// 遍历Map
m.Range(func(key, value interface{}) bool {
    fmt.Println("Key:", key, "Value:", value)
    return true
})
```

`sync.Map`内部被**分成了多个分片（shard）**，每个分片都包含了一个独立的映射，**以及一个互斥锁用于保护该分片的读写操作**。默认情况下，分**片的数量是32，可以通过调整`runtime.GOMAXPROCS()`来改变分片的数量。**当**需要读取或写入映射时，`sync.Map`会根据键的哈希值选择一个特定的分片**。对于读操作，只需要获取对应分片的互斥锁，并直接进行读取操作，不会对其他分片产生影响。**对于写操作，需要先获取对应分片的互斥锁，然后执行写入操作，并最后释放互斥锁。**这样就实现了对分片的并发访问控制，从而保证了并发安全性。

此外，`sync.Map`还使用了一种特殊的技术，即**当读操作遇到正在进行的写操作时，读操作会等待写操作完成，并重新加载最新的数据。**这种技术可以避免读操作读取到过期的数据，并确保读取的数据是最新的。

需要注意的是，**由于`sync.Map`采用了分片锁的机制，因此在高并发的情况下，对于频繁的写操作可能会导致锁的争用，进而影响性能**。因此，如果需要高性能的并发映射，可以**考虑使用其他第三方的并发安全映射库**，如`concurrent-map`、`go-cache`等，它们可能会采用更高效的底层实现方式。

# defer函数的使用场景（延迟Close、recover panic）

1. **解锁互斥锁**：在使用互斥锁进行临界区保护时，为了避免忘记解锁而导致死锁，可以**使用defer函数来确保在函数执行完毕后解锁互斥锁**。
2. 延迟**执行资源清理**：在函数执行过程中，如果需要进行一些**资源清理操作**，可以使用defer函数来延迟执行清理操作，确保无论函数执行的逻辑分支如何，都能够进行资源清理.(文件、数据库连接、网络连接)
3. **恢复错误**：在处理可能发生错误的代码段时，可以**使用defer函数来捕获和处理错误**，**避免错误传播到调用栈的更高层。**
4. **确保函数调用顺序：**当在函数中有多个函数调用的顺序很重要时，可以使用**defer函数来确保它们按照正确的顺序执行**。

# Map可以用数组作为Key吗（数组可以，切片不可以）

Go 语言中只要是可比较的类型都可以作为 key。除开 **slice，map，functions** 这几种类型，其他类型都是 OK 的。具体包括：布尔值、数字、字符串、指针、通道、接口类型、结构体、只包含上述类型的数组。

# Channel的阻塞和非阻塞（顺带问了select用法）

1. 阻塞通道：**默认情况下，通道是阻塞**的。**当发送方尝试将数据发送到通道时，如果通道已满（接收方没有准备好接收），发送方将被阻塞，直到有空间可用。类似地，当接收方尝试从通道中接收数据时，如果通道为空（发送方没有准备好发送），接收方将被阻塞，直到有数据可用**。这种阻塞行为使得通道成为一种有效的同步机制。
2. 非阻塞通道： 非阻塞通道的特点是在数据传输期间不会发生暂停，发送方和接收方可以继续执行其他操作。在 Go 中，**可以使用`select`语句结合`default`分支来实现非阻塞的通道操作**。`select`语句用于同时监视多个通道的状态，并执行相应的操作。

```go
ch := make(chan int) // 创建一个整数类型的通道

// 发送方（非阻塞）
select {
case ch <- 42: // 尝试发送数据到通道
    // 发送成功时执行的代码
default:
    // 通道已满时执行的代码或其他操作
}

// 接收方（非阻塞）
select {
case x := <-ch: // 尝试从通道接收数据
    // 接收到数据时执行的代码
default:
    // 通道为空时执行的代码或其他操作
}

```

# 实现一个接口C在指定时间内最大次数并发调用接口A与接口B

```go
duration := 5 * time.Second  // 指定时间段
startTime := time.Now()
// 循环调用接口A和接口B，直到达到指定时间或最大调用次数
for time.Since(startTime) < duration{
	wg.Add(1)
		// 并发调用接口A和接口B
		go func() {
			defer wg.Done()
			a := InterfaceA{}
			b := InterfaceB{}
			a.CallA()
			b.CallB()
		}()
}
// 等待所有调用完成
wg.Wait()
```

# 语法糖

在`Ethereum.StartMining()`函数中，出现了`if c, ok := s.engine.(*clique.Clique); ok` 的写法。这中写法是 Golang 中的语法糖，称为**类型断言**。具体的语法是 `value, ok := element.(T)`，它的含义是如果 `element` 是 `T` 类型的话，那么ok等于`True`, `value` 等于 `element` 的值。在 `if c, ok := s.engine.(*clique.Clique); ok` 语句中，就是在判断 `s.engine` 的是否为 `*clique.Clique` 类型。

```go
  var cli *clique.Clique
  if c, ok := s.engine.(*clique.Clique); ok {
   cli = c
  } else if cl, ok := s.engine.(*beacon.Beacon); ok {
   if c, ok := cl.InnerEngine().(*clique.Clique); ok {
    cli = c
   }
  }
```


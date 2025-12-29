# Go 语言并发编程 (07_concurrency)

本目录包含 Go 语言并发编程的学习示例代码,展示了 goroutines、channels、select 和 sync 等核心并发特性。主要参考自 [Go 语言之旅](https://tour.go-zh.org/)。

## 目录结构

### 核心概念
- [01_goroutines](#01_goroutines) - Go 协程
- [02_channels](#02_channels) - 信道 (Channel)
- [03_buffered-channels](#03_buffered-channels) - 带缓冲的信道
- [04_range-and-close](#04_range-and-close) - range 和 close
- [05_select](#05_select) - select 语句
- [06_default-selection](#06_default-selection) - 默认选择

### 练习与互斥锁
- [07_exercise-equivalent-binary-trees](#07_exercise-equivalent-binary-trees) - 练习: 等价二叉查找树
- [08_mutex-counter](#08_mutex-counter) - sync.Mutex 互斥锁
- [09_exercis-web-crawler](#09_exercis-web-crawler) - 练习:Web 爬虫

---

## 核心概念

### 01_goroutines

**Go 协程 (Goroutine)**

Go 协程是由 Go 运行时管理的轻量级线程。

**启动协程:**
```go
go f(x, y, z)  // 启动一个新的 goroutine 执行 f(x, y, z)
```

**示例:**
```go
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time. Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")  // 在新的 goroutine 中执行
    say("hello")     // 在主 goroutine 中执行
}
```

**输出 (交替出现):**
```
hello
world
hello
world
hello
world
hello
world
hello
```

**关键点:**
- `f`, `x`, `y`, `z` 的求值发生在**当前** goroutine 中
- `f` 的执行发生在**新的** goroutine 中
- Goroutine 在相同的地址空间中运行
- 访问共享内存时必须进行同步
- 非常轻量级,可以轻松创建成千上万个

**与线程对比:**

| 特性 | Goroutine | 线程 |
|------|-----------|------|
| 创建开销 | 约 2KB 栈空间 | 约 1-2MB 栈空间 |
| 调度方式 | Go 运行时调度 | 操作系统调度 |
| 切换成本 | 非常低 | 较高 |
| 数量限制 | 可创建百万级 | 通常几千个 |

---

### 02_channels

**信道 (Channel)**

信道是带有类型的管道,可以用信道操作符 `<-` 来发送或接收值。

**语法:**
```go
ch := make(chan int)  // 创建信道

ch <- v      // 将 v 发送至信道 ch
v := <-ch    // 从 ch 接收值并赋予 v
```

**示例 - 并行求和:**
```go
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum  // 发送 sum 到信道
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}
    
    c := make(chan int)
    go sum(s[: len(s)/2], c)   // 计算前半部分
    go sum(s[len(s)/2:], c)   // 计算后半部分
    x, y := <-c, <-c          // 接收两个结果
    
    fmt.Println(x, y, x+y)    // 输出: 17 -5 12
}
```

**关键点:**
- 箭头 `<-` 表示数据流的方向
- 信道在使用前必须用 `make` 创建
- 默认情况下,发送和接收操作会**阻塞**,直到另一端准备好
- 这种阻塞机制使得 goroutine 可以在没有显式锁的情况下进行同步

**信道操作:**

| 操作 | 语法 | 说明 |
|------|------|------|
| 创建 | `ch := make(chan T)` | 创建类型为 T 的信道 |
| 发送 | `ch <- v` | 将 v 发送到信道 ch |
| 接收 | `v := <-ch` | 从信道 ch 接收值 |
| 接收(忽略值) | `<-ch` | 接收但不使用值 |

---

### 03_buffered-channels

**带缓冲的信道**

信道可以是带缓冲的。将缓冲长度作为第二个参数提供给 `make`。

**语法:**
```go
ch := make(chan int, 100)  // 创建缓冲区大小为 100 的信道
```

**示例:**
```go
func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)  // 输出: 1
    fmt.Println(<-ch)  // 输出: 2
}
```

**阻塞规则:**
- **发送**:  仅当缓冲区**填满**后才会阻塞
- **接收**: 仅当缓冲区**为空**时才会阻塞

**对比:**

| 信道类型 | 发送阻塞时机 | 接收阻塞时机 |
|----------|-------------|-------------|
| 无缓冲信道 | 没有接收者时 | 没有发送者时 |
| 带缓冲信道 | 缓冲区满时 | 缓冲区空时 |

**关键点:**
- 缓冲信道是异步的
- 可以在没有接收者的情况下发送(直到缓冲区满)
- 缓冲区大小是容量,不是长度

---

### 04_range-and-close

**range 和 close**

发送者可以 `close` 信道来表示没有值需要发送了。

**语法:**
```go
close(ch)               // 关闭信道
v, ok := <-ch          // 测试信道是否关闭
for i := range ch {    // 不断接收值,直到信道关闭
    // 使用 i
}
```

**示例 - 斐波纳契数列:**
```go
func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)  // 发送完毕后关闭信道
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {  // 持续接收直到信道关闭
        fmt. Println(i)
    }
}
```

**输出:**
```
0
1
1
2
3
5
8
13
21
34
```

**重要规则:**

| 规则 | 说明 |
|------|------|
| 谁关闭 | **只应由发送者**关闭信道,不应由接收者关闭 |
| 关闭后发送 | 向已关闭信道发送数据会引发 **panic** |
| 关闭后接收 | 从已关闭信道接收会立即返回零值 |
| 是否必须关闭 | 通常无需关闭,只在需要通知接收者时才关闭 |

**检测信道是否关闭:**
```go
v, ok := <-ch
// ok 为 true:  信道未关闭,v 是接收到的值
// ok 为 false: 信道已关闭,v 是零值
```

---

### 05_select

**select 语句**

`select` 语句使一个 goroutine 可以等待多个通信操作。

**示例:**
```go
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x: 
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    
    fibonacci(c, quit)
}
```

**输出:**
```
0
1
1
2
3
5
8
13
21
34
quit
```

**select 特性:**

| 特性 | 说明 |
|------|------|
| 阻塞 | 阻塞直到某个分支可以继续执行 |
| 执行 | 执行准备好的分支 |
| 多个准备好 | 随机选择一个执行 |
| 语法 | 类似 switch,但每个 case 必须是通信操作 |

**使用场景:**
- 同时等待多个信道
- 实现超时机制
- 实现非阻塞的通信
- 协调多个 goroutine

---

### 06_default-selection

**默认选择 (default)**

当 `select` 中的其他分支都没有准备好时,`default` 分支就会执行。

**示例:**
```go
func main() {
    tick := time. Tick(100 * time. Millisecond)
    boom := time.After(500 * time.Millisecond)
    
    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom: 
            fmt.Println("BOOM!")
            return
        default:
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond)
        }
    }
}
```

**输出:**
```
    .
tick.
    .
tick.
    .
tick.
    .
tick.
    .
BOOM! 
```

**default 作用:**
- 实现**非阻塞**的发送或接收
- 避免 select 一直阻塞
- 可以执行其他工作

**非阻塞通信模式:**
```go
select {
case i := <-c:
    // 使用 i
default:
    // 从 c 接收会阻塞时执行
}

select {
case c <- v:
    // 发送成功
default:
    // 发送会阻塞时执行
}
```

---

## 练习与互斥锁

### 07_exercise-equivalent-binary-trees

**练习:  等价二叉查找树**

检查两个二叉树是否保存了相同的值序列。

**二叉树定义:**
```go
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```

**任务:**
1. 实现 `Walk` 函数 - 遍历树并将值发送到信道
2. 测试 `Walk` 函数
3. 实现 `Same` 函数 - 判断两个树是否包含相同值
4. 测试 `Same` 函数

**实现:**
```go
// Walk 遍历树 t,将所有值发送到信道 ch
func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)
}

// Same 判断 t1 和 t2 是否包含相同的值
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        Walk(t1, ch1)
        close(ch1)
    }()
    
    go func() {
        Walk(t2, ch2)
        close(ch2)
    }()
    
    for v1 := range ch1 {
        v2, ok := <-ch2
        if !ok || v1 != v2 {
            return false
        }
    }
    
    _, ok := <-ch2
    return ! ok
}
```

**测试:**
```go
func main() {
    t1 := tree.New(1)
    t2 := tree.New(1)
    t3 := tree.New(2)
    
    println(Same(t1, t2))  // true
    println(Same(t1, t3))  // false
}
```

**关键点:**
- 使用中序遍历获取排序后的值
- 使用 goroutine 并行遍历两棵树
- 使用 `close` 通知遍历完成
- 使用 `range` 简化接收逻辑

---

### 08_mutex-counter

**sync.Mutex 互斥锁**

信道非常适合各个 goroutine 间进行通信,但如果只是想保证每次只有一个 goroutine 能访问共享变量呢?

这就需要**互斥锁** (Mutex)。

**示例 - 并发安全计数器:**
```go
type SafeCounter struct {
    mu sync.Mutex
    v  map[string]int
}

// Inc 对给定键的计数加一
func (c *SafeCounter) Inc(key string) {
    c.mu.Lock()
    // 锁定使得一次只有一个 goroutine 可以访问 c.v
    c.v[key]++
    c.mu.Unlock()
}

// Value 返回给定键的计数值
func (c *SafeCounter) Value(key string) int {
    c.mu.Lock()
    defer c.mu.Unlock()  // 使用 defer 确保解锁
    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }
    
    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))  // 输出: 1000
}
```

**Mutex 方法:**

| 方法 | 说明 |
|------|------|
| `Lock()` | 获取锁,如果已被占用则阻塞 |
| `Unlock()` | 释放锁 |

**使用模式:**
```go
mu. Lock()
// 访问共享资源
mu.Unlock()

// 或者使用 defer
mu.Lock()
defer mu.Unlock()
// 访问共享资源
```

**关键点:**
- 互斥锁用于保护共享资源
- 一次只允许一个 goroutine 访问
- 必须成对使用 `Lock()` 和 `Unlock()`
- 推荐使用 `defer` 确保解锁
- Map 本身不是并发安全的,需要加锁

**通信 vs 互斥:**

| 场景 | 推荐方式 |
|------|---------|
| goroutine 间通信和数据传递 | 使用 Channel |
| 保护共享状态 | 使用 Mutex |

**Go 谚语:**
> "不要通过共享内存来通信,而应通过通信来共享内存。"

---

### 09_exercis-web-crawler

**练习: Web 爬虫**

使用 Go 的并发特性来并行化 Web 爬虫。

**任务:**
1. 并行地抓取 URL
2. 保证不重复抓取同一个 URL

**提示:**
- 可以用 map 缓存已获取的 URL
- 注意:  map 本身不是并发安全的! 

**实现:**
```go
func Crawl(url string, depth int, fetcher Fetcher) {
    ch := make(chan struct{})
    var crawl func(string, int)
    
    crawl = func(url string, depth int) {
        defer func() { ch <- struct{}{} }()
        
        if depth <= 0 {
            return
        }
        
        body, urls, err := fetcher.Fetch(url)
        if err != nil {
            fmt.Println(err)
            return
        }
        
        fmt.Printf("found: %s %q\n", url, body)
        
        for _, u := range urls {
            go crawl(u, depth-1)
        }
    }
    
    go crawl(url, depth)
    
    // 等待所有 goroutine 完成
    for i := 0; i < 1<<depth-1; i++ {
        <-ch
    }
}
```

**改进建议:**
1. **防止重复抓取:**
   ```go
   type SafeCache struct {
       mu    sync. Mutex
       cache map[string]bool
   }
   
   func (c *SafeCache) Set(key string) bool {
       c.mu. Lock()
       defer c.mu.Unlock()
       
       if c.cache[key] {
           return false  // 已存在
       }
       c.cache[key] = true
       return true  // 新增成功
   }
   ```

2. **使用 sync.WaitGroup:**
   ```go
   var wg sync.WaitGroup
   
   wg.Add(1)
   go func() {
       defer wg.Done()
       // 执行任务
   }()
   
   wg.Wait()  // 等待所有任务完成
   ```

**关键点:**
- 使用 goroutine 实现并行抓取
- 使用 channel 同步完成状态
- 使用 map + mutex 防止重复
- 或使用 sync.WaitGroup 更优雅地等待

---

## 并发模式总结

### 1. 生产者-消费者模式
```go
func producer(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}

func consumer(ch chan int) {
    for v := range ch {
        fmt. Println(v)
    }
}
```

### 2. Fan-Out / Fan-In 模式
```go
// Fan-out:  一个输入,多个处理器
for i := 0; i < numWorkers; i++ {
    go worker(input)
}

// Fan-in:  多个输入,一个输出
func fanIn(channels .. .<-chan int) <-chan int {
    out := make(chan int)
    for _, ch := range channels {
        go func(c <-chan int) {
            for v := range c {
                out <- v
            }
        }(ch)
    }
    return out
}
```

### 3. 超时模式
```go
select {
case result := <-ch:
    // 使用 result
case <-time.After(5 * time.Second):
    // 超时处理
}
```

### 4. 管道模式
```go
func gen(nums ... int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// 使用
c := gen(2, 3)
out := sq(c)
```

---

## 核心概念对比

### Goroutine vs Thread

| 特性 | Goroutine | Thread |
|------|-----------|--------|
| 栈大小 | 初始 2KB,按需增长 | 固定 1-2MB |
| 创建成本 | 极低 | 较高 |
| 切换成本 | 极低 | 较高 |
| 调度器 | Go 运行时 | OS 内核 |
| 数量 | 可百万级 | 一般几千 |

### Channel vs Mutex

| 场景 | Channel | Mutex |
|------|---------|-------|
| 通信和数据传递 | ✓ 推荐 | ✗ 不适合 |
| 保护共享状态 | 可以 | ✓ 推荐 |
| 所有权转移 | ✓ 推荐 | ✗ 不适合 |
| 同步 | ✓ 内置 | 需手动 |

### 有缓冲 vs 无缓冲信道

| 特性 | 无缓冲 | 有缓冲 |
|------|--------|--------|
| 同步 | 同步 | 异步 |
| 发送阻塞 | 无接收者时 | 缓冲区满时 |
| 接收阻塞 | 无发送者时 | 缓冲区空时 |
| 用途 | 强同步 | 削峰填谷 |

---

## 最佳实践

### 1. 避免数据竞争
```go
// ✗ 错误:  多个 goroutine 访问共享变量
var counter int
for i := 0; i < 1000; i++ {
    go func() {
        counter++  // 数据竞争! 
    }()
}

// ✓ 正确: 使用 mutex
var mu sync.Mutex
var counter int
for i := 0; i < 1000; i++ {
    go func() {
        mu.Lock()
        counter++
        mu.Unlock()
    }()
}
```

### 2. 使用 defer 确保解锁
```go
// ✓ 推荐
mu.Lock()
defer mu.Unlock()
// 复杂逻辑,可能有多个返回路径
```

### 3. 避免 goroutine 泄漏
```go
// ✗ 错误: goroutine 可能永远阻塞
go func() {
    result := <-ch  // 如果没有发送者,永远阻塞
}()

// ✓ 正确: 使用 context 或超时
go func() {
    select {
    case result := <-ch:
        // 处理 result
    case <-ctx.Done():
        return
    }
}()
```

### 4. 关闭信道的正确姿势
```go
// ✓ 只由发送者关闭
go func() {
    defer close(ch)
    for _, v := range data {
        ch <- v
    }
}()

// ✗ 接收者不应关闭信道
```

### 5. 使用带缓冲信道避免阻塞
```go
// ✗ 可能死锁
ch := make(chan int)
ch <- 1  // 阻塞,因为没有接收者

// ✓ 使用缓冲
ch := make(chan int, 1)
ch <- 1  // 不阻塞
```

---

## 运行示例

每个子目录都包含一个 `main.go` 文件,可以直接运行:

```bash
# 运行某个示例
cd 01_goroutines
go run main.go

# 或者直接指定路径
go run 01_goroutines/main. go
```

**需要外部包的示例:**
```bash
# 07_exercise-equivalent-binary-trees
go get golang.org/x/tour/tree
go run 07_exercise-equivalent-binary-trees/main.go
```

---

## 学习顺序建议

建议按照以下顺序学习: 

1. **Goroutine 基础** (01)
   - 理解轻量级线程的概念
   - 掌握 `go` 关键字的使用

2. **Channel 基础** (02-04)
   - 学习信道的创建和使用
   - 理解阻塞和同步机制
   - 掌握 close 和 range

3. **Select 语句** (05-06)
   - 理解多路复用
   - 掌握非阻塞通信

4. **实战练习** (07, 09)
   - 通过练习巩固并发知识
   - 学习实际问题的并发解决方案

5. **同步原语** (08)
   - 理解互斥锁的使用场景
   - 对比 Channel 和 Mutex

---

## 关键概念速查

### Goroutine
- 使用 `go` 关键字启动
- 轻量级,低开销
- 由 Go 运行时调度

### Channel
- 使用 `make(chan T)` 创建
- `<-` 操作符发送/接收
- 默认阻塞,用于同步
- 可以带缓冲

### Select
- 多路复用多个信道
- 随机选择准备好的分支
- `default` 实现非阻塞

### Mutex
- `Lock()` 和 `Unlock()`
- 保护共享资源
- 使用 `defer` 确保解锁

---

## 并发编程陷阱

### 1. 数据竞争
多个 goroutine 同时访问共享变量,至少一个在写入。

### 2. 死锁
所有 goroutine 都在等待,没有一个能继续执行。

### 3.  Goroutine 泄漏
Goroutine 被阻塞且永远无法结束。

### 4. 资源泄漏
未关闭的信道或未释放的锁。

### 5. 竞态条件
程序行为取决于不可控的执行顺序。

**检测工具:**
```bash
go run -race main.go  # 启用竞态检测器
```

---

## 参考资源

- [Go 语言之旅 - 并发](https://tour.go-zh.org/concurrency/1) - 官方交互式教程
- [Effective Go - 并发](https://go.dev/doc/effective_go#concurrency) - 官方最佳实践
- [Go Concurrency Patterns](https://go.dev/blog/pipelines) - 官方博客
- [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency) - 高级模式
- [Share Memory By Communicating](https://go.dev/blog/codelab-share) - 设计哲学
- [Go 并发编程实战](https://github.com/golang/go/wiki/LearnConcurrency) - 学习资源汇总

---

## Go 并发谚语

> "不要通过共享内存来通信,而应通过通信来共享内存。"
> — Rob Pike

> "并发不是并行。"
> — Rob Pike

> "信道编排,互斥保护。"
> — Go Proverbs

---

**注意:** 并发编程是 Go 语言最强大也最具特色的特性之一。掌握 goroutine 和 channel 是成为 Go 程序员的关键。建议通过大量实践来加深理解,并始终关注并发安全问题。
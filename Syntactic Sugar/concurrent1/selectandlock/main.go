// ============================= 1. Select基础使用 ====================
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func selectBasicDemo() {
	fmt.Println("=== Select基础使用 ===")

	// 1.1 创建多个管道
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan bool)

	defer func() {
		close(ch1)
		close(ch2)
		close(ch3)
	}()

	// 启动goroutine向不同管道发送数据
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "hello"
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		ch3 <- true
	}()

	// 1.2 select监听多个管道
	select {
	case n := <-ch1:
		fmt.Printf("从ch1接收到: %d\n", n)
	case s := <-ch2:
		fmt.Printf("从ch2接收到: %s\n", s)
	case b := <-ch3:
		fmt.Printf("从ch3接收到: %t\n", b)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("所有操作超时")
	}
}

// ============================= 2. Select循环监听 ====================
func selectLoopDemo() {
	fmt.Println("\n=== Select循环监听 ===")

	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})

	// 2.1 数据生产者
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 15; i++ {
			ch2 <- i
			time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		}
		close(ch2)
	}()

	// 2.2 使用for循环持续监听
	go func() {
		defer close(done)

		ch1Closed, ch2Closed := false, false

		for {
			select {
			case n, ok := <-ch1:
				if !ok {
					ch1Closed = true
					fmt.Println("ch1已关闭")
				} else {
					fmt.Printf("ch1: %d\n", n)
				}
			case n, ok := <-ch2:
				if !ok {
					ch2Closed = true
					fmt.Println("ch2已关闭")
				} else {
					fmt.Printf("ch2: %d\n", n)
				}
			case <-time.After(500 * time.Millisecond):
				fmt.Println("监控超时")
			}

			// 两个管道都关闭后退出
			if ch1Closed && ch2Closed {
				return
			}
		}
	}()

	<-done
}

// ============================= 3. Select超时控制 ====================
func selectTimeoutDemo() {
	fmt.Println("\n=== Select超时控制 ===")

	// 3.1 单个操作超时
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	select {
	case result := <-ch:
		fmt.Printf("操作成功: %d\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("操作超时!")
	}

	// 3.2 带默认分支的非阻塞操作
	fmt.Println("\n--- 非阻塞操作 ---")

	dataCh := make(chan int, 2)
	dataCh <- 1

	select {
	case dataCh <- 2:
		fmt.Println("成功写入数据")
	default:
		fmt.Println("管道已满，写入失败")
	}

	select {
	case val := <-dataCh:
		fmt.Printf("成功读取: %d\n", val)
	default:
		fmt.Println("没有数据可读")
	}

	close(dataCh)
}

// ============================= 4. 互斥锁使用 ====================
func mutexDemo() {
	fmt.Println("\n=== 互斥锁使用 ===")

	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)

	// 4.1 不加锁的问题演示
	fmt.Println("--- 不加锁的情况 ---")
	counter = 0
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			// 模拟竞态条件
			temp := counter
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			counter = temp + 1
			fmt.Printf("协程%d: 计数=%d\n", id, counter)
		}(i)
	}
	wg.Wait()
	fmt.Printf("最终计数(无锁): %d (应该为10)\n", counter)

	// 4.2 使用互斥锁保护共享数据
	fmt.Println("\n--- 使用互斥锁 ---")
	counter = 0
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()

			mutex.Lock()
			// 临界区开始
			temp := counter
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			counter = temp + 1
			fmt.Printf("协程%d: 计数=%d\n", id, counter)
			// 临界区结束
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("最终计数(有锁): %d\n", counter)
}

// ============================= 5. 读写锁使用 ====================
func rwMutexDemo() {
	fmt.Println("\n=== 读写锁使用 ===")

	var (
		data    int
		rwMutex sync.RWMutex
		wg      sync.WaitGroup
	)

	data = 0

	// 5.1 启动多个读协程
	wg.Add(8)
	for i := 0; i < 5; i++ {
		go reader(i, &data, &rwMutex, &wg)
	}

	// 5.2 启动多个写协程
	for i := 0; i < 3; i++ {
		go writer(i, &data, &rwMutex, &wg)
	}

	wg.Wait()
	fmt.Printf("最终数据值: %d\n", data)
}

func reader(id int, data *int, rwMutex *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		rwMutex.RLock() // 加读锁
		val := *data
		fmt.Printf("读者%d: 读取数据=%d\n", id, val)
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		rwMutex.RUnlock() // 解读锁

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func writer(id int, data *int, rwMutex *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		rwMutex.Lock() // 加写锁
		oldVal := *data
		newVal := oldVal + 1
		*data = newVal
		fmt.Printf("写者%d: %d -> %d\n", id, oldVal, newVal)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		rwMutex.Unlock() // 解写锁

		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
}

// ============================= 6. 条件变量使用 ====================
func condDemo() {
	fmt.Println("\n=== 条件变量使用 ===")

	var (
		queue    []int
		mutex    sync.Mutex
		cond     = sync.NewCond(&mutex)
		wg       sync.WaitGroup
		capacity = 3
	)

	// 6.1 生产者
	wg.Add(2)
	go producer(1, &queue, cond, &wg, capacity)
	go producer(2, &queue, cond, &wg, capacity)

	// 6.2 消费者
	wg.Add(2)
	go consumer(1, &queue, cond, &wg)
	go consumer(2, &queue, cond, &wg)

	wg.Wait()
	fmt.Println("所有生产消费完成")
}

func producer(id int, queue *[]int, cond *sync.Cond, wg *sync.WaitGroup, capacity int) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

		cond.L.Lock()

		// 等待队列有空间
		for len(*queue) >= capacity {
			fmt.Printf("生产者%d: 队列已满，等待...\n", id)
			cond.Wait()
		}

		// 生产数据
		*queue = append(*queue, i)
		fmt.Printf("生产者%d: 生产 %d, 队列长度: %d\n", id, i, len(*queue))

		cond.L.Unlock()
		cond.Broadcast() // 通知消费者
	}
}

func consumer(id int, queue *[]int, cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

		cond.L.Lock()

		// 等待队列有数据
		for len(*queue) == 0 {
			fmt.Printf("消费者%d: 队列为空，等待...\n", id)
			cond.Wait()
		}

		// 消费数据
		item := (*queue)[0]
		*queue = (*queue)[1:]
		fmt.Printf("消费者%d: 消费 %d, 队列长度: %d\n", id, item, len(*queue))

		cond.L.Unlock()
		cond.Broadcast() // 通知生产者
	}
}

// ============================= 7. 实用工具函数 ====================
func utilityFunctions() {
	fmt.Println("\n=== 实用工具函数 ===")

	// 7.1 非阻塞发送
	ch := make(chan int, 1)

	if trySend(ch, 100) {
		fmt.Println("发送成功")
	} else {
		fmt.Println("发送失败")
	}

	// 7.2 非阻塞接收
	if val, ok := tryRecv(ch); ok {
		fmt.Printf("接收成功: %d\n", val)
	} else {
		fmt.Println("接收失败")
	}

	close(ch)
}

// 非阻塞发送
func trySend(ch chan int, value int) bool {
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}

// 非阻塞接收
func tryRecv(ch chan int) (int, bool) {
	select {
	case val, ok := <-ch:
		return val, ok
	default:
		return 0, false
	}
}

// ============================= 主函数入口 ====================
func main() {
	rand.Seed(time.Now().UnixNano())

	selectBasicDemo()
	selectLoopDemo()
	selectTimeoutDemo()
	mutexDemo()
	rwMutexDemo()
	condDemo()
	utilityFunctions()

	fmt.Println("\n=== 所有示例执行完成 ===")
}

// ============================= 总结知识点 ====================
/*
Select和锁核心知识点总结：

1. Select多路复用：
   - 监听多个管道操作，任一case就绪即执行
   - 多个case就绪时随机选择一个执行
   - 没有case就绪且无default时会阻塞
   - 配合for循环实现持续监听

2. Select应用场景：
   - 超时控制：time.After()
   - 非阻塞操作：default分支
   - 多管道监听：同时监控多个通信通道
   - 退出信号：context.Done()监听

3. 互斥锁sync.Mutex：
   - Lock()/Unlock()保护临界区
   - 解决竞态条件问题
   - 保证数据操作的原子性
   - 不可重入，重复加锁会导致死锁

4. 读写锁sync.RWMutex：
   - RLock()/RUnlock()：读锁，共享访问
   - Lock()/Unlock()：写锁，独占访问
   - 读多写少场景性能更优
   - 读锁之间不互斥，写锁完全互斥

5. 条件变量sync.Cond：
   - Wait()：等待条件满足
   - Signal()：唤醒一个等待者
   - Broadcast()：唤醒所有等待者
   - 必须与锁配合使用，Wait前加锁

6. 最佳实践：
   - 使用defer确保锁释放
   - 条件变量检查使用for循环而非if
   - 锁粒度尽量小，减少持有时间
   - Select超时避免永久阻塞

7. 注意事项：
   - 避免锁嵌套导致的死锁
   - nil管道在select中会被忽略
   - 条件变量Wait前必须持有锁
   - 锁和条件变量应使用指针传递
*/

// ============================= 1. 无缓冲管道 ====================
package main

import (
	"fmt"
	"sync"
	"time"
)

func unbufferedChannelDemo() {
	fmt.Println("=== 无缓冲管道演示 ===")

	// 1.1 错误示例：同步读写会导致死锁
	/*
		ch := make(chan int) // 无缓冲
		defer close(ch)
		ch <- 123           // 阻塞，没有接收者
		fmt.Println(<-ch)   // 永远不会执行
	*/

	// 1.2 正确用法：在不同goroutine中读写
	ch := make(chan int)
	defer close(ch)

	go func() {
		fmt.Println("子协程发送数据: 123")
		ch <- 123 // 发送数据
	}()

	data := <-ch // 主协程接收数据
	fmt.Printf("主协程收到数据: %d\n", data)
}

// ============================= 2. 有缓冲管道 ====================
func bufferedChannelDemo() {
	fmt.Println("\n=== 有缓冲管道演示 ===")

	// 2.1 基本使用
	ch := make(chan int, 3) // 缓冲区大小为3
	defer close(ch)

	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4  // 如果取消注释会阻塞，因为缓冲区已满

	fmt.Printf("管道状态: 长度=%d, 容量=%d\n", len(ch), cap(ch))

	// 读取数据
	fmt.Println("读取:", <-ch)
	fmt.Println("读取:", <-ch)
	fmt.Println("读取:", <-ch)
}

// ============================= 3. 管道阻塞条件 ====================
func channelBlockingConditions() {
	fmt.Println("\n=== 管道阻塞条件演示 ===")

	// 3.1 无缓冲管道阻塞
	ch1 := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "hello"
	}()
	fmt.Println("等待无缓冲管道数据...", <-ch1)
	close(ch1)

	// 3.2 有缓冲管道满时阻塞
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println("缓冲区已满，继续写入会阻塞")
	go func() {
		time.Sleep(200 * time.Millisecond)
		<-ch2 // 释放一个位置
	}()
	ch2 <- 3 // 会等待直到有空间
	fmt.Println("成功写入第三个数据")
	close(ch2)

	// 3.3 nil管道阻塞（运行时错误）
	/*
		var ch3 chan int
		ch3 <- 1  // 对nil管道写入会永久阻塞
	*/
}

// ============================= 4. 管道panic情况 ====================
func channelPanicCases() {
	fmt.Println("\n=== 管道panic情况 ===")

	// 4.1 关闭nil管道会导致panic
	/*
		var ch1 chan int
		close(ch1) // panic: close of nil channel
	*/

	// 4.2 向已关闭管道写入数据会导致panic
	ch2 := make(chan int, 1)
	close(ch2)
	// ch2 <- 1 // panic: send on closed channel

	// 4.3 重复关闭管道会导致panic
	ch3 := make(chan int)
	close(ch3)
	// close(ch3) // panic: close of closed channel

	fmt.Println("所有panic情况已避免")
}

// ============================= 5. 单向管道 ====================
func directionalChannelsDemo() {
	fmt.Println("\n=== 单向管道演示 ===")

	// 5.1 创建双向管道
	ch := make(chan int, 2)

	// 转换为单向管道
	var sendCh chan<- int = ch // 只写管道
	var recvCh <-chan int = ch // 只读管道

	// 使用单向管道
	go producer(sendCh)
	consumer(recvCh)

	close(ch)
}

// 生产者：只能写入数据
func producer(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i * 10
		fmt.Printf("生产者发送: %d\n", i*10)
	}
}

// 消费者：只能读取数据
func consumer(ch <-chan int) {
	for i := 0; i < 3; i++ {
		data := <-ch
		fmt.Printf("消费者接收: %d\n", data)
	}
}

// ============================= 6. for range遍历管道 ====================
func channelRangeDemo() {
	fmt.Println("\n=== for range遍历管道 ===")

	// 6.1 正确用法：发送方关闭管道
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // 发送方负责关闭
	}()

	// 自动读取直到管道关闭
	for value := range ch {
		fmt.Println("遍历读取:", value)
	}

	// 6.2 已关闭管道的读取行为
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	// 可以继续读取已关闭管道中的数据
	fmt.Println("已关闭管道读取:")
	for i := 0; i < 4; i++ {
		value, ok := <-ch2
		fmt.Printf("值: %d, 能否读取: %t\n", value, ok)
	}
}

// ============================= 7. 管道作为同步工具 ====================
func channelAsSyncTool() {
	fmt.Println("\n=== 管道作为同步工具 ===")

	// 7.1 等待子协程完成
	done := make(chan struct{})

	go func() {
		fmt.Println("子协程开始工作...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("子协程工作完成")
		done <- struct{}{} // 发送完成信号
	}()

	<-done // 阻塞等待完成信号
	fmt.Println("主协程继续执行")
	close(done)

	// 7.2 管道作为互斥锁
	fmt.Println("\n--- 管道实现互斥锁 ---")
	var counter int
	lock := make(chan struct{}, 1) // 缓冲区为1的管道作为锁

	// 启动多个goroutine并发修改
	for i := 0; i < 5; i++ {
		go func(id int) {
			lock <- struct{}{} // 获取锁
			old := counter
			time.Sleep(10 * time.Millisecond) // 模拟处理时间
			counter = old + 1
			fmt.Printf("协程%d: 计数从%d增加到%d\n", id, old, counter)
			<-lock // 释放锁
		}(i)
	}

	time.Sleep(500 * time.Millisecond) // 等待所有goroutine完成
	fmt.Printf("最终计数: %d\n", counter)
	close(lock)
}

// ============================= 8. WaitGroup使用 ====================
func waitGroupDemo() {
	fmt.Println("\n=== WaitGroup使用演示 ===")

	var wg sync.WaitGroup

	// 8.1 基本使用
	wg.Add(3) // 设置等待3个协程

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done() // 确保Done被调用
			fmt.Printf("Worker %d 开始\n", id)
			time.Sleep(time.Duration(100+id*50) * time.Millisecond)
			fmt.Printf("Worker %d 完成\n", id)
		}(i)
	}

	fmt.Println("主协程等待所有worker完成...")
	wg.Wait() // 阻塞直到所有Done被调用
	fmt.Println("所有worker已完成!")

	// 8.2 动态添加任务
	fmt.Println("\n--- 动态添加任务 ---")
	var wg2 sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg2.Add(1) // 动态添加
		go func(id int) {
			defer wg2.Done()
			fmt.Printf("动态任务 %d 执行\n", id)
		}(i)
	}

	wg2.Wait()
	fmt.Println("动态任务全部完成")
}

// ============================= 9. WaitGroup注意事项 ====================
func waitGroupPrecautions() {
	fmt.Println("\n=== WaitGroup注意事项 ===")

	// 9.1 必须传递指针
	var wg sync.WaitGroup
	wg.Add(1)

	// 错误：传递值拷贝
	// go badWorker(wg) // 会导致死锁

	// 正确：传递指针
	go goodWorker(&wg)

	wg.Wait()
	fmt.Println("Worker完成")
}

func badWorker(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("badWorker执行")
}

func goodWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("goodWorker执行")
}

// ============================= 主函数入口 ====================
func main() {
	unbufferedChannelDemo()
	bufferedChannelDemo()
	channelBlockingConditions()
	channelPanicCases()
	directionalChannelsDemo()
	channelRangeDemo()
	channelAsSyncTool()
	waitGroupDemo()
	waitGroupPrecautions()

	fmt.Println("\n=== 所有示例执行完成 ===")
}

// ============================= 总结知识点 ====================
/*
管道和WaitGroup核心知识点总结：

1. 无缓冲管道：
   - 同步通信，发送和接收必须同时准备
   - 发送会阻塞直到有接收者，反之亦然
   - 必须在不同goroutine中使用

2. 有缓冲管道：
   - 异步通信，缓冲区满时发送阻塞，空时接收阻塞
   - 使用len()和cap()获取状态
   - 可以临时存储数据

3. 管道阻塞条件：
   - 无缓冲管道的同步读写
   - 有缓冲管道：读空缓冲区/写满缓冲区
   - 对nil管道进行任何操作

4. 管道panic情况：
   - 关闭nil管道
   - 向已关闭管道写入数据
   - 重复关闭管道

5. 单向管道：
   - chan<- Type: 只写管道
   - <-chan Type: 只读管道
   - 用于限制函数对管道的操作权限

6. for range遍历：
   - 自动读取直到管道关闭
   - 必须由发送方关闭管道
   - 已关闭管道可继续读取剩余数据

7. 管道同步应用：
   - 等待子协程完成
   - 实现简单的互斥锁
   - 协程间通信和协调

8. WaitGroup使用：
   - Add(): 设置等待数量
   - Done(): 标记任务完成
   - Wait(): 阻塞等待所有完成
   - 必须传递指针，不能复制

9. 最佳实践：
   - 发送方负责关闭管道
   - 使用defer确保资源释放
   - WaitGroup传递指针而非值
   - 避免死锁和竞态条件
*/

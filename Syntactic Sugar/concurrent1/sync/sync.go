// ============================= 1. sync.Once 使用 ====================
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func onceDemo() {
	fmt.Println("=== sync.Once 演示 ===")

	var (
		initCount int
		once      sync.Once
		wg        sync.WaitGroup
	)

	// 1.1 初始化函数（只会执行一次）
	initFunc := func() {
		initCount++
		fmt.Printf("初始化函数执行，计数: %d\n", initCount)
	}

	// 1.2 多个goroutine同时调用初始化
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			once.Do(initFunc) // 确保初始化只执行一次
			fmt.Printf("协程 %d 继续执行\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Printf("最终初始化计数: %d\n", initCount)
}

// ============================= 2. sync.Pool 使用 ====================
func poolDemo() {
	fmt.Println("\n=== sync.Pool 演示 ===")

	var (
		objectCount atomic.Int64
		pool        sync.Pool
		wg          sync.WaitGroup
	)

	// 2.1 定义池中对象的创建函数
	pool.New = func() interface{} {
		objectCount.Add(1)
		return &BigObject{Data: "新创建的对象"}
	}

	// 2.2 并发使用对象池
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(id int) {
			defer wg.Done()

			// 从池中获取对象
			obj := pool.Get().(*BigObject)
			obj.ID = id

			// 模拟使用对象
			// fmt.Printf("使用对象: %+v\n", obj)

			// 使用完毕后放回池中
			pool.Put(obj)
		}(i)
	}
	wg.Wait()

	fmt.Printf("总共创建的对象数量: %d (远小于100)\n", objectCount.Load())
}

type BigObject struct {
	ID   int
	Data string
}

// ============================= 3. sync.Map 使用 ====================
func syncMapDemo() {
	fmt.Println("\n=== sync.Map 演示 ===")

	var (
		syncMap sync.Map
		wg      sync.WaitGroup
	)

	// 3.1 并发安全的写入
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				syncMap.Store(key, id*100+j)
			}
		}(i)
	}
	wg.Wait()

	// 3.2 基本操作演示
	fmt.Println("--- 基本操作 ---")

	// 存储值
	syncMap.Store("name", "Alice")
	syncMap.Store("age", 25)

	// 加载值
	if name, ok := syncMap.Load("name"); ok {
		fmt.Printf("姓名: %v\n", name)
	}

	// 加载或存储
	if actual, loaded := syncMap.LoadOrStore("age", 30); loaded {
		fmt.Printf("年龄已存在: %v\n", actual)
	}

	// 删除并获取
	if value, loaded := syncMap.LoadAndDelete("name"); loaded {
		fmt.Printf("删除的姓名: %v\n", value)
	}

	// 3.3 遍历所有键值对
	fmt.Println("--- 遍历所有数据 ---")
	count := 0
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		count++
		return count < 10 // 限制输出数量
	})
}

// ============================= 4. atomic 原子操作 ====================
func atomicDemo() {
	fmt.Println("\n=== atomic 原子操作演示 ===")

	// 4.1 原子类型使用
	var (
		counter atomic.Int64
		flag    atomic.Bool
		wg      sync.WaitGroup
	)

	// 设置初始值
	counter.Store(100)
	flag.Store(true)

	fmt.Printf("初始值 - 计数器: %d, 标志: %t\n", counter.Load(), flag.Load())

	// 4.2 并发原子操作
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()

			// 原子增加
			counter.Add(int64(id + 1))

			// 原子交换
			oldFlag := flag.Swap(false)
			fmt.Printf("协程 %d: 旧标志=%t, 新标志=%t\n", id, oldFlag, flag.Load())

			// 原子比较并交换 (CAS)
			success := flag.CompareAndSwap(false, true)
			fmt.Printf("协程 %d CAS操作: %t\n", id, success)
		}(i)
	}
	wg.Wait()

	fmt.Printf("最终计数器值: %d\n", counter.Load())
}

// ============================= 5. CAS 乐观锁实现 ====================
func casDemo() {
	fmt.Println("\n=== CAS 乐观锁演示 ===")

	var (
		sharedValue int64
		wg          sync.WaitGroup
	)

	// 5.1 使用CAS实现无锁计数器
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()

			for {
				// 读取当前值
				current := atomic.LoadInt64(&sharedValue)
				// 尝试原子更新
				if atomic.CompareAndSwapInt64(&sharedValue, current, current+1) {
					fmt.Printf("协程 %d: 成功增加 %d -> %d\n", id, current, current+1)
					break
				}
				// 如果失败，循环重试
				// fmt.Printf("协程 %d: CAS失败，重试\n", id)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("CAS最终结果: %d\n", sharedValue)
}

// ============================= 6. atomic.Value 使用 ====================
func atomicValueDemo() {
	fmt.Println("\n=== atomic.Value 演示 ===")

	var (
		config atomic.Value
		wg     sync.WaitGroup
	)

	// 6.1 存储配置对象
	type Config struct {
		Server  string
		Port    int
		Timeout int
	}

	// 初始配置
	initialConfig := Config{
		Server:  "localhost",
		Port:    8080,
		Timeout: 30,
	}
	config.Store(initialConfig)

	// 6.2 并发读取配置
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()

			if cfg, ok := config.Load().(Config); ok {
				fmt.Printf("协程 %d 读取配置: %+v\n", id, cfg)
			}
		}(i)
	}

	// 6.3 更新配置
	newConfig := Config{
		Server:  "api.example.com",
		Port:    443,
		Timeout: 60,
	}
	config.Store(newConfig)
	fmt.Println("配置已更新")

	wg.Wait()

	// 6.4 错误用法演示
	fmt.Println("--- 错误用法 ---")
	var badValue atomic.Value
	// badValue.Store(nil) // 这会panic: store of nil value
	badValue.Store("hello")
	// badValue.Store(123) // 这会panic: store of inconsistently typed value
}

// ============================= 7. 性能对比示例 ====================
func performanceDemo() {
	fmt.Println("\n=== 性能对比演示 ===")

	var (
		mu        sync.Mutex
		atomicVal atomic.Int64
		mutexVal  int64
		wg        sync.WaitGroup
	)

	iterations := 10000
	wg.Add(2)

	// 7.1 原子操作性能
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			atomicVal.Add(1)
		}
	}()

	// 7.2 互斥锁性能
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			mu.Lock()
			mutexVal++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Printf("原子操作结果: %d\n", atomicVal.Load())
	fmt.Printf("互斥锁结果: %d\n", mutexVal)
}

// ============================= 主函数入口 ====================
func main() {
	onceDemo()
	poolDemo()
	syncMapDemo()
	atomicDemo()
	casDemo()
	atomicValueDemo()
	performanceDemo()

	fmt.Println("\n=== 所有sync示例执行完成 ===")
}

// ============================= 总结知识点 ====================
/*
sync包和atomic核心知识点总结：

1. sync.Once：
   - 确保函数只执行一次，无论并发调用多少次
   - 内部使用互斥锁+原子操作实现
   - 适用于懒加载、单例模式等场景

2. sync.Pool：
   - 临时对象池，减少GC压力
   - Get()获取对象，Put()放回对象
   - 对象可能被GC回收，不适合存储重要资源
   - 适用于创建成本高的临时对象

3. sync.Map：
   - 并发安全的Map实现
   - Store/Load/LoadOrStore/LoadAndDelete/Range操作
   - 读多写少场景性能更好
   - 比普通map+mutex更复杂的API

4. atomic原子操作：
   - 提供基础类型的原子操作
   - Load/Store/Swap/Add/CompareAndSwap
   - 无锁编程，性能优于互斥锁
   - 支持Int32/Int64/Uint32/Uint64/Bool/Pointer/Value

5. CAS乐观锁：
   - CompareAndSwap比较并交换
   - 无锁并发控制，失败时重试
   - 可能产生ABA问题
   - 适用于简单计数器等场景

6. atomic.Value：
   - 可以存储任意类型，但类型必须一致
   - 不能存储nil值
   - 适用于配置等需要原子更新的场景

7. 使用场景对比：
   - Once: 一次性初始化
   - Pool: 高频创建销毁的对象
   - Map: 并发安全键值存储
   - atomic: 简单计数器、标志位
   - Mutex: 复杂临界区保护

8. 最佳实践：
   - 选择合适的并发控制工具
   - 避免过度优化，先保证正确性
   - 注意原子类型的不可复制性
   - 使用Pool时要及时Put归还对象
*/

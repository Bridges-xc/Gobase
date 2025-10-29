// ============================= 1. sync.Once 使用 ====================
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

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

func main() {

	poolDemo()

	fmt.Println("\n=== 所有sync示例执行完成 ===")
}

// ============================= 1. 迭代器基础概念 ====================
// Go 1.23+ 支持 range over func，允许创建自定义迭代器
// 迭代器是一个函数，它将序列中的元素逐个传递给回调函数 yield

package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

func Fibonaccii(n int) func(yield func(int) bool) {
	a, b, c := 0, 1, 1
	return func(yield func(int) bool) {
		for range n {
			if !yield(a) {
				return
			}
			a, b = b, c
			c = a + b
		}
	}
} /**/

// ============================= 2. 推送式迭代器示例 ====================
// 推送式迭代器：由迭代器控制迭代逻辑，主动将元素推送给调用者

// Fibonacci 返回一个斐波那契数列迭代器
func Fibonacci(n int) iter.Seq[int] {
	a, b := 0, 1
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

// ============================= 3. 迭代器使用方式 ====================
func main() {
	fmt.Println("=== 基本迭代器使用 ===")

	// 3.1 基本迭代器使用
	fmt.Println("斐波那契数列:")
	for f := range Fibonacci(8) {
		fmt.Println(f)
	}

	// 3.2 切片迭代器
	fmt.Println("\n=== 切片迭代器 ===")
	numbers := []int{1, 2, 3, 4, 5}

	// slices.All - 带索引的迭代
	fmt.Println("带索引迭代:")
	for i, v := range slices.All(numbers) {
		fmt.Printf("索引:%d 值:%d\n", i, v)
	}

	// slices.Values - 仅值的迭代
	fmt.Println("仅值迭代:")
	for v := range slices.Values(numbers) {
		fmt.Println(v)
	}

	// 3.3 Map迭代器
	fmt.Println("\n=== Map迭代器 ===")
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println("Map键值对:")
	for k, v := range maps.All(data) {
		fmt.Printf("键:%s 值:%d\n", k, v)
	}

	fmt.Println("Map键集合:")
	keys := slices.Collect(maps.Keys(data))
	fmt.Println(keys)

	// ============================= 4. 拉取式迭代器 ====================
	// 拉取式迭代器：由调用者控制迭代逻辑，主动获取元素

	fmt.Println("\n=== 拉取式迭代器 ===")
	next, stop := iter.Pull(Fibonacci(5))
	defer stop() // 确保迭代器正确停止

	for {
		value, ok := next() // 主动获取下一个值
		if !ok {
			break // 没有更多值时退出
		}
		fmt.Println("拉取值:", value)
	}

	// ============================= 5. 错误处理迭代器 ====================
	fmt.Println("\n=== 带错误处理的迭代器 ===")

	// 模拟带错误的迭代器
	seqWithError := func(yield func(int, error) bool) {
		for i := 0; i < 3; i++ {
			var err error
			if i == 1 {
				err = fmt.Errorf("模拟错误")
			}
			if !yield(i, err) {
				return
			}
		}
	}

	for value, err := range seqWithError {
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			continue
		}
		fmt.Printf("值: %d\n", value)
	}

	// ============================= 6. 数据流处理示例 ====================
	fmt.Println("\n=== 数据流处理 ===")

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 过滤偶数并乘以2
	result := slices.Collect(
		iter.Seq[int](func(yield func(int) bool) {
			for v := range slices.Values(numbers) {
				if v%2 == 0 { // 过滤偶数
					if !yield(v * 2) { // 乘以2
						return
					}
				}
			}
		}),
	)

	fmt.Println("过滤偶数并乘以2:", result)

	// ============================= 7. 分块处理示例 ====================
	fmt.Println("\n=== 分块处理 ===")

	for chunk := range slices.Chunk(numbers, 3) {
		fmt.Println("分块:", chunk)
	}
}

// ============================= 8. 自定义迭代器类型 ====================
// 创建更复杂的迭代器类型

type CustomIterator[T any] struct {
	data []T
}

func NewCustomIterator[T any](data []T) *CustomIterator[T] {
	return &CustomIterator[T]{data: data}
}

// Seq 返回标准迭代器
func (ci *CustomIterator[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range ci.data {
			if !yield(item) {
				return
			}
		}
	}
}

// Filter 过滤元素
func (ci *CustomIterator[T]) Filter(predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range ci.data {
			if predicate(item) {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// ============================= 9. 使用自定义迭代器 ====================
func customIteratorDemo() {
	fmt.Println("\n=== 自定义迭代器 ===")

	data := []string{"apple", "banana", "cherry", "date"}
	iterator := NewCustomIterator(data)

	fmt.Println("所有元素:")
	for item := range iterator.Seq() {
		fmt.Println(item)
	}

	fmt.Println("过滤长度>5的元素:")
	for item := range iterator.Filter(func(s string) bool {
		return len(s) > 5
	}) {
		fmt.Println(item)
	}
}

// ============================= 总结知识点 ====================
/*
核心知识点总结:

1. 迭代器基础
   - Go 1.23+ 支持 range over func
   - 迭代器是接受 yield 回调的函数
   - 支持 0-2 个返回值 (Seq, Seq2)

2. 推送式迭代器 (推荐使用)
   - 迭代器控制逻辑，主动推送数据
   - 性能接近原生 for 循环
   - 语法: for v := range iterator()

3. 拉取式迭代器 (特殊场景使用)
   - 调用者通过 next() 主动拉取数据
   - 需要手动调用 stop() 释放资源
   - 性能开销较大

4. 标准库支持
   - slices.All/Values/Chunk/Collect
   - maps.Keys/Values/All/Collect
   - 提供丰富的数据处理功能

5. 错误处理模式
   - 通过多返回值传递错误
   - 在循环体内检查和处理错误
   - 支持错误后继续迭代

6. 数据流处理
   - 通过组合迭代器实现复杂处理
   - 支持过滤、映射、分块等操作
   - 可实现链式数据处理管道

7. 自定义迭代器
   - 封装复杂迭代逻辑
   - 支持链式操作方法
   - 提高代码复用性

8. 性能考虑
   - 推送式: 高性能，推荐使用
   - 拉取式: 低性能，特殊需求使用
   - 数据量大时注意内存使用

最佳实践:
1. 优先使用推送式迭代器
2. 及时处理迭代错误
3. 大数据集使用分块处理
4. 复杂逻辑封装为自定义迭代器
5. 注意迭代器的一次性使用特性
*/

// 注意: 需要 Go 1.23+ 版本编译运行

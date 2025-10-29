// ============================= 1. 泛型基础概念 ====================
// 泛型(参数化多态)让代码可以处理不同类型的数据，提高复用性和灵活性
// Go 1.18+ 支持泛型，采用Gcshape stenciling方案(单态化+字典的折中方案)

package main

import (
	"cmp"
	"fmt"
	"reflect"
	"sync"
)

// ============================= 2. 泛型函数 ====================
// 2.1 基础泛型函数示例
func Sum[T int | float64](a, b T) T {
	return a + b
}

// 2.2 类型约束使用接口
type Number interface {
	int | int32 | int64 | float32 | float64
}

func Multiply[T Number](a, b T) T {
	return a * b
}

// ============================= 3. 泛型类型 ====================
// 3.1 泛型切片
type GenericSlice[T int | string] []T

// 3.2 泛型映射 - 使用comparable约束键类型
type GenericMap[K comparable, V any] map[K]V

// 3.3 泛型结构体
type GenericStruct[T any] struct {
	Name string
	Data T
}

// 3.4 泛型方法(receiver使用泛型)
func (gs GenericStruct[T]) GetData() T {
	return gs.Data
}

// ============================= 4. 类型约束和类型集 ====================
// 4.1 类型集(接口作为类型约束)
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 // ~表示底层类型
}

type Integer interface {
	SignedInt | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// 4.2 使用类型约束
func ProcessNumbers[T Integer](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

// ============================= 5. 泛型接口 ====================
// 5.1 泛型接口定义
type Stringer[T any] interface {
	String() T
}

// 5.2 实现泛型接口
type Person struct {
	Name string
}

func (p Person) String() string {
	return "Person: " + p.Name
}

// 5.3 使用泛型接口约束
func PrintString[T Stringer[string]](s T) {
	fmt.Println(s.String())
}

// ============================= 6. 泛型数据结构示例 ====================
// 6.1 泛型队列
type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.items) == 0 {
		var zero T // 泛型零值
		return zero
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// 6.2 泛型对象池
type Pool[T any] struct {
	pool sync.Pool
}

func NewPool[T any](newFn func() T) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() interface{} { return newFn() },
		},
	}
}

func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

func (p *Pool[T]) Put(item T) {
	p.pool.Put(item)
}

// ============================= 7. 高级特性 ====================
// 7.1 类型推断示例
func inferTypes() {
	// 编译器自动推断类型
	result1 := Sum(1, 2)     // 推断为int
	result2 := Sum(1.1, 2.2) // 推断为float64

	fmt.Printf("Inferred int: %T, float64: %T\n", result1, result2)
}

// 7.2 泛型与反射结合
func TypeInfo[T any]() {
	var t T
	fmt.Printf("Type: %T, Kind: %v\n", t, reflect.TypeOf(t).Kind())
}

// ============================= 8. 实际应用示例 ====================
// 8.1 比较器函数类型
type Comparator[T any] func(a, b T) int

// 8.2 使用比较器的泛型函数
func FindMax[T any](items []T, compare Comparator[T]) T {
	if len(items) == 0 {
		var zero T
		return zero
	}

	max := items[0]
	for _, item := range items[1:] {
		if compare(item, max) > 0 {
			max = item
		}
	}
	return max
}

func main() {
	fmt.Println("=== 泛型函数演示 ===")
	fmt.Printf("Sum(int): %d\n", Sum(10, 20))
	fmt.Printf("Sum(float): %.2f\n", Sum(3.14, 2.71))

	fmt.Println("\n=== 泛型类型演示 ===")
	slice := GenericSlice[int]{1, 2, 3}
	mapp := GenericMap[string, int]{"a": 1, "b": 2}
	structt := GenericStruct[string]{Name: "test", Data: "hello"}

	fmt.Printf("Slice: %v, Map: %v, Struct: %+v\n", slice, mapp, structt)

	fmt.Println("\n=== 类型约束演示 ===")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of numbers: %d\n", ProcessNumbers(numbers))

	fmt.Println("\n=== 泛型接口演示 ===")
	person := Person{Name: "Alice"}
	PrintString(person)

	fmt.Println("\n=== 泛型数据结构演示 ===")
	queue := Queue[int]{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	fmt.Printf("Dequeue: %d\n", queue.Dequeue())

	fmt.Println("\n=== 类型推断演示 ===")
	inferTypes()

	fmt.Println("\n=== 实际应用演示 ===")
	ints := []int{3, 1, 4, 1, 5, 9}
	maxInt := FindMax(ints, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	fmt.Printf("Max integer: %d\n", maxInt)

	strings := []string{"apple", "banana", "cherry"}
	maxString := FindMax(strings, func(a, b string) int {
		return cmp.Compare(a, b)
	})
	fmt.Printf("Max string: %s\n", maxString)

	fmt.Println("\n=== 类型信息演示 ===")
	TypeInfo[int]()
	TypeInfo[string]()
}

// ============================= 总结知识点 ====================
// 1. 泛型基础: 类型参数[T], 类型约束[int|float64], 类型实参
// 2. 类型集: 使用接口定义类型约束, ~表示底层类型
// 3. 泛型类型: 切片、映射、结构体都可以是泛型的
// 4. 泛型接口: 接口也可以有类型参数
// 5. 类型推断: 编译器可以自动推断类型实参
// 6. 泛型数据结构: 队列、对象池等数据结构的泛型实现
// 7. 限制: 匿名结构体/函数不支持泛型, 不能有泛型方法
// 8. 最佳实践: 合理使用类型约束, 避免过度泛型化

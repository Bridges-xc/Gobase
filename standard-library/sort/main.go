package main

import (
	"fmt"
	"sort"
)

// ============================= 1. 基本类型排序 =============================
func basicSorting() {
	fmt.Println("=== 基本类型排序 ===")

	// 1.1 整型排序
	ints := []int{1, 2, 3, 111, 5, 99, 23, 5, 66}
	sort.Ints(ints)
	fmt.Printf("整型排序: %v\n", ints)

	// 1.2 浮点型排序
	floats := []float64{1.0, 2.5, 3.8, 1.11, 5.5, 99.99999, 23.9999, 5.66, 66}
	sort.Float64s(floats)
	fmt.Printf("浮点排序: %v\n", floats)

	// 1.3 字符串排序
	strings := []string{"helloworld", "aaa", "bbb", "ccc", "apple", "banana"}
	sort.Strings(strings)
	fmt.Printf("字符串排序: %v\n", strings)
}

// ============================= 2. 逆向排序 =============================

func reverseSorting() {
	fmt.Println("\n=== 逆向排序 ===")

	floats := []float64{1.0, 2.5, 3.8, 1.11, 5.5, 99.99999, 23.9999, 5.66, 66}

	// 2.1 使用sort.Reverse进行逆向排序
	sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
	fmt.Printf("逆向排序: %v\n", floats)
	// 2.2 字符串逆向排序
	strings := []string{"helloworld", "aaa", "bbb", "ccc", "apple", "banana"}
	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	fmt.Printf("字符串逆向: %v\n", strings)
}

// ============================= 3. 自定义结构体排序 =============================

// 3.1 定义Person结构体
type Person struct {
	UserId   string
	Username string
	Age      int
	Address  string
}

// 3.2 为Person切片定义类型，用于实现sort.Interface接口
type PersonSlice []Person

// 3.3 实现sort.Interface接口的三个方法
// Len() - 返回切片长度
func (p PersonSlice) Len() int {
	return len(p)
}

// Less() - 定义比较规则（按年龄升序）
func (p PersonSlice) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

// Swap() - 定义交换规则
func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// ============================= 4. 自定义排序实践 =============================

func customSorting() {
	fmt.Println("\n=== 自定义结构体排序 ===")

	persons := []Person{
		{
			UserId:   "1",
			Username: "wyh",
			Age:      18,
			Address:  "us",
		}, {
			UserId:   "2",
			Username: "jack",
			Age:      17,
			Address:  "ch",
		}, {
			UserId:   "3",
			Username: "mike",
			Age:      15,
			Address:  "india",
		},
	}
	fmt.Println("排序前:")
	for _, p := range persons {
		fmt.Printf("  %s(年龄:%d) ", p.Username, p.Age)
	}
	fmt.Println()
	// 4.1 进行自定义排序
	sort.Sort(PersonSlice(persons))

	fmt.Println("按年龄排序后:")
	for _, p := range persons {
		fmt.Printf("  %s(年龄:%d) ", p.Username, p.Age)
	}
	fmt.Println()
}

// ============================= 5. 切片有序性检查 =============================

func checkSorted() {
	fmt.Println("\n=== 切片有序性检查 ===")

	persons := []Person{
		{UserId: "1", Username: "wyh", Age: 15, Address: "us"},
		{UserId: "2", Username: "jack", Age: 17, Address: "ch"},
		{UserId: "3", Username: "mike", Age: 18, Address: "india"},
	}
	// 5.1 检查是否已排序（不会实际排序，只是检查）
	isSorted := sort.IsSorted(PersonSlice(persons))
	fmt.Printf("切片是否已按年龄排序: %t\n", isSorted)
	// 5.2 打乱顺序后再次检查
	persons[0], persons[1] = persons[1], persons[0]
	isSorted = sort.IsSorted(PersonSlice(persons))
	fmt.Printf("打乱后是否有序: %t\n", isSorted)
}

// ============================= 6. 使用sort.Slice进行快速自定义排序 ============================

func quickCustomSort() {
	fmt.Println("\n=== 快速自定义排序 (sort.Slice) ===")

	persons := []Person{
		{UserId: "1", Username: "wyh", Age: 18, Address: "us"},
		{UserId: "2", Username: "jack", Age: 17, Address: "ch"},
		{UserId: "3", Username: "mike", Age: 15, Address: "india"},
	}

	// 6.1 按用户名排序（不需要实现完整接口）
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Username < persons[j].Username
	})

	fmt.Println("按用户名排序:")
	for _, p := range persons {
		fmt.Printf("  %s ", p.Username)
	}
	fmt.Println()
}

// ============================= 主函数 =============================

func main() {
	// 执行所有排序示例
	basicSorting()
	reverseSorting()
	customSorting()
	checkSorted()
	quickCustomSort()

	fmt.Println("\n=== 学习总结 ===")
	fmt.Println("1. sort包为基本类型(Ints/Float64s/Strings)提供了开箱即用的排序方法")
	fmt.Println("2. 使用sort.Reverse可以实现逆向排序")
	fmt.Println("3. 自定义结构体排序需要实现Len()、Less()、Swap()三个方法")
	fmt.Println("4. sort.IsSorted可以检查切片是否已排序而不实际排序")
	fmt.Println("5. sort.Slice提供了更灵活的临时自定义排序方式")
}

/*=== 学习总结 ===
1. 基本类型: sort.Ints/Float64s/Strings 开箱即用
2. 逆向排序: 用 sort.Reverse 包装原切片
3. 自定义排序: 实现Len/Less/Swap接口 或 用sort.Slice快速排序
4. 有序检查: sort.IsSorted 或 XxxAreSorted 方法
5. 核心思想: 通过实现接口来扩展排序能力` + "\n")*/

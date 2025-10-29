// ============================= Go Map 全面学习 ============================

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// ======================= 1. 初始化 =======================

	// 1.1 字面量初始化
	mp1 := map[int]string{
		0: "零",
		1: "一",
		2: "二",
	}

	mp2 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}

	// 1.2 make函数初始化（推荐指定容量）
	mp3 := make(map[string]int, 10)  // 容量10
	mp4 := make(map[string][]int, 5) // 值为切片的map

	fmt.Println("初始化演示:")
	fmt.Printf("mp1: %v\n", mp1)
	fmt.Printf("mp2: %v\n", mp2)
	fmt.Printf("mp3: %v\n", mp3)
	fmt.Printf("mp4: %v\n", mp4)

	// ======================= 2. 访问操作 =======================

	fmt.Println("\n访问操作演示:")

	// 2.1 普通访问（不存在时返回零值）
	fmt.Printf("mp2['apple']: %d\n", mp2["apple"])
	fmt.Printf("mp2['not_exist']: %d (零值)\n", mp2["not_exist"])

	// 2.2 安全访问（检查key是否存在）
	if value, exists := mp2["banana"]; exists {
		fmt.Printf("键'banana'存在，值为: %d\n", value)
	} else {
		fmt.Println("键'banana'不存在")
	}

	// 2.3 获取map长度
	fmt.Printf("mp2的长度: %d\n", len(mp2))

	// ======================= 3. 存值操作 =======================

	fmt.Println("\n存值操作演示:")

	// 3.1 添加新键值对
	mp2["grape"] = 12
	fmt.Printf("添加grape后: %v\n", mp2)

	// 3.2 更新已存在的键
	mp2["apple"] = 10
	fmt.Printf("更新apple后: %v\n", mp2)

	// 3.3 NaN作为键的特殊情况（不推荐使用）
	nanMap := make(map[float64]string)
	nanMap[math.NaN()] = "first"
	nanMap[math.NaN()] = "second" // 不会覆盖，而是新增！
	fmt.Printf("NaN键的map: %v\n", nanMap)

	// ======================= 4. 删除操作 =======================

	fmt.Println("\n删除操作演示:")

	// 4.1 删除存在的键
	delete(mp2, "orange")
	fmt.Printf("删除orange后: %v\n", mp2)

	// 4.2 删除不存在的键（不会报错）
	delete(mp2, "not_exist")
	fmt.Printf("删除不存在的键后: %v\n", mp2)

	// 4.3 NaN键无法正常删除
	delete(nanMap, math.NaN())
	fmt.Printf("尝试删除NaN键后: %v (依然存在!)\n", nanMap)

	// ======================= 5. 遍历操作 =======================

	fmt.Println("\n遍历操作演示:")

	// 5.1 普通遍历（无序）
	fmt.Println("mp2遍历结果:")
	for key, value := range mp2 {
		fmt.Printf("  %s: %d\n", key, value)
	}

	// 5.2 遍历NaN键的map
	fmt.Println("NaN map遍历:")
	for key, value := range nanMap {
		fmt.Printf("  %f: %s\n", key, value)
	}

	// ======================= 6. 清空操作 =======================

	fmt.Println("\n清空操作演示:")

	// 6.1 使用clear函数（Go 1.21+）
	tempMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("清空前: %v\n", tempMap)
	clear(tempMap)
	fmt.Printf("清空后: %v\n", tempMap)

	// 6.2 传统清空方法（Go 1.21之前）
	tempMap2 := map[string]int{"x": 10, "y": 20, "z": 30}
	for key := range tempMap2 {
		delete(tempMap2, key)
	}
	fmt.Printf("传统方法清空后: %v\n", tempMap2)

	// ======================= 7. Set实现 =======================

	fmt.Println("\nSet实现演示:")

	// 7.1 使用map[type]struct{}实现Set
	set := make(map[int]struct{}, 10)

	// 添加元素
	for i := 0; i < 5; i++ {
		set[rand.Intn(100)] = struct{}{}
	}

	fmt.Printf("Set内容: %v\n", set)

	// 检查元素是否存在
	if _, exists := set[50]; exists {
		fmt.Println("50在Set中")
	} else {
		fmt.Println("50不在Set中")
	}

	// ======================= 8. 并发安全提示 =======================

	fmt.Println("\n并发安全提示:")
	fmt.Println("Map不是并发安全的！并发读写会触发fatal error")
	fmt.Println("高并发场景请使用sync.Map")

	// 演示并发不安全（实际运行可能触发panic）
	demoConcurrentIssue()
}

// 演示并发问题（运行时可能触发panic）
func demoConcurrentIssue() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到panic: %v\n", r)
			fmt.Println("这证明了map在并发读写时是不安全的！")
		}
	}()

	mp := make(map[int]int)

	// 启动goroutine进行写操作
	go func() {
		for i := 0; i < 1000; i++ {
			mp[i] = i
		}
	}()

	// 同时进行读操作（可能触发并发问题）
	for i := 0; i < 1000; i++ {
		_ = mp[i]
	}
}

// ============================= 总结知识点 ============================
/*
1. 初始化:
   - 字面量: map[K]V{k1:v1, k2:v2}
   - make函数: make(map[K]V, capacity)
   - 必须初始化后才能存值

2. 访问:
   - mp[key] 返回值，不存在时返回零值
   - val, exists := mp[key] 安全访问
   - len(mp) 获取元素个数

3. 存值:
   - mp[key] = value
   - 已存在的key会覆盖原值
   - 避免使用NaN作为key（行为异常）

4. 删除:
   - delete(mp, key)
   - 不存在的key不会报错
   - NaN key无法正常删除

5. 遍历:
   - for k, v := range mp
   - 遍历顺序不确定（哈希表特性）
   - 可以遍历到NaN键

6. 清空:
   - clear(mp) (Go 1.21+)
   - 传统方法: 遍历delete

7. Set实现:
   - map[T]struct{} 实现Set
   - 空结构体不占内存

8. 并发安全:
   - Map不是并发安全的
   - 并发读写会触发fatal error
   - 高并发用sync.Map

9. 键类型要求:
   - 必须是可比较类型
   - 不能是slice, map, function等不可比较类型
*/

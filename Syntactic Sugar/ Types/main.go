// ============================= 1. Go类型系统概述 ====================
// Go是静态强类型语言：类型在编译期确定，生命周期内不变
// 强类型：严格类型检查，不自动类型转换

package main

import (
	"fmt"
	"unsafe"
)

// ============================= 2. 静态强类型特性演示 ====================
func demonstrateStaticTyping() {
	// 2.1 类型一旦确定不能改变
	var a int = 64
	// a = "64" // 编译错误：不能将字符串赋给int变量

	// 2.2 强类型检查，不自动转换
	// fmt.Println(1 + "1") // 编译错误：类型不匹配

	fmt.Println("静态强类型验证通过", a)
}

// ============================= 3. 类型声明（创建新类型） ====================
// 3.1 基础类型声明
type MyInt int64
type MyFloat64 float64
type MyMap map[string]int

// 3.2 使用示例
func demonstrateTypeDeclaration() {
	var num1 MyInt = 100
	var num2 int64 = 200

	// num1 + num2 // 编译错误：不同类型不能运算
	fmt.Printf("MyInt值: %d, 基础类型: %T\n", num1, num1)
	fmt.Printf("int64值: %d, 类型: %T\n", num2, num2)
}

// ============================= 4. 类型别名（同类型的别名） ====================
// 4.1 类型别名定义
type IntAlias = int
type TwoDMap = map[string]map[string]int

// 4.2 使用示例
func demonstrateTypeAlias() {
	var a IntAlias = 10
	var b int = 20

	// 类型别名可以互相运算
	result := a + b
	fmt.Printf("类型别名运算结果: %d\n", result)

	// 复杂类型使用别名
	complexMap := TwoDMap{
		"a": {"x": 1, "y": 2},
		"b": {"z": 3},
	}
	fmt.Printf("二维Map: %v\n", complexMap)
}

// ============================= 5. 类型转换（显式转换） ====================
func demonstrateTypeConversion() {
	// 5.1 基础类型转换
	var myFloat MyFloat64 = 3.14
	var stdFloat float64 = 2.71

	// 必须显式转换
	converted := float64(myFloat) + stdFloat
	fmt.Printf("类型转换后相加: %.2f\n", converted)

	// 5.2 数值类型转换（注意溢出）
	var small int8 = 100
	var large int32 = 50000

	smallToLarge := int32(small) // 安全转换
	largeToSmall := int8(large)  // 可能溢出

	fmt.Printf("小转大: %d -> %d\n", small, smallToLarge)
	fmt.Printf("大转小(溢出): %d -> %d\n", large, largeToSmall)

	// 5.3 避免歧义的转换
	var x int = 42
	// 明确的转换语法
	pointer := (*int)(unsafe.Pointer(&x))
	fmt.Printf("指针转换: %v -> %v\n", x, *pointer)
}

// ============================= 6. 类型断言（接口类型检查） ====================
func demonstrateTypeAssertion() {
	// 6.1 基础类型断言
	var value interface{} = "hello world"

	// 安全断言方式
	if str, ok := value.(string); ok {
		fmt.Printf("类型断言成功: %s\n", str)
	} else {
		fmt.Println("类型断言失败")
	}

	// 6.2 多类型断言
	var num interface{} = 42

	switch v := num.(type) {
	case int:
		fmt.Printf("整数类型: %d\n", v)
	case string:
		fmt.Printf("字符串类型: %s\n", v)
	case float64:
		fmt.Printf("浮点数类型: %.2f\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

// ============================= 7. 类型判断（type switch） ====================
func demonstrateTypeSwitch() {
	// 7.1 处理多种类型
	processType := func(value interface{}) {
		switch value.(type) {
		case int:
			fmt.Println("处理整数类型")
		case float64:
			fmt.Println("处理浮点数类型")
		case string:
			fmt.Println("处理字符串类型")
		case bool:
			fmt.Println("处理布尔类型")
		default:
			fmt.Println("处理其他类型")
		}
	}

	processType(100)
	processType(3.14)
	processType("text")
	processType(true)
}

// ============================= 8. 实际应用示例 ====================
// 8.1 使用类型别名简化复杂函数签名
type Processor = func(string) (int, error)

func registerProcessor(name string, processor Processor) {
	fmt.Printf("注册处理器: %s\n", name)
}

// 8.2 类型安全的配置处理
type ConfigValue interface {
	string | int | bool
}

func GetConfig[T ConfigValue](key string, defaultValue T) T {
	// 模拟配置获取
	return defaultValue
}

func main() {
	fmt.Println("=== Go类型系统演示 ===\n")

	fmt.Println("1. 静态强类型演示:")
	demonstrateStaticTyping()

	fmt.Println("\n2. 类型声明演示:")
	demonstrateTypeDeclaration()

	fmt.Println("\n3. 类型别名演示:")
	demonstrateTypeAlias()

	fmt.Println("\n4. 类型转换演示:")
	demonstrateTypeConversion()

	fmt.Println("\n5. 类型断言演示:")
	demonstrateTypeAssertion()

	fmt.Println("\n6. 类型判断演示:")
	demonstrateTypeSwitch()

	fmt.Println("\n7. 实际应用演示:")
	registerProcessor("test", func(s string) (int, error) {
		return len(s), nil
	})

	strConfig := GetConfig("app.name", "default")
	intConfig := GetConfig("app.port", 8080)
	fmt.Printf("配置获取 - 字符串: %s, 整数: %d\n", strConfig, intConfig)

	fmt.Println("\n=== 演示完成 ===")
}

// ============================= 总结知识点 ====================
// 1. 静态强类型: 编译期确定类型，严格类型检查，无隐式转换
// 2. 类型后置: 变量名在前，类型在后，提高可读性
// 3. 类型声明(type): 创建新类型，与原类型不同，不能互相运算
// 4. 类型别名(=): 同类型的别名，可以互相运算，用于简化复杂类型
// 5. 类型转换: 必须显式转换，注意数值溢出问题
// 6. 类型断言: 用于接口类型检查，安全方式使用ok判断
// 7. 类型判断: 使用switch v.(type)处理多种类型情况
// 8. 实际应用: 类型别名简化代码，泛型约束确保类型安全

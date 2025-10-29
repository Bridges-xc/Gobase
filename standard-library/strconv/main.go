package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Go strconv 包使用指南 ===\n")

	// ============================= 1. 字符串与整型互转 =============================

	fmt.Println("1. 字符串与整型互转:")

	// 1.1 字符串转整型 - Atoi (ASCII to Integer)
	strNum := "456"
	intVal, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println("Atoi失败: %v\n", err)
	} else {
		fmt.Printf("   Atoi: '%s' → %d\n", strNum, intVal)
	}

	// 1.2 整型转字符串 - Itoa (Integer to ASCII)
	num := 114
	strVal := strconv.Itoa(num)
	fmt.Printf("   Itoa: %d → '%s'\n", num, strVal)
	// ============================= 2. 字符串与布尔值互转 =============================

	fmt.Println("\n2. 字符串与布尔值互转:")
	// 2.1 字符串转布尔值 - ParseBool
	// 支持: "1", "t", "T", "true", "TRUE", "True" → true
	//       "0", "f", "F", "false", "FALSE", "False" → false
	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Printf("   ParseBool失败: %v\n", err)
	}

	// 2.2 布尔值转字符串 - FormatBool
	trueStr := strconv.FormatBool(boolVal)
	falseStr := strconv.FormatBool(false)
	fmt.Println("FormatBool: true → '%s', false → '%s'\n", trueStr, falseStr)

	// ============================= 3. 字符串与浮点数互转 =============================
	fmt.Println("\n3. 字符串与浮点数互转:")

	// 3.1 字符串转浮点数 - ParseFloat
	// bitSize: 32(float32) 或 64(float64)
	floatStr := "3.1415926"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Printf("   ParseFloat失败: %v\n", err)
	} else {
		fmt.Printf("   ParseFloat: '%s' → %.6f\n", floatStr, floatVal)
	} // 3.2 浮点数转字符串 - FormatFloat
	// fmt参数说明:
	// 'f' - 普通小数格式 (-ddd.dddd) - 最常用
	// 'e' - 科学计数法 (-d.dddde±dd)
	// 'g' - 自动选择 'e' 或 'f' 格式
	pi := 3.141592653589793
	fmt.Printf("   FormatFloat-f: %.6f → '%s'\n", pi, strconv.FormatFloat(pi, 'f', 4, 64))
	fmt.Printf("   FormatFloat-e: %.6f → '%s'\n", pi, strconv.FormatFloat(pi, 'e', 4, 64))
	fmt.Printf("   FormatFloat-g: %.6f → '%s'\n", pi, strconv.FormatFloat(pi, 'g', 4, 64))

	// ============================= 4. 字符串与复数互转 =============================

	fmt.Println("\n4. 字符串与复数互转:")

	// 4.1 字符串转复数 - ParseComplex
	// 注意: 必须使用 'i' 作为虚数单位，不能用 'j'
	complexStr := "2.5+3.7i"
	complexVal, err := strconv.ParseComplex(complexStr, 128)
	if err != nil {
		fmt.Printf("   ParseComplex失败: %v\n", err)
	} else {
		fmt.Printf("   ParseComplex: '%s' → %v\n", complexStr, complexVal)
	}
	// 4.2 复数转字符串 - FormatComplex
	c := complex(1.5, 2.8)
	fmt.Printf("   FormatComplex: %v → '%s'\n", c, strconv.FormatComplex(c, 'f', 2, 128))
	// ============================= 5. 字符串引用和转义 =============================

	fmt.Println("\n5. 字符串引用和转义:")

	// 5.1 转换为带引号的Go字符串 - Quote
	text := "hello 世界"
	quoted := strconv.Quote(text)
	fmt.Printf("   Quote: '%s' → %s\n", text, quoted)
	// 5.2 转换为ASCII转义的Go字符串 - QuoteToASCII
	asciiQuoted := strconv.QuoteToASCII(text)
	fmt.Printf("   QuoteToASCII: '%s' → %s\n", text, asciiQuoted)

	// ============================= 6. 追加数据到字节切片 =============================

	fmt.Println("\n6. 追加数据到字节切片:")

	// 6.1 Append系列函数 - 用于构建复杂字符串
	// 在Go中不能直接用"+"连接不同类型，需要用Append函数
	base := []byte("数据")
	base = strconv.AppendInt(base, 100, 10)
	base = strconv.AppendFloat(base, 2.718, 'f', 3, 64)
	base = strconv.AppendBool(base, true)
	base = strconv.AppendQuote(base, "appended")
	fmt.Printf("   Append结果: %s\n", string(base))
	// ============================= 7. 其他实用函数 =============================

	fmt.Println("\n7. 其他实用函数:")

	// 7.1 判断字符是否可打印 - IsPrint
	char := 'A'
	fmt.Printf("   IsPrint('A'): %t\n", strconv.IsPrint(char))
	// 7.2 判断字符串是否可以不用转义表示 - CanBackquote
	testStr := "Hello\tWorld"
	fmt.Printf("   CanBackquote('Hello\\tWorld'): %t\n", strconv.CanBackquote(testStr))
	// ============================= 知识点总结 =============================

	fmt.Println("\n=== 知识点总结 ===")
	fmt.Println("1. Atoi/Itoa: 字符串与整型互转，最常用")
	fmt.Println("2. ParseBool/FormatBool: 布尔值转换，支持多种真值表示")
	fmt.Println("3. ParseFloat/FormatFloat: 浮点数转换，注意格式化参数选择")
	fmt.Println("4. ParseComplex/FormatComplex: 复数转换，虚数单位必须用'i'")
	fmt.Println("5. Quote/QuoteToASCII: 字符串转义，用于生成安全字符串")
	fmt.Println("6. Append系列: 高效构建字节切片，避免字符串连接性能问题")
	fmt.Println("7. 所有Parse函数都返回error，必须进行错误处理")
	fmt.Println("8. Format函数参数复杂但功能强大，根据需要选择合适的格式化选项")

}

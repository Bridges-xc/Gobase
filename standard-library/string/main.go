package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

func main() {
	// ============================= 1. 字符串基础操作 ====================

	// 1.1 字符串克隆 - 创建新内存副本
	original := "hello 世界"
	cloned := strings.Clone(original)
	fmt.Printf("克隆: %s → %s\n", original, cloned)
	// 1.2 字符串比较 - 字典序比较
	fmt.Printf("比较: %d\n", strings.Compare("abc", "abe"))
	fmt.Printf("忽略大小写相等: %t\n", strings.EqualFold("Hello", "hELLO"))
	// ============================= 2. 字符串包含检查 ====================

	// 2.1 检查子串包含
	fmt.Printf("包含子串: %t\n", strings.Contains("abcdefg", "abc"))

	// 2.2 检查字符集中任意字符
	fmt.Printf("包含任意字符: %t\n", strings.ContainsAny("abcdefg", "xyz"))

	// 2.3 检查特定rune字符
	fmt.Printf("包含rune: %t\n", strings.ContainsRune("你好世界", '你'))

	// ============================= 3. 子串统计与切割 ====================

	// 3.1 统计子串出现次数
	fmt.Printf("出现次数: %d\n", strings.Count("hello world", "l"))

	// 3.2 切割字符串 - 删除首次出现的子串
	before, after, found := strings.Cut("Hello world", " ")
	fmt.Printf("切割结果: %q, %q, %t\n", before, after, found)

	// ============================= 4. 字符串分割 ====================

	// 4.1 按空格分割
	words := strings.Fields(" a b c d e f g ")
	fmt.Printf("空格分割: %q\n", words)

	// 4.2 按自定义函数分割
	commaSplit := strings.FieldsFunc("a,b,c,d", func(r rune) bool {
		return r == ','
	})
	fmt.Printf("逗号分割: %q\n", commaSplit)

	// 4.3 按指定分隔符分割
	fmt.Printf("分隔符分割: %q\n", strings.Split("this,is,go", ","))
	fmt.Printf("限制分割次数: %q\n", strings.SplitN("this,is,go", ",", 2))
	// SplitAfter 保留分隔符
	fmt.Printf("%q\n", strings.SplitAfter(str, ",")) // ["a," "b," "c," "d"]

	// ============================= 5. 字符串查找 ====================

	// 5.1 前后缀检查
	str := "hello world"
	fmt.Printf("前缀检查: %t\n", strings.HasPrefix(str, "hello"))
	fmt.Printf("后缀检查: %t\n", strings.HasSuffix(str, "world"))

	// 5.2 子串位置查找
	fmt.Printf("首次出现位置: %d\n", strings.Index("abcdefg", "cd"))
	fmt.Printf("最后出现位置: %d\n", strings.LastIndex("abcdeafg", "a"))
	// ============================= 6. 字符串转换 ====================

	// 6.1 大小写转换
	fmt.Printf("转小写: %s\n", strings.ToLower("HELLO World"))
	fmt.Printf("转大写: %s\n", strings.ToUpper("hello world"))

	// 6.2 字符映射转换
	mapped := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - 32 // 小写转大写
		}
		return r
	}, "hello World")
	fmt.Printf("映射转换: %s\n", mapped)
	// ============================= 7. 字符串修改 ====================

	// 7.1 重复字符串
	fmt.Printf("重复字符串: %s\n", strings.Repeat("abc", 3))

	// 7.2 替换字符串
	fmt.Printf("替换: %s\n", strings.Replace("hello world", "world", "golang", 1))
	fmt.Printf("全部替换: %s\n", strings.ReplaceAll("oooops", "o", "a"))

	// 7.3 修剪字符串
	fmt.Printf("修剪两端: %s\n", strings.Trim("!!hello!!", "!"))
	fmt.Printf("修剪前缀: %s\n", strings.TrimPrefix("!!hello!!", "!!"))
	fmt.Printf("修剪后缀: %s\n", strings.TrimSuffix("!!hello!!", "!!"))
	// ============================= 8. 字符串构建器 ====================

	// 8.1 使用Builder高效构建字符串
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	fmt.Printf("Builder结果: %s, 长度: %d\n", builder.String(), builder.Len())
	// ============================= 9. 字符串替换器 ====================

	// 9.1 创建多组替换规则的替换器
	replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Printf("HTML转义: %s\n", replacer.Replace("This is <b>HTML</b>!"))

	// ============================= 10. 字符串读取器 ====================

	// 10.1 将字符串作为io.Reader读取
	reader := strings.NewReader("abcdefghijk")
	buffer := make([]byte, 5)
	n, err := reader.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("读取字节数: %d, 内容: %s\n", n, string(buffer[:n]))

	// ============================= 11. 特殊语言处理 ====================

	// 11.1 土耳其语特殊大小写处理
	turkishWord := "Örnek İş"
	lowerTurkish := strings.ToLowerSpecial(unicode.TurkishCase, turkishWord)
	fmt.Printf("土耳其语小写: %s\n", lowerTurkish)
}

// ============================= 总结知识点 ====================
/*
1. 字符串克隆: Clone() 创建独立内存副本
2. 字符串比较: Compare() 字典序比较, EqualFold() 忽略大小写比较
3. 包含检查: Contains() 检查子串, ContainsAny() 检查字符集, ContainsRune() 检查rune
4. 统计切割: Count() 统计次数, Cut() 删除首次出现的子串
5. 字符串分割: Fields() 按空格分割, Split() 按分隔符分割
6. 位置查找: Index()/LastIndex() 查找子串位置, HasPrefix()/HasSuffix() 检查前后缀
7. 大小写转换: ToLower()/ToUpper() 基础转换, ToLowerSpecial() 特殊语言处理
8. 字符串修改: Replace() 替换, Trim() 修剪, Repeat() 重复
9. 高效构建: Builder 用于大量字符串拼接, 避免内存分配
10. 批量替换: Replacer 用于多组替换规则
11. 字符串读取: Reader 将字符串作为流读取
12. 字符映射: Map() 通过函数转换每个字符

注意: Go字符串不可变, 所有修改操作都返回新字符串
Builder不能值传递, 否则会panic
*/

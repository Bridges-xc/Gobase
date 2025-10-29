package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

// ============================= 1. 快速开始 ====================
func basicExample() {
	fmt.Println("=== 1. 快速开始 ===")

	// 最简单的模板示例
	tmpl := `Hello, {{ .name }}! Welcome to {{ .place }}.`

	// 1. 创建模板
	t, err := template.New("greeting").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// 2. 准备数据
	data := map[string]interface{}{
		"name":  "Alice",
		"place": "Go Template World",
	}

	// 3. 执行模板
	fmt.Print("输出: ")
	t.Execute(os.Stdout, data)
	fmt.Println("\n")
}

// ============================= 2. 模板参数 ====================
func templateParameters() {
	fmt.Println("=== 2. 模板参数 ===")

	// 不同类型的数据演示
	datas := []interface{}{
		"simple string", // 字符串
		42,              // 整数
		3.14159,         // 浮点数
		map[string]interface{}{ // map
			"Name": "Bob",
			"Age":  25,
		},
		struct { // 结构体
			Title string
			Count int
		}{"Products", 100},
		[]string{"apple", "banana", "cherry"}, // 切片
	}

	templates := []string{
		"字符串: {{ . }}",
		"数字: {{ . }}",
		"浮点数: {{ . }}",
		"Map数据 - 姓名: {{ .Name }}, 年龄: {{ .Age }}",
		"结构体 - 标题: {{ .Title }}, 数量: {{ .Count }}",
		"切片访问: {{ index . 1 }}", // 访问切片第二个元素
	}

	for i, data := range datas {
		t, _ := template.New("test").Parse(templates[i])
		fmt.Print("输出: ")
		t.Execute(os.Stdout, data)
		fmt.Println()
	}
	fmt.Println()
}

// ============================= 3. 注释和空白控制 ====================
func commentsAndWhitespace() {
	fmt.Println("=== 3. 注释和空白控制 ===")

	tmpl := `开始{{- /* 这个注释不会输出 */ -}}
{{- "前面空白被消除" -}} 
中间{{ .message }}后面
{{- "后面空白也被消除" -}}结束`

	t, _ := template.New("whitespace").Parse(tmpl)
	data := map[string]interface{}{"message": "重要信息"}

	fmt.Print("输出: ")
	t.Execute(os.Stdout, data)
	fmt.Println("\n")
}

// ============================= 4. 变量声明和使用 ====================
func variablesExample() {
	fmt.Println("=== 4. 变量声明和使用 ===")

	tmpl := `{{ $firstName := "John" }}{{ $lastName := "Doe" }}
全名: {{ $firstName }} {{ $lastName }}
{{ $age := 30 }}年龄: {{ $age }}
{{ $score := 95.5 }}分数: {{ $score }}`

	t, _ := template.New("variables").Parse(tmpl)

	fmt.Print("输出: ")
	t.Execute(os.Stdout, nil)
	fmt.Println("\n")
}

// ============================= 5. 内置函数使用 ====================
func builtinFunctions() {
	fmt.Println("=== 5. 内置函数使用 ===")

	tmpl := `比较函数:
eq(5,5): {{ eq 5 5 }}
ne(5,3): {{ ne 5 3 }}
lt(3,5): {{ lt 3 5 }}
gt(5,3): {{ gt 5 3 }}

逻辑函数:
and(true,false): {{ and true false }}
or(true,false): {{ or true false }}
not(true): {{ not true }}

其他函数:
长度: {{ len .Items }}
索引访问: {{ index .Items 1 }}
格式化: {{ printf "价格: $%.2f" .Price }}
管道操作: {{ .Name | printf "欢迎 %s" | upper }}`

	// 自定义函数映射
	funcMap := template.FuncMap{
		"upper": func(s string) string {
			// 简单实现，实际应该用 strings.ToUpper
			return s + " (UPPERCASE)"
		},
	}

	data := map[string]interface{}{
		"Items": []string{"苹果", "香蕉", "橙子"},
		"Price": 19.99,
		"Name":  "Alice",
	}

	t, _ := template.New("functions").Funcs(funcMap).Parse(tmpl)

	fmt.Print("输出: ")
	t.Execute(os.Stdout, data)
	fmt.Println("\n")
}

// ============================= 6. 条件和循环 ====================
func conditionAndLoop() {
	fmt.Println("=== 6. 条件和循环 ===")

	tmpl := `{{ if .User.IsVIP }}
🌟 尊贵的VIP用户 {{ .User.Name }}
{{ else if .User.IsMember }}
👤 会员用户 {{ .User.Name }}  
{{ else }}
🚶 普通用户 {{ .User.Name }}
{{ end }}

购物车商品:
{{ range $index, $item := .Cart }}
  {{ add $index 1 }}. {{ $item.Name }} - ${{ $item.Price }}
{{ else }}
  购物车为空
{{ end }}

总价: ${{ .Total }}`

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	data := map[string]interface{}{
		"User": map[string]interface{}{
			"Name":     "Bob",
			"IsVIP":    false,
			"IsMember": true,
		},
		"Cart": []map[string]interface{}{
			{"Name": "笔记本电脑", "Price": 999.99},
			{"Name": "鼠标", "Price": 29.99},
			{"Name": "键盘", "Price": 79.99},
		},
		"Total": 1109.97,
	}

	t, _ := template.New("conditionLoop").Funcs(funcMap).Parse(tmpl)

	fmt.Print("输出: ")
	t.Execute(os.Stdout, data)
	fmt.Println("\n")
}

// ============================= 7. with语句和作用域 ====================
func withStatement() {
	fmt.Println("=== 7. with语句和作用域 ===")

	tmpl := `外部用户: {{ .User.Name }}
{{ with .User.Address }}
  地址信息:
  城市: {{ .City }}
  街道: {{ .Street }}
{{ else }}
  无地址信息
{{ end }}
外部访问: {{ .User.Name }}`

	data := map[string]interface{}{
		"User": map[string]interface{}{
			"Name": "Charlie",
			"Address": map[string]interface{}{
				"City":   "北京",
				"Street": "长安街",
			},
		},
	}

	t, _ := template.New("with").Parse(tmpl)

	fmt.Print("输出: ")
	t.Execute(os.Stdout, data)
	fmt.Println("\n")
}

// ============================= 8. 嵌套模板定义 ====================
func nestedTemplates() {
	fmt.Println("=== 8. 嵌套模板定义 ===")

	mainTmpl := `网站页面
{{ template "header" . }}
主要内容: {{ .Content }}
{{ template "footer" . }}`

	// 定义嵌套模板
	fullTmpl := `{{ define "header" }}== 网站头部 =={{ end }}
{{ define "footer" }}== 版权信息 © 2024 =={{ end }}
` + mainTmpl

	data := map[string]interface{}{
		"Content": "欢迎访问我们的网站！",
	}

	t, _ := template.New("main").Parse(fullTmpl)

	fmt.Print("输出: ")
	t.Execute(os.Stdout, data)
	fmt.Println("\n")
}

// ============================= 9. 模板关联 ====================
func associatedTemplates() {
	fmt.Println("=== 9. 模板关联 ===")

	// 创建多个模板
	t1, _ := template.New("header").Parse(`== {{ .Title }} ==`)
	t2, _ := template.New("content").Parse(`内容: {{ .Body }}`)

	// 主模板
	t3, _ := template.New("page").Parse(`页面开始
{{ template "header" .Header }}
{{ template "content" .Content }}
页面结束`)

	// 关联模板
	t3.AddParseTree("header", t1.Tree)
	t3.AddParseTree("content", t2.Tree)

	data := map[string]interface{}{
		"Header": map[string]interface{}{
			"Title": "我的网站",
		},
		"Content": map[string]interface{}{
			"Body": "这是页面内容",
		},
	}

	fmt.Print("输出: ")
	t3.Execute(os.Stdout, data)
	fmt.Println("\n")
}

// ============================= 10. 插槽功能 ====================
func slotTemplates() {
	fmt.Println("=== 10. 插槽功能 ===")

	baseTmpl := `用户信息:
姓名: {{ .Name }}
年龄: {{ .Age }}
{{ block "extra" . }}默认额外信息{{ end }}`

	extendedTmpl := `{{ template "base" . }}
{{ define "extra" }}
职业: {{ .Job }}
公司: {{ .Company }}
{{ end }}`

	// 先解析基础模板
	base, _ := template.New("base").Parse(baseTmpl)

	// 再解析扩展模板并关联基础模板
	extended, _ := template.New("extended").Parse(extendedTmpl)
	extended.AddParseTree("base", base.Tree)

	data := map[string]interface{}{
		"Name":    "David",
		"Age":     28,
		"Job":     "工程师",
		"Company": "Tech Corp",
	}

	fmt.Print("输出: ")
	extended.Execute(os.Stdout, data)
	fmt.Println("\n")
}

//go:embed templates/*.txt
var templateFS embed.FS

// ============================= 11. 文件模板 ====================
func fileTemplates() {
	fmt.Println("=== 11. 文件模板 ===")

	// 从文件系统加载模板
	t, err := template.ParseFS(templateFS, "templates/*.txt")
	if err != nil {
		fmt.Printf("加载模板文件失败: %v\n", err)
		return
	}

	data := map[string]interface{}{
		"Name":    "Emma",
		"Age":     32,
		"Email":   "emma@example.com",
		"Country": "中国",
	}

	fmt.Print("输出: ")
	// 执行特定模板
	err = t.ExecuteTemplate(os.Stdout, "user_profile.txt", data)
	if err != nil {
		fmt.Printf("执行模板失败: %v\n", err)
	}
	fmt.Println()
}

// ============================= 12. 自定义函数高级用法 ====================
func advancedCustomFunctions() {
	fmt.Println("=== 12. 自定义函数高级用法 ===")

	funcMap := template.FuncMap{
		"join": func(sep string, items []string) string {
			result := ""
			for i, item := range items {
				if i > 0 {
					result += sep
				}
				result += item
			}
			return result
		},
		"multiply": func(a, b int) int {
			return a * b
		},
		"formatDate": func(date string) string {
			return "2024-01-01 (" + date + ")"
		},
	}

	tmpl := `商品列表: {{ join ", " .Products }}
总价: {{ multiply .Price .Quantity }}
日期: {{ formatDate .Date }}`

	data := map[string]interface{}{
		"Products": []string{"手机", "电脑", "平板"},
		"Price":    100,
		"Quantity": 3,
		"Date":     "今天",
	}

	t, _ := template.New("advanced").Funcs(funcMap).Parse(tmpl)

	fmt.Print("输出: ")
	t.Execute(os.Stdout, data)
	fmt.Println("\n")
}

func main() {
	fmt.Println("Go 模板引擎学习示例")
	fmt.Println("===================\n")

	// 按顺序执行各个示例
	basicExample()            // 1. 快速开始
	templateParameters()      // 2. 模板参数
	commentsAndWhitespace()   // 3. 注释和空白控制
	variablesExample()        // 4. 变量声明和使用
	builtinFunctions()        // 5. 内置函数使用
	conditionAndLoop()        // 6. 条件和循环
	withStatement()           // 7. with语句和作用域
	nestedTemplates()         // 8. 嵌套模板定义
	associatedTemplates()     // 9. 模板关联
	slotTemplates()           // 10. 插槽功能
	fileTemplates()           // 11. 文件模板
	advancedCustomFunctions() // 12. 自定义函数高级用法

	fmt.Println("=== 学习完成 ===")
}

// ============================= 总结知识点 ====================
/*
Go 模板引擎核心知识点总结:

1. 基础流程:
   - template.New() 创建模板
   - Parse() 解析模板字符串
   - Execute() 应用数据并输出

2. 模板语法:
   - {{ .Field }} 访问数据字段
   - {{ . }} 根对象
   - {{ index .Slice 0 }} 访问切片/数组
   - {{- 消除左侧空白 -}} 空白控制

3. 变量操作:
   - {{ $var := value }} 变量声明
   - {{ $var }} 变量使用
   - 作用域遵循代码块规则

4. 函数系统:
   - 内置函数: eq, len, index, printf 等
   - 自定义函数: Funcs() 注册，支持多返回值
   - 管道操作: {{ . | func1 | func2 }}

5. 流程控制:
   - {{ if }} {{ else }} {{ end }} 条件判断
   - {{ range }} {{ else }} {{ end }} 循环迭代
   - {{ with }} {{ else }} {{ end }} 作用域控制

6. 模板组织:
   - {{ define "name" }} 定义命名模板
   - {{ template "name" . }} 引用模板
   - AddParseTree() 关联外部模板
   - {{ block "slot" . }} 插槽机制

7. 文件操作:
   - ParseFiles() 解析文件
   - ParseGlob() 通配符匹配
   - ParseFS() 从嵌入文件系统加载

8. 最佳实践:
   - 复杂HTML使用 html/template 更安全
   - 错误处理很重要
   - 合理使用自定义函数减少模板复杂度

适用场景:
- 服务端渲染HTML页面
- 生成邮件模板
- 代码生成工具
- 配置文件模板
- 报告文档生成

通过这一页代码，你已经掌握了Go模板引擎的核心用法！
*/

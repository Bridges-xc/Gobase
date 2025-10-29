package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"gopkg.in/yaml.v3"
)

// ============================= 1. 通用结构体定义 =============================

// Person 结构体用于演示各种序列化格式
// 注意：字段必须首字母大写（对外暴露）才能被序列化

type Person struct {
	UserID   string `xml:"id" yaml:"user_id" json:"id"`           // XML标签为id, YAML为user_id, JSON为id
	Username string `xml:"name" yaml:"username" json:"name"`      // 不同格式使用不同的字段名
	Age      int    `xml:"age" yaml:"age" json:"age"`             // 年龄字段
	Address  string `xml:"address" yaml:"address" json:"address"` // 地址字段
}

// ============================= 2. JSON 序列化 =============================

// 2.1 JSON序列化示例
func demoJSON() {
	fmt.Println("\n--- 2.1 JSON 序列化---")
	fmt.Println("JSON是最常用的数据交换格式，轻量且易读")

	person := Person{
		UserID:   "123",
		Username: "jack",
		Age:      18,
	}
	// 紧凑格式序列化（网络传输用）
	compactJSON, _ := json.Marshal(person)
	fmt.Printf("紧凑JSON: %s\n", string(compactJSON))
	// 格式化序列化（人类可读）
	prettyJSON, _ := json.MarshalIndent(person, "", "  ")
	fmt.Printf("格式化JSON:\n%s\n", string(prettyJSON))
}

// 2.2 JSON反序列化示例
func demoJSONUnmarshal() {
	fmt.Println("\n--- 2.2 JSON 反序列化 ---")
	jsonStr := `{"id":"120","name":"jack","age":18,"address":"usa"}`
	var person Person

	// 将JSON字符串解析为Go结构体
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Println("json解析失败", err)
		return
	}
	fmt.Println("解析结果： %+v\n", person)
}

// ============================= 3. XML 序列化 =============================

// 3.1 XML序列化示例
func demoXML() {
	fmt.Println("\n--- 3.1 XML 序列化 ---")
	fmt.Println("XML是较老的数据格式，但仍在一些传统系统中使用")

	person := Person{
		UserID:   "120",
		Username: "jack",
		Age:      18,
		Address:  "usa",
	}
	// XML格式化输出
	xmlData, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("xml序列化失败", err)
		return
	}
	fmt.Printf("XML数据:\n%s\n", string(xmlData))
}

// 3.2 XML反序列化示例
func demoXMLUnmarshal() {
	fmt.Println("\n--- 3.2 XML 反序列化 ---")
	xmlStr := `<Person>
<id>120</id>
<name>jack</name>
<age>18</age>
<address>usa</address>
</Person>`

	var person Person
	err := xml.Unmarshal([]byte(xmlStr), &person)
	if err != nil {
		fmt.Println("XML解析失败：", err)
		return
	}
	fmt.Println("解析结果： %+v\n", person)
}

// ============================= 4. YAML 序列化 =============================

// 4.1 YAML序列化示例
func demoYAML() {
	fmt.Println("\n--- 4.1 YAML 序列化 ---")
	fmt.Println("YAML常用于配置文件，语法简洁易读")

	person := Person{
		UserID:   "120",
		Username: "jack",
		Age:      18,
		Address:  "usa",
	}
	yamlData, err := yaml.Marshal(person)
	if err != nil {
		fmt.Println("YAML序列化失败：", err)
		return
	}
	fmt.Println("YAML数据：:\n%s\n", string(yamlData))
}

// 4.2 YAML文件读取示例
func demoYAMLFile() {
	fmt.Println("\n--- 4.2 YAML 文件读取 ---")
	// 模拟从文件读取YAML配置
	yamlContent := `use_id: "120""
username: "jack"
age: 18
address: "usa"`
	var person Person
	err := yaml.Unmarshal([]byte(yamlContent), &person)
	if err != nil {
		fmt.Println("yaml解析失败", err)
		return
	}
	fmt.Printf("%+v\n", person)
}

// ============================= 5. 数据格式对比总结 =============================
func compareFormats() {
	fmt.Println("\n--- 5. 数据格式对比 ---")

	person := Person{
		UserID:   "120",
		Username: "jack",
		Age:      18,
		Address:  "usa",
	}

	fmt.Println("同一数据在不同格式中的表现:")
	fmt.Println("\nJSON (最常用):")
	jsonData, _ := json.MarshalIndent(person, "", "  ")
	fmt.Println(string(jsonData))

	fmt.Println("\nXML (较冗长):")
	xmlData, _ := xml.MarshalIndent(person, "", "  ")
	fmt.Println(string(xmlData))

	fmt.Println("\nYAML (配置友好):")
	yamlData, _ := yaml.Marshal(person)
	fmt.Println(string(yamlData))
}

// ============================= 6. Protocol Buffers 说明 =============================

func demoProtobufInfo() {
	fmt.Println("\n--- 6. Protocol Buffers 说明 ---")
	fmt.Println("Protobuf是Google的高性能二进制序列化格式")
	fmt.Println("使用步骤:")
	fmt.Println("1. 定义 .proto 文件描述数据结构")
	fmt.Println("2. 使用 protoc 编译器生成对应语言代码")
	fmt.Println("3. 在代码中使用生成的类进行序列化")
	fmt.Println("")
	fmt.Println("特点:")
	fmt.Println("✅ 二进制格式，体积小")
	fmt.Println("✅ 序列化/反序列化速度快")
	fmt.Println("✅ 支持跨语言，类型安全")
	fmt.Println("✅ 适合RPC通信和高性能场景")
	fmt.Println("")
	fmt.Println("安装: go get github.com/golang/protobuf/proto")
}

// ============================= 7. 主函数 =============================

func main() {
	fmt.Println("🚀 Go 数据序列化完整示例")
	fmt.Println("===============================")
	fmt.Println("学习不同数据格式的序列化和反序列化方法")

	// 演示各种序列化格式
	demoJSON()
	demoJSONUnmarshal()

	demoXML()
	demoXMLUnmarshal()

	demoYAML()
	demoYAMLFile()

	compareFormats()
	demoProtobufInfo()

	fmt.Println("\n🎯 实际应用建议:")
	fmt.Println("1. Web API: 使用 JSON")
	fmt.Println("2. 配置文件: 使用 YAML")
	fmt.Println("3. 传统系统: 使用 XML")
	fmt.Println("4. 高性能场景: 使用 Protobuf")
	fmt.Println("5. 结构体字段必须首字母大写才能被序列化")
}

/*
🔍 核心知识点总结:

============================= 1. 序列化基础 =============================
✅ 序列化: 将Go对象转换为字符串/二进制数据
✅ 反序列化: 将数据转换回Go对象
✅ 结构体标签: 通过`xml:"id"`等形式控制字段名

============================= 2. JSON (推荐) =============================
✅ encoding/json 标准库
✅ Marshal(): 紧凑序列化
✅ MarshalIndent(): 格式化序列化
✅ Unmarshal(): 反序列化
✅ 适用: Web API、数据交换

============================= 3. XML (传统) =============================
✅ encoding/xml 标准库
✅ 语法冗长但结构清晰
✅ 适用: 传统系统、文档格式

============================= 4. YAML (配置) =============================
✅ gopkg.in/yaml.v3 第三方库
✅ 语法简洁，适合配置文件
✅ 缩进敏感，人类可读性好

============================= 5. Protobuf (高性能) =============================
✅ 二进制格式，性能最优
✅ 需要预定义.proto文件
✅ 适用: 微服务通信、高性能场景

============================= 6. 选择指南 =============================
🌐 网络传输: JSON (最通用)
⚙️ 配置文件: YAML (最直观)
🔧 传统集成: XML (兼容性)
🚀 性能优先: Protobuf (最高效)

💡 重要提醒:
1. 结构体字段必须首字母大写才能被序列化
2. 合理使用结构体标签控制字段名
3. 根据实际场景选择合适的数据格式
4. 错误处理很重要，记得检查err返回值

通过这个示例，你可以掌握Go语言中主要的数据序列化方式！
*/

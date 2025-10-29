package main

import (
	"flag"
	"fmt"
	"os"
)

// ============================= 1. 命令行参数基础 =============================

func main() {
	// 定义参数变量
	var (
		name   string
		age    int
		port   int
		debug  bool
		help   bool
		config string
	)

	// 2.1 使用变量绑定方式定义参数（推荐）
	flag.StringVar(&name, "name", "张三", "用户姓名")
	flag.IntVar(&age, "age", 18, "用户年龄")
	flag.IntVar(&port, "port", 8080, "服务端口")
	flag.BoolVar(&debug, "debug", false, "调试模式")
	flag.BoolVar(&help, "help", false, "显示帮助")
	flag.StringVar(&config, "config", "app.conf", "配置文件路径")

	// 2.2 自定义帮助信息
	flag.Usage = func() {
		fmt.Printf("用法: %s [选项] [参数...]\n\n", os.Args[0])
		fmt.Println("选项:")
		flag.PrintDefaults()
		fmt.Println("\n示例:")
		fmt.Printf("  %s -name 李四 -port 3000 -debug=true\n", os.Args[0])
		fmt.Printf("  %s -help\n", os.Args[0])
	}

	// 3. 解析命令行参数
	flag.Parse()

	// 4. 处理帮助请求
	if help {
		flag.Usage()
		return
	}

	// 5. 使用解析后的参数
	fmt.Println("=== 参数解析结果 ===")
	fmt.Printf("姓名: %s\n", name)
	fmt.Printf("年龄: %d\n", age)
	fmt.Printf("端口: %d\n", port)
	fmt.Printf("调试: %t\n", debug)
	fmt.Printf("配置: %s\n", config)

	// 6. 处理额外参数
	if flag.NArg() > 0 {
		fmt.Printf("\n额外参数(%d个):\n", flag.NArg())
		for i, arg := range flag.Args() {
			fmt.Printf("  [%d] %s\n", i, arg)
		}
	}
}

// ============================= 7. 其他重要知识点 =============================

/*
🔍 Flag 包核心总结:

✅ 参数定义方式:
   flag.TypeVar(&变量, "参数名", 默认值, "说明")  // 推荐
   flag.Type("参数名", 默认值, "说明")           // 返回指针

✅ 重要函数:
   flag.Parse()     // 解析参数（必须调用）
   flag.Args()      // 获取非选项参数
   flag.NArg()      // 非选项参数个数
   flag.NFlag()     // 已设置的选项参数个数

✅ 参数格式:
   -name value
   --name value
   -name=value
   --name=value

💡 使用技巧:
1. 布尔参数必须用等号: -debug=true
2. 解析在第一个非选项参数前停止
3. 自定义flag.Usage提供友好帮助
4. 及时处理-h/--help请求

🎯 编译测试:
go build -o myapp main.go
./myapp -name 李四 -port 3000 -debug=true arg1 arg2
./myapp -help
*/

package main

import (
	"log"
	"os"
)

// ============================= 1. 基础日志使用 =============================

func demoBasicLogging() {
	log.Println("\n--- 1. 基础日志使用 ---")

	// 基础日志输出
	log.Println("这是一条普通日志")

	// 注意：Fatal 和 Panic 会中断程序执行
	// log.Fatalln("Fatal日志 - 程序会退出") // 实际使用时取消注释测试
	// log.Panicln("Panic日志 - 会产生panic") // 实际使用时取消注释测试
}

// ============================= 2. 日志前缀设置 =============================

func demoLogPrefix() {
	log.Println("\n--- 2. 日志前缀设置 ---")

	// 获取当前前缀
	currentPrefix := log.Prefix()
	log.Printf("当前前缀: %s", currentPrefix)

	// 设置新的前缀
	log.SetPrefix("[APP] ")
	log.Println("这条日志带有自定义前缀")

	// 恢复默认前缀
	log.SetPrefix("")
}

// ============================= 3. 日志格式配置 =============================

func demoLogFlags() {
	log.Println("\n--- 3. 日志格式配置 ---")

	// 默认格式 (LstdFlags = Ldate | Ltime)
	log.Println("默认格式的日志")

	// 自定义格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix("[DEBUG] ")
	log.Println("带文件名和时间的详细日志")

	// 恢复默认设置
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("")
}

// ============================= 4. 自定义Logger实例 =============================

func demoCustomLogger() {
	log.Println("\n--- 4. 自定义Logger实例 ---")

	// 创建自定义Logger - 输出到文件
	file, err := os.Create("app.log")
	if err != nil {
		log.Fatal("创建日志文件失败:", err)
	}
	defer file.Close()

	// 自定义Logger：输出到文件，带前缀和格式
	customLog := log.New(file, "[CUSTOM] ", log.Ldate|log.Ltime|log.Lshortfile)
	customLog.Println("这条日志会写入文件")

	// 控制台和文件同时输出
	consoleAndFileLog := log.New(os.Stdout, "[CONSOLE] ", log.LstdFlags)
	consoleAndFileLog.Println("这条日志输出到控制台")
}

// ============================= 5. 日志级别模拟 =============================

func demoLogLevels() {
	log.Println("\n--- 5. 日志级别模拟 ---")

	// 在实际项目中，通常会使用更完善的日志库
	// 这里简单模拟不同日志级别

	// INFO 级别
	log.SetPrefix("[INFO] ")
	log.Println("应用程序启动")

	// WARN 级别
	log.SetPrefix("[WARN] ")
	log.Println("磁盘空间不足")

	// ERROR 级别
	log.SetPrefix("[ERROR] ")
	log.Println("数据库连接失败")

	// 恢复默认
	log.SetPrefix("")
}

// ============================= 6. 生产环境建议 =============================

func productionRecommendation() {
	log.Println("\n--- 6. 生产环境建议 ---")
	log.Println("标准库log包功能有限，生产环境推荐使用:")
	log.Println("✅ zap - 高性能日志库")
	log.Println("✅ logrus - 功能丰富的日志库")
	log.Println("✅ zerolog - 零分配JSON日志库")
}

// ============================= 7. 主函数 =============================

func main() {
	log.Println("🚀 Go Log 包使用示例")
	log.Println("========================")

	// 演示各种日志功能
	demoBasicLogging()
	demoLogPrefix()
	demoLogFlags()
	demoCustomLogger()
	demoLogLevels()
	productionRecommendation()

	log.Println("✅ 日志演示完成")
}

/*
🔍 Log 包核心知识点总结:

============================= 1. 基础日志函数 =============================
✅ log.Println()   - 普通日志
✅ log.Printf()    - 格式化日志
✅ log.Fatalln()   - 日志 + 退出程序 (os.Exit(1))
✅ log.Panicln()   - 日志 + panic

============================= 2. 前缀管理 =============================
✅ log.Prefix()    - 获取当前前缀
✅ log.SetPrefix() - 设置日志前缀

============================= 3. 格式标志 =============================
✅ Ldate         - 日期 (2009/01/23)
✅ Ltime         - 时间 (01:23:23)
✅ Lmicroseconds - 微秒精度
✅ Llongfile     - 完整文件路径
✅ Lshortfile    - 文件名和行号
✅ LUTC          - 使用UTC时间
✅ Lmsgprefix    - 前缀在消息前
✅ LstdFlags     - 标准标志 (Ldate | Ltime)

============================= 4. 自定义Logger =============================
✅ log.New() - 创建自定义Logger实例
✅ 可以指定输出位置、前缀、格式

============================= 5. 生产建议 =============================
🌱 开发环境: 标准log包足够
🚀 生产环境: 使用zap、logrus等第三方库

💡 使用技巧:
1. Fatal和Panic会中断程序，谨慎使用
2. 合理设置前缀便于日志分类
3. 根据环境调整日志详细程度
4. 生产环境考虑日志轮转和归档

这个示例涵盖了标准log包的核心用法，适合学习和开发阶段使用！
*/

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== Go 文件操作完整示例 ===\n")

	// ==================== 文件基础操作 ====================

	// 1. 文件打开与创建
	fmt.Println("1. 文件打开与创建:")
	// 使用 OpenFile 函数打开或创建文件
	// os.O_RDWR: 读写模式 | os.O_CREATE: 不存在则创建 | os.O_APPEND: 追加模式
	// 0644: 文件权限 (用户读写，组和其他只读)
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close() // defer 确保函数退出前关闭文件，释放资源
	fmt.Println("✅ 文件打开成功:", file.Name())

	// 2. 文件写入操作
	fmt.Println("\n2. 文件写入操作:")
	content := "这是第一行内容\n这是第二行内容\n"
	// 向文件写入字符串内容
	written, err := file.WriteString(content)
	if err != nil {
		fmt.Printf("写入失败: %v\n", err)
		return
	}
	fmt.Printf("✅ 写入 %d 字节\n", written)

	// 3. 文件读取操作
	fmt.Println("\n3. 文件读取操作:")
	// 将文件指针重置到开头，准备读取
	file.Seek(0, 0)

	// 方法1: 使用 io.ReadAll 读取全部内容 (适用于小文件)
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("读取失败: %v\n", err)
		return
	}
	fmt.Printf("✅ 文件内容:\n%s", string(data))

	// 方法2: 使用 os.ReadFile 简便读取 (内部会打开和关闭文件)
	data2, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Printf("简便读取失败: %v\n", err)
	} else {
		fmt.Printf("✅ 简便读取内容:\n%s", string(data2))
	}

	// ==================== 文件信息与检查 ====================

	fmt.Println("\n4. 文件信息检查:")
	// 获取文件的详细信息
	info, err := os.Stat("example.txt")
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
	} else {
		fmt.Printf("✅ 文件名: %s\n", info.Name())
		fmt.Printf("✅ 文件大小: %d 字节\n", info.Size())
		fmt.Printf("✅ 修改时间: %v\n", info.ModTime())
		fmt.Printf("✅ 是否是目录: %t\n", info.IsDir())
		fmt.Printf("✅ 权限: %s\n", info.Mode())
	}

	// 检查文件是否存在
	if _, err := os.Stat("example.txt"); err == nil {
		fmt.Println("✅ 文件存在")
	} else if os.IsNotExist(err) {
		fmt.Println("❌ 文件不存在")
	} else {
		fmt.Printf("❌ 检查文件出错: %v\n", err)
	}

	// ==================== 文件复制与重命名 ====================

	fmt.Println("\n5. 文件复制:")
	// 方法1: 使用 io.Copy 进行流式复制 (适用于大文件)
	srcFile, _ := os.Open("example.txt")        // 打开源文件
	dstFile, _ := os.Create("example_copy.txt") // 创建目标文件
	copied, err := io.Copy(dstFile, srcFile)    // 执行复制
	if err != nil {
		fmt.Printf("复制失败: %v\n", err)
	} else {
		fmt.Printf("✅ 复制成功，复制了 %d 字节\n", copied)
	}
	srcFile.Close() // 关闭源文件
	dstFile.Close() // 关闭目标文件

	fmt.Println("\n6. 文件重命名:")
	// 重命名文件 (也可用于移动文件)
	err = os.Rename("example_copy.txt", "renamed_example.txt")
	if err != nil {
		fmt.Printf("重命名失败: %v\n", err)
	} else {
		fmt.Println("✅ 重命名成功")
	}

	// ==================== 文件夹操作 ====================

	fmt.Println("\n7. 文件夹操作:")
	// 创建单级目录 (父目录必须存在)
	err = os.Mkdir("test_dir", 0755) // 0755: 用户读写执行，组和其他读执行
	if err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
	} else {
		fmt.Println("✅ 单级目录创建成功")
	}

	// 创建多级目录 (自动创建所有不存在的父目录)
	err = os.MkdirAll("parent/child/grandchild", 0755)
	if err != nil {
		fmt.Printf("创建多级目录失败: %v\n", err)
	} else {
		fmt.Println("✅ 多级目录创建成功")
	}

	// 在目录中创建测试文件
	os.WriteFile("parent/test_file.txt", []byte("测试文件内容"), 0644)

	// 读取目录内容
	fmt.Println("\n8. 目录遍历:")
	// 读取目录下的所有条目
	entries, err := os.ReadDir("parent")
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
	} else {
		fmt.Println("parent 目录内容:")
		for _, entry := range entries {
			info, _ := entry.Info() // 获取详细的文件信息
			fmt.Printf("  - %s (目录: %t, 大小: %d bytes)\n",
				entry.Name(), entry.IsDir(), info.Size())
		}
	}

	// 递归遍历目录
	fmt.Println("\n9. 递归遍历目录:")
	// filepath.Walk 会递归遍历指定目录及其所有子目录
	// 第二个参数是一个回调函数，对每个文件和目录都会调用
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果访问出错，返回错误
		}
		if info.IsDir() {
			fmt.Printf("📁 目录: %s\n", path)
		} else {
			fmt.Printf("📄 文件: %s (%d bytes)\n", path, info.Size())
		}
		return nil // 返回 nil 继续遍历
	})
	if err != nil {
		fmt.Printf("遍历失败: %v\n", err)
	}

	// ==================== 清理操作 ====================

	fmt.Println("\n10. 清理操作:")
	// 删除文件
	os.Remove("example.txt")
	os.Remove("renamed_example.txt")

	// 递归删除目录及其所有内容
	os.RemoveAll("test_dir")
	os.RemoveAll("parent")

	fmt.Println("✅ 清理完成")
}

// ==================== 核心知识点总结 ====================
/*
📌 文件打开模式常量:
  os.O_RDONLY    - 只读模式
  os.O_WRONLY    - 只写模式
  os.O_RDWR      - 读写模式
  os.O_CREATE    - 不存在时创建
  os.O_APPEND    - 追加模式
  os.O_TRUNC     - 打开时清空文件
  os.O_EXCL      - 与O_CREATE一起使用，文件必须不存在

📌 常用文件权限:
  0644 - 用户读写，组和其他只读
  0755 - 用户读写执行，组和其他读执行
  0600 - 仅用户读写

📌 核心函数说明:
  文件操作:
  - os.OpenFile()    - 最灵活的文件打开方式
  - os.Open()        - 只读方式打开文件
  - os.Create()      - 创建并打开文件(截断已存在文件)
  - file.WriteString() - 写入字符串
  - io.ReadAll()     - 读取全部内容
  - os.ReadFile()    - 简便的文件读取
  - os.WriteFile()   - 简便的文件写入
  - io.Copy()        - 流式复制文件
  - os.Rename()      - 重命名/移动文件
  - os.Remove()      - 删除单个文件或空目录

  文件夹操作:
  - os.Mkdir()       - 创建单级目录
  - os.MkdirAll()    - 创建多级目录
  - os.ReadDir()     - 读取目录内容
  - filepath.Walk()  - 递归遍历目录树
  - os.RemoveAll()   - 递归删除目录及其内容

💡 最佳实践:
  1. 总是检查错误返回值
  2. 使用 defer file.Close() 确保文件被关闭
  3. 大文件使用流式读写(io.Copy)避免内存问题
  4. 使用 filepath.Join() 处理跨平台路径
  5. 临时文件使用 os.CreateTemp()
  6. 重要操作前先备份文件
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// ============ERROR===========
// 自定义错误类型
type TimeError struct {
	Msg  string
	Time time.Time
}

func (e *TimeError) Error() string {
	return fmt.Sprintf("%s at %v", e.Msg, e.Time.Format("2006-01-02 15:04:05"))
}

// 错误创建和包装
func demoError() error {
	//基础错误
	err := errors.New("原始错误")

	//错误包装（链式错误）
	return fmt.Errorf("外层错误，%w", err)
}

// 错误检查
func checkErrors() {
	err := demoError()

	//检查错误类型
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("文件不存在")
	}

	//提取具体错误类型
	var timeErr *TimeError
	if errors.As(err, &timeErr) {
		fmt.Println("时间错误", timeErr.Time)
	}
}

// ==================== PANIC 和 RECOVER ====================
// panic 触发
func dangerError() {
	panic("严格错误发生！")
}

// panic 恢复
func safeOp() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("恢复的panic", r)
		}
	}()
	dangerError()
}

// 资源清理示例
func resourceOp() (err error) {
	//模拟资源打开
	fmt.Println("打开资源。。。")

	defer func() {
		//资源清理（无论是否panic都会执行）
		fmt.Println("清理资源。。。")
		//panic恢复
		if r := recover(); r != nil {
			err = fmt.Errorf("操作失败：&v", r)
		}
	}()

	//模拟可能panic的操作
	if time.Now().Second()%2 == 0 {
		panic("随机panic！")
	}
	fmt.Println("操作成功")
	return nil
}

// ==================== FATAL 错误 ====================
func fatalOp() {
	// 致命错误 - 立即退出，不执行defer
	fmt.Println("发生致命错误")
	os.Exit(1)
}

// ==================== 实际应用示例 ====================
func fileProcessor(filename string) error {
	//错误传递实例
	if len(filename) == 0 {
		return errors.New("文件名不能为空")
	}
	file, err := os.Open(filename)
	if err != nil {
		//包装错误信息
		return fmt.Errorf("打开文件失败", err)
	}
	defer file.Close()
	// 文件处理逻辑...
	return nil
}

func main() {
	fmt.Println("=== Error 处理示例 ===")
	checkErrors()

	fmt.Println("\n=== Panic 恢复实例===")
	safeOp()
	fmt.Println("程序继续运行...")

	fmt.Println("\n=== 带资源清理的操作 ===")
	if err := resourceOp(); err != nil {
		fmt.Println("错误:", err)
	}

	fmt.Println("\n=== 文件处理错误传递 ===")
	if err := fileProcessor(""); err != nil {
		fmt.Println("文件处理失败", err)

		//错误解包
		if original := errors.Unwrap(err); original != nil {
			fmt.Println("原始错误：", original)
		}
	}
	fmt.Println("\n程序正常结束")
}

// ==================== 关键要点总结 ====================
/*
ERROR (正常错误):
  - errors.New() 创建基础错误
  - fmt.Errorf("%w", err) 包装错误链
  - errors.Is() 检查错误类型
  - errors.As() 提取具体错误类型
  - 作为返回值传递，需要处理

PANIC (严重异常):
  - panic() 触发异常
  - recover() 在 defer 中恢复
  - 会执行当前函数的 defer 清理
  - 恢复后程序可继续运行

FATAL (致命错误):
  - os.Exit() 立即退出
  - 不执行 defer 清理
  - 仅用于无法恢复的情况

最佳实践:
  - 错误要处理或传递，不要忽略
  - panic 只用于真正异常情况
  - 使用 defer 确保资源清理
  - 在库函数中返回 error 而非 panic
*/

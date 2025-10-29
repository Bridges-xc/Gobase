package function

import "fmt"

func main() {
	// grow 是一个闭包函数，它"记住"了外部变量 e 和 n
	grow := Exp(2) // 调用Exp(2)，返回一个闭包函数

	for i := range 10 {
		// 每次调用 grow() 时，它都能记住上次执行后的状态
		fmt.Println("2^%d=%d\n", i, grow())
	}
}

func Exp(n int) func() int {
	e := 1 // 局部变量 e，初始值为1

	// 返回一个匿名函数 - 这就是闭包
	return func() int {
		temp := e   // 保存当前 e 的值（这是闭包捕获的外部变量）
		e *= n      // 更新 e = e * n（闭包修改了外部变量）
		return temp // 返回之前保存的值
	}
	// 注意：虽然 Exp 函数执行完毕，但 e 和 n 不会被回收！
	// 因为它们被返回的闭包函数引用着
}

//匿名函数 + 引用外部变量 = 闭包！！！！！！！！！！！！！！！！！

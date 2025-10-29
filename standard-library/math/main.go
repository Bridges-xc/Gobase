package main

import (
	"fmt"
	"math"
)

func main() {
	// ============================= 1. 数学常量 =============================
	fmt.Println("=== 数学常量 ===")
	fmt.Printf("自然常数 e: %.2f\n", math.E)
	fmt.Printf("圆周率 π: %.2f\n", math.Pi)
	fmt.Printf("黄金比例 φ: %.2f\n", math.Phi)
	fmt.Printf("2的平方根: %.2f\n", math.Sqrt2)

	// ============================= 2. 极值常量 =============================
	fmt.Println("\n=== 极值常量 ===")
	fmt.Printf("最大int64: %d\n", math.MaxInt64)
	fmt.Printf("最小int64: %d\n", math.MinInt64)
	fmt.Printf("最大float64: %e\n", math.MaxFloat64)
	fmt.Printf("最小非零float64: %e\n", math.SmallestNonzeroFloat64)

	// ============================= 3. 基本运算 =============================
	fmt.Println("\n=== 基本运算 ===")
	// 最大值最小值
	fmt.Printf("Max(3.5, 2.1) = %.1f\n", math.Max(3.5, 2.1))
	fmt.Printf("Min(3.5, 2.1) = %.1f\n", math.Min(3.5, 2.1))

	// 绝对值
	fmt.Printf("Abs(-15.7) = %.1f\n", math.Abs(-15.7))

	// 取余运算
	fmt.Printf("Mod(15, 4) = %.1f\n", math.Mod(15, 4))

	// ============================= 4. 取整函数 =============================
	fmt.Println("\n=== 取整函数 ===")
	num := 3.75
	fmt.Printf("原始值: %.2f\n", num)
	fmt.Printf("Trunc(截断): %.1f\n", math.Trunc(num))
	fmt.Printf("Floor(向下): %.1f\n", math.Floor(num))
	fmt.Printf("Ceil(向上): %.1f\n", math.Ceil(num))
	fmt.Printf("Round(四舍五入): %.1f\n", math.Round(num))

	// ============================= 5. 特殊值检测 =============================
	fmt.Println("\n=== 特殊值检测 ===")
	// NaN检测
	nanValue := math.NaN()
	fmt.Printf("IsNaN(NaN): %t\n", math.IsNaN(nanValue))

	// 无穷大检测
	posInf := math.Inf(1)  // 正无穷
	negInf := math.Inf(-1) // 负无穷
	fmt.Printf("IsInf(+∞, 1): %t\n", math.IsInf(posInf, 1))
	fmt.Printf("IsInf(-∞, -1): %t\n", math.IsInf(negInf, -1))

	// ============================= 6. 指数和对数 =============================
	fmt.Println("\n=== 指数和对数 ===")
	// 自然指数
	fmt.Printf("Exp(2) = e² = %.2f\n", math.Exp(2))

	// 幂运算
	fmt.Printf("Pow(2, 3) = 2³ = %.0f\n", math.Pow(2, 3))
	fmt.Printf("Pow(5, 2) = 5² = %.0f\n", math.Pow(5, 2))

	// 对数运算
	fmt.Printf("Log(100) = ln(100) = %.2f\n", math.Log(100))
	fmt.Printf("Log10(100) = log₁₀(100) = %.0f\n", math.Log10(100))

	// 换底公式计算 log₂(8)
	fmt.Printf("Log2(8) = ln(8)/ln(2) = %.0f\n", math.Log2(8))

	// ============================= 7. 开方运算 =============================
	fmt.Println("\n=== 开方运算 ===")
	// 平方根
	fmt.Printf("Sqrt(16) = %.0f\n", math.Sqrt(16))

	// 立方根
	fmt.Printf("Cbrt(27) = %.0f\n", math.Cbrt(27))

	// 开N次方：使用幂运算
	fmt.Printf("8^(1/3) = %.0f\n", math.Pow(8, 1.0/3))
	fmt.Printf("256^(1/4) = %.0f\n", math.Pow(256, 1.0/4))

	// ============================= 8. 三角函数 =============================
	fmt.Println("\n=== 三角函数 ===")
	angle := math.Pi / 6 // 30度
	fmt.Printf("角度: π/6 (30度)\n")
	fmt.Printf("Sin(π/6) = %.3f\n", math.Sin(angle))
	fmt.Printf("Cos(π/6) = %.3f\n", math.Cos(angle))
	fmt.Printf("Tan(π/6) = %.3f\n", math.Tan(angle))

	// 反三角函数
	fmt.Printf("Asin(0.5) = %.3f\n", math.Asin(0.5))
	fmt.Printf("Acos(0.866) = %.3f\n", math.Acos(0.866))

	// ============================= 9. 实用示例 =============================
	fmt.Println("\n=== 实用示例 ===")
	// 计算圆的面积
	radius := 5.0
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("半径%.0f的圆面积: %.2f\n", radius, area)

	// 计算直角三角形斜边
	a, b := 3.0, 4.0
	hypotenuse := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Printf("直角边%.0f和%.0f的斜边: %.1f\n", a, b, hypotenuse)
}

// ============================= 总结知识点 =============================
/*
1. 数学常量: E(自然常数), Pi(圆周率), Phi(黄金比例)等
2. 极值常量: 各数值类型的最大值最小值常量
3. 基本运算: Max/Min(最值), Abs(绝对值), Mod(取余)
4. 取整函数: Trunc(截断), Floor(向下), Ceil(向上), Round(四舍五入)
5. 特殊值: NaN(非数), Inf(无穷大)的创建和检测
6. 指数对数: Exp(e的幂), Pow(幂运算), Log(自然对数), Log10(常用对数)
7. 开方运算: Sqrt(平方根), Cbrt(立方根), Pow开N次方
8. 三角函数: Sin/Cos/Tan及其反函数，使用弧度制
9. 实际应用: 结合数学公式解决几何、物理等问题

注意: math包主要处理float64类型，整数运算建议使用其他方法
*/

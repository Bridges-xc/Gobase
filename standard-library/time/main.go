package main

import (
	"fmt"
	"time"
)

func main() {
	// ============================= 1. 时间获取与基础操作 ====================

	// 1.1 获取当前时间
	now := time.Now()
	fmt.Printf("当前时间: %v\n", now)
	// 1.2 时间比较操作
	later := now.Add(2 * time.Hour)
	fmt.Printf("时间比较 - 之后: %t, 之前: %t, 相等: %t\n",
		now.After(later), now.Before(later), now.Equal(later))
	// 1.3 时间差计算
	diff := now.Sub(now)
	fmt.Printf("时间差: %v\n", diff)

	// ============================= 2. 时间单位与常量 ====================
	// 2.1 时间单位常量使用
	fmt.Printf("时间单位 - 纳秒: %v, 微秒: %v\n", time.Nanosecond, time.Microsecond)
	fmt.Printf("时间单位 - 毫秒: %v, 秒: %v\n", time.Millisecond, time.Second)
	fmt.Printf("时间单位 - 分钟: %v, 小时: %v\n", time.Minute, time.Hour)
	// 2.2 时间加减操作
	newTime := now.Add(24 * time.Hour) // 加1天
	fmt.Printf("24小时后: %v\n", newTime)
	// ============================= 3. 时间格式化 ====================

	// 3.1 标准格式化 - 使用Go诞生日2006-01-02 15:04:05作为模板
	fmt.Printf("24小时格式: %s\n", now.Format("2006-01-02 15:04:05 Monday Jan"))
	fmt.Printf("只显示日期: %s\n", now.Format("2006-01-02"))
	fmt.Printf("12小时格式: %s\n", now.Format("03:04:05 PM"))
	fmt.Printf("自定义格式: %s\n", now.Format("2006年01月02日 15时04分"))
	// ============================= 4. 时间解析 ====================

	// 4.1 解析时间字符串
	timeStr := "2023-10-01 18:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
	} else {
		fmt.Printf("解析结果: %v\n", parsedTime)
	}
	// 4.2 带时区解析
	location, _ := time.LoadLocation("Asia/Shanghai")
	timeInLocation, err := time.ParseInLocation("2006/01/02", "2023/10/01", location)
	if err != nil {
		fmt.Printf("带时区解析错误: %v\n", err)
	} else {
		fmt.Printf("带时区解析: %v\n", timeInLocation)
	}
	// ============================= 5. 计时器 Timer ====================

	// 5.1 一次性计时器
	fmt.Println("启动2秒计时器...")
	timer := time.NewTimer(2 * time.Second)
	go func() {
		<-timer.C // 阻塞直到计时器触发
		fmt.Println("计时器触发!")
	}()
	// 停止计时器(如果还没触发)
	// timer.Stop()
	// ============================= 6. 定时器 Ticker ====================

	// 6.1 周期性定时器
	fmt.Println("启动定时器(3次触发)...")
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		count := 0
		for t := range ticker.C {
			fmt.Printf("定时器触发: %v\n", t.Format("15:04:05"))
			count++
			if count >= 3 {
				ticker.Stop() // 停止定时器
				break
			}
		}
	}()
	// ============================= 7. 延时操作 ====================

	// 7.1 使用Sleep延时
	start := time.Now()
	fmt.Printf("开始Sleep: %v\n", start.Format("15:04:05.000"))
	time.Sleep(1500 * time.Millisecond)
	end := time.Now()
	fmt.Printf("结束Sleep: %v, 实际休眠: %v\n",
		end.Format("15:04:05.000"), end.Sub(start))
	// 7.2 使用After延时 - 返回一个channel
	fmt.Println("等待1秒...")
	<-time.After(1 * time.Second)
	fmt.Println("1秒后!")
	//============================= 8. 时间戳操作 ====================

	// 8.1 获取Unix时间戳
	unixSec := now.Unix()
	unixMilli := now.UnixMilli()
	unixMicro := now.UnixMicro()
	unixNano := now.UnixNano()
	fmt.Printf("时间戳 - 秒: %d, 毫秒: %d\n", unixSec, unixMilli)
	fmt.Printf("时间戳 - 微秒: %d, 纳秒: %d\n", unixMicro, unixNano)
	// 8.2 从时间戳创建时间
	fromUnix := time.Unix(unixSec, 0)
	fmt.Printf("从时间戳创建: %v\n", fromUnix)
	// 等待所有goroutine完成
	time.Sleep(5 * time.Second)
	fmt.Println("程序结束")
	// ============================= 总结知识点 ====================
	/*
	   1. 时间获取: time.Now() 获取当前时间，返回Time结构体
	   2. 时间比较: After(), Before(), Equal() 方法进行时间先后比较
	   3. 时间计算: Sub() 计算时间差，Add() 增加时间间隔
	   4. 时间单位: 内置Nanosecond到Hour常量，用于时间运算
	   5. 时间格式化: 使用Go诞生日"2006-01-02 15:04:05"作为模板格式
	   6. 时间解析: Parse() 解析字符串为时间，ParseInLocation() 带时区解析
	   7. 计时器Timer: 一次性计时器，通过channel接收触发信号
	   8. 定时器Ticker: 周期性定时器，需要手动Stop()停止
	   9. 延时操作: Sleep() 阻塞当前goroutine，After() 返回触发channel
	   10. 时间戳: Unix()系列方法获取时间戳，Unix()从时间戳创建时间

	   注意:
	   - 格式化必须使用Go特定时间模板
	   - Timer和Ticker使用后要及时Stop避免资源泄露
	   - 时间操作要考虑时区影响
	   - 所有时间操作都是线程安全的
	*/
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

// ============================= 数据结构定义 =============================
type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

// ============================= HTTP 客户端 =============================

// 基础GET请求示例
func demoBasicGet() {
	fmt.Println("\n--- 基础GET请求 ---")

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("GET请求失败:", err)
		return
	}
	defer resp.Body.Close() // 必须关闭响应体

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", string(body))
}

// 基础POST请求示例
func demoBasicPost() {
	fmt.Println("\n--- 基础POST请求 ---")

	user := User{
		UserID:   "120",
		Username: "jack",
		Age:      18,
		Address:  "usa",
	}

	jsonData, _ := json.Marshal(user)
	reader := bytes.NewReader(jsonData)

	resp, err := http.Post("https://httpbin.org/post", "application/json", reader)
	if err != nil {
		fmt.Println("POST请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", string(body))
}

// 自定义客户端示例
func demoCustomClient() {
	fmt.Println("\n--- 自定义客户端 ---")

	client := &http.Client{
		Timeout: 10 * time.Second, // 设置超时时间
		// Transport: 可以配置更底层的网络设置
		// Jar:       Cookie管理
	}

	req, _ := http.NewRequest("GET", "https://httpbin.org/headers", nil)

	// 添加自定义Header
	req.Header.Add("Authorization", "Bearer token123")
	req.Header.Add("User-Agent", "MyGoClient/1.0")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("自定义客户端请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", string(body))
}

// ============================= HTTP 服务端 =============================

// 自定义处理器 - 实现Handler接口
type CustomHandler struct{}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "自定义处理器响应\n方法: %s\n路径: %s", r.Method, r.URL.Path)
}

// 处理器函数 - 更简洁的方式
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// RESTful用户API处理器
func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// 获取用户信息
		user := User{UserID: "123", Username: "John", Age: 25, Address: "NYC"}
		json.NewEncoder(w).Encode(user)
	case "POST":
		// 创建用户
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "无效的JSON数据", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, `{"status": "success", "message": "用户 %s 创建成功"}`, user.Username)
	default:
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
	}
}

// 启动自定义配置服务器
func startCustomServer() {
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second, // 读取超时
		WriteTimeout: 10 * time.Second, // 写入超时
		IdleTimeout:  30 * time.Second, // 空闲连接超时
	}

	// 注册路由
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "欢迎访问首页!\n路径: %s", r.URL.Path)
	}))
	http.Handle("/hello", http.HandlerFunc(helloHandler))
	http.Handle("/user", http.HandlerFunc(userHandler))
	http.Handle("/custom", &CustomHandler{})

	fmt.Println("自定义服务器运行在 http://localhost:8080")
	fmt.Println("可用路由:")
	fmt.Println("  GET  /              - 首页")
	fmt.Println("  GET  /hello?name=xxx - 问候页面")
	fmt.Println("  GET  /user          - 获取用户信息")
	fmt.Println("  POST /user          - 创建用户")
	fmt.Println("  ANY  /custom        - 自定义处理器")

	log.Fatal(server.ListenAndServe())
}

// ============================= 反向代理 =============================

// 反向代理处理器
func reverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	director := func(req *http.Request) {
		// 修改请求指向目标服务器
		req.URL.Scheme = "https"
		req.URL.Host = "httpbin.org"
		// 保持原始路径或自定义路径
		if req.URL.Path == "/forward" {
			req.URL.Path = "/get"
		}
		// 设置Host头
		req.Host = "httpbin.org"
		// 添加自定义头
		req.Header.Set("X-Proxy", "Go-Reverse-Proxy")
	}

	// 创建并执行反向代理
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}

// 启动反向代理服务器
func startProxyServer() {
	http.HandleFunc("/forward", reverseProxyHandler)

	fmt.Println("\n反向代理服务器运行在 http://localhost:8081")
	fmt.Println("访问 http://localhost:8081/forward 将代理到 https://httpbin.org/get")

	log.Fatal(http.ListenAndServe(":8081", nil))
}

// ============================= 主函数 =============================

func main() {
	fmt.Println("🚀 Go HTTP 编程完整示例")
	fmt.Println("========================")

	// 演示HTTP客户端功能
	fmt.Println("\n📡 客户端示例:")
	demoBasicGet()
	demoBasicPost()
	demoCustomClient()

	// 启动服务器
	fmt.Println("\n🌐 服务端示例:")

	// 启动自定义服务器
	go startCustomServer()

	// 启动反向代理服务器
	go startProxyServer()

	// 保持程序运行
	fmt.Println("\n⏳ 服务器已启动，按 Ctrl+C 退出...")
	select {}
}

/*
🔍 核心知识点总结:

HTTP 客户端:
✅ http.Get()/Post(): 简单快捷的请求方法
✅ http.Client{}: 可配置客户端(超时、传输层、Cookie等)
✅ http.NewRequest(): 创建复杂请求(自定义Header、Method等)
✅ 必须调用 defer resp.Body.Close() 释放连接

HTTP 服务端:
✅ http.ListenAndServe(): 快速启动默认服务器
✅ http.Server{}: 自定义服务器配置(超时、端口等)
✅ 两种处理器注册方式:
   - http.Handle(): 注册实现了Handler接口的对象
   - http.HandleFunc(): 注册普通函数作为处理器
✅ DefaultServeMux: 默认的多路复用器

反向代理:
✅ httputil.ReverseProxy: 内置反向代理功能
✅ Director函数: 用于修改转发的请求(URL、Header等)

最佳实践:
✅ 总是检查错误处理
✅ 及时关闭响应体避免资源泄漏
✅ 生产环境使用更安全的配置
✅ 考虑使用Context进行超时和取消控制
✅ 对于复杂项目，推荐使用成熟的Web框架(Gin、Echo等)

这个示例涵盖了Go语言net/http包的核心功能，是学习Web开发的坚实基础！
*/

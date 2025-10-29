package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func main() {
	//============================= 1. MAC地址解析 ===================
	hw, err := net.ParseMAC("00:1A:2B:3C:4D:5E")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1. MAC地址: %s\n", hw)

	//============================= 2. CIDR解析 ====================
	ip, ipnet, err := net.ParseCIDR("192.168.1.0/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("2. CIDR - IP: %s, 网络: %s\n", ip, ipnet)

	//============================= 3. IP地址解析 ===================
	ipv4Addr, _ := net.ResolveIPAddr("ip4", "192.168.1.1")
	ipv6Addr, _ := net.ResolveIPAddr("ip6", "::1")
	fmt.Printf("3. IPv4: %s, IPv6: %s\n", ipv4Addr, ipv6Addr)

	//============================= 4. TCP地址解析 ==================
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "localhost:8080")
	fmt.Printf("4. TCP地址: %s\n", tcpAddr)

	//============================= 5. UDP地址解析 ==================
	udpAddr, _ := net.ResolveUDPAddr("udp4", "localhost:9090")
	fmt.Printf("5. UDP地址: %s\n", udpAddr)

	//============================= 6. Unix域套接字 =================
	unixAddr, _ := net.ResolveUnixAddr("unix", "/tmp/test.sock")
	fmt.Printf("6. Unix地址: %s\n", unixAddr)

	//============================= 7. DNS查询 =====================
	// 7.1 查询主机地址
	hosts, _ := net.LookupHost("github.com")
	fmt.Printf("7.1 GitHub IP: %v\n", hosts[:2]) // 只显示前两个

	// 7.2 查询MX记录
	mxs, _ := net.LookupMX("github.com")
	if len(mxs) > 0 {
		fmt.Printf("7.2 MX记录: %s\n", mxs[0].Host)
	}

	//============================= 8. TCP服务器示例 ================
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// 8.1 启动TCP服务器
		listener, err := net.Listen("tcp", "localhost:12345")
		if err != nil {
			log.Fatal(err)
		}
		defer listener.Close()
		fmt.Println("8. TCP服务器启动在 localhost:12345")

		for {
			// 8.2 接受客户端连接
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}

			// 8.3 为每个连接创建goroutine处理
			go handleConnection(conn)
		}
	}()

	//============================= 9. TCP客户端示例 ================
	wg.Add(1)
	go func() {
		defer wg.Done()

		// 9.1 连接到服务器
		conn, err := net.Dial("tcp", "localhost:12345")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		// 9.2 发送数据
		message := "Hello TCP Server!"
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("9. 客户端发送: %s\n", message)

		// 9.3 读取响应
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		fmt.Printf("   收到服务器响应: %s\n", string(buffer[:n]))
	}()

	wg.Wait()
}

// ============================= 10. 连接处理函数 =================
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 10.1 读取客户端数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		if err != io.EOF {
			log.Print(err)
		}
		return
	}

	clientMsg := string(buffer[:n])
	fmt.Printf("10. 服务器收到: %s\n", clientMsg)

	// 10.2 发送响应
	response := "Message received!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Print(err)
	}
}

// 总结知识点:
// 1. MAC地址解析 - ParseMAC() 解析硬件地址
// 2. CIDR解析 - ParseCIDR() 解析IP地址和网络掩码
// 3. IP地址解析 - ResolveIPAddr() 支持IPv4/IPv6
// 4. TCP地址解析 - ResolveTCPAddr() 解析TCP端点地址
// 5. UDP地址解析 - ResolveUDPAddr() 解析UDP端点地址
// 6. Unix域套接字 - ResolveUnixAddr() 用于进程间通信
// 7. DNS查询 - LookupHost()查IP, LookupMX()查邮件交换记录
// 8. TCP服务器 - Listen()监听, Accept()接受连接, goroutine处理并发
// 9. TCP客户端 - Dial()建立连接, Write()发送, Read()接收
// 10. 网络编程模式 - 每个连接独立goroutine处理, 非阻塞IO
// 关键优势: Go的goroutine让网络编程简洁高效，轻松处理高并发连接

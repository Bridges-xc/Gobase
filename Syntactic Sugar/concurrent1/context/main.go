// ============================= 1. Context接口和基础 ====================
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Context接口的四个核心方法：
// Deadline() - 返回截止时间和是否设置
// Done() - 返回只读管道，用于接收取消信号
// Err() - 返回取消`原因
// Value() - 返回值

func contextInterfaceDemo() {
	fmt.Println("=== Context接口演示 ===")

	// 1.1 创建空的上下文
	background := context.Background() // 通常作为根上下文
	todo := context.TODO()             // 不确定使用哪种上下文时使用

	fmt.Printf("Background: %T\n", background)
	fmt.Printf("TODO: %T\n", todo)

	// 测试emptyCtx的方法
	deadline, ok := background.Deadline()
	fmt.Printf("Deadline: %v, Set: %t\n", deadline, ok)
	fmt.Printf("Err: %v\n", background.Err())
	fmt.Printf("Value: %v\n", background.Value("key"))
}

// ============================= 2. valueCtx使用 ====================
type keyType string

const (
	userKey      keyType = "user"
	requestIDKey keyType = "requestID"
)

func valueContextDemo() {
	fmt.Println("\n=== valueCtx值传递演示 ===")

	// 现在使用包级别的常量
	ctx := context.WithValue(context.Background(), userKey, "Alice")
	ctx = context.WithValue(ctx, requestIDKey, "12345")

	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	// 使用包级别的常量，类型一致
	user := ctx.Value(userKey)
	requestID := ctx.Value(requestIDKey)

	fmt.Printf("处理请求 - 用户: %v, 请求ID: %v\n", user, requestID)

	// 获取不存在的键
	notFound := ctx.Value("nonexistent")
	fmt.Printf("不存在的键: %v\n", notFound)
}

// ============================= 3. cancelCtx使用 ====================
func cancelContextDemo() {
	fmt.Println("\n=== cancelCtx取消控制演示 ===")

	var wg sync.WaitGroup

	// 3.1 创建可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx, &wg, "worker1")

	// 3秒后取消上下文
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("主协程发送取消信号...")
		cancel()
	})

	wg.Wait()
	fmt.Println("Worker已停止")
}

func worker(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 收到取消信号, 原因: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s: 工作中...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// ============================= 4. 多级上下文取消 ====================
func nestedCancelDemo() {
	fmt.Println("\n=== 多级上下文取消演示 ===")

	var wg sync.WaitGroup
	rootCtx, cancel := context.WithCancel(context.Background())

	wg.Add(3)

	// 创建子上下文
	childCtx1, _ := context.WithCancel(rootCtx)
	childCtx2, _ := context.WithCancel(rootCtx)

	go nestedWorker(rootCtx, &wg, "root-worker")
	go nestedWorker(childCtx1, &wg, "child1-worker")
	go nestedWorker(childCtx2, &wg, "child2-worker")

	// 2秒后取消根上下文，所有子上下文都会收到信号
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("取消根上下文...")
		cancel()
	})

	wg.Wait()
	fmt.Println("所有worker已停止")
}

func nestedWorker(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer wg.Done()

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 停止, 原因: %v\n", name, ctx.Err())
			return
		case <-ticker.C:
			fmt.Printf("%s: 运行中\n", name)
		}
	}
}

// ============================= 5. timerCtx超时控制 ====================
func timeoutContextDemo() {
	fmt.Println("\n=== timerCtx超时控制演示 ===")

	var wg sync.WaitGroup

	// 5.1 WithTimeout - 相对时间
	wg.Add(1)
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 良好实践：总是调用cancel

	go func() {
		defer wg.Done()
		processWithTimeout(timeoutCtx, "timeout-task")
	}()

	wg.Wait()

	// 5.2 WithDeadline - 绝对时间
	wg.Add(1)
	deadline := time.Now().Add(3 * time.Second)
	deadlineCtx, cancel2 := context.WithDeadline(context.Background(), deadline)
	defer cancel2()

	go func() {
		defer wg.Done()
		processWithTimeout(deadlineCtx, "deadline-task")
	}()

	wg.Wait()
}

func processWithTimeout(ctx context.Context, taskName string) {
	for i := 1; i <= 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 任务取消, 原因: %v\n", taskName, ctx.Err())
			return
		default:
			fmt.Printf("%s: 步骤 %d\n", taskName, i)
			time.Sleep(800 * time.Millisecond) // 模拟耗时操作
		}
	}
	fmt.Printf("%s: 任务完成\n", taskName)
}

// ============================= 6. 实际应用场景 ====================
func practicalExample() {
	fmt.Println("\n=== 实际应用场景演示 ===")

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 模拟HTTP请求处理
	wg.Add(1)
	go httpHandler(ctx, &wg, "/api/users")

	wg.Wait()
}

func httpHandler(ctx context.Context, wg *sync.WaitGroup, path string) {
	defer wg.Done()

	fmt.Printf("开始处理请求: %s\n", path)

	// 为不同的服务创建子上下文
	authCtx, authCancel := context.WithTimeout(ctx, 2*time.Second)
	defer authCancel()

	dbCtx, dbCancel := context.WithTimeout(ctx, 3*time.Second)
	defer dbCancel()

	// 并行执行认证和数据库操作
	authCh := make(chan bool, 1)
	dbCh := make(chan string, 1)

	go authenticate(authCtx, authCh)
	go queryDatabase(dbCtx, dbCh)

	// 等待结果或超时
	select {
	case <-ctx.Done():
		fmt.Printf("请求超时或取消: %v\n", ctx.Err())
	case authSuccess := <-authCh:
		if !authSuccess {
			fmt.Println("认证失败")
			return
		}
		userData := <-dbCh
		fmt.Printf("请求成功: %s\n", userData)
	}
}

func authenticate(ctx context.Context, ch chan<- bool) {
	// 模拟认证过程
	select {
	case <-time.After(1 * time.Second):
		ch <- true // 认证成功
	case <-ctx.Done():
		ch <- false // 认证超时
	}
}

func queryDatabase(ctx context.Context, ch chan<- string) {
	// 模拟数据库查询
	select {
	case <-time.After(2 * time.Second):
		ch <- "用户数据" // 查询成功
	case <-ctx.Done():
		ch <- "" // 查询超时
	}
}

// ============================= 7. Context最佳实践 ====================
func bestPractices() {
	fmt.Println("\n=== Context最佳实践 ===")

	// 7.1 总是传递Context作为第一个参数
	// 7.2 在超时或取消时及时释放资源
	// 7.3 使用WithValue时定义自定义键类型

	type contextKey string
	const traceIDKey contextKey = "traceID"

	ctx := context.WithValue(context.Background(), traceIDKey, "trace-123")
	processWithTrace(ctx, traceIDKey)
}

func processWithTrace(ctx context.Context, key interface{}) {
	if traceID := ctx.Value(key); traceID != nil {
		fmt.Printf("处理带追踪ID的请求: %v\n", traceID)
	}
}

// ============================= 主函数入口 ====================
func main() {
	contextInterfaceDemo()
	valueContextDemo()
	cancelContextDemo()
	nestedCancelDemo()
	timeoutContextDemo()
	practicalExample()
	bestPractices()

	fmt.Println("\n=== 所有Context示例执行完成 ===")
}

// ============================= 总结知识点 ====================
/*
Context核心知识点总结：

1. Context接口：
   - Deadline(): 返回截止时间和是否设置
   - Done(): 返回只读管道，接收取消信号
   - Err(): 返回取消原因
   - Value(): 根据键返回值

2. 上下文类型：
   - emptyCtx: 空上下文，通过Background()/TODO()创建
   - valueCtx: 带值的上下文，用于数据传递
   - cancelCtx: 可取消的上下文，支持手动取消
   - timerCtx: 带超时的上下文，支持超时自动取消

3. 创建方法：
   - Background(): 创建根上下文
   - TODO(): 创建待确定上下文
   - WithValue(): 创建带值的上下文
   - WithCancel(): 创建可取消的上下文
   - WithTimeout(): 创建相对超时的上下文
   - WithDeadline(): 创建绝对超时的上下文

4. 取消传播：
   - 父上下文取消时，所有子上下文都会自动取消
   - 取消信号通过Done()管道传播
   - 支持多级嵌套取消

5. 使用场景：
   - 请求超时控制
   - 跨协程取消操作
   - 传递请求范围的数据
   - 资源清理和超时管理

6. 最佳实践：
   - Context作为函数第一个参数
   - 及时调用cancel函数释放资源
   - 使用自定义类型作为WithValue的键
   - 检查Context是否已取消
   - 避免上下文泄漏

7. 注意事项：
   - 不要存储Context在结构体中，应该显式传递
   - 相同的键在不同包中可能冲突，使用自定义类型
   - 取消上下文会释放相关资源，确保及时取消
*/

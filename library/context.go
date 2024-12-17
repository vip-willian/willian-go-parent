package library

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func TestContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println("n=", n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {

	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func TestContextWithDeadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// 通常用于数据库或者网络连接的超时控制
func TestContextWithTimeOut() {

	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_NODE"), "1251232453")

	wg.Add(1)

	go worker1(ctx)

	// sleep3秒以免程序过快退出
	time.Sleep(time.Second * 3)

	cancel() // 通知子goroutine结束

	wg.Wait()
	fmt.Println("over")
}

var wg sync.WaitGroup

type TraceCode string

func worker1(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
	go worker2(ctx)
LOOP:
	for {
		fmt.Printf("working ..., trace code:%s\n", traceCode)
		time.Sleep(10 * time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

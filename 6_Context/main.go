package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // 5 seconds pass
		fmt.Println("finished")
	case <-ctx.Done(): // ctx is cancelled
		/*
			`Context` 的 `Done` 方法非常重要，它返回一个只接收的 `chan` 类型的通道。
			这个通道被用来发送信号，告知 `Context` 已经结束了，无论是被取消还是达到了设定的截止时间（超时）。
			这个机制允许协程在 `Context` 被取消时及时作出反应，如终止运行、回收资源等，以避免不必要的工作和资源浪费。

			### 功能
			- **取消通知**：`Done` 通道是协程之间进行同步的一种方式。
			当 `Context` 被取消或者到了预定的超时时间时，`Done` 通道会被关闭，此时所有监听了这个通道的协程都能接收到这一事件。
			- **超时处理**：这允许开发人员编写能够响应取消或超时事件的代码。
			通过这种方式，可以优雅地中断正在进行的操作，如数据库查询、网络请求等。

			### 使用方式
			接收到 `Done` 通道的信号后，可以进一步通过 `ctx.Err()` 来获取具体的错误原因，这在处理取消或超时事件时非常有用。
		*/
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
		/*
			检查 `ctx.Done()` 通道接收到取消信号后，`Context` 被取消或达到了其截止时间（如果设置了的话）的具体原因。
			当 `Context` 被取消或超时后，`ctx.Done()` 会接收到一个值，这时可以通过调用 `ctx.Err()` 来获取取消的原因。
			`ctx.Err()` 返回一个错误，指出是何种情况导致了 `Context` 的取消：
			- 如果 `Context` 被取消，`ctx.Err()` 会返回 `context.Canceled`。
			- 如果 `Context` 因为超时而结束，`ctx.Err()` 会返回 `context.DeadlineExceeded`。
			这样处理的目的是为了让开发者能够理解取消发生的上下文，便于调试和处理程序中可能依赖于这一取消信号的部分。
		*/
	}

}

func main() {
	//创建空context的两种方法
	ctx := context.Background() // 返回一个空的context，不能被cancel，kv为空

	//todoCtx := context.TODO() // 和Background类似，当你不确定时使用
	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
	//ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2 * time.Second))

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	doSomething(ctx)
}

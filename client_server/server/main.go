package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	n, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes.\n", n)
	}
	/*
		`Fprintf` 函数是 Go 语言标准库 `fmt` 包中的一个函数，用于将格式化的字符串输出到任何实现了 `io.Writer` 接口的对象。函数的原型如下：

		func Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
			p := newPrinter()
			p.doPrintf(format, a)
			n, err = w.Write(p.buf)
			p.free()
			return
		}

		- `n int`：表示成功写入到 `io.Writer` 中的字节数。这个数值告诉你有多少字节是实际被写入的。在某些情况下，写入的字节数可能会少于你预期，这通常意味着发生了一个写入错误。
		- `err error`：如果在写入的过程中遇到了任何错误，`err` 会包含相应的错误信息。如果写入操作成功，`err` 将为 `nil`。这可以用来检测执行 `Fprintf` 过程中是否遇到问题，例如无法写入目标、内存不足或者其他 IO 错误。

		简单来说，使用 `Fprintf` 可以很方便地将格式化后的字符串输出到文件、网络连接或其他任意可以写入的地方，其返回值帮助你确认写入操作的结果。
	*/
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	/*
		 `http.ListenAndServe` 函数的第二个参数期望的是一个实现了 `http.Handler` 接口的值。这个接口定义了一个方法 `ServeHTTP`，该方法用于处理 HTTP 请求。接口定义如下：

		type Handler interface {
		    ServeHTTP(w http.ResponseWriter, r *http.Request)
		}

		如果给 `ListenAndServe` 的第二个参数传递 `nil`，Go 会使用默认的多路复用器 `http.DefaultServeMux` 作为处理器。
		`http.DefaultServeMux` 是 `http.ServeMux` 的一个实例，它本质上是一个 HTTP 请求的路由器（或叫多路复用器），可以根据请求的 URL 的路径部分来将请求分派给不同的处理器函数。

		在您提供的代码示例中：
		- `http.HandleFunc("/", handler)` 函数调用实质上是往 `http.DefaultServeMux` 中注册了一个处理函数 `handler` 来响应路径为 `/` （即所有路径）的 HTTP 请求。
		- 传递 `nil` 作为 `http.ListenAndServe` 的第二个参数，告诉它使用 `http.DefaultServeMux` 作为请求处理器。

		因此，代码的工作流程是这样的：
		1. 当有 HTTP 请求到达时，`http.ListenAndServe` 启动的服务器会使用 `http.DefaultServeMux` 来处理请求。
		2. 由于您已经通过 `http.HandleFunc("/", handler)` 在 `http.DefaultServeMux` 注册了 `handler` 函数，所以对于任何到达服务器的 HTTP 请求，`http.DefaultServeMux` 都会调用 `handler` 函数来处理。
		3. 在 `handler` 函数中，您使用 `fmt.Fprintf` 将输出写入响应 `w`（`http.ResponseWriter`），这样就向客户端发送了响应数据。

		这种方式简化了路由和处理器的配置，对于许多 Web 应用程序来说既方侈又有效。
		如果您需要更复杂的路由逻辑（例如基于不同的请求路径或方法执行不同的处理逻辑），您可能会直接创建一个 `http.ServeMux` 实例，或使用第三方路由库来实现更细粒度的路由控制。
	*/
}

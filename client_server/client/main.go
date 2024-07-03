package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(req.Context())

	resp, err := http.Get("http://localhost:8080")
	// http.Get("http://localhost:8080") 方法尝试对指定的 URL 执行一个 HTTP GET 请求。
	// 这是 net/http 包提供的一个方便的函数，用于快速发起 GET 请求，并返回响应。这个函数返回两个值：resp 和 err。
	// resp 是一个 *http.Response 类型的对象，包含了服务器响应的所有信息，如状态码（StatusCode）、响应头（Header）、响应体（通过 Body 字段访问，它是一个 io.ReadCloser 接口）等。
	// err 表示请求过程中遇到的错误。如果请求成功发出并收到服务器响应，不管状态码是什么，err 都会是 nil。如果发生了网络错误或者不能解析服务器的回应等问题，err 会包含相应的错误信息。
	if err != nil {
		panic(err)
		/*
			panic(err) 语句用来中断当前程序的执行，并输出错误信息。
			panic 是 Go 语言的一个内置函数，用于处理无法恢复的错误情况。
			在实际开发中，你可能会用更温和的错误处理方式代替 panic，比如返回错误给上层调用者。
		*/
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	fmt.Println(respBytes)
	fmt.Println(string(respBytes))
	fmt.Printf("%s", respBytes)
}

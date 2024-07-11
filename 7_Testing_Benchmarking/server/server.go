package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/double", doubleHandler)
	http.ListenAndServe(":4000", nil)
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	// 获取表单数据
	text := r.FormValue("v")
	/*
		在Go语言的Web开发中，`r.FormValue` 是用来从 HTTP 请求中获取表单数据的方法。
		具体来说，`r.FormValue("v")` 会从请求中提取键为 "v" 的表单值。

		- `r` 是一个 `http.Request` 对象，代表了客户端的HTTP请求。
		- `FormValue` 是 `http.Request` 对象的一个方法，用于获取表单数据。它可以获取 URL 查询参数和 POST 表单数据中的值。
		- `"v"` 是你要获取的表单数据的键。

		例如，如果你有一个 HTML 表单：

		```html
		<form action="/double" method="post">
		  <input type="text" name="v" value="123">
		  <input type="submit" value="Submit">
		</form>
		```

		当用户提交这个表单时，浏览器会发送一个 POST 请求到 `/double`，并且请求体中包含表单数据 `v=123`。
		在服务器端，`r.FormValue("v")` 会返回字符串 `"123"`。
	*/

	// 如果客户端发送的请求没有包含键为"v"的参数，或者参数的值为空字符串，那么服务器会返回一个HTTP 400（Bad Request）错误，并显示"Missing value"的错误信息。
	if text == "" {
		http.Error(w, "Missing value", http.StatusBadRequest)
		return
	}

	// 将表单数据转换为整数
	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	if _, err = fmt.Fprintf(w, "%d", v*2); err != nil {
		// if _, err = ...; err != nil 结构用于检查 fmt.Fprintf 的错误返回值
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

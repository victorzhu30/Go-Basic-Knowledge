package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {

	testCases := []struct {
		name               string
		input              string
		expectedResult     string
		expectedStatusCode int
		err                string
	}{
		{name: "double of three", input: "3", expectedResult: "6", expectedStatusCode: http.StatusOK, err: ""},
		{name: "double of three", input: "4", expectedResult: "8", expectedStatusCode: http.StatusOK, err: ""},
	}

	/*
		使用一个for循环来遍历每个测试用例，并使用t.Run来运行每个测试用例。
		t.Run方法允许并发运行测试，并提供了更好的测试输出格式。

		t.Run 方法：
		t.Run(testCase.name, func(t *testing.T) {
		    // 这里是具体的测试逻辑
		})

		t.Run方法接受两个参数：测试用例的名称和一个匿名函数。在这个匿名函数中，我们可以编写具体的测试逻辑来验证代码的正确性。
	*/
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// 创建一个新的请求
			req := httptest.NewRequest("GET", "localhost:4000/double?v="+testCase.input, nil)
			// 创建一个响应记录器
			rec := httptest.NewRecorder()
			// 调用处理函数
			doubleHandler(rec, req)
			// 检查响应状态码
			if rec.Code != testCase.expectedStatusCode {
				t.Errorf("expected status OK; got %v", rec.Code)
			}
			// 检查响应体
			if rec.Body.String() != testCase.expectedResult {
				t.Errorf("expected body %q; got %q", testCase.expectedResult, rec.Body.String())
			}
		})
	}

	// TestdoubleHandler 测试名称格式错误: 'Test' 后面的首个字母不得为小写 TestXxxx 🆗
	// 创建一个新的请求
	_, err := http.NewRequest(http.MethodGet, "localhost:4000/double?v=2", nil)
	req := httptest.NewRequest("GET", "localhost:4000/double?v=2", nil)
	fmt.Println(req)
	// http.MethodGet 是 Go 语言标准库中定义的常量，表示HTTP GET方法。它的值为 "GET" method.go
	// httptest.NewRequest 是 net/http/httptest 包中的一个函数，用于在测试中创建一个HTTP请求，与 http.NewRequest 类似，但专门用于测试环境。它返回一个 *http.Request 而不返回错误，因此使用起来更简便。
	if err != nil {
		t.Fatalf("cound not create a new request: %v, err: %v", req, err)
	}

	// 创建一个响应记录器
	rec := httptest.NewRecorder()

	fmt.Println(rec, rec.Code, rec.Body)
	// httptest.NewRecorder() 返回一个 *httptest.ResponseRecorder 类型的指针。打印这样的指针会显示该结构体的详细内容。

	// 调用处理函数
	doubleHandler(rec, req)
	// StatusOK                   = 200 // RFC 9110, 15.3.1 status.go
	fmt.Println(rec, rec.Code, rec.Body)

	// 检查响应状态码
	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", rec.Code)
	}
	// 检查响应体
	if rec.Body.String() != "4" {
		t.Errorf("expected body %q; got %q", "4", rec.Body.String())
	}

	res := rec.Result()
	// rec.Result() 会返回一个 *http.Response，它表示 httptest.ResponseRecorder当前的HTTP响应。这个响应包括了所有头信息、状态码和响应体数据。

	fmt.Println(res)
	if res.StatusCode != http.StatusOK {
		t.Errorf("received status code %d, expect %d", res.StatusCode, http.StatusOK)
	}

	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("cannot read all from the response body, err %v", err)
	}

	result, err := strconv.Atoi(string(resBytes))
	if err != nil {
		t.Fatalf("cannot convert response body to int, err %v", err)
	}

	if result != 4 {
		t.Fatalf("expected 4, got %d", result)
	}
}

/*
在Go语言中，单元测试是一种用来验证代码功能的小型测试。Go语言的标准库提供了一个名为testing的包，用来编写和运行单元测试。

测试文件：测试文件的文件名通常以 _test.go 结尾。
测试函数：测试函数的名字以 Test 开头，且接受一个 *testing.T 参数。例如：
运行测试：使用 go test 命令运行测试。这将运行所有以 _test.go 结尾文件中的测试函数，并报告测试结果。
*/

/*

在Go语言中，*testing.T 是一个指向 testing.T 类型的指针，testing.T 类型用于管理测试的状态并支持格式化的测试输出。
*testing.T 作为参数传递给每个测试函数，使得测试函数可以使用其方法来记录测试的成功与失败。

*testing.T 提供了多个方法来报告测试结果和日志信息。常用的方法包括：

t.Log：记录测试日志，使用 t.Log 方法记录的信息在测试失败时会显示出来。
t.Error：报告测试错误，但不会立即停止测试，允许测试函数继续执行。
t.Errorf：类似于 t.Error，但提供格式化输出。
t.Fatal：报告测试失败，并立即停止执行当前测试函数。
t.Fatalf：类似于 t.Fatal，但提供格式化输出。
*/

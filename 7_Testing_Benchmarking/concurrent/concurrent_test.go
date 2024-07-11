package main

import "testing"

func BenchmarkConcurrentAtomicAdd(b *testing.B) {
	b.ResetTimer()
	// 复位计时器：在基准测试中，使用b.ResetTimer()复位计时器，以确保初始化代码的时间不被计入测试结果中。
	println(b.N)
	for i := 0; i < b.N; i++ {
		ConcurrentAtomicAdd()
	}
}

/*
### 1. 环境设置信息

GOROOT=D:\go1.22.4 #gosetup
GOPATH=C:\Users\Victor\go #gosetup
D:\go1.22.4\bin\go.exe test -c -o C:\Users\Victor\AppData\Local\JetBrains\GoLand2024.1\tmp\GoLand\___7__Testing_Benchmarking_concurrent__BenchmarkConcurrentAtomicAdd.test.exe 7__Testing_Benchmarking/concurrent #gosetup
C:\Users\Victor\AppData\Local\JetBrains\GoLand2024.1\tmp\GoLand\___7__Testing_Benchmarking_concurrent__BenchmarkConcurrentAtomicAdd.test.exe -test.v -test.paniconexit0 -test.bench ^\QBenchmarkConcurrentAtomicAdd\E$ -test.run ^$ #gosetup

这部分输出是IDE（JetBrains GoLand）设置和执行测试的相关信息，包括`GOROOT`、`GOPATH`等环境变量的设置，以及`go test`命令的执行路径和参数。

### 2. 基准测试的输出

goos: windows
goarch: amd64
pkg: 7__Testing_Benchmarking/concurrent
cpu: 13th Gen Intel(R) Core(TM) i7-13700H

这部分输出是基准测试的结果。我们逐行解释：

- `goos: windows`：操作系统是Windows。
- `goarch: amd64`：处理器架构是AMD64。
- `pkg: 7__Testing_Benchmarking/concurrent`：正在测试的包名是`7__Testing_Benchmarking/concurrent`。
- `cpu: 13th Gen Intel(R) Core(TM) i7-13700H`：CPU信息。

### 3. `b.N`的打印值

100
BenchmarkConcurrentAtomicAdd
6111

- `100`：这是第一个打印值`println(b.N)`的输出，它表示基准测试的初始运行次数`b.N`是100。基准测试框架会多次调整`b.N`的值以确定最佳的运行次数以获取稳定的结果。
- `BenchmarkConcurrentAtomicAdd`：基准测试的名称。
- `6111`：这是经过调整后的`b.N`的最终值，表示基准测试最终运行了6111次。

### 4. 基准测试结果

BenchmarkConcurrentAtomicAdd-20    	    6111	    185624 ns/op
PASS

- `BenchmarkConcurrentAtomicAdd-20`：基准测试的名称以及运行时使用的并发goroutine数量（这里是20）。
- `6111`：表示基准测试运行了6111次。
- `185624 ns/op`：每次操作的平均时间是185624纳秒。
- `PASS`：表示基准测试通过。

### 综上所述

基准测试`BenchmarkConcurrentAtomicAdd`在一台配置为13代英特尔i7-13700H的Windows机器上运行。
测试框架首先设置`b.N`为100次运行以预热，然后调整为6111次运行以获取稳定的结果，最终每次运行的平均时间是185624纳秒。
*/

func BenchmarkConcurrentMutexAdd(b *testing.B) {
	b.ResetTimer()
	println(b.N)
	for i := 0; i < b.N; i++ {
		ConcurrentMutexAdd()
	}
}

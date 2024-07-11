package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func say(id string) {
	time.Sleep(time.Second)
	fmt.Println("I am done! id:" + id)
	wg.Done()
}

func player(name string, ch chan int) {

	defer wg.Done()

	for {
		ball, ok := <-ch //怎样从通道里拿值
		if !ok {         // 通道关闭
			fmt.Printf("channel is closed! %s wins!\n", name)
			return
		}

		n := rand.Intn(100)

		if n%10 == 0 {
			// 把球大飞
			close(ch)
			fmt.Printf("%s missed the ball! %s loses!\n", name, name)
			return
		}
		ball++
		fmt.Println(name, ball)
		ch <- ball // 把球传给对手
	}

	//defer wg.Done()
}

var wg sync.WaitGroup

var counter int32

var mtx sync.Mutex

var ch = make(chan int, 0)

// 把0改成1避免死锁

func UnsafeIncCouter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		counter++
	}
}

func MutexIncCouter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mtx.Lock()
		counter++
		mtx.Unlock()
	}
}

func AtomicIncCouter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func ChannelIncCouter() {
	defer wg.Done()
	count := <-ch
	count++
	ch <- count
}

func main() {
	//go say("hello")
	//go say("world")
	/*
		主 Goroutine 退出：
		Goroutine 是并发执行的，但如果主 Goroutine 退出，程序将立即结束，所有未完成的 Goroutine（包括子 Goroutine）也将停止。
	*/
	//say("world")

	//wg.Add(3) // 总共有2个任务
	//
	//go say("hello")
	//go say("world")
	//go func(id string) {
	//	fmt.Println(id)
	//	wg.Done()
	//}("hi")
	//
	//wg.Wait() // 等待所有任务完成，卡住如果wg不是0
	//fmt.Println("All done!")

	//wg.Add(2)
	//
	//// channel
	//ch := make(chan int, 0) // unbuffered channel
	//
	////ch <- 0 // main卡在这 无接受者
	//
	//go player("heli", ch)
	//go player("chong", ch)
	//
	//ch <- 0 // 开球
	//
	//wg.Wait()
	wg.Add(2)
	//go UnsafeIncCouter()
	//go UnsafeIncCouter()

	//go MutexIncCouter()
	//go MutexIncCouter()

	//go AtomicIncCouter()
	//go AtomicIncCouter()

	go ChannelIncCouter()
	go ChannelIncCouter()

	ch <- 0

	wg.Wait()
	fmt.Println(counter)
}

/*
在Go语言中，defer语句会将函数压入栈中，当外围函数执行完毕后，这些被压入栈的函数才会顺序执行（后进先出：LIFO）。
`defer wg.Done()`语句是在函数`player`退出时才会执行的，无论函数是正常结束还是由于return提前结束。
defer就是在包含它的函数结束时（也就是return之后）才会被执行，不管函数是从哪个路径返回的，defer都能确保被调用。
在你的代码中，`ball++`、`fmt.Println(name, ball)`和`ch <- ball`在一个无限循环中，这意味着，除非函数在某个地方返回（即其中一个if语句成立并执行return），否则这个循环一直会运行，不会有机会给`defer wg.Done()`执行。
所以，每次`player`函数结束，不管是正常结束还是因为条件触发了`return`，`defer wg.Done()`都会执行，`WaitGroup`的计数器都会减一。
简单来说，`defer wg.Done()`的执行不依赖于特定的if条件，而是依赖于函数的结束。
只要函数结束了，无论如何它都会被执行。因此，不需要担心`wg`会不会减一，因为每个启动的`player`函数在结束的时候都会保证把`wg`减一。
*/

/*
你的代码虽然是在并发环境中运行，但由于使用了无缓冲通道，所以这段代码不会出现一个球员连续接球的情况。
在Go语言中，无缓冲通道有一个特性，那就是发送操作会阻塞，直到有其他Goroutine执行接收操作，才会解除阻塞。
因此，当一个球员发送（或说投递）球到通道中之后，这个球员的Goroutine就会阻塞，不能再执行下一个循环，除非有另一个球员接（或说取）出这个球。
换句话说，球员在投递完球后必须等待对方接球，自己无法再次投递，因此无法实现一个人连续接球。
对于你的乒乓球比赛模拟，这种特性正好能确保每次只有一个球员在打球，不会出现一个球员连续打球的情况。
这也正是无缓冲通道在Go并发编程中常常被用来作为同步工具的原因。
*/

/*
具体来说，两个goroutine（也就是乒乓球比赛中的两个球员）一开始都会因为等待从通道接收数据而阻塞。
然后，当主函数把一个0发送到通道中时，其中一个goroutine会接收到这个0，并开始执行。
这个goroutine会增加球的计数（比如，0变为1），然后把新的计数发送回通道。
然后这个goroutine就会再次阻塞，等待从通道中接收新的数据。
而另一个goroutine此时就能从通道中接收到这个1，然后同样地增加球的计数并发送回通道。
因为通道的这种阻塞特性，两个goroutine就会以这种方式轮流执行，模拟了乒乓球比赛中球员交替击球的情况。
直到其中一个goroutine遇到一个满足退出条件的随机数，使得它结束执行并关闭通道。
*/

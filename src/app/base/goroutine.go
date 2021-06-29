package util

import (
	"fmt"
	"sync"
	"time"
)

func MyTest() {

	// pipeTest1()
	// pipeTest2()
	// pipeTest3()
	// pipeTest4()
	// pipeTest5()
	pipeTest6()

}

func ss(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y
	}
	close(mychan)
}

func pipeTest3() {
	pipline := make(chan int, 10)
	go ss(pipline)

	for k := range pipline {
		fmt.Println(k)
	}

}

func pipeTest1() {
	pipline := make(chan int, 10)
	fmt.Println("信道可缓冲 %d 个数据\n", cap(pipline))
	pipline <- -1
	fmt.Println("信道中当前有 %d 个数据", len(pipline))
}

func pipeTest2() {
	pipline := make(chan int)

	go func() {
		fmt.Println("准备发送数据")
		pipline <- 2000
		fmt.Println("信道中当前有 %d 个数据", len(pipline))
	}()

	go func() {
		println("准备接收数据")
		data := <-pipline
		println("接收到data:", data)
	}()

	time.Sleep(time.Second)
}

func test() {
	fmt.Println("mytest")
	go mygo("协程1号")
	go mygo("协程2号")
	time.Sleep(time.Second)
}

func mygo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine %s\n", name)
		// 为了避免第一个协程执行过快，观察不到并发的效果，加个休眠
		time.Sleep(10 * time.Millisecond)

	}
}

func pipeTest4() {
	pipline := make(chan bool, 1)
	var x int
	for i := 0; i < 1000; i++ {
		go increment(pipline, &x)
	}
	time.Sleep(time.Second)
	fmt.Println("x 的值：", x)
}

func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}

func pipeTest5() {
	done := make(chan bool, 1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}

func pipeTest6() {
	var wg sync.WaitGroup
	wg.Add(2)
	go woker(5, &wg)
	go woker(10, &wg)
	wg.Wait()
}

func woker(x int, wg *sync.WaitGroup) {
	wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}

}

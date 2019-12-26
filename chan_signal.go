package main

import (
	"fmt"
)

var (
	ErrRunning  = fmt.Errorf("service running")
	ErrWillStop = fmt.Errorf("service will stop")
	ErrStopped  = fmt.Errorf("service stopped")
)

func main08() {
	var stop chan int
	stop = make(chan int)
	// wg := sync.WaitGroup{}

	// defer close(stop)

Loop:
	for {
		select {
		case <-stop:
			break Loop
		default:
			// go func() {
			// wg.Add(1)
			// defer wg.Done()
			do(stop)
			// }()
		}
		fmt.Println(ErrRunning)
	}

	fmt.Println(ErrStopped)
	// wg.Wait()
}

func do(stop chan int) {
	// stop <- 1
	// fmt.Println(ErrWillStop)
	close(stop) // 关闭channel时, 发送整型:(stop<-0), <-stop输入0
}

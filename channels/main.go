package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int, 1)

	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)
	// recive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0

		// myCh <- 5
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}

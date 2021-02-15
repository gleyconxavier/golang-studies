package racecondition

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

// Call - Calls the racecondition contents
func Call() {

	fmt.Println("Race Condition")
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	/*
		When two or more Go Routines have access not syncronized to
		a shared resource, they try to read and write in this resource
		at same time. Race condition are the motive that concurrency programming
		is complicated and have more potential to bugs.

		Read and Write operations in a shared resource, must be atomic, or in other words
		must be done by one goroutine per time.
	*/
	wg.Wait()
	fmt.Println("Final Counter:", counter)

	return
}

// incCounter with Race Condition bug
// func incCounter(id int) {
// 	defer wg.Done()

// 	for count := 0; count < 2; count++ {
// 		value := counter

// 		// Liberate the goroutine and puts back in stack
// 		runtime.Gosched()

// 		value++

// 		counter = value
// 	}
// }

// incCounter - Fixed using Atomic
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Code commented we don't need anymore in this case
		// value := counter

		// Counting with Atomic on a safe mode
		atomic.AddInt64(&counter, 1)

		// Liberate the goroutine and puts back in stack
		runtime.Gosched()

		// value++

		// counter = value
	}
}

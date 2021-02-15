package mutexes

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup

	// It's use to define a critical section in code
	mutex sync.Mutex
)

// Call - Calls the mutexes content
func Call() {
	wg.Add(2)

	fmt.Println("Mutexes")

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)

	return
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Allows only one gorountine in this critical section
		mutex.Lock()
		{
			// Capture the counter value
			value := counter

			// Liberate the thread and put in the stack
			runtime.Gosched()

			// Increase local counter value
			value++

			// Store the value again in on counter
			counter = value
		}
		mutex.Unlock()
		// Liberate the lock and allows any goroutine waiting continue
	}
}

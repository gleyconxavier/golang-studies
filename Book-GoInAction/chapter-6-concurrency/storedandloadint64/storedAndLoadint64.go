package storedandloadint64

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

// Call - Calls the StoredInt64 & LoadInt64 contents
func Call() {

	fmt.Println("StoredInt64 & LoadInt64")

	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// Give time to gorountines finish
	time.Sleep(1 * time.Second)

	// Sinalizes in safe mode that's time to end
	fmt.Println("Shutdown now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()

	return
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Verify if it's time to end
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

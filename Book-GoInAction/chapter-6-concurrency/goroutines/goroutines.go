package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

// Call - Calls the concurrency content
func Call() {

	fmt.Println("Go Routines")

	/*
		Go concurrency (goroutines) runs in a layer above the processor threads,
		and use a paradigm called CSP(Communicating Sequential Process).
		CSP is a model that works through tha data communication between goroutines.
		The OS escalonate Threads to be executed by physical processors, Go escalonates
		Go routines to be executed by logical processors.

		If a Go Routine need to make a call through R/W a network, this goroutine gonna
		be dessociated of the logical processor, and it's transfer to the integrated
		network poller of the runtime. When it's ready it's.

		Concurrency is diferent of paralelism, in paralelism a bunch of codes is
		executed at same time, through physical processors, paralelism is about to do
		a bunch of tasks at same time, concurrency it's about administrate a bunch of tasks.
	*/

	// Allocates one logical processor
	runtime.GOMAXPROCS(1)

	// Wg utilize to wait the program ends
	// Add a counter equal two to the two goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Go Routines")

	go func() {
		// Escalonate the wg.Done call to inform main we finished
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		// Escalonate the wg.Done call to inform main we finished
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating program")

	return
}

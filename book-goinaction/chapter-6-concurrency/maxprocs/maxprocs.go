package maxprocs

import (
	"fmt"
	"runtime"
	"sync"
)

// Call - Calls the contents of maxprocs
func Call() {

	fmt.Println("GOMAXPROCS")

	// With two procs, things start to get messy(Paralelism)
	runtime.GOMAXPROCS(2)

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

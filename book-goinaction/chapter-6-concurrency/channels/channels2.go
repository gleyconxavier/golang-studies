package channels

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func CallChannels2() {
	fmt.Println("\nChannels 2nd part")

	// Creat a channel without buffering
	baton := make(chan int)

	// Add counter equal to one to the last runner
	wg2.Add(1)

	// First runner in the mark
	go Runner(baton)

	// Start the race
	baton <- 1

	// Wait the race end
	wg2.Wait()

}

// Runner simulates a person running in the circuit
func Runner(baton chan int) {

	var newRunner int

	// Waits to receive the baton
	runner := <-baton

	// Start to run in the circuit
	fmt.Printf("Runner %d Running with Baton\n", runner)

	// New runner on the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// Running through the circuit
	time.Sleep(100 * time.Millisecond)

	// Verifica se a corrida terminou
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg2.Done()
		return
	}

	// Change the baton to the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)

	baton <- newRunner
}

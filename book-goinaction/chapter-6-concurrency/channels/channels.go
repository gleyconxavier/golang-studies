package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Call - Calls the channels content
func Call() {
	fmt.Println("Channels")

	// /*
	// 	Channels is act as a conductor between goroutines,
	// 	to create a channel we can use the function make
	// */

	// // Integer channel without buffer
	// unbuffered := make(chan int)

	// // String channel buffered
	// buffered := make(chan string, 10)

	// // sending a value or pointer on a channel needs the use of the operator <-
	// buffered <- "Gopher"

	// // Receive a string from the channel
	// value := <-buffered

	// // A unbuffered channel it's a channel without the hability to store values,
	// // till they have received one

	// This example simulates a tennis game between goroutines
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	// Start the set
	court <- 1

	// Wait the game end
	wg.Wait()

	return
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// Waits the ball throw to us
		ball, ok := <-court
		if !ok {
			// If channel close we win
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Chose a rand number to see if we miss the ball
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to sinalize we lost
			close(court)
			return
		}

		// Show the hit counter and increase by one
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Throw the ball again to the adversary
		court <- ball
	}
	return
}

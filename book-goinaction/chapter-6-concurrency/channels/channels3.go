package channels

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10 // Quantity of tasks to being process
)

func init() {
	rand.Seed(time.Now().Unix())
}

func CallChannels3() {

	// Create a buffered channel
	tasks := make(chan string, taskLoad)

	// Start goroutines to process the tasks
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// Add some tasks to be process
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// Close the channel for goroutines end
	// when all the tasks be finished
	close(tasks)

	// Await all the tasks end
	wg.Wait()

	return
}

func worker(tasks chan string, worker int) {
	// Inform what we have done
	defer wg.Done()

	for {
		// Await all the tasks being atribuited
		task, ok := <-tasks
		if !ok {
			// That mean the channel is empty and closed
			fmt.Printf("Worker: %d: Shutting Down\n", worker)
			return
		}

		// Shows that we are staring a task
		fmt.Printf("Worker: %d: Started %s\n", worker, task)

		// Await randomly to simulate the task time execution
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Shows that we end the task
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

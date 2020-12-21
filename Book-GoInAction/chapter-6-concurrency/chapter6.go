package chapter6

import (
	"fmt"

	"github.com/golang-studies/Book-GoInAction/chapter-6-concurrency/mutexes"
)

// Run - Run the calls on chapter 6
func Run() {

	fmt.Println("Chapter 6 - Concurrency")

	// goroutines.Call()
	// maxprocs.Call()
	// racecondition.Call()
	// storedandloadint64.Call()
	mutexes.Call()

	return
}

package chapter6

import (
	"fmt"

	"github.com/golang-studies/Book-GoInAction/chapter-6-concurrency/channels"
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
	channels.Call()
	channels.CallChannels2()
	channels.CallChannels3()

	return
}

package chapter5

import (
	"fmt"

	"github.com/golang-studies/Book-GoInAction/chapter-5-Types/interfaces"
	"github.com/golang-studies/Book-GoInAction/chapter-5-Types/methods"
	"github.com/golang-studies/Book-GoInAction/chapter-5-Types/structs"
)

// Run - Run the calls on chapter 5
func Run() {
	fmt.Println("Chapter 5")

	methods.Call()
	structs.Call()
	interfaces.Call()

	return
}

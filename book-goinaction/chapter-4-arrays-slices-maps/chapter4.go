package chapter4

import (
	"fmt"

	"github.com/golang-studies/Book-GoInAction/chapter-4-arrays-slices-maps/arrays"
	"github.com/golang-studies/Book-GoInAction/chapter-4-arrays-slices-maps/maps"
	"github.com/golang-studies/Book-GoInAction/chapter-4-arrays-slices-maps/slices"
)

func Run() {
	fmt.Println("Chapter 4")

	arrays.Call()
	maps.Call()
	slices.Call()

	return
}

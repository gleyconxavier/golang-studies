package chapter4

import (
	"fmt"

	"github.com/golang-estudos/Book-GoInAction/chapter-4-arrays-slices-maps/arrays"
	"github.com/golang-estudos/Book-GoInAction/chapter-4-arrays-slices-maps/maps"
	"github.com/golang-estudos/Book-GoInAction/chapter-4-arrays-slices-maps/slices"
)

func Run() {
	fmt.Println("Chapter4")

	arrays.Call()
	slices.Call()
	maps.Call()

	return
}

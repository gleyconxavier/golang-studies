package main

import "fmt"

func main() {

	fmt.Println("Chapter 3 - Arrays, Slices and Maps")

	/*
		Array declaration sintax with values:
		Check that arrays have a max length predetermined, if the values at
		those positions are not informed, they will be evalued with the default
		value of the array type.
	*/
	array := [5]int{10, 20, 30}
	fmt.Println("Array declared:", array)

	/*
		If we inform values beyound the max length, Go compiler will not accept
		We can also use "go vet" command to check this concept out
	*/
	// array2 := [5]int{10, 20, 30, 40, 50, 60}
	// fmt.Println("Array declared:", array2)

	// We can also declared a array, calculating it's max length
	array3 := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Array declared with calculated length:", array3)

	return
}

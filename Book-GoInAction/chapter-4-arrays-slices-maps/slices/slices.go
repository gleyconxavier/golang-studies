package slices

import "fmt"

func Call() {

	fmt.Println("Slices")

	/*
		Slices are structures that offer a collection of data and a way to
		administrate that. Slices are built with the concept of dynamic arrays
		that mean they have they size as long as necessary.

		They offer all the indexation vantages, iteration and optimization for
		garbage collector, cause the adjacent memory is allocated in blocks.
	*/

	// Declaring slices with make, this way we can specify the size
	// and capacity

	// When we specify only the size, the capacity will be the same
	slice := make([]string, 5)
	fmt.Println("Slice with only size(5) specified:", slice)

	slice = make([]string, 3, 5)
	fmt.Println("Slice with size(3) and capacity(5) specified:", slice)

	// We cannot specify a slice with less capacity than it's size
	// slice = make([]int, 5, 3)

	/*
		A idiomatic way to declare a slice, it's the same as declaring a array
		but with no value on operator []
	*/
	slice = []string{"Red", "Green", "Blue", "Yellow", "Pink"}
	fmt.Println("Slice of strings:", slice)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice of ints:", slice2)

	// Initialize the 100th element with a empty string
	// Reminder: If you especify on [] operator, will be a array
	slice = []string{99: ""}
	fmt.Println("Slice of strings, with 100th index as empty string:", slice)

	// Declarying nil slices
	var slice3 []int

	// Declarying empty slices
	slice4 := make([]int, 0)

	fmt.Println(slice3, slice4)

	// Creating a slice of slices
	slice5 := []int{10, 20, 30, 40, 50}

	// Create a slice of slice, cutting from index 1 to 3
	sliceOfSlice := slice5[1:3]
	fmt.Println("Slice of slices:", sliceOfSlice)

	return
}

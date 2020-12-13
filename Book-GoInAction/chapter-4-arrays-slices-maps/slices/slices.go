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

	// The slice of a slice, use the same subjacent array but see the array
	// in a different way
	slice5[1] = 11
	fmt.Println("Slice of slices, share the same array from the original slice:",
		sliceOfSlice)

	// Incriasing a slice
	slice6 := []int{10, 20, 30, 40, 50}
	newSlice := slice6[1:3]
	fmt.Println("Slice original:", slice6)
	fmt.Println("Slice of slice original:", newSlice)
	newSlice = append(newSlice, 60)
	fmt.Println("Slice increased:", slice6)
	fmt.Println("Slice of slice increased:", newSlice)

	/*
		When we append a to a newSlice, a value that goes beyond the original
		slice capacity, the capacity is increased by 2x, that escale to the max
		of 1.000 elements. Beyound that, from the 1.001 element, go increase the
		capacity by x1.25, that can change as the language evolve
	*/

	// There's a third index we can use in a slice of slice, to limit
	// it's capacity
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice = source[2:3:4]
	// Formula to know size and capacity(3 - 2 = 1 size, 4 -2 = 2 capacity)
	fmt.Println(slice, len(slice), cap(slice))

	// Vantages of defining size and capazity with the same value
	slice = source[2:3:3]
	slice = append(slice, "kiwi")

	/*
		Without the limit of same size and capacity, the original subjacent
		array would be changed "Banana" to "Kiwi", with that specification
		only the slice of slice receive the new index.
	*/
	fmt.Println("Source slice:", source)
	fmt.Println("Slice of slice:", slice)

	// Slices can be pass through functions without worrying about memory

	return
}

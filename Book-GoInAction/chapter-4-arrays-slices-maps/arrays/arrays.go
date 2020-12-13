package arrays

import "fmt"

// Return the concepts of the Chapter 4, about arrays.
func Call() {
	fmt.Println("Chapter 4 - Arrays")

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
	fmt.Println("Array3 declared with calculated length:", array3)

	// Declaring a array with specific elements
	array4 := [5]int{1: 10, 2: 20}
	fmt.Println("Array4 declared with specific elements:", array4)

	// Accessing array elements
	array5 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array5 declared:", array5)

	// Changing value of array5 at index 2
	array5[2] = 35
	fmt.Println("Array5 received value 35 at index 2:", array5)

	// We can also have a array of pointers, the new bult-in function allocate
	// memory to create the pointer
	array6 := [5]*int{0: new(int), 1: new(int)}
	fmt.Println("Array6, array of int pointers:", array6)

	// Atribuing values to the pointers
	*array6[0] = 10
	*array6[1] = 20
	fmt.Println("Array6, received 10 at index 0 and 20 at index 1:\n", array6)
	/*
		This is only to be clear what a *(pointer) represent
		and a &(ampersand) represent
	*/
	fmt.Println("*array6[1], represent memory value:", *array6[1])
	fmt.Println("&array6[1], represent memory adress:", &array6[1])

	/*
		In Go, arrays are a value, with that in mind we can use them in
		operations of atribuation, only arrays of same type can be atribuated
		one to another
	*/
	var array7 [5]string
	array8 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	array7 = array8

	// We have made a complete copy of array8, check the memory adresses
	fmt.Println("array7 index 1 memory adress:", &array7[1])
	fmt.Println("array8 index 1 memory adress:", &array8[1])

	/*
		Error of atribuiting arrays with different types, Go undestand
		that [4]string and [5]string are different types
	*/
	// var array10 [4]string
	// array9 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// array10 = array9

	// Obs: Copying a array of pointers, copy the pointers values
	var array11 [3]*string
	array12 := [3]*string{new(string), new(string), new(string)}

	*array12[0] = "Red"
	*array12[1] = "Green"
	*array12[2] = "Blue"

	array11 = array12

	// Two arrays pointing to the same value but different adresses
	fmt.Println(&array12[0], &array11[0])

	// Declare bidimensional arrays (matrix) 4x2
	var array13 [4][2]int
	fmt.Println(array13)

	// Use a literal array to declare and initialize a bidimensional array
	array14 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array14)

	// Declare and initialize the values of the external array
	array15 := [4][2]int{1: {20, 21}, 2: {40, 41}}
	fmt.Println(array15)

	// Declare and initialize the values of the internal array
	array16 := [4][2]int{1: {0: 21}, 3: {1: 41}}
	fmt.Println(array16)

	// Accessing a bidimensional array value
	fmt.Println(array16[1][0])

	// Copying bidimensional arrays 2x2
	var matrix1 [2][2]int
	var matrix2 [2][2]int

	matrix2[0][0] = 10
	matrix2[0][1] = 20
	matrix2[1][0] = 30
	matrix2[1][1] = 40

	matrix1 = matrix2

	fmt.Printf("matrix1: %v\nmatrix2: %v\n", matrix1, matrix2)

	// Copy a index to another array of the same type
	var matrix3 [2]int = matrix1[1]
	fmt.Println(matrix3)

	/*
		Passing a array between functions can be a operation who use a lot of
		resources, cause functions pass variables by value. If your variable are
		a array, that mean your array idenpendent of it size, will be copy and
		passed to the function.
	*/

	// array with a million elements
	var array1e6 [1e6]int

	// This will pass that big array and create a copy in the function 8MB
	foo(array1e6)

	// This will pass a pointer of that big array to the function 8bytes
	fooPointingIsBetter(&array1e6)

	/*
		Slices deal with those problems in a inherint way
	*/
}

// Warning this can be very dangerous ;)
func foo(array [1e6]int) {
	fmt.Println("Passed a 8MB array and created a copy, results in 16MB")
	return
}

// Pointing is a better way, but beware that this variable is shared
// and that can be tricky
func fooPointingIsBetter(array *[1e6]int) {
	fmt.Println("Passed a 8bytes pointer of a 8MB array")
	return
}

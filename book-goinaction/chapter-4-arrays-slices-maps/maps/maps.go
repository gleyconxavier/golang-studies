package maps

import "fmt"

func Call() {
	fmt.Println("Maps")

	/*
		Maps are unordered data structures that returns pairs of keys and values
		they strong point is fast return data through a key, maps are unordered
		cause they are implemented as hash tables
	*/

	// Declaring maps

	// Map with string key and a int value
	dict := make(map[string]int)

	// Map declared with values
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict, dict2)

	// We can use slice as value in maps
	dict3 := map[int][]string{}
	fmt.Println("Map with empty string slice:", dict3)

	// Error assign value to nil map
	var colors map[string]string
	// colors["Red"] = "#da1337"

	// We can also check if a key exists
	value, exists := colors["blue"]
	fmt.Println(value, exists)

	// Also can remove a key/value pair
	delete(colors, "Coral")

	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	return
}

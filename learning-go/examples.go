package main

import "fmt"

// This file contains examples of basic Go concepts
// Uncomment sections to try them out!

func examples() {
	// Variables
	var name string = "Go"
	age := 25 // Short variable declaration
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Arrays and Slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Maps
	person := map[string]string{
		"name": "Alice",
		"city": "New York",
	}
	fmt.Println("Person:", person)

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Printf("Loop iteration: %d\n", i)
	}

	// Conditionals
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Functions
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)
}

// Function example
func add(a int, b int) int {
	return a + b
}


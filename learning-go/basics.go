package main

import "fmt"

// ============================================================================
// LESSON 1: Variables and Types
// ============================================================================

func lesson1_VariablesAndTypes() {
	fmt.Println("\n=== Lesson 1: Variables and Types ===")
	fmt.Println()

	// Method 1: Explicit type declaration
	var name string = "Alice"
	var age int = 30
	var isActive bool = true

	// Method 2: Type inference (Go figures out the type)
	var city = "New York"
	var temperature = 72.5

	// Method 3: Short variable declaration (most common)
	country := "USA"
	population := 330000000

	// Multiple variable declaration
	var x, y int = 10, 20
	a, b := "hello", "world"

	// Constants
	const pi = 3.14159
	const greeting = "Hello, Go!"

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", name, age, isActive)
	fmt.Printf("City: %s, Temperature: %.1f\n", city, temperature)
	fmt.Printf("Country: %s, Population: %d\n", country, population)
	fmt.Printf("x: %d, y: %d\n", x, y)
	fmt.Printf("a: %s, b: %s\n", a, b)
	fmt.Printf("Pi: %.5f\n", pi)
	fmt.Printf("Greeting: %s\n", greeting)

	// Common Go types
	var (
		text      string  = "text"
		number    int     = 42
		decimal   float64 = 3.14
		flag      bool    = false
		character rune    = 'A' // rune is an alias for int32
	)

	fmt.Printf("\nTypes:\n")
	fmt.Printf("  string: %s\n", text)
	fmt.Printf("  int: %d\n", number)
	fmt.Printf("  float64: %.2f\n", decimal)
	fmt.Printf("  bool: %t\n", flag)
	fmt.Printf("  rune: %c (%d)\n", character, character)
}

// ============================================================================
// LESSON 2: Functions
// ============================================================================

func lesson2_Functions() {
	fmt.Println("\n=== Lesson 2: Functions ===")
	fmt.Println()

	// Call various function examples
	result1 := addNumbers(5, 3)
	result2 := multiply(4, 7)
	result3 := greet("Go Learner")
	sum, product := calculate(10, 5)

	fmt.Printf("addNumbers(5, 3) = %d\n", result1)
	fmt.Printf("multiply(4, 7) = %d\n", result2)
	fmt.Printf("greet() = %s\n", result3)
	fmt.Printf("calculate(10, 5) = sum: %d, product: %d\n", sum, product)

	// Variadic function (accepts variable number of arguments)
	total := sumNumbers(1, 2, 3, 4, 5)
	fmt.Printf("sumNumbers(1,2,3,4,5) = %d\n", total)
}

// Simple function with parameters and return value
func addNumbers(a int, b int) int {
	return a + b
}

// Function with same parameter types (can simplify)
func multiply(a, b int) int {
	return a * b
}

// Function that returns a string
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Function with multiple return values
func calculate(a, b int) (int, int) {
	return a + b, a * b
}

// Variadic function (accepts any number of int arguments)
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// ============================================================================
// LESSON 3: Control Structures (if, for, switch)
// ============================================================================

func lesson3_ControlStructures() {
	fmt.Println("\n=== Lesson 3: Control Structures ===")
	fmt.Println()

	// IF statements
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// IF with initialization statement
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// FOR loops - Method 1: Traditional
	fmt.Println("\nTraditional for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  Iteration %d\n", i)
	}

	// FOR loops - Method 2: While-style
	fmt.Println("\nWhile-style loop:")
	count := 0
	for count < 3 {
		fmt.Printf("  Count: %d\n", count)
		count++
	}

	// FOR loops - Method 3: Infinite loop (with break)
	fmt.Println("\nInfinite loop with break:")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("  i: %d\n", i)
		i++
	}

	// FOR loops - Method 4: Range over slice
	fmt.Println("\nRange over slice:")
	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Printf("  Index %d: %d\n", index, value)
	}

	// FOR loops - Method 5: Range (ignore index)
	fmt.Println("\nRange (values only):")
	for _, value := range numbers {
		fmt.Printf("  Value: %d\n", value)
	}

	// SWITCH statement
	fmt.Println("\nSwitch statement:")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("  Start of the work week")
	case "Friday":
		fmt.Println("  Almost weekend!")
	case "Saturday", "Sunday":
		fmt.Println("  Weekend!")
	default:
		fmt.Println("  Regular day")
	}

	// SWITCH without condition (like if-else chain)
	fmt.Println("\nSwitch without condition:")
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("  Good morning!")
	case hour < 18:
		fmt.Println("  Good afternoon!")
	default:
		fmt.Println("  Good evening!")
	}
}

// ============================================================================
// LESSON 4: Arrays and Slices
// ============================================================================

func lesson4_ArraysAndSlices() {
	fmt.Println("\n=== Lesson 4: Arrays and Slices ===")
	fmt.Println()

	// ARRAYS (fixed size, less commonly used)
	var arr [5]int // Array of 5 integers (all zeros)
	arr[0] = 10
	arr[1] = 20
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Array length: %d\n", len(arr))

	// Array literal
	arr2 := [3]string{"apple", "banana", "cherry"}
	fmt.Printf("String array: %v\n", arr2)

	// SLICES (dynamic size, commonly used)
	// Method 1: Slice literal
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice 1: %v (length: %d, capacity: %d)\n", slice1, len(slice1), cap(slice1))

	// Method 2: Make a slice
	slice2 := make([]int, 3) // length 3, capacity 3
	slice2[0] = 10
	slice2[1] = 20
	slice2[2] = 30
	fmt.Printf("Slice 2: %v\n", slice2)

	// Method 3: Make with capacity
	slice3 := make([]int, 0, 5) // length 0, capacity 5
	fmt.Printf("Slice 3: %v (length: %d, capacity: %d)\n", slice3, len(slice3), cap(slice3))

	// Appending to slices
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2, 3, 4)
	fmt.Printf("Slice 3 after append: %v\n", slice3)

	// Slicing a slice
	original := []int{0, 1, 2, 3, 4, 5}
	subSlice := original[1:4] // elements at index 1, 2, 3
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Sub-slice [1:4]: %v\n", subSlice)

	// Copying slices
	source := []int{1, 2, 3}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Printf("Copied slice: %v\n", destination)

	// Iterating over slices
	fmt.Println("\nIterating over slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("  Index %d: %s\n", i, fruit)
	}
}

// ============================================================================
// LESSON 5: Maps
// ============================================================================

func lesson5_Maps() {
	fmt.Println("\n=== Lesson 5: Maps ===")
	fmt.Println()

	// Method 1: Map literal
	person := map[string]string{
		"name": "Alice",
		"city": "New York",
		"job":  "Developer",
	}
	fmt.Printf("Person map: %v\n", person)

	// Method 2: Make a map
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("Scores: %v\n", scores)

	// Accessing values
	fmt.Printf("Alice's score: %d\n", scores["Alice"])
	fmt.Printf("Person's name: %s\n", person["name"])

	// Checking if key exists
	score, exists := scores["Bob"]
	if exists {
		fmt.Printf("Bob's score exists: %d\n", score)
	} else {
		fmt.Println("Bob's score not found")
	}

	// Deleting from map
	delete(scores, "Charlie")
	fmt.Printf("Scores after deletion: %v\n", scores)

	// Iterating over map
	fmt.Println("\nIterating over map:")
	for key, value := range person {
		fmt.Printf("  %s: %s\n", key, value)
	}

	// Map with different value types
	info := map[string]interface{}{
		"name":   "Alice",
		"age":    30,
		"active": true,
	}
	fmt.Printf("\nMixed types map: %v\n", info)
}

// ============================================================================
// LESSON 6: Structs
// ============================================================================

// Define a struct type
type Person struct {
	Name   string
	Age    int
	City   string
	Active bool
}

type Rectangle struct {
	Width  float64
	Height float64
}

func lesson6_Structs() {
	fmt.Println("\n=== Lesson 6: Structs ===")
	fmt.Println()

	// Method 1: Struct literal
	person1 := Person{
		Name:   "Alice",
		Age:    30,
		City:   "New York",
		Active: true,
	}
	fmt.Printf("Person 1: %+v\n", person1)
	fmt.Printf("Person 1 name: %s\n", person1.Name)

	// Method 2: Struct literal (shorthand, must include all fields in order)
	person2 := Person{"Bob", 25, "Boston", true}
	fmt.Printf("Person 2: %+v\n", person2)

	// Method 3: Create and set fields
	person3 := Person{}
	person3.Name = "Charlie"
	person3.Age = 35
	person3.City = "Chicago"
	person3.Active = false
	fmt.Printf("Person 3: %+v\n", person3)

	// Using struct methods
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	perimeter := rect.Perimeter()
	fmt.Printf("\nRectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", area)
	fmt.Printf("Perimeter: %.2f\n", perimeter)

	// Struct with pointer
	person4 := &Person{Name: "Diana", Age: 28, City: "Seattle", Active: true}
	fmt.Printf("\nPerson 4 (pointer): %+v\n", person4)
	fmt.Printf("Person 4 name: %s\n", person4.Name) // Go automatically dereferences
}

// Methods on structs
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

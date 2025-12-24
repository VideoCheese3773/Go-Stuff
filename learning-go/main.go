package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		// If argument provided, run specific lesson
		runLesson(os.Args[1])
	} else {
		// Show menu
		showMenu()
	}
}

func showMenu() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   Go Basics Learning Path              ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println("\nSelect a lesson to run:")
	fmt.Println("  1 - Variables and Types")
	fmt.Println("  2 - Functions")
	fmt.Println("  3 - Control Structures (if, for, switch)")
	fmt.Println("  4 - Arrays and Slices")
	fmt.Println("  5 - Maps")
	fmt.Println("  6 - Structs")
	fmt.Println("  all - Run all lessons")
	fmt.Println("\nUsage: go run . <lesson_number>")
	fmt.Println("Example: go run . 1")
	fmt.Println("\nOr run all lessons: go run . all")
}

func runLesson(lesson string) {
	switch lesson {
	case "1":
		lesson1_VariablesAndTypes()
	case "2":
		lesson2_Functions()
	case "3":
		lesson3_ControlStructures()
	case "4":
		lesson4_ArraysAndSlices()
	case "5":
		lesson5_Maps()
	case "6":
		lesson6_Structs()
	case "all":
		lesson1_VariablesAndTypes()
		lesson2_Functions()
		lesson3_ControlStructures()
		lesson4_ArraysAndSlices()
		lesson5_Maps()
		lesson6_Structs()
	default:
		fmt.Printf("Unknown lesson: %s\n", lesson)
		fmt.Println("Run without arguments to see the menu.")
	}
}

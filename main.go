package main

import "fmt"

// Define variables at the package level
var (
	shortGoLang = "List Team 1"
	fullGoLang  = "List Team 2"
	rewards     = "Rewards"
)

// Define a slice at the package level
var taskItems = []string{shortGoLang, fullGoLang, rewards}

func main() {
	// Call a function to print the task items
	fmt.Println("Initial List:")
	printTaskItems(taskItems)

	// Correctly update and reassign taskItems
	taskItems = addTaskItem(taskItems, "Testing add function written")
	taskItems = addTaskItem(taskItems, "Running")

	// Print updated list
	fmt.Println("\nUpdated List:")
	printTaskItems(taskItems)
}

// Function to print task items
func printTaskItems(items []string) {
	for index, item := range items {
		fmt.Printf("%d. %s\n", index+1, item)
	}
}

// Function to add a new task item and return updated slice
func addTaskItem(taskItems []string, newTask string) []string {
	return append(taskItems, newTask)
}

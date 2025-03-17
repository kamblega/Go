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
	printTaskItems(taskItems)
	addTaskItem(taskItems, "Testing add funtion written")
}

// Function to print task items
func printTaskItems(items []string) {
	for index, item := range items {
		fmt.Printf("%d. %s\n", index+1, item)
	}
}

func addTaskItem(taskItems []string, newTask string) {
	var updatedTaskItem = append(taskItems, newTask)
	printTaskItems(updatedTaskItem)
}

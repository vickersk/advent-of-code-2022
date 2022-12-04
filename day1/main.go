/*
Day 1: Calorie Counting

Part 1 Answer: 74711
Part 2 Answer: 209481
*/

package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"advent/day1/maxheap"
)

const FilePath = "./day1/input.txt"

// readInputFile reads in the file at FilePath. The data from the file is
// cast and returned as a string.
func readInputFile() (string, error) {
	input, err := os.ReadFile(FilePath)
	if err != nil {
		return "", err
	}
	return string(input), nil
}

// getCalories sums up the number of calories in a given food supply and
// returns the total number. Acts as a helper function for part1 and part2.
func getCalories(food []string) (int, error) {
	calories := 0
	for _, c := range food {

		cal, err := strconv.Atoi(strings.Trim(c, "\r\n"))
		if err != nil {
			return 0, err

		} else {
			calories += cal
		}
	}
	return calories, nil
}

// Corresponds to part 1 of the day 1 puzzle.
// Calculates the number of calories each Elf is carrying and returns the
// maximum number of calories.
func part1() (int, error) {
	// Read the input file
	input, err := readInputFile()
	if err != nil {
		return 0, err
	}

	// Split input on empty lines to group together the food supply of each elf
	elfFood := strings.Split(input, "\n\r")

	maxCalories := 0
	for _, s := range elfFood {

		// Split each elf's food supply into separate calories
		food := strings.Split(s, "\r\n")
		calories, err := getCalories(food)
		if err != nil {
			return 0, nil
		}

		// Update the maximum calories carried by an elf
		if calories > maxCalories {
			maxCalories = calories
		}
	}

	return maxCalories, nil
}

// Corresponds to part 2 of the day 1 puzzle.
// Calculates and returns the number of calories the top three Elves are
// carrying total.
func part2() (int, error) {
	// Read the input file
	input, err := readInputFile()
	if err != nil {
		return 0, err
	}

	// Split input on empty lines to group together the food supply of each elf
	elfFood := strings.Split(input, "\n\r")
	maxHeap := &maxheap.MaxHeap{}

	for _, s := range elfFood {

		// Split each elf's food supply into separate calories
		food := strings.Split(s, "\r\n")
		calories, err := getCalories(food)
		if err != nil {
			return 0, nil
		}

		heap.Push(maxHeap, calories)
	}

	maxCalories := 0

	// Sum up the calories of the top three evles based on the heap
	for i := 0; i < 3; i++ {
		maxCalories += heap.Pop(maxHeap).(int)
	}

	return maxCalories, nil
}

func main() {
	// Maximum number of calories being carried by any of the Elves
	maxCalories, err := part1()
	if err != nil {
		fmt.Printf("The following error occurred in part1:\n%q", err)
		return
	}
	fmt.Printf("The most calories carried is: %d\n", maxCalories)

	// Total number of calories being caried by the top three Elves
	maxCalories, err = part2()
	if err != nil {
		fmt.Printf("The following error occurred in part2:\n%q", err)
		return
	}
	fmt.Printf("The top three Elves are carrying: %d Calories\n", maxCalories)
}

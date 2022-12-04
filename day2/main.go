/*
Day 2: Rock Paper Scissors

Part 1 Answer: 8890
Part 2 Answer: 10238
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FilePath = "./day2/input.txt"

// Corresponds to part 1 of the day 2 puzzle.
// Calculates the total score if everything goes exactly according to
// the strategy guide.
func part1() (int, error) {
	// Maps possible shape selections to outcome scores
	var winLossDraw = map[string]int{
		"A X": 3,
		"A Y": 6,
		"A Z": 0,
		"B X": 0,
		"B Y": 3,
		"B Z": 6,
		"C X": 6,
		"C Y": 0,
		"C Z": 3,
	}

	// Maps shapes to their score value
	var shapeValue = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := 0
	scanner := bufio.NewScanner(file)
	// Sum the score for each round by determining if the round is a win,
	// loss, or draw, and then adding the value of the played shape.
	for scanner.Scan() {
		roundScore := winLossDraw[scanner.Text()]
		shape := strings.Split(scanner.Text(), " ")[1]
		shapeScore := shapeValue[shape]

		score += roundScore + shapeScore
	}

	return score, nil
}

// Corresponds to part 2 of the day 2 puzzle.
// Calculates the total score if everything goes exactly according to
// the strategy guide, in addition to following the Elf's instructions.
func part2() (int, error) {
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var shapeScore = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	var winRules = map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	var loseRules = map[string]string{
		"C": "A",
		"A": "B",
		"B": "C",
	}

	score := 0
	scanner := bufio.NewScanner(file)

	// Sum the score for each round by determining if the round is a win,
	// loss, or draw, and then adding the value of the played shape.
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")

		switch {
		// Round needs to end in a loss
		case round[1] == "X":
			score += 0 + shapeScore[winRules[round[0]]]
		// Round needs to end in a draw
		case round[1] == "Y":
			score += 3 + shapeScore[round[0]]
		// Round needs to end in a win
		case round[1] == "Z":
			score += 6 + shapeScore[loseRules[round[0]]]
		}
	}

	return score, nil
}

func main() {
	// Total score following the strategy guide
	totalScore, err := part1()
	if err != nil {
		fmt.Printf("The following error occurred in part1:\n%q", err)
		return
	}
	fmt.Printf("The total score following the strategy is: %d\n", totalScore)

	// Total score following the strategy guide and the Elf's instructions
	totalScore, err = part2()
	if err != nil {
		fmt.Printf("The following error occured in part2:\n%q", err)
		return
	}
	fmt.Printf("The total score following the strategy and the Elf's instructions is: %d\n", totalScore)
}

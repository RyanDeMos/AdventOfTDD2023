package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	// file, err := os.Open("./Day4/Part1/input/testInput.txt")
	file, err := os.Open("./Day4/Part1/input/InputFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalPoints := 0
	scanner := bufio.NewScanner(file)
	lineCounter := 1
	for scanner.Scan() {
		winningNumbers, ourNumbers := parseLine(scanner.Text())
		matchCounter := getMatchCount(winningNumbers, ourNumbers)

		// fmt.Printf("Line %d points: %d\n", lineCounter, getPointsFromMatchCount(matchCounter))
		// fmt.Printf("Number of matchs: %d\n", matchCounter)
		lineCounter += 1
		totalPoints += getPointsFromMatchCount(matchCounter)
	}
	fmt.Printf("Total points: %d\n", totalPoints)
}

func parseLine(line string) ([]string, []string) {
	colonIndex := strings.Index(line, ":")
	splitWinningFromOurs := strings.Split(line[colonIndex+1:], "|")
	winningNumbers := strings.Trim(splitWinningFromOurs[0], " ")
	ourNumbers := strings.Trim(splitWinningFromOurs[1], " ")
	return strings.Split(winningNumbers, " "), strings.Split(ourNumbers, " ")
}

func getMatchCount(winningNumbers []string, ourNumbers []string) int {
	matchCounter := 0
	for _, winningNumber := range winningNumbers {
		for _, ourNumber := range ourNumbers {
			if winningNumber == ourNumber && !(ourNumber == "" || winningNumber == "") {
				// fmt.Printf("Found matching number %s\n", ourNumber)
				matchCounter += 1
			}
		}
	}
	// fmt.Printf("Number of matchs: %d\n", matchCounter)
	return matchCounter
}

// Matchs: 0 1 2 3 4 5  6
// Points: 0 1 2 4 8 16 32

func getPointsFromMatchCount(matchCount int) int {
	if matchCount == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matchCount-1)))
}

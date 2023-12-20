package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// file, err := os.Open("./Day4/Part1/input/testInput.txt")
	file, err := os.Open("./Day4/Part1/input/InputFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	cardDictionary := map[int]int{}
	for i := 1; i <= 201; i++ {
		cardDictionary[i] = 1
	}

	scanner := bufio.NewScanner(file)
	cardNumber := 1
	for scanner.Scan() {
		winningNumbers, ourNumbers := parseLine(scanner.Text())

		for copies := 1; copies <= cardDictionary[cardNumber]; copies++ {
			matchCounter := getMatchCount(winningNumbers, ourNumbers)
			getCopies(cardDictionary, matchCounter, cardNumber)
		}

		cardNumber += 1
	}

	totalCopies := 0
	for key, copies := range cardDictionary {
		fmt.Printf("Key %v, value %v\n", key, copies)
		totalCopies += copies
	}
	fmt.Print(totalCopies)
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

func getCopies(cardDictionary map[int]int, matchCount int, cardNumber int) {
	for i := 1; i <= matchCount; i++ {
		cardDictionary[cardNumber+i] += 1
	}
}

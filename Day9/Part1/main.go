package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fileLines := ReadFileIntoStringSlice("./Day9/Part1/input/testInput.txt")
	fileLines := ReadFileIntoStringSlice("./Day9/Part1/input/InputFile.txt")
	extrapolatedValueSum := 0
	for idx, line := range fileLines {
		extrapolatedValue := findNextInSequence(line)
		extrapolatedValueSum += extrapolatedValue
		fmt.Printf("History %d has the next number in sequence: %d\n", idx, extrapolatedValue)
	}
	fmt.Printf("Total extrapolated value sum is: %d\n", extrapolatedValueSum)
}

func ReadFileIntoStringSlice(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Get each line in the file into a []string
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	file.Close()

	return fileLines
}

func findNextInSequence(line string) int {
	sequence := parseSingleLine(line)
	diffNotAllZero := true
	nextInSequence := sequence[len(sequence)-1]
	for diffNotAllZero {
		sequence, diffNotAllZero = nextDiff(sequence)
		nextInSequence += sequence[len(sequence)-1]
	}
	return nextInSequence
}

func parseSingleLine(line string) []int {
	sequence := []int{}
	for _, char := range strings.Split(line, " ") {
		number, err := strconv.Atoi(char)
		if err != nil {
			log.Fatal(err)
		}
		sequence = append(sequence, number)
	}
	return sequence
}

func nextDiff(sequence []int) ([]int, bool) {
	nextDifferenceSequence := []int{}
	for idx := 1; idx < len(sequence); idx++ {
		nextdiff := sequence[idx] - sequence[idx-1]
		nextDifferenceSequence = append(nextDifferenceSequence, nextdiff)
	}

	// If any diff is non-zero we need to keep calculating diffs
	for _, diff := range nextDifferenceSequence {
		if diff != 0 {
			return nextDifferenceSequence, true
		}
	}
	return nextDifferenceSequence, false
}

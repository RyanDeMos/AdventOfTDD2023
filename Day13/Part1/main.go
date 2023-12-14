package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	Part1()
}

func Part1() {
	fileLines := ReadFileIntoStringSlice("./Day13/Part1/input/InputFile.txt")
	patterns := getPatterns(fileLines)
	totalPoints := 0
	for _, pattern := range patterns {
		transposedPattern := transpose(pattern)
		matchingRows := findMatchingRows(pattern)
		matchingColumns := findMatchingRows(transposedPattern)

		for idx, matchingRow := range matchingRows {
			isRowReflection := checkForReflection(pattern, matchingRow)
			if isRowReflection {
				totalPoints += getPoints(matchingRows[idx], false)
			}
		}
		for idx, matchingColumn := range matchingColumns {
			isColumnReflection := checkForReflection(transposedPattern, matchingColumn)
			if isColumnReflection {
				totalPoints += getPoints(matchingColumns[idx], true)
			}
		}
	}
	fmt.Printf("Total Points: %v\n", totalPoints)
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

func getPatterns(fileLines []string) [][]string {
	patterns := [][]string{}
	block := []string{}
	for _, line := range fileLines {
		if line != "" {
			block = append(block, line)
		} else {
			patterns = append(patterns, block)
			block = []string{}
		}
		// fmt.Printf("Block: %v\n", block)
	}
	return patterns
}

func transpose(pattern []string) []string {
	xlength := len(pattern[0])
	ylength := len(pattern)
	result := make([]string, xlength)
	for i := range result {
		result[i] = ""
	}
	for i := 0; i < xlength; i++ {
		for j := 0; j < ylength; j++ {
			result[i] += string(pattern[j][i])
		}
	}
	return result
}

func findMatchingRows(pattern []string) []int {
	matchingRows := []int{}
	for i := 1; i < len(pattern); i++ {
		if pattern[i-1] == pattern[i] {
			matchingRows = append(matchingRows, i)
		}
	}
	return matchingRows
}

func checkForReflection(pattern []string, matchingRowIndexs int) bool {
	i := 0
	for i+matchingRowIndexs < len(pattern) {
		if matchingRowIndexs-1-i < 0 {
			return true
		}
		if pattern[matchingRowIndexs+i] != pattern[matchingRowIndexs-1-i] {
			return false
		}
		i += 1
	}
	return true
}

func getPoints(matchingRowIndexs int, transposed bool) int {
	if transposed {
		return matchingRowIndexs
	}
	return (matchingRowIndexs) * 100
}

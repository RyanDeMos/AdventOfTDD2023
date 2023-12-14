package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	Part2()
}

func Part2() {
	// fileLines := ReadFileIntoStringSlice("./Day13/Part2/input/InputFile.txt")
	fileLines := ReadFileIntoStringSlice("./Day13/Part2/input/testInput.txt")
	patterns := getPatterns(fileLines)
	totalPoints := 0
	for _, pattern := range patterns {
		transposedPattern := transpose(pattern)
		matchingRows := findMatchingRows(pattern)
		matchingColumns := findMatchingRows(transposedPattern)
		fmt.Printf("Matching Rows: %v\n Matching Columns: %v\n", matchingRows, matchingColumns)

		newMatchingRows := []int{}
		newMatchingColumns := []int{}
		for _, matchingRow := range matchingRows {
			isRowReflection := checkForReflection(pattern, matchingRow)
			if !isRowReflection {
				newMatchingRows = append(newMatchingRows, matchingRow)
			}
		}
		for _, matchingColumn := range matchingColumns {
			isColumnReflection := checkForReflection(transposedPattern, matchingColumn)
			if !isColumnReflection {
				newMatchingColumns = append(newMatchingColumns, matchingColumn)
			}
		}

		fmt.Printf("New Matching Rows: %v\n New Matching Columns: %v\n", newMatchingRows, newMatchingColumns)
		for idx, matchingRow := range newMatchingRows {
			isRowReflection := checkForReflectionButWithChanges(pattern, matchingRow)
			if isRowReflection {
				totalPoints += getPoints(matchingRows[idx], false)
			}
		}
		for idx, matchingColumn := range newMatchingColumns {
			isColumnReflection := checkForReflectionButWithChanges(transposedPattern, matchingColumn)
			if isColumnReflection {
				totalPoints += getPoints(matchingColumns[idx], true)
			}
		}
		fmt.Printf("Total Points: %v\n", totalPoints)

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

func checkForReflectionButWithChanges(pattern []string, matchingRowIndexs int) bool {
	i := 0
	for i+matchingRowIndexs < len(pattern) {
		if matchingRowIndexs-1-i < 0 {
			return true
		}
		if pattern[matchingRowIndexs+i] != pattern[matchingRowIndexs-1-i] {
			for idx := range pattern[matchingRowIndexs+1] {
				if pattern[matchingRowIndexs+i][idx] != pattern[matchingRowIndexs-1-i][idx] {
					if string(pattern[matchingRowIndexs+i][idx]) == "#" {
						fmt.Printf("Index: %v Pattern[matchingRowIndexs+i]: %v\n", matchingRowIndexs+1, pattern[matchingRowIndexs+i])
						pattern[matchingRowIndexs+i] = pattern[matchingRowIndexs+i][:idx] + "." + pattern[matchingRowIndexs+i][idx+1:]
						fmt.Printf("Pattern[matchingRowIndexs+1]: %v\n", pattern[matchingRowIndexs+i])
						return checkForReflection(pattern, matchingRowIndexs)
					} else {
						fmt.Printf("Index: %v Pattern[matchingRowIndexs+1]: %v\n", matchingRowIndexs+1, pattern[matchingRowIndexs+i])
						pattern[matchingRowIndexs+i] = pattern[matchingRowIndexs+i][:idx] + "#" + pattern[matchingRowIndexs+i][idx+1:]
						fmt.Printf("Pattern[matchingRowIndexs+1]: %v\n", pattern[matchingRowIndexs+1])
						return checkForReflection(pattern, matchingRowIndexs)
					}
				}
			}
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

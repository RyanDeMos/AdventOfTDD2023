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
	fileLines := ReadFileIntoStringSlice("./Day13/Part2/input/InputFile.txt")
	// fileLines := ReadFileIntoStringSlice("./Day13/Part2/input/testInput.txt")
	patterns := getPatterns(fileLines)
	totalPoints := 0
	for idx, pattern := range patterns {
		fmt.Printf("Processing pattern: %v\n", idx)
		changeInPoints := processOnePatternForPart2(pattern)
		fmt.Printf("Change in Points: %v\n", changeInPoints)
		fmt.Printf("\n")
		if changeInPoints == 0 {
			for _, line := range pattern {
				fmt.Printf("%v\n", line)
			}
			panic("A pattern returned with no change in points")
		}
		totalPoints += changeInPoints
	}
	fmt.Printf("Total Points: %v\n", totalPoints)
}

func processOnePatternForPart2(pattern []string) int {
	// Get the transposed pattern for an easier time looking at columns
	transposedPattern := transpose(pattern)

	// Find all locations where there are matching rows
	matchingRows := findMatchingRows(pattern)
	matchingColumns := findMatchingRows(transposedPattern)
	// fmt.Printf("Matching Rows: %v\n Matching Columns: %v\n", matchingRows, matchingColumns)

	// Remove the matching rows that gave the solution for part 1 (as when we change a smudge they will no longer be solutions)
	newMatchingRows, newMatchingColumns := removePart1Matches(pattern, transposedPattern, matchingRows, matchingColumns)

	// Check if any of these give a reflection
	changeInPoints := 0
	// fmt.Printf("New Matching Rows: %v\n New Matching Columns: %v\n", newMatchingRows, newMatchingColumns)
	changeInPoints += checkAllMatchingRowsForReflectionWithSmudges(pattern, newMatchingRows, false)
	changeInPoints += checkAllMatchingRowsForReflectionWithSmudges(transposedPattern, newMatchingColumns, true)

	// If non of those gave a reflection, then we must check to see if we can change a smudge to creatue a pair of matching rows (or columns)
	if changeInPoints == 0 {
		fmt.Printf("Need to smudge to find point of reflection\n")
		smudgedMatchingRows := findMatchingRowsWithSmudges(pattern)
		changeInPoints += checkAllMatchingRowsForReflection(pattern, smudgedMatchingRows, false)

		if changeInPoints == 0 {
			fmt.Printf("Smudge made vertically\n")
			smudgedMatchingColumns := findMatchingRowsWithSmudges(transposedPattern)
			changeInPoints += checkAllMatchingRowsForReflection(transposedPattern, smudgedMatchingColumns, true)
		}
	}
	return changeInPoints
}

func findMatchingRowsWithSmudges(pattern []string) []int {
	matchingRows := []int{}
	for i := 1; i < len(pattern); i++ {
		if pattern[i-1] != pattern[i] {
			changedOneCharBool := false
			for idx := range pattern[i-1] {
				if !changedOneCharBool {
					if pattern[i-1][idx] != pattern[i][idx] {
						changedOneCharBool = true
						if string(pattern[i-1][idx]) == "." {
							newLine := pattern[i-1][:idx] + "#" + pattern[i-1][idx+1:]
							if newLine == pattern[i] {
								matchingRows = append(matchingRows, i)
							}
						} else {
							newLine := pattern[i-1][:idx] + "." + pattern[i-1][idx+1:]
							if newLine == pattern[i] {
								matchingRows = append(matchingRows, i)
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("Matching Rows %v\n", matchingRows)
	return matchingRows
}

func checkAllMatchingRowsForReflectionWithSmudges(pattern []string, matchingRows []int, transposed bool) int {
	changeInPoints := 0
	for idx, matchingRow := range matchingRows {
		isRowReflection := checkForReflectionButWithSmudges(pattern, matchingRow)
		if isRowReflection {
			changeInPoints += getPoints(matchingRows[idx], transposed)
		}
	}
	return changeInPoints
}

func checkAllMatchingRowsForReflection(pattern []string, matchingRows []int, transposed bool) int {
	changeInPoints := 0
	for idx, matchingRow := range matchingRows {
		isRowReflection := checkForReflection(pattern, matchingRow)
		if isRowReflection {
			changeInPoints += getPoints(matchingRows[idx], transposed)
		}
	}
	return changeInPoints
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
	i := 1
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

func checkForReflectionButWithSmudges(pattern []string, matchingRowIndexs int) bool {
	patternCopy := make([]string, len(pattern))
	copy(patternCopy, pattern)
	i := 1
	if matchingRowIndexs == 1 || matchingRowIndexs == len(pattern) {
		return true
	}
	for i+matchingRowIndexs < len(pattern) {
		if matchingRowIndexs-1-i < 0 {
			return true
		}
		if patternCopy[matchingRowIndexs+i] != patternCopy[matchingRowIndexs-1-i] {
			for idx := range patternCopy[matchingRowIndexs+1] {
				if patternCopy[matchingRowIndexs+i][idx] != patternCopy[matchingRowIndexs-1-i][idx] {
					if string(patternCopy[matchingRowIndexs+i][idx]) == "#" {
						patternCopy[matchingRowIndexs+i] = patternCopy[matchingRowIndexs+i][:idx] + "." + patternCopy[matchingRowIndexs+i][idx+1:]
						return checkForReflection(patternCopy, matchingRowIndexs)
					} else {
						patternCopy[matchingRowIndexs+i] = patternCopy[matchingRowIndexs+i][:idx] + "#" + patternCopy[matchingRowIndexs+i][idx+1:]
						return checkForReflection(patternCopy, matchingRowIndexs)
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

func removePart1Matches(pattern []string, transposedPattern []string, matchingRows []int, matchingColumns []int) ([]int, []int) {
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
	return newMatchingRows, newMatchingColumns
}

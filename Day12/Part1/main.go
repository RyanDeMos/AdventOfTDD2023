package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stringCondition struct {
	condition string
	groups    []int
}

func main() {
	fileLines := ReadFileIntoStringSlice("./Day12/Part1/input/testInput.txt")
	fmt.Printf("fileLines:\n %v\n", fileLines)
	stringConditions := parselines(fileLines)
	fmt.Printf("String Conditions: %v\n", stringConditions)
	totalCombinations := getPossibleCombinationsCount(stringConditions)
	fmt.Printf("Total Combinations: %v\n", totalCombinations)
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

func parselines(fileLines []string) []stringCondition {
	stringConditions := []stringCondition{}
	for _, line := range fileLines {
		splitLine := strings.Split(line, " ")
		pattern := splitLine[0]

		groups := strings.Split(splitLine[1], ",")
		groupLengths := []int{}
		for _, group := range groups {
			groupLength, err := strconv.Atoi(group)
			if err != nil {
				log.Fatal(err)
			}
			groupLengths = append(groupLengths, groupLength)
		}
		stringConditions = append(stringConditions, stringCondition{pattern, groupLengths})
	}

	return stringConditions
}

func recursionNightmare(condition string, groups []int) {
	fmt.Printf("\n")
	fmt.Printf("Condition: %v\n", condition)
	fmt.Printf("Groups: %v\n", groups)
	if len(condition) == 0 && len(groups) == 0 {
		count++
		fmt.Printf("COUNT GOT UPDATED")
	} else if len(condition) != 0 {
		if string(condition[0]) == "." {
			recursionNightmare(condition[1:], groups)
			// *counter += recursionNightmare(condition[1:], groups, numberOfValidArrangments)
		}
		if string(condition[0]) == "?" {
			recursionNightmare("."+condition[1:], groups)
			// *counter += recursionNightmare("."+condition[1:], groups, numberOfValidArrangments)
			recursionNightmare("#"+condition[1:], groups)
			// *counter += recursionNightmare("#"+condition[1:], groups, numberOfValidArrangments)
		}
		if string(condition[0]) == "#" && len(groups) > 0 {
			fmt.Print("AAAAAAAAAAAAAA\n")
			groupLength := groups[0]
			fmt.Printf("Group Length: %v\n", groupLength)
			idx := 0
			character := "#"
			for idx < groupLength {
				idx += 1
				if idx < len(condition) {
					character = string(condition[idx])
				}
				if character == "." {
					break
				}
			}
			fmt.Printf("idx: %v\n", idx)
			if idx >= groupLength && groupLength < len(condition) {
				if len(groups) != 0 {
					recursionNightmare(condition[groupLength:], groups[1:])
				}
				// *counter += recursionNightmare(condition[groupLength:], groups[1:], numberOfValidArrangments)
			}
		}
	}

}

var count int = 0

func getPossibleCombinationsCount(stringConditions []stringCondition) int {
	totalCombinations := 0
	for _, stringCondition := range stringConditions {
		count = 0
		fmt.Printf("STARTING NEW LINE AHHHHH\n")
		recursionNightmare(stringCondition.condition, stringCondition.groups)
		totalCombinations += count
	}
	return totalCombinations
}

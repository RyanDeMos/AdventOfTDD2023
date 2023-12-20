package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	threeLines := map[string]string{
		"line1": "............................................................................................................................................",
		"line2": "............................................................................................................................................",
		"line3": "............................................................................................................................................",
	}
	// "............................................................................................................................................",
	//"..........",

	file, err := os.Open("./Day3/Part1/input/InputFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%v\n", threeLines)
	digitSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		updateLines(scanner.Text(), threeLines)
		// fmt.Printf("%v\n", threeLines)
		digitsInLine2 := parseSecondLine(threeLines["line2"])
		// fmt.Printf("%v\n", digitsInLine2)
		for _, digitInfo := range digitsInLine2 {
			if checkSurrondingsForSpecialCharacters(threeLines, digitInfo) {
				digitSum += digitInfo.digit
				// fmt.Printf("Added digit %d\n", digitInfo.digit)
				// fmt.Printf("Current sum is %d\n", digitSum)
			}
		}
	}
	fmt.Printf("Total sum is %d", digitSum)
}

func updateLines(newLine string, threeLines map[string]string) {
	threeLines["line1"] = threeLines["line2"]
	threeLines["line2"] = threeLines["line3"]
	threeLines["line3"] = newLine
}

type digitAndLocations struct {
	digit      int
	startIndex int
	endIndex   int
}

func parseSecondLine(line string) []digitAndLocations {
	digitsInLine := []digitAndLocations{}

	startIdx := 0
	for startIdx < len(line) {
		if unicode.IsDigit(rune(line[startIdx])) {
			endIdx := startIdx + 1
		findEndIdx:
			for endIdx < len(line) {
				if unicode.IsDigit(rune(line[endIdx])) {
					endIdx += 1
				} else {
					break findEndIdx
				}
			}

			foundDigit, err := strconv.Atoi(line[startIdx:endIdx])
			if err != nil {
				log.Fatal(err)
			}

			digitsInLine = append(digitsInLine, digitAndLocations{
				digit:      foundDigit,
				startIndex: startIdx,
				endIndex:   endIdx,
			})
			startIdx = endIdx
		} else {
			startIdx += 1
		}
	}
	return digitsInLine
}

func checkSurrondingsForSpecialCharacters(threeLines map[string]string, digitInformation digitAndLocations) bool {
	specialChars := "!@#$%^&*()_-+=,?/;:"

	var aboveChars string
	var belowChars string
	if (digitInformation.endIndex + 1) < len(threeLines["line1"]) {
		if digitInformation.startIndex != 0 {
			aboveChars = threeLines["line1"][digitInformation.startIndex-1 : digitInformation.endIndex+1]
			belowChars = threeLines["line3"][digitInformation.startIndex-1 : digitInformation.endIndex+1]
		} else {
			aboveChars = threeLines["line1"][digitInformation.startIndex : digitInformation.endIndex+1]
			belowChars = threeLines["line3"][digitInformation.startIndex : digitInformation.endIndex+1]
		}
	} else {
		aboveChars = threeLines["line1"][digitInformation.startIndex-1:]
		belowChars = threeLines["line3"][digitInformation.startIndex-1:]
	}

	var leftChar string
	if digitInformation.startIndex != 0 {
		leftChar = string(threeLines["line2"][digitInformation.startIndex-1])
	} else {
		leftChar = string(threeLines["line2"][digitInformation.startIndex])
	}

	var rightChar string
	if digitInformation.endIndex >= len(threeLines["line2"]) {
		rightChar = string(threeLines["line2"][digitInformation.endIndex-1])
	} else {
		rightChar = string(threeLines["line2"][digitInformation.endIndex])
	}

	for _, specialChar := range specialChars {
		special := string(specialChar)
		if strings.Contains(aboveChars, special) || strings.Contains(belowChars, special) || strings.Contains(leftChar, special) || strings.Contains(rightChar, special) {
			return true
		}
	}
	return false
}

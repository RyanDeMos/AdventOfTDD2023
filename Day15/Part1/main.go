package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// fileLines := ReadFileIntoStringSlice("./Day15/Part1/input/testInput.txt")
	fileLines := ReadFileIntoStringSlice("./Day15/Part1/input/InputFile.txt")
	valueSum := Part1(fileLines)
	fmt.Printf("Part 1 solution: %v\n", valueSum)
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

func Part1(fileLines []string) int {
	// Split line on ,
	sequenceArray := strings.Split(fileLines[0], ",")
	// loop over each thing in split line getting value
	valueSum := 0
	for _, sequence := range sequenceArray {
		// call a function to get value from sequence
		valueSum += getSequenceValue(sequence)
	}
	return valueSum
}

func getSequenceValue(sequence string) int {
	value := 0
	for _, character := range sequence {
		value = getCharacterValue(character, value)
	}
	return value
}

func getCharacterValue(character rune, value int) int {
	// Get ascii number
	asciiNumber := int(character)
	value += asciiNumber
	value *= 17
	value = value % 256
	return value
}

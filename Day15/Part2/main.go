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
	// fileLines := ReadFileIntoStringSlice("./Day15/Part2/input/testInput.txt")
	fileLines := ReadFileIntoStringSlice("./Day15/Part2/input/InputFile.txt")
	Part2(fileLines)
	// fmt.Printf("Part 2 solution: %v\n", valueSum)
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

func Part2(fileLines []string) {
	boxes := map[int][][]string{}
	// Split line on ,
	sequenceArray := strings.Split(fileLines[0], ",")
	// loop over each thing in split line getting value
	for _, sequence := range sequenceArray {
		splitSequence := []string{}

		// Check to see if string contains =
		equalIndex := strings.Index(sequence, "=")
		if equalIndex != -1 {
			splitSequence = strings.Split(sequence, "=")
			addToBox(splitSequence, boxes)
		}

		// Check to see if string contains -
		dashIndex := strings.Index(sequence, "-")
		if dashIndex != -1 {
			splitSequence = strings.Split(sequence, "-")
			removeFromBox(splitSequence[0], boxes)
		}
	}
	fmt.Printf("Boxes: \n%v\n", boxes)
	totalPoints := getValueFromBoxes(boxes)
	fmt.Printf("Total Points: %v\n", totalPoints)

}

func getValueFromBoxes(boxes map[int][][]string) int {
	totalPoints := 0
	for boxNumber, box := range boxes {
		pointsFromBoxNumber := boxNumber + 1
		for idx, lens := range box {
			focalLength := lens[1]
			pointsFromFocalLength, err := strconv.Atoi(focalLength)
			if err != nil {
				log.Fatal(err)
			}
			totalPoints += pointsFromBoxNumber * (idx + 1) * pointsFromFocalLength
		}

	}
	return totalPoints
}

func removeFromBox(label string, boxes map[int][][]string) {
	boxNumber := getSequenceValue(label)
	for lensIndex, lens := range boxes[boxNumber] {
		if lens[0] == label {
			if lensIndex != len(boxes[boxNumber])-1 {
				boxes[boxNumber] = append(boxes[boxNumber][:lensIndex], boxes[boxNumber][lensIndex+1:]...)
			} else {
				boxes[boxNumber] = boxes[boxNumber][:lensIndex]
			}
			return
		}
	}
}

func addToBox(splitSequence []string, boxes map[int][][]string) {
	label := splitSequence[0]
	focalLength := splitSequence[1]
	boxNumber := getSequenceValue(label)
	replaceBool := false
	for _, lens := range boxes[boxNumber] {
		if lens[0] == label {
			lens[1] = focalLength
			replaceBool = true
		}
	}
	if !replaceBool {
		boxes[boxNumber] = append(boxes[boxNumber], splitSequence)
	}
}

func getSequenceValue(sequence string) int {
	boxNumber := 0
	for _, character := range sequence {
		boxNumber = getCharacterValue(character, boxNumber)
	}
	return boxNumber
}

func getCharacterValue(character rune, value int) int {
	// Get ascii number
	asciiNumber := int(character)
	value += asciiNumber
	value *= 17
	value = value % 256
	return value
}

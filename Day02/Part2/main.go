package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

var MaxMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open("./Day2/Part1/input/InputFile.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}

	totalMinimumPower := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedLine := parseLine(scanner.Text())

		colourCounter := map[string]int{
			"green": getMaxColourCount(parsedLine, "green"),
			"blue":  getMaxColourCount(parsedLine, "blue"),
			"red":   getMaxColourCount(parsedLine, "red"),
		}

		minimumPower := colourCounter["green"] * colourCounter["blue"] * colourCounter["red"]
		totalMinimumPower += minimumPower

	}
	fmt.Printf("Total minimum power is: %d", totalMinimumPower)
}

func getGameID(line string) int {
	gameID := line[5:]
	gameIDInt, err := strconv.Atoi(gameID)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return gameIDInt
}

func getMaxColourCount(parsedLine []string, colour string) int {
	ColourCount := 0
	for _, sections := range parsedLine {
		if idx := strings.Index(sections, colour); idx != -1 {
			InstanceCount, err := strconv.Atoi(sections[1 : idx-1])
			if err != nil {
				log.Fatal(err)
			}
			if InstanceCount > ColourCount {
				ColourCount = InstanceCount
			}
		}
	}
	return ColourCount
}

func parseLine(line string) []string {
	return strings.FieldsFunc(line, Split)
}

func Split(r rune) bool {
	return r == ';' || r == ',' || r == ':'
}

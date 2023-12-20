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

type SeedsAndType struct {
	seedNumber int
	typeOfSeed string
}

func main() {
	// file, err := os.Open("./Day5/Part1/input/testInput.txt")
	file, err := os.Open("./Day5/Part1/input/InputFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	lineNumber := 1
	seeds := []SeedsAndType{}
	currentType := "seed"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if lineNumber == 1 {
			seedsList := parseSeeds(scanner.Text())
			for _, seedNumber := range seedsList {
				seeds = append(seeds, SeedsAndType{seedNumber, "seed"})
			}
			fmt.Printf("Second seed: %v\n", seeds[1])
		} else {
			destination, source, rangeLength, newType := parseLine(scanner.Text(), currentType)
			delta := findDelta(destination, source)
			seeds = convertToNext(seeds, destination, delta, rangeLength, newType)
			fmt.Printf("Second seed: %v, Line Number %d\n", seeds[1], lineNumber)
			currentType = newType
		}
		lineNumber += 1
	}

	fmt.Print(seeds)
	minimumLoc := findMinimum(seeds)
	fmt.Printf("\nMinimum location: %v\n", minimumLoc)
}

func parseSeeds(line string) []int {
	seeds := strings.Split(line[7:], " ")
	IntSeeds := []int{}
	for _, seed := range seeds {
		intseed, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatal(err)
		}
		IntSeeds = append(IntSeeds, intseed)
	}

	return IntSeeds
}

func parseLine(line string, currentType string) (int, int, int, string) {
	if line == "" {
		return 0, 0, 0, currentType
	}

	if !unicode.IsDigit(rune(line[0])) {
		newType := strings.Split(line, "-to-")
		newType = strings.Split(newType[1], " ")
		fmt.Printf("New type: %v\n", newType[0])
		return 0, 0, 0, newType[0]
	}

	result := strings.Split(line, " ")

	destination, err := strconv.Atoi(result[0])
	if err != nil {
		log.Fatal(err)
	}

	source, err := strconv.Atoi(result[1])
	if err != nil {
		log.Fatal(err)
	}

	rangeLength, err := strconv.Atoi(result[2])
	if err != nil {
		log.Fatal(err)
	}

	return destination, source, rangeLength, currentType
}

func findDelta(destination int, source int) int {
	return destination - source
}

func convertToNext(seeds []SeedsAndType, destination int, delta int, rangeLength int, newType string) []SeedsAndType {
	modifiedSeeds := []SeedsAndType{}
	for _, seed := range seeds {
		if seed.seedNumber+delta >= destination && seed.seedNumber+delta <= destination+rangeLength-1 {
			if seed.typeOfSeed != newType {
				modifiedSeeds = append(modifiedSeeds, SeedsAndType{seed.seedNumber + delta, newType})
			} else {
				modifiedSeeds = append(modifiedSeeds, seed)
			}
		} else {
			modifiedSeeds = append(modifiedSeeds, seed)
		}
	}
	return modifiedSeeds
}

func findMinimum(seeds []SeedsAndType) int {
	if len(seeds) == 0 {
		log.Fatalf("YOU PASSED AN EMPTY ARRAY DONT DO THAT")
	}
	minimum := seeds[0].seedNumber

	for _, seed := range seeds {
		if seed.seedNumber < minimum {
			minimum = seed.seedNumber
		}
	}

	return minimum
}

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

	// Get each line in the file into a []string
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	file.Close()

	// Used to track the minimum of the locations from each range
	minimums := []int{}

	// Get the seed ranges
	seeds := []SeedsAndType{}
	SeedRanges := parseSeeds(fileLines[0])
	for seedRange := 1; seedRange < len(SeedRanges); seedRange += 2 {
		startLocation := SeedRanges[seedRange-1]
		rangeLength := SeedRanges[seedRange]
		seeds = parseRange(startLocation, rangeLength)
		currentType := "seed"
		for _, line := range fileLines[1:] {
			destination, source, rangeLength, newType := parseLine(line, currentType)
			delta := findDelta(destination, source)
			seeds = convertToNext(seeds, destination, delta, rangeLength, newType)
			currentType = newType
			// fmt.Printf("Second seed: %v, Line Number %d\n", seeds[1], lineNumber)
		}
		minimums = append(minimums, findMinimum(seeds))
		fmt.Printf("%v\n", minimums)
	}

	actualMinimum := seeds[0].seedNumber
	for _, seed := range seeds {
		if seed.seedNumber < actualMinimum {
			actualMinimum = seed.seedNumber
		}
	}
	fmt.Printf("Actual Minimum: %d", actualMinimum)
	// lineNumber := 1
	// seeds := []SeedsAndType{}
	// currentType := "seed"
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	if lineNumber == 1 {
	// 		seedsList := parseSeeds(scanner.Text())
	// 		for _, seedNumber := range seedsList {
	// 			seeds = append(seeds, SeedsAndType{seedNumber, "seed"})
	// 		}
	// 		fmt.Printf("Second seed: %v\n", seeds[1])
	// 	} else {
	// 		destination, source, rangeLength, newType := parseLine(scanner.Text(), currentType)
	// 		delta := findDelta(destination, source)
	// 		seeds = convertToNext(seeds, destination, delta, rangeLength, newType)
	// 		// fmt.Printf("Second seed: %v, Line Number %d\n", seeds[1], lineNumber)
	// 		currentType = newType
	// 	}
	// 	lineNumber += 1
	// }

	// fmt.Print(seeds)
	// minimumLoc := findMinimum(seeds)
	// fmt.Printf("\nMinimum location: %v\n", minimumLoc)
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

	// actualSeedInt := []int{}
	// for idx, number := range IntSeeds {
	// 	if idx%2 == 0 {
	// 		fmt.Printf("idx: %d", idx)
	// 		continue
	// 	} else {
	// 		for j := 0; j < number; j++ {
	// 			actualSeedInt = append(actualSeedInt, IntSeeds[idx-1]+j)
	// 		}
	// 	}
	// }

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

func parseRange(minimum int, length int) []SeedsAndType {
	seedsList := []SeedsAndType{}
	for seed := minimum; seed < minimum+length; seed++ {
		seedsList = append(seedsList, SeedsAndType{seed, "seed"})
	}
	return seedsList
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

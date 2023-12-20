package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type mapping struct {
	destination int
	source      int
	length      int
}

func main() {
	// file, err := os.Open("./Day5/Part2/input/testInput.txt")
	file, err := os.Open("./Day5/Part2/input/InputFile.txt")
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
	minimumLocation := math.MaxInt

	// Get Mappings
	allMappings := parseMappings(fileLines)

	// Get the seed ranges
	SeedRanges := parseSeeds(fileLines[0])
	for i := 0; i < len(SeedRanges); i += 2 {
		// SeedRanges always has two ints, the starting seed and the length of the range
		seedStart, seedRangeLength := SeedRanges[i], SeedRanges[i+1]
		for seed := seedStart; seed <= seedStart+seedRangeLength-1; seed++ {
			// We don't want to change the value of the actual seed as that is what we are looping over
			tmpSeed := seed
			// Loop over each mapping converting it all the way to a location
			for _, nextMapping := range allMappings {
				tmpSeed = convertToNext(nextMapping, tmpSeed)
			}
			if tmpSeed < minimumLocation {
				minimumLocation = tmpSeed
			}
		}
	}
	fmt.Printf("Actual Minimum: %d", minimumLocation)
}

func parseSeeds(line string) []int {
	// The first int on the first line starts at position 7
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

func parseRange(line string) mapping {
	pieces := strings.Split(line, " ")
	if len(pieces) != 3 {
		log.Fatal("Should be length 3")
	}
	// These blocks should always match destination, source, length
	destination, err := strconv.Atoi(pieces[0])
	source, err := strconv.Atoi(pieces[1])
	length, err := strconv.Atoi(pieces[2])
	if err != nil {
		log.Fatal(err)
	}
	return mapping{destination, source, length}
}

func parseMappings(lines []string) [][]mapping {
	var allMappings [][]mapping
	// We increment up by two as we will only increment here when we've found an empty line with the inner for loop
	// Incrementing by two allows us to skip the line with test for with x-to-y the next block is for
	for i := 3; i < len(lines); i += 2 {
		var toNextLines []string
		// Here we loop through the block of the ranges adding the lines
		for i < len(lines) && lines[i] != "" {
			toNextLines = append(toNextLines, lines[i])
			i += 1
		}
		// Now we change those lines into the mapping struct
		var toNext []mapping
		for _, line := range toNextLines {
			toNext = append(toNext, parseRange(line))
		}
		allMappings = append(allMappings, toNext)
	}
	return allMappings
}

func convertToNext(rangeMap []mapping, seeds int) int {
	for _, Range := range rangeMap {
		if Range.source <= seeds && seeds <= Range.source+Range.length-1 {
			return Range.destination - Range.source + seeds
		}
	}
	return seeds
}

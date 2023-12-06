package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type mapping struct {
	destination int
	source      int
	length      int
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
	minimumLocation := math.MaxInt

	// Get Mappings
	allMappings := parseMappings(fileLines)

	// Get the seed ranges
	SeedRanges := parseSeeds(fileLines[0])
	for i := 0; i < len(SeedRanges); i += 2 {
		seedStart, seedRangeLength := SeedRanges[i], SeedRanges[i+1]
		for seed := seedStart; seed <= seedStart+seedRangeLength; seed++ {
			tmpSeed := seed
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
	for i := 3; i < len(lines); i += 2 {
		var toNextLines []string
		for i < len(lines) && lines[i] != "" {
			toNextLines = append(toNextLines, lines[i])
			i += 1
		}
		var toNext []mapping
		for _, line := range toNextLines {
			toNext = append(toNext, parseRange(line))
		}
		sort.Slice(toNext, func(j int, k int) bool {
			return toNext[j].source < toNext[k].source
		})
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

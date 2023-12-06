package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	// file, err := os.Open("./Day6/Part1/input/testInput.txt")
	file, err := os.Open("./Day6/Part1/input/InputFile.txt")
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

	raceDurations := parseLine(fileLines[0])
	raceDistances := parseLine(fileLines[1])

	possibleDistances := getPossibleDistances(raceDistances, raceDurations)
	totalProduct := 1
	for _, distance := range possibleDistances {
		totalProduct *= distance
	}
	fmt.Printf("Total product is %v\n", totalProduct)
	elapsed := time.Since(start)
	fmt.Printf("Total runtime: %s\n", elapsed)
}

func parseLine(line string) []int {
	firstArray := strings.Split(line[9:], " ")
	filteredInts := []int{}
	for _, word := range firstArray {
		if word != "" {
			digit, err := strconv.Atoi(word)
			if err != nil {
				log.Fatal(err)
			}
			filteredInts = append(filteredInts, digit)
		}
	}
	return filteredInts
}

func getPossibleDistances(raceDistances []int, raceDurations []int) []int {
	possibleDistances := []int{}
	for i := 0; i < len(raceDistances); i++ {
		time := raceDurations[i]
		distance := raceDistances[i]
		lowerRoot := ((float64(time)) - math.Sqrt(math.Pow(float64(time), 2)-4*float64(distance))) / (2)
		upperRoot := ((float64(time)) + math.Sqrt(math.Pow(float64(time), 2)-4*float64(distance))) / (2)
		// If upperRoot is an int (unlikely) then the amount of integers between the two roots is different
		if upperRoot != float64(int(upperRoot)) {
			possibleDistances = append(possibleDistances, int(upperRoot)-int(lowerRoot))
		} else {
			possibleDistances = append(possibleDistances, int(upperRoot)-int(lowerRoot)-1)
		}

	}
	return possibleDistances
}

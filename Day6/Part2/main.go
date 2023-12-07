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

	fileLines := ReadFileIntoStringSlice("./Day6/Part1/input/InputFile.txt")
	// fileLines := ReadFileIntoStringSlice("./Day6/Part1/input/testInput.txt")

	raceDurations := parseLine(fileLines[0])
	raceDistances := parseLine(fileLines[1])

	amountOfSolutions := getPossibleDistances(raceDistances, raceDurations)

	fmt.Printf("Amount of solutions is %v\n", amountOfSolutions)

	elapsed := time.Since(start)
	fmt.Printf("Total runtime: %s\n", elapsed)
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

func parseLine(line string) int {
	combinedDigits := strings.ReplaceAll(line[9:], " ", "")
	singleInt, err := strconv.Atoi(combinedDigits)
	if err != nil {
		log.Fatal(err)
	}
	return singleInt
}

func getPossibleDistances(distance int, raceDuration int) int {
	lowerRoot := ((float64(raceDuration)) - math.Sqrt(math.Pow(float64(raceDuration), 2)-4*float64(distance))) / (2)
	upperRoot := ((float64(raceDuration)) + math.Sqrt(math.Pow(float64(raceDuration), 2)-4*float64(distance))) / (2)
	// If upperRoot is an int (unlikely) then the amount of integers between the two roots is different
	if upperRoot != float64(int(upperRoot)) {
		return int(upperRoot) - int(lowerRoot)
	} else {
		return int(upperRoot) - int(lowerRoot) - 1
	}
}

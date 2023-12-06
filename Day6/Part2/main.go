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
	amountOfSolutions := getPossibleDistances(raceDistances, raceDurations)

	fmt.Printf("Amount of solutions is %v\n", amountOfSolutions)

	elapsed := time.Since(start)
	fmt.Printf("Total runtime: %s\n", elapsed)
}

func parseLine(line string) int {
	firstArray := strings.Split(line[9:], " ")
	combinedDigits := ""
	for _, word := range firstArray {
		if word != "" {
			combinedDigits += word
		}
	}
	singleInt, err := strconv.Atoi(combinedDigits)
	if err != nil {
		log.Fatal(err)
	}
	return singleInt
}

func getPossibleDistances(distance int, time int) int {
	fmt.Printf("Time is: %d\n", time)
	fmt.Printf("DIstance is: %d\n", distance)
	lowerRoot := ((float64(time)) - math.Sqrt(math.Pow(float64(time), 2)-4*float64(distance))) / (2)
	upperRoot := ((float64(time)) + math.Sqrt(math.Pow(float64(time), 2)-4*float64(distance))) / (2)
	fmt.Printf("Lower bound :%v\n", lowerRoot)
	fmt.Printf("Upper bound :%v\n", upperRoot)
	amountOfSolutions := 0
	for j := int(lowerRoot) + 1; float64(j) < upperRoot; j++ {
		amountOfSolutions += 1
	}
	return amountOfSolutions
}

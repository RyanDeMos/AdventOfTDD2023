package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type LeftRightDirections struct {
	L string
	R string
}

func main() {
	fileLines := ReadFileIntoStringSlice("./Day8/Part1/input/InputFile.txt")

	// Parse first row of input file into Directions
	Directions := fileLines[0]

	// Parse input file into Map
	Map := parseMap(fileLines[2:]) //map[string]LeftRightDirections{}

	fmt.Printf("Directions: %v\n", Directions)
	fmt.Printf("Map: %v\n", Map)

	currentLocation := "AAA"

	stepsTaken := travelToZZZ(currentLocation, Directions, Map)
	fmt.Printf("Total steps taken: %d\n", stepsTaken)
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

func parseMap(filesLines []string) map[string]LeftRightDirections {
	Map := map[string]LeftRightDirections{}
	for _, line := range filesLines {
		line = strings.ReplaceAll(line, " ", "")
		keyAndValue := strings.Split(line, "=")
		key := keyAndValue[0]
		values := strings.Split(keyAndValue[1][1:8], ",")
		leftAndRightDirections := LeftRightDirections{values[0], values[1]}
		Map[key] = leftAndRightDirections
	}
	return Map
}

func travelToNext(currentPosition string, nextDirection string, Map map[string]LeftRightDirections) string {
	if nextDirection == "L" {
		return Map[currentPosition].L
	}
	return Map[currentPosition].R
}

func travelToZZZ(currentLocation string, Directions string, Map map[string]LeftRightDirections) int {
	stepsTaken := 0
getToZZZ:
	for currentLocation != "ZZZ" {
		for idx, direction := range Directions {
			currentLocation = travelToNext(currentLocation, string(direction), Map)
			fmt.Printf("Current Location: %v, direction travelled: %v\n", currentLocation, string(direction))
			if currentLocation == "ZZZ" {
				stepsTaken += idx + 1
				break getToZZZ
			}
		}
		stepsTaken += len(Directions)
	}
	return stepsTaken
}

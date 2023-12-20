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
	fileLines := ReadFileIntoStringSlice("./Day8/Part2/input/InputFile.txt")
	// fileLines := ReadFileIntoStringSlice("./Day8/Part2/input/testInput.txt")

	// Parse first row of input file into Directions
	Directions := fileLines[0]

	// Parse input file into Map
	Map := parseMap(fileLines[2:]) //map[string]LeftRightDirections{}

	// fmt.Printf("Directions: %v\n", Directions)
	// fmt.Printf("Map: %v\n", Map)

	startingLocations := getAllStartingNodes(fileLines)
	fmt.Printf("All starting locations: %v\n", startingLocations)

	stepsTaken := travelAllToZ(startingLocations, Directions, Map)
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

func travelAllToZ(currentLocations []string, Directions string, Map map[string]LeftRightDirections) int {
	stepsTakenArray := []int{}
getToZZZ:
	for _, node := range currentLocations {
		stepsTaken := 0
		for string(node[2]) != "Z" {
			for idx, direction := range Directions {
				node = travelToNext(node, string(direction), Map)
				if string(node[2]) == "Z" {
					stepsTaken += idx + 1
					stepsTakenArray = append(stepsTakenArray, stepsTaken)
					continue getToZZZ
				}
			}
			stepsTaken += len(Directions)
		}
	}
	lcm := LCM(stepsTakenArray[0], stepsTakenArray[1], stepsTakenArray[2:]...)
	return lcm
}

func getAllStartingNodes(lines []string) []string {
	startingLocations := []string{}
	for _, line := range lines[2:] {
		if string(line[2]) == "A" {
			startingLocations = append(startingLocations, line[:3])
		}
	}
	return startingLocations
}

// Euclidean Algo
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Get the LCM from the GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

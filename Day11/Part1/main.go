package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	fileLines := ReadFileIntoStringSlice("./Day11/Part1/input/testInput.txt")
	// fileLines := ReadFileIntoStringSlice("./Day11/Part1/input/InputFile.txt")
	expandedUniverse := expandUniverse(fileLines)
	fmt.Printf("Expanded Rows:\n")
	for _, line := range expandedUniverse {
		fmt.Printf("%v\n", line)
	}
	galaxyPositions := findAllGalaxyPositions(expandedUniverse)
	fmt.Printf("Galaxy Positions: %v\n", galaxyPositions)

	distances := findAllDistances(galaxyPositions)
	fmt.Printf("Distances: %v\n", distances)
	fmt.Printf("Amount of distance: %v\n", len(distances))

	totalDistance := findTotalDistance(distances)
	fmt.Printf("Total distance: %v\n", totalDistance)
}

func findTotalDistance(distances []int) int {
	totalDistance := 0
	for _, distance := range distances {
		totalDistance += distance
	}
	return totalDistance
}

func findAllDistances(galaxyPositions [][]int) []int {
	distances := []int{}
	for i := 0; i < len(galaxyPositions); i++ {
		for j := i + 1; j < len(galaxyPositions); j++ {
			distance := math.Abs(float64(galaxyPositions[i][0]-galaxyPositions[j][0])) + math.Abs(float64(galaxyPositions[i][1]-galaxyPositions[j][1]))
			distances = append(distances, int(distance))
		}
	}
	return distances
}

func findAllGalaxyPositions(expandedUniverse []string) [][]int {
	galaxyPositions := [][]int{}
	for row, line := range expandedUniverse {
		for column, character := range line {
			if string(character) == "#" {
				galaxyPositions = append(galaxyPositions, []int{column, row})
			}
		}
	}
	return galaxyPositions
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

func expandUniverse(fileLines []string) []string {
	expandedUniverse := expandRows(fileLines)
	expandedUniverse = expandColumns(expandedUniverse)
	return expandedUniverse
}

func expandRows(fileLines []string) []string {
	expandedUniverse := []string{}
expandUniverse:
	for _, line := range fileLines {
		expandedUniverse = append(expandedUniverse, line)
		// if line contains no # then
		for _, character := range line {
			if string(character) != "." {
				continue expandUniverse
			}
		}
		expandedUniverse = append(expandedUniverse, strings.Repeat(".", len(line)))
	}
	return expandedUniverse
}

func expandColumns(fileLines []string) []string {
	expandedUniverse := make([]string, len(fileLines))
	for column := range fileLines[0] {
		containsGalaxy := false
		for row := 0; row < len(fileLines); row++ {
			expandedUniverse[row] += string(fileLines[row][column])
			if string(fileLines[row][column]) != "." {
				containsGalaxy = true
			}
		}
		if !containsGalaxy {
			for row := 0; row < len(fileLines); row++ {
				expandedUniverse[row] += "."
			}
		}
	}
	return expandedUniverse
}

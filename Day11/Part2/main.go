package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	// fileLines := ReadFileIntoStringSlice("./Day11/Part2/input/testInput.txt")
	fileLines := ReadFileIntoStringSlice("./Day11/Part2/input/InputFile.txt")
	expandedRows, expandedColumns := expandUniverse(fileLines)
	// fmt.Printf("Expanded Rows: %v\n", expandedRows)
	// fmt.Printf("Expanded Columns: %v\n", expandedColumns)

	galaxyPositions := findAllGalaxyPositions(fileLines)
	// fmt.Printf("Galaxy Positions: %v\n", galaxyPositions)

	allDistances := findAllDistances(galaxyPositions, expandedRows, expandedColumns)
	// fmt.Printf("All distances: %v\n", allDistances)

	totalDistance := findTotalDistance(allDistances)
	fmt.Printf("Total Distance: %v\n", totalDistance)
}

func findTotalDistance(distances []int) int {
	totalDistance := 0
	for _, distance := range distances {
		totalDistance += distance
	}
	return totalDistance
}

func findAllDistances(galaxyPositions [][]int, expandedRows map[int]int, expandedColumns map[int]int) []int {
	distances := []int{}
	for i := 0; i < len(galaxyPositions); i++ {
		for j := i + 1; j < len(galaxyPositions); j++ {
			distance := 0
			maxX := int(math.Max(float64(galaxyPositions[i][0]), float64(galaxyPositions[j][0])))
			minX := int(math.Min(float64(galaxyPositions[i][0]), float64(galaxyPositions[j][0])))
			maxY := int(math.Max(float64(galaxyPositions[i][1]), float64(galaxyPositions[j][1])))
			minY := int(math.Min(float64(galaxyPositions[i][1]), float64(galaxyPositions[j][1])))
			for x := maxX - 1; x >= minX; x-- {
				distance += expandedColumns[x]
			}
			for y := maxY - 1; y >= minY; y-- {
				distance += expandedRows[y]
			}
			distances = append(distances, distance)
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

func expandUniverse(fileLines []string) (map[int]int, map[int]int) {
	rowValues := expandRows(fileLines)
	columnValues := expandColumns(fileLines)
	return rowValues, columnValues
}

func expandRows(fileLines []string) map[int]int {
	rowValues := map[int]int{}
	for row, line := range fileLines {
		// if line contains no # then
		containsNoGalaxy := true
		for _, character := range line {
			if string(character) != "." {
				containsNoGalaxy = false
			}
		}
		if containsNoGalaxy {
			rowValues[row] = 1000000
			// rowValues[row] = 2
		} else {
			rowValues[row] = 1
		}
	}
	return rowValues
}

func expandColumns(fileLines []string) map[int]int {
	expandedColumns := map[int]int{}
	for column := range fileLines[0] {
		containsNoGalaxy := true
		for row := 0; row < len(fileLines); row++ {
			if string(fileLines[row][column]) != "." {
				containsNoGalaxy = false
			}
		}
		if containsNoGalaxy {
			expandedColumns[column] = 1000000
			// expandedColumns[column] = 2
		} else {
			expandedColumns[column] = 1
		}
	}
	return expandedColumns
}

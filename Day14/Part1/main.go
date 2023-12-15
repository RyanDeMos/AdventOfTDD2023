package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// fileLines := ReadFileIntoStringSlice("./Day14/Part1/input/testInput.txt")
	fileLines := ReadFileIntoStringSlice("./Day14/Part1/input/InputFile.txt")
	Part1(fileLines)
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

func Part1(fileLines []string) {
	grid := [][]string{}
	for _, line := range fileLines {
		grid = append(grid, strings.Split(line, ""))
	}
	moveAllStonesNorth(grid)
	totalLoad := getLoad(grid)
	fmt.Printf("Total Load: %d\n", totalLoad)
}

func getLoad(grid [][]string) int {
	totalLoad := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if string(grid[row][column]) == "O" {
				totalLoad += len(grid) - row
			}
		}
	}
	return totalLoad
}

func moveAllStonesNorth(grid [][]string) {
	for row := 1; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if string(grid[row][column]) == "O" {
				moveStoneNorth(grid, row, column, true)
			}
		}
	}
}

func moveStoneNorth(grid [][]string, row int, column int, firstCall bool) {
	if firstCall {
		grid[row][column] = "."
	}
	if row != 0 {
		if grid[row-1][column] == "." {
			moveStoneNorth(grid, row-1, column, false)
		} else {
			grid[row][column] = "O"
		}
	} else {
		grid[row][column] = "O"
	}
}

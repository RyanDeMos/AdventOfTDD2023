package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileLines := ReadFileIntoStringSlice("./Day14/Part2/input/testInput.txt")
	// fileLines := ReadFileIntoStringSlice("./Day14/Part2/input/InputFile.txt")
	Part2(fileLines)
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

func Part2(fileLines []string) {
	grid := [][]string{}
	for _, line := range fileLines {
		grid = append(grid, strings.Split(line, ""))
	}

	possibleSolutions := map[int]bool{}
	seenGrids := map[string]bool{}

	cycle := 0
	for cycle < (300) {

		moveAllStonesNorth(grid)
		moveAllStonesWest(grid)
		moveAllStonesSouth(grid)
		moveAllStonesEast(grid)
		allLines := ""
		for _, line := range grid {
			for _, character := range line {
				allLines += string(character)
			}
		}

		// seenGrids[allLines] = true
		// possibleSolutions[getLoad(grid)] = true
		cycle += 1
		fmt.Printf("Cycle: %v\n", cycle)
		if cycle > 94 {
			if seenGrids[allLines] {
				break
			}
			seenGrids[allLines] = true
			possibleSolutions[getLoad(grid)] = true
		}
	}
	totalLoad1 := getLoad(grid)
	fmt.Printf("Len Possible Solutions %v\n", len(possibleSolutions))
	fmt.Printf("Len seen Grids Solutions %v\n", len(seenGrids))
	fmt.Printf("Total Load1: %d\n", totalLoad1)

	// grid2 := [][]string{}
	// for _, line := range fileLines {
	// 	grid2 = append(grid2, strings.Split(line, ""))
	// }
	// for j := 0; j < (1000000000%len(possibleSolutions)); j++ {
	// 	moveAllStonesNorth(grid2)
	// 	moveAllStonesWest(grid2)
	// 	moveAllStonesSouth(grid2)
	// 	moveAllStonesEast(grid2)
	// }
	// totalLoad := getLoad(grid2)
	// fmt.Printf("Total Load: %d\n", totalLoad)
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

func moveStoneWest(grid [][]string, row int, column int, firstCall bool) {
	if firstCall {
		grid[row][column] = "."
	}
	if column != 0 {
		if grid[row][column-1] == "." {
			moveStoneWest(grid, row, column-1, false)
		} else {
			grid[row][column] = "O"
		}
	} else {
		grid[row][column] = "O"
	}
}

func moveAllStonesWest(grid [][]string) {
	for row := 0; row < len(grid); row++ {
		for column := 1; column < len(grid[row]); column++ {
			if string(grid[row][column]) == "O" {
				moveStoneWest(grid, row, column, true)
			}
		}
	}
}

func moveAllStonesSouth(grid [][]string) {
	for row := len(grid) - 1; row >= 0; row-- {
		for column := 0; column < len(grid[row]); column++ {
			if string(grid[row][column]) == "O" {
				moveStoneSouth(grid, row, column, true)
			}
		}
	}
}

func moveStoneSouth(grid [][]string, row int, column int, firstCall bool) {
	if firstCall {
		grid[row][column] = "."
	}
	if row != len(grid)-1 {
		if grid[row+1][column] == "." {
			moveStoneSouth(grid, row+1, column, false)
		} else {
			grid[row][column] = "O"
		}
	} else {
		grid[row][column] = "O"
	}
}

func moveStoneEast(grid [][]string, row int, column int, firstCall bool) {
	if firstCall {
		grid[row][column] = "."
	}
	if column != len(grid[row])-1 {
		if grid[row][column+1] == "." {
			moveStoneEast(grid, row, column+1, false)
		} else {
			grid[row][column] = "O"
		}
	} else {
		grid[row][column] = "O"
	}
}

func moveAllStonesEast(grid [][]string) {
	for row := 0; row < len(grid); row++ {
		for column := len(grid[row]) - 1; column >= 0; column-- {
			if string(grid[row][column]) == "O" {
				moveStoneEast(grid, row, column, true)
			}
		}
	}
}

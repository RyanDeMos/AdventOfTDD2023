package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
)

type position struct {
	pipe                         string
	previous_direction_of_travel string
	xposition                    int
	yposition                    int
}

func main() {
	// fileLines := ReadFileIntoStringSlice("./Day10/Part2/input/testInput.txt")
	// fileLines := ReadFileIntoStringSlice("./Day10/Part2/input/testInput2.txt")
	fileLines := ReadFileIntoStringSlice("./Day10/Part2/input/testInput3.txt")
	// fileLines := ReadFileIntoStringSlice("./Day10/Part2/input/testInput4.txt")
	// fileLines := ReadFileIntoStringSlice("./Day10/Part2/input/InputFile.txt")

	loopBoundary := [][]int{}

	// Get the starting position of S and the pipes connecting to it
	startingX, startingY := getStartingLocation(fileLines)
	loopBoundary = append(loopBoundary, []int{startingX, startingY})

	// Loop over the two paths traversing them until they are at the same position
	stepsTaken := 1
	connectingToStart := getConnectingPipesToStart(startingX, startingY, fileLines)
	firstStream := connectingToStart[0]
	loopBoundary = append(loopBoundary, []int{firstStream.xposition, firstStream.yposition})
	for (firstStream.xposition != startingX) || (firstStream.yposition != startingY) {
		stepsTaken += 1
		firstStream = travelToNextPipe(firstStream, fileLines)
		loopBoundary = append(loopBoundary, []int{firstStream.xposition, firstStream.yposition})
	}
	fmt.Printf("Total distance in pipe is: %v\n", stepsTaken)

	connections := pipelineConnections(loopBoundary, fileLines)
	fmt.Printf("Connections: %v\n", connections)

	totalEnclosedArea := findEnclosedArea(loopBoundary, fileLines, connections)
	fmt.Printf("Total enclosed area: %v\n", totalEnclosedArea)

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

func getStartingLocation(fileLines []string) (int, int) {
	for y, line := range fileLines {
		for x, pipe := range line {
			if string(pipe) == "S" {
				return x, y
			}
		}
	}
	// Couldn't find S :/
	return -1, -1
}

func getConnectingPipesToStart(startingX int, startingY int, fileLines []string) []position {
	connectingCoords := []position{}
	// Check to the North
	if string(fileLines[startingY-1][startingX]) == "|" || string(fileLines[startingY-1][startingX]) == "7" || string(fileLines[startingY-1][startingX]) == "F" {
		connectingCoords = append(connectingCoords, position{
			pipe:                         string(fileLines[startingY-1][startingX]),
			previous_direction_of_travel: "north",
			xposition:                    startingX,
			yposition:                    startingY - 1,
		})
	}
	// Check to the South
	if string(fileLines[startingY+1][startingX]) == "|" || string(fileLines[startingY+1][startingX]) == "L" || string(fileLines[startingY+1][startingX]) == "J" {
		connectingCoords = append(connectingCoords, position{
			pipe:                         string(fileLines[startingY+1][startingX]),
			previous_direction_of_travel: "south",
			xposition:                    startingX,
			yposition:                    startingY + 1,
		})
		if len(connectingCoords) == 2 {
			return connectingCoords
		}
	}
	// Check to the East
	if string(fileLines[startingY][startingX+1]) == "-" || string(fileLines[startingY][startingX+1]) == "J" || string(fileLines[startingY][startingX+1]) == "7" {
		connectingCoords = append(connectingCoords, position{
			pipe:                         string(fileLines[startingY][startingX+1]),
			previous_direction_of_travel: "east",
			xposition:                    startingX + 1,
			yposition:                    startingY,
		})
		if len(connectingCoords) == 2 {
			return connectingCoords
		}
	}
	// Check to the West
	if string(fileLines[startingY][startingX-1]) == "-" || string(fileLines[startingY][startingX-1]) == "L" || string(fileLines[startingY][startingX-1]) == "F" {
		connectingCoords = append(connectingCoords, position{
			pipe:                         string(fileLines[startingY][startingX-1]),
			previous_direction_of_travel: "west",
			xposition:                    startingX - 1,
			yposition:                    startingY,
		})
		if len(connectingCoords) == 2 {
			return connectingCoords
		}
	}

	if len(connectingCoords) != 2 {
		panic("Should have 2 connecting but dont :(")
	}
	return connectingCoords
}

var pipeTypes = map[string]map[string]string{
	//   {"previous direction of travel": "next direction to travel", "previous direction of travel": "next direction to travel"},
	"|": {"south": "south", "north": "north"},
	"-": {"east": "east", "west": "west"},
	"L": {"south": "east", "west": "north"},
	"J": {"south": "west", "east": "north"},
	"7": {"east": "south", "north": "west"},
	"F": {"north": "east", "west": "south"},
}

var directionTravel = map[string][]int{
	"north": {0, -1},
	"south": {0, 1},
	"east":  {1, 0},
	"west":  {-1, 0},
}

func travelToNextPipe(currentPosition position, fileLines []string) position {
	nextPosition := position{}

	// Find the new direction of travel
	nextPosition.previous_direction_of_travel = pipeTypes[currentPosition.pipe][currentPosition.previous_direction_of_travel]

	// Find the new position
	coordChange := directionTravel[nextPosition.previous_direction_of_travel]
	nextPosition.xposition = currentPosition.xposition + coordChange[0]
	nextPosition.yposition = currentPosition.yposition + coordChange[1]

	// Find the new pipe type
	nextPosition.pipe = string(fileLines[nextPosition.yposition][nextPosition.xposition])
	return nextPosition
}

func findEnclosedArea(loopBoundary [][]int, fileLines []string, connections map[int][][]int) int {
	enclosedArea := 0
	fmt.Printf("Loop: %v\n", loopBoundary)
	for y, lines := range fileLines {
		pipesInThisLine := connections[y]
		if len(pipesInThisLine) == 0 {
			continue
		}
		crossedBoundaryCounter := 0
		for x := range lines {
			if !onLoop(loopBoundary, x, y) {
				for _, boundary := range pipesInThisLine {
					if x > boundary[0] {
						crossedBoundaryCounter += 1
					}
				}
			}
			if crossedBoundaryCounter%2 == 1 {
				enclosedArea += 1
			}
		}
	}
	return enclosedArea
}

func onLoop(loopBoundary [][]int, x int, y int) bool {
	for _, coorPair := range loopBoundary {
		if reflect.DeepEqual([]int{x, y}, coorPair) {
			return true
		}
	}
	return false
}

func pipelineConnections(loopBoundary [][]int, fileLines []string) map[int][][]int {
	pipeLineBoundariesAtYLevels := map[int][][]int{}
	// Loop over each line in the grid
	for y := 0; y < len(fileLines); y++ {
		startingPositionsOfBoundariesAtYLevel := [][]int{}
		// loop over each cooridnate in the pipeline
		for j := 0; j < len(loopBoundary)-1; j++ {
			// if the y position matches our line number
			if loopBoundary[j][1] == y {
				minXFound := loopBoundary[j][0]
				maxXFound := loopBoundary[j][0]

				//
			borderFollowingLevelY:
				for j < len(loopBoundary) {
					if loopBoundary[j][1] == y {
						if loopBoundary[j][0] < minXFound {
							minXFound = loopBoundary[j][0]
						}
						if loopBoundary[j][0] > maxXFound {
							maxXFound = loopBoundary[j][0]
						}

						j += 1
					} else {
						break borderFollowingLevelY
					}
				}
				startingPositionsOfBoundariesAtYLevel = append(startingPositionsOfBoundariesAtYLevel, []int{minXFound, maxXFound})
			}
		}
		sort.SliceStable(startingPositionsOfBoundariesAtYLevel, func(i, j int) bool {
			return startingPositionsOfBoundariesAtYLevel[i][0] < startingPositionsOfBoundariesAtYLevel[j][0]
		})
		pipeLineBoundariesAtYLevels[y] = startingPositionsOfBoundariesAtYLevel
	}
	return pipeLineBoundariesAtYLevels
}

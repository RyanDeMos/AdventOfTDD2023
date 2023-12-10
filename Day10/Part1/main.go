package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type position struct {
	pipe                         string
	previous_direction_of_travel string
	xposition                    int
	yposition                    int
}

func main() {
	// fileLines := ReadFileIntoStringSlice("./Day10/Part1/input/testInput.txt")
	// fileLines := ReadFileIntoStringSlice("./Day10/Part1/input/testInput2.txt")
	fileLines := ReadFileIntoStringSlice("./Day10/Part1/input/InputFile.txt")

	// Get the starting position of S and the pipes connecting to it
	startingX, startingY := getStartingLocation(fileLines)
	connectingToStart := getConnectingPipesToStart(startingX, startingY, fileLines)

	// Loop over the two paths traversing them until they are at the same position
	stepsTaken := 1
	firstStream := connectingToStart[0]
	secondStream := connectingToStart[1]
	for (firstStream.xposition != secondStream.xposition) || (firstStream.yposition != secondStream.yposition) {
		stepsTaken += 1
		firstStream = travelToNextPipe(firstStream, fileLines)
		secondStream = travelToNextPipe(secondStream, fileLines)
	}
	fmt.Printf("Total distance in pipe is: %v\n", stepsTaken)
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

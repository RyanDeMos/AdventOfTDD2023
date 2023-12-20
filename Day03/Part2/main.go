package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	threeLines := map[string]string{
		"line1": "............................................................................................................................................",
		"line2": "............................................................................................................................................",
		"line3": "............................................................................................................................................",
	}
	// "............................................................................................................................................",
	//"..........",

	file, err := os.Open("./Day3/Part1/input/InputFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%v\n", threeLines)
	digitSum := 0
	lineCounter := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%d\n", lineCounter)
		lineCounter += 1
		updateLines(scanner.Text(), threeLines)
		starLocations := parseSecondLine(threeLines["line2"])
		for _, location := range starLocations {
			digitSum += checkSurrondingsForDigits(threeLines, location)
		}
	}
	fmt.Printf("Total sum is %d", digitSum)
}

func updateLines(newLine string, threeLines map[string]string) {
	threeLines["line1"] = threeLines["line2"]
	threeLines["line2"] = threeLines["line3"]
	threeLines["line3"] = newLine
}

func parseSecondLine(line string) []int {
	starsInLine := []int{}
	for idx, character := range line {
		if string(character) == "*" {
			starsInLine = append(starsInLine, idx)
		}
	}
	return starsInLine
}

func checkSurrondingsForDigits(threeLines map[string]string, starLocation int) int {
	digitCounter := 0
	gearRatio := 1
	if starLocation == 0 {
		// Star on left
		// Above
		if unicode.IsDigit(rune(threeLines["line1"][starLocation])) || unicode.IsDigit(rune(threeLines["line1"][starLocation+1])) {
			// Start at either 0th or 1st position
			startIdx := 0
			if unicode.IsDigit(rune(threeLines["line1"][starLocation+1])) && !unicode.IsDigit(rune(threeLines["line1"][starLocation])) {
				startIdx = 1
			}
			digitCounter += 1
			endIdx := 0
			for idx, char := range threeLines["line1"][startIdx:] {
				if !unicode.IsDigit(char) {
					endIdx = idx
					break
				}
			}
			aboveDigit, err := strconv.Atoi(threeLines["line1"][startIdx:endIdx])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= aboveDigit
		}
		// Below
		if unicode.IsDigit(rune(threeLines["line3"][starLocation])) || unicode.IsDigit(rune(threeLines["line3"][starLocation+1])) {
			// Start at either 0th or 1st position
			startIdx := 0
			if unicode.IsDigit(rune(threeLines["line3"][starLocation+1])) && !unicode.IsDigit(rune(threeLines["line3"][starLocation])) {
				startIdx = 1
			}
			digitCounter += 1
			endIdx := 0
			for idx, char := range threeLines["line3"][startIdx:] {
				if !unicode.IsDigit(char) {
					endIdx = idx
					break
				}
			}
			aboveDigit, err := strconv.Atoi(threeLines["line3"][startIdx:endIdx])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= aboveDigit
		}
		// Right
		if unicode.IsDigit(rune(threeLines["line2"][starLocation+1])) {
			startIdx := 1
			digitCounter += 1
			endIdx := 0
			for idx, char := range threeLines["line2"][startIdx:] {
				if !unicode.IsDigit(char) {
					endIdx = idx
					break
				}
			}
			aboveDigit, err := strconv.Atoi(threeLines["line2"][startIdx : endIdx+1])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= aboveDigit
		}
	} else if starLocation == len(threeLines["line2"])-1 {
		// Star on right
		// Above
		if unicode.IsDigit(rune(threeLines["line1"][starLocation])) || unicode.IsDigit(rune(threeLines["line1"][starLocation-1])) {
			digitCounter += 1

			// Start at either 0th or 1st position
			endIdx := starLocation
			if unicode.IsDigit(rune(threeLines["line1"][starLocation-1])) && !unicode.IsDigit(rune(threeLines["line1"][starLocation])) {
				endIdx = starLocation - 1
			}

			startIdx := 0
			for idx := endIdx; idx >= 0; idx-- {
				if !unicode.IsDigit(rune(threeLines["line1"][idx])) {
					startIdx = idx
					break
				}
			}

			aboveDigit, err := strconv.Atoi(threeLines["line1"][startIdx+1:])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= aboveDigit
		}
		// Below
		if unicode.IsDigit(rune(threeLines["line3"][starLocation])) || unicode.IsDigit(rune(threeLines["line3"][starLocation-1])) {
			digitCounter += 1

			// Start at either 0th or 1st position
			endIdx := starLocation
			if unicode.IsDigit(rune(threeLines["line3"][starLocation-1])) && !unicode.IsDigit(rune(threeLines["line3"][starLocation])) {
				endIdx = starLocation - 1
			}

			startIdx := 0
			for idx := endIdx; idx >= 0; idx-- {
				if !unicode.IsDigit(rune(threeLines["line3"][idx])) {
					startIdx = idx
					break
				}
			}

			belowDigit, err := strconv.Atoi(threeLines["line3"][startIdx+1:])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= belowDigit
		}
		// Left
		if unicode.IsDigit(rune(threeLines["line2"][starLocation-1])) {
			endIdx := starLocation - 1
			digitCounter += 1

			startIdx := 0
			for idx := endIdx; idx >= 0; idx-- {
				if !unicode.IsDigit(rune(threeLines["line2"][idx])) {
					startIdx = idx
					break
				}
			}

			aboveDigit, err := strconv.Atoi(threeLines["line2"][startIdx+1 : endIdx+1])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= aboveDigit
		}
	} else {
		// Star in middle
		// Above
		if unicode.IsDigit(rune(threeLines["line1"][starLocation-1])) && unicode.IsDigit(rune(threeLines["line1"][starLocation+1])) && !unicode.IsDigit(rune(threeLines["line1"][starLocation])) {
			// Left and Right above but not middle
			digitCounter += 2

			// Get Left Digit
			leftEndIdx := starLocation
			leftStartidx := starLocation
			for idx := leftEndIdx - 1; idx >= 0; idx-- {
				if !unicode.IsDigit(rune(threeLines["line1"][idx])) {
					leftStartidx = idx + 1
					break
				}
				if idx == 0 {
					leftStartidx = 0
				}
			}

			if leftStartidx != 0 {
				leftDigit, err := strconv.Atoi(threeLines["line1"][leftStartidx:leftEndIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= leftDigit
			} else {
				leftDigit, err := strconv.Atoi(threeLines["line1"][leftStartidx:leftEndIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= leftDigit
			}
			// Get Right Digit
			rightEndIdx := starLocation
			rightStartidx := starLocation + 1
			for idx := rightStartidx; idx < len(threeLines["line1"]); idx++ {
				if !unicode.IsDigit(rune(threeLines["line1"][idx])) {
					rightEndIdx = idx
					break
				}
				if idx == len(threeLines["line1"])-1 {
					rightEndIdx = idx + 1
				}
			}
			rightDigit, err := strconv.Atoi(threeLines["line1"][rightStartidx:rightEndIdx])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= rightDigit
		} else if unicode.IsDigit(rune(threeLines["line1"][starLocation-1])) || unicode.IsDigit(rune(threeLines["line1"][starLocation])) || unicode.IsDigit(rune(threeLines["line1"][starLocation+1])) {
			// Any of the above but the are not disconnected
			digitCounter += 1

			// Find left most digit
			startIdx := 0
			if unicode.IsDigit(rune(threeLines["line1"][starLocation-1])) {
				startIdx = starLocation - 1
			} else if unicode.IsDigit(rune(threeLines["line1"][starLocation])) {
				startIdx = starLocation
			} else {
				startIdx = starLocation + 1
			}
			for idx := startIdx; idx >= 0; idx-- {
				if !unicode.IsDigit(rune(threeLines["line1"][idx])) {
					startIdx = idx + 1
					break
				}
				if idx == 0 {
					startIdx = 0
				}
			}

			//Find right most digit
			endIdx := startIdx + 1
			for idx := endIdx; idx < len(threeLines["line1"]); idx++ {
				if !unicode.IsDigit(rune(threeLines["line1"][idx])) {
					endIdx = idx
					break
				}
				if idx == len(threeLines["line1"])-1 {
					endIdx = idx + 1
				}
			}
			if startIdx != 0 {
				aboveDigit, err := strconv.Atoi(threeLines["line1"][startIdx:endIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= aboveDigit
			} else {
				aboveDigit, err := strconv.Atoi(threeLines["line1"][startIdx:endIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= aboveDigit
			}
		}

		// Below
		// Below
		if unicode.IsDigit(rune(threeLines["line3"][starLocation-1])) && unicode.IsDigit(rune(threeLines["line3"][starLocation+1])) && !unicode.IsDigit(rune(threeLines["line3"][starLocation])) {
			// Left and Right above but not middle
			digitCounter += 2

			// Get Left Digit
			leftEndIdx := starLocation
			leftStartidx := starLocation
			for idx := leftEndIdx - 1; idx >= 0; idx-- {
				if !unicode.IsDigit(rune(threeLines["line3"][idx])) {
					leftStartidx = idx
					break
				}
				if idx == 0 {
					leftStartidx = 0
				}
			}

			if leftStartidx != 0 {
				leftDigit, err := strconv.Atoi(threeLines["line3"][leftStartidx+1 : leftEndIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= leftDigit
			} else {
				leftDigit, err := strconv.Atoi(threeLines["line3"][leftStartidx:leftEndIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= leftDigit
			}
			// Get Right Digit
			rightEndIdx := starLocation
			rightStartidx := starLocation + 1
			for idx := rightStartidx; idx < len(threeLines["line3"]); idx++ {
				if !unicode.IsDigit(rune(threeLines["line3"][idx])) {
					rightEndIdx = idx
					break
				}
				if idx == len(threeLines["line3"])-1 {
					rightEndIdx = idx + 1
				}
			}
			rightDigit, err := strconv.Atoi(threeLines["line3"][rightStartidx:rightEndIdx])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= rightDigit
		} else if unicode.IsDigit(rune(threeLines["line3"][starLocation-1])) || unicode.IsDigit(rune(threeLines["line3"][starLocation])) || unicode.IsDigit(rune(threeLines["line3"][starLocation+1])) {
			// Any of the above but the are not disconnected
			digitCounter += 1

			// Find left most digit
			startIdx := 0
			if unicode.IsDigit(rune(threeLines["line3"][starLocation-1])) {
				startIdx = starLocation - 1
			} else if unicode.IsDigit(rune(threeLines["line3"][starLocation])) {
				startIdx = starLocation
			} else {
				startIdx = starLocation + 1
			}
			for idx := startIdx; idx >= 0; idx-- {
				if !unicode.IsDigit(rune(threeLines["line3"][idx])) {
					startIdx = idx + 1
					break
				}
				if idx == 0 {
					startIdx = 0
				}
			}

			//Find right most digit
			endIdx := startIdx + 1
			for idx := endIdx; idx < len(threeLines["line3"]); idx++ {
				if !unicode.IsDigit(rune(threeLines["line3"][idx])) {
					endIdx = idx
					break
				}
				if idx == len(threeLines["line3"])-1 {
					endIdx = idx + 1
				}
			}
			if startIdx != 0 {
				aboveDigit, err := strconv.Atoi(threeLines["line3"][startIdx:endIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= aboveDigit
			} else {
				aboveDigit, err := strconv.Atoi(threeLines["line3"][startIdx:endIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= aboveDigit
			}
		}

		// Left same line
		if unicode.IsDigit(rune(threeLines["line2"][starLocation-1])) {
			// Left and Right above but not middle
			digitCounter += 1

			// Get Left Digit
			leftEndIdx := starLocation
			leftStartidx := starLocation
			for idx := leftEndIdx - 1; idx >= 0; idx-- {
				if !unicode.IsDigit(rune(threeLines["line2"][idx])) {
					leftStartidx = idx
					break
				}
				if idx == 0 {
					leftStartidx = 0
				}
			}

			if leftStartidx != 0 {
				leftDigit, err := strconv.Atoi(threeLines["line2"][leftStartidx+1 : leftEndIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= leftDigit
			} else {
				leftDigit, err := strconv.Atoi(threeLines["line2"][leftStartidx:leftEndIdx])
				if err != nil {
					log.Fatal(err)
				}
				gearRatio *= leftDigit
			}

		}
		// Right same line
		if unicode.IsDigit(rune(threeLines["line2"][starLocation+1])) {
			// Left and Right above but not middle
			digitCounter += 1

			// Get Right Digit
			rightEndIdx := starLocation
			rightStartidx := starLocation + 1
			for idx := rightStartidx; idx < len(threeLines["line2"]); idx++ {
				if !unicode.IsDigit(rune(threeLines["line2"][idx])) {
					rightEndIdx = idx
					break
				}
				if idx == len(threeLines["line2"])-1 {
					rightEndIdx = idx + 1
				}
			}
			rightDigit, err := strconv.Atoi(threeLines["line2"][rightStartidx:rightEndIdx])
			if err != nil {
				log.Fatal(err)
			}
			gearRatio *= rightDigit
		}
	}

	if digitCounter >= 2 {
		return gearRatio
	}
	return 0
}

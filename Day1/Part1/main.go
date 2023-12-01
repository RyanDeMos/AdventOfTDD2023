package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	file, err := os.Open("./Day1/Part1/input/InputFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total_calibration_sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		firstDigit := loopForwards(scanner.Text())
		secondDigit := loopBackwards(scanner.Text())
		total_calibration_sum += combineDigits(firstDigit, secondDigit)
	}
	fmt.Printf("Total result is %d", total_calibration_sum)
}

func loopForwards(line string) string {
	for _, character := range line {
		if unicode.IsDigit(character) {
			return string(character)
		}
	}
	return ""
}

func loopBackwards(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		character := line[i]
		r, _ := utf8.DecodeRuneInString(string(character))
		if unicode.IsDigit(r) {
			return string(character)
		}
	}
	return ""
}

func combineDigits(firstDigit string, secondDigit string) int {
	combinedString := firstDigit + secondDigit
	combinedInt, err := strconv.Atoi(combinedString)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return combinedInt
}

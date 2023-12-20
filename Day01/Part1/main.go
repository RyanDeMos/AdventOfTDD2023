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

	// Channels to store results of each line
	firstChan := make(chan string)
	secondChan := make(chan string)

	// Loop through each line in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go loopForwards(scanner.Text(), firstChan)
		go loopBackwards(scanner.Text(), secondChan)
		firstDigit := <-firstChan
		secondDigit := <-secondChan
		total_calibration_sum += combineDigits(firstDigit, secondDigit)
	}
	fmt.Printf("Total result is %d", total_calibration_sum)
}

func loopForwards(line string, ch chan string) {
	for _, character := range line {
		if unicode.IsDigit(character) {
			ch <- string(character)
			return
		}
	}
	return
}

func loopBackwards(line string, ch chan string) {
	for i := len(line) - 1; i >= 0; i-- {
		character := line[i]
		r, _ := utf8.DecodeRuneInString(string(character))
		if unicode.IsDigit(r) {
			ch <- string(character)
			return
		}
	}
	return
}

func combineDigits(firstDigit string, secondDigit string) int {
	combinedString := firstDigit + secondDigit
	combinedInt, err := strconv.Atoi(combinedString)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return combinedInt
}

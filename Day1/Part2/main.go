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
	file, err := os.Open("./Day1/Part2/input/InputFile.txt")
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
	for i := 0; i < len(line); i++ {
		character := line[i]
		r, _ := utf8.DecodeRuneInString(string(character))
		if unicode.IsDigit(r) {
			ch <- string(character)
			return
		} else {
			numberFromWord := getNumberFromWord(line, i)
			if numberFromWord != "" {
				ch <- numberFromWord
				return
			}
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
		} else {
			numberFromWord := getNumberFromWord(line, i)
			if numberFromWord != "" {
				ch <- numberFromWord
				return
			}
		}
	}
	return
}

func getNumberFromWord(line string, startPosition int) string {
	wordNumberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	// Max word length is 5 so endposition will be at most 4 more
	for endPosition := 4; endPosition >= 2; endPosition-- {
		if startPosition+endPosition < len(line) {
			if val, ok := wordNumberMap[line[startPosition:startPosition+endPosition+1]]; ok {
				return val
			}
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

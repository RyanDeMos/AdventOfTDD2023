package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	bid       int
	cards     string
	handValue int
}

var cardValueMap = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

func main() {
	filesLines := ReadFileIntoStringSlice("./Day7/Part1/input/InputFile.txt")
	// filesLines := ReadFileIntoStringSlice("./Day7/Part1/input/testInput.txt")
	allHands := []hand{}
	for _, line := range filesLines {
		allHands = append(allHands, parseLineIntoHand(line))
	}
	allHands = sortHands(allHands)
	fmt.Printf("All hands: %v\n", allHands)
	for _, hands := range allHands {
		fmt.Printf("Hand %v, value %d\n", hands, hands.handValue)
	}
	totalWinnings := getTotalWinnings(allHands)
	fmt.Printf("Total winnings: %d\n", totalWinnings)
}

func getTotalWinnings(hands []hand) int {
	winnings := 0
	for rank, hand := range hands {
		winnings += (rank + 1) * hand.bid
	}
	return winnings
}

func sortHands(hands []hand) []hand {
	sort.SliceStable(hands, func(i, j int) bool {
		// If the two hands have different values then we can sort right away by value
		if hands[i].handValue != hands[j].handValue {
			return hands[i].handValue < hands[j].handValue
		}
		// If the two cards have the same value then we sort by the individual card value
		for cardIDx := 0; cardIDx < 5; cardIDx++ {
			hand1Card := string(hands[i].cards[cardIDx])
			hand2Card := string(hands[j].cards[cardIDx])
			hand1CardValue := cardValueMap[hand1Card]
			hand2CardValue := cardValueMap[hand2Card]
			if hand1CardValue != hand2CardValue {
				return hand1CardValue < hand2CardValue
			}
		}
		return false
	})
	return hands
}

func findHandValue(cards string) int {
	handCardMap := map[string]int{}
	for _, character := range cards {
		handCardMap[string(character)] += 1
	}
	maxValue := 0
	for _, val := range handCardMap {
		if val > maxValue {
			if val == 3 && maxValue == 2 { // Fullhouse
				return 4
			}
			maxValue = val
		} else if maxValue == 2 && val == 2 { // 2 pair
			return 2
		} else if maxValue == 3 && val == 2 { //Fullhouse
			return 4
		}
	}
	if maxValue == 5 { // 5 of a kind
		return 6
	} else if maxValue == 4 { // 4 of a kind
		return 5
	} else if maxValue == 3 { //3 of kind
		return 3
	} else if maxValue == 2 { //1 pair
		return 1
	}
	return 0 // high card
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

func parseLineIntoHand(line string) hand {
	splitLine := strings.Split(line, " ")
	cards := splitLine[0]
	bid, err := strconv.Atoi(splitLine[1])
	if err != nil {
		log.Fatal(err)
	}
	handValue := findHandValue(cards)
	return hand{bid, cards, handValue}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
	"J": -1,
}

func main() {
	filesLines := ReadFileIntoStringSlice("./Day7/Part2/input/InputFile.txt")
	// filesLines := ReadFileIntoStringSlice("./Day7/Part2/input/testInput.txt")
	allHands := []hand{}
	for idx, line := range filesLines {
		fmt.Printf("Parse line %v into a hand. Line: %s\n", idx, line)
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

	if handCardMap["J"] == 5 {
		return 6
	}

	cardCounts := []int{}
	for key, val := range handCardMap {
		if key != "J" {
			cardCounts = append(cardCounts, val)
		}
	}
	slices.Sort(cardCounts)

	// Always add more cards to the card type with the most
	cardCounts[len(cardCounts)-1] += handCardMap["J"]

	if len(cardCounts) == 0 {
		panic("WHAT THE FUCK")
	}
	if cardCounts[len(cardCounts)-1] == 5 { //five of a kind
		return 6
	} else if cardCounts[len(cardCounts)-1] == 4 { // four of a kind
		return 5
	} else if cardCounts[len(cardCounts)-1] == 3 && cardCounts[len(cardCounts)-2] >= 2 { //Fullhouse
		return 4
	} else if cardCounts[len(cardCounts)-1] == 3 { // Three of a kind
		return 3
	} else if cardCounts[len(cardCounts)-1] == 2 && cardCounts[len(cardCounts)-2] == 2 && cardCounts[len(cardCounts)-3] != 2 { // two pair
		return 2
	} else if cardCounts[len(cardCounts)-1] == 2 {
		return 1
	} else {
		return 0
	}
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

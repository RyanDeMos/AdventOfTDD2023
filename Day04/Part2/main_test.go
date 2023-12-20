package main

import (
	"reflect"
	"testing"
)

func Test_parseLine(t *testing.T) {
	tests := []struct {
		name           string
		line           string
		winningNumbers []string
		ourNumbers     []string
	}{
		{
			name:           "Test",
			line:           "Card   1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			winningNumbers: []string{"41", "48", "83", "86", "17"},
			ourNumbers:     []string{"83", "86", "", "6", "31", "17", "", "9", "48", "53"},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			winningNumberResult, ourNumberResult := parseLine(testcase.line)

			if !reflect.DeepEqual(winningNumberResult, testcase.winningNumbers) {
				t.Fatalf("Wanted %v, but got %v", testcase.winningNumbers, winningNumberResult)
			}

			if !reflect.DeepEqual(ourNumberResult, testcase.ourNumbers) {
				t.Fatalf("Wanted %s, but got %s", testcase.ourNumbers, ourNumberResult)
			}
		})
	}
}

func Test_getMatchCount(t *testing.T) {
	tests := []struct {
		name           string
		winningNumbers []string
		ourNumbers     []string
		want           int
	}{
		{
			name:           "Test",
			winningNumbers: []string{"41", "48", "83", "86", "17"},
			ourNumbers:     []string{"83", "86", "", "6", "31", "17", "", "9", "48", "53"},
			want:           4,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := getMatchCount(testCase.winningNumbers, testCase.ourNumbers)
			if result != testCase.want {
				t.Fatalf("Wanted: %v, but got %v", testCase.want, result)
			}
		})
	}
}

func Test_getCopies(t *testing.T) {
	tests := []struct {
		name           string
		cardDictionary map[int]int
		matchCount     int
		cardNumber     int
		want           map[int]int
	}{
		{
			name: "Test Copies",
			cardDictionary: map[int]int{
				1: 1,
				2: 1,
				3: 1,
				4: 1,
				5: 1,
			},
			matchCount: 3,
			cardNumber: 1,
			want: map[int]int{
				1: 1,
				2: 2,
				3: 2,
				4: 2,
				5: 1,
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			getCopies(testCase.cardDictionary, testCase.matchCount, testCase.cardNumber)
			if !reflect.DeepEqual(testCase.cardDictionary, testCase.want) {
				t.Fatalf("Wanted: %v, but got %v", testCase.want, testCase.cardDictionary)
			}
		})
	}
}

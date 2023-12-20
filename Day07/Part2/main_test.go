package main

import (
	"reflect"
	"testing"
)

//
//
//
//

func Test_parseLineIntoHand(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected hand
	}{
		{
			name: "Fifth Last (AKA First)",
			line: "32T3K 765",
			expected: hand{
				bid:       765,
				cards:     "32T3K",
				handValue: 1,
			},
		},
		{
			name: "Fourth Last",
			line: "T55J5 684",
			expected: hand{
				bid:       684,
				cards:     "T55J5",
				handValue: 3,
			},
		},
		{
			name: "Third Last",
			line: "KK677 28",
			expected: hand{
				bid:       28,
				cards:     "KK677",
				handValue: 2,
			},
		},
		{
			name: "Second Last",
			line: "KTJJT 220",
			expected: hand{
				bid:       220,
				cards:     "KTJJT",
				handValue: 2,
			},
		},
		{
			name: "Last",
			line: "QQQJA 483",
			expected: hand{
				bid:       483,
				cards:     "QQQJA",
				handValue: 3,
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := parseLineIntoHand(testcase.line)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Expected %v, but got %v", testcase.expected, result)
			}
		})
	}
}

func Test_ReadFileIntoStringSlice(t *testing.T) {
	tests := []struct {
		name       string
		pathToFile string
		expected   []string
	}{
		{
			name:       "testFile",
			pathToFile: "../../testFile.txt",
			expected:   []string{"Line 1", "Line 2", "Line 3", "Line 4", ""},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := ReadFileIntoStringSlice(testcase.pathToFile)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Expected %v, but got %v", testcase.expected, result)
			}
		})
	}
}

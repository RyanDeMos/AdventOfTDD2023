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

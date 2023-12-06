package main

import (
	"reflect"
	"testing"
)

func Test_parseLine(t *testing.T) {
	tests := []struct {
		name      string
		inputLine string
		expected  []int
	}{
		{
			name:      "First Line",
			inputLine: "Time:        45     97     72     95",
			expected:  []int{45, 97, 72, 95},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := parseLine(testcase.inputLine)
			if !reflect.DeepEqual(testcase.expected, result) {
				t.Fatalf("Expected: %v, but got %v", testcase.expected, result)
			}
		})
	}
}

func Test_getPossibleDistances(t *testing.T) {
	tests := []struct {
		name          string
		expected      []int
		raceDurations []int
		raceDistance  []int
	}{
		{
			name:          "Test Possible Distances",
			expected:      []int{28, 72, 27, 48},
			raceDurations: []int{45, 97, 72, 95},
			raceDistance:  []int{305, 1062, 1110, 1695},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := getPossibleDistances(testcase.raceDistance, testcase.raceDurations)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Got: %v, but expected %v", result, testcase.expected)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_parseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{
			name:     "Working test",
			input:    "Time:        45     97     72     95",
			expected: 45977295,
		},
		{
			name:     "Single number",
			input:    "Time:                1       ",
			expected: 1,
		},
		{
			name:     "Distance",
			input:    "Distance:   305   1062   1110   1695",
			expected: 305106211101695,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := parseLine(testcase.input)
			if result != testcase.expected {
				t.Fatalf("Got: %v, expected: %v", result, testcase.expected)
			}
		})
	}
}

func Test_getPossibleDistances(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		distance float64
		time     float64
	}{
		{
			name:     "test",
			expected: 29891250,
			distance: 305106211101695,
			time:     45977295,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := getPossibleDistances(testcase.distance, testcase.time)
			if result != testcase.expected {
				t.Fatalf("Got: %v, but expected: %v", result, testcase.expected)
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

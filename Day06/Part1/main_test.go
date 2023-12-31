package main

import (
	"reflect"
	"testing"
)

func Benchmark_main(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

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

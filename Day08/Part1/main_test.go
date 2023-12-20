package main

import (
	"reflect"
	"testing"
)

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

func Test_parseMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]LeftRightDirections
	}{
		{
			name:  "two line test",
			input: []string{"AAA = (BBB, CCC)", "BBB = (DDD, EEE)"},
			expected: map[string]LeftRightDirections{
				"AAA": {"BBB", "CCC"},
				"BBB": {"DDD", "EEE"},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := parseMap(testcase.input)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Expected %v, but got %v", testcase.expected, result)
			}
		})
	}
}

func Test_travelToNext(t *testing.T) {
	tests := []struct {
		name            string
		currentPosition string
		nextDirection   string
		Map             map[string]LeftRightDirections
		expected        string
	}{
		{
			name:            "two line test",
			currentPosition: "AAA",
			nextDirection:   "R",
			Map:             map[string]LeftRightDirections{"AAA": {"BBB", "CCC"}},
			expected:        "CCC",
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := travelToNext(testcase.currentPosition, testcase.nextDirection, testcase.Map)
			if result != testcase.expected {
				t.Fatalf("Expected %v, but got %v", testcase.expected, result)
			}
		})
	}
}

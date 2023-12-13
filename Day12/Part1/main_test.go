package main

import (
	"fmt"
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

func Test_parselines(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []stringCondition
	}{
		{
			name: "Test 1",
			input: []string{
				"???.### 1,1,3",
				".??..??...?##. 1,1,3",
				"?#?#?#?#?#?#?#? 1,3,1,6",
				"????.#...#... 4,1,1",
				"????.######..#####. 1,6,5",
				"?###???????? 3,2,1",
			},
			expected: []stringCondition{
				{"???.###", []int{1, 1, 3}},
				{".??..??...?##.", []int{1, 1, 3}},
				{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}},
				{"????.#...#...", []int{4, 1, 1}},
				{"????.######..#####.", []int{1, 6, 5}},
				{"?###????????", []int{3, 2, 1}},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := parselines(testcase.input)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Expected %v, but got %v", testcase.expected, result)
			}
		})
	}
}

func Test_recursionNightmare(t *testing.T) {
	tests := []struct {
		name     string
		input    stringCondition
		expected int
	}{
		{
			name: "Test 1",
			input: stringCondition{
				condition: "???.###",
				groups:    []int{1, 1, 3},
			},
			expected: 1,
		},
		{
			name: "Test 2",
			input: stringCondition{
				condition: ".??..??...?##.",
				groups:    []int{1, 1, 3},
			},
			expected: 4,
		},
		{
			name: "Test 3",
			input: stringCondition{
				condition: "?#?#?#?#?#?#?#?",
				groups:    []int{1, 3, 1, 6},
			},
			expected: 1,
		},
		{
			name: "Test 4",
			input: stringCondition{
				condition: "????.#...#...",
				groups:    []int{4, 1, 1},
			},
			expected: 1,
		},
		{
			name: "Test 5",
			input: stringCondition{
				condition: "?###????????",
				groups:    []int{3, 2, 1},
			},
			expected: 10,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			recursionNightmare(testcase.input.condition, testcase.input.groups)
			result := count
			fmt.Printf("result: %v\n", result)
			if result != testcase.expected {
				t.Fatalf("Expected %v, but got %v", testcase.expected, result)
			}
		})
	}
}

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

func Test_getPatterns(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name: "test",
			input: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
				"",
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
				"",
			},
			expected: [][]string{
				[]string{
					"#.##..##.",
					"..#.##.#.",
					"##......#",
					"##......#",
					"..#.##.#.",
					"..##..##.",
					"#.#.##.#.",
				},
				[]string{
					"#...##..#",
					"#....#..#",
					"..##..###",
					"#####.##.",
					"#####.##.",
					"..##..###",
					"#....#..#",
				},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := getPatterns(testcase.input)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Expected %v\n, but got %v\n", testcase.expected, result)
			}
		})
	}
}

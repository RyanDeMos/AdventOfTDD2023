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

func Test_getStartingLocation(t *testing.T) {
	tests := []struct {
		name      string
		fileLines []string
		expectedx int
		expectedy int
	}{
		{
			name: "testFile",
			fileLines: []string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
			expectedy: 2,
			expectedx: 0,
		},
		{
			name: "testFile2",
			fileLines: []string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
			expectedy: 1,
			expectedx: 1,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			resultx, resulty := getStartingLocation(testcase.fileLines)
			if resultx != testcase.expectedx || resulty != testcase.expectedy {
				t.Fatalf("Expected x position: %v, but got %v\n Expected y position %v, but got, %v", testcase.expectedx, resultx, testcase.expectedy, resulty)
			}
		})
	}
}

func Test_getConnectingPipesToStart(t *testing.T) {
	tests := []struct {
		name      string
		fileLines []string
		startingX int
		startingY int
		expected  []position
	}{
		{
			name: "east south",
			fileLines: []string{
				" ..F7.",
				" .FJ|.",
				" SJ.L7",
				" |F--J",
				" LJ...",
			},
			startingX: 1,
			startingY: 2,
			expected: []position{
				{pipe: "|", previous_direction_of_travel: "south", xposition: 1, yposition: 3},
				{pipe: "J", previous_direction_of_travel: "east", xposition: 2, yposition: 2},
			},
		},
		{
			name: "another east south",
			fileLines: []string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
			startingX: 1,
			startingY: 1,
			expected: []position{
				{pipe: "|", previous_direction_of_travel: "south", xposition: 1, yposition: 2},
				{pipe: "-", previous_direction_of_travel: "east", xposition: 2, yposition: 1},
			},
		},
		{
			name: "north south",
			fileLines: []string{
				".|...",
				".S.7.",
				".|.|.",
				".L-J.",
				".....",
			},
			startingX: 1,
			startingY: 1,
			expected: []position{
				{pipe: "|", previous_direction_of_travel: "north", xposition: 1, yposition: 0},
				{pipe: "|", previous_direction_of_travel: "south", xposition: 1, yposition: 2},
			},
		},
		{
			name: "north and west",
			fileLines: []string{
				".F...",
				"-S.7.",
				"...|.",
				".L-J.",
				".....",
			},
			startingX: 1,
			startingY: 1,
			expected: []position{
				{pipe: "F", previous_direction_of_travel: "north", xposition: 1, yposition: 0},
				{pipe: "-", previous_direction_of_travel: "west", xposition: 0, yposition: 1},
			},
		},
		{
			name: "south and west",
			fileLines: []string{
				".....",
				"-S.7.",
				".L.|.",
				".L-J.",
				".....",
			},
			startingX: 1,
			startingY: 1,
			expected: []position{
				{pipe: "L", previous_direction_of_travel: "south", xposition: 1, yposition: 2},
				{pipe: "-", previous_direction_of_travel: "west", xposition: 0, yposition: 1},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := getConnectingPipesToStart(testcase.startingX, testcase.startingY, testcase.fileLines)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expected, result)
			}
		})
	}
}

func Test_travelToNextPipe(t *testing.T) {
	tests := []struct {
		name            string
		fileLines       []string
		currentPosition position
		expected        position
	}{
		{
			name: "Travel East",
			fileLines: []string{
				" ..L7.",
				" .FJ|.",
				" SJ.L7",
				" |F--J",
				" LJ...",
			},
			currentPosition: position{
				pipe:                         "F",
				previous_direction_of_travel: "north",
				xposition:                    2,
				yposition:                    1,
			},
			expected: position{
				pipe:                         "J",
				previous_direction_of_travel: "east",
				xposition:                    3,
				yposition:                    1,
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := travelToNextPipe(testcase.currentPosition, testcase.fileLines)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expected, result)
			}
		})
	}
}

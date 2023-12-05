package main

import (
	"reflect"
	"testing"
)

// eeds []int, destination int, delta int, rangeLength int
func Test_convertToNext(t *testing.T) {
	tests := []struct {
		name             string
		InputSeeds       []int
		InputDestination int
		InputDelta       int
		InputRangeLength int
		expected         []int
	}{
		{
			name:             "No changes",
			InputSeeds:       []int{79, 14, 55, 13},
			InputDestination: 50,
			InputDelta:       -48,
			InputRangeLength: 2,
			expected:         []int{79, 14, 55, 13},
		},
		{
			name:             "Rangelength 0 results in no change",
			InputSeeds:       []int{79, 98, 55, 99},
			InputDestination: 50,
			InputDelta:       -48,
			InputRangeLength: 0,
			expected:         []int{79, 98, 55, 99},
		},
		{
			name:             "Some Changes",
			InputSeeds:       []int{79, 14, 55, 13},
			InputDestination: 52,
			InputDelta:       2,
			InputRangeLength: 48,
			expected:         []int{81, 14, 57, 13},
		},
		{
			name:             "Change down",
			InputSeeds:       []int{79, 98, 55, 99},
			InputDestination: 50,
			InputDelta:       -48,
			InputRangeLength: 2,
			expected:         []int{79, 50, 55, 51},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := convertToNext(testcase.InputSeeds, testcase.InputDestination, testcase.InputDelta, testcase.InputRangeLength)
			if !reflect.DeepEqual(result, testcase.expected) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expected, result)
			}
		})
	}
}

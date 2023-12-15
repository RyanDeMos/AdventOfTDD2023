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

func Test_getCharacterValue(t *testing.T) {
	tests := []struct {
		name           string
		inputCharacter rune
		inputValue     int
		expected       int
	}{
		{
			name:           "H",
			inputCharacter: 'H',
			inputValue:     0,
			expected:       200,
		},
		{
			name:           "A",
			inputCharacter: 'A',
			inputValue:     200,
			expected:       153,
		},
		{
			name:           "S",
			inputCharacter: 'S',
			inputValue:     153,
			expected:       172,
		},
		{
			name:           "H2",
			inputCharacter: 'H',
			inputValue:     172,
			expected:       52,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := getCharacterValue(testcase.inputCharacter, testcase.inputValue)
			if result != testcase.expected {
				t.Fatalf("Expected %v, but got %v\n", testcase.expected, result)
			}
		})
	}
}

func Test_getSequenceValue(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "HASH",
			expected: 52,
		},
		{
			input:    "rn=1",
			expected: 30,
		},
		{
			input:    "cm-",
			expected: 253,
		},
		{
			input:    "qp=3",
			expected: 97,
		},
		{
			input:    "cm=2",
			expected: 47,
		},
		{
			input:    "qp-",
			expected: 14,
		},
		{
			input:    "pc=4",
			expected: 180,
		},
		{
			input:    "ot=9",
			expected: 9,
		},
		{
			input:    "ab=5",
			expected: 197,
		},
		{
			input:    "pc-",
			expected: 48,
		},
		{
			input:    "pc=6",
			expected: 214,
		},
		{
			input:    "ot=7",
			expected: 231,
		},
		{
			input:    "rn",
			expected: 0,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.input, func(t *testing.T) {
			result := getSequenceValue(testcase.input)
			if result != testcase.expected {
				t.Fatalf("Expected %v, but got %v\n", testcase.expected, result)
			}
		})
	}
}

func Test_addToBox(t *testing.T) {
	tests := []struct {
		name          string
		splitSequence []string
		boxes         map[int][][]string
		expected      map[int][][]string
	}{
		{
			name:          "Basic Add",
			splitSequence: []string{"rn", "1"},
			boxes:         map[int][][]string{},
			expected: map[int][][]string{
				0: {{"rn", "1"}},
			},
		},
		{
			name:          "Add to end",
			splitSequence: []string{"cm", "2"},
			boxes: map[int][][]string{
				0: {{"rn", "1"}},
			},
			expected: map[int][][]string{
				0: {{"rn", "1"}, {"cm", "2"}},
			},
		},
		{
			name:          "Basic add 2",
			splitSequence: []string{"qp", "3"},
			boxes:         map[int][][]string{},
			expected: map[int][][]string{
				1: {{"qp", "3"}},
			},
		},
		{
			name:          "Overwrite existing",
			splitSequence: []string{"ot", "7"},
			boxes: map[int][][]string{
				0: {{"rn", "1"}, {"cm", "2"}},
				1: {{"qp", "3"}},
				3: {{"ot", "9"}},
			},
			expected: map[int][][]string{
				0: {{"rn", "1"}, {"cm", "2"}},
				1: {{"qp", "3"}},
				3: {{"ot", "7"}},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			addToBox(testcase.splitSequence, testcase.boxes)
			if !reflect.DeepEqual(testcase.boxes, testcase.expected) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expected, testcase.boxes)
			}
		})
	}
}

func Test_removeFromBox(t *testing.T) {
	tests := []struct {
		name     string
		label    string
		boxes    map[int][][]string
		expected map[int][][]string
	}{
		{
			name:  "Basic Add",
			label: "rn",
			boxes: map[int][][]string{
				0: {{"rn", "1"}},
			},
			expected: map[int][][]string{0: {}},
		},
		{
			name:  "Add to end",
			label: "cm",
			boxes: map[int][][]string{
				0: {{"rn", "1"}, {"cm", "2"}},
			},
			expected: map[int][][]string{
				0: {{"rn", "1"}},
			},
		},
		{
			name:  "Basic add 2",
			label: "qp",
			boxes: map[int][][]string{
				1: {{"qp", "3"}},
			},
			expected: map[int][][]string{1: {}},
		},
		{
			name:  "Overwrite existing",
			label: "ot",
			boxes: map[int][][]string{
				0: {{"rn", "1"}, {"cm", "2"}},
				1: {{"qp", "3"}},
				3: {{"ot", "7"}},
			},
			expected: map[int][][]string{
				0: {{"rn", "1"}, {"cm", "2"}},
				1: {{"qp", "3"}},
				3: {},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			removeFromBox(testcase.label, testcase.boxes)
			if !reflect.DeepEqual(testcase.boxes, testcase.expected) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expected, testcase.boxes)
			}
		})
	}
}

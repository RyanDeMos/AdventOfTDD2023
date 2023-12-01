package main

import (
	"testing"
)

func Test_loopForwards(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Digit in first position",
			input:    "1abc2",
			expected: "1",
		},
		{
			name:     "Digit in last position",
			input:    "aabc2",
			expected: "2",
		},
		{
			name:     "Digits not at ends",
			input:    "pqr3stu8vwx",
			expected: "3",
		},
		{
			name:     "Multiple Digits",
			input:    "a1b2c3d4e5f",
			expected: "1",
		},
		{
			name:     "Only one digit",
			input:    "treb7uchet",
			expected: "7",
		},
		{
			name:     "one in front",
			input:    "onetreb7uchet",
			expected: "1",
		},
		{
			name:     "one after digit",
			input:    "treb7oneuchet",
			expected: "7",
		},
		{
			name:     "word at end of string",
			input:    "trebuchetthree",
			expected: "3",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := loopForwards(testcase.input)
			want := testcase.expected
			if got != want {
				t.Fatalf("Wanted %s, but got %s", want, got)
			}
		})
	}
}

func Test_loopBackwards(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Digit in last positions",
			input:    "abc2",
			expected: "2",
		},
		{
			name:     "Digit in first positions",
			input:    "1abc",
			expected: "1",
		},
		{
			name:     "Digits not at ends",
			input:    "pqr3stu8vwx",
			expected: "8",
		},
		{
			name:     "Multiple Digits",
			input:    "a1b2c3d4e5f",
			expected: "5",
		},
		{
			name:     "Only one digit",
			input:    "treb7uchet",
			expected: "7",
		},
		{
			name:     "one in front",
			input:    "onetreb7uchet",
			expected: "7",
		},
		{
			name:     "one after digit",
			input:    "treb7oneuchet",
			expected: "1",
		},
		{
			name:     "word at end of string",
			input:    "trebuchetthree",
			expected: "3",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := loopBackwards(testcase.input)
			want := testcase.expected
			if got != want {
				t.Fatalf("Wanted %s, but got %s", want, got)
			}
		})
	}
}

func Test_combineDigits(t *testing.T) {
	tests := []struct {
		name        string
		inputDigit1 string
		inputDigit2 string
		expected    int
	}{
		{
			name:        "First",
			inputDigit1: "1",
			inputDigit2: "2",
			expected:    12,
		},
		{
			name:        "Second",
			inputDigit1: "3",
			inputDigit2: "8",
			expected:    38,
		},
		{
			name:        "Third",
			inputDigit1: "1",
			inputDigit2: "5",
			expected:    15,
		},
		{
			name:        "Fourth",
			inputDigit1: "7",
			inputDigit2: "7",
			expected:    77,
		},
		{
			name:        "First digit is 0",
			inputDigit1: "0",
			inputDigit2: "7",
			expected:    7,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := combineDigits(testcase.inputDigit1, testcase.inputDigit2)
			want := testcase.expected
			if got != want {
				t.Fatalf("Wanted %d, but got %d", want, got)
			}
		})
	}
}

func Test_getNumberFromWord(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		position int
		expected string
	}{
		{
			name:     "one in front",
			line:     "onetreb7uchet",
			position: 0,
			expected: "1",
		},
		{
			name:     "one after digit",
			line:     "trebtoneuchet",
			position: 5,
			expected: "1",
		},
		{
			name:     "word at end of string",
			line:     "trebuchetthree",
			position: 9,
			expected: "3",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := getNumberFromWord(testcase.line, testcase.position)
			want := testcase.expected
			if got != want {
				t.Fatalf("Wanted %s, but got %s", want, got)
			}
		})
	}
}

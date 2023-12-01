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
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			ch := make(chan string)
			go loopForwards(testcase.input, ch)
			got := <-ch
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
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			ch := make(chan string)
			go loopBackwards(testcase.input, ch)
			got := <-ch
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

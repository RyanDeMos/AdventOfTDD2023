package main

import (
	"reflect"
	"testing"
)

func Test_getGameID(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "One digit game id",
			input: []string{"Game 1", " 3 blue", " 4 red", " 1 red", " 2 green", " 6 blue", " 2 green"},
			want:  1,
		},
		{
			name:  "Multiple digit game id",
			input: []string{"Game 121", " 3 blue", " 4 red", " 1 red", " 2 green", " 6 blue", " 2 green"},
			want:  121,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			want := testcase.want
			result := getGameID(testcase.input[0])
			if want != result {
				t.Fatalf("Expected %v, but got %v", want, result)
			}
		})
	}
}

func Test_getColourCount(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		colour string
		want   int
	}{
		{
			name:   "Green Scenario",
			input:  []string{"Game 1", " 3 blue", " 4 red", " 1 red", " 2 green", " 6 blue", " 2 green"},
			colour: "green",
			want:   2,
		},
		{
			name:   "Blue Scenario",
			input:  []string{"Game 1", " 3 blue", " 4 red", " 1 red", " 2 green", " 6 blue", " 2 green"},
			colour: "blue",
			want:   6,
		},
		{
			name:   "Red Scenario",
			input:  []string{"Game 1", " 3 blue", " 4 red", " 1 red", " 2 green", " 6 blue", " 2 green"},
			colour: "red",
			want:   4,
		},
		{
			name:   "None of given colour",
			input:  []string{"Game 1", " 3 blue", " 4 red", " 1 red", " 6 blue"},
			colour: "green",
			want:   0,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			want := testcase.want
			result := getMaxColourCount(testcase.input, testcase.colour)
			if want != result {
				t.Fatalf("Expected %v, but got %v", want, result)
			}
		})
	}
}

func Test_parseLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "Green Scenario",
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:  []string{"Game 1", " 3 blue", " 4 red", " 1 red", " 2 green", " 6 blue", " 2 green"},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			want := testcase.want
			result := parseLine(testcase.input)
			if !reflect.DeepEqual(want, result) {
				t.Fatalf("Expected %v, but got %v", want, result)
			}
		})
	}
}

func Test_Split(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		want  bool
	}{
		{
			name:  ";",
			input: ';',
			want:  true,
		},
		{
			name:  ":",
			input: ':',
			want:  true,
		},
		{
			name:  ",",
			input: ',',
			want:  true,
		},
		{
			name:  "a",
			input: 'a',
			want:  false,
		},
		{
			name:  "2",
			input: '2',
			want:  false,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			want := testcase.want
			result := Split(testcase.input)
			if want != result {
				t.Fatalf("Expected %v, but got %v", want, result)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_updateLines(t *testing.T) {
	tests := []struct {
		name     string
		newLine  string
		mapInput map[string]string
		want     map[string]string
	}{
		{
			name:    "Test",
			newLine: "Blah Blah Blah",
			mapInput: map[string]string{
				"line1": "First line",
				"line2": "Second line",
				"line3": "Third line",
			},
			want: map[string]string{
				"line1": "Second line",
				"line2": "Third line",
				"line3": "Blah Blah Blah",
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			want := testcase.want
			updateLines(testcase.newLine, testcase.mapInput)
			result := testcase.mapInput
			if !reflect.DeepEqual(result, want) {
				t.Fatalf("Expected %v, but got %v", want, result)
			}
		})
	}
}

func Test_parseSecondLine(t *testing.T) {
	tests := []struct {
		name string
		Line string
		want []digitAndLocations
	}{
		{
			name: "Test Middle digits",
			Line: "..35..633.",
			want: []digitAndLocations{
				{
					digit:      35,
					startIndex: 2,
					endIndex:   4,
				},
				{
					digit:      633,
					startIndex: 6,
					endIndex:   9,
				},
			},
		},
		{
			name: "Test no digits",
			Line: "...*......",
			want: []digitAndLocations{},
		},
		{
			name: "Test starting digits",
			Line: "467..114..",
			want: []digitAndLocations{
				{
					digit:      467,
					startIndex: 0,
					endIndex:   3,
				},
				{
					digit:      114,
					startIndex: 5,
					endIndex:   8,
				},
			},
		},
		{
			name: "Test ending digits",
			Line: "..23...675",
			want: []digitAndLocations{
				{
					digit:      23,
					startIndex: 2,
					endIndex:   4,
				},
				{
					digit:      675,
					startIndex: 7,
					endIndex:   10,
				},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			want := testcase.want
			result := parseSecondLine(testcase.Line)
			if !reflect.DeepEqual(result, want) {
				t.Fatalf("Expected %v, but got %v", want, result)
			}
		})
	}
}

func Test_checkSurrondingsForSpecialCharacters(t *testing.T) {
	tests := []struct {
		name             string
		lines            map[string]string
		digitInformation digitAndLocations
		want             bool
	}{
		{
			name: "Special Character Above",
			lines: map[string]string{
				"line1": "..*....%..",
				"line2": "..35..633.",
				"line3": "..........",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: true,
		},
		{
			name: "Special Character Above Diagonal Right",
			lines: map[string]string{
				"line1": "..*......%",
				"line2": "..35..633.",
				"line3": "..........",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: true,
		},
		{
			name: "Special Character Above Diagonal Left",
			lines: map[string]string{
				"line1": "..*..*....",
				"line2": "..35..633.",
				"line3": "..........",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: true,
		},
		{
			name: "Special Character Below",
			lines: map[string]string{
				"line1": "..*.......",
				"line2": "..35..633.",
				"line3": ".......*..",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: true,
		},
		{
			name: "Special Character Below Diagonal Right",
			lines: map[string]string{
				"line1": "..*.......",
				"line2": "..35..633.",
				"line3": ".........*",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: true,
		},
		{
			name: "Special Character Below Diagonal Left",
			lines: map[string]string{
				"line1": "..*.......",
				"line2": "..35..633.",
				"line3": ".....*....",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: true,
		},
		{
			name: "Special Character Left",
			lines: map[string]string{
				"line1": "..*.......",
				"line2": "..35.*633.",
				"line3": "..........",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: true,
		},
		{
			name: "Special Character Right",
			lines: map[string]string{
				"line1": "..*.......",
				"line2": "..35..633*",
				"line3": "..........",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: true,
		},
		{
			name: "No Special Character around given digit",
			lines: map[string]string{
				"line1": "..*.......",
				"line2": "..35..633.",
				"line3": "..........",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   9,
			},
			want: false,
		},
		{
			name: "Digits on end",
			lines: map[string]string{
				"line1": "..*.....*",
				"line2": "..35..633",
				"line3": ".........",
			},
			digitInformation: digitAndLocations{
				digit:      633,
				startIndex: 6,
				endIndex:   8,
			},
			want: true,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			want := testcase.want
			result := checkSurrondingsForSpecialCharacters(testcase.lines, testcase.digitInformation)
			if result != want {
				t.Fatalf("Expected %v, but got %v", want, result)
			}
		})
	}
}

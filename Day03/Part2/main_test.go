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
		want []int
	}{
		{
			name: "No Stars",
			Line: "..35..633.",
			want: []int{},
		},
		{
			name: "Stars in Middle",
			Line: "..*..*.",
			want: []int{2, 5},
		},
		{
			name: "Stars at start and middle",
			Line: "*.*..*.",
			want: []int{0, 2, 5},
		},
		{
			name: "Stars at middle and end",
			Line: "..*..**",
			want: []int{2, 5, 6},
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
		name         string
		lines        map[string]string
		starLocation int
		want         int
	}{
		// * on left
		{
			name: "* on left with above & below",
			lines: map[string]string{
				"line1": "35*....%..",
				"line2": "*.35..633.",
				"line3": "100.......",
			},
			starLocation: 0,
			want:         3500,
		},
		{
			name: "* on left with above & right",
			lines: map[string]string{
				"line1": "35*....%..",
				"line2": "*100..633.",
				"line3": "..........",
			},
			starLocation: 0,
			want:         3500,
		},
		{
			name: "* on left with right & below",
			lines: map[string]string{
				"line1": "..*....%..",
				"line2": "*100..633.",
				"line3": "35........",
			},
			starLocation: 0,
			want:         3500,
		},
		{
			name: "* on left with only 1 surrounding digit",
			lines: map[string]string{
				"line1": "..*....%..",
				"line2": "*100..633.",
				"line3": "..........",
			},
			starLocation: 0,
			want:         0,
		},
		// * on right
		{
			name: "* on right with above & below",
			lines: map[string]string{
				"line1": "35*....%35",
				"line2": "..35.....*",
				"line3": "100....100",
			},
			starLocation: 9,
			want:         3500,
		},
		{
			name: "* on right with above & left",
			lines: map[string]string{
				"line1": "35*....%35",
				"line2": "..35..100*",
				"line3": "100.......",
			},
			starLocation: 9,
			want:         3500,
		},
		{
			name: "* on right with left & below",
			lines: map[string]string{
				"line1": "35*....%..",
				"line2": "..35..100*",
				"line3": "100.....35",
			},
			starLocation: 9,
			want:         3500,
		},
		{
			name: "* on right with only 1 surrounding number",
			lines: map[string]string{
				"line1": "35*....%..",
				"line2": "..35..100*",
				"line3": "100.......",
			},
			starLocation: 9,
			want:         0,
		},
		// * in middle
		{
			name: "* in middle with 2 above",
			lines: map[string]string{
				"line1": "...35.100.",
				"line2": ".....*....",
				"line3": "..........",
			},
			starLocation: 5,
			want:         3500,
		},
		{
			name: "* in middle with 1 above",
			lines: map[string]string{
				"line1": "...351100.",
				"line2": ".....*....",
				"line3": "..........",
			},
			starLocation: 5,
			want:         0,
		},
		{
			name: "* in middle with 2 below",
			lines: map[string]string{
				"line1": "..........",
				"line2": ".....*....",
				"line3": "...35.100.",
			},
			starLocation: 5,
			want:         3500,
		},
		{
			name: "* in middle with 1 below",
			lines: map[string]string{
				"line1": "..........",
				"line2": ".....*....",
				"line3": "...351100.",
			},
			starLocation: 5,
			want:         0,
		},
		{
			name: "* in middle with 1 below and 1 above",
			lines: map[string]string{
				"line1": ".....35...",
				"line2": ".....*....",
				"line3": "......100.",
			},
			starLocation: 5,
			want:         3500,
		},
		{
			name: "* in middle with 1 right and 1 left",
			lines: map[string]string{
				"line1": "...........",
				"line2": "...35*100.",
				"line3": "..........",
			},
			starLocation: 5,
			want:         3500,
		},
		{
			name: "* in middle with left, right, up",
			lines: map[string]string{
				"line1": ".....10....",
				"line2": "...35*100.",
				"line3": "..........",
			},
			starLocation: 5,
			want:         35000,
		},
		{
			name: "* in middle with left, right, up, down",
			lines: map[string]string{
				"line1": ".....10....",
				"line2": "...35*100.",
				"line3": ".....10...",
			},
			starLocation: 5,
			want:         350000,
		},
		{
			name: "* in middle with left, right, up, down",
			lines: map[string]string{
				"line1": ".....100...",
				"line2": "...35*100.",
				"line3": ".....10...",
			},
			starLocation: 5,
			want:         3500000,
		},
		{
			name: "* in middle testing right corner",
			lines: map[string]string{
				"line1": ".....10000",
				"line2": "...35*100",
				"line3": ".....1000",
			},
			starLocation: 5,
			want:         35000000000,
		},
		{
			name: "* in middle testing left corner",
			lines: map[string]string{
				"line1": "100.......",
				"line2": "35*.......",
				"line3": "100.......",
			},
			starLocation: 2,
			want:         350000,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			want := testcase.want
			result := checkSurrondingsForDigits(testcase.lines, testcase.starLocation)
			if result != want {
				t.Fatalf("Expected %v, but got %v", want, result)
			}
		})
	}
}

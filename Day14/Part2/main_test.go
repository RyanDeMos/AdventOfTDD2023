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

func Test_moveStoneNorth(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		row    int
		column int
		expect [][]string
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    2,
			column: 2,
			expect: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "."},
			},
		},
		{
			name: "test 2",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    2,
			column: 1,
			expect: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 3",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{"O", "O", "O"},
			},
			row:    0,
			column: 0,
			expect: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{"O", "O", "O"},
			},
		},
		{
			name: "test 4",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{"O", "O", "O"},
			},
			row:    2,
			column: 0,
			expect: [][]string{
				{"O", ".", "."},
				{"O", "#", "."},
				{".", "O", "O"},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			moveStoneNorth(testcase.grid, testcase.row, testcase.column, true)
			if !reflect.DeepEqual(testcase.grid, testcase.expect) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, testcase.grid)
			}
		})
	}
}

func Test_moveAllStonesNorth(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		expect [][]string
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			expect: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "."},
			},
		},
		{
			name: "test 2",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{"O", "O", "O"},
			},
			expect: [][]string{
				{"O", ".", "O"},
				{"O", "#", "."},
				{".", "O", "."},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			moveAllStonesNorth(testcase.grid)
			if !reflect.DeepEqual(testcase.grid, testcase.expect) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, testcase.grid)
			}
		})
	}
}

func Test_moveStoneWest(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		row    int
		column int
		expect [][]string
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    2,
			column: 2,
			expect: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 2",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    2,
			column: 1,
			expect: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{"O", ".", "O"},
			},
		},
		{
			name: "test 3",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{"O", "O", "O"},
			},
			row:    0,
			column: 0,
			expect: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{"O", "O", "O"},
			},
		},
		{
			name: "test 4",
			grid: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    0,
			column: 2,
			expect: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 5",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "O"},
				{".", "O", "O"},
			},
			row:    1,
			column: 2,
			expect: [][]string{
				{"O", ".", "."},
				{".", "#", "O"},
				{".", "O", "O"},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			moveStoneWest(testcase.grid, testcase.row, testcase.column, true)
			if !reflect.DeepEqual(testcase.grid, testcase.expect) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, testcase.grid)
			}
		})
	}
}

func Test_moveAllStonesWest(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		expect [][]string
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
			expect: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{"O", "O", "."},
			},
		},
		{
			name: "test 2",
			grid: [][]string{
				{"O", ".", "O"},
				{".", "#", "O"},
				{"O", "O", "."},
			},
			expect: [][]string{
				{"O", "O", "."},
				{".", "#", "O"},
				{"O", "O", "."},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			moveAllStonesWest(testcase.grid)
			if !reflect.DeepEqual(testcase.grid, testcase.expect) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, testcase.grid)
			}
		})
	}
}

func Test_moveStoneSouth(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		row    int
		column int
		expect [][]string
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    0,
			column: 0,
			expect: [][]string{
				{".", ".", "."},
				{".", "#", "."},
				{"O", "O", "O"},
			},
		},
		{
			name: "test 2",
			grid: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    2,
			column: 1,
			expect: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 3",
			grid: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    0,
			column: 1,
			expect: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 4",
			grid: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    0,
			column: 2,
			expect: [][]string{
				{"O", ".", "."},
				{".", "#", "O"},
				{".", "O", "O"},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			moveStoneSouth(testcase.grid, testcase.row, testcase.column, true)
			if !reflect.DeepEqual(testcase.grid, testcase.expect) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, testcase.grid)
			}
		})
	}
}

func Test_moveAllStonesSouth(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		expect [][]string
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
			expect: [][]string{
				{".", ".", "."},
				{".", "#", "O"},
				{"O", "O", "O"},
			},
		},
		{
			name: "test 2",
			grid: [][]string{
				{"O", "O", "O"},
				{".", "#", "O"},
				{"O", "O", "."},
			},
			expect: [][]string{
				{".", "O", "."},
				{"O", "#", "O"},
				{"O", "O", "O"},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			moveAllStonesSouth(testcase.grid)
			if !reflect.DeepEqual(testcase.grid, testcase.expect) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, testcase.grid)
			}
		})
	}
}

func Test_moveStoneEast(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		row    int
		column int
		expect [][]string
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", ".", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    0,
			column: 0,
			expect: [][]string{
				{".", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 2",
			grid: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    2,
			column: 1,
			expect: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 3",
			grid: [][]string{
				{"O", "O", "."},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    0,
			column: 1,
			expect: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 4",
			grid: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
			row:    0,
			column: 0,
			expect: [][]string{
				{".", "O", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 5",
			grid: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{"O", ".", "O"},
			},
			row:    2,
			column: 0,
			expect: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 6",
			grid: [][]string{
				{".", ".", "O"},
				{"O", "#", "."},
				{"O", ".", "O"},
			},
			row:    1,
			column: 0,
			expect: [][]string{
				{".", ".", "O"},
				{"O", "#", "."},
				{"O", ".", "O"},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			moveStoneEast(testcase.grid, testcase.row, testcase.column, true)
			if !reflect.DeepEqual(testcase.grid, testcase.expect) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, testcase.grid)
			}
		})
	}
}

func Test_moveAllStonesEast(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		expect [][]string
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", ".", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
			expect: [][]string{
				{".", "O", "O"},
				{".", "#", "."},
				{".", "O", "O"},
			},
		},
		{
			name: "test 2",
			grid: [][]string{
				{".", ".", "."},
				{"O", "#", "O"},
				{"O", "O", "."},
			},
			expect: [][]string{
				{".", ".", "."},
				{"O", "#", "O"},
				{".", "O", "O"},
			},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			moveAllStonesEast(testcase.grid)
			if !reflect.DeepEqual(testcase.grid, testcase.expect) {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, testcase.grid)
			}
		})
	}
}

func Test_getLoad(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]string
		expect int
	}{
		{
			name: "test 1",
			grid: [][]string{
				{"O", "O", "O", "O", ".", "#", ".", "O", ".", "."},
				{"O", "O", ".", ".", "#", ".", ".", ".", ".", "#"},
				{"O", "O", ".", ".", "O", "#", "#", ".", ".", "O"},
				{"O", ".", ".", "#", ".", "O", "O", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
				{".", ".", "#", ".", ".", ".", ".", "#", ".", "#"},
				{".", ".", "O", ".", ".", "#", ".", "O", ".", "O"},
				{".", ".", "O", ".", ".", ".", ".", ".", ".", "."},
				{"#", ".", ".", ".", ".", "#", "#", "#", ".", "."},
				{"#", ".", ".", ".", ".", "#", ".", ".", ".", "."},
			},
			expect: 136,
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			result := getLoad(testcase.grid)
			if result != testcase.expect {
				t.Fatalf("Expected %v, but got %v\n", testcase.expect, result)
			}
		})
	}
}

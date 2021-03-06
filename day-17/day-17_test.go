package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	. "utils"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			112,
		},
	}

	for _, test := range tests {
		a := puzzle1(getInput(test.input))
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(getInput("./input.txt")))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			848,
		},
	}
	for _, test := range tests {
		a := puzzle2(getInput(test.input))
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}

		a = puzzle2Concurrent(getInput(test.input))
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(getInput("./input.txt")))
}

func BenchmarkPuzzle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle2(getInput("./input.txt"))
	}
}

func BenchmarkPuzzle2Concurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle2Concurrent(getInput("./input.txt"))
	}
}

func getInput(path string) string {
	file, err := ioutil.ReadFile(path)
	Check(err)
	return string(file)
}

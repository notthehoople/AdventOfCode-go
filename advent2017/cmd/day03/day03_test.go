package main

import (
	"testing"
)

func TestCalcSteps(t *testing.T) {
	var testsA = []struct {
		input    int
		expected int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}
	var testsB = []struct {
		input    int
		expected int
	}{
		{9, 10},
		{30, 54},
		{60, 122},
		{200, 304},
	}

	// Test for Part A
	for _, test := range testsA {
		if output := calcSteps(test.input, false); output != test.expected {
			t.Errorf("Part A Test Failed: {%v} inputted, {%v} expected, recieved: {%v}\n", test.input, test.expected, output)
		}
	}

	// Test for Part B
	for _, test := range testsB {
		if output := calcAdjacentValues(test.input, false); output != test.expected {
			t.Errorf("Part A Test Failed: {%v} inputted, {%v} expected, recieved: {%v}\n", test.input, test.expected, output)
		}
	}
}

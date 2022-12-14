package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

// Part B: Looking for any overlap in clearning duties at all
func elfCleaningPartialOverlap(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	var partialOverlap int
	for _, cleaningAssignments := range puzzleInput {
		var elf1Start, elf1End, elf2Start, elf2End int

		fmt.Sscanf(cleaningAssignments, "%d-%d,%d-%d\n",
			&elf1Start, &elf1End, &elf2Start, &elf2End)

		if (elf1Start >= elf2Start && elf1Start <= elf2End) ||
			(elf1End >= elf2Start && elf1End <= elf2End) ||
			(elf2Start >= elf1Start && elf2Start <= elf1End) ||
			(elf2End >= elf1Start && elf2End <= elf1End) {
			partialOverlap++
		}
	}

	return partialOverlap
}

// Part A: Looking for situations where one Elf's assignment is entirely inside another
func elfCleaning(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)

	var fullOverlap int
	for _, cleaningAssignments := range puzzleInput {
		var elf1Start, elf1End, elf2Start, elf2End int

		fmt.Sscanf(cleaningAssignments, "%d-%d,%d-%d\n",
			&elf1Start, &elf1End, &elf2Start, &elf2End)

		if debug {
			fmt.Printf("Start: %d End: %d | Start: %d End: %d\n",
				elf1Start, elf1End, elf2Start, elf2End)
		}

		if (elf1Start >= elf2Start && elf1End <= elf2End) ||
			(elf2Start >= elf1Start && elf2End <= elf1End) {
			fullOverlap++
		}

	}

	return fullOverlap
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", elfCleaning(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", elfCleaningPartialOverlap(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}

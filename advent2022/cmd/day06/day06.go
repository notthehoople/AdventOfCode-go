package main

import (
	"AdventOfCode-go/advent2022/utils"
	"fmt"
)

// Check if the 14 characters passed through are a start of message (i.e. are all different)
func isMessage(testBlock string) bool {

	for i := 0; i < len(testBlock)-1; i++ {
		for j := i + 1; j < len(testBlock); j++ {
			if testBlock[i] == testBlock[j] {
				return false
			}
		}
	}
	return true
}

// Part A: look for the start of message (14 chars all different)
func findStartOfMessage(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	communication := puzzleInput[0]

	endOfString := len(communication)

	for pos := 0; pos < endOfString-13; pos++ {
		if isMessage(communication[pos : pos+14]) {
			fmt.Println("Found at", communication[pos:pos+14])
			return pos + 14
		}

	}

	return 0
}

// Check if the 4 characters passed through are a start of packet marker (i.e. are all different)
func isPacketMarker(testBlock string) bool {
	if (testBlock[0] != testBlock[1]) && (testBlock[0] != testBlock[2]) && (testBlock[0] != testBlock[3]) {
		if (testBlock[1] != testBlock[2]) && (testBlock[1] != testBlock[3]) {
			if testBlock[2] != testBlock[3] {
				return true
			}
		}
	}
	return false
}

// Part A: look for the start of packet marker
func findStartOfPacket(filename string, part byte, debug bool) int {

	puzzleInput, _ := utils.ReadFile(filename)
	communication := puzzleInput[0]

	endOfString := len(communication)

	for pos := 0; pos < endOfString-3; pos++ {
		if isPacketMarker(communication[pos : pos+4]) {
			fmt.Println("Found at", communication[pos:pos+4])
			return pos + 4
		}

	}

	return 0
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	switch execPart {
	case 'a':
		fmt.Printf("Result is: %d\n", findStartOfPacket(filenamePtr, execPart, debug))
	case 'b':
		fmt.Printf("Result is: %d\n", findStartOfMessage(filenamePtr, execPart, debug))
	case 'z':
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	}
}

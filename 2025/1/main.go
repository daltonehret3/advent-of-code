package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

const DIAL_MAX = 99
const DIAL_MIN = 0

var DIAL_POSITION = 0

func partOne(input string, currentPosition int) int {
	dialPosition := currentPosition

	direction := input[0]

	var steps int

	// TODO: Very hacky way to get steps, needs improvement for Part Two
	if len(input) > 2 {
		steps, _ = strconv.Atoi(input[len(input)-2:])
	} else {
		steps, _ = strconv.Atoi(input[1:])
	}

	if direction == 'L' {
		dialPosition -= steps
		if dialPosition < DIAL_MIN {
			dialPosition += (DIAL_MAX + 1)
		}
	} else if direction == 'R' {
		dialPosition += steps
		if dialPosition > DIAL_MAX {
			dialPosition -= (DIAL_MAX + 1)
		}
	}

	return dialPosition
}

func partTwo(input string) int {
	return 0
}

func main() {
	dialTurns := []string{}
	password := 0
	dialPosition := 50

	file, _ := os.Open("2025/1/input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dialTurns = append(dialTurns, scanner.Text())
	}

	for _, turn := range dialTurns {
		newDialPosition := partOne(turn, dialPosition)

		dialPosition = newDialPosition

		if dialPosition == 0 {
			password ++
		}
	}

	fmt.Println("Password:", password)

	file.Close()
}

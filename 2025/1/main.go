package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

const DIAL_MAX = 99
const DIAL_MIN = 0

var dialPosition = 50

func partOne(input string, currentPosition int) (int, int) {
	password := 0
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

	if currentPosition == 0 {
		password++
	}

	return dialPosition, password
}

func partTwo(input string, startingDialPosition int) (int, int) {
	hasPassed := false
	currentPosition := startingDialPosition
	direction, rest := input[0], input[1:]

	numberOfTurns, _ := strconv.Atoi(rest)

	var steps int

	if len(input) > 2 {
		steps, _ = strconv.Atoi(input[len(input)-2:])
	} else {
		steps, _ = strconv.Atoi(input[len(input)-1:])
	}

	switch direction {
	case 'L':
		currentPosition -= steps
		if currentPosition < DIAL_MIN {
			currentPosition += (DIAL_MAX + 1)
			hasPassed = true
		}
	case 'R':
		currentPosition += steps
		if currentPosition > DIAL_MAX {
			currentPosition -= (DIAL_MAX + 1)
			hasPassed = true
		}
	}

	final := numberOfTurns / 100

	if (hasPassed || currentPosition == 0) && startingDialPosition != 0 {
		final += 1
	}

	return currentPosition, int(final)
}

func main() {
	dialTurns := []string{}
	password := 0

	file, _ := os.Open("2025/1/input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dialTurns = append(dialTurns, scanner.Text())
	}

	for _, turn := range dialTurns {
		newDialPosition, final := partTwo(turn, dialPosition)

		dialPosition = newDialPosition

		password += final

	}

	fmt.Println("Password:", password)

	file.Close()
}

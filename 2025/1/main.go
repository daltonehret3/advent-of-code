package main

import "fmt"

const DIAL_MAX = 99
const DIAL_MIN = 0

var DIAL_POSITION = 0

func turnDial(input string, currentPosition int) int {
	dialPosition := currentPosition

	direction := input[0]

	var steps int

	fmt.Sscanf(input[1:], "%d", &steps)

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

func main() {
	dialPosition := turnDial("R100", DIAL_POSITION)

	fmt.Printf("Final Dial Position: %d\n", dialPosition)

	fmt.Println("Hello, World!")
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func partOne(input string) int {
	total := 0
	splitInput := strings.Split(input, "-")

	leadingZeros := 0

	for _, split := range splitInput {
		if strings.HasPrefix(split, "0") {
			total, _ = strconv.Atoi(split)	
		}
	}
	
	fmt.Println("Split Input", splitInput)

	fmt.Println("Total", total)


	return leadingZeros
}

func main() {
	total := 0

	invalidId := partOne("01234-4351")

	total += invalidId

	fmt.Println("InvalidId", invalidId)
}

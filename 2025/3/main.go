package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func step1(input string) int {
	maxNum := -1
	maxPos := -1

	for i, character := range input {
		num := int(character - '0')
		if num > maxNum && i != len(input)-1 {
			maxNum = num
			maxPos = i
		}
	}

	secondMax := -1
	for j := maxPos + 1; j < len(input); j++ {
		num := int(input[j] - '0')
		if num > secondMax {
			secondMax = num
		}
	}

	finalString := strconv.Itoa(maxNum) + strconv.Itoa(secondMax)
	final, _ := strconv.Atoi(finalString)
	return final
}

func step2(input string) int {
	lastPosition := 11
	finalString := ""
	startingPos := 0

	for lastPosition >= 0 {
		maxNum := -1

		for i, character := range input[: len(input) - lastPosition] {
			num := int(character - '0')

			if i < startingPos {
				continue
			}

			if num > maxNum {
				maxNum = num
				startingPos = i +1
			}
		}

		finalString += strconv.Itoa(maxNum)
		lastPosition--
	}

	final, _ := strconv.Atoi(finalString)

	fmt.Println("Final: ", final)
	fmt.Println("LENGTH: ", len(finalString))

	return final
}

func main() {
	totalPower := 0

	file, _ := os.Open("2025/3/input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		totalPower += step2(scanner.Text())
	}

	fmt.Println("Final Power:", totalPower)
}

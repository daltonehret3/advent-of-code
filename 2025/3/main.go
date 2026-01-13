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

func main() {
	totalPower := 0

	file, _ := os.Open("2025/3/input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		totalPower += step1(scanner.Text())
	}

	fmt.Println("Final Power:", totalPower)
}

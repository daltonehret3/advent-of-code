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
	powerBanks := []string{}
	totalPower := 0

	file, _ := os.Open("2025/3/input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		powerBanks = append(powerBanks, scanner.Text())
	}

	for _, bank := range powerBanks {
		power := step1(bank)

		totalPower += power
	}

	fmt.Println("Final Power: ", totalPower)
}

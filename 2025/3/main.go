package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func step1(input string) int {
	final := 0
	maxNumber := 0
	secondMax := 0
	position := 0

	for i, character := range input {
		val := string(character)

		number, _ := strconv.Atoi(val)

		if (number > maxNumber) && i != len(input) -1 {
			position = i
			maxNumber = number	
		}

	}

	for j := (position + 1); j < len(input); j++ {
		secondVal := string(input[j])
		secondNum, _ := strconv.Atoi(secondVal)

		if secondNum > secondMax && j != position {
			secondMax = secondNum
		}
	}

	finalString := strconv.Itoa(maxNumber) + strconv.Itoa(secondMax)

	final, _ = strconv.Atoi(finalString)

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

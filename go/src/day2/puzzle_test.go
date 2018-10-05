package day2_test

import (
	"day2"
	"strconv"
	"strings"
	"testing"
	"utils"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1GivenExample(t *testing.T) {
	input := toNumberArray([]string{"5 1 9 5", "7 5 3", "2 4 6 8"})

	output := day2.SolvePart1(input)

	assert.Equal(t, 18, output)
}

func TestSolvePart1(t *testing.T) {
	input := toNumberArray(file.ReadLines("input"))

	output := day2.SolvePart1(input)

	assert.Equal(t, 50376, output)
}

func TestSolvePart2GivenExample(t *testing.T) {
	input := toNumberArray([]string{"5 9 2 8", "9 4 7 3", "3 8 6 5"})

	output := day2.SolvePart2(input)

	assert.Equal(t, 9, output)
}

func TestSolvePart2(t *testing.T) {
	input := toNumberArray(file.ReadLines("input"))

	output := day2.SolvePart2(input)

	assert.Equal(t, 267, output)
}

func toNumberArray(rows []string) [][]int {
	numberArray := [][]int{}

	for _, row := range rows {
		numbers := []int{}
		numbersAsStrings := strings.Split(row, " ")

		for _, numberAsString := range numbersAsStrings {
			number, _ := strconv.Atoi(numberAsString)
			numbers = append(numbers, number)
		}

		numberArray = append(numberArray, numbers)
	}

	return numberArray
}

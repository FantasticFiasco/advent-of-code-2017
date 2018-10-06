package day5_test

import (
	"strconv"
	"testing"

	"github.com/FantasticFiasco/advent-of-code-2017/day5"
	"github.com/FantasticFiasco/advent-of-code-2017/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1GivenExample(t *testing.T) {
	jumpInstructions := toNumberArray([]string{"0", "3", "0", "1", "-3"});

	var actualSteps = day5.SolvePart1(jumpInstructions);

	assert.Equal(t, 5, actualSteps)
}

func TestSolvePart1(t *testing.T) {
	jumpInstructions := toNumberArray(file.ReadLines("input"))

	actualSteps := day5.SolvePart1(jumpInstructions)

	assert.Equal(t, 388611, actualSteps)
}

func TestSolvePart2GivenExample(t *testing.T) {
	jumpInstructions := toNumberArray([]string{"0", "3", "0", "1", "-3"});

	var actualSteps = day5.SolvePart2(jumpInstructions);

	assert.Equal(t, 10, actualSteps)
}

func TestSolvePart2(t *testing.T) {
	jumpInstructions := toNumberArray(file.ReadLines("input"))

	actualSteps := day5.SolvePart2(jumpInstructions)

	assert.Equal(t, 27763113, actualSteps)
}

func toNumberArray(rows []string) []int {
	numbers := []int{}

	for _, row := range rows {
		number, _ := strconv.Atoi(row)
		numbers = append(numbers, number)
	}

	return numbers
}

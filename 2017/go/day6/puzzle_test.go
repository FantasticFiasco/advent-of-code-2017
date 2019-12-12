package day6_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/FantasticFiasco/advent-of-code-2017/day6"
	"github.com/FantasticFiasco/advent-of-code-2017/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1GivenExample(t *testing.T) {
	var redistributionCycles = day6.SolvePart1([]int{0, 2, 7, 0})

	assert.Equal(t, 5, redistributionCycles)
}

func TestSolvePart1(t *testing.T) {
	memoryBanks := toNumberArray(file.ReadLines("input"))

	redistributionCycles := day6.SolvePart1(memoryBanks)

	assert.Equal(t, 11137, redistributionCycles)
}

func TestSolvePart2GivenExample(t *testing.T) {
	var redistributionCycles = day6.SolvePart2([]int{2, 4, 1, 2});

	assert.Equal(t, 4, redistributionCycles)
}

func TestSolvePart2(t *testing.T) {
	var redistributionCycles = day6.SolvePart2([]int{14, 13, 12, 11, 9, 8, 8, 6, 6, 4, 4, 3, 1, 1, 0, 12});

	assert.Equal(t, 1037, redistributionCycles)
}

func toNumberArray(rows []string) []int {
	numbers := []int{}

	for _, numberAsString := range strings.Split(rows[0], ",") {
		number, _ := strconv.Atoi(numberAsString)
		numbers = append(numbers, number)
	}

	return numbers
}

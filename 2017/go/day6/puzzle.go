package day6

import (
	"strings"
)

func SolvePart1(memoryBanks []int) int {
	redistributionCycles := 0

	history := map[string]bool{}

	for !historyContains(history, memoryBanks) {
		addToHistory(history, memoryBanks)

		index := indexOfMax(memoryBanks)

		reallocate(memoryBanks, index)

		redistributionCycles++
	}

	return redistributionCycles
}

func SolvePart2(memoryBanks []int) int {
	// The solution to part 2 is the same as to part 1, the only thing that differs is the initial memory banks
    return SolvePart1(memoryBanks)
}

func reallocate(memoryBanks []int, index int) {
	blockCount := memoryBanks[index]

	// Clear block count
	memoryBanks[index] = 0

	// Reallocate
	for offset := 1; offset <= blockCount; offset++ {
		currentIndex := (index + offset) % len(memoryBanks)

		memoryBanks[currentIndex]++
	}
}

func historyContains(history map[string]bool, memoryBanks []int) bool {
	return history[toString(memoryBanks)]
}

func addToHistory(history map[string]bool, memoryBanks []int) {
	history[toString(memoryBanks)] = true
}

func toString(values []int) string {
	valuesAsStrings := []string{}

	for _, value := range values {
		valuesAsStrings = append(valuesAsStrings, string(value))
	}

	return strings.Join(valuesAsStrings, ",")
}

func indexOfMax(values []int) int {
	index := 0
	max := values[0]

	for i, value := range values {
		if value > max {
			index = i
			max = value
		}
	}

	return index
}

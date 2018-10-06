package day2

func SolvePart1(input [][]int) int {
	checksum := 0

	for _, rowOfNumbers := range input {
		min := min(rowOfNumbers)
		max := max(rowOfNumbers)

		checksum += max - min
	}

	return checksum
}

func SolvePart2(input [][]int) int {
	checksum := 0

	for _, rowOfNumbers := range input {
		for i := 0; i < len(rowOfNumbers); i++ {
			for j := 0; j < len(rowOfNumbers); j++ {
				if i != j && rowOfNumbers[j] != 0 && rowOfNumbers[i]%rowOfNumbers[j] == 0 {
					checksum += rowOfNumbers[i] / rowOfNumbers[j]
				}
			}
		}
	}

	return checksum
}

func min(numbers []int) int {
	min := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	return min
}

func max(numbers []int) int {
	max := numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

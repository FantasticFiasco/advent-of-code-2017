package day1

func SolvePart1(input []int) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		current := input[i]
		next := input[(i + 1) % len(input)]

		if current == next {
			sum += current
		}
	}

	return sum
}

func SolvePart2(input []int) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		current := input[i]
		halfwayAround := input[(i + len(input) / 2) % len(input)];

		if current == halfwayAround {
			sum += current
		}
	}

	return sum
}
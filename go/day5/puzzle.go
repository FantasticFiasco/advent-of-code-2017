package day5

func SolvePart1(jumpInstructions []int) int {
	currentIndex := 0
	steps := 0

	for currentIndex > -1 && currentIndex < len(jumpInstructions) {
		nextIndex := currentIndex + jumpInstructions[currentIndex];

		// Increment current instruction by 1
		jumpInstructions[currentIndex]++;

		currentIndex = nextIndex;
		steps++;
	}

	return steps
}

func SolvePart2(jumpInstructions []int) int {
	currentIndex := 0
	steps := 0

	for currentIndex > -1 && currentIndex < len(jumpInstructions) {
		nextIndex := currentIndex + jumpInstructions[currentIndex];

		// Decrement or increment current instruction depending on its value
		if jumpInstructions[currentIndex] >= 3 {
			jumpInstructions[currentIndex]--
		} else {
			jumpInstructions[currentIndex]++
		}

		currentIndex = nextIndex
		steps++
	}

	return steps
}

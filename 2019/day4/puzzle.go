package day4

import (
	"math"
	"strconv"
)

/*
--- Day 4: Secure Container ---

You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on
a sticky note, but someone threw it out.

However, they do remember a few key facts about the password:

- It is a six-digit number.
- The value is within the range given in your puzzle input.
- Two adjacent digits are the same (like 22 in 122345).
- Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

Other than the range rule, the following are true:

- 111111 meets these criteria (double 11, never decreases).
- 223450 does not meet these criteria (decreasing pair of digits 50).
- 123789 does not meet these criteria (no double).

How many different passwords within the range given in your puzzle input meet these criteria?
*/

type PasswordRange struct {
	Min int
	Max int
}

func SolvePart1(passwordRange PasswordRange) int {
	valid := 0
	for candidate := passwordRange.Min; candidate <= passwordRange.Max; candidate++ {
		success := checkTwoAdjacentDigits(candidate)
		if !success {
			continue
		}
		success = checkIncreasingOrEvenSequence(candidate)
		if !success {
			continue
		}
		valid++
	}

	return valid
}

/*
An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of
matching digits.

Given this additional criterion, but still ignoring the range rule, the following are now true:

- 112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
- 123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
- 111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).

How many different passwords within the range given in your puzzle input meet all of the criteria?
*/

func SolvePart2(passwordRange PasswordRange) int {
	valid := 0
	for candidate := passwordRange.Min; candidate <= passwordRange.Max; candidate++ {
		success := checkOnlyTwoAdjacentDigits(candidate)
		if !success {
			continue
		}
		success = checkIncreasingOrEvenSequence(candidate)
		if !success {
			continue
		}
		valid++
	}

	return valid
}

func checkTwoAdjacentDigits(password int) bool {
	passwordText := strconv.Itoa(password)
	for i := 1; i < len(passwordText); i++ {
		if passwordText[i] == passwordText[i-1] {
			return true
		}
	}
	return false
}

func checkOnlyTwoAdjacentDigits(password int) bool {
	passwordText := strconv.Itoa(password)
	for i := 1; i < len(passwordText); i++ {
		if (i == 1 || passwordText[i-2] != passwordText[i-1]) &&
			passwordText[i-1] == passwordText[i] &&
			(i == len(passwordText)-1 || passwordText[i] != passwordText[i+1]) {
			return true
		}
	}
	return false
}

func checkIncreasingOrEvenSequence(password int) bool {
	for i := 1; i < len(strconv.Itoa(password)); i++ {
		current := (password / int(math.Pow10(5-i))) % 10
		previous := (password / int(math.Pow10(5-i+1))) % 10
		if current < previous {
			return false
		}
	}
	return true
}

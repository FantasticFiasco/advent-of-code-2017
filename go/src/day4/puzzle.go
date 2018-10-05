package day4

import (
	"sort"
	"strings"
)

func SolvePart1(passphrase string) bool {
	words := map[string]bool{}

	for _, word := range strings.Split(passphrase, " ") {
		if words[word] == true {
			return false
		}

		words[word] = true
	}

	return true
}

func SolvePart2(passphrase string) bool {
	words := map[string]bool{}

	for _, word := range strings.Split(passphrase, " ") {
		runes := []rune(word)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		word := string(runes)

		if words[word] == true {
			return false
		}

		words[word] = true
	}

	return true
}

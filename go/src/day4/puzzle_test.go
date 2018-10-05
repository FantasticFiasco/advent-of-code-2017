package day4_test

import (
	"day4"
	"testing"
	"utils"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1GivenExamples(t *testing.T) {
	samples := []struct {
		passphrase string
		expected   bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, sample := range samples {
		output := day4.SolvePart1(sample.passphrase)

		assert.Equal(t, sample.expected, output)
	}
}

func TestSolvePart1(t *testing.T) {
	passphrases := file.ReadLines("input")

	validPassphrasCount := 0

	for _, passphrase := range passphrases {
		output := day4.SolvePart1(passphrase)

		if output {
			validPassphrasCount++
		}
	}

	assert.Equal(t, 451, validPassphrasCount)
}

func TestSolvePart2GivenExamples(t *testing.T) {
	samples := []struct {
		passphrase string
		expected   bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	for _, sample := range samples {
		output := day4.SolvePart2(sample.passphrase)

		assert.Equal(t, sample.expected, output)
	}
}

func TestSolvePart2(t *testing.T) {
	passphrases := file.ReadLines("input")

	validPassphrasCount := 0

	for _, passphrase := range passphrases {
		output := day4.SolvePart2(passphrase)

		if output {
			validPassphrasCount++
		}
	}

	assert.Equal(t, 223, validPassphrasCount)
}

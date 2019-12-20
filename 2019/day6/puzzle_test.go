package day6_test

import (
	"github.com/FantasticFiasco/advent-of-code/2019/day6"
	"github.com/FantasticFiasco/advent-of-code/2019/fileutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1GivenExample(t *testing.T) {
	in := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	got := day6.SolvePart1(in)
	assert.Equal(t, 42, got)
}

func TestSolvePart1(t *testing.T) {
	in := fileutils.ReadStrings("input")
	got := day6.SolvePart1(in)
	assert.Equal(t, 171213, got)
}

func TestSolvePart2GivenExample(t *testing.T) {
	in := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	got := day6.SolvePart2(in)
	assert.Equal(t, 4, got)
}

func TestSolvePart2(t *testing.T) {
	in := fileutils.ReadStrings("input")
	got := day6.SolvePart2(in)
	assert.Equal(t, 292, got)
}

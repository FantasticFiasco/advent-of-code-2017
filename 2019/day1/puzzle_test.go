package day1_test

import (
	"github.com/FantasticFiasco/advent-of-code/2019/day1"
	"github.com/FantasticFiasco/advent-of-code/2019/fileutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1GivenExample(t *testing.T) {
	for _, test := range []struct{
		in []int
		expect int
	} {
		{ in: []int{ 12 }, expect: 2 },
		{ in: []int{ 14 }, expect: 2 },
		{ in: []int{ 1969 }, expect: 654 },
		{ in: []int{ 100756 }, expect: 33583 },
	} {
		got := day1.SolvePart1(test.in)
		assert.Equal(t, test.expect, got)
	}
}

func TestSolvePart1(t *testing.T) {
	in := fileutils.ReadIntegers("input")
	got := day1.SolvePart1(in)
	assert.Equal(t, 3426455, got)
}

func TestSolvePart2GivenExample(t *testing.T) {
	for _, test := range []struct{
		in []int
		expect int
	} {
		{ in: []int{ 1969 }, expect: 966 },
		{ in: []int{ 100756 }, expect: 50346 },
	} {
		got := day1.SolvePart2(test.in)
		assert.Equal(t, test.expect, got)
	}
}

func TestSolvePart2(t *testing.T) {
	in := fileutils.ReadIntegers("input")
	got := day1.SolvePart2(in)
	assert.Equal(t, 5136807, got)
}

package day4_test

import (
	"github.com/FantasticFiasco/advent-of-code/2019/day4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1GivenExample(t *testing.T) {
	for _, test := range []struct {
		in     day4.PasswordRange
		expect int
	}{
		{in: day4.PasswordRange{Min: 111111, Max: 111111}, expect: 1},
		{in: day4.PasswordRange{Min: 223450, Max: 223450}, expect: 0},
		{in: day4.PasswordRange{Min: 123789, Max: 123789}, expect: 0},
	} {
		got := day4.SolvePart1(test.in)
		assert.Equal(t, test.expect, got)
	}
}

func TestSolvePart1(t *testing.T) {
	passwordRange := day4.PasswordRange{
		Min: 128392,
		Max: 643281,
	}
	got := day4.SolvePart1(passwordRange)
	assert.Equal(t, 2050, got)
}

func TestSolvePart2GivenExample(t *testing.T) {
	for _, test := range []struct {
		in     day4.PasswordRange
		expect int
	}{
		{in: day4.PasswordRange{Min: 112233, Max: 112233}, expect: 1},
		{in: day4.PasswordRange{Min: 123444, Max: 123444}, expect: 0},
		{in: day4.PasswordRange{Min: 111122, Max: 111122}, expect: 1},
	} {
		got := day4.SolvePart2(test.in)
		assert.Equal(t, test.expect, got)
	}
}

func TestSolvePart2(t *testing.T) {
	passwordRange := day4.PasswordRange{
		Min: 128392,
		Max: 643281,
	}
	got := day4.SolvePart2(passwordRange)
	assert.Equal(t, 1390, got)
}

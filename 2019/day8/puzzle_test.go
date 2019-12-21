package day8_test

import (
	"fmt"
	"github.com/FantasticFiasco/advent-of-code/2019/day8"
	"github.com/FantasticFiasco/advent-of-code/2019/fileutils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSolvePart1GivenExample(t *testing.T) {
	in := split("123456789012")
	got := day8.SolvePart1(3, 2, in)
	assert.Equal(t, 1, got)
}

func TestSolvePart1(t *testing.T) {
	in := split(fileutils.ReadString("input"))
	got := day8.SolvePart1(25, 6, in)
	assert.Equal(t, 2480, got)
}

func TestSolvePart2GivenExample(t *testing.T) {
	in := split("0222112222120000")
	got := day8.SolvePart2(2, 2, in)
	assert.Equal(t, []int{0, 1, 1, 0}, got)
	printScreen(2, 2, got)
}

func TestSolvePart2(t *testing.T) {
	in := split(fileutils.ReadString("input"))
	got := day8.SolvePart2(25, 6, in)
	expect := []int{
		1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0,
		0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0,
		0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0,
		0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0,
		1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0,
		1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0,
	}
	assert.Equal(t, expect, got)
	printScreen(25, 6, got)
}

func split(s string) (numbers []int) {
	for _, c := range s {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	return
}

func printScreen(width, height int, pixels []int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(pixels[y*width+x])
		}
		fmt.Print("\n")
	}
}

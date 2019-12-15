package day2_test

import (
	"fmt"
	"github.com/FantasticFiasco/advent-of-code/2019/day2"
	"github.com/FantasticFiasco/advent-of-code/2019/fileutils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestSolvePart1GivenExample(t *testing.T) {
	for _, test := range []struct{
		in string
		expect string
	} {
		{ in: "1,9,10,3,2,3,11,0,99,30,40,50", expect: "3500,9,10,70,2,3,11,0,99,30,40,50" },
		{ in: "1,0,0,0,99", expect: "2,0,0,0,99" },
		{ in: "2,3,0,3,99", expect: "2,3,0,6,99" },
		{ in: "2,4,4,5,99,0", expect: "2,4,4,5,99,9801" },
		{ in: "1,1,1,4,99,5,6,0,99", expect: "30,1,1,4,2,5,6,0,99" },
	} {
		got := join(day2.SolvePart1(split(test.in)))
		assert.Equal(t, test.expect, got)
	}
}

func TestSolvePart1(t *testing.T) {
	in := fileutils.ReadString("input")
	got := first(day2.SolvePart1(split(in)))
	assert.Equal(t, 3760627, got)
}

func TestSolvePart2(t *testing.T) {
	in := fileutils.ReadString("input")
	gotNoun, gotVerb := day2.SolvePart2(split(in), 19690720)
	assert.Equal(t, 71, gotNoun)
	assert.Equal(t, 95, gotVerb)
	fmt.Printf("100 * noun + verb = %d\n", 100 * gotNoun + gotVerb)
}

func first(numbers []int) int {
	return numbers[0]
}

func split(sequence string) []int {
	var numbers []int
	for _, s := range strings.Split(sequence, ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func join(numbers []int) string {
	var numbersText []string
	for i := range numbers {
		s := strconv.Itoa(numbers[i])
		numbersText = append(numbersText, s)
	}
	return strings.Join(numbersText, ",")
}

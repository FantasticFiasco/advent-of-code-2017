package day7_test

import (
	"github.com/FantasticFiasco/advent-of-code/2017/go/fileutils"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/FantasticFiasco/advent-of-code/2017/go/day7"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1GivenExample(t *testing.T) {
	programs := []day7.Program{
		{
			"pbga",
			66,
			[]string{},
		},
		{
			"xhth",
			57,
			[]string{},
		},
		{
			"ebii",
			61,
			[]string{},
		},
		{
			"havc",
			66,
			[]string{},
		},
		{
			"ktlj",
			57,
			[]string{},
		},
		{
			"fwft",
			72,
			[]string{"ktlj", "cntj", "xhth"},
		},
		{
			"qoyq",
			66,
			[]string{},
		},
		{
			"padx",
			45,
			[]string{"pbga", "havc", "qoyq"},
		},
		{
			"tknk",
			41,
			[]string{"ugml", "padx", "fwft"},
		},
		{
			"jptl",
			61,
			[]string{},
		},
		{
			"ugml",
			68,
			[]string{"gyxo", "ebii", "jptl"},
		},
		{
			"gyxo",
			61,
			[]string{},
		},
		{
			"cntj",
			57,
			[]string{},
		},
	}

	bottomProgram := day7.SolvePart1(programs)

	assert.Equal(t, "tknk", bottomProgram.Name)
}

func TestSolvePart1(t *testing.T) {
	programs := toProgramArray(fileutils.ReadLines("input"))

	bottomProgram := day7.SolvePart1(programs)

	assert.Equal(t, "hmvwl", bottomProgram.Name)
}

func toProgramArray(rows []string) []day7.Program {
	programs := []day7.Program{}

	for _, row := range rows {
		matches := parse(row)

		name := matches["name"]
		weight, _ := strconv.Atoi("weight")
		programsAbove := []string{}
		if matches["programsAbove"] != "" {
			for _, programAbove := range strings.Split(matches["programsAbove"], ",") {
				programsAbove = append(programsAbove, strings.TrimSpace(programAbove))
			}
		}

		program := day7.Program{
			name,
			weight,
			programsAbove,
		}

		programs = append(programs, program)
	}

	return programs
}

var rowRegex = regexp.MustCompile(`^(?P<name>\w+) \((?P<weight>\d+)\)( -> (?P<programsAbove>.*))?$`)

func parse(row string) map[string]string {
	result := map[string]string{}

	matches := rowRegex.FindStringSubmatch(row)

	for i, name := range rowRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	return result
}

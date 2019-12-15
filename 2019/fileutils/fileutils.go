package fileutils

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadString(filename string) string {
	path := filepath.Join("testdata", filename)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ReadStrings(filename string) []string {
	path := filepath.Join("testdata", filename)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bytes), "\n")
}

func ReadIntegers(filename string) []int {
	var integers []int
	for _, v := range ReadStrings(filename) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		integers = append(integers, integer)
	}
	return integers
}



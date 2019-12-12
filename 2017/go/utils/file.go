package file

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ReadLines(filename string) []string {
	path := filepath.Join("testdata", filename)
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}

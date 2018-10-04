package file

import (
	"io/ioutil"
	"path/filepath"
)

func ReadLines(filename string) string {
	path := filepath.Join("testdata", filename)
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

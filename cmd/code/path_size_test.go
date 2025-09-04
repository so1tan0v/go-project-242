package main

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getWd() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)

		return ""
	}
	fmt.Printf("Current working directory: %s\n", wd)

	return wd
}

func TestGetSizeFile(t *testing.T) {
	pathToFile := path.Join(getWd(), "testdata", "size.test")

	fmt.Println(pathToFile)
	result, err := GetSize(pathToFile)
	assert.Nil(t, err, "There aren't error")

	fmt.Println(result)
	assert.Equal(t, "19B	size.test", result)
}

func TestGetSizeDir(t *testing.T) {
	pathToFile := path.Join(getWd(), "testdata")

	result, err := GetSize(pathToFile)
	assert.Nil(t, err)

	fmt.Println(result)
	assert.Equal(t, "128B	testdata", result)
}

func TestGetSizeError(t *testing.T) {
	pathToFile := path.Join(getWd(), "has_no_file")
	result, err := GetSize(pathToFile)

	assert.NotNil(t, err, "There Error!")
	assert.Equal(t, "", result)

	pathToFile = path.Join(getWd(), "testdata", "has_no_file.json")
	result, err = GetSize(pathToFile)

	assert.NotNil(t, err, "There Error!")
	assert.Equal(t, "", result)
}

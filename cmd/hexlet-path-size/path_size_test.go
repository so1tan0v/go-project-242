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
	var pathToFile = path.Join(getWd(), "testdata", "size.test")
	var err error
	var result string

	// Combination fileName/false/false
	result, err = GetSize(pathToFile, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4480B	size.test", result)

	// Combination fileName/true/false
	err = nil
	result = ""

	result, err = GetSize(pathToFile, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.5KB	size.test", result)

	// Combination fileName/true/true
	err = nil
	result = ""

	result, err = GetSize(pathToFile, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.5KB	size.test", result)

	// Combination fileName/false/true
	err = nil
	result = ""

	result, err = GetSize(pathToFile, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4480B	size.test", result)
}

func TestGetSizeDir(t *testing.T) {
	var pathToDir = path.Join(getWd(), "testdata")
	var err error
	var result string

	// Combination fileName/false/false
	result, err = GetSize(pathToDir, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4704B	testdata", result)

	// Combination fileName/true/false
	err = nil
	result = ""

	result, err = GetSize(pathToDir, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.7KB	testdata", result)

	// Combination fileName/true/true
	err = nil
	result = ""

	result, err = GetSize(pathToDir, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.9KB	testdata", result)

	// Combination fileName/false/true
	err = nil
	result = ""

	result, err = GetSize(pathToDir, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4928B	testdata", result)
}

func TestGetSizeError(t *testing.T) {
	pathToFile := path.Join(getWd(), "has_no_file")
	result, err := GetSize(pathToFile, false, false)

	assert.NotNil(t, err, "There Error!")
	assert.Equal(t, "", result)

	pathToFile = path.Join(getWd(), "testdata", "has_no_file.json")
	result, err = GetSize(pathToFile, false, false)

	assert.NotNil(t, err, "There Error!")
	assert.Equal(t, "", result)
}

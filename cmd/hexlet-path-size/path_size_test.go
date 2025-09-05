package main

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDir = "testdata"

func TestGetSize(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result int64
	var isDir bool

	// Test file
	result, isDir, err = GetSize(0, pathToFile, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4480), result)
	assert.Equal(t, false, isDir)

	// Test dir
	err = nil
	result = 0
	isDir = false

	result, isDir, err = GetSize(0, testDir, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4704), result)
	assert.Equal(t, true, isDir)
}

func TestGetSizeAll(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result int64
	var isDir bool

	// Test file
	result, isDir, err = GetSize(0, pathToFile, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4480), result)
	assert.Equal(t, false, isDir)

	// Test dir
	err = nil
	result = 0
	isDir = false

	result, isDir, err = GetSize(0, testDir, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4928), result)
	assert.Equal(t, true, isDir)
}

func TestGetSizeRecursive(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result int64
	var isDir bool

	// Test file
	result, isDir, err = GetSize(0, pathToFile, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4480), result)
	assert.Equal(t, false, isDir)

	// Test dir
	err = nil
	result = 0
	isDir = false

	result, isDir, err = GetSize(0, testDir, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(5152), result)
	assert.Equal(t, true, isDir)
}

func TestGetSizeRecursiveAll(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result int64
	var isDir bool

	// Test file
	result, isDir, err = GetSize(0, pathToFile, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4480), result)
	assert.Equal(t, false, isDir)

	// Test dir
	err = nil
	result = 0
	isDir = false

	result, isDir, err = GetSize(0, testDir, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(6496), result)
	assert.Equal(t, true, isDir)
}

func TestGetResult(t *testing.T) {
	result, err := GetResult(path.Join(testDir, "size.test"), false, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4480B	testdata/size.test", result)
}

func TestGetResultHuman(t *testing.T) {
	result, err := GetResult(path.Join(testDir, "size.test"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.50KB	testdata/size.test", result)
}

func TestGetResultHumanAll(t *testing.T) {
	result, err := GetResult(testDir, true, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.90KB	testdata/", result)
}

func TestGetResultHumanRecursive(t *testing.T) {
	result, err := GetResult(testDir, true, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "6.50KB	testdata/", result)
}

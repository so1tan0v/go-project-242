package tests

import (
	"path"
	"testing"

	"code"

	"github.com/stretchr/testify/assert"
)

var testDir = "testdata"

func TestGetSize(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result int64

	// Test file
	result, err = code.GetSize(0, pathToFile, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4480), result)

	// Test dir
	err = nil
	result = 0

	result, err = code.GetSize(0, testDir, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4704), result)
}

func TestGetSizeAll(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result int64

	// Test file
	result, err = code.GetSize(0, pathToFile, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4480), result)

	// Test dir
	err = nil
	result = 0

	result, err = code.GetSize(0, testDir, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4928), result)

	// Test dir
	err = nil
	result = 0

	result, err = code.GetSize(0, path.Join(testDir, "dir"), true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(448), result)

	// Test dir_rec
	err = nil
	result = 0

	result, err = code.GetSize(0, path.Join(testDir, "dir", "dir_rec"), true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(448), result)
}

func TestGetSizeRecursive(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result int64

	// Test file
	result, err = code.GetSize(0, pathToFile, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4480), result)

	// Test dir
	err = nil
	result = 0

	result, err = code.GetSize(0, testDir, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(5603), result)
}

func TestGetSizeRecursiveAll(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result int64

	// Test file
	result, err = code.GetSize(0, pathToFile, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(4480), result)

	// Test dir
	err = nil
	result = 0

	result, err = code.GetSize(0, testDir, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(8291), result)

	// Test dir
	err = nil
	result = 0

	result, err = code.GetSize(0, path.Join(testDir, "dir"), true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(1120), result)

	// Test dir_rec
	err = nil
	result = 0

	result, err = code.GetSize(0, path.Join(testDir, "dir", "dir_rec"), true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(448), result)
}

func TestGetSizeRecursiveLessThen1024(t *testing.T) {
	var pathToDir = path.Join(testDir, "dir_less_1024b")
	var err error
	var result int64

	// Test dir
	err = nil
	result = 0

	result, err = code.GetSize(0, pathToDir, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, int64(3), result)
}

func TestGetResult(t *testing.T) {
	result, err := code.GetResult(path.Join(testDir, "size.test"), false, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4480B	testdata/size.test", result)
}

func TestGetResultHuman(t *testing.T) {
	result, err := code.GetResult(path.Join(testDir, "size.test"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.4KB	testdata/size.test", result)
}

func TestGetResultHumanAll(t *testing.T) {
	result, err := code.GetResult(testDir, true, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.8KB	testdata", result)
}

func TestGetResultHumanRecursive(t *testing.T) {
	result, err := code.GetResult(testDir, true, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "8.1KB	testdata", result)
}

func TestGetPathSize(t *testing.T) {
	var pathToFile = path.Join(testDir, "size.test")
	var err error
	var result string

	// Test file
	result, err = code.GetPathSize(pathToFile, true, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "4.4KB", result)

	// Test dir
	err = nil
	result = ""

	result, err = code.GetPathSize(testDir, true, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "8.1KB", result)

	// Test dir
	err = nil
	result = ""

	result, err = code.GetPathSize(path.Join(testDir, "dir"), true, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "1.1KB", result)

	// Test dir_rec
	err = nil
	result = ""

	result, err = code.GetPathSize(path.Join(testDir, "dir", "dir_rec"), true, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "448B", result)
}

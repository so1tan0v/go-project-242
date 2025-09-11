package tests

import (
	"path"
	"testing"

	"code"

	"github.com/stretchr/testify/assert"
)

var testDir = "testdata"

func TestGetPathSize_Basic(t *testing.T) {
	var err error
	var result string

	// Test file
	result, err = code.GetPathSize(path.Join(testDir, "file"), false, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2048B", result)

	result, err = code.GetPathSize(path.Join(testDir, ".file"), false, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "20B", result)

	// Test dir
	result, err = code.GetPathSize(path.Join(testDir), false, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2048B", result)

	result, err = code.GetPathSize(path.Join(testDir, "dir"), false, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "20B", result)

	result, err = code.GetPathSize(path.Join(testDir, ".dir"), false, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "20B", result)
}

func TestGetPathSize_Recursive(t *testing.T) {
	var err error
	var result string

	// Test file
	result, err = code.GetPathSize(path.Join(testDir, "file"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2048B", result)

	result, err = code.GetPathSize(path.Join(testDir, ".file"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "20B", result)

	// Test dir
	result, err = code.GetPathSize(path.Join(testDir), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2088B", result)

	result, err = code.GetPathSize(path.Join(testDir, "dir"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "40B", result)

	result, err = code.GetPathSize(path.Join(testDir, ".dir"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "40B", result)
}

func TestGetPathSize_All(t *testing.T) {
	var err error
	var result string

	// Test file
	result, err = code.GetPathSize(path.Join(testDir, "file"), false, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2048B", result)

	result, err = code.GetPathSize(path.Join(testDir, ".file"), false, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "20B", result)

	// Test dir
	result, err = code.GetPathSize(path.Join(testDir), false, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2068B", result)

	result, err = code.GetPathSize(path.Join(testDir, "dir"), false, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "20B", result)

	result, err = code.GetPathSize(path.Join(testDir, ".dir"), false, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "20B", result)
}

func TestGetPathSize_EmptyDir(t *testing.T) {
	var err error
	var result string

	result, err = code.GetPathSize(path.Join(testDir, "empty"), false, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "0B", result)

	result, err = code.GetPathSize(path.Join(testDir, "empty"), true, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "0B", result)
}

func TestGetPathSize_SingleFile(t *testing.T) {
	var err error
	var result string

	result, err = code.GetPathSize(path.Join(testDir, "file"), false, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2048B", result)

	result, err = code.GetPathSize(path.Join(testDir, "file"), false, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2.0KB", result)

	result, err = code.GetPathSize(path.Join(testDir, "file"), false, false, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2048B", result)

	result, err = code.GetPathSize(path.Join(testDir, "file"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2048B", result)
}

func TestGetPathSize_HiddenFilesIgnored(t *testing.T) {
	var err error
	var result string

	result, err = code.GetPathSize(path.Join(testDir), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2088B", result)

	result, err = code.GetPathSize(path.Join(testDir, "dir"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "40B", result)

	result, err = code.GetPathSize(path.Join(testDir, ".dir"), true, false, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "40B", result)
}

func TestGetPathSize_HumanReadable(t *testing.T) {
	var err error
	var result string

	result, err = code.GetPathSize(path.Join(testDir, "file"), false, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2.0KB", result)

	result, err = code.GetPathSize(path.Join(testDir, ".file"), false, true, false)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "20B", result)

	result, err = code.GetPathSize(path.Join(testDir), false, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2.0KB", result)

	result, err = code.GetPathSize(path.Join(testDir), true, true, true)

	assert.Nil(t, err, "There aren't error")
	assert.Equal(t, "2.2KB", result)
}

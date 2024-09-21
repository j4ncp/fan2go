package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileHasPermissionsUserIsRoot(t *testing.T) {
	if os.Getuid() != 0 {
		t.Skip("Skipping tests which require root")
	}

	// GIVEN
	filePath := "./testfile"

	filePerm := os.FileMode(0o700)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePerm)
	assert.NoError(t, err)
	err = os.Chown(filePath, 0, 1000)
	assert.NoError(t, err)
	err = os.Chmod(filePath, filePerm)
	assert.NoError(t, err)

	defer file.Close()
	defer os.Remove(filePath)

	// WHEN
	result, err := CheckFilePermissionsForExecution(filePath)

	// THEN
	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestFileHasPermissionsGroupIsRootAndHasWrite(t *testing.T) {
	if os.Getuid() != 0 {
		t.Skip("Skipping tests which require root")
	}

	// GIVEN
	filePath := "./testfile"

	filePerm := os.FileMode(0o770)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePerm)
	assert.NoError(t, err)
	err = os.Chown(filePath, 0, 0)
	assert.NoError(t, err)
	err = os.Chmod(filePath, filePerm)
	assert.NoError(t, err)

	defer file.Close()
	defer os.Remove(filePath)

	// WHEN
	result, err := CheckFilePermissionsForExecution(filePath)

	// THEN
	assert.Equal(t, true, result)
	assert.NoError(t, err)
}

func TestFileHasPermissionsGroupOtherThanRootHasWritePermission(t *testing.T) {
	if os.Getuid() != 0 {
		t.Skip("Skipping tests which require root")
	}

	// GIVEN
	filePath := "./testfile"

	filePerm := os.FileMode(0o720)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePerm)
	assert.NoError(t, err)
	err = os.Chown(filePath, 0, 1000)
	assert.NoError(t, err)
	err = os.Chmod(filePath, filePerm)
	assert.NoError(t, err)

	defer file.Close()
	defer os.Remove(filePath)

	// WHEN
	result, err := CheckFilePermissionsForExecution(filePath)

	// THEN
	assert.Equal(t, false, result)
	assert.Error(t, err)
}

func TestFileHasPermissionsOtherHasWritePermission(t *testing.T) {
	if os.Getuid() != 0 {
		t.Skip("Skipping tests which require root")
	}

	// GIVEN
	filePath := "./testfile"

	filePerm := os.FileMode(0o702)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePerm)
	assert.NoError(t, err)
	err = os.Chown(filePath, 0, 1000)
	assert.NoError(t, err)
	err = os.Chmod(filePath, filePerm)
	assert.NoError(t, err)

	defer file.Close()
	defer os.Remove(filePath)

	// WHEN
	result, err := CheckFilePermissionsForExecution(filePath)

	// THEN
	assert.Equal(t, false, result)
	assert.Error(t, err)
}

func TestReadIntFromFile_Success(t *testing.T) {
	// GIVEN
	filePath := "../../test/file_fan_rpm"

	// WHEN
	result, err := ReadIntFromFile(filePath)

	// THEN
	assert.Equal(t, 2150, result)
	assert.NoError(t, err)
}

func TestReadIntFromFile_FileNotFound(t *testing.T) {
	// GIVEN
	filePath := "../../not exists"

	// WHEN
	result, err := ReadIntFromFile(filePath)

	// THEN
	assert.Equal(t, -1, result)
	assert.Error(t, err)
}

func TestReadIntFromFile_FileEmpty(t *testing.T) {
	// GIVEN
	filePath := "./empty_file"
	os.Create(filePath)
	defer os.Remove(filePath)

	// WHEN
	result, err := ReadIntFromFile(filePath)

	// THEN
	assert.Equal(t, -1, result)
	assert.Error(t, err)
}

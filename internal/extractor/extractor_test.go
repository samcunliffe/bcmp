package extractor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/samcunliffe/bcmp/internal/bcmptest"
	"github.com/samcunliffe/bcmp/internal/organiser"
	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/Artist - Album.zip"

	err := Extract(testfile, destination)
	assert.NoError(t, err, "Extract(%q, %q) returned error: %v", testfile, destination, err)

	// Check that expected files exist
	wantDestination := filepath.Join(destination, "Artist", "Album")
	assert.DirExists(t, wantDestination)

	wantFiles := []string{
		"01 First Track.flac",
		"02 Second Track.flac",
		"03 Third Track.flac",
	}
	for _, wantFile := range wantFiles {
		wantPath := filepath.Join(wantDestination, wantFile)
		assert.FileExists(t, wantPath, "Expected file %s does not exist after extraction", wantPath)
	}
}

func TestExtractDryRun(t *testing.T) {
	destination := t.TempDir()
	bcmptest.AssertDirEmpty(t, destination, "Setup failed: destination %q not empty", destination)
	testfile := "testdata/Artist - Album.zip"

	organiser.Config.DryRun = true
	defer func() { organiser.Config.DryRun = false }()

	err := Extract(testfile, destination)
	assert.NoError(t, err, "Extract(%q, %q) in dry run mode returned error: %v", testfile, destination, err)

	// Check that no files were created
	bcmptest.AssertDirEmpty(t, destination, "Extract in dry run mode modified destination %q", destination)
}

func TestExtractErrors(t *testing.T) {
	testCases := []struct {
		wantContains string
		testfile     string
	}{
		{
			wantContains: "no such file or directory",
			testfile:     "testdata/Artist - Album That Doesnt Exist.zip",
		},
		{
			wantContains: "not a zip archive",
			testfile:     "testdata/Artist - Album.txt",
		},
		{
			wantContains: "expected only one ' - ' separator",
			testfile:     "testdata/Artist - Album - WhatIsThis.zip",
		},
	}

	for _, tc := range testCases {
		err := Extract(tc.testfile, "./")
		assert.ErrorContains(t, err, tc.wantContains, "Expected error '%s', got: %v", tc.wantContains, err)
	}
}

func TestEmptyArchive(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/empty.zip"

	err := unzipAndRename(testfile, destination)
	assert.ErrorContains(t, err, "not a valid zip file", "Expected error for empty archive, got: %v", err)
}

func TestInvalidArchive(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/one-file-no-music.zip"

	err := unzipAndRename(testfile, destination)
	assert.ErrorContains(t, err, "filename does not have a valid music file suffix",
		"Expected error for invalid archive, got: %v", err)
}

func TestNoFilePermissions(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/Artist - Album.zip"

	// Remove write permissions from destination
	err := os.Chmod(destination, 0555) // TODO can I avoid magic numbers and platform specific?
	assert.NoError(t, err, "Failed to change permissions in test setup: %v", err)
	defer os.Chmod(destination, 0755) // Restore permissions after test

	err = unzipAndRename(testfile, destination)
	assert.ErrorContains(t, err, "permission denied", "Expected permission denied error, got: %v", err)
}

func TestArchiveOnlyDirectory(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/no-files-only-directories.zip"

	err := unzipAndRename(testfile, destination)
	assert.ErrorContains(t, err, "contains a directory", "Expected error about directories, got: %v", err)
}

// Test for Zip Slip vulnerability
// https://security.snyk.io/research/zip-slip-vulnerability
func TestArchiveWithZipSlip(t *testing.T) {
	destination := t.TempDir()

	// Archive contains file: '../../ive-escaped'
	// which should not escape the destination directory
	testfile := "testdata/archive-with-path-outside-slip.zip"

	err := unzipAndRename(testfile, destination)
	assert.ErrorContains(t, err, "invalid file path",
		"Expected error about invalid file path, got: %v", err.Error())
}

func TestArchiveWithUnParsableFileName(t *testing.T) {
	destination := t.TempDir()

	// Archive contains file: 'Artist - Album - Track One.flac'
	testfile := "testdata/unparsable-filename.zip"

	err := unzipAndRename(testfile, destination)
	assert.ErrorContains(t, err, "error parsing music file name",
		"Expected error about unparsable filename, got: %v", err)
}

package extractor

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractAndRename(t *testing.T) {

	destination := t.TempDir()
	testfile := "testdata/Artist - Album.zip"

	err := ExtractAndRename(testfile, destination)
	if err != nil {
		t.Fatalf("ExtractAndRename returned an error: %v", err)
	}

	// Check that expected files exist
	wantFiles := []string{
		"01 First Track.flac",
		"02 Second Track.flac",
		"03 Third Track.flac",
	}
	for _, wantFile := range wantFiles {
		wantPath := filepath.Join(destination, wantFile)
		if _, err := os.Stat(wantPath); os.IsNotExist(err) {
			t.Errorf("Expected file %s does not exist after extraction", wantPath)
		}
	}
}

func TestEmptyArchive(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/empty.zip"

	err := ExtractAndRename(testfile, destination)
	if err == nil {
		t.Errorf("Expected an error for an empty archive")
	}
	if !strings.Contains(err.Error(), "not a valid zip file") {
		t.Errorf("Expected error message about empty archive, got: %v", err.Error())
	}
}

func TestInvalidArchive(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/one-file-no-music.zip"

	err := ExtractAndRename(testfile, destination)
	if err == nil {
		t.Fatal("Expected an error for an invalid archive")
	}
	if !strings.Contains(err.Error(), "filename does not have a valid music file suffix") {
		t.Errorf("Expected error message about no valid music files, got: %v", err.Error())
	}
}

func TestNoFilePermissions(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/Artist - Album.zip"

	// Remove write permissions from destination
	err := os.Chmod(destination, 0555) // TODO can I avoid magic numbers and platform specific?
	if err != nil {
		t.Fatalf("Failed to change permissions of destination directory: %v", err)
	}
	defer os.Chmod(destination, 0755) // Restore permissions after test

	err = ExtractAndRename(testfile, destination)
	if err == nil {
		t.Errorf("Expected an error due to lack of write permissions")
	}
	if !strings.Contains(err.Error(), "permission denied") {
		t.Errorf("Expected permission denied error, got: %v", err.Error())
	}
}

func TestArchiveOnlyDirectory(t *testing.T) {
	destination := t.TempDir()
	testfile := "testdata/no-files-only-directories.zip"

	err := ExtractAndRename(testfile, destination)
	if err == nil {
		t.Errorf("Expected an error for archive with only directories")
	}
	if !strings.Contains(err.Error(), "contains a directory") {
		t.Errorf("Expected error message about directories only, got: %v", err.Error())
	}
}

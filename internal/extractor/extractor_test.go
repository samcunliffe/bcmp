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
		t.Errorf("Expected an error for an invalid archive")
	}
	if !strings.Contains(err.Error(), "filename does not have a valid music file suffix") {
		t.Errorf("Expected error message about no valid music files, got: %v", err.Error())
	}
}

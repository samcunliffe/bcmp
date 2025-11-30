package extractor

import (
	"os"
	"path/filepath"
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

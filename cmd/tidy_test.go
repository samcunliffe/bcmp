package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestTidyEndToEndProcessing(t *testing.T) {
	destination := t.TempDir()
	source := "testdata/Artist - Album - 01 Track.flac"
	defer func() {
		f, _ := os.Create(source)
		os.WriteFile(source, []byte("Just a non-empty test file."), 0644)
		f.Close()
	}()

	// Buffer to capture output
	buf := &bytes.Buffer{}
	rootCmd.SetOut(buf)

	// Execute the tidy command
	rootCmd.SetArgs([]string{
		"tidy",
		"testdata/Artist - Album - 01 Track.flac",
		"--destination",
		destination,
	})
	err := rootCmd.Execute()
	rootCmd.SetArgs(nil)

	if err != nil {
		t.Errorf("Expected no error for tidy command, got %e", err)
	}
}

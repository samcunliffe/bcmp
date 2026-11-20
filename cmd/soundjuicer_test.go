package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/samcunliffe/bcmp/internal/bcmptest"
	"github.com/stretchr/testify/assert"
)

func TestSoundJuicerEndToEndProcessing(t *testing.T) {
	// Setup: copy test data to a temporary directory
	destination := t.TempDir()
	source := "testdata/Artist"
	copied := destination + "/Artist/Album"

	err := os.CopyFS(destination, os.DirFS(source))
	assert.NoError(t, err, "Error in test setup copying over test files: %v", err)
	assert.DirExists(t, copied, "Expected directory %q to exist after copying", copied)

	// Buffer to capture output
	buf := &bytes.Buffer{}
	rootCmd.SetOut(buf)

	// Execute the tidy command
	rootCmd.SetArgs([]string{
		"soundjuicer",
		copied,
	})
	defer rootCmd.SetArgs(nil)

	err = rootCmd.Execute()

	assert.NoError(t, err, "Expected no error for soundjuicer command, got %v", err)

	nfiles, err := bcmptest.DirCount(copied)
	assert.NoError(t, err, "Error counting files in %q: %v", copied, err)
	assert.Equal(t, 2, nfiles, "Expected 2 files in %q after soundjuicer command, got %d", copied, nfiles)

	expectedFiles := []string{
		copied + "/01 First track.flac",
		copied + "/02 Second track.flac",
	}
	for _, file := range expectedFiles {
		assert.FileExists(t, file, "Expected file %q to exist after soundjuicer command", file)
	}
}

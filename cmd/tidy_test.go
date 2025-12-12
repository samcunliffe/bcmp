package cmd

import (
	"bytes"
	"testing"

	"github.com/samcunliffe/bcmp/internal/bcmptest"
	"github.com/stretchr/testify/assert"
)

func TestTidyEndToEndProcessing(t *testing.T) {
	destination := t.TempDir()
	source := "testdata/Artist - Album - 01 Track.flac"
	defer bcmptest.PutFileBack(t, source)

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
	defer rootCmd.SetArgs(nil)

	err := rootCmd.Execute()

	assert.NoError(t, err, "Expected no error for tidy command, got %v", err)
}

package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractEndToEndProcessing(t *testing.T) {
	destination := t.TempDir()
	source := "testdata/Artist - Album.zip"

	// Buffer to capture output
	buf := &bytes.Buffer{}
	rootCmd.SetOut(buf)

	// Actually execute the command
	rootCmd.SetArgs([]string{
		"extract",
		source,
		"--destination",
		destination,
	})
	defer rootCmd.SetArgs(nil)
	err := rootCmd.Execute()

	got := buf.String()
	assert.NoError(t, err, "Expected no error when processing %s", source)
	assert.NotContains(t, got, "Error", "Expected no error messages in output for %s", source)
}

package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractEndToEndProcessing(t *testing.T) {
	destination := t.TempDir()

	testCases := []struct {
		filename  string
		wantError bool
	}{
		{"testdata/Artist - Album.zip", false},
		{"testdata/Artist - Nonexistent Album.zip", true},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {

			// Buffer to capture output
			buf := &bytes.Buffer{}
			rootCmd.SetOut(buf)

			// Actually execute the command
			rootCmd.SetArgs([]string{"extract", tc.filename, "--destination", destination})
			err := rootCmd.Execute()
			defer rootCmd.SetArgs(nil)

			if tc.wantError {
				assert.Error(t, err, "Expected error when processing %s", tc.filename)
				return
			}

			got := buf.String()
			assert.NoError(t, err, "Expected no error when processing %s", tc.filename)
			assert.NotContains(t, got, "Error", "Expected no error messages in output for %s", tc.filename)
		})
	}
}

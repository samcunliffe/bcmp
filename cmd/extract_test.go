package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestExtractEndToEndProcessing(t *testing.T) {
	destination := t.TempDir()

	testCases := []struct {
		filename   string
		want_error bool
	}{
		{"testdata/Artist - Album.zip", false},
		{"testdata/Artist - Nonexistent Album.zip", true},
	}

	for _, testcase := range testCases {

		// Buffer to capture output
		buf := &bytes.Buffer{}
		rootCmd.SetOut(buf)

		// Actually execute the command
		rootCmd.SetArgs([]string{"extract", testcase.filename, "--destination", destination})
		err := rootCmd.Execute()
		rootCmd.SetArgs(nil) // cleanup for next iteration or next execution

		gotOutput := buf.String()
		if testcase.want_error {
			if err == nil {
				t.Errorf("Expected error when processing %s, got %q", testcase.filename, gotOutput)
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error when processing %s, got %v", testcase.filename, err)
			}
			if strings.Contains(gotOutput, "Error") {
				t.Errorf("Expected no error messages in output, got %q", gotOutput)
			}
		}
	}
}

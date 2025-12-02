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
		// {"testdata/Artist - Nonexistent Album.zip", true},
	}

	for _, testcase := range testCases {

		// Buffer to capture output
		buf := &bytes.Buffer{}
		extractCmd.SetOut(buf)

		// Actually execute the command
		extractCmd.SetArgs([]string{testcase.filename, "--destination", destination})
		defer extractCmd.SetOut(nil)
		err := extractCmd.Execute()

		output := buf.String()
		if testcase.want_error {
			if err == nil {
				t.Errorf("Expected error when processing %s, got %q", testcase.filename, output)
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error when processing %s, got %v", testcase.filename, err)
			}
			if strings.Contains(output, "Error") {
				t.Errorf("Expected no error messages in output, got %q", output)
			}
		}
	}
}

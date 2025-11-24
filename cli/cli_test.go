package cli

import (
	"bytes"
	"strings"
	"testing"
)

// TODO: this modifies the global cmd variable without modifying it in between.
// Figure out the idiomatic Go way to not do this.

func TestCLINoArgs(t *testing.T) {
	// Buffer to capture output
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)

	// Actually execute the command with no args - should prompt for zip file
	err := cmd.Execute()
	if err != nil {
		t.Errorf("Expected no error when executing with no args, got %v", err)
	}
	output := buf.String()

	// Check output contains expected prompt
	expect := "Please provide the path to a Bandcamp zip file."
	if !strings.Contains(output, expect) {
		t.Errorf("Expected output to contain %q, got %q", expect, output)
	}
}

func TestCLIHelp(t *testing.T) {
	// Exceptable ways to ask for help
	helpFlags := []string{"-h", "--help"}

	// Helpful substrings expected in the help output
	wantSubstrings := []string{
		"Extract and organise Bandcamp music files.",
		"Usage:",
		"Flags:",
	}

	for _, flag := range helpFlags {
		// Buffer to capture output
		buf := &bytes.Buffer{}
		cmd.SetOut(buf)

		// Set the help flag as argument
		args := []string{flag}
		cmd.SetArgs(args)

		// Actually execute the command - should show help without error
		err := cmd.Execute()
		if err != nil {
			t.Errorf("Expected no error for help flag %s, got %v", flag, err)
		}

		// Check output contains helpful substrings
		output := buf.String()
		for _, want := range wantSubstrings {
			if !strings.Contains(output, want) {
				t.Errorf("Help output '%s' missing expected substring: %s", want, output)
			}
		}
	}
}

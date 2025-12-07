package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCmdNoArgs(t *testing.T) {
	// Buffer to capture output
	buf := &bytes.Buffer{}
	SetOut(buf)

	// Actually execute the command with no args - should show description and usage
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Expected no error when executing with no args, got %v", err)
	}
	output := buf.String()

	// Check output contains expected prompt
	expect := "Extract and organise Bandcamp music files."
	if !strings.Contains(output, expect) {
		t.Errorf("Expected output to contain %q, got %q", expect, output)
	}
}

func TestRootCmdHelp(t *testing.T) {
	// Acceptable ways to ask for help
	helpFlags := []string{"-h", "--help", "help"}

	// Helpful substrings expected in the help output
	wantSubstrings := []string{
		"Extract and organise Bandcamp music files.",
		"Usage:",
		"Flags:",
	}

	for _, flag := range helpFlags {
		// Buffer to capture output
		buf := &bytes.Buffer{}
		SetOut(buf)

		// Set the help flag as argument
		args := []string{flag}
		rootCmd.SetArgs(args)

		// Actually execute the command - should show help without error
		err := rootCmd.Execute()
		rootCmd.SetArgs(nil) // cleanup for next iteration or next execution
		if err != nil {
			t.Errorf("Expected no error for help flag %s, got %v", flag, err)
		}

		// Check output contains helpful substrings
		output := buf.String()
		for _, want := range wantSubstrings {
			if !strings.Contains(output, want) {
				t.Errorf("Help output '%s' missing expected substring: %s", output, want)
			}
		}
	}
}

func TestRootExecuteSmoke(t *testing.T) {
	Execute()
}

func TestRootExecuteError(t *testing.T) {

	// Monkey patch osExit to prevent actual exit, record exit code
	originalOsExit := osExit
	defer func() { osExit = originalOsExit }()
	got, want := 0, 1
	osExit = func(code int) {
		got = code
	}

	// Temporarily replace rootCmd with one that always errors
	originalRootCmd := rootCmd
	defer func() { rootCmd = originalRootCmd }()

	rootCmd = &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("simulated error")
		},
	}

	// Now run the function
	Execute()

	if got != want {
		t.Errorf("Expected exit code: %d, got: %d", want, got)
	}
}

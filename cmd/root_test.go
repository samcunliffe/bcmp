package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRootCmdNoArgs(t *testing.T) {
	// Buffer to capture output
	buf := &bytes.Buffer{}
	SetOut(buf)

	// Actually execute the command with no args - should show description and usage
	err := rootCmd.Execute()
	assert.NoError(t, err, "Expected no error when executing with no args, got %v", err)

	// Check got contains expected prompt
	got := buf.String()
	want := "Extract and organise Bandcamp music files."
	assert.Contains(t, got, want, "Expected output to contain %q, got %q", want, got)
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
		t.Run(flag, func(t *testing.T) {
			// Buffer to capture output
			buf := &bytes.Buffer{}
			SetOut(buf)

			// Set the help flag as argument
			args := []string{flag}
			rootCmd.SetArgs(args)
			defer rootCmd.SetArgs(nil)

			// Actually execute the command - should show help without error
			err := rootCmd.Execute()

			assert.NoError(t, err, "Expected no error for help flag %s, got %v", flag, err)

			// Check output contains helpful substrings
			output := buf.String()
			for _, want := range wantSubstrings {
				assert.Contains(t, output, want, "Help output for flag %s missing expected substring: %s", flag, want)
			}
		})
	}
}

func TestRootExecuteSmoke(t *testing.T) {
	assert.NotPanics(t, Execute, "Expected Execute() not to panic")
}

func TestRootExecuteError(t *testing.T) {
	// Monkey patch osExit to prevent actual exit, record exit code
	originalOsExit := osExit
	defer func() { osExit = originalOsExit }()
	got, want := 0, 1
	osExit = func(code int) { got = code }

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

	assert.Equal(t, want, got, "Expected exit code: %d, got: %d", want, got)
}

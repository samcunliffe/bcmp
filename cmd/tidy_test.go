package cmd

import (
	"strings"
	"testing"
)

func TestTidyNoOp(t *testing.T) {
	t.Skip()
	rootCmd.SetArgs([]string{"tidy"})
	defer rootCmd.SetArgs(nil)

	err := rootCmd.Execute()
	if err == nil {
		t.Errorf("Expected error for tidy command, got nil")
	}

	want := "not implemented"
	got := err.Error()

	if !strings.Contains(got, want) {
		t.Errorf("Expected error %q, got %q", want, got)
	}
}

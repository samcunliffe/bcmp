package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/samcunliffe/bcmp/cmd"
)

// Just ensure main() runs without error, panic, or non-zero exit.
func TestMainSmoke(t *testing.T) {

	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	main()
	got := buf.String()

	for _, dontWant := range []string{"error", "ERROR", "Error", "panic"} {
		if strings.Contains(got, dontWant) {
			t.Errorf("Expected no error messages in output, got %q", got)
		}
	}

	if !strings.Contains(got, "Usage:") {
		t.Errorf("Expected usage message in output, got %q", got)
	}
}

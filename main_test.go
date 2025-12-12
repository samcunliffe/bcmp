package main

import (
	"bytes"
	"testing"

	"github.com/samcunliffe/bcmp/cmd"
	"github.com/stretchr/testify/assert"
)

// Just ensure main() runs without error, panic, or non-zero exit.
func TestMainSmoke(t *testing.T) {
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)

	main()

	got := buf.String()
	assert.Contains(t, got, "Usage:")
	for _, dontWant := range []string{"error", "ERROR", "Error", "panic"} {
		assert.NotContains(t, got, dontWant)
	}
}

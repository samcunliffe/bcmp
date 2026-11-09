package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRomanNumeral(t *testing.T) {
	var testCases = []struct {
		input string
		want  bool
	}{
		{"I", true},
		{"II", true},
		{"III", true},
		{"IV", true},
		{"V", true},
		{"VI", true},
		{"VII", true},
		{"VIII", true},
		{"IX", true},
		{"X", true},
		{"MCMXCIX", true}, // 1999
		{"MM", true},      // 2000
		{"CMXLIX", true},  // 949
		{"CDXC", true},    // 490
		// --
		{"", false},
		{"SSS", false},
		{"SAM", false},
		{"A", false},
		{"B", false},
	}
	for _, tc := range testCases {
		got := isRomanNumeral(tc.input)
		assert.Equal(t, tc.want, got, "isRomanNumeral(%q) = %v; want %v", tc.input, got, tc.want)
	}
}

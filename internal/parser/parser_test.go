package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtension(t *testing.T) {
	var testCases = []struct {
		input string
		want  string
	}{
		{"song.flac", ".flac"},
		{"song.mp3", ".mp3"},
		{"song.ogg", ".ogg"},
		{"archive.zip", ".zip"},
		{"no_extension", ""},
		{"multiple.dots.in.name.txt", ".txt"},
	}
	for _, tc := range testCases {
		got := Extension(tc.input)
		assert.Equal(t, got, tc.want, "Extension(%q) = %q; want %q", tc.input, got, tc.want)
	}
}

func TestParserConfig(t *testing.T) {
	// Defaults should be false
	assert.False(t, Config.TitleCase, "Expected default TitleCase to be false, got %v", Config.TitleCase)
	assert.False(t, Config.Debug, "Expected default Debug to be false, got %v", Config.Debug)

	// Sanity check that we can set a value
	Config.TitleCase = true
	defer func() { Config.TitleCase = false }()
	assert.True(t, Config.TitleCase, "Expected TitleCase to be true after setting, got %v", Config.TitleCase)
}

func TestToTitleCase(t *testing.T) {
	var testCases = []struct {
		input string
		want  string
	}{
		{"SYLOSIS", "Sylosis"},                                       // Made up word
		{"ROSALÍA - SAOKO", "Rosalía - Saoko"},                       // With accent
		{"IVAR BJØRNSON", "Ivar Bjørnson"},                           // Scandinavian letter
		{"THE OTHER SIDE OF ANGER", "The Other Side of Anger"},       // 'Small' word from UPPER
		{"bleed from within", "Bleed from Within"},                   // 'Small' word from lower
		{"THE BEGINNING OF THE END", "The Beginning of the End"},     // Multiple 'small' words
		{"the end of all we know", "The End of All We Know"},         // Multiple 'small' words
		{"BABYMETAL - メギツネ", "Babymetal - メギツネ"},                     // Katakana
		{"BABYMETAL - SONG 4 (4の歌)", "Babymetal - Song 4 (4の歌)"},     // Hiragana and Kanji
		{"06 THE OTHER SIDE OF ANGER", "06 The Other Side of Anger"}, // Leading track number and beginning with a 'small' word
	}
	for _, tc := range testCases {
		got := toTitleCase(tc.input)
		assert.Equal(t, got, tc.want, "toTitleCase(%q) = %q; want %q", tc.input, got, tc.want)
	}
}

func TestNumberPrefix(t *testing.T) {
	var testCases = []struct {
		input string
		want  int
	}{
		{"01 Song Title", 1},
		{"1 Song Title", 1},
		{"10 Another Song", 10},
		{"99 Last Song", 99},
		{"No Number Here", -1},
		{"", -1},
	}
	for _, tc := range testCases {
		expectError := tc.want == -1
		got, _, err := numberPrefix(tc.input)
		assert.Equal(t, got, tc.want, "numberPrefix(%q) = %d; want %d", tc.input, got, tc.want)

		if expectError {
			assert.ErrorContains(t, err, "track number extraction",
				"numberPrefix(%q) did not return error when expected", tc.input)
			return
		}
		assert.NoError(t, err, "numberPrefix(%q) returned error: %v", tc.input, err)
	}
}

func TestHasTrackNumber(t *testing.T) {
	var testCases = []struct {
		input string
		want  bool
	}{
		{"01 Song Title", true},
		{"10 Another Song", true},
		{"99 Last Song", true},
		{"No Number Here", false},
		{"", false},
	}
	for _, tc := range testCases {
		got := hasTrackNumber(tc.input)
		assert.Equal(t, got, tc.want, "hasTrackNumber(%q) = %v; want %v", tc.input, got, tc.want)
	}
}

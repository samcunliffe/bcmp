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
		assert.Equal(t, tc.want, got, "Extension(%q) = %q; want %q", tc.input, got, tc.want)
	}
}

func TestParserConfig(t *testing.T) {
	// Defaults should be false
	assert.False(t, Config.ApplyTitleCase, "Expected default ApplyTitleCase to be false, got %v", Config.ApplyTitleCase)

	// Sanity check that we can set a value
	Config.ApplyTitleCase = true
	defer func() { Config.ApplyTitleCase = false }()
	assert.True(t, Config.ApplyTitleCase, "Expected ApplyTitleCase to be true after setting, got %v", Config.ApplyTitleCase)
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
		{"AMENRA - MASS IV", "Amenra - Mass IV"},                     // Roman numerals
		{"amenra - mass iv", "Amenra - Mass IV"},                     // Roman numerals
		{"BATMAN V SUPERMAN", "Batman V Superman"},                   // 'v' is an abbreviation but also a Roman numeral so expect capital (Roman numerals have priority - an unlikely corner case)
		{"Animals As Leaders - CAFO", "Animals as Leaders - Cafo"},   // Abbreviations aren't handled so CAFO becomes Cafo
	}
	for _, tc := range testCases {
		got := toTitleCase(tc.input)
		assert.Equal(t, tc.want, got, "toTitleCase(%q) = %q; want %q", tc.input, got, tc.want)
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
		assert.Equal(t, tc.want, got, "numberPrefix(%q) = %d; want %d", tc.input, got, tc.want)

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
		{"01 Song Title", true}, // Normal track numbers
		{"10 Another Song", true},
		{"99 Last Song", true},
		{"No Number Here", false}, // No track number
		{"01", true},              // Only track number, no title
		{"0001", true},            // Track number with leading zeros
		{"010", true},             // Track number with leading zeros
		{"1", true},               // Single digit track number
		{"", false},
		{"Song Title First 01", false},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := hasTrackNumber(tc.input)
			assert.Equal(t, tc.want, got, "hasTrackNumber(%q) = %v; want %v", tc.input, got, tc.want)
		})
	}
}

func TestSplitOnHyphen(t *testing.T) {
	var testCases = []struct {
		input string
		want1 string
		want2 string
	}{
		{"Artist - Album", "Artist", "Album"},
		{"Artist - Album - Song", "Artist", "Album - Song"},
		{"Artist - Album - Song - Extra", "Artist", "Album - Song - Extra"},
		{"No Hyphen Here", "", "No Hyphen Here"},
		{" - Leading Hyphen", "", "Leading Hyphen"},
		{"Trailing Hyphen - ", "Trailing Hyphen", ""},
		{"", "", ""},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got1, got2 := splitOnHyphen(tc.input)
			assert.Equal(t, tc.want1, got1, "splitOnHyphen(%q) = %q, %q; want %q, %q", tc.input, got1, got2, tc.want1, tc.want2)
			assert.Equal(t, tc.want2, got2, "splitOnHyphen(%q) = %q, %q; want %q, %q", tc.input, got1, got2, tc.want1, tc.want2)
		})
	}
}

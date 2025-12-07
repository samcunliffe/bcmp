package parser

import "testing"

func TestParserConfig(t *testing.T) {
	// Defaults should be false
	if Config.TitleCase {
		t.Errorf("Expected default TitleCase to be false, got %v", Config.TitleCase)
	}
	if Config.Debug {
		t.Errorf("Expected default Debug to be false, got %v", Config.Debug)
	}

	// Sanity check that we can set a value
	Config.TitleCase = true
	defer func() { Config.TitleCase = false }()
	if !Config.TitleCase {
		t.Errorf("Expected TitleCase to be true after setting, got %v", Config.TitleCase)
	}
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
	for _, testcase := range testCases {
		got := toTitleCase(testcase.input)
		if got != testcase.want {
			t.Errorf("toTitleCase(%q) = %q; want %q", testcase.input, got, testcase.want)
		}
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
	for _, testcase := range testCases {
		got, _, err := numberPrefix(testcase.input)
		if got != testcase.want {
			t.Errorf("numberPrefix(%q) = %d; want %d", testcase.input, got, testcase.want)
		}
		if testcase.want == -1 {
			// Then there should be an error...
			if err == nil {
				t.Errorf("numberPrefix(%q) expected error, got %d", testcase.input, got)
			}
		}
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
	for _, testcase := range testCases {
		got := hasTrackNumber(testcase.input)
		if got != testcase.want {
			t.Errorf("hasTrackNumber(%q) = %v; want %v", testcase.input, got, testcase.want)
		}
	}
}

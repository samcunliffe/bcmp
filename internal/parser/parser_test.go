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

func TestIsCoverArtFile(t *testing.T) {
	var testCases = []struct {
		input string
		want  bool
	}{
		{"cover.jpg", true},
		{"COVER.PNG", true},
		{"folder.jpg", true},
		{"FOLDER.PNG", true},
		{"not_cover.jpg", false},
		{"flibble", false},
		{"cover.jpeg", false},
		{"some_other_file.mp3", false},
		{"some_other_file.flac", false},
	}
	for _, testcase := range testCases {
		got := IsCoverArtFile(testcase.input)
		if got != testcase.want {
			t.Errorf("IsCoverArtFile(%q) = %v; want %v", testcase.input, got, testcase.want)
		}
	}
}

func TestIsValidMusicFile(t *testing.T) {
	var testCases = []struct {
		input string
		want  bool
	}{
		{"Crypta - Shades of Sorrow - 01 The Aftermath.txt", false}, // Invalid file extension
		{"Crypta - Shades of Sorrow - 01 The Aftermath", false},     // No file extension
		{"Crypta - Shades of Sorrow - 01 The Aftermath.mp3", true},  // Valid files
		{"Crypta - Shades of Sorrow - 01 The Aftermath.flac", true},
	}
	for _, testcase := range testCases {
		if IsValidMusicFile(testcase.input) != testcase.want {
			t.Errorf("IsValidMusicFile(%q) = %v", testcase.input, !testcase.want)
		}
	}
}

func TestToTitleCase(t *testing.T) {
	var testCases = []struct {
		input string
		want  string
	}{
		{"SYLOSIS", "Sylosis"},                                   // Made up word
		{"ROSALÍA - SAOKO", "Rosalía - Saoko"},                   // With accent
		{"IVAR BJØRNSON", "Ivar Bjørnson"},                       // Scandinavian letter
		{"THE OTHER SIDE OF ANGER", "The Other Side of Anger"},   // 'Small' word from UPPER
		{"bleed from within", "Bleed from Within"},               // 'Small' word from lower
		{"THE BEGINNING OF THE END", "The Beginning of the End"}, // Multiple 'small' words
		{"the end of all we know", "The End of All We Know"},     // Multiple 'small' words
		{"BABYMETAL - メギツネ", "Babymetal - メギツネ"},                 // Katakana
		{"BABYMETAL - SONG 4 (4の歌)", "Babymetal - Song 4 (4の歌)"}, // Hiragana and Kanji
	}
	for _, testcase := range testCases {
		got := toTitleCase(testcase.input)
		if got != testcase.want {
			t.Errorf("toTitleCase(%q) = %q; want %q", testcase.input, got, testcase.want)
		}
	}
}

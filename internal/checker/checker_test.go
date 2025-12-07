package checker

import "testing"

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

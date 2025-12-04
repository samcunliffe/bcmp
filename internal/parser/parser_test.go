package parser

import "testing"

// func TestSplitOnHyphen(t *testing.T)
// func TestRemoveParenthesis(t *testing.T)

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
	var testCases = []string{
		"Crypta - Shades of Sorrow - 01 The Aftermath.txt", // Invalid file extension
		"Crypta - Shades of Sorrow - 01 The Aftermath",     // No file extension
	}
	for _, testcase := range testCases {
		if IsValidMusicFile(testcase) {
			t.Errorf("IsValidMusicFile(%q) = true; want false", testcase)
		}
	}
}

package parser

import "testing"

func TestExtractNumberPrefix(t *testing.T) {
	var testCases = []struct {
		input_track    string
		want_number    int
		want_trackname string
	}{
		{"01 The Aftermath", 1, "The Aftermath"},
		{"02 Dark Clouds", 2, "Dark Clouds"},
		{"06 The Other Side of Anger", 6, "The Other Side of Anger"},
		{"12 Lord of Ruins", 12, "Lord of Ruins"},
	}
	for _, testcase := range testCases {
		n, s, err := extractNumberPrefix(testcase.input_track)
		if err != nil {
			t.Errorf("extractNumberPrefix returned error for %q: %v", testcase.input_track, err)
		}
		if n != testcase.want_number || s != testcase.want_trackname {
			t.Errorf("extractNumberPrefix failed for %q: got (%d, %q), want (%d, %q)", testcase.input_track, n, s, testcase.want_number, testcase.want_trackname)
		}
	}
}

func TestParseMusicFileName(t *testing.T) {
	var testCases = []struct {
		inputFilename string
		wantNumber    int
		wantTitle     string
		wantTrack     string
		wantSuffix    string
	}{
		{"Crypta - Shades of Sorrow - 01 The Aftermath.flac", 1, "The Aftermath", "01 The Aftermath", ".flac"},
		{"Crypta - Shades of Sorrow - 02 Dark Clouds.mp3", 2, "Dark Clouds", "02 Dark Clouds", ".mp3"},
		{"Crypta - Shades of Sorrow - 06 The Other Side of Anger.ogg", 6, "The Other Side of Anger", "06 The Other Side of Anger", ".ogg"},
	}
	for _, testcase := range testCases {
		_, got, err := ParseMusicFileName(testcase.inputFilename)
		if err != nil {
			t.Errorf("ParseMusicFileName(%q) returned error: %v", testcase.inputFilename, err)
		}
		if got.Number != testcase.wantNumber {
			t.Errorf("ParseMusicFileName(%q) Number = %d; want %d", testcase.inputFilename, got.Number, testcase.wantNumber)
		}
		if got.Title != testcase.wantTitle {
			t.Errorf("ParseMusicFileName(%q) Title = %q; want %q", testcase.inputFilename, got.Title, testcase.wantTitle)
		}
		if got.FullTrack != testcase.wantTrack {
			t.Errorf("ParseMusicFileName(%q) FullTrack = %q; want %q", testcase.inputFilename, got.FullTrack, testcase.wantTrack)
		}
		if got.FileType != testcase.wantSuffix {
			t.Errorf("ParseMusicFileName(%q) FileType = %q; want %q", testcase.inputFilename, got.FileType, testcase.wantSuffix)
		}
	}
}

func TestParseMusicFilenameErrors(t *testing.T) {
	var errorCases = []string{
		"Crypta - Shades of Sorrow - The Aftermath.flac",          // Missing track number
		"Crypta - Shades of Sorrow - 01 The Aftermath.txt",        // Invalid file extension
		"Crypta - Shades of Sorrow - 01 The Aftermath",            // No file extension
		"Crypta - Shades of Sorrow - 00 The Aftermath.flac",       // Track zero
		"Crypta - Shades of Sorrow - Track One The Aftermath.mp3", // Non-numeric track number
		"Crypta Shades of Sorrow 01 The Aftermath.flac",           // No hyphens
		" - 01 The Aftermath.flac",                                // Missing artist and album
		"Just the Song Title.flac",                                // Missing artist, album, track number
	}
	for _, filename := range errorCases {
		_, _, err := ParseMusicFileName(filename)
		if err == nil {
			t.Errorf("ParseMusicFileName(%q) expected to return error, but got nil", filename)
		}
	}
}

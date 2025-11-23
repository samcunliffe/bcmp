package fileparser

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

func TestParseMusicFilename(t *testing.T) {
	var testCases = []struct {
		input_filename string
		want_number    int
		want_title     string
		want_track     string
	}{
		{"Crypta - Shades of Sorrow - 01 The Aftermath.flac", 1, "The Aftermath", "01 The Aftermath"},
		{"Crypta - Shades of Sorrow - 02 Dark Clouds.mp3", 2, "Dark Clouds", "02 Dark Clouds"},
		{"Crypta - Shades of Sorrow - 06 The Other Side of Anger.ogg", 6, "The Other Side of Anger", "06 The Other Side of Anger"},
	}
	for _, testcase := range testCases {
		got, err := ParseMusicFileName(testcase.input_filename)
		if err != nil {
			t.Errorf("ParseMusicFileName(%q) returned error: %v", testcase.input_filename, err)
		}
		if got.Number != testcase.want_number || got.Title != testcase.want_title || got.FullTrack != testcase.want_track {
			t.Errorf("ParseMusicFileName(%q) = number: %d, title: %q, fulltrack: %q; want number: %d, title: %q, fulltrack: %q",
				testcase.input_filename, got.Number, got.Title, got.FullTrack, testcase.want_number, testcase.want_title, testcase.want_track)
		}
	}
}

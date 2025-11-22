package fileparser

import "testing"

var testCases = []struct {
	input_filename string
	want_artist    string
	want_album     string
}{
	{"Architects - For Those That Wish To Exist.zip", "Architects", "For Those That Wish To Exist"},
	{"Bloodywood - Nu Delhi.zip", "Bloodywood", "Nu Delhi"},
	{"Crypta - Shades of Sorrow (pre-order).zip", "Crypta", "Shades of Sorrow"},
	{"Crypta - Shades of Sorrow.zip", "Crypta", "Shades of Sorrow"},
	{"Enslaved - Heimdal.zip", "Enslaved", "Heimdal"},
	{"Immovable Stone - Sylosis.zip", "Immovable Stone", "Sylosis"},
	{"Lokust - Infidel.zip", "Lokust", "Infidel"},
	{"Malevolence - Malicious Intent.zip", "Malevolence", "Malicious Intent"},
	{"Nervosa - Jailbreak (pre-order).zip", "Nervosa", "Jailbreak"},
	{"Nervosa - Jailbreak.zip", "Nervosa", "Jailbreak"},
	{"Orbit Culture - Death Above Life (24-bit HD audio) (pre-order).zip", "Orbit Culture", "Death Above Life"},
	{"Orbit Culture - Rasen.zip", "Orbit Culture", "Rasen"},
	{"Pallbearer - Mind Burns Alive (pre-order).zip", "Pallbearer", "Mind Burns Alive"},
}

func TestParseZipFileName(t *testing.T) {
	for _, testcase := range testCases {
		got, err := ParseZipFileName(testcase.input_filename)
		if err != nil {
			t.Errorf("ParseZipArchiveFileName(%q) returned error: %v", testcase.input_filename, err)
		}
		if got.Artist != testcase.want_artist || got.Title != testcase.want_album {
			t.Errorf("ParseZipArchiveFileName(%q) = artist: %q, album: %q; want artist: %q, album: %q",
				testcase.input_filename, got.Artist, got.Title, testcase.want_artist, testcase.want_album)
		}
	}
}

var errorTestCases = []string{
	"NoHyphenHere.zip",
	"MissingSuffix",
	"Also Missing Hyphen.zip",
}

func TestParseZipFileNameErrors(t *testing.T) {
	for _, filename := range errorTestCases {
		_, err := ParseZipFileName(filename)
		if err == nil {
			t.Errorf("ParseZipFileName(%q) expected error but got nil", filename)
		}
	}
}

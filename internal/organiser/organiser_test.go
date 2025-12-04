package organiser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/samcunliffe/bcmp/internal/parser"
)

func TestDefaultDestination(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	want := filepath.Join(home, "Music")
	got := DefaultDestination()
	if got != want {
		t.Errorf("DefaultDestination() = %q; want %q", got, want)
	}
}

func TestDetermineDefaultDestinationNoHome(t *testing.T) {
	t.Setenv("HOME", "")

	want := filepath.Join(".", "Music")
	got := DefaultDestination()
	if got != want {
		t.Errorf("DefaultDestination() = %q; want %q", got, want)
	}
}

func TestCreateDestination(t *testing.T) {
	base := t.TempDir()

	testCases := []struct {
		album parser.Album
		want  string
	}{
		{
			album: parser.Album{Artist: "Crypta", Title: "Shades of Sorrow"},
			want:  filepath.Join(base, "Crypta", "Shades of Sorrow"),
		},
		{
			album: parser.Album{Artist: "Orbit Culture", Title: "Death Above Life"},
			want:  filepath.Join(base, "Orbit Culture", "Death Above Life"),
		},
	}

	for _, testcase := range testCases {
		got, err := CreateDestination(testcase.album, base)
		if err != nil {
			t.Errorf("CreateDestination(%v, %q) returned error: %v", testcase.album, base, err)
			continue
		}
		if got != testcase.want {
			t.Errorf("CreateDestination(%v, %q) = %q; want %q", testcase.album, base, got, testcase.want)
		}
		if _, err := os.Stat(got); os.IsNotExist(err) {
			t.Errorf("CreateDestination did not create directory %q", got)
		}
	}
}

func TestCreateDestinationNonExistentBase(t *testing.T) {
	base := filepath.Join(t.TempDir(), "some_nonexistent_subdir")

	album := parser.Album{Artist: "Crypta", Title: "Shades of Sorrow"}
	got, err := CreateDestination(album, base)
	if err != nil {
		t.Errorf("CreateDestination(%v, %q) returned error: %v", album, base, err)
	}
	want := filepath.Join(base, "Crypta", "Shades of Sorrow")
	if got != want {
		t.Errorf("CreateDestination(%v, %q) = %q; want %q", album, base, got, want)
	}
	// TODO: Might actually be better to do something else than write a warning to stdout...
	// Investigate logging and a -v/--verbose flag?
}

func TestCheckFileDirectory(t *testing.T) {
	err := CheckFile("testdata/directory")
	if err == nil {
		t.Errorf("CheckFile on directory did not return error")
	}
}

func TestCheckFileNonExistent(t *testing.T) {
	err := CheckFile("testdata/nonexistent.zip")
	if err == nil {
		t.Errorf("CheckFile on nonexistent file did not return error")
	}
}

func TestCheckFileEmptyFile(t *testing.T) {
	err := CheckFile("testdata/emptyfile")
	if err == nil {
		t.Errorf("CheckFile on empty file did not return error")
	}
}

func TestCheckFileValidZipFile(t *testing.T) {
	err := CheckFile("testdata/validfile.zip")
	if err != nil {
		t.Errorf("CheckFile on valid file returned error: %v", err)
	}
}

func TestCheckFileValidMusicFile(t *testing.T) {
	err := CheckFile("testdata/ding.flac")
	if err != nil {
		t.Errorf("CheckFile on valid music file returned error: %v", err)
	}
}

func TestTrackDestinationPath(t *testing.T) {
	track := parser.Track{
		Number:    1,
		Title:     "First Track",
		FullTrack: "01 First Track",
		FileType:  ".flac",
	}
	want := "01 First Track.flac"
	if TrackDestinationPath(track, "./") != want {
		t.Errorf("TrackDestinationPath() = %q; want %q", TrackDestinationPath(track, "./"), want)
	}
}

func TestTrackDestinationPathWeirdPaths(t *testing.T) {
	track := parser.Track{
		Number:    1,
		Title:     "First Track",
		FullTrack: "01 First Track",
		FileType:  ".flac",
	}
	testCases := []struct {
		input string
		want  string
	}{
		{"/a/b/c/../../", "/a/01 First Track.flac"},
		{"a/b/c/../..", "a/01 First Track.flac"},
		{"./", "01 First Track.flac"},
		{"./music/", "music/01 First Track.flac"},
	}

	for _, testcase := range testCases {
		got := TrackDestinationPath(track, testcase.input)
		if got != testcase.want {
			t.Errorf("TrackDestinationPath(_, %q) = %q; want %q", testcase.input, got, testcase.want)
		}
	}
}

package organiser

import (
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

func TestTrackDestination(t *testing.T) {
	track := parser.Track{
		Number:    1,
		Title:     "First Track",
		FullTrack: "01 First Track",
		FileType:  ".flac",
	}
	want := "01 First Track.flac"
	if TrackDestination(track, "./") != want {
		t.Errorf("TrackDestinationPath() = %q; want %q", TrackDestination(track, "./"), want)
	}
}

func TestTrackDestinationWeirdPaths(t *testing.T) {
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
		got := TrackDestination(track, testcase.input)
		if got != testcase.want {
			t.Errorf("TrackDestinationPath(_, %q) = %q; want %q", testcase.input, got, testcase.want)
		}
	}
}

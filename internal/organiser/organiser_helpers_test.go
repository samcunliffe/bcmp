package organiser

import (
	"path/filepath"
	"testing"

	"github.com/samcunliffe/bcmp/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestDefaultDestination(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	want := filepath.Join(home, "Music")
	got := DefaultDestination()
	assert.Equal(t, want, got, "DefaultDestination() = %q; want %q", got, want)
}

func TestDetermineDefaultDestinationNoHome(t *testing.T) {
	t.Setenv("HOME", "")

	want := filepath.Join(".", "Music")
	got := DefaultDestination()

	assert.Equal(t, want, got, "DefaultDestination() = %q; want %q", got, want)
}

func TestTrackDestination(t *testing.T) {
	track := parser.Track{
		Number:    1,
		Title:     "First Track",
		FullTrack: "01 First Track",
		FileType:  ".flac",
	}
	want := "01 First Track.flac"
	got := TrackDestination(track, "./")
	assert.Equal(t, want, got, "TrackDestinationPath() = %q; want %q", got, want)
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

	for _, tc := range testCases {
		got := TrackDestination(track, tc.input)
		assert.Equal(t, tc.want, got, "TrackDestinationPath(_, %q) = %q; want %q", tc.input, got, tc.want)
	}
}

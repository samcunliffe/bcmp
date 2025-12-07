package checker

import (
	"strings"

	p "github.com/samcunliffe/bcmp/internal/parser"
)

var coverArtFilenames = []string{"cover.jpg", "cover.png", "folder.jpg", "folder.png"}
var validMusicFiles = []string{".flac", ".mp3", ".ogg"}

// Is the file a zip archive?
func IsZipFile(s string) bool {
	return p.Extension(s) == ".zip"
}

// Is the file a known cover art file?
func IsCoverArtFile(name string) bool {
	name = strings.ToLower(name)
	for _, coverName := range coverArtFilenames {
		if name == coverName {
			return true
		}
	}
	return false
}

// Is the file a known (supported) music file?
func IsValidMusicFile(name string) bool {
	ext := p.Extension(name)

	for _, fileType := range validMusicFiles {
		if ext == fileType {
			return true
		}
	}
	return false
}

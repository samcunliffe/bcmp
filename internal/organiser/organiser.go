package organiser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/samcunliffe/bcmp/internal/parser"
)

// Determine default destination for music files
//
// Typically $HOME/Music. If $HOME cannot be determined, use current directory.
// Note: this does not check that the directory exists.
func DefaultDestination() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, "Music")
}

// Create the directory structure for the album under base
//
// e.g. if base is $HOME/Music, create $HOME/Music/Artist/Album
func CreateDestination(album parser.Album, base string) (string, error) {
	if _, err := os.Stat(base); os.IsNotExist(err) {
		// It's noteworthy if, e.g. the user doesn't have a Music folder
		fmt.Printf("Warning: base destination path %s does not exist. Will create it.\n", base)
	}
	destination := filepath.Join(base, album.Artist, album.Title)
	return destination, os.MkdirAll(destination, os.ModePerm)
}

// Ensure a zip file or music file exists and is not a directory
func CheckFile(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("error accessing zip file: %v", err)
	}
	if fi.IsDir() || fi.Size() == 0 {
		return fmt.Errorf("the zip file: %v is not valid", fi.Name())
	}
	return nil
}

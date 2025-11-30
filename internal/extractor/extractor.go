package extractor

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/samcunliffe/bcmptidy/internal/parser"
)

func trackDestinationPath(t parser.Track, destination string) string {
	filename := fmt.Sprintf("%s%s", t.FullTrack, t.FileType)
	return filepath.Join(destination, filename)
}

func processTrack(f *zip.File, destination string) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("impossible to open file in archive: %s", err)
	}
	defer rc.Close()

	track, err := parser.ParseMusicFileName(f.Name)
	if err != nil {
		return fmt.Errorf("error parsing music file name %s: %v", f.Name, err)
	}

	// Actually extract and write the file
	fd, err := os.Create(trackDestinationPath(track, destination))
	if err != nil {
		return fmt.Errorf("impossible to create destination file: %s", err)
	}
	defer fd.Close()
	_, err = io.Copy(fd, rc)
	if err != nil {
		return fmt.Errorf("impossible to copy file contents to destination: %s", err)
	}

	fmt.Printf("Extracted track %d: %s\n", track.Number, track.Title)
	return nil
}

func ExtractAndRename(zipPath, destination string) error {
	// Open zip archive for reading.
	rc, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("impossible to open zip reader: %s", err)
	}
	defer rc.Close()

	// Iterate through the files in the archive,
	for _, f := range rc.File {
		fmt.Printf("Unzipping %s:\n", f.Name)

		// Ignore directories
		if f.FileInfo().IsDir() {
			fmt.Printf("Skipping directory: %s\n", f.Name)
			continue
		}

		// Ignore cover art for now
		if parser.IsCoverArtFile(f.Name) {
			fmt.Printf("Skipping: %s\n", f.Name)
			continue
		}

		err := processTrack(f, destination)
		if err != nil {
			return fmt.Errorf("error processing track %s: %v", f.Name, err)
		}
	}
	return nil
}

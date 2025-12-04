package extractor

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	o "github.com/samcunliffe/bcmp/internal/organiser"
	p "github.com/samcunliffe/bcmp/internal/parser"
)

func processTrack(f *zip.File, destination string) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("impossible to open file in archive: %s", err)
	}
	defer rc.Close()

	_, track, err := p.ParseMusicFileName(f.Name)
	if err != nil {
		return fmt.Errorf("error parsing music file name %s: %v", f.Name, err)
	}

	// Actually extract and write the file
	fd, err := os.Create(o.TrackDestination(track, destination))
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

		// Any directories - this is likely not a bandcamp file
		if f.FileInfo().IsDir() {
			return fmt.Errorf("archive contains a directory, not a valid bandcamp zip")
		}

		// Ignore cover art for now
		if p.IsCoverArtFile(filepath.Base(f.Name)) {
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

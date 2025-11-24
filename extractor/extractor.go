package extractor

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
)

func ExtractAndRename(zipPath, destinationPath string) error {
	return nil
	// Open a zip archive for reading.
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Fatalf("impossible to open zip reader: %s", err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	for k, f := range r.File {
		fmt.Printf("Unzipping %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatalf("impossible to open file n°%d in archine: %s", k, err)
		}
		defer rc.Close()
		// define the new file path
		newFilePath := fmt.Sprintf("uncompressed/%s", f.Name)

		// CASE 1 : we have a directory
		if f.FileInfo().IsDir() {
			// if we have a directory we have to create it
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				log.Fatalf("impossible to MkdirAll: %s", err)
			}
			// we can go to next iteration
			continue
		}
	}
	return nil
}

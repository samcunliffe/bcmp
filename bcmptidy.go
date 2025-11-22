package main

import (
	"fmt"
	"os"

	"github.com/samcunliffe/bcmptidy/fileparser"
)

func main() {

	input := os.Args[1:]
	if len(input) != 1 {
		fmt.Println("Usage: bcmptidy <zipfile>")
		os.Exit(1)
	}
	fmt.Println("Input file: %s", input[0])
	album, err := fileparser.ParseZipFileName(input[1])
	if err != nil {
		os.Exit(1)
	}

	// outputDirectory := fmt.Sprintf("%s/%s/%s", MUSICLIBRARYPATH, album.Artist, album.Title)
	// mkdir(outputDirectory)
	// zip.Extract(input)
	fmt.Println("Hello, %s", album.Artist)
}

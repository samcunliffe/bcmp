package parser

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Config = struct {
	TitleCase bool
	Debug     bool
}{}

type Album struct {
	Artist string
	Title  string
	// Tracks []Track
}

type Track struct {
	Number    int
	Title     string
	FullTrack string
	FileType  string
}

var coverArtFilenames = []string{"cover.jpg", "cover.png", "folder.jpg", "folder.png"}
var validMusicFiles = []string{".flac", ".mp3", ".ogg"}

// Get the filename and its extenstion in lower case
func splitNameAndExtension(s string) (string, string) {
	extension := filepath.Ext(s)
	name := strings.TrimSuffix(s, extension)
	return name, strings.ToLower(extension)
}

// Get only the extension of a filename in lower case
func extension(s string) string {
	_, ext := splitNameAndExtension(s)
	return ext
}

// Is the file a zip archive?
func isZipFile(s string) bool {
	return extension(s) == ".zip"
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
	ext := extension(name)

	for _, fileType := range validMusicFiles {
		if ext == fileType {
			return true
		}
	}
	return false
}

func toTitleCase(s string) string {
	smallWords := " a an and as at but by for from if in nor of on or the to v von vs "
	words := strings.Split(s, " ")
	for i, word := range words {
		lowerPaddedWord := " " + strings.ToLower(word) + " "
		if i != 0 && strings.Contains(smallWords, lowerPaddedWord) {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = cases.Title(language.English).String(word)
		}
	}
	return strings.Join(words, " ")
}

func splitOnHyphen(s string) (string, string) {
	ss := strings.SplitN(s, " - ", 2)
	return ss[0], ss[1]
}

func removeParenthesis(s string) string {
	re := regexp.MustCompile(`\s*\(.*?\)`)
	return strings.TrimSpace(re.ReplaceAllString(s, ""))
}

func numberPrefix(s string) (int, string, error) {
	re := regexp.MustCompile(`^(\d+)\s*(.*)`)

	// Expect "XX Track Name", "XX", "Track Name"
	matches := re.FindStringSubmatch(s)
	if len(matches) != 3 {
		return -1, s, fmt.Errorf("error in regex match for track number extraction")
	}

	// Convert track number from string
	number, err := strconv.Atoi(matches[1])
	if err != nil || number < 1 {
		return -1, s, fmt.Errorf("failed to convert track number '%s' to int: %v", matches[1], err)
	}
	return number, strings.TrimSpace(matches[2]), nil
}

func ParseZipFileName(name string) (Album, error) {
	if !isZipFile(name) {
		return Album{}, fmt.Errorf("file is not a zip archive")
	}
	name, _ = splitNameAndExtension(name)

	// Split into artist and album
	if !strings.Contains(name, " - ") {
		return Album{}, fmt.Errorf("filename does not contain ' - ' separator")
	}
	if strings.Count(name, " - ") != 1 {
		return Album{}, fmt.Errorf("expected only one ' - ' separator: '%s'", name)
	}
	artist, album := splitOnHyphen(name)
	if Config.TitleCase {
		artist = toTitleCase(artist)
		album = toTitleCase(album)
	}
	return Album{Artist: artist, Title: removeParenthesis(album)}, nil
}

func ParseMusicFileName(name string) (Album, Track, error) {
	if !IsValidMusicFile(name) {
		return Album{}, Track{}, fmt.Errorf("file is not a valid music file")
	}
	name, fileType := splitNameAndExtension(name)

	// Should be two hyphens 'Artist - Album - XX Title'
	if !strings.Contains(name, " - ") {
		return Album{}, Track{}, fmt.Errorf("filename does not contain ' - ' separator")
	}
	if strings.Count(name, " - ") != 2 {
		return Album{}, Track{}, fmt.Errorf("expected two ' - ' separators: '%s'", name)
	}

	// Split into artist, album, number, track title
	artist, albumAndTrack := splitOnHyphen(name)
	albumTitle, fullTrack := splitOnHyphen(albumAndTrack)

	// Convert to title case if configured
	if Config.TitleCase {
		artist = toTitleCase(artist)
		albumTitle = toTitleCase(albumTitle)
	}

	album := Album{Artist: artist, Title: removeParenthesis(albumTitle)}

	number, songTitle, err := numberPrefix(fullTrack)
	if err != nil {
		return album, Track{Title: name}, fmt.Errorf("failed to extract track number and title: %v", err)
	}

	if Config.TitleCase {
		songTitle = toTitleCase(songTitle)
		fullTrack = fmt.Sprintf("%02d %s", number, songTitle)
	}

	track := Track{
		Number:    number,
		Title:     songTitle,
		FullTrack: fullTrack,
		FileType:  fileType,
	}

	return album, track, nil
}

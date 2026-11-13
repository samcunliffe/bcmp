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

// Get the filename and its extension in lower case
func splitNameAndExtension(s string) (string, string) {
	extension := filepath.Ext(s)
	name := strings.TrimSuffix(s, extension)
	return name, strings.ToLower(extension)
}

// Get only the extension of a filename in lower case
func Extension(s string) string {
	_, ext := splitNameAndExtension(s)
	return ext
}

func isRomanNumeral(s string) bool {
	pattern := regexp.MustCompile(`^M*(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	return s != "" && pattern.MatchString(strings.ToUpper(s))
}

func toTitleCase(s string) string {
	smallWords := " a an and as at but by for from if in nor of on or the to von vs "
	words := strings.Split(s, " ")

	// The first word in the string might be a track number:
	// "01 The Beginning of the End"
	// -> ignore "01" -> the first word is "The", capitalize
	firstWord := 0
	if hasTrackNumber(s) {
		firstWord = 1
	}

	for i, word := range words {
		lowerPadded := " " + strings.ToLower(word) + " "
		if i > firstWord && strings.Contains(smallWords, lowerPadded) {
			// If the word is a 'small' word that doesn't start the title, make it lower case
			lower := strings.ToLower(word)
			words[i] = lower
		} else if isRomanNumeral(word) {
			// If the word is a roman numeral, make it upper case
			upper := strings.ToUpper(word)
			words[i] = upper
		} else {
			// Otherwise, capitalize the word
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

func hasTrackNumber(s string) bool {
	re := regexp.MustCompile(`^(\d+)\s+.*`)
	return re.MatchString(s)
}

func ParseZipFileName(name string) (Album, error) {
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
	name, fileType := splitNameAndExtension(name)

	// Should be one or two hyphens either a single: 'Artist - Title' or an album track: 'Artist - Album - XX Title'
	nSeparators := strings.Count(name, " - ")
	if nSeparators < 1 {
		return Album{}, Track{}, fmt.Errorf("filename does not contain ' - ' separator")
	}
	if nSeparators != 1 && nSeparators != 2 {
		return Album{}, Track{}, fmt.Errorf("expected one or two ' - ' separators: '%s'", name)
	}

	// Split into artist, album, number, track title
	artist, albumAndTrack := splitOnHyphen(name)
	albumTitle, fullTrack := splitOnHyphen(albumAndTrack)
	if nSeparators == 1 {
		// If there's only one hyphen, then the album title is the same as the track title
		albumTitle = fullTrack
	}

	// Convert to title case if configured
	if Config.TitleCase {
		artist = toTitleCase(artist)
		albumTitle = toTitleCase(albumTitle)
		fullTrack = toTitleCase(fullTrack)
	}

	album := Album{Artist: artist, Title: removeParenthesis(albumTitle)}

	number, songTitle, err := numberPrefix(fullTrack)
	if err != nil && nSeparators == 2 {
		return album, Track{Title: name}, fmt.Errorf("failed to extract track number and title: %v", err)
	}
	if err != nil && nSeparators == 1 {
		// If there's only one hyphen, then the track title is the same as the album title
		songTitle = fullTrack
		number = 0
	}

	track := Track{
		Number:    number,
		Title:     songTitle,
		FullTrack: fullTrack,
		FileType:  fileType,
	}

	return album, track, nil
}

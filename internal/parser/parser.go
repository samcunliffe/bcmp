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

// Global configuration for the parser package.
var Config = struct {
	ApplyTitleCase bool
}{}

// Album represents an album of one or more tracks (including LP, single with B-side etc).
// An empty title is allowed for singles (where there is no album title).
type Album struct {
	Artist string
	Title  string
}

// Track represents one track in an album, including its number, title, and file extension.
// A number of zero is allowed for singles (where there is no track number).
type Track struct {
	Number    int
	Title     string
	FullTrack string
	Extension string
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

// Is the string a valid Roman numeral?
func isRomanNumeral(s string) bool {
	pattern := regexp.MustCompile(`^M*(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	return s != "" && pattern.MatchString(strings.ToUpper(s))
}

// Convert the string to title case with special handling for small words and Roman numerals.
// Roman numerals are treated at a preference and are uppercase.
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

// Split the string on the first hyphen, returning the two parts.
// If there is no hyphen, the first part is empty and the second part is the whole string.
func splitOnHyphen(s string) (string, string) {
	ss := strings.SplitN(s, " - ", 2)
	switch len(ss) {
	case 0:
		return "", ""
	case 1:
		return "", ss[0]
	}
	return ss[0], ss[1]
}

// Remove any text within parentheses from the string.
func removeParenthesis(s string) string {
	re := regexp.MustCompile(`\s*\(.*?\)`)
	return strings.TrimSpace(re.ReplaceAllString(s, ""))
}

// Get the track number and title from the full track string '01 Track Name'.
func numberPrefix(s string) (int, string, error) {
	re := regexp.MustCompile(`^(\d+)\s*(.*)`)

	matches := re.FindStringSubmatch(s)
	if len(matches) != 3 {
		return -1, s, fmt.Errorf("error in regex match for track number extraction")
	}

	number, err := strconv.Atoi(matches[1])
	if err != nil || number < 1 {
		return -1, s, fmt.Errorf("failed to convert track number '%s' to int: %v", matches[1], err)
	}
	return number, strings.TrimSpace(matches[2]), nil
}

// Check if the string has a track number prefix (e.g., '01 Track Name').
func hasTrackNumber(s string) bool {
	re := regexp.MustCompile(`^\d+(\s+.*)?$`)
	return re.MatchString(s)
}

// Check the album and track for sanity, returning an error if any field is invalid.
func check(album Album, track Track, isSingle bool) error {
	if album.Artist == "" {
		return fmt.Errorf("album artist is empty")
	}
	if track.FullTrack == "" || track.Title == "" {
		return fmt.Errorf("track title is empty")
	}
	if track.Extension == "" {
		return fmt.Errorf("track extension is empty")
	}
	if !isSingle && track.Number < 1 {
		return fmt.Errorf("track number is invalid: %d", track.Number)
	}
	if album.Title == "" && !isSingle {
		return fmt.Errorf("album title is empty")
	}
	return nil
}

// Parse the zip file name into an Album struct, returning an error if the name is invalid.
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
	if Config.ApplyTitleCase {
		artist = toTitleCase(artist)
		album = toTitleCase(album)
	}
	return Album{Artist: artist, Title: removeParenthesis(album)}, nil
}

// Parse the music file name into an Album and Track struct, returning an error if the name is invalid.
func ParseMusicFileName(name string) (Album, Track, error) {
	name, extension := splitNameAndExtension(name)

	// Should be one or two hyphens either a single: 'Artist - Title' or an
	// album track: 'Artist - Album - XX Title'
	nSeparators := strings.Count(name, " - ")
	if nSeparators < 1 {
		return Album{}, Track{}, fmt.Errorf("filename does not contain ' - ' separator")
	}
	if nSeparators != 1 && nSeparators != 2 {
		return Album{}, Track{}, fmt.Errorf("expected one or two ' - ' separators: '%s'", name)
	}

	// Only one hyphen means it's probably a single (without an album)
	isSingle := nSeparators == 1

	// Split into artist, album, number, track title
	artist, albumAndTrack := splitOnHyphen(name)
	albumTitle, fullTrack := splitOnHyphen(albumAndTrack)

	// Convert to title case if configured
	if Config.ApplyTitleCase {
		artist = toTitleCase(artist)
		albumTitle = toTitleCase(albumTitle)
		fullTrack = toTitleCase(fullTrack)
	}

	album := Album{Artist: artist, Title: removeParenthesis(albumTitle)}

	if isSingle && !hasTrackNumber(fullTrack) {
		singleTrack := Track{
			Number:    0,
			Title:     fullTrack,
			FullTrack: fullTrack,
			Extension: extension,
		}
		err := check(album, singleTrack, isSingle)
		return album, singleTrack, err
	}

	// Try to extract the track trackNumber. We've dealt with the single case so if
	// it fails something has gone wrong so return an error.
	trackNumber, songTitle, err := numberPrefix(fullTrack)
	if err != nil {
		return album, Track{}, fmt.Errorf("failed to extract track number and title: %v", err)
	}

	albumTrack := Track{
		Number:    trackNumber,
		Title:     songTitle,
		FullTrack: fullTrack,
		Extension: extension,
	}
	err = check(album, albumTrack, isSingle)
	return album, albumTrack, err
}

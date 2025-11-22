package datamodel

type Album struct {
	Artist string
	Title  string
	Tracks []Track
}

type Track struct {
	TrackNum int
	Title    string
}

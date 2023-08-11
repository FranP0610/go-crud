package domain

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model

	Name           string
	Country        string
	FoundationYear string
	Albums         []Album
}

type Album struct {
	gorm.Model

	Name        string
	ArtistID    uint
	ReleaseYear uint
	Tracks      []Track
}

type Track struct {
	gorm.Model

	Name     string
	AlbumID  uint
	Duration float32
}

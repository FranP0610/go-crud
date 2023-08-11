package domain

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model

	Name           string `gorm:"unique;not null" json:"name"`
	Country        string `gorm:"not null" json:"country"`
	FoundationYear uint   `gorm:"not null" json:"foundation_year"`
	Albums         []Album
}

type Album struct {
	gorm.Model

	Name        string `gorm:"unique;not null" json:"name"`
	ArtistID    uint   `gorm:"not null" json:"artist_id"`
	ReleaseYear uint   `gorm:"not null" json:"release_year"`
	Tracks      []Track
}

type Track struct {
	gorm.Model

	Name     string  `gorm:"unique;not null" json:"name"`
	AlbumID  uint    `json:"album_id"`
	Duration float32 `gorm:"not null" json:"duration"`
}

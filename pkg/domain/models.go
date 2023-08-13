package domain

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	// `gorm:"unique;not null;default:null" json:"name"` -> not null and default:null are used together
	// to prevent empty rows
	Name           string `gorm:"unique;not null;default:null" json:"name"`
	Country        string `gorm:"not null;default:null" json:"country"`
	FoundationYear uint   `gorm:"not null;default:null" json:"foundation_year"`
	Albums         []Album
}

type Album struct {
	gorm.Model

	Name        string `gorm:"unique;not null;default:null" json:"name"`
	ArtistID    uint   `gorm:"not null;default:null" json:"artist_id"`
	ReleaseYear uint   `gorm:"not null;default:null" json:"release_year"`
	Tracks      []Track
}

type Track struct {
	gorm.Model

	Name     string  `gorm:"unique;not null;default:null" json:"name"`
	AlbumID  uint    `json:"album_id"`
	Duration float32 `gorm:"not null" json:"duration"`
}

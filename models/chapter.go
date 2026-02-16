package models

type Chapter struct {
	ID      uint   `gorm:"primaryKey"`
	MangaID string
	Title string
	Number int
}
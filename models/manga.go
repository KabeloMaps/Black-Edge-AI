package models

type Manga struct {
	ID     uint    `gorm:"primaryKey"`
	Title  string  `gorm:"index"`
	Description string
	Author string
	Status string
	CoverImage string 
}
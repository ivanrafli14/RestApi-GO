package model

type Song struct {
	ID string `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:VARCHAR(300)" json:"title" binding:"required"`
	Year int64 `gorm:"type:INTEGER" json:"year" binding:"required"`
	Genre string `gorm:"type:VARCHAR(300)" json:"genre" binding:"required"`
	Performer string `gorm:"type:VARCHAR(300)" json:"performer" binding:"required"`
	Duration int64 `gorm:"type:INTEGER" json:"duration"`
	AlbumId string `gorm:"type:VARCHAR(300)" json:"albumId"`
}
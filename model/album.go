package model

type Album struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:VARCHAR(300)" json:"name"  binding:"required"`
	Year int64 	`gorm:"type:INT" json:"year"  binding:"required"`
}
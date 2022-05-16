package book

import "time"

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required" gorm:"type:varchar(100);not null"`
	Price       int    `json:"price" binding:"required,number" gorm:"type:int;not null"`
	Description string `json:"description" binding:"required" gorm:"type:varchar(255);not null"`
	Rating      int    `json:"rating" binding:"required" gorm:"type:int;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

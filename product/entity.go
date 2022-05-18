package product

import "time"

type Product struct {
	ID        int    `json:"id"`
	Product   string `json:"product" binding:"required" gorm:"type:varchar(100);not null"`
	Price     int    `json:"price" binding:"required,number" gorm:"type:int;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

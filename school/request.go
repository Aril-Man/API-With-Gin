package school

type SchoolRequest struct {
	Name    string `json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Address string `json:"address" binding:"required" gorm:"type:varchar(225);not null"`
	Class   string `json:"class" binding:"required" gorm:"type:varchar(3);not null"`
	Major   string `json:"major" binding:"required" gorm:"type:varchar(4);not null"`
}

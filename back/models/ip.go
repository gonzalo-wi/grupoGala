package models

type VisitorIP struct {
	ID        uint   `gorm:"primaryKey"`
	IP        string `gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}

package model

import "time"

type TokenDiblokir struct {
	ID        uint   `gorm:"primaryKey"`
	TokenID   string `gorm:"size:100;uniqueIndex;not null"`
	ExpiredAt time.Time
	CreatedAt time.Time
}

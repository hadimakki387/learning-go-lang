package models

import "github.com/google/uuid"

type Post struct {
	ID      uuid.UUID `json:"id" gorm:"primaryKey"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid"`
}

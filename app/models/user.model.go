package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Name         string         // A regular string field
	Email        string         // A pointer to a string, allowing for null values
	Password     string         // A regular string field that is not exported
	Age          uint8          // An unsigned 8-bit integer
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	Posts        []Post         `gorm:"foreignKey:UserID"`
}

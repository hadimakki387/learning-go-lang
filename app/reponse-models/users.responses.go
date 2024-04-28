package reponsemodels

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID           uuid.UUID      `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Age          uint8          `json:"age"`
	MemberNumber sql.NullString `json:"member_number"`
	ActivatedAt  sql.NullTime   `json:"activated_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Token        string         `json:"token"`
}

package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	// Add other fields as needed
	//gorm.Model
}

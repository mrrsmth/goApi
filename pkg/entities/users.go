package entities

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"firstName,omitempty" db:"first_name"`
	LastName  string    `json:"lastName,omitempty" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

package models

import (
	"time"
)

// User : User model
type User struct {
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Date     time.Time          `json:"date,omitempty" bson:"date,omitempty"`
}

package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// User : User model
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Date     time.Time          `json:"date,omitempty" bson:"date,omitempty"`
}

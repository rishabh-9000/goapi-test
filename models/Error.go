package models

// Err : Error response format
type Err struct {
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

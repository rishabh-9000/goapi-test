package models

// JWT : Model for JWT
type JWT struct {
	Token string `json:"token,omitempty" bson:"token,omitempty"`
}

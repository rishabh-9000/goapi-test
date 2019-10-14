package validators

import (
	"social-api/models"
)

// NameValidation : Validator for Name
func NameValidation(name string) models.Err {
	var response models.Err

	if name == "" {
		response.Message = "Name cannot be empty"
	}

	return response
}

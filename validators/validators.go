package validators

import (
	"regexp"
)

// IsEmpty : Check if string is empty
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

// IsChar : Check if string only constains characters
func IsChar(s string) bool {
	re := regexp.MustCompile(`[a-zA-Z]`)
	return re.Match([]byte(s))
}

// IsNumeric : Check if string only contains numbers
func IsNumeric(s string) bool {
	re := regexp.MustCompile(`[0-9]`)
	return re.Match([]byte(s))
}

// IsEmail : Check if string is email
func IsEmail(s string) bool {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z]+\.[a-zA-Z]+`)
	return re.Match([]byte(s))
}

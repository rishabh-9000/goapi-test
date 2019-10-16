package validators

import (
	"regexp"
)

// IsEmpty : Check if string is empty
func IsEmpty(s string) (bool, string) {
	re := regexp.MustCompile(`\.+`)
	return re.Match([]byte(s)), "Input Empty if True."
}

// IsChar : Check if string only constains characters
func IsChar(s string) (bool, string) {
	re := regexp.MustCompile(`[a-zA-Z]`)
	return re.Match([]byte(s)), "Input is not just characters."
}

// IsNumeric : Check if string only contains numbers
func IsNumeric(s string) (bool, string) {
	re := regexp.MustCompile(`[0-9]`)
	return re.Match([]byte(s)), "Input is not just numeric."
}

// IsEmail : Check if string is email
func IsEmail(s string) (bool, string) {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z]+\.[a-zA-Z]+`)
	return re.Match([]byte(s)), "Input is not an email."
}

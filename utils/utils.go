package utils

import "regexp"

// Contains checks if a slice contains a specific value
func Contains(slice []string, value string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}

	return false
}

func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

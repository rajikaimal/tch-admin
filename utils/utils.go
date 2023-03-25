package utils

// Contains checks if a slice contains a specific value
func Contains(slice []string, value string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}

	return false
}

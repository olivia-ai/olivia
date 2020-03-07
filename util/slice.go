package util

// Contains checks if a string slice contains a specified string
func Contains(slice []string, text string) bool {
	for _, item := range slice {
		if item == text {
			return true
		}
	}

	return false
}

// Index returns a string index in a string slice
func Index(slice []string, text string) int {
	for i, item := range slice {
		if item == text {
			return i
		}
	}

	return 0
}

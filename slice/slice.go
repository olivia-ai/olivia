package slice

// Check if a slice contains a specific item
func SliceContains(slice []string, text string) bool {
	for _, item := range slice {
		if item == text {
			return true
		}
	}

	return false
}
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

// Difference returns the difference of slice and slice2
func Difference(slice []string, slice2 []string) (difference []string) {
	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				difference = append(difference, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice, slice2 = slice2, slice
		}
	}

	return difference
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

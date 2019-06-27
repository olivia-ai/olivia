package util

func Contains(slice []string, text string) bool {
	for _, item := range slice {
		if item == text {
			return true
		}
	}

	return false
}

func Index(slice []string, text string) int {
	for i, item := range slice {
		if item == text {
			return i
		}
	}

	return 0
}

func Remove(slice []string, s string) []string {
	index := Index(slice, s)
	return append(slice[:index], slice[index+1:]...)
}
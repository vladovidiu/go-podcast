package utils

// CheckNullString checks if string is nil
func CheckNullString(str string) *string {
	if len(str) == 0 {
		return nil
	}

	return &str
}

package helper

import "regexp"

// CheckFilename validates file name (path and ./.. is invalid!)
func CheckFilename(date string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

	return reg.MatchString(date)
}

// CheckDateSeperated validates XXXX-YY-ZZ formats
func CheckDateSeperated(date string) bool {
	reg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	return reg.MatchString(date)
}

// CheckDate validates XXXXYYZZ formats
func CheckDate(date string) bool {
	reg := regexp.MustCompile(`^\d{4}\d{2}\d{2}$`)

	return reg.MatchString(date)
}

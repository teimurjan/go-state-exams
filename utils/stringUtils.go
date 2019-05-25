package utils

import "strings"

// CaseInsensetiveContains checks if a substring is in the string not taking case into account
func CaseInsensetiveContains(s string, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}

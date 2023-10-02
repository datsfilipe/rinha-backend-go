package utils

import (
	"time"
)

func IsValidDate(inputDate string) bool {
	_, err := time.Parse("2006-01-02", inputDate)
	return err == nil
}

func VerifyLength(input string, min int, max int) bool {
	return len(input) >= min && len(input) <= max
}

func ValidateUUID(input string) bool {
	return len(input) == 36
}

func ValidSearchTerm(t string) bool {
	return len(t) >= 1
}

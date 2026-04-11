// Package contact manages contact-related functionalities.
package contact

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
}

func New(name, phone string) Contact {
	return Contact{name, phone}
}

func (c Contact) Display() string {
	c.formatDRCPhoneNumber()

	return c.Name + " - " + c.Phone
}

// formatDRCPhoneNumber normalizes and formats a phone number to the DRC (RDC) standard.
// It removes separators, ensures the +243 prefix, and applies the format: +243 XXX XXX XXX.
// The method modifies the Contact's Phone field in place.
func (c *Contact) formatDRCPhoneNumber() {
	c.Phone = normalizeDRCPhoneNumber(c.Phone)
}

// normalizeDRCPhoneNumber cleans, normalizes, and formats a DRC phone number.
// Returns the formatted number or the original input if normalization fails.
func normalizeDRCPhoneNumber(raw string) string {
	cleaned := cleanPhoneNumber(raw)
	normalized := ensureDRCPrefix(cleaned)

	if !isValidDRCLength(normalized) {
		return raw // fallback to original input
	}

	return formatDRCNumber(normalized)
}

// cleanPhoneNumber removes all common separators from a phone number.
func cleanPhoneNumber(raw string) string {
	replacer := strings.NewReplacer(
		" ", "",
		"-", "",
		"(", "",
		")", "",
		".", "",
		"/", "",
	)
	return replacer.Replace(raw)
}

// ensureDRCPrefix guarantees the DRC prefix (+243) is present.
// Handles cases: no prefix, +0 prefix, non-DRC prefix.
func ensureDRCPrefix(num string) string {
	// Add leading + if missing
	if !strings.HasPrefix(num, "+") {
		num = "+" + num
	}

	// Handle +0... pattern (e.g., +0XXXXXXXXX)
	if strings.HasPrefix(num, "+0") {
		return "+243" + strings.TrimPrefix(num, "+0")
	}

	// Replace any non-DRC prefix with +243
	if !strings.HasPrefix(num, "+243") {
		return "+243" + strings.TrimPrefix(num, "+")
	}

	return num
}

// isValidDRCLength validates that the number has exactly 9 digits after the prefix.
func isValidDRCLength(num string) bool {
	digits := strings.TrimPrefix(num, "+243")
	return len(digits) == 9 && isAllDigits(digits)
}

// isAllDigits checks if a string contains only numeric characters.
func isAllDigits(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

// formatDRCNumber applies the standard DRC format: +243 XXX XXX XXX
func formatDRCNumber(num string) string {
	digits := strings.TrimPrefix(num, "+243")

	if len(digits) != 9 {
		return "+243 " + digits // fallback for non-standard length
	}

	return fmt.Sprintf("+243 %s %s %s",
		digits[0:3],
		digits[3:6],
		digits[6:9],
	)
}

func formatPhone(phone string) string {
	return strings.ReplaceAll(phone, " ", "")
}

package str

import (
	"strings"
)

// Contains checks whether collection contains element
func Contains(collection []string, element string) bool {
	for _, el := range collection {
		if el == element {
			return true
		}
	}
	return false
}

// JoinIgnoreEmpty joins string parts with the given separator ignoring empty values
func JoinIgnoreEmpty(sep string, parts ...string) string {
	var nonEmpty []string
	for _, part := range parts {
		if len(part) > 0 {
			nonEmpty = append(nonEmpty, part)
		}
	}
	return strings.Join(nonEmpty, sep)
}

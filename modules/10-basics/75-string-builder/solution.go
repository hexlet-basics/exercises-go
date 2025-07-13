package solution

import (
	"strings"
	"unicode"
)

// BEGIN

func FilterString(text string, exclude rune) string {
	var result strings.Builder
	excludeLower := unicode.ToLower(exclude)

	for _, r := range text {
		if unicode.ToLower(r) != excludeLower {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// END

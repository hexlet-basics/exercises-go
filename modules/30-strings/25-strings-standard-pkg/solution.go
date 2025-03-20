package solution

import (
	"strings"
	"unicode"
)

// BEGIN (write your solution here)
func latinLetters(s string) string {
	sb := &strings.Builder{}

	for _, r := range s {
		if unicode.Is(unicode.Latin, r) {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// END

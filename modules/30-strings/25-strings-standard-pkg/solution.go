package solution

import (
	"strings"
)

// BEGIN

// latinLetters returns only latin letters from the string s filtering all other chars.
func latinLetters(s string) string {
	sb := &strings.Builder{}

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// END

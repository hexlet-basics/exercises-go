package solution

import "unicode"

// BEGIN

// isASCII checks whether the string s contains only ASCII chars.
func isASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// END

package solution

import (
	"strings"
)

// BEGIN

func ModifySpaces(s, mode string) string {
	var replacement string

	switch mode {
	case "dash":
		replacement = "-"
	case "underscore":
		replacement = "_"
	default:
		replacement = "*"
	}

	return strings.ReplaceAll(s, " ", replacement)
}

// END

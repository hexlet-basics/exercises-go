package solution

// BEGIN

import (
	"strings"
)

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

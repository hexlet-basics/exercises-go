package solution

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
)

// BEGIN

// Greetings returns the greetings string for a name.
func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = cases.Title(name)

	return fmt.Sprintf("Привет, %s!", name)
}

// END

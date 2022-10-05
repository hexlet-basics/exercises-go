package solution

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// BEGIN

// Greetings returns the greetings string for a name.
func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = cases.Title(language.Russian).String(name)

	return fmt.Sprintf("Привет, %s!", name)
}

// END

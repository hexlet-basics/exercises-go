package solution

import (
	"fmt"
	"strings"
)

// BEGIN

// Greetings returns the greetings string for a name.
func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)

	return fmt.Sprintf("Привет, %s!", name)
}

// END

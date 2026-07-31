package solution

// BEGIN

import (
	"fmt"
)

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}

	return fmt.Sprintf("%s.%s", locale, domain)
}

// END

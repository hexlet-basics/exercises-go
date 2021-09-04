package solution

import (
	"fmt"
)

// BEGIN

// DomainForLocale adds a subdomain to a domain depending on locale.
func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}

	return fmt.Sprintf("%s.%s", locale, domain)
}

// END

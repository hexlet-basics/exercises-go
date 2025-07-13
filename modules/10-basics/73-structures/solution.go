package solution

import "fmt"

// BEGIN

type Package struct {
	ID        string
	Delivered bool
}

func (p Package) Status() string {
	if p.Delivered {
		return fmt.Sprintf("Package %s has been delivered", p.ID)
	}
	return fmt.Sprintf("Package %s is still in transit", p.ID)
}

// END

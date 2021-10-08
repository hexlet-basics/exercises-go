package solution

// Person is a struct that keeps info about person's age
type Person struct {
	Age uint8
}

// BEGIN

// PersonList is a list of persons.
type PersonList []Person

// GetAgePopularity calculates and returns popularity for each age in the person list.
func (pl PersonList) GetAgePopularity() map[uint8]int {
	popularity := make(map[uint8]int)
	for _, p := range pl {
		popularity[p.Age]++
	}

	return popularity
}

// END

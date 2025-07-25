package solution

// BEGIN

func GetWorkHours(schedule []int) []int {
	if len(schedule) < 5 {
		return []int{}
	}
	return schedule[:5]
}

// END

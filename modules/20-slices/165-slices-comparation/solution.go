package solution

import "slices"

// BEGIN

func AreOrderHistoriesEqual(history1, history2 [][]string) bool {
	if len(history1) != len(history2) {
		return false
	}
	for i := range history1 {
		if !slices.Equal(history1[i], history2[i]) {
			return false
		}
	}
	return true
}

// END

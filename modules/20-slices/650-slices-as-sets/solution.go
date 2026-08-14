package solution

import "github.com/samber/lo"

// BEGIN

func CompareProductLists(oldList, newList []string) (added, removed []string) {
	added, removed = lo.Difference(newList, oldList)
	return added, removed
}

// END

package solution

import "slices"

// BEGIN

func CompareProductLists(oldList, newList []string) (added, removed []string) {
	for _, item := range newList {
		if !slices.Contains(oldList, item) {
			added = append(added, item)
		}
	}

	for _, item := range oldList {
		if !slices.Contains(newList, item) {
			removed = append(removed, item)
		}
	}

	return added, removed
}

// END

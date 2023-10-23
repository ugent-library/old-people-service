package gismo

import (
	"slices"
)

func UniqStrings(vals []string) (newVals []string) {
	for _, val := range vals {
		if !slices.Contains(newVals, val) {
			newVals = append(newVals, val)
		}
	}
	return
}

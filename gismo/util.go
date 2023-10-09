package gismo

func InArray(values []string, val string) bool {
	for _, v := range values {
		if val == v {
			return true
		}
	}
	return false
}

func UniqStrings(vals []string) (newVals []string) {
	for _, val := range vals {
		if !InArray(newVals, val) {
			newVals = append(newVals, val)
		}
	}
	return
}

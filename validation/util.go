package validation

func InArray(values []string, val string) bool {
	for _, v := range values {
		if val == v {
			return true
		}
	}
	return false
}

package arrayutil

func Uniq[T comparable](values []T) []T {
	newValues := make([]T, 0)

	for _, v := range values {
		found := false
		for _, newV := range newValues {
			if v == newV {
				found = true
				break
			}
		}
		if !found {
			newValues = append(newValues, v)
		}
	}

	return newValues
}

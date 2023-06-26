package helper

func InSlice[T string | int](slice []T, value T) (result bool) {
	for _, val := range slice {
		if val == value {
			result = true
			break
		}
	}

	return
}

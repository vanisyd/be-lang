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

func Except[T string | int](slice map[T]interface{}, keys []T) map[T]interface{} {
	resultMap := make(map[T]interface{})

	for key, val := range slice {
		if !InSlice(keys, key) {
			resultMap[key] = val
		}
	}

	return resultMap
}

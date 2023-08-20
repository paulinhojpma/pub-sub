package helpers

func CheckIfElementExistInSliceThatsMatchs[K comparable, T interface{}](value K, elements []T, getValue func(element T) K) bool {
	for _, element := range elements {
		if value == getValue(element) {
			return true
		}
	}
	return false
}

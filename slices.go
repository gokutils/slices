package slices

/*
Filter remove from slice an elements
*/
func Filter[T comparable](arr []T, v T) []T {
	ret := []T{}
	for i := range arr {
		if arr[i] == v {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}

/*
Containe search if elements as been in slice
*/
func Containe[T comparable](arr []T, v T) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}

/*
Unique remove all duplicate elements
*/
func Unique[T comparable](arr []T) []T {
	ret := []T{}
	for i := range arr {
		if Containe(ret, arr[i]) {
			continue
		}
		ret = append(ret, arr[i])
	}
	return ret
}

/*
Find call iterator on each elements, if iterator return true return object and true else return (default, false)
*/
func Find[T comparable](arr []T, iterator func(item T) bool) (T, bool) {
	for i := range arr {
		if iterator(arr[i]) {
			return arr[i], true
		}
	}
	var noop T
	return noop, false
}

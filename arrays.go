package ffp

// Filter is a generic function that filters the input slice `vals` based on the provided filtering function `f`.
// The function `f` should return true for elements that need to be included in the filtered slice.
// It returns a new slice containing the filtered elements.
func Filter[T any](vals []T, f func(T) bool) []T {
	filtered := make([]T, 0)

	for _, val := range vals {
		if f(val) {
			filtered = append(filtered, val)
		}
	}
	return filtered
}

// Map is a generic function that maps the input slice `vals` to another type using the provided mapping function `f`.
// The function `f` should return a value of the desired output type for each input element.
// It returns a new slice containing the mapped elements.
func Map[T any, V any](vals []T, f func(T) V) []V {
	mapped := make([]V, len(vals))

	for i, val := range vals {
		mappedVal := f(val)
		mapped[i] = mappedVal
	}

	return mapped
}

// MapResult is a generic function that maps the input slice `vals` to another type using the provided mapping function `f`.
// The function `f` should return a value of the desired output type for each input element, or an error.
// MapResult returns a new slice of Results of the given output type.
func MapResult[T any, V any](vals []T, f func(T) (V, error)) []Result[V] {
	mapped := make([]Result[V], len(vals))

	for i, val := range vals {
		mappedVal := NewResult(f(val))
		mapped[i] = mappedVal
	}

	return mapped
}

// MapOrEmpty takes in a slice of Results of type T. It calls .OrEmpty on each element,
// and returns a slice of the resulting values.
func MapOrEmpty[T any](vals []Result[T]) [](T) {
	mapped := make([]T, len(vals))

	for i, val := range vals {
		mapped[i] = val.OrEmpty()
	}

	return mapped
}

// Every is a generic function that checks if all the elements in the input slice `vals` satisfy the provided condition `f`.
// The function `f` should return true if the element meets the condition, and false otherwise.
// It returns true if all elements satisfy the condition, and false otherwise.
func Every[T any](vals []T, f func(T) bool) bool {
	for _, val := range vals {
		if !f(val) {
			return false
		}
	}
	return true
}

// ForEach runs thge given function for each element in the array
func ForEach[T any](vals []T, f func(T)) {
	for _, val := range vals {
		f(val)
	}
}

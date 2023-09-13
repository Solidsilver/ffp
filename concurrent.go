package ffp

import "sync"

// MapConcurrent applies a function to each value in the input slice concurrently and returns a slice of the results.
// Note: the mapped result can be out of order of the original
func MapConcurrent[T any, V any](vals []T, f func(T) V) []V {
	var wg sync.WaitGroup
	c := make(chan V, len(vals))

	// Iterate through the input values and apply the function concurrently.
	for _, val := range vals {
		wg.Add(1)
		go func(fnInput T) {
			defer wg.Done()
			result := f(fnInput)
			c <- result
		}(val)
	}

	// Close the channel once all the input values have been processed.
	go func() {
		wg.Wait()
		close(c)
	}()

	// Collect the results from the channel and append them to the results slice.
	results := []V{}
	for result := range c {
		results = append(results, result)
	}
	return results
}

// MapConcurrentResult is similar to MapConcurrent but expects the function to return a value and an error.
// The results are returned as a slice of ConcurrentRunResult.
func MapConcurrentResult[T any, V any](vals []T, f func(T) (V, error)) []Result[V] {
	var wg sync.WaitGroup
	c := make(chan Result[V], len(vals))

	// Iterate through the input values and apply the function concurrently.
	for _, val := range vals {
		wg.Add(1)
		go func(fnInput T) {
			defer wg.Done()
			result, err := f(fnInput)
			c <- Result[V]{
				value: &result,
				err:   err,
			}
		}(val)
	}

	// Close the channel once all the input values have been processed.
	go func() {
		wg.Wait()
		close(c)
	}()

	// Collect the results from the channel and append them to the results slice.
	results := make([]Result[V], len(vals))
	idx := 0
	for result := range c {
		results[idx] = result
		idx++
	}
	return results
}

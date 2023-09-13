package ffp

import (
	"testing"

	"golang.org/x/exp/slices"
)

func defaultArray() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func TestFilter(t *testing.T) {
	arr := defaultArray()
	filteredArr := Filter(arr, func(item int) bool {
		return item%2 == 1
	})
	if slices.Compare(filteredArr, []int{1, 3, 5, 7, 9}) != 0 {
		t.Errorf("Slice was not filtered properly: %+v", filteredArr)
	}
}

func TestMap(t *testing.T) {
	arr := defaultArray()
	mappedArray := Map(arr, func(item int) int {
		return item * 2
	})
	if slices.Compare(mappedArray, []int{2, 4, 6, 8, 10, 12, 14, 16, 18}) != 0 {
		t.Errorf("Slice was not mapped properly: %+v", mappedArray)
	}
}

func TestEvery(t *testing.T) {
	arr := defaultArray()
	smallArray := Every(arr, func(item int) bool {
		return item < 10
	})
	if !smallArray {
		t.Errorf("Every function did not work properly on the slice")
	}
}

func TestMapConcurrent(t *testing.T) {
	arr := defaultArray()
	mappedArray := MapConcurrent(arr, func(item int) int {
		return item * 2
	})
	slices.Sort(mappedArray)
	if slices.Compare(mappedArray, []int{2, 4, 6, 8, 10, 12, 14, 16, 18}) != 0 {
		t.Errorf("Slice was not mapped properly: %+v", mappedArray)
	}
}

func TestMapConcurrentResult(t *testing.T) {
	arr := defaultArray()
	mappedArray := MapConcurrentResult(arr, func(item int) (int, error) {
		return item * 2, nil
	})
	slices.SortFunc(mappedArray, func(a, b Result[int]) int {
		if a.IsOk() && b.IsOk() {
			return a.OrEmpty() - b.OrEmpty()
		}
		if b.IsOk() {
			return -1
		}
		if a.IsOk() {
			return 1
		}
		return 0

	})
	if slices.Compare(MapOrEmpty(mappedArray), []int{2, 4, 6, 8, 10, 12, 14, 16, 18}) != 0 {
		t.Errorf("Slice was not mapped properly: %+v", mappedArray)
	}
}

package quicksort

func Quicksort(array []int) {
	partition := func(array []int, low int, high int) (int, int) {
		mid := low + (high - low) / 2
		pivot := array[mid]
		lesserIndex := low
		equalIndex := low
		greaterIndex := high

		for equalIndex <= greaterIndex {
			if array[equalIndex] < pivot {
				
			} else if array[equalIndex] > pivot {
				
			} else {
				
			}
		}

		return lesserIndex, greaterIndex
	}

	var f func([]int, int, int)
	f = func(array []int, low int, high int) {
		if low < high {
			lesserIndex, greaterIndex := partition(array, low, high)
			f(array, low, lesserIndex - 1)
			f(array, greaterIndex + 1, high)
		}
	}

	f(array, 0, len(array) - 1)
}

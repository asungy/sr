package quicksort

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	array := GenerateRandomArray(20, 0, 100)
	Quicksort(array)
	for i := 1; i < len(array); i++ {
		assert.LessOrEqual(t, array[i-1], array[i])
	}

	t.Log(StringifyArray(array))
}

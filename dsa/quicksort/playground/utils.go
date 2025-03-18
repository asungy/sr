package quicksort

import (
	"fmt"
	"math/rand"
	"strings"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func GenerateRandomArray(size, min, max int) []int {
	result := make([]int, size)
	for i := range size {
		result[i] = randInt(min, max)
	}

	return result
}

func StringifyArray(array []int) string {
	var sb strings.Builder

	sb.WriteRune('[')
	for _, num := range array[:len(array)-1] {
		sb.WriteString(fmt.Sprintf("%d, ", num))
	}

	return sb.String()
}

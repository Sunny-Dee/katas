package mergersort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	arrA := make([]int, 4)
	arrA[0] = 2
	arrA[1] = 6

	fmt.Println(arrA)

	arrB := []int{1, 9}

	expected := []int{1, 2, 6, 9}
	MergeSort(&arrA, &arrB, 2)
	assert.Equal(t, expected, arrA)
}

func TestMergeSortLongArray(t *testing.T) {
	arrA := make([]int, 6)
	arrA[0] = 2
	arrA[1] = 10

	fmt.Println(arrA)

	arrB := []int{3, 9, 12, 14}

	expected := []int{2, 3, 9, 10, 12, 14}
	MergeSort(&arrA, &arrB, 2)
	assert.Equal(t, expected, arrA)
}

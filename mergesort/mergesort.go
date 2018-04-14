package mergersort

import "fmt"

// MergeSort must provide the value of A to know
// where the buffered values start
func MergeSort(arrA, arrB *[]int, lenA int) *[]int {
	a := *arrA
	b := *arrB

	i := lenA - 1
	j := len(b) - 1
	x := len(a) - 1
	for x > 0 {
		if a[i] < b[j] {
			(*arrA)[x] = b[j]
			j--
		} else {
			(*arrA)[x] = a[i]
			i--
		}
		x--
	}
	if a[0] < b[0] {
		(*arrA)[0] = a[0]
	} else {
		(*arrA)[0] = b[0]
	}

	fmt.Println(arrA)
	return arrA
}

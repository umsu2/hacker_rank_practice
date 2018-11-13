package main

import (
	"fmt"
	"sort"
)

var array1 = []int{1, 4, 5, 1, 3, 2, 45, 6, 3, 4, 67, 7}
var array2 = []int{3, 4, 6, 7, 8, 39, 20, 34}

func main() {
	v := findMaxValue(array1, array2, 4)
	fmt.Println(v)
}

// finds the closest value to target but not greater than target

// if array1 is size N, array2 is size M, and if M > N, the overall speed is O(mLogn)
// m binary searches on array n O(mLogn) plus nLogn on sorting of the smaller array
func findMaxValue(array1, array2 []int, target int) int {
	sortedArray := array1
	iteratingArray := array2
	if len(array1) < len(array2) {
		sort.Ints(array1)
	} else {
		sortedArray = array2
		iteratingArray = array1
		sort.Ints(array2)
	}
	sum := -1
	for _, v := range iteratingArray {
		val := target - v
		r := modifiedBinarySearch(sortedArray, val)
		if r < 0 {
			continue
		}
		s := r + v
		if s == target {
			return target
		} else if s > sum {
			sum = s
		}
	}
	return sum
}

// modifiedBinarySearch attempts to find a value closest to val, but not greater than val, if no value is found, return -1
// assuming sorted ASC
func modifiedBinarySearch(sorted []int, val int) int {
	// assuming nothing cost 0 dollars and you must buy two things.
	// some small optimizations
	if val <= 0 {
		return -1
	}
	if len(sorted) == 0 {
		return -1
	}
	left := 0
	right := len(sorted) - 1
	min := sorted[left]
	max := sorted[right]
	if val < min {
		return -1
	}
	if val == min || val == max {
		return val
	}
	// do binary search

	for true {
		middle := (left + right) / 2
		if middle == left {
			return sorted[left]
		}
		if sorted[middle] == val {
			return val
		} else if sorted[middle] > val {
			right = middle
		} else {
			left = middle
		}
	}
	return -1
}

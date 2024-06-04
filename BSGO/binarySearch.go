package main

import (
	"fmt"
	"sort"
)

func recursiveFunction(l int, r int, slice []int, target int) int {

	mid := (l + r) / 2
	fmt.Println("Mid", mid)
	if l > r {
		return -1
	}

	if target > slice[mid] {
		fmt.Println("Upper", l)
		l = mid + 1
	} else if target < slice[mid] {
		fmt.Println("Downer", r)
		r = mid - 1
	} else {
		return mid
	}
	fmt.Println(slice[l : r+1])
	return recursiveFunction(l, r, slice, target)
}

func beginningFuction(arr [6]int, target int) int {
	fmt.Println(arr, target)
	slice := arr[:]

	sort.Ints(slice)
	l := 0
	r := len(slice) - 1

	fmt.Println(slice, l, r)

	return recursiveFunction(l, r, slice, target)
}

func main() {
	fmt.Println(beginningFuction([6]int{5, 4, 1, 3, 2, 6}, 7))
}

package main

import (
	"fmt"
)

func Bsearch(l, r int, a []int, goal int) int {
	if l <= r {
		mid := (l + r) / 2
		if a[mid] == goal {
			return mid
		} else if a[mid] > goal {
			return bsearch(0, mid-1, a, goal)
		} else if a[mid] < goal {
			return bsearch(mid+1, r, a, goal)
		}
	}

	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	goal := 2

	res := Bsearch(0, len(a)-1, a, goal)

	fmt.Println(res)
}

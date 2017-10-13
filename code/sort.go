package main

import (
	"fmt"
)

/*******************insert****************/
func insert(a []int, idx, value int) {
	pos := idx - 1
	for pos >= 0 && a[pos] > value {
		a[pos+1] = a[pos]
		pos--
	}
	a[pos+1] = value
}

func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		insert(a, i, a[i])
	}
}

/*******************insert****************/

/*******************quic******************/
func partition(l, r int, a []int) int {
	mv := a[l]

	for l < r {
		for l < r && a[r] >= mv {
			r--
		}
		if l < r {
			a[l] = a[r]
			//l++
			//r--
		}

		for l < r && a[l] <= mv {
			l++
		}
		if l < r {
			a[r] = a[l]
			//r--
			//l++
		}
	}

	a[l] = mv
	return l
}

func qsort(l, r int, a []int) {
	if l < r {
		piv := partition(l, r, a)
		qsort(l, piv-1, a)
		qsort(piv+1, r, a)
	}
}

func QuickSort(a []int) {
	qsort(0, len(a)-1, a)
}

/*******************quic******************/

func HeapSort(a []int) {}

func CounterSort(a []int) {}

func BucketSort(a []int) {}

func main() {
	a := []int{1, 3, 2, 5, 4, 8, 7, 0, 6, 10}

	//InsertSort(a)
	QuickSort(a)

	fmt.Println(a)
}

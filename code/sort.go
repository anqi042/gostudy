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

/*******************quick******************/
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

/*******************quick******************/
func heapify(a []int, idx, max int) {
	left := 2*idx + 1
	right := 2*idx + 2
	largest := idx

	if left < max && a[left] > a[idx] {
		largest = left
	}

	if right < max && a[right] > a[largest] {
		largest = right
	}

	if largest != idx {
		a[largest], a[idx] = a[idx], a[largest]
		heapify(a, largest, max)
	}
}

func buildHeap(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		heapify(a, i, len(a))
	}
}

func HeapSort(a []int) {
	buildHeap(a)
	for i := len(a) - 1; i >= 1; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, 0, i)
	}
}

/*******************counter****************/
func CounterSort(a []int) {
	temp := make([]int, 20)

	for i := 0; i < len(a); i++ {
		temp[a[i]]++
	}

	counter := 0
	for i := 0; i < len(temp); i++ {
		for 0 != temp[i] {
			a[counter] = i
			temp[i]--
			counter++
		}
	}
}

/*******************counter****************/
func hash(v int) int {
	return v / 3
}
func BucketSort(a []int) {
	bucket := make([][]int, len(a))

	for i := 0; i < len(a); i++ {
		index := hash(a[i])
		if nil == bucket[index] {
			bucket[index] = make([]int, 0)
		}
		bucket[index] = append(bucket[index], a[i])
	}

	counter := 0
	for i := 0; i < len(a); i++ {
		if 0 != len(bucket[i]) {
			InsertSort(bucket[i])
			for j := 0; j < len(bucket[i]); j++ {
				a[counter] = bucket[i][j]
				counter++
			}
		}
	}
}

/******************merge*******************/
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

func MergeSort(a []int) []int {
	length := len(a)
	if length <= 1 {
		return a
	}

	mid := length / 2
	right := MergeSort(a[:mid])
	left := MergeSort(a[mid:])
	return merge(left, right)
}

/******************merge*******************/

func main() {
	a := []int{1, 3, 2, 5, 4, 8, 7, 0, 6, 10}

	//InsertSort(a)
	//QuickSort(a)
	//a = MergeSort(a)
	//CounterSort(a)
	//BucketSort(a)
	HeapSort(a)
	fmt.Println(a)

}

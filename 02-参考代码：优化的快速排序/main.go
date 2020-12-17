package main

import (
	"fmt"

)
func main() {
	arr:=[]int{45,23,78,97,43,23,12,34}
	quickSortNew(arr,0,len(arr)-1)
	fmt.Println(arr)
}
func quickSortNew(arr []int, l, r int) {
	for l < r {
		x := l
		y := r
		z := selectValue(arr[l], arr[r], arr[(l+r)>>1])
		for x <= y {
			for arr[x] < z {
				x++
			}
			for arr[y] > z {
				y--
			}
			if x <= y {
				arr[x], arr[y] = arr[y], arr[x]
				x++
				y--
			}
		}
		quickSortNew(arr, x, r)
		r = y
	}
}

func selectValue(a, b, c int) int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	return b
}
package main

import (
	"fmt"
)

var a []int = make([]int, 100000)

func swap(i, j int){
	a[i], a[j] = a[j], a[i]
}
func less(i, j int) bool{
	if a[i] < a[j]{
		return true
	}

	return false
}

func partition(l, r int) int{
	pivot := r
	i := l - 1
	for j := l; j < r; j++ {
		if less(j, pivot){
			i++
			swap(i, j)
		}
	}
	swap(i + 1, r)
	return i + 1
}

func quicksort(l, r int){
	if l >= r {
		return
	}
	pi := partition(l, r)
	
	quicksort(l, pi - 1)
	quicksort(pi+1, r)
}
func qsort(n int, less func(i, j int) bool, swap func(i, j int) ){
	quicksort(0, n - 1)
}

func main(){
	var n int32;
	fmt.Scan(&n)
	for i := 0; i < int(n); i++{
		fmt.Scan(&a[i])
	}
	qsort(int(n), less, swap)
	for i := 0; i < int(n); i++{
		fmt.Print(a[i], " ")
	}
}
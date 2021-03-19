package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){
	var n int32
	fmt.Scanln(&n)
	var a []string = make([]string, n)
	for i := int32(0); i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func (i, j int) bool{
		var first, _ = strconv.ParseInt(a[i] + a[j], 10, 64)
		var second, _ = strconv.ParseInt(a[j] + a[i], 10, 64)
		return first > second
	})
	for i := int32(0); i < n; i++ {
		fmt.Print(a[i])
	}
}
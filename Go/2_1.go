package main

import "fmt"
import "math"
import "sort"

func main(){
	var n int
	var a []int
	fmt.Scanln(&n)
	for i := 1; i <= int(math.Sqrt(float64(n))); i++{
		if (n % i == 0){
			if (float64(i) == math.Sqrt(float64(n))){
				a = append(a, i)
			} else {
				a = append(a, i)
				a = append(a, n/i)
			}
		}
	}
	sort.Ints(a)
	fmt.Println("graph {")
	for i := len(a) - 1; i >= 0; i--{
		fmt.Println("   ", a[i])
	}
	for i := len(a) - 1; i > 0; i--{
		for j := i - 1; j >= 0; j--{
			if (a[i] % a[j] != 0){
				continue
			}
			to_write := true
			for k := i - 1; k > j; k--{
				if (a[i] % a[k] == 0) && (a[k] % a[j] == 0){
					to_write = false
					break
				}
			}
			if (to_write){
				fmt.Println("   ", a[i], "--", a[j])
			}
		}
	}
	fmt.Println("}")
}
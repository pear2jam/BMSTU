package main

import "fmt"
import "math"
import "github.com/skorobogatov/input"

func main(){
	var (
		a_rune, f, s []rune
		pos int32 = -1e9
		first, second string
	)
	a_rune = []rune(input.Gets())
	input.Scanf("%s", &first)
	input.Scanf("%s", &second)
	f = []rune(first)
	s = []rune(second)
	var res int32 = 1e9

	for i := 0; i < len(a_rune); i++{
		if a_rune[i] == f[0]{
			pos = int32(i)
		}
		if a_rune[i] == s[0]{
			res = int32(math.Min(float64(res), float64(int32(i) - pos)))
		}
	}
	pos = -1e9
	for i := 0; i < len(a_rune); i++{
		if a_rune[i] == s[0]{
			pos = int32(i)
		}
		if a_rune[i] == f[0]{
			res = int32(math.Min(float64(res), float64(int32(i) - pos)))
		}
	}

	fmt.Println(res - 1)
}
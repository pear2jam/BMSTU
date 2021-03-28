package main

import "fmt"
import "sort"
import "math"

type frac struct{
	a, b int64
}

//работа с дробями

func gcd(a, b int64) int64{
	if (b != 0){
		return gcd(b, a % b)
	}
	return a
}

func m_f(a, b int64) frac{
	var res frac
	res.a = a
	res.b = b
	if (b < 0){
		res.a *= -1
		res.b *= -1
	}
	return reduce(res)
}

func reduce(a frac) frac{
	var x, y float64 = float64(a.a), float64(a.b)
	var c int64 = gcd(int64(math.Abs(x)), int64(y))
	a.a /= c
	a.b /= c
	return a
}

func sum(x, y frac) frac{
	var res frac
	var a, b, c, d int64 = x.a, x.b, y.a, y.b
	res.a = a*d + b*c
	res.b = b*d
	return reduce(res)
}

func sub(x, y frac) frac{
	return sum(x, m_f(-y.a, y.b))
}

func mul(x, y frac) frac{
	var res frac
	var a, b, c, d int64 = x.a, x.b, y.a, y.b
	res.a = a*c
	res.b = b*d
	return reduce(res)
}

func div(x, y frac) frac{
	var res frac
	var a, b, c, d int64 = x.a, x.b, y.a, y.b
	res.a = a*d
	res.b = b*c
	if (res.b < 0){
		res.a *= -1
		res.b *= -1
	}
	return reduce(res)
}
func equal(a, b frac) bool{
	if (a.a == 0 && b.a == 0){
		return true
	}
	if (a.a == b.a && a.b == b.b){
		return true
	}
	return false
}

func look_frac(a frac){
	fmt.Print("(frac ",a.a, "/", a.b, ") ")
}

func look(a frac){
	fmt.Print(a.a, "/", a.b, "\n")
}
//считает определитель
func det(a [][]frac) frac{
	if (len(a) == 1){
		return a[0][0]
	}
	var res frac = m_f(0, 1)
	for i := 0; i < len(a); i++ {
		new_a := make([][]frac, len(a)-1)
		//здесь создаем матрицу для минора
		for j := 1; j < len(a); j++{ //строка
			for k := 0; k < len(a); k++{ //столбец - тут пропускаем индекс i
				if (k != i){
					new_a[j-1] = append(new_a[j-1], a[j][k])
				}
			}
		}
		var koef frac = m_f(1, 1)
		if (i % 2 == 0){
			koef = m_f(-1, 1)
		}
		koef = mul(koef, a[0][i])
		res = sum(res, mul(koef, det(new_a)))
	}
	return res
}

//сортирует матрицу по убыванию ведущих нулей
func matrix_sort(a [][]frac) [][]frac{
	new_a := make([][]frac, len(a))
	order := make([][]float64, len(a))
	for i := 0; i < len(a); i++{
		var j int8 = 0
		for ; equal(a[i][j], m_f(0, 1)); j++{}
		order[i] = append(order[i], float64(j), float64(i))
	}
	sort.SliceStable(order, func(i, j int) bool {
		return order[i][0] < order[j][0]
	})
	for i := 0; i < len(a); i++{
		new_a[i] = a[int(order[i][1])]
	}
	return new_a
}

//приводит матрицу к верхнетреугольному виду

func traingle(a [][]frac) [][]frac{
	for i := 0; i < len(a) - 1; i++{
		a = matrix_sort(a)
		for j := i + 1; j < len(a); j++{
			if (!equal(a[j][i], m_f(0,1))){
				var dv frac = div(a[j][i], a[i][i])
				for k := 0; k <= len(a); k++{
					a[j][k] = sub(a[j][k], mul(a[i][k], dv))
				}
			}
		}
	}
	return a
}

func solve(a [][]frac) []frac{
	res := make([]frac, len(a))
	var to_sub frac = m_f(0, 1)
	for i := 0; i < len(a); i++{
		for j := len(a) - 1; j > len(a) - i - 1; j--{
			to_sub = sum(to_sub, mul(a[len(a)-i-1][j], res[j]))
		}
		res[len(a) - i - 1] = div(sub(a[len(a)-i-1][len(a)], to_sub), a[len(a)-i-1][len(a)-i-1])
		to_sub = m_f(0, 1)
	}
	return res
}

func main(){
	var n int8
	var val int64
	fmt.Scan(&n)
	a := make([][]frac, n);
	for i := int8(0); i < n; i++{
		to_append := make([]frac, n + 1)
		for j := int8(0); j <= n; j++{
			fmt.Scan(&val)
			to_append[j] = m_f(val, 1)
		}
		a[i] = to_append
	}
	if (equal(det(a), m_f(0,1))){
		fmt.Print("No solution")
		return
	}
	a = traingle(a)
	for _, i := range solve(a){
		look(i)
	}
}

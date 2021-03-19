package main

import "fmt"
import "sort"

//для красивого вывода двумерных срезов для дебаггинга
func look(a [][]float64){
	fmt.Println(">")
	for i := 0; i < len(a); i++{
		for j := 0; j < len(a[i]); j++{
			fmt.Print(a[i][j], " ")
		}
		fmt.Print("\n")
	}
}
func look_(a float64){
	fmt.Println(">")
	fmt.Println(a)


	
}
func warn(){
	fmt.Print("!!!___!!!")
}



//считает определитель
func det(a [][]float64) float64{
	if (len(a) == 1){
		return a[0][0]
	}
	var res float64 = 0
	for i := 0; i < len(a); i++ {
		new_a := make([][]float64, len(a)-1)
		//здесь создаем матрицу для минора
		for j := 1; j < len(a); j++{ //строка
			for k := 0; k < len(a); k++{ //столбец - тут пропускаем индекс i
				if (k != i){
					new_a[j-1] = append(new_a[j-1], a[j][k])
				}
			}
		}
		var koef float64 = 1
		if (i % 2 == 0){
			koef = -1
		}
		koef *= float64(a[0][i])
		res += koef * det(new_a)
	}
	return res
}

//сортирует матрицу по убыванию ведущих нулей
func matrix_sort(a [][]float64) [][]float64{
	new_a := make([][]float64, len(a))
	order := make([][]float64, len(a))
	for i := 0; i < len(a); i++{
		var j int8 = 0
		for ; a[i][j] == 0; j++{}
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

func traingle(a [][]float64) [][]float64{
	for i := 0; i < len(a) - 1; i++{
		a = matrix_sort(a)
		for j := i + 1; j < len(a); j++{
			if (a[j][i] != 0){
				var div float64 = (a[j][i]/a[i][i])
				for k := 0; k <= len(a); k++{
					a[j][k] -= a[i][k]*div
				}
			}
		}
	}
	return a
}

func solve(a [][]float64) []float64{
	res := make([]float64, len(a))
	var to_sub float64 = 0
	for i := 0; i < len(a); i++{
		for j := len(a) - 1; j > len(a) - i - 1; j--{
			to_sub += a[len(a)-i-1][j]*res[j]
		}
		res[len(a) - i - 1] = (a[len(a)-i-1][len(a)] - to_sub)/a[len(a)-i-1][len(a)-i-1]
		to_sub = 0
	}
	return res
}

func main(){
	a := make([][]float64, 3);
	a[0] = append(a[0], -4, -1, 8, 2)
	a[1] = append(a[1], 7, -7, 7, 3)
	a[2] = append(a[2], 5, -1, -4, 7)

	if (det(a) == 0){
		fmt.Print("No solution")
		return
	}
	a = traingle(a)
	look(a)
	for _, i := range solve(a){
		
		fmt.Println(i)
	}
}

//нужно 

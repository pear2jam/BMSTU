package main

import "fmt"

type Pair struct {
	signal string
	next int
}

func main() {
	var n, m, start int
	s_list := "abcdefghijklmnopqrstuvwxyz"
	fmt.Scan(&n, &m, &start)
	mat := make([][]Pair, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]Pair, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&mat[i][j].next)	
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j ++ {
			fmt.Scan(&mat[i][j].signal)	
		}
	}

	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")
	fmt.Println("dummy [label = \"\", shape = none]")
	for i := 0; i < n; i++ {
		fmt.Println(i, "[shape = circle]")
	}
	fmt.Println("dummy ->", start)
	
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(i, " -> ", mat[i][j].next, " [label = \"", string(s_list[j]), "(", mat[i][j].signal, ")\"]\n")
		}
	}
	fmt.Println("}")
}
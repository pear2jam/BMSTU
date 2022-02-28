package main

import "fmt"

type Pair struct {
	signal string
	next int
}

func dfs(mat [][]Pair, vis []bool, to_num []int, to_anti []int, v int, count *int) {
	(*count)++

	to_num[*count - 1] = v
	to_anti[v] = *count - 1
	vis[v] = true

	for i := 0; i < len(mat[v]); i++ {
		if (vis[mat[v][i].next]) {
			continue
		}
		dfs(mat, vis, to_num, to_anti, mat[v][i].next, count)
	}
}

func main() {
	var n, m, start, count int = 0,0,0,0
	fmt.Scan(&n, &m, &start)
	mat := make([][]Pair, n)
	vis := make([]bool, n)

	to_num := make([]int, n)
	to_anti := make([]int, n)

	//input
	for i := 0; i < n; i++ {
		mat[i] = make([]Pair, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&(mat[i][j].next))	
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&(mat[i][j].signal))	
		}
	}

	dfs(mat, vis, to_num, to_anti, start, &count)

	fmt.Println(count)
	fmt.Println(m)
	fmt.Println(0)
	for i := 0; i < count; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(to_anti[mat[to_num[i]][j].next], " ")
		}
		fmt.Println()
	}
	for i := 0; i < count; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(mat[to_num[i]][j].signal, " ")
		}
		fmt.Println()
	}
}
package main

import (
	"fmt"
)

type Pair struct {
	signal string
	next int
}

func get(key int, prev_list []int) int{
    if (key == prev_list[key]) {
        return key
    }

    prev_list[key] = get(prev_list[key], prev_list)
    return prev_list[key]
}

func to_merge(a, b int, prev_list []int, ord []int) {
    first := get(a, prev_list)
    sec := get(b, prev_list)
    if (first != sec) {
		if (ord[first] == ord[sec]) {
            ord[first] += 1
        }
        if (ord[first] < ord[sec]) {
            first, sec = sec, first
        }
        prev_list[sec] = first
    }
}

func to_split(mat [][]Pair, a []int, prev_list []int, ord []int) int {
	for i := 0; i < len(mat); i++ {
		prev_list[i] = i
	}
	n, m := len(mat), 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if ( (get(i, prev_list) != get(j, prev_list)) && (a[i] == a[j]) ) {
				m += 1
				merge_bool := true
				for k := 0; k < len(mat[0]); k += 1 { 
					if (a[mat[i][k].next] != a[mat[j][k].next]) {
						merge_bool = false
						break
					}
				}
				if (merge_bool) {
					to_merge(i, j, prev_list, ord)
					m -= 1
					n -= 1
				}
			}
		}
	}
	for i := 0; i < len(mat); i += 1 {
		a[i] = get(i, prev_list)
	}
	return n
}


func to_minimize(mat [][]Pair, prev_list []int, ord []int) ([]int, []int, [][]Pair) {
	for i := 0; i < len(mat); i++ {
		prev_list[i] = i
	}
	n, m := len(mat), 0

	mat_new := make([][]Pair, 0)
	vis := make(map[int]bool)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if (get(i, prev_list) != get(j, prev_list)) {
				merge_bool := true
				m += 1
				for k := 0; k < len(mat[0]); k++ {
					if (mat[j][k].signal != mat[i][k].signal) {
						merge_bool = false
						break
					}
				}
				if (merge_bool) {
					to_merge(i, j, prev_list, ord)
					m -= 1
					n -= 1
				}
			}
		}
	}
	cur := 0
	a := make([]int, len(mat))
	for i := 0; i < len(mat); i++ {
		a[i] = get(i, prev_list)
	}
	for ;; {
		s := to_split(mat, a, prev_list, ord)
		if (s == n) {
			break
		}
		n = s
	}
	res := make([]int, len(mat))
	for i := 0; i < len(mat); i++ {
		j := a[i]
		res[i] = cur
		if (!vis[j]) {
			vis[j] = true
			cur += 1
		}
	}
	cur = 0
	
	vis = make(map[int]bool)
	for i := 0; i < len(mat); i++ {
		k := a[i]
		if (!vis[k]) {
			mat_new = append(mat_new, make([]Pair, len(mat[0])))
			vis[k] = true
			for j := 0; j < len(mat[0]); j++ {
				next := a[mat[i][j].next]
				mat_new[cur][j].signal = mat[i][j].signal
				mat_new[cur][j].next = res[next]
			}
			cur += 1
		}
	}
	return res, a, mat_new
}

func dfs(mat [][]Pair, vis []bool, to_num []int, to_anti []int, q int, cnt *int) {
	to_num[*cnt] = q
	to_anti[q] = *cnt

	vis[q] = true
	(*cnt) += 1

	for i := 0; i < len(mat[q]); i++ {
		if (vis[mat[q][i].next]) {
			continue
		}
		dfs(mat, vis, to_num, to_anti, mat[q][i].next, cnt)
	}
}

func main() {
	var n, m, start, cnt int = 0,0,0,0
	var prev_list, ord []int
	s_list := "abcdefghijklmnopqrstuvwxyz"
	fmt.Scan(&n, &m, &start)

	mat := make([][]Pair, n)
	prev_list = make([]int, n)
	to_num := make([]int, n)
	to_anti := make([]int, n)

	ord = make([]int, n)
	//input
	for i := 0; i < n; i++ {
		mat[i] = make([]Pair, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&mat[i][j].next)	
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&mat[i][j].signal)	
		}
	}

	res, a, mat_new := to_minimize(mat, prev_list, ord)

	vis := make([]bool, n)
	dfs(mat_new, vis, to_num, to_anti, res[a[start]], &cnt)

	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")
	fmt.Println("dummy [label = \"\", shape = none]")

	for i := 0; i < cnt; i++ {
		fmt.Print(i, " [shape = circle]", "\n")
	}
	fmt.Println("dummy ->", 0)

	for i := 0; i < cnt; i++ {
		for j := 0; j < len(mat_new[i]); j++ {
			fmt.Print(i, " -> ", to_anti[mat_new[to_num[i]][j].next], " [label = \"", string(s_list[j]), "(", mat_new[to_num[i]][j].signal, ")\"]\n")
		}
	}
	fmt.Println("}")
}


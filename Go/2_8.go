package main

import "fmt"

type Vert struct {
	compare, ind, min_val int
	used, vis bool
	next []int

}

func top_sort(g []Vert, queue *[]int, v int) {
	g[v].used = true
	for i := 0; i < len(g[v].next); i++ {
		if (!g[g[v].next[i]].used) {
			top_sort(g, queue, g[v].next[i])
		}
	}
	*queue = append(*queue, v)
	g[v].vis = true	
}

func combine(g []Vert, v int, compare int, g_cond []Vert) {
	g[v].used = true
	g[v].compare = compare
	if (g_cond[compare].min_val > v) {
		g_cond[compare].min_val = v
	}
	for i := 0; i < len(g[v].next); i++ {
		if g[g[v].next[i]].used { 
			if (g[g[v].next[i]].compare != compare) {
				g_cond[compare].next = append(g_cond[compare].next, g[g[v].next[i]].compare)
				g_cond[g[g[v].next[i]].compare].ind += 1
			}
			continue 
		}
		combine(g, g[v].next[i], compare, g_cond)
	}
}

func reset(g []Vert) {
	for i := 0; i < len(g); i++ {
		g[i].used = false
	}
}

func get_ver() Vert{
	var ver Vert
	ver.next = make([]int, 0)
	return ver
}

func get_ver_c() Vert{
	var ver Vert
	ver.min_val = 100000
	ver.next = make([]int, 0)
	return ver
}

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m)
	var g []Vert
	g_cond := make([]Vert, 0)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		g = append(g, get_ver())
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		g[a].next = append(g[a].next, b)
	}
	for i := 0; i < len(g); i++ {
		if (!g[i].used) {
			top_sort(g, &queue, i)
		}
	}
	reset(g)
	
	for i := 0; i < len(queue); i++ {
		if (!g[queue[i]].used) {
			g_cond = append(g_cond, get_ver_c())
			combine(g, queue[i], len(g_cond) - 1, g_cond)
		}
	}
	for i := 0; i < len(g_cond); i++ {
		if (g_cond[i].ind == 0) {
			fmt.Print(g_cond[i].min_val, "\n")
		}
	}
}
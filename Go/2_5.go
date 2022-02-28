package main

import "fmt"

type Vert struct {
	next, d []int
}

func get_vert() Vert {
	var v Vert
	v.next = make([]int, 0)
	v.d = make([]int, 0)
	return v
}

func bfs(root, id int, g []Vert) {
	start := 0
	g[root].d[id] = 0

	queue := make([]int, 0)
	queue = append(queue, root)
	for j := 0 ;start < len(queue); j += 1 {
		v := queue[start]
		start += 1
		for i := 0; i < len(g[v].next); i++ {
			if (g[v].d[id] + 1 > g[g[v].next[i]].d[id]) {
				continue
			}
			queue = append(queue, g[v].next[i])
			g[g[v].next[i]].d[id] = g[v].d[id] + 1
			
		}
	}
}

func main() {
	var n, m, a, b, q int	
	fmt.Scan(&n, &m)

	var g []Vert
	for i := 0; i < n; i++ {
		g = append(g, get_vert())
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		if (a != b) {
			g[a].next = append(g[a].next, b)
		}
		g[b].next = append(g[b].next, a)
	}
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var v int
		fmt.Scan(&v)
		for j := 0; j < n; j++ {
			g[j].d = append(g[j].d, n)
		}

		bfs(v, i, g)
	}
	cnt := 0
	for i := 0; i < n; i++ {
		equal := true
		for j := 1; j < q; j++ {
			if ((g[i].d[j] == n) || (g[i].d[j - 1] != g[i].d[j]) ) {
				equal = false
				break
			}
		}
		if (equal) {
			fmt.Print(i, " ")
			cnt += 1
		}
	}
	if (cnt == 0) {
		fmt.Print("-")
	}
}
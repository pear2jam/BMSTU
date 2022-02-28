package main

import(
	"fmt"
	"math"
)

type Vert struct {
	in, from int
	next []int
	vis bool
}

func get_vert() Vert{
	var ver Vert
	ver.in = 0
	ver.from = 0

	ver.next = make([]int, 0)
	ver.vis = false
	return ver
}

func dfs(g []Vert, t, cnt *int, bridg_list *map[int]bool, vt, p int) {
	g[vt].vis = true
	(*t) += 1
	g[vt].in = *t
	g[vt].from = *t
	//////////////////
	for i := 0; i <= len(g[vt].next) - 1; i++ {
		if (g[vt].next[i] == p){
			continue
		}
		if g[g[vt].next[i]].vis {
			g[vt].from = int(math.Min (float64(g[vt].from),float64(g[g[vt].next[i]].in)));
		} else {
			//рекурсивно вызываем
			dfs(g, t, cnt, bridg_list, g[vt].next[i], vt);

			g[vt].from = int(math.Min(float64(g[vt].from), float64(g[g[vt].next[i]].from)));
			if ( (g[vt].in < g[g[vt].next[i]].from) && (!(*bridg_list)[len(g)*g[vt].next[i]+vt])){
				(*cnt) += 1
				(*bridg_list)[vt * len(g) + g[vt].next[i]] = true
				(*bridg_list)[g[vt].next[i] * len(g) + vt] = true
				
			}
		}
	}
}

func main() {
	var n, m, cnt, t, a, b int
	fmt.Scan(&n);
	fmt.Scan(&m);


	var g []Vert
	for i := 0; i < n; i++ {
		g = append(g, get_vert())
	}
	//reading graph (g)
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		g[a].next = append(g[a].next, b)
		g[b].next = append(g[b].next, a)
	}

	t = 0
	cnt = 0
	bridg_list := make(map[int] bool)
	for i := 0; i < n; i++ {
		if (!g[i].vis) {
			dfs(g, &t, &cnt, &bridg_list, i, -1)
		}
	}
	fmt.Printf("%d\n", cnt)
}
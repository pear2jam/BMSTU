package main
import "fmt"



type Connection struct {
	d, wt int
}

type Vertex struct {
	vis bool
	next []Connection
}

type Edge struct {
	vert, id int
}

func get_g() Vertex{
	var to_return Vertex
	to_return.next = make([]Connection, 0)
	to_return.vis = false
	return to_return
}

func get_connection(a, b int) Connection{
	var to_return Connection
	to_return.d = a
	to_return.wt = b
	return to_return
}

func get_edge(a, b int) Edge{
	var to_return Edge
	to_return.vert = a
	to_return.id = b
	return to_return
}

func main() {
	var n, m, a, b, val int
	fmt.Scan(&n, &m)
	edge_list := make([]Edge, 0)

	res := 0

	var g []Vertex
	for i := 0; i < n; i++ {
		g = append(g, get_g())
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b, &val)
		g[a].next = append(g[a].next, get_connection(b, val))
		g[b].next = append(g[b].next, get_connection(a, val))
	}

	for i := 0; i < len(g[0].next); i++ {
		edge_list = append(edge_list, get_edge(0, i))
	}

	for ; len(edge_list) > 0; {
		min_edge := 0
		for i := 0; i < len(edge_list); i++ {
			cond := g[edge_list[min_edge].vert].next[edge_list[min_edge].id].wt > g[edge_list[i].vert].next[edge_list[i].id].wt
			if (cond) { min_edge = i }
		}
		v := edge_list[min_edge].vert
		to := g[v].next[edge_list[min_edge].id].d
		g[edge_list[min_edge].vert].vis = true
		
		if (!g[to].vis) {
			res += g[v].next[edge_list[min_edge].id].wt
			g[to].vis = true
			for i := 0; i < len(g[to].next); i++ {
				next_d := g[to].next[i].d
				if (g[next_d].vis) {
					continue
				}
				edge_list = append(edge_list, get_edge(to, i))
			}
		}
		
		edge_list = append(edge_list[ :min_edge], edge_list[min_edge+1:] ...)
	}

	fmt.Print(res)
}
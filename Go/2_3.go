package main


import "fmt"

type Vert struct {
	next []int
	vis bool
}

type Component struct {
	vertices, edges, min_vert int
	set_vert map[int]bool
	
}

func get_vert() Vert {
	var ver Vert
	ver.next = make([]int, 0)
	ver.vis = false
	return ver
}

func get_comp(n int) Component {
	var comp Component
	comp.vertices = 0
	comp.edges = 0
	comp.min_vert = n
	comp.set_vert = make(map[int]bool)
	return comp
}

func compare(a, b Component) bool {
	if (a.vertices != b. vertices) {
		return a.vertices < b.vertices
	}
	if (a.edges != b. edges) {
		return a.edges < b.edges
	}
	return a.min_vert > b.min_vert
}

func dfs(graph []Vert, comp []Component, v, id int) {
	if (comp[id].min_vert > v) {
		comp[id].min_vert = v
	}

	comp[id].vertices += 1
	comp[id].set_vert[v] = true
	graph[v].vis = true

	for i := 0; i < len(graph[v].next); i++ {
		comp[id].edges += 1
		if (graph[graph[v].next[i]].vis) {
			continue
		}
		dfs(graph, comp, graph[v].next[i], id);
	}
}

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m)
	var graph []Vert
	for i := 0; i < n; i++ {
		graph = append(graph, get_vert())
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		if (a != b) {
			graph[a].next = append(graph[a].next, b)
		}

		graph[b].next = append(graph[b].next, a)
	}

	c_list := make([]Component, 0)

	for i, id := 0, 0; i < n; i++ {
		if (!graph[i].vis) {
			c_list = append(c_list, get_comp(n))
			dfs(graph, c_list, i, id)
			id += 1
		}
	}
	comp_max := 0

	for i := 1; i < len(c_list); i++ {
		if (compare(c_list[comp_max], c_list[i])) {
			comp_max = i
		}
	}
	fmt.Print("graph {\n")

	for i := 0; i < len(graph); i++ {
		fmt.Print(i)
		if (c_list[comp_max].set_vert[i]) {
			fmt.Printf(" [color = red]")
		}
		fmt.Print("\n")
	}
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i].next); j++ {
			if (i > graph[i].next[j]) {
				continue
			}
			fmt.Print(i, " -- ", graph[i].next[j])
			if (c_list[comp_max].set_vert[i]) {
				fmt.Print(" [color = red]")
			}
			fmt.Print("\n")
		}
	}
	fmt.Print("}\n")
}
package main

import (
	"fmt"
	"sort"
)

type Component struct {
	bal int
	left map[int]bool
	right map[int]bool
}

type Vert struct {
	g int
	neighbors []int
	vis bool
}

func get_vert() Vert {
	var ver Vert
	ver.neighbors = make([]int, 0)
	ver.g = 0
	ver.vis = false
	return ver
}

func dfs(v int, pos *bool, graph []Vert, comp *Component) {
	graph[v].vis = true
	(*comp).bal += graph[v].g

	if (graph[v].g != -1) {
		(*comp).right[v] = true
	} else {
		(*comp).left[v] = true
	}

	for i := 0; i <= len(graph[v].neighbors) - 1; i++ {
		if (graph[v].g*graph[graph[v].neighbors[i]].g == 1) {
			*pos = false
			return
		}

		if (graph[graph[v].neighbors[i]].vis) {
			continue
		}

		graph[graph[v].neighbors[i]].g = -graph[v].g;
		dfs (graph[v].neighbors[i], pos, graph, comp);
	}
}

func comp(arr_a, arr_b []int) bool {
	if len(arr_a) < len(arr_b) { return true }
	if len(arr_a) > len(arr_b) { return false }
	for i := 0; i < len(arr_a); i++ {
		if arr_a[i] < arr_b[i] { return true }
		if arr_a[i] > arr_b[i] { return false }
	}
	return false
}

func get_mask(mask []bool, masks *[][]bool, n, n_cur int) {
	if (n == n_cur) {
		*masks = append(*masks, mask)
		return
	}
	false_mask := make([]bool, len(mask))
	true_mask := make([]bool, len(mask))
	//////////////////////////
	copy(false_mask, mask)
	copy(true_mask, mask)
	true_mask = append(true_mask, true)
	false_mask = append(false_mask, false)

	get_mask(true_mask, masks, n,n_cur + 1)
	get_mask(false_mask, masks, n, n_cur + 1)
}


func main() {
	var n int
	var s string
	fmt.Scan(&n)

	var graph []Vert
	for i := 0; i <= n - 1; i++ {
		graph = append(graph, get_vert())
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&s)
			if (s[0] == '+') {
				graph[i].neighbors = append(graph[i].neighbors, j)
			}
		}
	}

	components := make([]Component, 0)
	pos := true
	for i := 0; i < n; i++ {
		
		if (!graph[i].vis) {
			var c Component
			c.bal = 0
			c.left = make(map[int]bool)
			c.right = make(map[int]bool)
			graph[i].g = 1
			////////////////////////// dfs
			dfs(i, &pos, graph, &c)

			components = append(components, c)
		}
	}

	mask := make([]bool, 0)
	masks := make([][]bool, 0)
	
	get_mask(mask, &masks, len(components), 0)

	min_s := n
	if (!pos){
		fmt.Println("No solution")
		return
	}
	
	s_mask := make([]int, len(masks))
	for i := 0; i < len(masks); i++ {
		sum := 0
		for j := 0; j < len(components); j++ {
			mod := 1
			if (!masks[i][j]) { 
				mod = -1 
			}
			sum += mod * components[j].bal
		}

		if (sum < 0) {
			sum *= -1
		}
		s_mask[i] = sum

		if (min_s > sum) {
			min_s = sum
		}
	}

	minlist := make([]int, 0)

	for i := 0; i < len(masks); i++ {
		if (s_mask[i] == min_s) {
			set := make(map[int]bool)
			for j := 0; j < len(components); j++ {
				subset := components[j].right
				if (masks[i][j]) {
					subset = components[j].left
				}
				for key, value := range subset {
					set[key] = value
				}
			}
			i := 0
			keys := make([]int, len(set))
			for k, _ := range set {
				keys[i] = k
				i++
			}
			sort.Ints(keys)
			
			if (len(minlist) == 0 || comp(keys, minlist)) {
				minlist = keys
			}
		}
	}

	for _, vertex := range minlist {
		fmt.Print(vertex + 1, " ")
	}
}
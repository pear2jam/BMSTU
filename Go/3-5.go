package main

import "fmt"

type Pair struct {
	from, to int
}

func main() {
	var n, result_n, m, v int
	var s string
	io_map := make(map[string]int) // mapping to out

	corresp_m := make([][]int, 0)  // for Moore
	result, map_m := make([]int, 0), make([]int, 0)

	fmt.Scan(&n)
	v = 0
	// writing
	code_list := make([]string, n)
	for i := 0; i < n; i++{
		fmt.Scan(&code_list[i])
	}
	fmt.Scan(&result_n)
	res_list := make([]string, result_n)
	
	for i := 0; i < result_n; i++{
		fmt.Scan(&res_list[i])
		v += 1
	}
	for i := 0; i < result_n; i++{
		io_map[res_list[i]] = i
	}
	fmt.Scan(&m)
	corresp := make([][]Pair, m)
	for i := 0; i < m; i++{
		corresp[i] = make([]Pair, n)
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			fmt.Scan(&corresp[i][j].to)	
		}
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			fmt.Scan(&s)
			corresp[i][j].from = io_map[s]
		}
	}
	to_input := make([]map[int]bool, m)
	out_mapping := make([][]int, 0)
	for i := 0; i < m; i++{
		out_mapping = append(out_mapping, make([]int, n))
		to_input[i] = make(map[int]bool)
	}

	k := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++{ to_input[i][corresp[i][j].from] = true }
	}

	visited := make(map[int]bool)
	reverse_map := make(map[int]map[int]int)
	for i := 0; i < m; i++{
		reverse_map[i] = make(map[int]int)
		for from := range res_list {
			map_m = append(map_m, i)
			v += 1
			corresp_m = append(corresp_m, make([]int, n))
			for j := 0; j < n; j++{
				corresp_m[k][j] = corresp[i][j].to
				out_mapping[corresp[i][j].to][from] = k
			}
			reverse_map[i][from] = k
			k += 1
		}

		for from := range res_list {
			result = append(result, from)
		}
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			v_new := reverse_map[corresp[i][j].to][corresp[i][j].from]
			visited[v_new] = true
			for v, _ := range reverse_map[i] {
				corresp_m[reverse_map[i][v]][j] = v_new
			}
		}
	}
	fmt.Println("digraph {\nrankdir = LR")
	for i := 0; i < k; i++{
	    if (visited[i]) {
	    	fmt.Print(i, " [label = \"(", map_m[i], ",", res_list[result[i]], ")\"]\n")

		    for j := 0; j < len(corresp_m[i]); j++ {
			    fmt.Print(i, " -> ", corresp_m[i][j], " [label = \"", code_list[j], "\"]\n")
		    }
	    }
	}
	fmt.Println("}")
}
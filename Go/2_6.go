package main

import(
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func get_point(a, b float64) Point{
	var p Point
	p.x = a
	p.y = b
	return p
}

func main() {
	var n int
	var x, y, res float64

	fmt.Scan(&n)
	p := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		p[i] = get_point(x, y)
	}
	dist := make([][]float64, n)
	vis := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			dist[i][j] = math.Sqrt((p[i].x-p[j].x)*(p[i].x-p[j].x) + (p[i].y-p[j].y)*(p[i].y-p[j].y))
		}	
	}
	s_list := make([]int, n)
	min_list := make([]float64, n)
	
	for i := 0; i < n; i++ {
		s_list[i] = -1
		min_list[i] = 1e10
	}
	//start init
	min_list[0] = 0

	for i := 0; i < n; i++ {
		cur := -1

		for j := 0; j < n; j++ {
			if (vis[j]){
				continue
			}
			if ((cur == -1) || (min_list[cur] > min_list[j])) {
				cur = j
			}
		}
		vis[cur] = true;
		res += min_list[cur]

		for i := 0; i < n; i++ {
			//updating
			if (min_list[i] > dist[cur][i]) {
				s_list[i] = cur;
				min_list[i] = dist[cur][i];
			}
		}
	}
	fmt.Printf("%.2f", math.Round(res * 100) / 100.)
}
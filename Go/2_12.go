package main

import (
	"fmt"
	"errors"
	"container/heap"
)

type Place struct {
	score, d int
}

type heaplist []*Item

type Item struct {
	ind int
	order float64
	weight interface{}
}

type Point struct {
	x, y int
}

type PQueue struct {
	to_next map[interface{}]*Item
	heaplist *heaplist
}

func (p *PQueue) Len() int {
	return p.heaplist.Len()
}

func (item_heap *heaplist) Swap(i, j int) {
	(*item_heap)[i], (*item_heap)[j] = (*item_heap)[j], (*item_heap)[i]
	(*item_heap)[j].ind = j
	(*item_heap)[i].ind = i
}

func (item_heap *heaplist) Less(i, j int) bool {
	return (*item_heap)[j].order > (*item_heap)[i].order
}

func (item_heap *heaplist) Push(x interface{}) {
	it := x.(*Item)
	it.ind = len(*item_heap)
	*item_heap = append(*item_heap, it)
}

func (p *PQueue) Insert(v interface{}, order float64) {
	_, to_return := p.to_next[v]
	if to_return {return}
	newItem := &Item{
		order: order,
		weight: v,
	}
	heap.Push(p.heaplist, newItem)
	p.to_next[v] = newItem
}

func (p *PQueue) Pop() (interface{}, error) {
	if (len(*p.heaplist) == 0) {return nil, errors.New("Empty")}
	Item := heap.Pop(p.heaplist).(*Item)
	delete(p.to_next, Item.weight)
	return Item.weight, nil
}

func get_point(a, b int) Point{
	var p Point
	p.x = a
	p.y = b
	return p
}

func (item_heap *heaplist) Pop() interface{} {
	old := *item_heap
	l := len(old) - 1
	Item := old[l]
	*item_heap = old[0:l]
	return Item
}

func (item_heap *heaplist) Len() int {
	return len(*item_heap)
}

func main() {
	var n int
	fmt.Scan(&n)
	matrix := make([][]Place, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]Place, n)
		for j := 0; j < n; j++ {
			matrix[i][j].d = 1e8
			fmt.Scan(&matrix[i][j].score)
		}
	}
	m := n - 1
	q := PQueue{
		heaplist: &heaplist{},
		to_next: make(map[interface{}]*Item),		
	}

	q.Insert(get_point(0, 0), 0)
	a := make([]Point, 4)

	matrix[0][0].d = matrix[0][0].score

	a[0], a[1], a[2], a[3] = get_point(0,1), get_point(0,-1), get_point(1,0), get_point(-1,0)
	for ; q.Len() > 0; {
		p_inter, _ := q.Pop()
		p_inter2 := p_inter.(Point)
		for _, con := range a{
			p := get_point(p_inter2.x+con.x, p_inter2.y+con.y)
			if (p.x >= n || p.y >= n || p.x < 0 || p.y < 0) {continue}
			if (matrix[p_inter2.x][p_inter2.y].d + matrix[p.x][p.y].score < matrix[p.x][p.y].d){
				matrix[p.x][p.y].d = matrix[p_inter2.x][p_inter2.y].d + matrix[p.x][p.y].score

				q.Insert(p, (float64)(matrix[p.x][p.y].d))
			}
		}
	}
	
	fmt.Println(matrix[m][m].d)
}
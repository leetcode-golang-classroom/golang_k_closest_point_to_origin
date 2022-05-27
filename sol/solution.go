package sol

import (
	"container/heap"
)

type Point struct {
	x, y int
}

type MaxPointHeap []Point

func (h *MaxPointHeap) Len() int {
	return len(*h)
}

func (h *MaxPointHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxPointHeap) Less(i, j int) bool {
	return (*h)[i].EuclideanDist() > (*h)[j].EuclideanDist()
}

func (p *Point) EuclideanDist() int {
	return ((*p).x)*((*p).x) + ((*p).y)*((*p).y)
}

func (h *MaxPointHeap) Push(val interface{}) {
	*h = append(*h, val.(Point))
}
func (h *MaxPointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func kClosest(points [][]int, k int) [][]int {
	pLen := len(points)
	if k >= pLen {
		return points
	}
	pq := &MaxPointHeap{}
	heap.Init(pq)
	for _, point := range points {
		x := point[0]
		y := point[1]
		newPoint := Point{x: x, y: y}
		if pq.Len() < k {
			heap.Push(pq, newPoint)
		} else {
			if pq.Len() > 0 {
				top := heap.Pop(pq).(Point)
				if top.EuclideanDist() > newPoint.EuclideanDist() {
					heap.Push(pq, newPoint)
				} else {
					heap.Push(pq, top)
				}
			}
		}
	}

	result := [][]int{}
	for idx := 0; idx < k; idx++ {
		top := heap.Pop(pq).(Point)
		result = append(result, []int{top.x, top.y})
	}
	return result
}

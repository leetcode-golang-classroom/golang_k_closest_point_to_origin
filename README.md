# golang_k_closest_point_to_origin

Given an array of `points` where `points[i] = [xi, yi]` represents a point on the **X-Y** plane and an integer `k`, return the `k` closest points to the origin `(0, 0)`.

The distance between two points on the **X-Y** plane is the Euclidean distance (i.e., $\sqrt((x_1-x_2)^2+(y_1-y_2)^2)$).

You may return the answer in **any order**. The answer is **guaranteed** to be **unique** (except for the order that it is in).

## Examples
**Example 1:**

![https://assets.leetcode.com/uploads/2021/03/03/closestplane1.jpg](https://assets.leetcode.com/uploads/2021/03/03/closestplane1.jpg)

```
Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].

```

**Example 2:**

```
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.

```

**Constraints:**

- `1 <= k <= points.length <= $10^4$`
- $`-10^4$ < xi, yi < $10^4$`

## 解析

給定一個座標陣列 points ,

定義 Euclidean distance = $\sqrt((x_1-x_2)^2+(y_1-y_2)^2)$

 題目要求寫出一個演算法來找出前 k 個與原點的Euclidean distance 最小的座標

這題可以透過 min-heap 把這些點都放入

然後再 pop 出 k 個

這樣的時間複雜度是 O(nlogn)

空間複雜度是 O(n)

優化上面的作法是透過 MaxHeap

先輸入 k 個元素到 MaxHeap

第 k+1 個時， 取出目前 MaxHeap top 那個元素

如果這個元素比 目前最大的值小，則把最大的值移除掉，放入這個值

這樣的時間複雜度是 O(nlogk)

空間複雜度是 O(k)

## 程式碼
```go
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
```
## 困難點

1. 需要替換比大小的部份
2. 理解 MaxHeap 運作原理

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity
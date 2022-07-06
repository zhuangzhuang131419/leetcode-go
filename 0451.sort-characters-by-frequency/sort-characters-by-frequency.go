package leetcode

import (
	"container/heap"
	"fmt"
)

func frequencySort(s string) string {
	counts := map[int]int{}

	for _, w := range s {
		counts[int(w)]++
	}

	fmt.Printf("counts = %v", counts)
	h := &IHeap{}
	heap.Init(h)

	for k, v := range counts {
		heap.Push(h, [2]int{k, v})
	}

	result := ""

	for h.Len() > 0 {
		count := heap.Pop(h).([2]int)
		for i := 0; i < count[1]; i++ {
			result += string(count[0])
		}
	}

	return result

}

type IHeap [][2]int

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	result := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return result
}

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

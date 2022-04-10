package stream

/**
 * Created by frankieci on 2022/3/28 10:28 pm
 */

// elemHeap implements sort.Interface can be sorted by the routines in this sort package.
type elemHeap []Element

func (h *elemHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *elemHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

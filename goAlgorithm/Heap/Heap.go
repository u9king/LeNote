package Heap

import (
	. "../LinkList"
)

//Heap 顺序存储实现最小堆
type Heap []*ListNode

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
} // 最小堆

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}
func (h *Heap) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

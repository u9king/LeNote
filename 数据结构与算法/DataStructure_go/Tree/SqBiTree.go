package main

import "fmt"

type SqBiTree struct {
	data  []int
	count int
}

func NewSqBiTree(size int) *SqBiTree {
	return &SqBiTree{
		data:  make([]int, size),
		count: 0,
	}
}

func (bt *SqBiTree) Insert(value int) bool {
	if bt.count >= len(bt.data) {
		return false // 二叉树已满
	}

	bt.data[bt.count] = value
	bt.count++
	return true
}

func (bt *SqBiTree) Traverse() {
	for i := 0; i < bt.count; i++ {
		fmt.Printf("%d ", bt.data[i])
	}
	fmt.Println()
}

func main() {
	bt := NewSqBiTree(10)

	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(3)
	bt.Insert(4)
	bt.Insert(5)
	bt.Insert(6)
	bt.Insert(7)

	fmt.Println("Binary Tree:")
	bt.Traverse()
}
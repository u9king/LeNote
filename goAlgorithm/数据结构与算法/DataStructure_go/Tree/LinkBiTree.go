package main

import (
	"container/list"
	"fmt"
)

type BiNode struct {
	data  int
	left  *BiNode
	right *BiNode
}

func NewBiNode(value int) *BiNode {
	return &BiNode{
		data: value,
		left: nil,
		right: nil,
	}
}

func (node *BiNode) Insert(value int) {
	if value < node.data {
		if node.left == nil {
			node.left = NewBiNode(value)
		} else {
			node.left.Insert(value)
		}
	} else {
		if node.right == nil {
			node.right = NewBiNode(value)
		} else {
			node.right.Insert(value)
		}
	}
}

//PreOrderTraverse 先序遍历,根左右
func (node *BiNode) PreOrderTraverse() {
	if node != nil {
		fmt.Printf("%d ", node.data) // 访问根节点
		node.left.PreOrderTraverse() // 遍历左子树
		node.right.PreOrderTraverse() // 遍历右子树
	}
}

//InOrderTraverse 中序遍历
func (node *BiNode) InOrderTraverse() {
	if node != nil {
		node.left.InOrderTraverse()
		fmt.Printf("%d ", node.data)
		node.right.InOrderTraverse()
	}
}

//PostOrderTraverse 后序遍历
func (node *BiNode) PostOrderTraverse() {
	if node != nil {
		node.left.PostOrderTraverse()  // 遍历左子树
		node.right.PostOrderTraverse() // 遍历右子树
		fmt.Printf("%d ", node.data)    // 访问根节点
	}
}

//LevelOrderTraversal 层次遍历
func LevelOrderTraversal(root *BiNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		node := queue.Front().Value.(*BiNode)
		queue.Remove(queue.Front())

		fmt.Printf("%d ", node.data)

		if node.left != nil {
			queue.PushBack(node.left)
		}

		if node.right != nil {
			queue.PushBack(node.right)
		}
	}
}

//CreateBiTree 创建二叉树递归
func CreateBiTree(data []interface{}) *BiNode {
	if len(data) == 0 {
		return nil
	}

	// 从切片中取出根节点的数据
	rootData := data[0].(int)
	root := &BiNode{data: rootData}

	// 递归构建左子树和右子树
	if len(data) >= 2 {
		root.left = CreateBiTree(data[1].([]interface{}))
	}
	if len(data) >= 3 {
		root.right = CreateBiTree(data[2].([]interface{}))
	}

	return root
}

func main() {
	root := NewBiNode(5)
	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	fmt.Println("In-Order Traversal:")
	root.InOrderTraverse()
	fmt.Println()
}

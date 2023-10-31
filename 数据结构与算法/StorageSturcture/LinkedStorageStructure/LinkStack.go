package main

import "fmt"

//StackNode 定义单链表节点结构
type StackNode struct {
	data ElemType
	next *StackNode
}

//LinkStack 定义栈结构
type LinkStack struct {
	Head *StackNode
}

//InitStack 初始化一个空栈
func InitStack() *LinkStack {
	return &LinkStack{}
}

//Push 入栈操作
func (s *LinkStack) Push(item ElemType) {
	newNode := &StackNode{data: item, next: s.Head}
	s.Head = newNode
}

//Pop 出栈操作
func (s *LinkStack) Pop() ElemType {
	if s.Head == nil {
		return 0 // 返回默认值或者适当的错误处理
	}

	item := s.Head.data
	s.Head = s.Head.next
	return item
}

//Peek 获取栈顶元素
func (s *LinkStack) Peek() ElemType {
	if s.Head == nil {
		return 0 // 返回默认值或者适当的错误处理
	}
	return s.Head.data
}

//IsEmpty 判断栈是否为空
func (s *LinkStack) IsEmpty() bool {
	return s.Head == nil
}

func main() {
	stack := InitStack()

	// 入栈操作
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// 出栈操作
	fmt.Println("Pop:", stack.Pop())

	// 获取栈顶元素
	fmt.Println("Peek:", stack.Peek())

	// 判断栈是否为空
	fmt.Println("Is Empty:", stack.IsEmpty())
}
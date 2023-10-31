package main

import "fmt"

/*栈本身就是线性表，所以栈也有顺序存储和链式存储两种形式*/


//Stack 定义顺序栈结构
type Stack struct {
	data []interface{}
}

//NewStack 初始化一个空栈
func NewStack() *Stack {
	return &Stack{}
}

//Push 入栈操作
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

//Pop 出栈操作
func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

//Peek 获取栈顶元素
func (s *Stack) Peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

//IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

//Size 获取栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	stack := NewStack()

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

	// 获取栈的大小
	fmt.Println("Size:", stack.Size())
}
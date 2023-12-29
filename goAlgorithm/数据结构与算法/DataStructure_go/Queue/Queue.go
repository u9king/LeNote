package main

import (
"fmt"
)

//Queue 定义队列结构
type Queue struct {
	data []interface{}
}

//NewQueue 初始化一个空队列
func NewQueue() *Queue {
	return &Queue{}
}

//Enqueue 入队操作
func (q *Queue) Enqueue(item interface{}) {
	q.data = append(q.data, item)
}

//Dequeue 出队操作
func (q *Queue) Dequeue() interface{} {
	if len(q.data) == 0 {
		return nil
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

//Peek 获取队列头元素
func (q *Queue) Peek() interface{} {
	if len(q.data) == 0 {
		return nil
	}
	return q.data[0]
}

//IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

//Size 获取队列大小
func (q *Queue) Size() int {
	return len(q.data)
}

func main() {
	queue := NewQueue()

	// 入队操作
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// 出队操作
	fmt.Println("Dequeue:", queue.Dequeue())

	// 获取队列头元素
	fmt.Println("Peek:", queue.Peek())

	// 判断队列是否为空
	fmt.Println("Is Empty:", queue.IsEmpty())

	// 获取队列的大小
	fmt.Println("Size:", queue.Size())
}
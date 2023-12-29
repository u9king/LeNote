package main

import (
	"errors"
	"fmt"
)

type QueueNode struct {
	data interface{}
	next *QueueNode
}

type LinkQueue struct {
	front *QueueNode
	rear  *QueueNode
}

func NewLinkQueue() *LinkQueue {
	// 初始化一个空的链式队列
	node := &QueueNode{
		data: nil,
		next: nil,
	}
	return &LinkQueue{
		front: node,
		rear:  node,
	}
}

//IsEmpty 为空
func (q *LinkQueue) IsEmpty() bool {
	return q.front == q.rear
}

//Enqueue 入队
func (q *LinkQueue) Enqueue(data interface{}) {
	newNode := &QueueNode{
		data: data,
		next: nil,
	}
	q.rear.next = newNode
	q.rear = newNode
}

//Dequeue 出队
func (q *LinkQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	dequeuedData := q.front.next.data
	q.front.next = q.front.next.next
	if q.front.next == nil {
		q.rear = q.front
	}
	return dequeuedData, nil
}

func main() {
	queue := NewLinkQueue()

	// Enqueue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Dequeue
	item, err := queue.Dequeue()
	if err == nil {
		fmt.Printf("Dequeued item: %v\n", item)
	}

	// Check if Queue is empty
	if queue.IsEmpty() {
		fmt.Println("Queue is empty")
	}
}

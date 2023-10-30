package main

import (
	"errors"
	"fmt"
)

type SqQueue struct {
	data  []interface{}
	front int //头指针
	rear  int //尾指针
}

//NewSqQueue 创建循环队列
func NewSqQueue(capacity int) *SqQueue {
	return &SqQueue{
		data:  make([]interface{}, capacity),
		front: 0,
		rear:  0,
	}
}

//IsEmpty 判断队列是否为空
func (q *SqQueue) IsEmpty() bool {
	return q.front == q.rear
}

//IsFull 判断队列是否满
func (q *SqQueue) IsFull() bool {
	return (q.rear+1)%len(q.data) == q.front
}

//Length 循环队列长度
func (q *SqQueue) Length() int {
	maxQsize := len(q.data)
	return (q.rear - q.front + maxQsize) % maxQsize
}

//Enqueue 循环队列入队
func (q *SqQueue) Enqueue(item interface{}) error {
	if q.IsFull() {
		return errors.New("Queue is full")
	}
	q.data[q.rear] = item
	q.rear = (q.rear + 1) % len(q.data)
	return nil
}

//Dequeue 循环队列出队
func (q *SqQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	item := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	return item, nil
}

//GetHead 取队头元素
func (q *SqQueue) GetHead() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return q.data[q.front], nil
}

func main() {
	queue := NewSqQueue(5)

	// Enqueue
	err := queue.Enqueue(1)
	err = queue.Enqueue(2)
	err = queue.Enqueue(3)
	err = queue.Enqueue(4)
	err = queue.Enqueue(5)
	if err != nil {
		return
	}

	// Dequeue
	item, err := queue.Dequeue()
	if err == nil {
		fmt.Printf("Dequeued item: %v\n", item)
	}

	// Queue Length
	fmt.Printf("Queue Length: %d\n", queue.Length())

	// Enqueue more items
	queue.Enqueue(6)

	// Check if Queue is full
	if queue.IsFull() {
		fmt.Println("Queue is full")
	}
}

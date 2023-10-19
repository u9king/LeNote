package main

import "fmt"

type ElemType int

//DuLNode 双向链表节点
type DuLNode struct {
	Prior *DuLNode
	Data ElemType
	Next *DuLNode
}

//DuLinkList 双向链表头
type DuLinkList struct {
	Head *DuLNode
}

//InitDuList 初始化双链表
func InitDuList() *DuLinkList {
	return &DuLinkList{}
}

//DuListInsert 在指定节点之前插入新节点
func (l *DuLinkList) DuListInsert(i int, data ElemType) {
	cur := l.GetDuListNode(i)
	if cur == nil {
		fmt.Println("没这个结点")
		return
	}

	newNode := &DuLNode{Data: data, Prior: cur.Prior, Next: cur}
	cur.Prior.Next = newNode  //新节点接前面的尾
	cur.Prior = newNode  //新节点接后面的头
	return
}

// DuListDelete 删除指定节点
func (l *DuLinkList) DuListDelete(i int) {
	cur := l.GetDuListNode(i)
	if cur == nil {
		fmt.Println("没有找到匹配的节点")
		return
	}
	cur.Prior.Next = cur.Next
	cur.Next.Prior = cur.Prior
}

//GetDuListNode 按序号获取节点
func (l *DuLinkList) GetDuListNode(i int) *DuLNode {
	cur := l.Head.Next
	j := 1
	for cur != nil && j < i {
		cur = cur.Next
		j++
	}
	return cur
}
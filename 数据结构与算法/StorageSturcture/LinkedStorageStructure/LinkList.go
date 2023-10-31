package main

import "fmt"

// ListNode 带头结点的单链表
type ListNode struct {
	Data ElemType
	Next *ListNode
}

type LinkList struct {
	Head *ListNode
}

//InitList 初始化单链表
func InitList() *LinkList {
	return &LinkList{}
}

//CreateListL 头插法建立单链表
func CreateListL(data []ElemType) *LinkList {
	l := InitList() // 初始化单链表
	for _, value := range data {
		newNode := &ListNode{Data: value}
		newNode.Next = l.Head.Next
		l.Head.Next = newNode
	}
	return l
}

//CreateListR 尾插法建立单链表
func CreateListR(data []ElemType) *LinkList {
	l := InitList()
	tail := l.Head
	for _, value := range data {
		newNode := &ListNode{Data: value}
		tail.Next = newNode
		tail = newNode
	}
	return l
}

//ListEmpty 判断链表是否为空
func (l *LinkList) ListEmpty() bool {
	if l.Head == nil {
		return true
	}
	return false
}

//Length 计算链表长度
func (l *LinkList) Length() int {
	cur := l.Head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

//GetElem 按序号取值
func (l *LinkList) GetElem(i int) ElemType {
	cur := l.Head
	j := 1
	for cur != nil && j < i {
		cur = cur.Next
		j++
	}
	if cur != nil {
		return cur.Data
	}
	return -1 // 未找到匹配的节点
}

//LocateELem 按值查找
func (l *LinkList) LocateELem(e ElemType) *ListNode {
	cur := l.Head
	for cur != nil && cur.Data != e {
		cur = cur.Next
	}
	return cur
}

//ListInsert 指定位置插入
func (l *LinkList) ListInsert(i int, e ElemType) {
	if i < 1 {
		fmt.Println("插入位置不合法")
		return
	}
	cur := l.Head
	j := 0
	for cur != nil && j < i-1 {
		cur = cur.Next
		j++
	}
	if cur == nil || j > i-1 {
		fmt.Println("插入位置不合理")
		return
	}
	s := &ListNode{Data: e}
	s.Next = cur.Next
	cur.Next = s
	return
}

//ListDelete 删除指定位置的元素(无法删除首元节点)
func (l *LinkList) ListDelete(i int) {
	if i < 1 {
		fmt.Println("删除位置不合法")
		return
	}
	cur := l.Head
	j := 0
	for cur.Next != nil && j < i-1 {
		cur = cur.Next
		j++
	}
	if cur == nil || j > i-1 {
		fmt.Println("删除位置不合理")
		return
	}
	cur.Next = cur.Next.Next
	return
}



func main() {
	l := InitList() // 初始化单链表

	data := []ElemType{1, 2, 3, 4, 5, 6, 7}

	for _, elem := range data {
		// 在链表中添加元素
		newNode := &ListNode{Data: elem, Next: nil}
		if l.Head == nil {
			l.Head = newNode
		} else {
			current := l.Head
			for current.Next != nil {
				current = current.Next
			}
			current.Next = newNode
		}
	}

	//fmt.Println(l.GetElem(2))
	l.ListDelete(1)

}

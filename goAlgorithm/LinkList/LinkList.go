package LinkList

//将方法 (l *ListNode) Traverse() 放在 ListNode 类型上是 Go 中一种常见的做法，
//称为“接收者(receiver)”或者“方法”。这种方式将方法与特定类型关联，允许在该类型的实例上直接调用方法。
//关于你的问题，是的，使用接收者函数 (l *ListNode) Traverse() 是一种常见的方式，特别适合链表节点类型，
//因为它直接作用于节点实例。这种方法能够以更直观、更具可读性的方式处理链表节点的遍历和操作。
//但是，在某些情况下，你可能更倾向于不将方法与特定类型直接绑定，而是像 Length(head *ListNode) int 这样将链表头部作为参数传递给函数。
//这种方法更加通用，因为它可以处理不同的链表实例，而不仅仅局限于特定类型上。
//所以，选择将方法作为接收者或者将链表头部作为参数传递取决于你的需求和设计偏好。
//在简化和直接性方面，使用接收者函数可能更有优势，但在更通用的情况下，将链表头部作为参数传递可能更为合适。

import "fmt"

//type Element interface{}

// ListNode 单链表结点
type ListNode struct {
	Val  int       //数据
	Next *ListNode //下个指针
}

//InitList 初始化单链表
func InitList() *ListNode {
	return &ListNode{}
}

func CreateList(Val []int) *ListNode {
	l := InitList() // 初始化单链表
	tail := l
	for _, value := range Val {
		if l.Val == 0 {
			l.Val = value
			continue
		}
		node := &ListNode{Val: value}
		tail.Next = node
		tail = node
	}
	return l
}

//Length 计算链表长度
func Length(head *ListNode) int {
	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

//Traverse 遍历列表打印输出
func Traverse(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print("->")
		}
		cur = cur.Next
	}
}

//CreateListL 头插法建立单链表
func CreateListL(Val []int) *ListNode {
	l := InitList() // 初始化单链表
	for _, value := range Val {
		newNode := &ListNode{Val: value}
		newNode.Next = l.Next
		l.Next = newNode
	}
	return l
}

//CreateListR 尾插法建立单链表
func CreateListR(Val []int) *ListNode {
	l := InitList()
	tail := l
	for _, value := range Val {
		newNode := &ListNode{Val: value}
		tail.Next = newNode
		tail = newNode
	}
	return l
}

//ListEmpty 判断链表是否为空
func ListEmpty(head *ListNode) bool {
	if head == nil {
		return true
	}
	return false
}

//GetElem 按序号取值
func GetElem(head *ListNode, i int) int {
	cur := head
	j := 1
	for cur != nil && j < i {
		cur = cur.Next
		j++
	}
	if cur != nil {
		return cur.Val
	}
	return -1 // 未找到匹配的节点
}

//LocateELem 按值查找
func LocateELem(head *ListNode, e int) *ListNode {
	cur := head
	for cur != nil && cur.Val != e {
		cur = cur.Next
	}
	return cur
}

//ListInsert 指定位置插入
func ListInsert(head *ListNode, i int, e int) {
	if i < 1 {
		fmt.Println("插入位置不合法")
		return
	}
	cur := head
	j := 0
	for cur != nil && j < i-1 {
		cur = cur.Next
		j++
	}
	if cur == nil || j > i-1 {
		fmt.Println("插入位置不合理")
		return
	}
	s := &ListNode{Val: e}
	s.Next = cur.Next
	cur.Next = s
	return
}

//ListDelete 删除指定位置的元素(无法删除首元节点)
func ListDelete(head *ListNode, i int) {
	if i < 1 {
		fmt.Println("删除位置不合法")
		return
	}
	cur := head
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

//ReverseList 反转链表
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		nxt := cur.Next //nxt保存反转前下一个节点
		cur.Next = pre  //反转cur节点的方向
		pre = cur       //向下一个节点移动
		cur = nxt       //向下一个节点移动
	}
	return pre
}

func main() {
	l := InitList() // 初始化单链表

	Val := []int{1, 2, 3, 4, 5, 6, 7}

	for _, elem := range Val {
		// 在链表中添加元素
		newNode := &ListNode{Val: elem, Next: nil}
		if l == nil {
			l = newNode
		} else {
			current := l
			for current.Next != nil {
				current = current.Next
			}
			current.Next = newNode
		}
	}

	//fmt.Println(l.GetElem(2))
	ListDelete(l, 1)

}

package LinkList

import "fmt"

//type Element interface{}

// ListNode 单链表结点
type ListNode struct {
	Val int   //数据
	Next *ListNode //下个指针
}

//LinkList 单链表
type LinkList struct {
	Head *ListNode //头结点
}

//InitList 初始化单链表
func InitList() *LinkList {
	head := &ListNode{} //建立头结点，值需要赋值函数单独处理  //todo思考是否采用哨兵节点的方式处理
	return &LinkList{Head: head}
}

func CreateList(Val []int) *LinkList {
	l := InitList() // 初始化单链表
	tail := l.Head
	for _, value := range Val {
		if l.Head.Val == 0 {
			l.Head.Val = value
			continue
		}
		node := &ListNode{Val: value}
		tail.Next = node
		tail = node
	}
	return l
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

//Traverse 遍历列表打印输出
func (l *LinkList) Traverse() {
	cur := l.Head
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print("->")
		}
		cur = cur.Next
	}
}

//CreateListL 头插法建立单链表
func CreateListL(Val []int) *LinkList {
	l := InitList() // 初始化单链表
	for _, value := range Val {
		newNode := &ListNode{Val: value}
		newNode.Next = l.Head.Next
		l.Head.Next = newNode
	}
	return l
}

//CreateListR 尾插法建立单链表
func CreateListR(Val []int) *LinkList {
	l := InitList()
	tail := l.Head
	for _, value := range Val {
		newNode := &ListNode{Val: value}
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

//GetElem 按序号取值
func (l *LinkList) GetElem(i int) int {
	cur := l.Head
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
func (l *LinkList) LocateELem(e int) *ListNode {
	cur := l.Head
	for cur != nil && cur.Val != e {
		cur = cur.Next
	}
	return cur
}

//ListInsert 指定位置插入
func (l *LinkList) ListInsert(i int, e int) {
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
	s := &ListNode{Val: e}
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

	Val := []int{1, 2, 3, 4, 5, 6, 7}

	for _, elem := range Val {
		// 在链表中添加元素
		newNode := &ListNode{Val: elem, Next: nil}
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

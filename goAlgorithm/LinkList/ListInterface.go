package LinkList



type AllLinkList interface {
	Length() int
	GetElem(i int) Element
	LocateELem(e Element) *ListNode
	ListInsert(i int, e Element)
	ListDelete(i int)
}



//func main() {
//	l := InitList() // 初始化单链表
//
//	data := []Element{1, 2, 3, 4, 5, 6, 7}
//
//	for _, elem := range data {
//		// 在链表中添加元素
//		newNode := &ListNode{Data: elem, Next: nil}
//		if l.Head == nil {
//			l.Head = newNode
//		} else {
//			current := l.Head
//			for current.Next != nil {
//				current = current.Next
//			}
//			current.Next = newNode
//		}
//	}
//
//	//fmt.Println(l.GetElem(2))
//	l.ListDelete(1)
//
//}

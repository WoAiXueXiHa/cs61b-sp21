package LList

type IntNode struct {
	Value int
	Next  *IntNode
}

type LList struct {
	dummy *IntNode
	Size  int
}

func InitLList(val int) *LList {
	Node := &IntNode{
		Value: val,
		Next:  nil,
	}
	return &LList{
		dummy: Node,
		Size:  0,
	}
}

func Size(list *LList) int {
	return list.Size
}

func addFirst(list *LList, val int) {
	Node := &IntNode{
		Value: val,
	}

	Node.Next = list.dummy.Next
	list.dummy.Next = Node
	list.Size++
}

func addLast(list *LList, val int) {
	Node := &IntNode{
		Value: val,
	}

	cur := list.dummy
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = Node

	list.Size++
}

func GetFirst(list *LList) *IntNode {
	return list.dummy.Next
}

func Get(list *LList, val int) (*IntNode, bool) {
	if list.Size == 0 {
		return nil, false
	}
	for cur := list.dummy; cur != nil; cur = cur.Next {
		if cur.Value == val {
			return cur, true
		}
	}
	return nil, false
}

func GetByIndex(list *LList, index int) (int, bool) {
	if list.Size == 0 {
		return -1, false
	}
	if index < 0 || index > Size(list) {
		return -1, false
	}

	cnt := -1
	for cur := list.dummy; cur != nil; cur = cur.Next {
		if cnt == index {
			return cur.Value, true
		}
		cnt++
	}
	return -1, false
}

func DelFirst(list *LList) {
	if list.Size == 0 {
		return
	}
	list.dummy.Next = list.dummy.Next.Next
	list.Size--
}

func DelLast(list *LList) {
	if list.Size == 0 {
		return
	}
	cur := list.dummy
	var prev *IntNode
	for ; cur.Next != nil; cur = cur.Next {
		prev = cur
	}
	prev.Next = prev.Next.Next
	list.Size--
}

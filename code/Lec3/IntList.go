package IntList

type IntList struct {
	Next  *IntList
	Value int
}

// 递归
func RecursionSize(list *IntList) int {
	if list == nil {
		return 0
	}
	// 我当前节点长度是1，加上剩余链表长度就是全链表长度
	return 1 + RecursionSize(list.Next)
}

// 迭代
func IterativeSize(list *IntList) int {
	cnt := 0
	for cur := list; cur != nil; cur = cur.Next {
		cnt++
	}
	return cnt
}

// ========================================================================================

// 迭代法根据索引获取节点值
func GetByIteration(list *IntList, index int) (int, bool) {
	if list == nil {
		return -1, false
	}
	// 10 20 30
	cnt := 0
	for cur := list; cur != nil; cur = cur.Next {
		if cnt == index {
			return cur.Value, true
		}
		cnt++
	}
	return -1, false
}

// 递归法根据索引获取节点值
func GetByRecursion(list *IntList, index int) (int, bool) {
	if list == nil || index < 0 {
		return -1, false
	}

	if index == 0 {
		return list.Value, true
	}
	// 想要得到索引 n 对应节点的值，就是找当前节点之后索引为 n-1 节点的值
	// 索引：0	 1	 2	 3	  4
	// 链表：1 ->2 ->3 ->4 -> 5
	// 假设找索引 2 节点的值：
	// 当前节点是 1，下一个节点：2 ->3 ->4 -> 5，找索引为 1
	// 当前节点是 1， 下一个节点：3 ->4 -> 5，找索引为 0
	return GetByIteration(list.Next, index-1)
}

// ========================================================================================

// 迭代法，不修改原链表，把每个节点的值加 delta，返回新链表
func IncrListByIteration(list *IntList, delta int) *IntList {
	if list == nil {
		return nil
	}
	dummy := &IntList{}
	cur := dummy

	for tmp := list; tmp != nil; tmp = tmp.Next {
		node := &IntList{
			Value: tmp.Value + delta,
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

// 递归法，不修改原链表，把每个节点的值加 delta，返回新链表
func IncrListByRecursion(list *IntList, delta int) *IntList {
	// 给每个节点创建新的值，接着处理当前节点之后的链表
	if list == nil {
		return list
	}

	return &IntList{
		Value: list.Value + delta,
		Next:  IncrListByRecursion(list.Next, delta),
	}
}

// 迭代法，修改原链表，把每个节点值加 delta
func DIncrListByIteration(list *IntList, delta int) {
	if list == nil {
		return
	}
	for cur := list; cur != nil; cur = cur.Next {
		cur.Value += delta
	}
}

// 递归法，修改原链表，把每个节点值加 delta
func DIncrListByRecuision(list *IntList, delta int) {
	if list == nil {
		return
	}
	list.Value += delta
	DIncrListByRecuision(list.Next, delta)
}

// 总结一下递归的思路，链表天生就是递归的结构，我知道了当前节点，那么就能直到下一个节点
// 所以我只用处理当前节点，下一个节点交给下一层处理

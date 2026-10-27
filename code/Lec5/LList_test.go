package LList

import "testing"

// want ...int 是变长参数，调用时可以传入任意数量的期望节点值。
func assertList(t *testing.T, list *LList, want ...int) {
	// 标记当前函数为测试辅助函数，让失败位置指向调用 assertList 的那一行。
	t.Helper()

	cur := GetFirst(list)
	for index, wantValue := range want {
		if cur == nil {
			t.Fatalf("node %d is nil, want value %d", index, wantValue)
		}
		if cur.Value != wantValue {
			t.Errorf("node %d value = %d, want %d", index, cur.Value, wantValue)
		}
		cur = cur.Next
	}

	if cur != nil {
		t.Errorf("list contains extra node with value %d", cur.Value)
	}
	if Size(list) != len(want) {
		t.Errorf("size = %d, want %d", Size(list), len(want))
	}
}

func assertGetByIndex(t *testing.T, list *LList, index, want int, wantOK bool) {
	t.Helper()

	got, ok := GetByIndex(list, index)
	if got != want || ok != wantOK {
		t.Errorf("GetByIndex(list, %d) = (%d, %v), want (%d, %v)",
			index, got, ok, want, wantOK)
	}
}

func TestAddFirst(t *testing.T) {
	list := InitLList(0)
	assertList(t, list)

	addFirst(list, 10)
	assertList(t, list, 10)

	addFirst(list, 20)
	assertList(t, list, 20, 10)
}

func TestAddLast(t *testing.T) {
	list := InitLList(0)
	assertList(t, list)

	addLast(list, 10)
	assertList(t, list, 10)

	addLast(list, 20)
	assertList(t, list, 10, 20)
}

func TestGet(t *testing.T) {
	list := InitLList(0)
	addLast(list, 10)
	addLast(list, 20)

	node, ok := Get(list, 10)
	if !ok || node == nil {
		t.Fatal("Get(list, 10) did not find 10")
	}
	if node.Value != 10 {
		t.Errorf("node.Value = %d, want 10", node.Value)
	}

	node, ok = Get(list, 30)
	if ok || node != nil {
		t.Errorf("Get(list, 30) = (%+v, %v), want (nil, false)", node, ok)
	}

	emptyList := InitLList(0)
	node, ok = Get(emptyList, 10)
	if ok || node != nil {
		t.Errorf("Get(emptyList, 10) = (%+v, %v), want (nil, false)", node, ok)
	}
}

func TestGetByIndex(t *testing.T) {
	list := InitLList(0)
	addLast(list, 10)
	addLast(list, 20)

	assertGetByIndex(t, list, 0, 10, true)
	assertGetByIndex(t, list, 1, 20, true)
	assertGetByIndex(t, list, -9, -1, false)
	assertGetByIndex(t, list, 9, -1, false)

	emptyList := InitLList(0)
	assertGetByIndex(t, emptyList, 0, -1, false)
}

func TestDelFirst(t *testing.T) {
	list := InitLList(0)
	addLast(list, 10)
	addLast(list, 20)

	DelFirst(list)
	assertList(t, list, 20)

	DelFirst(list)
	assertList(t, list)

	DelFirst(list)
	assertList(t, list)
}

func TestDelLast(t *testing.T) {
	list := InitLList(0)
	addLast(list, 10)
	addLast(list, 20)

	DelLast(list)
	assertList(t, list, 10)

	DelLast(list)
	assertList(t, list)

	DelLast(list)
	assertList(t, list)
}

/*
原始测试版本（保留用于和精简版对照，不参与运行）

func TestAddFirst(t *testing.T) {
	list := InitLList(0)

	if Size(list) != 0 {
		t.Errorf("new list size = %d, want 0\n", Size(list))
	}

	addFirst(list, 10)
	if Size(list) != 1 {
		t.Errorf("after addFirst(10), size = %d, want 1\n", Size(list))
	}

	first := GetFirst(list)
	if first == nil {
		t.Errorf("after addFirst(10), first node is nil")
	}
	if first.Value != 10 {
		t.Errorf("after addFirst(10), first value = %d, want 10\n", first.Value)
	}

	addFirst(list, 20)
	if Size(list) != 2 {
		t.Errorf("after addFirst(20), size = %d, want 2", Size(list))
	}

	first = GetFirst(list)
	if first == nil {
		t.Fatal("after addFirst(20), first node is nil")
	}
	if first.Value != 20 {
		t.Errorf("first value = %d, want 20", first.Value)
	}
	if first.Next == nil {
		t.Fatal("first.Next is nil, want node containing 10")
	}
	if first.Next.Value != 10 {
		t.Errorf("second value = %d, want 10", first.Next.Value)
	}
}

func TestAddLast(t *testing.T) {
	list := InitLList(0)

	addLast(list, 10)
	if Size(list) != 1 {
		t.Errorf("after addLast(10), size = %d, want 1", Size(list))
	}

	first := GetFirst(list)
	if first == nil {
		t.Errorf("first node is nil")
	}
	if first.Value != 10 {
		t.Errorf("after addLast(10), first value = %d, want 10", first.Value)
	}

	addLast(list, 20)
	if Size(list) != 2 {
		t.Errorf("after addLast(20), size = %d, want 2", Size(list))
	}
	first = GetFirst(list)
	if first.Next == nil {
		t.Errorf("after addLast(20), second node is nil")
	}
	if first.Next.Value != 20 {
		t.Errorf("after addLast(20), second value = %d, want 20", first.Next.Value)
	}
}

func TestGet(t *testing.T) {
	list := InitLList(0)
	addFirst(list, 20)
	addFirst(list, 10)

	node, ok := Get(list, 10)
	if !ok {
		t.Fatal("Get(list, 10) did not find 10")
	}
	if node == nil {
		t.Fatal("Get(list, 10) returned nil node")
	}
	if node.Value != 10 {
		t.Errorf("node.Value = %d, want 10", node.Value)
	}

	node1, ok1 := Get(list, 30)
	if ok1 {
		t.Errorf("Get found unexpected node: %+v", node1)
	}
	if node1 != nil {
		t.Errorf("node = %+v, want nil", node1)
	}

	emptyList := InitLList(-1)
	node2, ok2 := Get(emptyList, 10)
	if ok2 {
		t.Error("Get on empty list returned true, want false")
	}
	if node2 != nil {
		t.Errorf("Get on empty list returned %+v, want nil", node2)
	}
}

func TestGetByIndexFirst(t *testing.T) {
	list := InitLList(0)
	addFirst(list, 20)
	addFirst(list, 10)

	got, ok := GetByIndex(list, 0)
	if !ok {
		t.Fatal("GetByIndex(list, 0) returned false, want true")
	}
	want := 10
	if got != want {
		t.Errorf("GetByIndex(list, 0) = %d, want %d", got, want)
	}
}

func TestGetByIndexSecond(t *testing.T) {
	list := InitLList(0)
	addFirst(list, 20)
	addFirst(list, 10)

	got, ok := GetByIndex(list, 1)
	if !ok {
		t.Fatal("GetByIndex(list, 1) returned false, want true")
	}
	if got != 20 {
		t.Errorf("GetByIndex(list, 1) = %d, want 20", got)
	}
}

func TestGetByIndexOutOfRange(t *testing.T) {
	list := InitLList(0)
	addLast(list, 10)
	addLast(list, 20)

	got1, ok1 := GetByIndex(list, -9)
	if ok1 {
		t.Errorf("GetByIndex returned (%d, true), want(-1, false)", got1)
	}
	if got1 != -1 {
		t.Errorf("GetByIndex value = %d, want -1", got1)
	}

	got2, ok2 := GetByIndex(list, 9)
	if ok2 {
		t.Errorf("GetByIndex returned (%d, true), want(-1, false)", got2)
	}
	if got2 != -1 {
		t.Errorf("GetByIndex value = %d, want -1", got2)
	}
}

func TestDelFirst(t *testing.T) {
	list := InitLList(0)
	addLast(list, 10)
	addLast(list, 20)

	DelFirst(list)
	gotSize := Size(list)
	if gotSize != 1 {
		t.Errorf("size after DelFirst = %d, want 1", gotSize)
	}
	first := GetFirst(list)
	if first == nil {
		t.Fatal("first node is nil, want node containing 20")
	}
	if first.Value != 20 {
		t.Errorf("first value after DelFirst = %d, want 20", first.Value)
	}

	DelFirst(list)
	size := Size(list)
	if size != 0 {
		t.Errorf("size after DelFirst = %d, want 0", size)
	}
	first1 := GetFirst(list)
	if first1 != nil {
		t.Fatal("first node was not nil, want nil node")
	}

	DelFirst(list)
	if Size(list) != 0 {
		t.Errorf("size after deleting from empty list = %d, want 0", Size(list))
	}
	if GetFirst(list) != nil {
		t.Errorf("first node after deleting from empty list = %+v, want nil", GetFirst(list))
	}
}

func TestDelLast(t *testing.T) {
	list := InitLList(0)
	addLast(list, 10)
	addLast(list, 20)

	DelLast(list)
	if Size(list) != 1 {
		t.Errorf("after delLast, size = %d, want 1", Size(list))
	}
	first := GetFirst(list)
	if first == nil {
		t.Fatal("after delLast, first node was nil, want one node")
	}
	if first.Value != 10 {
		t.Errorf("after delLast, first node value = %d, want 10", first.Value)
	}
	if first.Next != nil {
		t.Fatal("after delLast, second node was not nil, want nil")
	}

	DelLast(list)
	if Size(list) != 0 {
		t.Errorf("after delLast, size = %d, want 0", Size(list))
	}
	first = GetFirst(list)
	if first != nil {
		t.Fatal("after delLast, first node was not nil, want nil")
	}

	DelLast(list)
	if Size(list) != 0 {
		t.Errorf("after delLast, size = %d, want 0", Size(list))
	}
	first = GetFirst(list)
	if first != nil {
		t.Fatal("after delLast, first node was not nil, want nil")
	}
}
*/

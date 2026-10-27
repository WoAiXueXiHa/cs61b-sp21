package IntList

import "testing"

func TestSize(t *testing.T) {
	tests := []struct {
		name string
		list *IntList
		want int
	}{
		{
			name: "empty list",
			list: nil,
			want: 0,
		},
		{
			name: "one node",
			list: &IntList{
				Value: 10,
			},
			want: 1,
		},
		{
			name: "three node",
			list: &IntList{
				Value: 1,
				Next: &IntList{
					Value: 2,
					Next: &IntList{
						Value: 3,
					},
				},
			},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Run("recursive", func(t *testing.T) {
				got := RecursionSize(test.list)

				if got != test.want {
					t.Errorf("RecursionSize() = %d, want %d\n", got, test.want)
				}
			})
		})

		t.Run("iterative", func(t *testing.T) {
			got := IterativeSize(test.list)

			if got != test.want {
				t.Errorf("IterationSize() = %d, want %d\n", got, test.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	list := &IntList{
		Value: 10,
		Next: &IntList{
			Value: 20,
			Next: &IntList{
				Value: 30,
			},
		},
	}

	tests := []struct {
		name      string
		list      *IntList
		index     int
		wantValue int
		wantOk    bool
	}{
		{
			name:      "first node",
			list:      list,
			index:     0,
			wantValue: 10,
			wantOk:    true,
		},
		{
			name:      "middle node",
			list:      list,
			index:     1,
			wantValue: 20,
			wantOk:    true,
		},
		{
			name:      "last node",
			list:      list,
			index:     2,
			wantValue: 30,
			wantOk:    true,
		},
		{
			name:      "negative index",
			list:      list,
			index:     -1,
			wantValue: -1,
			wantOk:    false,
		},
		{
			name:      "out of range",
			list:      list,
			index:     3,
			wantValue: -1,
			wantOk:    false,
		},
		{
			name:      "empty list",
			list:      nil,
			index:     0,
			wantValue: -1,
			wantOk:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Run("Recursive", func(t *testing.T) {
				gotValue, gotOk := GetByRecursion(test.list, test.index)

				if gotValue != test.wantValue || gotOk != test.wantOk {
					t.Errorf("GetByRecursion(index=%d) = (%d, %v), want(%d, %v)\n",
						test.index,
						gotValue,
						gotOk,
						test.wantValue,
						test.wantOk)
				}
			})

			t.Run("Iterative", func(t *testing.T) {
				gotValue, gotOk := GetByRecursion(test.list, test.index)

				if gotValue != test.wantValue || gotOk != test.wantOk {
					t.Errorf("GetByIteration(index=%d) = (%d, %v), want(%d, %v)\n",
						test.index,
						gotValue,
						gotOk,
						test.wantValue,
						test.wantOk)
				}
			})
		})
	}
}

func equalList(a, b *IntList) bool {
	for a != nil && b != nil {
		if a.Value != b.Value {
			return false
		}
		a = a.Next
		b = b.Next
	}

	return a == nil && b == nil
}

func listValue(list *IntList) []int {
	var vals []int

	for cur := list; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Value)
	}
	return vals
}

func TestIncr(t *testing.T) {
	list := &IntList{
		Value: 10,
		Next: &IntList{
			Value: 20,
			Next: &IntList{
				Value: 30,
			},
		},
	}

	tests := []struct {
		name     string
		list     *IntList
		wantList *IntList
		delta    int
	}{
		{
			name:     "nil list",
			list:     nil,
			wantList: nil,
			delta:    1,
		},
		{
			name: "list",
			list: list,
			wantList: &IntList{
				Value: 11,
				Next: &IntList{
					Value: 21,
					Next: &IntList{
						Value: 31,
					},
				},
			},
			delta: 1,
		},
	}

	for _, test := range tests {
		t.Run("Recursive", func(t *testing.T) {
			gotList := IncrListByRecursion(test.list, test.delta)
			original := &IntList{
				Value: 10,
				Next: &IntList{
					Value: 20,
					Next: &IntList{
						Value: 30,
					},
				},
			}

			if !equalList(gotList, test.wantList) {
				t.Errorf("IncrListByRecursion() = %v, want %v\n",
					listValue(gotList),
					listValue((test.wantList)))
			}

			if test.list != nil && !equalList(test.list, original) {
				t.Errorf("original list was modified: got %v, want %v",
					listValue(test.list),
					listValue(original))
			}
		})

		t.Run("Iterative", func(t *testing.T) {
			gotList := IncrListByIteration(test.list, test.delta)
			original := &IntList{
				Value: 10,
				Next: &IntList{
					Value: 20,
					Next: &IntList{
						Value: 30,
					},
				},
			}

			if !equalList(gotList, test.wantList) {
				t.Errorf("IncrListByIteration() = %v, want %v\n",
					listValue(gotList),
					listValue((test.wantList)))
			}

			if test.list != nil && !equalList(test.list, original) {
				t.Errorf("original list was modified: got %v, want %v",
					listValue(test.list),
					listValue(original))
			}
		})
	}
}

func makeList(values ...int) *IntList {
	var list *IntList

	for i := len(values) - 1; i >= 0; i-- {
		list = &IntList{
			Value: values[i],
			Next:  list,
		}
	}

	return list
}

func TestDIncr(t *testing.T) {
	implementations := []struct {
		name string
		run  func(*IntList, int)
	}{
		{
			name: "iterative",
			run:  DIncrListByIteration,
		},
		{
			name: "recursive",
			run:  DIncrListByRecuision,
		},
	}

	tests := []struct {
		name   string
		values []int
		delta  int
		want   []int
	}{
		{
			name:   "nil list",
			values: nil,
			delta:  1,
			want:   nil,
		},
		{
			name:   "positive delta",
			values: []int{10, 20, 30},
			delta:  1,
			want:   []int{11, 21, 31},
		},
		{
			name:   "negative delta",
			values: []int{10, 20, 30},
			delta:  -5,
			want:   []int{5, 15, 25},
		},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					// 每个子测试创建全新的链表
					list := makeList(test.values...)
					wantList := makeList(test.want...)

					implementation.run(list, test.delta)

					if !equalList(list, wantList) {
						t.Errorf(
							"got %v, want %v",
							listValue(list),
							listValue(wantList),
						)
					}
				})
			}
		})
	}
}

# Go 单元测试实战笔记

## 1. 核心思路

测试是在定义函数的行为契约：

```text
准备输入和期望结果（Arrange）
→ 调用函数（Act）
→ 比较实际结果与期望结果（Assert）
```

不要只判断“结果发生了变化”，必须判断结果是否**准确等于期望值**。

## 2. 基本约定

- 测试文件以 `_test.go` 结尾。
- 测试函数格式为 `func TestXxx(t *testing.T)`。
- `got` 表示实际结果，`want` 表示期望结果。
- 每个函数至少考虑：正常、最小、边界、非法输入。

## 3. 单案例模板

```go
func TestFunction(t *testing.T) {
	input := ...
	want := ...

	got := Function(input)

	if got != want {
		t.Errorf("Function(%v) = %v, want %v", input, got, want)
	}
}
```

## 4. 表驱动测试模板

```go
func TestFunction(t *testing.T) {
	tests := []struct {
		name  string
		input InputType
		want  ResultType
	}{
		{"normal", inputValue, wantValue},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Function(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
```

同一套案例测试多个实现时，把函数放进表中：

```go
implementations := []struct {
	name string
	run  func(*IntList, int)
}{
	{"iterative", DIncrListByIteration},
	{"recursive", DIncrListByRecursion},
}
```

破坏式函数会修改输入，因此每个子测试必须创建全新的测试数据。

## 5. 链表测试要点

链表不能用 `==` 比较内容；`==` 比较的是指针地址。应同时遍历两条链表，比较每个值，并确认它们同时到达 `nil`。

```go
func equalList(a, b *IntList) bool {
	for a != nil && b != nil {
		if a.Value != b.Value {
			return false
		}
		a, b = a.Next, b.Next
	}
	return a == nil && b == nil
}
```

- `IncrList`：新链表值正确、原链表不变、新旧链表不共享节点。
- `DIncrList`：修改后的值准确、节点地址保持不变、`nil` 输入不 panic。

`makeList` 不是必需功能，只是避免重复创建链表，并保证每个修改型测试拿到独立输入。

## 6. 今天踩过的坑

1. `GetByIteration` 移动了节点指针，却忘记 `cnt++`，导致只能找到索引 `0`。
2. “迭代版”测试误调用了递归函数，测试名称与实际调用不一致。
3. `cur := dummy.Next` 得到 `nil`，随后访问 `cur.Next` 引发空指针 panic；尾指针应从 `dummy` 开始。
4. 循环中反复使用相同的 `t.Run("Iterative", ...)`，Go 自动生成 `#01`；应先用案例名分组，再嵌套实现名。
5. 两个破坏式实现复用同一条链表，第二个实现拿到的是已修改数据；每个子测试必须重新构造输入。
6. 只比较“修改前后不相等”不够；错误修改同样会不相等，必须与精确的 `want` 比较。

## 7. Panic 排查顺序

```text
确定失败的子测试
→ 看 panic 类型
→ 跳过 runtime/testing 标准库堆栈
→ 找第一条属于项目的 .go 文件和行号
→ 检查该行哪个指针可能是 nil
→ 向上追踪它在哪里赋值
→ 修复后重新运行全部测试
```

例如：

```text
IntList.go:76  cur.Next = node
```

先检查 `cur`，再追踪到：

```go
cur := dummy.Next // dummy.Next 的零值是 nil
```

## 8. 常用命令

```powershell
go test ./...
go test -v ./...
go test -run TestFunction ./...
go test -run "TestFunction/subtest" ./...
go test -count=1 ./...
```

修复后必须重新运行原失败测试，再运行全部测试。这就是回归测试。

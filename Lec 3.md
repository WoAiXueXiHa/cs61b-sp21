## 1. Bits 和 Go 的基本类型

bit 是是计算机存储信息的最小单位，只有 0 和 1，8 bit 构成 1 byte

bit 本身是没有含义的，假设内存中有：

```text
01000001
```

它可以被解释成：

- 无符号整数 65
- ASCII 字符：A
- 开关的状态

**数据的意义不在 bit 本身，而在解释规则**，类型就是解释规则的一部分：

```go
var number uint8 = 65
var letter byte = 65

fmt.Println(number)        // 65
fmt.Printf("%c\n", letter) // A
```

`byte` 是 `uint8` 的别名，二者底层模式相同，但是可以用不同的输出方式代表数字 or 字符

对于无符号整数，`uint8` 使用全部 8 个 bit 表示非负整数：

```text
2^7 2^6 2^5 2^4 2^3 2^2 2^1 2^0
 0	0	0	0	0	 0   0	 0	= 	0
 0	0	0	0	0	 0   0	 1	= 	1
... 
 1	1	1	1	1	 1	 1	 1  =   255
```

8 个 bit 共有 $2^8 = 256$ 种不同组合，范围就是 \[ 0\sim 2^8-1=255 \]

对于有符号整数，采用**二进制补码**表示，以 `int8` 为例：

```text
符号 2^6 2^5 2^4 2^3 2^2 2^1 2^0
 0	 0	 0	 0	 0	  0   0	  0	 = 	0
 0	 0	 0	 0	 0	  0   0	  1  = 	1
 1   0   0   0   0    0   0   0  = -128
 1	 1	 1	 1	 1	  1	  1	  1  =  -1
```

范围是：\[ -2^7\sim 2^7-1 \]

**为什么 11111111 表示 -1？**首先必须明确：**8 个 bit 只有 256 种状态**，`int8`和`uint8`都用这256种状态

如果把所有的状态排成一个圈：

![image-20261024203941717](https://gitee.com/binary-whispers/pic/raw/master///20261024203944675.png)

最关键的交界处：

```text
01111111  →  10000000
   127          128       按 uint8 解释
   127         -128       按 int8 解释
```

- 对于 `uint8`，编号从 `0` 一直增加到 `255`
- 对于 `int8`，前半圈编号不变，后半圈改成负数

同时有个有趣的结论：**`uint8` - 256 = `int8`**，相同的 0 1 串被解释成不同的语义



## 2. 变量、内存和复制

```go
var x int32 = 42
```

这段代码的意识是：**声明一个 int32 类型的叫做 x 的盒子，把 42 按照比特位排列塞到这个盒子里**

什么是赋值？

```go
x := int32(42)
y := x
```

执行第一行：x 的盒子里装着 42，执行第二行：把 x 里的值 copy 到 y 这个盒子，两个盒子的类型一样

后续执行：

```go
y = 100
```

修改了 y，不会影响 x，因为 x 和 y 是两个不同的盒子

**= 是赋值符号，是把右边的值塞进左边的盒子**



## 3. 指针、地址和共享修改

假设有：

```go
x := 10
```

内存布局是这样的：

![image-20261024205411536](https://gitee.com/binary-whispers/pic/raw/master///20261024205413155.png)

使用 `&x` 可以得到 `x` 的地址：

```go
p ：= &x
```

![image-20261024205613402](https://gitee.com/binary-whispers/pic/raw/master///20261024205615025.png)

`p` 保存的是地址 `0x1000`，因此它能够找到 `x`

`&`作用是取得变量的地址，把这个地址保存到一个变量中

`*`作用是读取变量保存的地址，访问这个地址中的值

Go 提供了内置函数 `new`：

```go
p := new(T)
```

它会给一个 `T` 类型的值准备空间，初始化成零值，然后返回 T 类型的指针

对于结构体：

```go
type Student struct {
    Name string
    Age  int
}

s := Student{
    Name: "Alice",
    Age:  20,
}
```

那么 `s` 就装了结构体对应的内容

执行：

```go
t := s
```

会复制整个结构体

![image-20261024210258643](https://gitee.com/binary-whispers/pic/raw/master///20261024210300467.png)

随后对`t`的修改不会影响`s`，因为这是两个不同的盒子

如果现在不复制结构体，而是取得它的地址：

```go
p := &s
```

![image-20261024210432862](https://gitee.com/binary-whispers/pic/raw/master///20261024210434407.png)

通过指针修改字段：

```go
p.Age = 30		// (*p).Age = 30
```

做个总结：**赋值始终复制值；指针赋值复制的是地址值**，两个指针变量独立，但是指向的是同一个目标



## 3. 函数传参

**Go 函数参数永远按值传递**，调用函数时，会把实参的值复制给形参

```go
func addOne(n int) {
    n++
}

func main() {
    x := 10
    addOne(x)
}
```

`x` 是调用函数提供的实参

`n` 是函数定义接收值的形参

![image-20261024211828246](https://gitee.com/binary-whispers/pic/raw/master///20261024211832875.png)

`n` 和 `x` 是两个独立的盒子，修改 `n` 影响不到 `x`，所以 `x` 最后还是 10

那如何真正修改调用者的变量值呢？传地址

```go
func addOne(n *int) {
    (*n)++
}

func main() {
    x := 10
    addOne(&x)
    fmt.Println(x) // 11
}
```

![image-20261024212206754](https://gitee.com/binary-whispers/pic/raw/master///20261024212210115.png)

再看个例子，深刻理解，Go 函数传参都是值传递

```go
func redirect(n *int) {
    y := 20
    n = &y
}

func main() {
    x := 10
    p := &x

    redirect(p)

    fmt.Println(*p) // 10
}
```

![image-20261024214531316](https://gitee.com/binary-whispers/pic/raw/master///20261024214533275.png)



## 5. 数组和切片

**Go 的数组是值类型。数组变量直接包含所有元素，赋值和传参会复制整个数组，而切片是对数组的一段描述，赋值和传参会复制三个字段**

数组的的长度也是数组的一部分，Go 中的数组是定长的

我们重点来看切片：

```go
a := []int{10, 20, 30}
b := a		// 复制了 a 的三个字段

b[0] = 99	// a 和 b 共享一份底层数组
```

![image-20261024215333324](https://gitee.com/binary-whispers/pic/raw/master///20261024215337276.png)

切片核心内容请移步：[slice](https://code-learn-build-evolve.vercel.app/posts/go-slice/)
package base

import (
	"fmt"
	"unsafe"
)

// 全局变量声明
var x, y int
var (
	m bool
	n string
)
var t1, t2 float32 = 1, 2
var o, p = 123, "456"

// LENGTH 全局声明常量
const LENGTH int = 10
const WIDTH int = 5
const PI, SIZE = 3.14, 1024

// 常量用于枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

// 常量可以使用内置函数
const (
	A = "abc"            // abc
	B = len(A)           // 3
	C = unsafe.Sizeof(A) //16
)

// iota，特殊常量，可以认为是一个可以被编译器修改的常量
const (
	I1 = iota // 0
	I2 = iota // 1
	I3 = iota // 2
)

const (
	I4 = iota // 0
	I5        // 1
	I6        // 2
)

func TestConst() {
	fmt.Println(LENGTH, WIDTH, PI, SIZE, Unknown, Female, Male, A, B, C, I1, I2, I3, I4, I5, I6)
}

func TestIota() {

	const (
		a = iota // 0
		b        // 1
		c        // 2
		d = "ha" // "ha" iota += 1
		e        // "ha" iota += 1
		f = 100  // 100 iota +=1
		g        // 100 iota +=1
		h = iota // 7,恢复计数
		i        // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		// iota从0开始加
		m = 1 << iota // 1
		n = 3 << iota // 6
		x             // 12
		y             // 24
	)
	fmt.Println(m, n, x, y)
}

func TestGlobalVariable() {
	fmt.Println(x, y, m, n, t1, t2, o, p)
}

func TestTempVariable() {

	// 第一种方式: var 变量名 变量类型
	var a int
	fmt.Println(a)
	// 给变量赋值
	a = 3
	fmt.Println(a)

	// 第二种方式【多个变量一起声明】:var 变量名1,变量名2 变量类型
	var b, c string
	fmt.Println(b, c)

	// 第三种方式【类型推导】: var 变量名 = 变量值
	var e = 3.14
	fmt.Println(e)

	// 第四种方式【类型推导】: var 变量名1,变量名2 变量类型 = 变量1值,变量2值
	var f, g float64 = 25.80, 30.60
	fmt.Println(f, g)

	// 第五种方式【类型推导(简化)】: 变量名 := 变量值
	h := false
	fmt.Println(h)
	i, j := 3, "hello go"
	fmt.Println(i, j)
}

func TestDefaultValueVariable() {

	var a int
	var b bool
	var c string
	var d float64
	var e []int
	var f [5]int
	var g map[string]string
	var h *string
	var i chan int
	var j error
	var k interface{} // nil
	// var m func(string) int nil
	fmt.Printf("a=[%d],b=[%t],c=[%s],d=[%f],e=[%v],f=[%v],g=[%v],h=[%p],i=[%v],j=[%v],k=[%v]", a, b, c, d, e, f, g, h, i, j, k)
}

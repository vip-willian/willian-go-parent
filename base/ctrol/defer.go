package _ctrol

import "fmt"

func TestDeferStack() {
	var numsIndex [5]struct{}
	for i := range numsIndex {
		defer fmt.Println(i)
	}
}

func TestDeferClosePackage() {
	var numsIndex [5]struct{}
	for i := range numsIndex {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func TestDeferCopy() {

	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制

	x += 10
	y += 100
	println("x =", x, "y =", y)
}

func TestDeferCase() {
	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              // 无效！
	defer fmt.Println(recover()) // 无效！
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！
		}()
	}()
	panic("test panic")
}

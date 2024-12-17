package base

import (
	"fmt"
)

func TestAnonymousFunction() {
	// 方式一: 定义一个匿名函数并将其赋值给变量add
	add := func(a, b int) int {
		return a + b
	}
	result1 := add(1, 2)
	fmt.Printf("1 + 2 = %d\n", result1)
	// 方式二: 定义匿名函数并直接调用
	result2 := func(a, b int) int { return a - b }(10, 20)
	fmt.Printf("10 + 20 = %d\n", result2)
}

func TestFuncVariable() {

	// 定义一个匿名函数并将其赋值给变量add
	add := func(a, b int) int {
		return a + b
	}
	// 调用匿名函数
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	// 在函数内部使用匿名函数
	multiply := func(x, y int) int {
		return x * y
	}

	product := multiply(4, 6)
	fmt.Println("4 * 6 =", product)

	// 将匿名函数作为参数传递给其他函数
	// operation func(int, int) int 函数作为参数
	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8)
	fmt.Println("2 + 8 =", sum)

	// 也可以直接在函数调用中定义匿名函数
	// 函数可以作为变量
	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)
	fmt.Println("10 - 4 =", difference)

	// 函数可以作为一个结构体的字段
	funcAsField := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	fmt.Println(funcAsField.fn())

	// 函数可以作为集合的元素
	funcAsSliceElement := []func(x int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	fmt.Println(funcAsSliceElement[0](100))

	// --- 函数可以作为通道的返回 ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	fmt.Println((<-fc)())
}

// TestFuncClosePackage 支持匿名函数，可作为闭包
func TestFuncClosePackage() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

// 闭包函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// TestValueTran 值传递
func TestValueTran() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	swapValue(a, b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

// 引用传递
func TestRefTran() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	swapRef(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

// 函数返回两个数的最大值
func Max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 多个返回值
func Test(x, y string) (string, string) {
	return y, x
}

// 值传递
func swapValue(x, y int) int {

	temp := x
	x = y
	y = temp
	return temp
}

// 引用传递
func swapRef(x, y *int) int {

	// 拿到x指针地址上的值
	temp := *x
	// x指针地址上的值 = y指针地址上的值
	*x = *y
	// y指针地址上的值 = x指针地址上的值
	*y = temp
	return temp
}

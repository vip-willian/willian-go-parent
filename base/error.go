package base

import (
	"errors"
	"fmt"
)

/*
Go 的错误处理主要围绕以下机制展开：
	1、error 接口：标准的错误表示。
		type error interface {
			Error() string
		}
	2、显式返回值：通过函数的返回值返回错误。
	3、自定义错误：可以通过标准库或自定义的方式创建错误。
	4、panic 和 recover：处理不可恢复的严重错误。
*/

func TestError() {
	// 定义一个异常
	error := errors.New("this is an error")
	fmt.Println(error)

	// 显示指定异常返回
	_, error = Sqrt(-1)
	if error != nil {
		fmt.Println(error)
	}
	// 自定义异常
	_, error = divide(10, 0)
	if error != nil {
		// error.Is() 判断是什么类型异常
		if errors.Is(error, &DivideError{}) {
			fmt.Println(error) // 输出：cannot divide 10 by 0
		}
		// errors.As() 将错误转换为特定类型以便进一步处理
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// 显示的返回错误
		return 0, errors.New("math: square root of negative number")
	}
	return f, nil
}

// 自定义错误
type DivideError struct {
	Dividend int
	Divisor  int
}

// 实现Error接口方法
func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

/*
*
panic 和 recover
Go 的 panic 用于处理不可恢复的错误，recover 用于从 panic 中恢复。
panic:

	导致程序崩溃并输出堆栈信息。
	常用于程序无法继续运行的情况。

recover:

	捕获 panic，避免程序崩溃。
*/
func TestPanicAndRecover() {
	fmt.Println("开始执行")
	safeFunction()
	fmt.Println("执行完成")
}

func safeFunction() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到错误信息:", err)
		}
	}()
	panic("出现运行错误")
}

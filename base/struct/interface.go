package _struct

import (
	"fmt"
)

// Shape 定义新接口
type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

func TestInterface() {
	c := Circle{radius: 5}
	var s Shape = &c

	if perimeter, error := s.Perimeter(); error == nil {
		fmt.Printf("圆的周长:%f\n", perimeter)
	}
	if area, error := s.Area(); error == nil {
		fmt.Printf("圆的面积:%f\n", area)
	}

	r := Rectangle{length: 20, width: 15.2}
	s = &r

	if perimeter, error := s.Perimeter(); error == nil {
		fmt.Printf("长方形的周长:%f\n", perimeter)
	}
	if area, error := s.Area(); error == nil {
		fmt.Printf("长方形的面积:%f\n", area)
	}

	// 动态类型
	var i interface{} = 42
	fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)
}

// TestEmptyInterface
// 空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。
// 任意类型都实现了空接口。
// 常用于需要存储任意类型数据的场景，如泛型容器、通用参数等。
func TestEmptyInterface() {
	printValue(42)          // int
	printValue("hello")     // string
	printValue(3.14)        // float64
	printValue([]int{1, 2}) // slice
}

type Reader interface {
	read() string
}
type Writer interface {
	write(string)
}

// ReadWriter 接口组合
type ReadWriter interface {
	Reader
	Writer
}

// TestInterfaceTypeConverter 接口类型转换
func TestInterfaceTypeConverter() {
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

// 判断空接口类型
func TestJudgeEmptyInterfaceType() {
	printValue(42)
	printValue("hello")
	printValue(3.14)
}

func printValue(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("Int: %v, Type: %T\n", v, v)
	case string:
		fmt.Printf("String: %v, Type: %T\n", v, v)
	default:
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
}

package base

import (
	"fmt"
	"strconv"
)

// TestInt8ToInt64 整数int8 -> 整数int64
func TestInt8ToInt64() {

	var i int8 = 20
	var j int64 = int64(i)
	fmt.Printf("整数8位【%d】转换为整数64位为【%d】\n", i, j)
}

// TestInt64ToInt8 整数int64 -> 整数int8
func TestInt64ToInt8() {

	var i int64 = 9999232
	var j int8 = int8(i)
	fmt.Printf("整数64位【%d】转换为整数8位为【%d】\n", i, j)
}

// TestIntToFloat 整数 -> float64
func TestIntToFloat() {

	i := 20
	j := float64(i)
	fmt.Printf("整数【%d】转换为浮点数为【%f】\n", i, j)
}

// TestFloatToInt float64 -> 整数
func TestFloatToInt() {

	i := 20.00
	j := int(i)
	fmt.Printf("浮点数【%f】转换为整数为【%d】\n", i, j)
}

// TestStringToInt string -> 整数
func TestStringToInt() {

	// 第一种方式
	var str = "10"
	i, _ := strconv.Atoi(str)
	fmt.Printf("字符串【'%s'】转换为整数为【%d】\n", str, i)
}

// TestIntToString 整数 -> string
func TestIntToString() {

	num := 123
	str := strconv.Itoa(num)
	fmt.Printf("整数【%d】转换为字符串为【'%s'】\n", num, str)
}

// TestStringToFloat string -> float
func TestStringToFloat() {

	var str = "10.12"
	num, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("字符串【'%s'】转换为浮点数为【%f】\n", str, num)
}

// TestFloatToString float -> string
func TestFloatToString() {

	num := 3.14
	str := strconv.FormatFloat(num, 'f', 2, 64)
	fmt.Printf("浮点数【%f】转换为字符串为【'%s'】\n", num, str)
}

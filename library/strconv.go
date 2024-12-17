package library

import (
	"fmt"
	"strconv"
)

// string to int
func TestStringToInt() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) // type:int value:100
	}
}

// int to string
func TestIntToString() {
	i := 200
	s := strconv.Itoa(i)
	fmt.Printf("type:%T value:%#v\n", i, s)
}

// string to bool
func TestStringToBool() {
	s1 := "1"
	s2 := "0"
	s3 := "t"
	s4 := "f"
	s5 := "T"
	s6 := "F"
	s7 := "true"
	s8 := "false"
	s9 := "True"
	s10 := "False"
	s11 := "TRUE"
	s12 := "FALSE"

	b1, _ := strconv.ParseBool(s1)
	b2, _ := strconv.ParseBool(s2)
	b3, _ := strconv.ParseBool(s3)
	b4, _ := strconv.ParseBool(s4)
	b5, _ := strconv.ParseBool(s5)
	b6, _ := strconv.ParseBool(s6)
	b7, _ := strconv.ParseBool(s7)
	b8, _ := strconv.ParseBool(s8)
	b9, _ := strconv.ParseBool(s9)
	b10, _ := strconv.ParseBool(s10)
	b11, _ := strconv.ParseBool(s11)
	b12, _ := strconv.ParseBool(s12)

	fmt.Println(b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12)
}

// bool to string
func TestBoolToString() {
	b := false
	fmt.Println(strconv.FormatBool(b))
}

/*
bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
*/
// string to float
func TestStringToFloat() {

	s := "3.1415"
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Println(f)
}

// float to string
func TestFloatToString() {
	f := 3.1415
	fmt.Println(strconv.FormatFloat(f, 'E', -1, 64))
}

/*
base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
*/
func TestParseInt() {

	s := "-2"
	i, _ := strconv.ParseInt(s, 10, 64)
	fmt.Println(i)
}

// ParseUint() 用于无符号整型
func TestParseUnit() {

	s := "2"
	i, _ := strconv.ParseUint(s, 10, 64)
	fmt.Println(i)
}

// func FormatUint(i uint64, base int) string
// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。
func TestFormatInt() {
	fmt.Println(strconv.FormatInt(-20, 10))
}

func TestFormatUint() {
	fmt.Println(strconv.FormatUint(22, 10))
}

/*
函数将浮点数表示为字符串并返回。
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
*/
func TestFormatFloat() {
	fmt.Println(strconv.FormatFloat(22.86, 'E', -1, 64))
}

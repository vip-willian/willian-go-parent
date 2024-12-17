package library

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Print
// 将内容输出到系统的标准输出
func TestPrint() {
	fmt.Print("在终端打印该信息。")
	name := "willian"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}

// Fprint
// 将内容输出到一个io.Writer接口类型的变量w中,常用这个函数往文件中写入内容。
func TestFprint() {
	// 写入标准输出
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")

	// 写入文件
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "Hello Go"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写入信息：%s", name)
}

// Sprint
// 将传入的数据生成并返回一个字符串。
func TestSprint() {

	name := "willian"
	age := 18
	s1 := fmt.Sprint("willian", "hello")
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("willian")
	fmt.Println(s1, s2, s3)
}

// Errorf
// 根据format参数生成格式化字符串并返回一个包含该字符串的错误。
func TestErrorf() {

	i := 0
	err := fmt.Errorf("这是一个错误%d", i)
	fmt.Println(err)
}

// 格式化占位符
func TestFormatter() {

	fmt.Println("***通用占位符***")
	fmt.Printf("%v\n", 100) // 值的默认格式
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"willian"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o) // 类似%v，但输出结构体时会添加字段名
	fmt.Printf("%#v\n", o) // 值的Go语法表示
	fmt.Printf("%T\n", o)  // 打印值的类型
	fmt.Printf("100%%\n")  // 百分号

	fmt.Println("***布尔类型***")
	b := true
	fmt.Printf("%t\n", b) // 布尔型

	fmt.Println("***数字类型***")
	n := 65
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%c\n", n) // unicode码值
	fmt.Printf("%d\n", n) // 十进制
	fmt.Printf("%o\n", n) // 八进制
	fmt.Printf("%x\n", n) // 十六进制，使用a-f
	fmt.Printf("%X\n", n) // 十六进制，使用A-F

	fmt.Println("***浮点类型***")
	f := 12.34
	fmt.Printf("%b\n", f)    // 无小数部分、二进制指数的科学计数法，如-123456p-78
	fmt.Printf("%e\n", f)    // 科学计数法，如-1234.456e+78
	fmt.Printf("%E\n", f)    // 科学计数法，如-1234.456E+78
	fmt.Printf("%f\n", f)    // 有小数部分但无指数部分 默认宽度，默认精度
	fmt.Printf("%9f\n", f)   // 宽度9，默认精度
	fmt.Printf("%.2f\n", f)  // 默认宽度，精度2
	fmt.Printf("%9.2f\n", f) // 宽度9，精度2
	fmt.Printf("%9.f\n", f)  // 宽度9，精度0
	fmt.Printf("%g\n", f)    // 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	fmt.Printf("%G\n", f)    // 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	fmt.Println("***字符换和[]byte类型***")
	s := "中国"
	fmt.Printf("%s\n", s) // 直接输出字符串或者[]byte
	fmt.Printf("%q\n", s) // 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	fmt.Printf("%x\n", s) // 每个字节用两字符十六进制数表示（使用a-f)
	fmt.Printf("%X\n", s) // 每个字节用两字符十六进制数表示（使用A-F）

	fmt.Println("***指针类型***")
	p := 18
	fmt.Printf("%p\n", &p) // 表示为十六进制，并加上前导的0x
	fmt.Printf("%#p\n", &p)

	fmt.Println("***其他类型***")
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)
}

// 获取输入Scan
func TestScan() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("获取输入结果 name:%s age:%d married:%t\n", name, age, married)
}

// 按指定格式获取输入
func TestScanf() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s,2:%d,3:t", &name, &age, &married)
	fmt.Printf("获取输入结果 name:%s age:%d married:%t\n", name, age, married)
}

// 读取标准输入
func TestReadStdIn() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容: ")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

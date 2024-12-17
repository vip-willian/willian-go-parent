package _struct

import (
	"errors"
	"fmt"
	"math"
)

// 定义一个结构体
type Books struct {
	Title   string
	Author  string
	Subject string
	BookId  int
}

// 支持匿名字段
type Address struct {
	province string
	city     string
}

func (a *Address) ToString() string {
	return fmt.Sprintf("Address: %p, %v", a, a)
}

type User struct {
	userId int
	name   string
	age    uint8
	Address
}

func (u *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

type Circle struct {
	radius float64
}

// 定义方法
func (c *Circle) Area() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("半径不能为负数")
	}
	return math.Pi * c.radius * c.radius, nil
}

func (c *Circle) Perimeter() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("半径不能为负数")
	}
	return 2 * math.Pi * c.radius, nil
}

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Area() (float64, error) {
	if r.length < 0 || r.width < 0 {
		return 0, errors.New("长或宽不能为负数")
	}
	return r.length * r.width, nil
}

func (r *Rectangle) Perimeter() (float64, error) {
	if r.length < 0 || r.width < 0 {
		return 0, errors.New("长或宽不能为负数")
	}
	return 2 * (r.length + r.width), nil
}

func TestStruct() {

	// 第一种方式: 直接按顺序赋值
	book1 := Books{"go语言", "willian", "值得学习的好语言", 1001}
	fmt.Printf("book1=%#v\n", book1)
	// 第二种方式: 按key-value赋值
	book2 := Books{title: "java语言", author: "willian", subject: "都是很不错的", bookId: 1002}
	fmt.Printf("book2=%#v\n", book2)
	// 第三种方式: 给指定的key-value赋值
	book3 := Books{title: "python语言", bookId: 1003}
	fmt.Printf("book3=%#v\n", book3)
	// 第四种方式: 使用new关键字 返回的是指针类型
	book4 := new(Books)

	// 读取数据
	printBookOfValue("book1", book1)
	printBookOfPointer("book2", &book2)
	printBookOfValue("book3", book3)
	printBookOfPointer("book4", book4)
}

func TestStructPointer() {

	// book1 指针类型
	book1 := &Books{"go语言", "willian", "值得学习的好语言", 1001}
	fmt.Printf("book1=%#v", book1)

	printBookOfPointer("book1", book1)
	printBookOfValue("book1", *book1)
}

// 匿名字段
func TestAnonymousField() {
	user := User{1, "willian", 20, Address{"浙江省", "杭州市"}}
	fmt.Println(user.ToString())
	fmt.Println(user.Address.ToString())
}

func printBookOfValue(bookStr string, book Books) {

	fmt.Printf("通过值对象打印【%s】的数据内容\n", bookStr)
	fmt.Printf("%s title : %s\n", bookStr, book.title)
	fmt.Printf("%s author : %s\n", bookStr, book.author)
	fmt.Printf("%s subject : %s\n", bookStr, book.subject)
	fmt.Printf("%s bookId : %d\n", bookStr, book.bookId)
}

func printBookOfPointer(bookStr string, book *Books) {

	fmt.Printf("通过指针打印【%s】的数据内容\n", bookStr)
	fmt.Printf("%s title : %s\n", bookStr, book.title)
	fmt.Printf("%s author : %s\n", bookStr, book.author)
	fmt.Printf("%s subject : %s\n", bookStr, book.subject)
	fmt.Printf("%s bookId : %d\n", bookStr, book.bookId)
}

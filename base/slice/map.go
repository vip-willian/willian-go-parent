package _slice

import (
	"fmt"
	"willian-go-parent/base/struct"
)

func TestCreateMap() {
	// 第一种方式 【使用make函数】
	var map1 map[string]int = make(map[string]int)
	fmt.Printf("map1 = %#v\n", map1)

	// 第二种方式 【使用make函数】指定初始容量
	var map2 = make(map[string]int, 10)
	fmt.Printf("map2 = %#v\n", map2)

	// 第三种方式 直接指定并初始化
	map3 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Printf("map3 = %#v\n", map3)
}

func TestGetSetMap() {
	bookMap := map[int]*_struct.Books{
		// 可以省去&Book
		1001: {Title: "go 语言", Author: "tom", Subject: "这是一门不错的语言", BookId: 1001},
		1002: {Title: "java 语言", Author: "jack", Subject: "java老古董", BookId: 1002},
		1003: {Title: "python", Author: "hi", Subject: "一门解释性语言", BookId: 1003},
	}
	book1 := bookMap[1001]
	fmt.Printf("修改前 book 的值: %#v\n", book1)

	book1.Author = "jerry"
	// bookMap[1001] = book1

	book2 := bookMap[1001]
	fmt.Printf("修改后 book 的值: %#v\n", book2)
}

func TestDeleteIteratorMap() {
	bookMap := map[int]*_struct.Books{
		// 可以省去&Book
		1001: {Title: "go 语言", Author: "tom", Subject: "这是一门不错的语言", BookId: 1001},
		1002: {Title: "java 语言", Author: "jack", Subject: "java老古董", BookId: 1002},
		1003: {Title: "python", Author: "hi", Subject: "一门解释性语言", BookId: 1003},
	}
	delete(bookMap, 1002)

	// 遍历map
	for key, book := range bookMap {
		fmt.Printf("map[%d] = %#v\n", key, book)
	}
}

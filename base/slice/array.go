package _slice

import "fmt"

// TestCreateArray var arrayName [size]dataType
func TestCreateArray() {
	// 第一种方式: 初始化数组，默认值全为0
	var nums1 [5]int
	// 第二种方式: 初始化数组，初始化值
	var nums2 = [5]int{1, 2, 3, 4, 5}
	// 第三种方式【简化】: 初始化数组，初始化值
	nums3 := [5]int{1, 2, 3, 4, 5}
	// 第四种方式【数组长度不确定】: ...
	nums4 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	// 第五种方式【指定索引下标初始化】:  将索引为 1 和 3 的元素初始化
	nums5 := [5]float32{1: 2.0, 3: 7.0}

	fmt.Println("nums1=", nums1)
	fmt.Println("nums2=", nums2)
	fmt.Println("nums3=", nums3)
	fmt.Println("nums4=", nums4)
	fmt.Println("nums5=", nums5)
}

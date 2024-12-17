package _slice

import "fmt"

// TestCreateSlice var spliceName []dataType
func TestCreateSlice() {
	// 第一种方式: 初始化切片，默认值全为0
	var nums1 []int
	printSlice("nums1", nums1)

	// 第二种方式: 初始化切片，初始化值
	var nums2 = []int{1, 2, 3, 4, 5}
	printSlice("nums2", nums2)

	// 第三种方式【简化】: 初始化数组，初始化值
	nums3 := []int{1, 2, 3, 4, 5}
	printSlice("nums3", nums3)

	// 第四种方式【使用make】: 指定长度
	var nums4 []int = make([]int, 5)
	printSlice("nums4", nums4)

	// 第五种方式【使用make】: 指定长度 和 初始容量大小
	nums5 := make([]int, 3, 6)
	printSlice("nums5", nums5)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 第六种方式【使用make】: 指定长度 和 初始容量大小
	// data: 3 ~ 5 len = [endIndex - startIndex] cap = [len(arr) - startIndex]
	nums6 := arr[3:6]
	printSlice("nums6", nums6)

	nums7 := arr[2:]
	printSlice("nums7", nums7)

	nums8 := arr[:8]
	printSlice("nums8", nums8)
}

func TestDoubleSplice() {
	var nums1 [][]int
	printDoubleSlice("nums1", nums1)

	nums2 := [][]int{{3, 4, 5, 6}, {7, 8, 9, 1}, {2, 6, 7, 3}, {35, 67, 34, 87}}
	printDoubleSlice("nums2", nums2)

	var nums3 [][]int = make([][]int, 5, 10)
	printDoubleSlice("nums3", nums3)

	nums4 := make([][]int, 10)
	printDoubleSlice("nums4", nums4)

}

func TestAppendSlice() {

	var numbers []int
	printSlice("numbers", numbers)

	// 添加一个元素
	numbers = append(numbers, 1)
	printSlice("numbers", numbers)

	// 添加多个元素
	numbers = append(numbers, 5, 6, 7)
	printSlice("numbers", numbers)
}

func TestCopySlice() {

	num1 := []int{23, 34, 36}
	printSlice("num1", num1)

	// 创建一个新的切片，容量为原来的2倍
	num2 := make([]int, len(num1), cap(num1)*2)

	// copy num1数组的数据到num2
	copy(num2, num1)
	printSlice("num2", num2)
}

func printSlice(s string, nums []int) {
	fmt.Printf("%s%v, len=%d, cap=%d\n", s, nums, len(nums), cap(nums))
}

func printDoubleSlice(s string, nums [][]int) {
	fmt.Printf("%s%v, len=%d, cap=%d\n", s, nums, len(nums), cap(nums))
}

func PrintSlice() {
	fmt.Println("****一维数组遍历****")
	nums1 := []int{5, 6, 9, 10, 12}
	for index, value := range nums1 {
		fmt.Printf("nums[%d] = %d\n", index, value)
	}
	fmt.Println("****二维数组遍历****")
	nums2 := [][]int{{3, 4, 5, 6}, {7, 8, 9, 1}, {2, 6, 7, 3}, {35, 67, 34, 87}}
	for row, cols := range nums2 {
		for col, value := range cols {
			fmt.Printf("nums[%d][%d]=%d  ", row, col, value)
		}
		fmt.Println()
	}
}

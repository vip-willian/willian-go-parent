package _ctrol

import (
	"fmt"
	"time"
)

// for init; condition; post {}
// init： 一般为赋值表达式，给控制变量赋初值 [可选]
// condition： 关系表达式或逻辑表达式，循环控制条件 [可选]
// post： 一般为赋值表达式，给控制变量增量或减量 [可选]
func TestFor() {
	for i := 0; i < 10; i++ {
		fmt.Printf("这是for循环的第%d次打印", i)
	}
}

func TestForWhile() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func TestForRange() {
	strings := []string{"google", "willian"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	numbers := [6]int{1, 2, 3, 5}
	for i, v := range numbers {
		fmt.Printf("第 %d 位的值 = %d\n", i, v)
	}
}

func TestForRangeMap() {

	// 初始化一个Map
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
}

func TestDoubleFor() {
	// 定义局部变量
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			// 如果发现因子，则不是素数
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	fmt.Println("******九九乘法表*******")
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
		}
		fmt.Println("")
	}
}

func TestForBreak() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			/* a 大于 15 时使用 break 语句跳出循环 */
			break
		}
	}
}

func TestForBreakLabel() {
	// 不使用标记
	fmt.Println("---- break not label ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}

func TestSwitchBreak() {
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
		// 默认会有break
		// break
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	}
}

func TestSelectBreak() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("Received from ch1.")
	case <-ch2:
		fmt.Println("Received from ch2.")
		break // 跳出 select 语句
	}
}

func TestContinue() {
	var a int = 10
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}

// goto : 将控制转移到被标记的语句。
func TestGoto() {
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

func TestForRangeStr() {
	s := "willian"
	for i, v := range s {
		fmt.Printf("key : %v, value : %c\n", i, v)
	}
}

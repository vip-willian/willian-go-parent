package _ctrol

import (
	"fmt"
	"time"
)

// if语句
func TestIf() {

	var a int = 10
	if a < 20 {
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

// if-else语句
func TestIfElse() {
	var a int = 100
	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

// if初始化
func TestIfInit() {
	var a int
	if a = 300; a < 20 {
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
	if b := 10; a < 20 {
		fmt.Printf("b 小于 20, b = %d\n", b)
	}
}

// switch
func TestSwitch() {
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

func TestTypeSwitch() {
	// 空接口可以支持任意类型
	var x interface{}
	switch i := x.(type) {
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("x 的类型 :%T", i)
	}
}

// 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为true。
func TestSwitchFallthrough() {
	a := true
	b := false
	switch {
	case b:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case a: // 打印
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case b: // 打印
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case a: // 打印
		fmt.Println("4、case 条件语句为 true")
	case b:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

// TestSelect
// select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
// select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
// 如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。
/*
1、每个 case 都必须是一个通道
2、所有 channel 表达式都会被求值
3、所有被发送的表达式都会被求值
4、如果任意某个通道可以进行，它就执行，其他被忽略。
5、如果有多个 case 都可以运行，select 会随机公平地选出一个执行，其他不会执行。

	  否则：
		1）如果有 default 子句，则执行该语句。
		2）如果没有 default 子句，select 将阻塞，直到某个通道可以运行；Go 不会重新对 channel 或值进行求值。
*/
func TestSelect() {

	// 建立2个通道c1,c2
	c1 := make(chan string)
	c2 := make(chan string)

	// 开启一个协程往c1通道写入数据one
	go func() {
		for {
			// time.Sleep(1 * time.Second)
			// 写入数据到通道中
			c1 <- "one"
		}
	}()

	// 开启一个协程往c2通道写入数据two
	go func() {
		for {
			// time.Sleep(2 * time.Second)
			// 写入数据到通道中
			c2 <- "two"
		}
	}()

	for {
		// 随机从通道中获取数据
		select {
		// 从c1通道读取数据
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		// 从c2通道读取数据
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		// 如果2个通道都没获取到数据,走默认的 ; 没默认语句,通道阻塞
		default:
			fmt.Println("no message received")
		}
	}
}

func TestSelectChan() {

	c := 0
	dataChan := make(chan int)
	stopChan := make(chan bool)

	go Channel(dataChan, stopChan)

	for {
		select {
		// C S 二选一输出打印
		case c = <-dataChan:
			fmt.Println("Receive C ", c)
		case s := <-dataChan:
			fmt.Println("Receive S ", s)
		// 输出完了，往stop通道里输入true
		case bool := <-stopChan:
			if bool {
				return
			}
		}
	}
}

func Channel(ch chan int, stopCh chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)
	}
	stopCh <- true
}

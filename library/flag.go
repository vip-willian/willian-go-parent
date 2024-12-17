package library

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// 简单获取命令行参数
func TestOsArgs() {
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

/*
支持的命令行参数格式有以下几种：

	-flag xxx （使用空格，一个-符号）
	--flag xxx （使用空格，两个-符号）
	-flag=xxx （使用等号，一个-符号）
	--flag=xxx （使用等号，两个-符号）

flag.Args() // 返回命令行参数后的其他参数，以[]string类型
flag.NArg() // 返回命令行参数后的其他参数个数
flag.NFlag() // 返回使用的命令行参数个数
*/
func TestFlagParser() {

	// 定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	// flag名 默认值 帮助信息
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	// 定义命令行参数方式2
	// name := flag.String("name", "张三", "姓名")
	// age := flag.Int("age", 18, "年龄")
	// married := flag.Bool("married", false, "婚否")
	// delay := flag.Duration("d", 0, "延迟的时间间隔")

	// 解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

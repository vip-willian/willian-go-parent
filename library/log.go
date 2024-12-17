package library

import "log"

// Print系列(Print|Printf|Println）

// Fatal系列（Fatal|Fatalf|Fatalln）: 写入日志信息后调用os.Exit(1)

// Panic系列（Panic|Panicf|Panicln）: 写入日志信息后panic

func TestLog() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}

// 可选项
//
//	Ldate         = 1 << iota     // 日期：2009/01/23
//	Ltime                         // 时间：01:23:23
//	Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
//	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
//	Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
//	LUTC                          // 使用UTC时间
//	LstdFlags     = Ldate | Ltime // 标准logger的初始值
func TestSetLogFormatter() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}

// 一般使用第三方日志库: logrus、zap
func TestSetLogPrefix() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[willian]")
	log.Println("这是一条很普通的日志。")
}

package library

import (
	"fmt"
	"time"
)

func TestTime() {

	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestTimestamp() {

	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func TestTimestampConvert() {

	timeObj := time.Unix(1734340328, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)

	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestDuration() {
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)
}

func TestTimeAdd() {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
}

// go诞生的时间
const time_template = "2006-01-02 15:04:05"

// 求两个时间之间的差值：
func TestTimeSub() {

	time1 := "2024-12-16 12:23:29"
	time2 := "2024-12-16 15:50:29"

	t1, err1 := time.Parse(time_template, time1)
	t2, err2 := time.Parse(time_template, time2)

	if err1 == nil && err2 == nil {
		d := t2.Sub(t1)
		fmt.Println("【2024-12-16 15:50:29】与【2024-12-16 12:23:29】相差: ", d)
	}
}

// 判断时间是否相等
func TestTimeEqual() {

	time1 := "2024-12-16 12:23:29"
	time2 := "2024-12-16 15:50:29"

	t1, err1 := time.Parse(time_template, time1)
	t2, err2 := time.Parse(time_template, time2)

	if err1 == nil && err2 == nil {
		fmt.Println("【2024-12-16 15:50:29】与【2024-12-16 12:23:29】是否相等: ", t1.Equal(t2))
	}
}

// 如果t代表的时间点在u之前，返回真；否则返回假。
func TestTimeBefore() {

	time1 := "2024-12-16 12:23:29"
	time2 := "2024-12-16 15:50:29"

	t1, err1 := time.Parse(time_template, time1)
	t2, err2 := time.Parse(time_template, time2)

	if err1 == nil && err2 == nil {
		fmt.Println("t2=【2024-12-16 15:50:29】是否在 t1=【2024-12-16 12:23:29】之前: ", t2.Before(t1))
	}
}

// 如果t代表的时间点在u之后，返回真；否则返回假。
func TestTimeAfter() {
	time1 := "2024-12-16 12:23:29"
	time2 := "2024-12-16 15:50:29"

	t1, err1 := time.Parse(time_template, time1)
	t2, err2 := time.Parse(time_template, time2)

	if err1 == nil && err2 == nil {
		fmt.Println("t2=【2024-12-16 15:50:29】是否在 t1=【2024-12-16 12:23:29】之后: ", t2.After(t1))
	}
}

// 时间格式化
func TestTimeFormatter() {
	now := time.Now()
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

// 定时器
func TestTick() {
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

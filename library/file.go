package library

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// 标准输入输出文件
/*
	os.Stdin：标准输入的文件实例，类型为*File
	os.Stdout：标准输出的文件实例，类型为*File
	os.Stderr：标准错误输出的文件实例，类型为*File
*/

// 创建文件
// 根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666
func TestCreateFile() {

}

/*
func Create(name string) (file *File, err Error)
	根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666
func NewFile(fd uintptr, name string) *File
	根据文件描述符创建相应的文件，返回一个文件对象
func Open(name string) (file *File, err Error)
	只读方式打开一个名称为name的文件
func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
	打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
func (file *File) Write(b []byte) (n int, err Error)
	写入byte类型的信息到文件
func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
	在指定位置开始写入byte类型的信息
func (file *File) WriteString(s string) (ret int, err Error)
	写入string信息到文件
func (file *File) Read(b []byte) (n int, err Error)
	读取数据到b中
func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
	从off开始读取数据到b中
func Remove(name string) Error
	删除文件名为name的文件
*/

func TestOpen() {

	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./time.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 读取文件内容
	readFile(file)
	// 关闭文件
	defer file.Close()
}

func TestWriteFileData() {

	// 创建一个文件
	file, err := os.Create("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 写入数据
	for i := 0; i < 5; i++ {
		// 文件中写入字符串
		file.WriteString("ab\n")
		// 以byte方式写入
		file.Write([]byte("cd\n"))
	}
}

// 文件拷贝
func TestFileCopy() {
	// 源文件
	srcFile, err := os.Open("./source.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 目标文件
	dstFile, err2 := os.Create("./target.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	// 缓冲读取
	buf := make([]byte, 1024)
	for {
		// 从源文件读数据
		len, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		// 写入到目标文件
		dstFile.Write(buf[:len])
	}

	srcFile.Close()
	dstFile.Close()
}

func readFile(file *os.File) {

	// 定义接收文件内容的字节数组
	var buf [128]byte

	var content []byte
	// 不断循环读取，直到结束
	for {
		len, err := file.Read(buf[:])
		// 说明读取结束
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file error", err)
			return
		}
		// 将读取到的内容全部写入content
		content = append(content, buf[:len]...)
	}
	// 打印读取到的内容
	fmt.Println(string(content))
}

func TestWriteFile() {
	err := os.WriteFile("./yyy.txt", []byte(""), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestReadFile() {
	content, err := os.ReadFile("./yyy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func TestCat() {
	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') //注意是字符
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

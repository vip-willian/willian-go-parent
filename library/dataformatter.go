package library

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"math/rand"
	"os"
)

type Person struct {
	Name  string
	Hobby string
}

type Student struct {
	Id    int     `json:"-"`
	Name  string  `json:"name"`
	Class string  `json:"class"`
	Score float64 `json:"score"`
}

// TestMarshal 编码json
func TestMarshal() {

	p := Person{"willian", "男"}
	content, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json error", err)
	}
	fmt.Println("编码后的json内容: ", string(content))

	// 格式化输出
	s := Student{1001, "willian", "go学习班", 89.98}
	content, err = json.MarshalIndent(s, "", "     ")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println("格式化编码后的json内容: ", string(content))
}

func TestMapMarshal() {
	student := make(map[string]interface{})
	student["name"] = "willian"
	student["age"] = 18
	student["sex"] = "man"
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}

func TestUnmarshal() {

	jsonStr := `{"Name":"willian","Hobby":"男"}`
	var p Person
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", p)
}

// Servers 读取xml
type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version int      `xml:"version"`
	Servers []Server `xml:"server"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func TestReadXml() {
	// 从文件读取内容
	data, err := os.ReadFile("./doc/server.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 解析xml内容成servers对象
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xml: %#v\n", servers)
}

// User msgpack包读取写入json
type User struct {
	Name string
	Age  int
	Sex  string
}

func TestJsonFileWriteRead() {
	err := writeJson("./doc/person.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = readJson("./doc/person.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeJson(fileName string) (err error) {
	var users []*User
	// 假数据
	for i := 0; i < 10; i++ {
		u := &User{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  "male",
		}
		users = append(users, u)
	}
	data, err := msgpack.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile(fileName, data, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// 二进制读取
func readJson(fileName string) (err error) {
	var users []*User
	// 读文件
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 反序列化
	err = msgpack.Unmarshal(data, &users)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range users {
		fmt.Printf("%#v\n", v)
	}
	return
}

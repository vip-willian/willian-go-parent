package library

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func TestHttpGetBaidu() {

	executeGetRequest("https://www.baidu.com/")
}

// TestHttpGetWithParameter GET 携带请求参数
func TestHttpGetWithParameter() {

	executeGetRequest(getUrlAddress())
}

func executeGetRequest(url string) {
	response, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("出现错误")
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}
	fmt.Println(string(body))
}

func executePostRequest(url string) {

	contentType := "application/json"
	data := `{name:"willian","age":18}`
	response, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func getUrlAddress() string {

	requestUrl := "http://127.0.0.1:9090/get"
	data := url.Values{}
	data.Set("name", "willian")
	data.Set("age", "18")

	url, err := url.ParseRequestURI(requestUrl)

	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	url.RawQuery = data.Encode()
	fmt.Println(url.String())

	return url.String()
}

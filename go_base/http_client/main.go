package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http Client

//func get_baidu() {
//	ret, err := http.Get("https://www.runoob.com/js/js-if-else.html")
//	if err != nil {
//		fmt.Println("get url failed,err:", err)
//	}
//	defer ret.Body.Close()
//	info, _ := ioutil.ReadAll(ret.Body)
//	fmt.Println(string(info))
//}

func getServerDemo() {
	ret, err := http.Get("http://127.0.0.1:9090/get/?name=test&age=18")
	if err != nil {
		fmt.Println("get url failed,err:", err)
		return
	}
	defer ret.Body.Close()
	info, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		fmt.Println("read body failed,err:", err)
		return
	}
	fmt.Println(string(info))
}

func getBlogInfo() {
	data := url.Values{} // url values
	urlObj, _ := url.Parse("http://47.94.142.215:32222/api/v1.0/blog/backend/article?page=1&page_size=10")
	data.Set("name", "one")
	data.Set("age", "20")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	ret, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("read body failed,err:", err)
		return
	}
	// 禁用KeepAlive 到close 的代码段，用于请求不频繁的连接
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	client.Do(req)
	defer ret.Body.Close()

	info, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		fmt.Println("read body failed,err:", err)
		return
	}
	fmt.Println(string(info))
}

func main() {
	//get_baidu()
	getServerDemo()
	getBlogInfo()

}

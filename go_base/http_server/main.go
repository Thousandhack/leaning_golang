package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func stringHandlerFunc(w http.ResponseWriter, r *http.Request) {
	str := `<h1 style="color:red">hello 南山!</h1>`
	w.Write([]byte(str))
}

func readTxtHandlerFunc(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		w.Write([]byte("暂无内容~"))
	}
	w.Write(b)
}

func readHtmlHandlerFunc(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./demo.html")
	if err != nil {
		w.Write([]byte("暂无内容~"))
	}
	w.Write(b)
}

func testHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	// 对于get请求，参数都放在url
	query := r.URL.Query() // 自动识别URL中的参数 变成map的数据
	name := query.Get("name")
	age := query.Get("age")
	fmt.Println(name, age) // 自动识别URL中的参数 变成map的数据
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/posts/Go/", stringHandlerFunc)
	http.HandleFunc("/posts/txt/", readTxtHandlerFunc)
	http.HandleFunc("/posts/html/", readHtmlHandlerFunc)
	http.HandleFunc("/get/", testHandlerFunc)
	http.ListenAndServe("0.0.0.0:9090", nil)
}

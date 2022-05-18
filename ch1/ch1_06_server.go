// 一些简单的server

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func server1() {
	// server1是一个迷你回声服务器
	http.HandleFunc("/", handler1) // 回声请求调用处理程序
	log.Fatal()
}

// 处理程序回显请求 URL 的路径部分
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

var mu sync.Mutex
var count int

func server2() {
	// server2是一个迷你的回声和计数器服务器
	http.HandleFunc("/handler2", handler2) // 回声请求调用处理程序
	http.HandleFunc("/handler3", handler3) // 回声请求调用处理程序
	http.HandleFunc("/count", counter2)    // 回声请求调用处理程序
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}) // 利萨茹图形
}

// 处理程序回显请求 URL 的路径部分
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter 回显目前为止调用的次数
func counter2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func main() {
	//server1()
	server2()
}

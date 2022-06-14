package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	/*fmt.Fprintf(w, "<h1>Hello,这里是goblog</h1>")
	fmt.Fprintf(w, "请求路径为："+r.URL.Path)*/
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>这里是GoBlog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprintf(w, "此博客是用于记录编程笔记，如您有反馈和建议，请联系 "+"<a href=\"zhongtao1024@gmail.com\">zhongtao1024@gmail.com</a>")
	} else {
		fmt.Fprintf(w, "<h1>页面未找到 :( </h1>"+"<p>如有疑惑，请联系我们。</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}

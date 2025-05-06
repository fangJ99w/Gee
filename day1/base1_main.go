package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

//置了2个路由，/和/hello，分别绑定 indexHandler 和 helloHandler ， 根据不同的HTTP请求会调用不同的处理函数。访问/，响应是URL.Path = /，而/hello的响应则是请求头(header)中的键值对信息。
//
//用 curl 这个工具测试一下，将会得到如下的结果。
//$ curl http://localhost:9999/
//URL.Path = "/"
//$ curl http://localhost:9999/hello
//Header["Accept"] = ["*/*"]
//Header["User-Agent"] = ["curl/7.54.0"]

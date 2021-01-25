package main

import (
	"fmt"
	"log"
	"net/http"

	apm "github.com/Richeir/rycked"
)

//最终的页面
func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

//首页的触发链接
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href='/query'> Click here to start a request </a>`))
}

func main() {
	//http服务器实例
	mux := http.NewServeMux()
	//端口号
	port := 5000

	//注册路由
	mux.HandleFunc("/", indexHandler)
	finalHandler := http.HandlerFunc(final)
	mux.Handle("/query", apm.AddAPMMiddleware((finalHandler)))

	//启动本地服务器
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	fmt.Printf("Web server start at: http://localhost:%d/  \n", port)
	log.Print(http.ListenAndServe(addr, mux))
}

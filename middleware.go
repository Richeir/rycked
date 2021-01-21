package rycked

import (
	"log"
	"net/http"
)

// ApmMiddleware 中间件
func ApmMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//HTTP处理之前发生的事情
		log.Println("Executing middleware before")

		//判断 tacer 是否存在，存在就读取，不存在就创建一个新的

		//创建Span

		//调用正常的HTTP处理过程
		next.ServeHTTP(w, r)

		//完成 span 各种信息

		//写入ES

		//HTTP处理之后发生的事情
		log.Println("Executing middleware after")
	})
}

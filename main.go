package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}
func main() {

	//1.注册一个给定模式的处理器函数到DefaultServeMux
	http.HandleFunc("/", sayHello)

	//2.设置监听的TCP地址并启动服务
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	log.Fatal(err)
}

package main

import (
	"io"
	"net/http"
)

func Listen() {
	http.HandleFunc("/", heartbeat)      //设定访问的路径
	go http.ListenAndServe(":8899", nil) //设定端口和handler
}

func heartbeat(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "ok")
}

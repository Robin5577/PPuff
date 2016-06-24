package main

import (
	//"fmt"
	"time"
)

var ZkSevs = []string{""}
var ZkTimeout = time.Second * 10
var ZkNodes = []string{"/mirror", "/test01"}
var ZkDebug = false
var ZkLogPath = ".\\b\\a\\log"
var ZKRootPath = ""

var ZkLoger *ZkLog

func main() {

	/*defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("%s\n", err.(error).Error())
		}
	}()*/
	ZkLoger = logSe()
	ZkInit()
}


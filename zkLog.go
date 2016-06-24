package main

import (
	"fmt"
	"log"
	"sync"
)

const (
	bufLen = 1024
	bufCap = 100
)

type logBuffer byte
type ZkLog struct {
	loger *log.Logger
}

var logB logBuffer
var logHandler *log.Logger
var buffers [][]byte
var lock *sync.Mutex = &sync.Mutex{}
var zkf *ZkF = &ZkF{}

func (b *logBuffer) Write(p []byte) (n int, err error) {
	buf := logPop()
	l := len(p)
	if l > bufLen {
		l = bufLen
	}
	buf = buf[0:l]
	n = copy(buf, p[0:l])
	err = nil
	if ZkDebug {
		fmt.Println("\033[40;32m" + string(buf) + "\033[0m")
	} else {
		zkf.PutAppend(ZkLogPath, buf)
	}
	logPush(buf)
	return
}

//pop action of Buffer Cache
func logPop() []byte {

	lock.Lock()

	if len(buffers) < 1 {
		_buf := make([]byte, 0, bufLen)
		buffers = append(buffers, _buf)
	}
	_b := buffers[0:1]
	buf := _b[0][:]
	buffers = buffers[1:]
	lock.Unlock()

	return buf
}

//push action of Buffer Cache
func logPush(b []byte) {

	if len(buffers) < bufCap {
		buffers = append(buffers, b[0:0])
	}
}

//Logserver
func logSe() *ZkLog {

	logHandler = &log.Logger{}
	logHandler.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	buffers = make([][]byte, 0, 100)
	logB = logBuffer(1)
	logHandler.SetOutput(&logB)

	return &ZkLog{logHandler}
}

func (log *ZkLog) Printf(f string, s ...interface{}) {
	//if ZkDebug {
	log.loger.SetPrefix("[DEBUG]")
	log.loger.Output(2, fmt.Sprintf(f, s))
	//}
}

func (log *ZkLog) PrintLog(f string, s ...interface{}) {
	log.loger.SetPrefix("[TRACE]")
	log.loger.Output(2, fmt.Sprintf(f, s))
}


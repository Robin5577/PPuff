//Logger Lib
package core

import (
   "log"
   "sync"
   "fmt"
)

type logBuffer byte 

var logB logBuffer
var logHandler *log.Logger
var buffers [][]byte
var lock *sync.Mutex = &sync.Mutex{}

func (b *logBuffer) Write(p []byte)(n int,err error){
     buf := logPop()
     n = copy(buf,p)
     err = nil
     if Debug  {
        fmt.Println("\033[40;32m"+string(buf)+"\033[0m")
     }
     logPush(buf)
     return
}

//pop action of Buffer Cache 
func logPop() []byte {

    lock.Lock()
    
    if len(buffers) < 1 {
       _buf := make([]byte,256)
       buffers = append(buffers,_buf)
    }
    _b := buffers[0:1]
    buf := _b[0][:]
    buffers = buffers[1:]
    lock.Unlock()

    return buf
}
//push action of Buffer Cache
func logPush(b []byte){

    if len(buffers) < 100 {
       buffers = append(buffers,b)
    }
}

//Logserver 
func logSe(){

   logHandler = &log.Logger{}
   logHandler.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)  
   logB = logBuffer(1)
   logHandler.SetOutput(&logB)
}

func logDebug(s string){
   if Debug {
      logHandler.SetPrefix("[DEBUG]")
      logHandler.Output(2,s)
   }
}

func logTrace(s string){
   logHandler.SetPrefix("[TRACE]")
   logHandler.Output(2,s)
}

func logError(s string){
   logHandler.SetPrefix("[ERROR]")
   logHandler.Output(2,s)
}

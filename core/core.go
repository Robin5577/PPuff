package core

import (
   "flag"
   "os"
   "strings"
)

var Dae bool = false
var Debug bool = false

func Start(){

    defer func(){
	if err := recover(); err != nil {
           logError(err.(error).Error())
        }
    }()
    
    flagSe()
}

func flagSe(){

   daemon := flag.Bool("daemon", false, "deamon mode ?") 
   debug  := flag.Bool("debug", false, "debug mode ?")
   conf   := flag.String("conf", "", "the config file is ...")

   flag.Parse()

   Dae = *daemon
   Debug = *debug

   if Dae {
      Daemon()
      os.Exit(0)
   }

   logSe()

   if strings.TrimSpace(*conf) == "" {
      logError("the conf is unknow")   
   }
   
   logTrace("Starting ...")
   SetSignal()
}
 

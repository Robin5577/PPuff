//Deamon mode
package core

import (
   "os"
   "os/signal"
   "strings"
   "os/exec"
)

func Daemon() {

   var se string = os.Args[0]
   var args []string = os.Args[1:]
   filter(args)
   cmd := exec.Command(se, args...)
   err := cmd.Start()

   if err != nil {
      panic("Start unsuccessfully")
   }
}

//Filter the flag of daemon
func filter(args []string) {

   for arg := range args {
       if strings.EqualFold(args[arg], "--daemon") || strings.EqualFold(args[arg], "-daemon") {
          args[arg] = ""
	  return
       }
   }

}

//initialize the signal controal
func SetSignal(){

   var c chan os.Signal
   c = make(chan os.Signal) 

   go func(){
      signal.Notify(c, os.Interrupt,os.Kill)
   }()

   for {
      switch <-c {
          case os.Interrupt:
               logTrace("The PPuff was interrupted")
               return 
          case os.Kill:
               logTrace("The PPuff was killed")
               return
      }
   }
}

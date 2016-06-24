package main

import (
	"os"
	"os/exec"
	"os/signal"
	"strings"
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
		if strings.EqualFold(args[arg], "--debug") || strings.EqualFold(args[arg], "-debug") {
			args[arg] = ""
			return
		}
	}

}

//initialize the signal controal
func SetSignal() {

	defer ShutDown()

	var c chan os.Signal
	c = make(chan os.Signal)

	go func() {
		signal.Notify(c, os.Interrupt, os.Kill)
	}()

	for {
		switch <-c {
		case os.Interrupt:
			ZkLoger.PrintLog("%s", "The PPuff was interrupted")
			return
		case os.Kill:
			ZkLoger.PrintLog("%s", "The PPuff was killed")
			return
		}
	}
}

func ShutDown() {
	if err := recover(); err != nil {
		ZkLoger.PrintLog("%+v", err)
	}
	ZKDes()
}


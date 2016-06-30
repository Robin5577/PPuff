package main

import (
	"flag"
	"os"
)

var ZkConfPath string
var ZkDebug bool
var ZkLoger *ZkLog

func main() {

	defer ShutDown()

	initialize()
}

func initialize() {

	flag.BoolVar(&ZkDebug, "debug", false, "debug mode ?")
	flag.StringVar(&ZkConfPath, "conf", "./ppuff.yaml", "the config file is ...")

	flag.Parse()

	ParseConf()

	if os.Getppid() != 1 && !ZkDebug {
		Daemon()
		os.Exit(0)
	}

	ZkLoger = logSe()
	ZkInit()
	ZkLoger.PrintLog("%s", "The Ppuff starting ...")
	SetSignal()
}

package main

import (
	"gopkg.in/yaml.v2"
	"os"
)

type ZkC struct {
	ZkSevs     []string `yaml:"zkServers,flow"`
	ZkTimeout  int      `yaml:"zkTimeout"`
	ZkNodes    []string `yaml:"zkNodes,flow"`
	ZkLogPath  string   `yaml:"zkLogPath,omitempty"`
	ZkRootPath string   `yaml:"zkRootPath"`
}

var ZKConf *ZkC = &ZkC{[]string{}, 10, []string{}, "./ppuff.log", "./"}

/*var ZkSevs = []string{"10.30.6.180:2181"}
var ZkTimeout = time.Second * 10
var ZkNodes = []string{"/test"}
var ZkLogPath = "./ppuff.log"
var ZKRootPath = "./"*/

func ParseConf() {
	f, err := os.Open(ZkConfPath)

	if err != nil {
		panic(err)
	}

	stat, _ := f.Stat()

	buf := make([]byte, stat.Size())
	if _, err := f.Read(buf); err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(buf, ZKConf); err != nil {
		panic(err)
	}
}


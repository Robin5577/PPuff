package main

import (
	"gopkg.in/yaml.v2"
	"os"
)

type ZkC struct {
	ListenPort string   `yaml:"listenPort"`
	ZkSevs     []string `yaml:"zkServers,flow"`
	ZkTimeout  int      `yaml:"zkTimeout"`
	ZkNodes    []string `yaml:"zkNodes,flow"`
	ZkLogPath  string   `yaml:"zkLogPath,omitempty"`
	ZkRootPath string   `yaml:"zkRootPath"`
}

var ZKConf *ZkC = &ZkC{"127.0.0.1:8899", []string{}, 10, []string{}, "./ppuff.log", "./"}

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

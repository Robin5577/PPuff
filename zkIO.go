package main

import (
	"os"
	"strings"
)

type ZkF struct {
	fd *os.File
}

func (zkf *ZkF) PutAppend(name string, b []byte) {

	fd, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_SYNC, 0664)

	if err != nil {
		if os.IsPermission(err) {
			panic(err)
		}
		if os.IsNotExist(err) {
			zkf.mkdir(name)
		}
		fd, err = os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_SYNC, 0664)
		if err != nil {
			panic(err)
		}
	}
	zkf.fd = fd
	zkf.write(b)
}

func (zkf *ZkF) Put(name string, b []byte) {
	fd, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_SYNC, 0664)

	if err != nil {
		if os.IsPermission(err) {
			panic(err)
		}
		if os.IsNotExist(err) {
			zkf.mkdir(name)
		}
		fd, err = os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_SYNC, 0664)
		if err != nil {
			panic(err)
		}
	}
	zkf.fd = fd
	zkf.write(b)
}

func (zkf *ZkF) mkdir(name string) {
	paths := strings.Split(name, string(os.PathSeparator))

	paths = paths[0 : len(paths)-1]

	if err := os.MkdirAll(strings.Join(paths, string(os.PathSeparator)), 0776); err != nil {
		panic(err)
	}
}

func (zkf *ZkF) write(b []byte) {

	defer zkf.fd.Close()

	if _, err := zkf.fd.Write(b); err != nil {
		//panic(err)
	}
}


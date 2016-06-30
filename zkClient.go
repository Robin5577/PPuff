package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

var putData *ZkF = &ZkF{}
var zkConn *zk.Conn
var zkInitState = false

func ZkInit() {
	zkconn, zkevent, zkerror := zk.Connect(ZKConf.ZkSevs, time.Duration(ZKConf.ZkTimeout)*time.Second)

	if zkerror != nil {
		panic(zkerror)
	}

	zkconn.SetLogger(ZkLoger)
	zkConn = zkconn

	for _, node := range ZKConf.ZkNodes {
		go zkChWatch(node)
	}

	go func() {
		for {
			select {
			case ch := <-zkevent:
				if ch.State == zk.StateConnected {
					if zkInitState {
						ZkLoger.Printf("=====%s=======", "Reconnect")
						zkConn.ClearWatches()
					} else {
						zkInitState = true
					}
				} else {
					ZkLoger.Printf("%+v", ch)
				}
			}
		}
	}()
}

func ZKDes() {
	zkConn.Close()
}

func zkChWatch(path string) {

	nodesWatchers := make(map[string]chan string)

	for {
		nodes, _, watch, err := zkConn.ChildrenW(path)

		if err != nil {
			if err == zk.ErrNoServer {
				ZkLoger.PrintLog("=====ErrConnectionClosed========: %s,the error msg: %+v", path, err)
				continue
			}
			ZkLoger.PrintLog("Thr child: %s,the error msg: %+v", path, err)
			return
		}

		ZkLoger.Printf("The childrens info: %+v", nodes)

		for n, nw := range nodesWatchers {
			delWatcher := true
			for _, node := range nodes {
				if strings.EqualFold(n, node) {
					delWatcher = false
				}
			}

			if delWatcher {
				close(nw)
				delete(nodesWatchers, n)
				ZkLoger.Printf("%s is delete 1", n)
			}
		}

		for _, node := range nodes {
			if _, ok := nodesWatchers[node]; !ok {
				nodesWatchers[node] = make(chan string, 1)
				go zkNodeWatch(path+"/"+node, nodesWatchers[node])
			}
		}

		select {
		case <-watch:
		}
	}
}

func zkNodeWatch(path string, nodeWatcher chan string) {

	for {
		select {
		case <-nodeWatcher:
			ZkLoger.Printf("%s is delete 2", path)
			return
		default:
		}

		node, _, watch, err := zkConn.GetW(path)

		if err != nil {
			ZkLoger.PrintLog("The node: %s,the error msg:%+v", path, err)
			select {
			case <-time.After(time.Millisecond * 500):
			}
			continue
		}

		putData.Put(ZKConf.ZkRootPath+path, node)
		ZkLoger.Printf("The node info: %s", node)

		select {
		case <-watch:
		}
	}
}

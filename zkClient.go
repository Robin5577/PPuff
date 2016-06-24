package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

var zkConn *zk.Conn

func ZkInit() {
	zkconn, zkevent, zkerror := zk.Connect(ZkSevs, ZkTimeout)

	if zkerror != nil {
		panic(zkerror)
	}

	zkconn.SetLogger(ZkLoger)
	zkConn = zkconn

	for _, node := range ZkNodes {
		go zkChWatch(node)
	}
	//go func() {
	for {
		select {
		case ch := <-zkevent:
			if ch.State == zk.StateExpired {
				ZkLoger.Printf("%+v", ch)
			} else {
				ZkLoger.Printf("%+v", ch)
			}
		}
	}
	//}()
}

func zkChWatch(path string) {

	nodesWatchers := make(map[string]chan string)

	for {
		nodes, _, watch, err := zkConn.ChildrenW(path)

		if err != nil {
			ZkLoger.PrintLog("The child: %s,the error msg: %+v", path, err)
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

		ZkLoger.Printf("The node info: %s", node)

		select {
		case <-watch:
		}
	}
}

package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
)

func main() {
	var addr []string
	addr = make([]string, 1)
	addr[0] = "127.0.0.1:2181"

	conn, chW, _ := zk.Connect(addr, 5*time.Second)
	_, err := conn.Create("/main", []byte(""), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("0000000")
	_, _, ch, _ := conn.ExistsW("/main")
	log.Println("111111")
	go func(ch <-chan zk.Event) {
		ev := <-ch
		log.Printf("%v", ev)
	}(ch)
	go func(chW <-chan zk.Event) {
		for {
			ev := <-chW
			log.Printf("%v", ev)
		}

	}(chW)
	conn.Set("/main", []byte("haoren"), -1)
	conn.Close()
	time.Sleep(1000 * time.Second)

}

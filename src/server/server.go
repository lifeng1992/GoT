package main

import (
	"fmt"
	"net/rpc"
	"net"
	"log"
	"net/http"
	"time"
	"io/ioutil"
)

type Argv struct  {
	msg string
}

type Rpg struct {}

func (t Rpg)SyaHello(args string, reply *time.Time) error {
	fmt.Println("Remote call...")
	fmt.Println(args)
	*reply = time.Now()

	return nil
}

func (t Rpg)WriteCert(args string, reply *bool) error {
	err := ioutil.WriteFile("/README.md", []byte(args), 0644)
	if err != nil {
		*reply = true
	}

	return err
}

func main() {
	rpc.Register(Rpg{})
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8964")
	if e != nil {
		log.Fatal("Listen error:", e)
	}

	fmt.Println("Rpc Listen At localhost:8964 ......")
	http.Serve(l, nil)
}


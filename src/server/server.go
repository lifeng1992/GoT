// env GOOS=linux GOARCH=amd64 go build -v -o server-linux src/server/server.go
package main

import (
	"fmt"
	"net/rpc"
	"net"
	"log"
	"net/http"
	"io/ioutil"
)

type Rpg struct {}

func (t Rpg)WriteCert(args string, reply *bool) error {
	err := ioutil.WriteFile("/README.md", []byte(args), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	*reply = true
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


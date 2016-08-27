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
	b, err := ioutil.ReadFile("/root/.ssh/authorized_keys")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	
	err = ioutil.WriteFile("/root/.ssh/authorized_keys", b, 0777)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	*reply = true
	return nil
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


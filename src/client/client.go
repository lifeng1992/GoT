package main

import (
	"net/rpc"
	"log"
	"fmt"
	"flag"
)

func main() {
	var ip = flag.String("ip", "127.0.0.1", "Server IP")
	flag.Parse()

	client, err := rpc.DialHTTP("tcp", *ip + ": 8964")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	msg := "Hello World\n"
	reply := false
	err = client.Call("Rpg.WriteCert", msg, &reply)
	if err != nil {
		log.Fatal("Rpg error:", err)
	}
	fmt.Println("Remote Call Return:", reply)
}


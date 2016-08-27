ckage main

import (
	"net/rpc"
	"log"
	"fmt"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1" + ": 8964")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	msg := "Hello World"
	reply := false
	err = client.Call("Rpg.WriteCert", msg, &reply)
	if err != nil {
		log.Fatal("Rpg error:", err)
	}
	fmt.Println("Remote Call Return:", reply)
}


package gysrpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func ServeRpc(){
	log.Println("Starting rpc server")
	r := new(RPCHandler)
	err := rpc.Register(r)
	log.Println(err)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":9000")
	log.Println(err)
	// Close the listener whenever we stop
	defer listener.Close()
	// Wait for incoming connections
	//rpc.Accept(listener)
	if err := http.Serve(listener, nil); err != nil{
		log.Fatal(err)
	}
}
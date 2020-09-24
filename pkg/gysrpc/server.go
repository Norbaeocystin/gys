package gysrpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc() {
	log.Println("Starting rpc server")
	r := new(RPCHandler)
	srv := rpc.NewServer()
	err := srv.Register(r)
	log.Println(err)
	listener, err := net.Listen("tcp", ":9000")
	log.Println(err)
	// Close the listener whenever we stop
	defer listener.Close()
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
//Gob serve RPC
//func ServeRpc(){
//	log.Println("Starting rpc server")
//	r := new(RPCHandler)
//	err := rpc.Register(r)
//	log.Println(err)
//	rpc.HandleHTTP()
//	listener, err := net.Listen("tcp", ":9000")
//	log.Println(err)
//	// Close the listener whenever we stop
//	defer listener.Close()
//	// Wait for incoming connections
//	//rpc.Accept(listener)
//	//if err := http.Serve(listener, nil); err != nil{
//	//	log.Fatal(err)
//	//}
//	rpc.Accept(listener)
//}
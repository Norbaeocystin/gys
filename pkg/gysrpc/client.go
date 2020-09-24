package gysrpc

import (
	"log"
	"net/rpc/jsonrpc"
	"os"
)

// {"Url":"https://www.zoznam.sk/firma/2550207/Dajan-Daniela-Valkova-Sobrance","Selector":"div[class='col-md-8 profile middle-content']","Type":"one","Subselectors":[{"Selector":"div.row","Attribute":"text","Name":"","Split":":","Default":""},{"Selector":"div.row","Attribute":"text","Name":"CompanyName","Split":"","Default":""}]}

func CallClient(){
	log := log.New(os.Stderr, "",log.Lmicroseconds)
	log.Println("Calling rpc client")
	//HTTP client
	//client, _ := rpc.DialHTTP("tcp","127.0.0.1:9000")
	//TCP client using GOB protocol
	//client, _ := rpc.Dial("tcp","127.0.0.1:9000")
	client, _ := jsonrpc.Dial("tcp","127.0.0.1:9000")
	var res Response
	request := GysRpc{
		Url:          "https://www.zoznam.sk/firma/2550207/Dajan-Daniela-Valkova-Sobrance",
		Selector:     "div[class='col-md-8 profile middle-content']",
		Type:         "one",
		Subselectors: []Subselector{{"div.row", "text", "", ":", ""}, {"div.row", "text", "CompanyName", "", ""}},
	}
	_ = client.Call("RPCHandler.Execute", request, &res)
	defer client.Close()
	log.Println(&res)
	log.Println("closing rpc client")
}

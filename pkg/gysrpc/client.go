package gysrpc

import (
	"encoding/json"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

// {"Url":"https://www.zoznam.sk/firma/2550207/Dajan-Daniela-Valkova-Sobrance","Selector":"div[class='col-md-8 profile middle-content']","Type":"one","Subselectors":[{"Selector":"div.row","Attribute":"text","Name":"","Split":":","Default":""},{"Selector":"div.row","Attribute":"text","Name":"CompanyName","Split":"","Default":""}]}

func CallClient() ResultHash{
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
	request2 := GysMain{
		Iterator:      Iterator{"https://www.zoznam.sk/katalog/Spravodajstvo-informacie/Abecedny-zoznam-firiem/A/sekcia.fcgi?sid=1173&so=&page=PAGE", "PAGE", 1,2},
		Identificator: Identificator{ "a.link_title", "href","url","many" ,"","https://www.zoznam.sk"},
		Extractor:     Extractor{},
	}
	var res2 ResultHash
	b, _ := json.Marshal(request2)
	log.Println(string(b))
	_ = client.Call("RPCHandler.Iterate", request2, &res2)
	log.Println(&res2)
	log.Println("closing rpc client")
	return res2
}

func CallClient2(rh ResultHash) []string {
	log := log.New(os.Stderr, "",log.Lmicroseconds)
	log.Println("Calling rpc client")
	client, _ := jsonrpc.Dial("tcp","127.0.0.1:9000")
	var res IteratorResponse
	_ = client.Call("RPCHandler.FindIteration", rh, &res)
	defer client.Close()
	log.Println(&res)
	return res
}

func CallClient3(array []string) ResultHash{
	request2 := GysMain{
		Iterator:      Iterator{},
		Identificator: Identificator{},
		Extractor:     Extractor{array, "div", "one", []Subselector{{"div.row", "text", "CompanyName", "", ""},}},
	}
	client, _ := jsonrpc.Dial("tcp","127.0.0.1:9000")
	var res ResultHash
	_ = client.Call("RPCHandler.ExtractAll", request2, &res)
	defer client.Close()
	log.Println(&res)
	return res
}

func CallClient4(rh ResultHash){
	log := log.New(os.Stderr, "",log.Lmicroseconds)
	log.Println("Calling rpc client")
	client, _ := jsonrpc.Dial("tcp","127.0.0.1:9000")
	var res Response
	_ = client.Call("RPCHandler.FindExtract", rh, &res)
	defer client.Close()
	log.Println(&res)
}
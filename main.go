// ./gys --config iterator.yaml | (read u; ./gys --config extractor.yaml --urls $u)
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"gys/gysyaml"
	gys2 "gys/gysyaml/gys"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	urls := flag.String("urls","","urls delimited by comma")
	configfile := flag.String("config","iterator.yaml","yaml config file to parse")
	flag.Parse()
	b, _ := ioutil.ReadFile(*configfile)
	gys := gysyaml.Gys{}
	_ = yaml.Unmarshal(b, &gys)
	if *urls != ""{
		gys.Extractor.Urls = *urls
	}
	if gys.Extractor.Urls != ""{
		extractor(gys)
	}else {
		iterator(gys)
	}
}

func extractor(gys gysyaml.Gys){
	result := gys2.Extract(gys)
	for _, r := range result{
		clean := make(map[string]string)
		for k,v := range r{
			key := strconv.QuoteToASCII(string(k))
			value := strconv.QuoteToASCII(string(v))
			clean[key] = value
		}
		jsonBytes, _ := json.Marshal(clean)
		fmt.Println(string(jsonBytes))
		}
	}


func iterator(gys gysyaml.Gys) {
	links := gys2.GenerateLinks(gys.Iterator.Url, gys.Iterator.Replace, gys.Iterator.Min, gys.Iterator.Max)
	results := make([]string,0)
	for _, link := range links{
		doc := gys2.GetDoc(link)
		result := gys2.ProcessMessage(doc, gys.Identificator.Selector, gys.Identificator.Attribute,gys.Identificator.Type,gys.Identificator.Default, gys.Identificator.Base)
		results = append(results, result...)
	}
	r := strings.Join(results, gys.Output.Delimiter)
	fmt.Print(r)
}
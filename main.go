// ./gys --config iterator.yaml | (read u; ./gys --config extractor.yaml --urls $u)
package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"gys/gysyaml"
	gys2 "gys/gysyaml/gys"
	"io/ioutil"
	"log"
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
		for k,v := range r{
			log.Println(k,v)
		}
	}
}


func iterator(gys gysyaml.Gys) {
	log.Println("hovno")
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
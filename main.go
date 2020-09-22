package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"gys/gysyaml"
	gys2 "gys/gysyaml/gys"
	"io/ioutil"
)

func main() {
	configfile := flag.String("config","iterator.yaml","yaml config file to parse")
	b, _ := ioutil.ReadFile(*configfile)
	gys := gysyaml.Gys{}
	_ = yaml.Unmarshal(b, &gys)
	links := gys2.GenerateLinks(gys.Iterator.Url, gys.Iterator.Replace, gys.Iterator.Min, gys.Iterator.Max)
	for _, link := range links{
		doc := gys2.GetDoc(link)
		result := gys2.ProcessMessage(doc, gys.Identificator.Selector, gys.Identificator.Attribute,gys.Identificator.Type,gys.Identificator.Default, gys.Identificator.Base)
		for _, r := range result{
			fmt.Println(r)
		}
	}
}

// ./gys --config iterator.yaml | (read u; ./gys --config generator.yaml --urls $u)
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"gys/pkg"
	gys2 "gys/pkg/core"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	urls := flag.String("urls","","urls delimited by comma")
	configfile := flag.String("config","iterator.yaml","yaml config file to parse")
	flag.Parse()
	b, _ := ioutil.ReadFile(*configfile)
	gys := pkg.Gys{}
	_ = yaml.Unmarshal(b, &gys)
	if gys.Extractor.Filewithurls != "" {
		b, _ := ioutil.ReadFile(gys.Extractor.Filewithurls)
		urls := strings.ReplaceAll(string(b), "\n", ",")
		gys.Extractor.Urls = urls
	}
	if *urls != ""{
		gys.Extractor.Urls = *urls
	}
	if gys.Extractor.Urls != ""{
		extractor(gys)
	}else {
		iterator(gys)
	}
}

func extractor(gys pkg.Gys){
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


func iterator(gys pkg.Gys) {
	links := gys2.GenerateLinks(gys.Iterator.Url, gys.Iterator.Replace, gys.Iterator.Min, gys.Iterator.Max)
	filename := gys.Output.Filename
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	for _, link := range links{
		doc := gys2.GetDoc(link)
		result := gys2.ProcessMessage(doc, gys.Identificator.Selector, gys.Identificator.Attribute,gys.Identificator.Type,gys.Identificator.Default, gys.Identificator.Base)
		r := strings.Join(result, "\n")
		fmt.Println(r)
		f.WriteString(r)
	}
	f.Sync()
}
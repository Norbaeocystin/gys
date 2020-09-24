package gysrpc

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/PuerkitoBio/goquery"
	gys2 "gys/pkg/core"
	"log"
	"strings"
	"time"
)

type GysRpc struct {
	Url string
	Selector string
	Type string
	Subselectors []Subselector
}

type Subselector struct{
	Selector string
	Attribute string
	Name string
	Split string
	Default string
}

type 	Iterator   struct {
	Url     string
	Replace string
	Min     int
	Max     int
}

type 	Identificator struct {
	Attribute string
	Selector string
	Name string
	Type string
	Default string
	Base string
}

type Extractor struct {
	Urls string
	Selector string
	Type string
	Subselectors []Subselector
}

type GysMain struct {
	Iterator   Iterator
	Identificator Identificator
	Extractor Extractor
}

type Response []map[string]string

type IteratorResponse []string

type Storage map[string][]map[string]string

type IteratorStorage map[string][]string

type ResultHash struct {
	Hash string
}


type RPCHandler struct {
	Storage Storage
	IteratorStorage IteratorStorage
}

func ( h *RPCHandler) Execute(request *GysRpc, response *Response) error {
	log.Println("Extracting info")
	log.Println(request)
	r := Extract(*request)
	*response = r
	return nil
}

func  (h *RPCHandler)ExtractAll(request *GysMain, response *ResultHash) error {
	ex := request.Extractor
	stringbeforehash := time.Now().String() + ex.Urls + ex.Type + ex.Selector
	hash := SHA1(stringbeforehash)
	response.Hash = hash
	go func(){
		r := ExtractGysmain(*request)
		h.Storage[response.Hash] = r
	}()
	return nil
}

func  (h *RPCHandler)Iterate(request *GysMain, response *ResultHash) error {
	it := request.Iterator
	id := request.Identificator
	stringbeforehash := time.Now().String() + it.Replace + it.Url + id.Name + id.Base
	hash := SHA1(stringbeforehash)
	response.Hash = hash
	go func(){
		r := Iterate(*request)
		h.IteratorStorage[response.Hash] = r
	}()
	return nil
}

func ( h *RPCHandler) FindExtract(hash *ResultHash, response *Response) error {
	log.Println("Searching info")
	r, ok := h.Storage[hash.Hash]
	if ok {
		delete (h.Storage, hash.Hash)
		*response = r
	}else{
		res := make(map[string]string)
		res["success"] = "false"
		resp := []map[string]string{res}
		*response = resp
	}
	return nil
}
func ( h *RPCHandler) FindIteration(hash *ResultHash, response *IteratorResponse) error {
	log.Println("Searching info")
	r, ok := h.IteratorStorage[hash.Hash]
	if ok {
		delete (h.IteratorStorage, hash.Hash)
		*response = r
	}else{
		resp := []string{"false"}
		*response = resp
	}
	return nil
}



func Extract( ext GysRpc) []map[string]string {
	res := ExtractInfoUrl(ext.Url, &ext)
	return res
}

func ExtractInfoUrl(urlstring string, gys *GysRpc) []map[string]string {
	doc := gys2.GetDoc(urlstring)
	typ := gys.Type
	switch typ{
	case "many":
		results := make([]map[string]string,0)
		doc.Find(gys.Selector).Each(func(i int, s *goquery.Selection){
			for _, sub := range gys.Subselectors{
				result := make(map[string]string)
				ExtractSubselector(sub.Selector,sub.Attribute,sub.Default, sub.Name, sub.Split, result, *s)
				result["urlsource"] = urlstring
				results = append(results, result)
			}
		})
		return results
	case "one":
		result := make(map[string]string)
		result["urlsource"] = urlstring
		r := doc.Find(gys.Selector)
		for _, sub := range gys.Subselectors{
			ExtractSubselector(sub.Selector,sub.Attribute,sub.Default, sub.Name, sub.Split, result, *r)
		}
		res := []map[string]string{result}
		return res
	}
	return make([]map[string]string,0)
}

func GetAttr(sel *goquery.Selection, attribute, defaultvalue string) string {
	switch attribute{
	case "text":
		s := sel.Text()
		return s
	case "innerHTML":
		s := sel.Text()
		return s
	case "textContent":
		s := sel.Text()
		return s
	case "outerHTML":
		s, _ := sel.Html()
		return s
	default:
		s := sel.AttrOr(attribute, defaultvalue)
		return s
	}
}

func ExtractGysmain( gys GysMain) []map[string]string {
	ext := gys.Extractor
	urls := strings.Split(ext.Urls, ",")
	results := make([]map[string]string,0)
	for _, url := range urls{
		res := ExtractInfoUrlGysmain(url, &gys)
		results = append(results, res...)
	}
	return results
}

func ExtractInfoUrlGysmain(urlstring string, gys *GysMain) []map[string]string {
	ext := gys.Extractor
	doc := gys2.GetDoc(urlstring)
	typ := ext.Type
	switch typ{
	case "many":
		results := make([]map[string]string,0)
		doc.Find(ext.Selector).Each(func(i int, s *goquery.Selection){
			for _, sub := range ext.Subselectors{
				result := make(map[string]string)
				ExtractSubselector(sub.Selector,sub.Attribute,sub.Default, sub.Name, sub.Split, result, *s)
				result["urlsource"] = urlstring
				results = append(results, result)
			}
		})
		return results
	case "one":
		result := make(map[string]string)
		result["urlsource"] = urlstring
		r := doc.Find(ext.Selector)
		for _, sub := range ext.Subselectors{
			ExtractSubselector(sub.Selector,sub.Attribute,sub.Default, sub.Name, sub.Split, result, *r)
		}
		res := []map[string]string{result}
		return res
	}
	return make([]map[string]string,0)
}


func ExtractSubselector(selector, attribute, defaultvalue, name, split string, result map[string]string, selection goquery.Selection){
	if len(split) > 0{
		selection.Find(selector).Each(func(i int, sel *goquery.Selection) {
			//log.Println(sel.Text())
			s := GetAttr(sel, attribute, defaultvalue)
			//log.Println(s)
			if strings.Contains(s, split) {
				splits := strings.SplitN(s, split, 2)
				result[splits[0]] = splits[1]
			}
		})
	}else{
		sel := selection.Find(selector) //.AttrOr(attribute, defaultvalue)
		s := GetAttr(sel, attribute, defaultvalue)
		result[name] = s
	}
}

func Iterate(gys GysMain) []string {
	links := gys2.GenerateLinks(gys.Iterator.Url, gys.Iterator.Replace, gys.Iterator.Min, gys.Iterator.Max)
	results := make([]string, 0)
	for _, link := range links{
		doc := gys2.GetDoc(link)
		result := gys2.ProcessMessage(doc, gys.Identificator.Selector, gys.Identificator.Attribute,gys.Identificator.Type,gys.Identificator.Default, gys.Identificator.Base)
		results = append(results, result...)
	}
	return results
}

func SHA1(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return string(hex.EncodeToString(algorithm.Sum(nil)))
}
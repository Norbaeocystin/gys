package gysrpc

import (
	"github.com/PuerkitoBio/goquery"
	"gys/pkg/core"
	"log"
	"strings"
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

type Response []map[string]string

type RPCHandler struct {}

func ( h *RPCHandler) Execute(request *GysRpc, response *Response) error {
	log.Println("Extracting info")
	log.Println(request)
	r := Extract(*request)
	*response = r
	return nil
}


func Extract( ext GysRpc) []map[string]string {
	res := ExtractInfoUrl(ext.Url, &ext)
	return res
}

func ExtractInfoUrl(urlstring string, gys *GysRpc) []map[string]string {
	doc := core.GetDoc(urlstring)
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


func ExtractSubselector(selector, attribute, defaultvalue, name, split string, result map[string]string, selection goquery.Selection){
	if len(split) > 0{
		selection.Find(selector).Each(func(i int, sel *goquery.Selection) {
			s := GetAttr(sel, attribute, defaultvalue)
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
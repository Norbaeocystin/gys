package gys

import (
	"github.com/PuerkitoBio/goquery"
	"gys/gysyaml"
	"strings"
)

func Extract( gys gysyaml.Gys) []map[string]string {
	ext := gys.Extractor
	urls := strings.Split(ext.Urls, ",")
	results := make([]map[string]string,0)
	for _, url := range urls{
		res := ExtractInfoUrl(url, &gys)
		results = append(results, res...)
	}
	return results
}

func ExtractInfoUrl(urlstring string, gys *gysyaml.Gys) []map[string]string {
	ext := gys.Extractor
	doc := GetDoc(urlstring)
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
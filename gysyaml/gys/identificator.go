package gys

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"time"
)

func ProcessMessage(doc *goquery.Document, selector, attribute, typ, defaultvalue, base string) []string {
	result := make([]string, 0)
	switch typ{
	case "many":
		doc.Find(selector).Each(func(i int, s *goquery.Selection){
			r := s.AttrOr(attribute, defaultvalue)
			if r != defaultvalue{
				r = base + r
			}
			result = append(result, r)
		})
	case "one":
		r := doc.Find(selector).AttrOr(attribute, defaultvalue)
		result = append(result,r)
	}
	return result
}

func GetDoc(urlstring string) *goquery.Document {
	header := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36"
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest("GET", urlstring, nil)
	if err != nil {
		panic(err)
	}
	//do not forget!!!
	request.Header.Set("User-Agent", header)

	// Make request
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil{
		panic(err)
	}
	return doc
}

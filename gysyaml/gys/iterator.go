package gys

import (
	"strconv"
	"strings"
)

//GenerateLinks
func GenerateLinks(urlstring string, replace string, min, max int) []string {
	result := make([]string,0)
	for i := min; i <= max; i++{
		replacement := strconv.Itoa(i)
		newurlstring := strings.Replace(urlstring, replace, replacement, 1)
		result = append(result, newurlstring)
	}
	return result
}
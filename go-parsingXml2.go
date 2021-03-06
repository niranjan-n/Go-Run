package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Location string
	Keyword  string
}

func main() {

	var s SiteMapIndex
	var n News
	//	newsMap := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	//for-loop
	for _, Location := range s.Locations {
		resp, _ := http.Get(strings.TrimSpace(Location))
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for k, _ := range n.Titles {
			newsMap[n.Titles[k]] = NewsMap{n.Locations[k], n.Keywords[k]}
		}
		for index, data := range newsMap {

			fmt.Println("\n\n\n", index)
			fmt.Println("\n", data.Keyword)
			fmt.Println("\n", data.Location)

		}

	}

}

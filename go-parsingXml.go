package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SiteMapIndex struct {
	Locations []location `xml:"sitemap"`
}

type location struct {
	Loc string `xml:"loc"`
}

func (l location) String() string {
	return fmt.Sprintf(l.Loc)
}
func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	//for-loop
	for _, Location := range s.Locations {
		fmt.Printf("\n %s", Location)
	}

}

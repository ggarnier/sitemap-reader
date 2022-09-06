package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Urlset struct {
	Urls []Url `xml:"url"`
}

type Url struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

func parseXml(filename string) int {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var urlset Urlset
	xml.Unmarshal(byteValue, &urlset)

	return len(urlset.Urls)
}

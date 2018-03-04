package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type StringArray struct {
	Name string   `xml:"name,attr"`
	Item []string `xml:"item"`
}

type String struct {
	Name   string `xml:"name,attr"`
	Item   string `xml:",innerxml"`
	Ignore string `xml:"tools:ignore,attr"`
}

type Resource struct {
	StringArrays []StringArray `xml:"string-array"`
	Strings      []String      `xml:"string"`
}

func main() {
	dat, err := ioutil.ReadFile("./strings_model.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("dat = ", string(dat))

	var res Resource

	err = xml.Unmarshal(dat, &res)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for _, a := range res.StringArrays {
		fmt.Printf("%#v\n", a)
	}
	for _, l := range res.Strings {
		fmt.Printf("%#v\n", l)
	}
}

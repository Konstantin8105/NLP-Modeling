package nlp

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

func LoadResource(lang Language) (lines map[string]string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Error in LoadResource. err = %v", err)
		}
	}()
	var langList []Language
	langList, err = GetLanguageList()
	if err != nil {
		return
	}

	langList = append(langList, lang)

	// Load all resources for all language
	// Load `lang` language at the last
	for _, lang := range langList {
		var l map[string]string
		l, err = load(ResourceLocation + "/" + xmlLocation + "-" + string(lang))
		if err != nil {
			return
		}
		lines = append(lines, l)
	}

	return
}

func load(path string) (lines map[string]string, err error) {
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

package nlp

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	ResourceLocation = "./res"
	xmlLocation      = "values"
)

// Language type of language
type Language string

// GetLanguageList return list of allowable language
func GetLanguageList() (langs []Language, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cannot GetLanguageList. err = %v", err)
		}
	}()
	var list []string
	list, err = filepath.Glob(ResourceLocation + "/*")
	if err != nil {
		return
	}
	for i := 0; i < len(list); i++ {
		if !strings.Contains(list[i], xmlLocation) {
			list = append(list[:i], list[i+1:]...)
			i--
			continue
		}
		index := strings.Index(list[i], xmlLocation)
		list[i] = list[i][index+len(xmlLocation)+1:]
		langs = append(langs, Language(list[i]))
	}
	return
}

type NamedResource interface {
	GetName() string
}

type StringArray struct {
	Name string   `xml:"name,attr"`
	Item []string `xml:"item"`
}

func (sa StringArray) GetName() string {
	return sa.Name
}

type String struct {
	Name   string `xml:"name,attr"`
	Item   string `xml:",innerxml"`
	Ignore string `xml:"tools:ignore,attr"`
}

func (s String) GetName() string {
	return s.Name
}

type Resource struct {
	StringArrays []StringArray `xml:"string-array"`
	Strings      []String      `xml:"string"`
}

func LoadResource(lang Language) (lines map[Language]string, err error) {
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
		var l map[Language]string
		l, err = load(ResourceLocation + "/" + xmlLocation + "-" + string(lang))
		if err != nil {
			return
		}
		_ = l
		// lines[lang] = l
	}

	return
}

func load(path string) (lines map[Language]string, err error) {
	dat, err := ioutil.ReadFile(path + "/" + "strings_model.xml")
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

	return
}

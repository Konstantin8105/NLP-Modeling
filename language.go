package nlp

import (
	"fmt"
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

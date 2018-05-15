package main

import (
	"fmt"

	"github.com/Konstantin8105/nlp"
)

func main() {
	fmt.Println("Start of NLP Modeling")
	l, err := nlp.GetLanguageList()
	fmt.Println("l : ", l)
	fmt.Println("e : ", err)

	ls, err := nlp.LoadResource(nlp.Language("ru"))
	fmt.Println("ls  : ", ls)
	fmt.Println("err : ", err)
}

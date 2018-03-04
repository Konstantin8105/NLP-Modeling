package main

import (
	"fmt"

	nlp "github.com/Konstantin8105/NLP-Modeling"
)

func main() {
	fmt.Println("Start of NLP Modeling")
	l, err := nlp.GetLanguageList()
	fmt.Println("l : ", l)
	fmt.Println("e : ", err)
}

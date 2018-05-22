package nlp

import (
	"testing"
)

func TestResource(t *testing.T) {
	l, err := GetLanguageList()
	if err != nil {
		t.Fatalf("Error : %v", err)
	}
	if len(l) == 0 {
		t.Fatalf("Language list is empty")
	}
	ls, err := LoadResource(Language("ru"))
	if err != nil {
		t.Fatalf("Error resource : %v", err)
	}
	if len(ls) == 0 {
		t.Fatalf("Resource list is empty")
	}
}

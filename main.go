package web

import (
	"fmt"
	"studying/web/vocabulary"
)

func Execute() {
	// words := vocabulary.AddWord(vocabulary.Word{
	// 	Text:       "новий",
	// 	LanguageID: 2,
	// 	Type:       int(vocabulary.Base),
	// })
	words, _ := vocabulary.GetWords(vocabulary.Language{
		ID: 1,
	})
	fmt.Printf("%s", words)
}

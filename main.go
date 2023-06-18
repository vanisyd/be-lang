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
	words := vocabulary.GetWords(vocabulary.Language{
		ID: 2,
	})

	fmt.Printf("%v", words)
}

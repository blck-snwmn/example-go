package main

import (
	"fmt"

	"golang.org/x/text/language"
)

func main() {
	acceptLangageHeader := "fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5"
	{
		matcher := language.NewMatcher([]language.Tag{
			language.Japanese,
			language.English,
			language.French,
		})
		t, i := language.MatchStrings(matcher, acceptLangageHeader)
		fmt.Printf("tag: %v, index: %d\n", t, i)
	}
	{
		matcher := language.NewMatcher([]language.Tag{
			language.French,
			language.Japanese,
			language.English,
		})
		t, i := language.MatchStrings(matcher, acceptLangageHeader)
		fmt.Printf("tag: %v, index: %d\n", t, i)
	}
	{
		matcher := language.NewMatcher([]language.Tag{
			language.English,
			language.French,
			language.Japanese,
		})
		t, i := language.MatchStrings(matcher, acceptLangageHeader)
		fmt.Printf("tag: %v, index: %d\n", t, i)
	}
}

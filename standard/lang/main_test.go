package main

import (
	"testing"

	"golang.org/x/text/language"
)

const acceptLangageHeader = "fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5"

func Test_Matcher1(t *testing.T) {
	matcher := language.NewMatcher([]language.Tag{
		language.Japanese,
		language.English,
		language.French,
	})
	tg, i := language.MatchStrings(matcher, acceptLangageHeader)
	if tg != language.French || i != 2 {
		t.Errorf("tag: %v, index: %d", tg, i)
	}
}

func Test_Matcher2(t *testing.T) {
	matcher := language.NewMatcher([]language.Tag{
		language.French,
		language.Japanese,
		language.English,
	})
	tg, i := language.MatchStrings(matcher, acceptLangageHeader)
	if tg != language.French || i != 0 {
		t.Errorf("tag: %v, index: %d", tg, i)
	}
}

func Test_Matcher3(t *testing.T) {
	matcher := language.NewMatcher([]language.Tag{
		language.English,
		language.French,
		language.Japanese,
	})
	tg, i := language.MatchStrings(matcher, acceptLangageHeader)
	if tg != language.French || i != 1 {
		t.Errorf("tag: %v, index: %d", tg, i)
	}
}

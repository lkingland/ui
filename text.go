package ui

import "github.com/lkingland/html"

type Text struct {
	html.C
}

func NewText(text string) *Text {
	return &Text{C: html.C(text)}
}

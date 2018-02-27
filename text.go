package ui

import "kingland.io/html"

type Text struct {
	html.C
}

func NewText(text string) *Text {
	return &Text{C: html.C(text)}
}

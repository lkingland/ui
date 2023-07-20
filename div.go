package ui

import "github.com/lkingland/html"

type Div struct {
	*html.DIV
}

func NewDiv() *Div {
	return &Div{
		DIV: html.Div(),
	}
}

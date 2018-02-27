package ui

import "kingland.io/html"

type Div struct {
	*html.DIV
}

func NewDiv() *Div {
	return &Div{
		DIV: html.Div(),
	}
}

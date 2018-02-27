package ui

import "kingland.io/html"

type Image struct {
	*html.IMG
}

func NewImage() *Image {
	return &Image{
		IMG: html.Img(),
	}
}

func (i *Image) SetAlt(alt string) *Image {
	i.IMG.Set("alt", alt)
	return i
}

func (i *Image) SetSrc(src string) *Image {
	i.IMG.Set("src", src)
	return i
}

func (i *Image) SetWidth(width string) *Image {
	i.IMG.Set("width", width)
	return i
}

func (i *Image) SetHeight(height string) *Image {
	i.IMG.Set("height", height)
	return i
}

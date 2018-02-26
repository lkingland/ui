package ui

import (
	. "kingland.io/html"
)

const (
	DefaultStyle    = "/css/index.css"
	DefaultScript   = "/js/index.js"
	DefaultLanguage = "en"
)

// Document is an abtraction of a standard html document with a
// large number of omissions, opinions and reductions in scope.
// Rather than expand this abstractin with too many mutative acccessors
// and the requisite increase in states, just create a new abstraction
// for a tight set of permutations, or create the HTML elements directly.
// The intended use of this document is for "paved path" docs with little
// variance.  As such it will not work for many, even most, use cases.
type Document struct {
	id       string
	title    string
	desc     string
	keywords string
	style    string
	script   string
	language string

	children []Renderable
}

func NewDocument() *Document {
	return &Document{
		style:    DefaultStyle,
		script:   DefaultScript,
		language: DefaultLanguage,
	}
}

func (s *Document) SetID(id string) *Document {
	s.id = id
	return s
}

func (s *Document) SetTitle(text string) *Document {
	s.title = text
	return s
}

func (s *Document) SetDesc(text string) *Document {
	s.desc = text
	return s
}

func (s *Document) SetKeywords(text string) *Document {
	s.keywords = text
	return s
}

func (s *Document) SetStyle(style string) *Document {
	s.style = style
	return s
}

func (s *Document) SetScript(script string) *Document {
	s.script = script
	return s
}

func (s *Document) Add(r ...Renderable) *Document {
	s.children = append(s.children, r...)
	return s
}

func (s *Document) Render(i int) string {
	root := Root().
		Add(Doctype()).
		Add(Html().Set("lang", s.language).
			Add(Head().
				Add(Title().Add(C(s.title))).
				Add(Base().Set("href", "/")).
				Add(Meta().Set("charset", "UTF-8")).
				Add(Meta().Set("viewport", "width=device-width, initial-scale=1")).
				Add(Meta().Set("description", s.desc)).
				Add(Meta().Set("keywords", s.keywords)).
				Add(Link().Set("type", "text/css").Set("rel", "stylesheet").Set("href", s.style)).
				Add(Script().Set("type", "text/javascript").Set("src", s.script))).
			Add(Body().Set("id", s.id).
				Add(s.children...)))

	return root.Render(i)
}

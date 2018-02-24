package ui

import (
	"fmt"
	"time"

	. "kingland.io/html"
)

const (
	DefaultBanner = "/img/splash.jpg"
	DefaultLogo   = "/img/logo.png"
)

// Splash is a simplistic splashscreen with minimal configuraiton,
// suitable as a placeholder for new web locations.
type Splash struct {
	title    string
	desc     string
	keywords string
	copyYear int
	version  string
}

func NewSplash() *Splash {
	return &Splash{}
}

func (s *Splash) SetTitle(text string) {
	s.title = text
}

func (s *Splash) SetDesc(text string) {
	s.desc = text
}

func (s *Splash) SetKeywords(text string) {
	s.keywords = text
}

func (s *Splash) SetCopyrightYear(year int) {
	s.copyYear = year
}

func (s *Splash) SetVersion(version string) {
	s.version = version
}

func (s *Splash) copyrightNotice() string {
	return fmt.Sprintf("&copy; %d - %d", s.copyYear, time.Now().Year())
}

func (s *Splash) Bytes() []byte {
	root := Root().
		Add(Doctype()).
		Add(Html().Set("lang", "en").
			Add(Head().
				Add(Title().Add(C(s.title))).
				Add(Base().Set("href", "/")).
				Add(Meta().Set("charset", "UTF-8")).
				Add(Meta().Set("viewport", "width=device-width, initial-scale=1")).
				Add(Link().Set("type", "text/css").Set("rel", "stylesheet").Set("href", "/css/splash.css")).
				Add(Script().Set("type", "text/javascript").Set("src", "/js/splash.js"))).
			Add(Body().
				Add(Img().Set("src", DefaultBanner)).
				Add(Img().Set("src", DefaultLogo)).
				Add(Div().Add(C(s.copyrightNotice()))).
				Add(Div().Add(C(s.version)))))

	return []byte(root.Render(0))
}

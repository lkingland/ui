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
	*Document
	copyYear int
	version  string
}

func NewSplash() *Splash {
	return &Splash{
		Document: NewDocument(), // Inherit defaults
	}
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

func (s *Splash) Render(i int) string {
	s.Document.
		Add(Img().Set("src", DefaultBanner).Set("class", "banner")).
		Add(Img().Set("src", DefaultLogo).Set("class", "logo").
			Set("width", "300").Set("height", "60")).
		Add(Div().Set("class", "copyright").Add(C(s.copyrightNotice()))).
		Add(Div().Set("class", "version").Add(C(s.version)))

	return s.Document.Render(i)
}

package ui

import (
	"fmt"
	"time"

	"github.com/lkingland/html"
)

const (
	DefaultBanner     = "/img/splash.jpg"
	DefaultLogo       = "/img/logo.png"
	DefaultLogoWidth  = "auto"
	DefaultLogoHeight = "auto"
)

// Splash is a simplistic splashscreen with minimal configuraiton,
// suitable as a placeholder for new web locations.
type Splash struct {
	*Document
	copyYear    int
	version     string
	banner      string
	logo        string
	logoWidth   string
	logoHeight  string
	disableLogo bool
}

func NewSplash() *Splash {
	return &Splash{
		banner:     DefaultBanner,
		logo:       DefaultLogo,
		logoWidth:  DefaultLogoWidth,
		logoHeight: DefaultLogoHeight,
		Document:   NewDocument(), // Inherit defaults
	}
}

func (s *Splash) SetDisableLogo(disable bool) {
	s.disableLogo = disable
}

func (s *Splash) SetLogoWidth(width string) {
	s.logoWidth = width
}

func (s *Splash) SetLogoHeight(height string) {
	s.logoHeight = height
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
		Add(html.Img().Set("src", DefaultBanner).Set("class", "banner"))
	if !s.disableLogo {
		s.Document.Add(html.Img().Set("src", DefaultLogo).Set("class", "logo").
			Set("width", s.logoWidth).Set("height", s.logoHeight))
	}
	s.Document.
		Add(html.Div().Set("class", "copyright").Add(html.C(s.copyrightNotice()))).
		Add(html.Div().Set("class", "version").Add(html.C(s.version)))

	return s.Document.Render(i)
}

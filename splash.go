package ui

import (
	"fmt"
	"time"

	. "kingland.io/html"
)

const (
	DefaultBanner     = "/img/splash.jpg"
	DefaultLogo       = "/img/logo.png"
	DefaultLogoWidth  = -1
	DefaultLogoHeight = -1
)

// Splash is a simplistic splashscreen with minimal configuraiton,
// suitable as a placeholder for new web locations.
type Splash struct {
	*Document
	copyYear   int
	version    string
	banner     string
	logo       string
	logoWidth  int
	logoHeight int
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

func (s *Splash) SetLogoWidth(width int) {
	s.logoWidth = width
}

func (s *Splash) SetLogoHeight(height int) {
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
		Add(Img().Set("src", DefaultBanner).Set("class", "banner")).
		Add(Img().Set("src", DefaultLogo).Set("class", "logo").
			Set("width", px(s.logoWidth)).Set("height", px(s.logoHeight))).
		Add(Div().Set("class", "copyright").Add(C(s.copyrightNotice()))).
		Add(Div().Set("class", "version").Add(C(s.version)))

	return s.Document.Render(i)
}

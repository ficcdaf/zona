package builder

import (
	"embed"
)

const (
	DefaultHeader     = ""
	DefaultFooter     = ""
	DefaultStylesheet = "/style/zonaDefault.css"
	DefaultIcon       = ""
	DefaultTemplate   = `<!doctype html>
<html>
  <head>
    <title>{{ .Title }}</title>
    <link rel="icon" href="{{ .Icon }}" type="image/x-icon" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta charset="UTF-8" />
    <link
      href="{{ .Stylesheet }}"
      rel="stylesheet"
      type="text/css"
      media="all"
    />
  </head>
  <body>
    <div id="container">
      <header id="header">{{ .Header }}</header>
      <article id="content">
        {{ .Content }}
        <nav id="nextprev">
          {{ .NextPost }}<br />
          {{ .PrevPost }}
        </nav>
      </article>
      <footer id="footer">{{ .Footer }}</footer>
    </div>
  </body>
</html>`
)

//go:embed embed
var embedDir embed.FS

type Settings struct {
	Header          string
	Footer          string
	Stylesheet      string
	Icon            string
	DefaultTemplate string
}

func NewSettings(header string, footer string, style string, icon string, temp string) *Settings {
	return &Settings{
		header,
		footer,
		style,
		icon,
		temp,
	}
}

func GetSettings() *Settings {
	// TODO: Read a config file to override defaults
	// "Defaults" should be a default config file via embed package,
	// so the settings func should need to handle one case:
	// check if config file exists, if not, use embedded one
	return NewSettings(DefaultHeader, DefaultFooter, DefaultStylesheet, DefaultIcon, DefaultTemplate)
}

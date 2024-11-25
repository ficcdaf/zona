package build

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

type Settings struct {
	Header     string
	Footer     string
	Stylesheet string
	Icon       string
}

func NewSettings(header string, footer string, style string, icon string) *Settings {
	return &Settings{
		header,
		footer,
		style,
		icon,
	}
}

package builder

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ficcdaf/zona/internal/util"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// This function takes a Markdown document and returns an HTML document.
func MdToHTML(md []byte) []byte {
	// create parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// build HTML renderer
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := newZonaRenderer(opts)

	return markdown.Render(doc, renderer)
}

// PathIsValid checks if a path is valid.
// If requireFile is set, directories are not considered valid.
func PathIsValid(path string, requireFile bool) bool {
	s, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else if requireFile {
		// fmt.Printf("Directory status: %s\n", strconv.FormatBool(s.IsDir()))
		return !s.IsDir()
	}
	return err == nil
}

func processLink(p string) string {
	// fmt.Println("Processing link...")
	ext := filepath.Ext(p)
	// Only process if it points to an existing, local markdown file
	if ext == ".md" && filepath.IsLocal(p) {
		// fmt.Println("Markdown link detected...")
		return util.ChangeExtension(p, ".html")
	} else {
		return p
	}
}

func renderImage(w io.Writer, node *ast.Image, entering bool) {
	// we add image-container div tag
	// here before the opening img tag
	if entering {
		fmt.Fprintf(w, "<div class=\"image-container\">\n")
		fmt.Fprintf(w, `<img src="%s" title="%s">`, node.Destination, node.Title)
	} else {
		// if it's the closing img tag
		// we close the div tag *after*
		fmt.Fprintf(w, `</div>`)
		fmt.Println("Image node not entering??")
	}
}

func renderLink(w io.Writer, l *ast.Link, entering bool) {
	if entering {
		destPath := processLink(string(l.Destination))
		fmt.Fprintf(w, `<a href="%s"`, destPath)
		for _, attr := range html.BlockAttrs(l) {
			fmt.Fprintf(w, ` %s`, attr)
		}
		io.WriteString(w, ">")
	} else {
		io.WriteString(w, "</a>")
	}
}

// htmlRenderHook hooks the HTML renderer and overrides the rendering of certain nodes.
func htmlRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if link, ok := node.(*ast.Link); ok {
		renderLink(w, link, entering)
		return ast.GoToNext, true
	} else if image, ok := node.(*ast.Image); ok {
		// TODO: should do something more interesting with the alt text -- like put it in a <small> tag?
		renderImage(w, image, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func newZonaRenderer(opts html.RendererOptions) *html.Renderer {
	opts.RenderNodeHook = htmlRenderHook
	return html.NewRenderer(opts)
}

package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"strings"

	"github.com/ficcdaf/zona/internal/util"
	"gopkg.in/yaml.v3"
)

type PageData struct {
	Title      string
	Icon       string
	Stylesheet string
	HeaderName string
	Header     template.HTML
	Content    template.HTML
	NextPost   string
	PrevPost   string
	FooterName string
	Footer     template.HTML
	Template   string
}

type Metadata map[string]interface{}

func processWithYaml(f []byte) (Metadata, []byte, error) {
	// Check if the file has valid metadata
	if !bytes.HasPrefix(f, []byte("---\n")) {
		// No valid yaml, so return the entire content
		return nil, f, nil
	}
	// Separate YAML from rest of document
	split := strings.SplitN(string(f), "---\n", 3)
	if len(split) < 3 {
		return nil, nil, fmt.Errorf("Invalid frontmatter format.")
	}
	var meta Metadata
	// Parse YAML
	if err := yaml.Unmarshal([]byte(split[1]), &meta); err != nil {
		return nil, nil, err
	}
	return meta, []byte(split[2]), nil
}

func buildPageData(m Metadata, in string, out string, settings *Settings) *PageData {
	p := &PageData{}
	if title, ok := m["title"].(string); ok {
		p.Title = util.WordsToTitle(title)
	} else {
		p.Title = util.PathToTitle(in)
	}
	if icon, ok := m["icon"].(string); ok {
		p.Icon = icon
	} else {
		p.Icon = settings.IconName
	}
	var stylePath string
	if style, ok := m["style"].(string); ok {
		stylePath = style
	} else {
		stylePath = settings.StylePath
	}
	curDir := filepath.Dir(out)
	relPath, err := filepath.Rel(curDir, stylePath)
	// fmt.Printf("fp: %s, sp: %s, rp: %s\n", curDir, stylePath, relPath)
	if err != nil {
		log.Fatalln("Error calculating stylesheet path: ", err)
	}
	p.Stylesheet = relPath
	if header, ok := m["header"].(string); ok {
		p.HeaderName = header
		// for now we use default anyways
		p.Header = settings.Header
	} else {
		p.HeaderName = settings.HeaderName
		p.Header = settings.Header
	}
	if footer, ok := m["footer"].(string); ok {
		p.FooterName = footer
		p.Footer = settings.Footer
	} else {
		p.FooterName = settings.FooterName
		p.Footer = settings.Footer
	}
	if t, ok := m["type"].(string); ok && t == "article" || t == "post" {
		p.Template = (settings.ArticleTemplate)
	} else {
		p.Template = (settings.DefaultTemplate)
	}
	return p
}

func ConvertFile(in string, out string, settings *Settings) error {
	mdPre, err := util.ReadFile(in)
	if err != nil {
		return err
	}
	metadata, md, err := processWithYaml(mdPre)
	if err != nil {
		return err
	}
	pd := buildPageData(metadata, in, out, settings)
	fmt.Println("Title: ", pd.Title)

	// build according to template here
	html := MdToHTML(md)
	pd.Content = template.HTML(html)

	tmpl, err := template.New("webpage").Parse(pd.Template)
	if err != nil {
		return err
	}

	var output bytes.Buffer
	if err := tmpl.Execute(&output, pd); err != nil {
		return err
	}

	err = util.WriteFile(output.Bytes(), out)
	return err
}

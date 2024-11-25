package build

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type PageData struct {
	Title      string
	Icon       string
	Stylesheet string
	Header     string
	Content    template.HTML
	NextPost   string
	PrevPost   string
	Footer     string
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

// this function converts a file path to its title form
func pathToTitle(path string) string {
	stripped := ChangeExtension(filepath.Base(path), "")
	replaced := strings.NewReplacer("-", " ", "_", " ", `\ `, " ").Replace(stripped)
	return strings.ToTitle(replaced)
}

func buildPageData(m Metadata, path string) *PageData {
	p := &PageData{}
	if title, ok := m["title"].(string); ok {
		p.Title = title
	} else {
		p.Title = pathToTitle(path)
	}
	if icon, ok := m["icon"].(string); ok {
		p.Icon = icon
	} else {
		p.Icon = DefaultIcon
	}
	if style, ok := m["style"].(string); ok {
		p.Stylesheet = style
	} else {
		p.Stylesheet = DefaultStylesheet
	}
	if header, ok := m["header"].(string); ok {
		p.Header = header
	} else {
		p.Header = DefaultHeader
	}
	if footer, ok := m["footer"].(string); ok {
		p.Footer = footer
	} else {
		p.Footer = DefaultFooter
	}
	return p
}

func ConvertFile(in string, out string) error {
	mdPre, err := ReadFile(in)
	if err != nil {
		return err
	}
	metadata, md, err := processWithYaml(mdPre)
	if err != nil {
		return err
	}
	pd := buildPageData(metadata, in)

	// build according to template here
	html, err := MdToHTML(md)
	if err != nil {
		return err
	}
	pd.Content = template.HTML(html)

	tmpl, err := template.New("webpage").Parse(DefaultTemplate)
	if err != nil {
		return err
	}

	var output bytes.Buffer
	if err := tmpl.Execute(&output, pd); err != nil {
		return err
	}

	err = WriteFile(output.Bytes(), out)
	return err
}

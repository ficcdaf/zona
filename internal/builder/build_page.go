package builder

import (
	"bytes"
	"errors"
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
	Type       string
}

type Metadata map[string]any

type FrontMatter struct {
	Title  string `yaml:"title"`
	Icon   string `yaml:"icon"`
	Style  string `yaml:"style"`
	Header string `yaml:"header"`
	Footer string `yaml:"footer"`
	Type   string `yaml:"type"`
}

func processWithYaml(f []byte) (*FrontMatter, []byte, error) {
	// Check if the file has valid metadata
	trimmed := bytes.TrimSpace(f)
	normalized := strings.ReplaceAll(string(trimmed), "\r\n", "\n")
	if !strings.HasPrefix(normalized, ("---\n")) {
		// No valid yaml, so return the entire content
		return nil, f, nil
	}
	// Separate YAML from rest of document
	split := strings.SplitN(normalized, "---\n", 3)
	if len(split) < 3 {
		return nil, nil, fmt.Errorf("invalid frontmatter format")
	}
	var meta FrontMatter
	// Parse YAML
	if err := yaml.Unmarshal([]byte(split[1]), &meta); err != nil {
		return nil, nil, err
	}
	return &meta, []byte(split[2]), nil
}

func buildPageData(m *FrontMatter, in string, out string, settings *Settings) *PageData {
	p := &PageData{}
	if m != nil && m.Title != "" {
		p.Title = util.WordsToTitle(m.Title)
	} else {
		p.Title = util.PathToTitle(in)
	}
	if m != nil && m.Icon != "" {
		i, err := util.NormalizePath(m.Icon, in)
		if err != nil {
			p.Icon = settings.IconName
		} else {
			p.Icon = i
		}
	} else {
		p.Icon = settings.IconName
	}
	var stylePath string
	if m != nil && m.Style != "" {
		stylePath = m.Style
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

	if m != nil && m.Header != "" {
		p.HeaderName = m.Header
		// for now we use default anyways
		p.Header = settings.Header
	} else {
		p.HeaderName = settings.HeaderName
		p.Header = settings.Header
	}
	if m != nil && m.Footer != "" {
		p.FooterName = m.Footer
		p.Footer = settings.Footer
	} else {
		p.FooterName = settings.FooterName
		p.Footer = settings.Footer
	}
	// TODO: Don't hard code posts dir name
	if (m != nil && (m.Type == "article" || m.Type == "post")) || util.InDir(in, "posts") {
		p.Template = (settings.ArticleTemplate)
		p.Type = "post"
	} else {
		p.Template = (settings.DefaultTemplate)
		p.Type = ""
	}
	return p
}

func _BuildHtmlFile(in string, out string, settings *Settings) error {
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

func BuildFile(f *File, settings *Settings) error {
	if f.ShouldCopy {
		if err := util.CreateParents(f.OutPath); err != nil {
			return err
		}
		if err := util.CopyFile(f.InPath, f.OutPath); err != nil {
			return errors.Join(errors.New("Error processing file "+f.InPath), err)
		} else {
			return nil
		}
	}

	if err := util.CreateParents(f.OutPath); err != nil {
		return err
	}
	if err := BuildHtmlFile(f.FrontMatterLen, f.InPath, f.OutPath, f.PageData, settings); err != nil {
		return errors.Join(errors.New("Error processing file "+f.InPath), err)
	} else {
		return nil
	}
}

func BuildHtmlFile(l int, in string, out string, pd *PageData, settings *Settings) error {
	// WARN: ReadLineRange is fine, but l is the len of the frontmatter
	// NOT including the delimiters!
	start := l
	// if the frontmatter exists (len > 0), then we need to
	// account for two lines of delimiter!
	if l != 0 {
		start += 2
	}
	md, err := util.ReadLineRange(in, start, -1)
	if err != nil {
		return err
	}
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

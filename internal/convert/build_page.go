package convert

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"gopkg.in/yaml.v3"
)

type PageData struct {
	Title      string
	Icon       string
	Stylesheet string
	Header     template.HTML
	Content    template.HTML
	NextPost   template.HTML
	PrevPost   template.HTML
	Footer     template.HTML
}

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
	var metadata Metadata
	// Parse YAML
	if err := yaml.Unmarshal([]byte(split[1]), &metadata); err != nil {
		return nil, nil, err
	}
	return metadata, []byte(split[2]), nil
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

	title, ok := metadata["title"].(string)
	if !ok {
		fmt.Println("No title in page.")
	} else {
		fmt.Println("Title found: ", title)
	}

	html, err := MdToHTML(md)
	if err != nil {
		return err
	}
	err = WriteFile(html, out)
	return err
}

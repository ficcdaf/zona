package builder

import (
	"embed"
	"html/template"
	"log"
	"path/filepath"

	"github.com/ficcdaf/zona/internal/util"
	"gopkg.in/yaml.v3"
)

const (
	DefConfigName       = "config.yml"
	DefHeaderName       = "header.md"
	DefFooterName       = "footer.md"
	DefStylesheetName   = "style.css"
	DefIconName         = "icon.png"
	DefTemplateName     = "default.html"
	ArticleTemplateName = "article.html"
)

//go:embed embed
var embedDir embed.FS

type Settings struct {
	Header              template.HTML
	HeaderName          string
	Footer              template.HTML
	FooterName          string
	StylesheetName      string
	IconName            string
	DefaultTemplate     string
	DefaultTemplateName string
	ArticleTemplate     string
	Stylesheet          []byte
	Icon                []byte
}

func buildSettings(f []byte) (*Settings, error) {
	s := &Settings{}
	var c map[string]interface{}
	// Parse YAML
	if err := yaml.Unmarshal(f, &c); err != nil {
		return nil, err
	}
	if headerName, ok := c["header"].(string); ok {
		header, err := util.ReadFile(headerName)
		s.HeaderName = headerName
		if err != nil {
			return nil, util.ErrorPrepend("Could not read header specified in config: ", err)
		}
		s.Header = template.HTML(MdToHTML(header))
	} else {
		header := readEmbed(DefHeaderName)
		s.Header = template.HTML(MdToHTML(header))
		s.HeaderName = DefHeaderName
	}
	if footerName, ok := c["footer"].(string); ok {
		footer, err := util.ReadFile(footerName)
		s.FooterName = footerName
		if err != nil {
			return nil, util.ErrorPrepend("Could not read footer specified in config: ", err)
		}
		s.Footer = template.HTML(MdToHTML(footer))
	} else {
		footer := readEmbed(DefFooterName)
		s.Footer = template.HTML(MdToHTML(footer))
		s.FooterName = DefFooterName
	}
	if stylesheetName, ok := c["stylesheet"].(string); ok {
		stylesheet, err := util.ReadFile(stylesheetName)
		if err != nil {
			return nil, util.ErrorPrepend("Could not read stylesheet specified in config: ", err)
		}
		s.StylesheetName = stylesheetName
		s.Stylesheet = stylesheet
	} else {
		stylesheet := readEmbed(DefStylesheetName)
		s.Stylesheet = stylesheet
		s.StylesheetName = DefStylesheetName
	}
	if iconName, ok := c["icon"].(string); ok {
		icon, err := util.ReadFile(iconName)
		if err != nil {
			return nil, util.ErrorPrepend("Could not read icon specified in config: ", err)
		}
		s.Icon = icon
		s.IconName = iconName
	} else {
		icon := readEmbed(DefIconName)
		s.Icon = icon
		s.IconName = DefIconName
	}
	if templateName, ok := c["template"].(string); ok {
		temp, err := util.ReadFile(templateName)
		if err != nil {
			return nil, util.ErrorPrepend("Could not read template specified in config: ", err)
		}
		s.DefaultTemplate = string(temp)
		s.DefaultTemplateName = templateName
	} else {
		temp := readEmbed(DefTemplateName)
		s.DefaultTemplate = string(temp)
		s.DefaultTemplateName = DefTemplateName
	}
	artTemp := readEmbed(ArticleTemplateName)
	s.ArticleTemplate = string(artTemp)

	return s, nil
}

// readEmbed reads a file inside the embedded dir
func readEmbed(name string) []byte {
	f, err := embedDir.ReadFile(name)
	if err != nil {
		log.Fatalln("Fatal internal error: Could not read embedded default config!")
	}
	return f
}

func GetSettings(root string) *Settings {
	var config []byte
	configPath := filepath.Join(root, DefConfigName)
	if !util.FileExists(configPath) {
		// Config file does not exist, we used embedded default
		config = readEmbed(configPath)
	} else {
		var err error
		config, err = util.ReadFile(configPath)
		if err != nil {
			log.Fatalln("Fatal internal error: Config file exists but could not be read!")
		}
	}
	s, err := buildSettings(config)
	if err != nil {
		log.Fatalf("Fatal error: could not parse config: %u\n", err)
	}
	return s
}

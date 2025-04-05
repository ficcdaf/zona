package builder

import (
	"embed"
	"html/template"
	"log"
	"path/filepath"

	"github.com/ficcdaf/zona/internal/util"
	"gopkg.in/yaml.v3"
)

var defaultNames = map[string]string{
	"config":    ".zona.yml",
	"header":    "header.md",
	"footer":    "footer.md",
	"style":     "style.css",
	"stylePath": filepath.Join("style", "style.css"),
	"icon":      "favicon.png",
	"article":   "article.html",
	"template":  "default.html",
}

//go:embed embed/article.html
//go:embed embed/.zona.yml
//go:embed embed/default.html
//go:embed embed/favicon.png
//go:embed embed/footer.md
//go:embed embed/header.md
//go:embed embed/style.css
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
	StylePath           string
	Icon                []byte
}

var isDefaultStyle bool

// processSetting checks the user's configuration for
// each option. If set, reads the specified file. If not,
// default option is used.
func processSetting(c map[string]interface{}, s string) (string, []byte, error) {
	if name, ok := c[s].(string); ok {
		val, err := util.ReadFile(name)
		if err != nil {
			return "", nil, util.ErrorPrepend("Could not read "+s+" specified in config: ", err)
		}
		return name, val, nil
	} else {
		val := readEmbed(defaultNames[s])
		isDefaultStyle = true
		return defaultNames[s], val, nil
	}
}

// buildSettings constructs the Settings struct.
func buildSettings(f []byte, outRoot string) (*Settings, error) {
	s := &Settings{}
	var c map[string]interface{}
	// Parse YAML
	if err := yaml.Unmarshal(f, &c); err != nil {
		return nil, err
	}
	n, v, err := processSetting(c, "header")
	if err != nil {
		return nil, err
	}
	s.HeaderName = n
	s.Header = template.HTML(MdToHTML(v))
	n, v, err = processSetting(c, "footer")
	if err != nil {
		return nil, err
	}
	s.FooterName = n
	s.Footer = template.HTML(MdToHTML(v))
	isDefaultStyle = false
	n, v, err = processSetting(c, "style")
	if err != nil {
		return nil, err
	}
	s.StylesheetName = n
	s.Stylesheet = v

	if isDefaultStyle {
		stylePath := filepath.Join(outRoot, defaultNames["stylePath"])
		s.StylePath = stylePath
		err := util.CreateParents(stylePath)
		if err != nil {
			return nil, util.ErrorPrepend("Could not create default stylesheet directory: ", err)
		}
		err = util.WriteFile(s.Stylesheet, stylePath)
		if err != nil {
			return nil, util.ErrorPrepend("Could not create default stylesheet: ", err)
		}
	}

	n, v, err = processSetting(c, "icon")
	if err != nil {
		return nil, err
	}
	s.IconName = n
	s.Icon = v
	n, v, err = processSetting(c, "template")
	if err != nil {
		return nil, err
	}
	s.DefaultTemplateName = n
	s.DefaultTemplate = string(v)
	artTemp := readEmbed(string(defaultNames["article"]))
	s.ArticleTemplate = string(artTemp)

	return s, nil
}

// readEmbed reads a file inside the embedded dir
func readEmbed(name string) []byte {
	f, err := embedDir.ReadFile("embed/" + name)
	if err != nil {
		// panic(0)
		log.Fatalf("Fatal internal error: Could not read embedded default %s! %u", name, err)
	}
	return f
}

func GetSettings(root string, outRoot string) *Settings {
	var config []byte
	configPath := filepath.Join(root, defaultNames["config"])
	if !util.FileExists(configPath) {
		// Config file does not exist, we used embedded default
		config = readEmbed(defaultNames["config"])
	} else {
		var err error
		// config, err = util.ReadFile(filepath.Join(root, configPath))
		config, err = util.ReadFile(configPath)
		if err != nil {
			log.Fatalln("Fatal internal error: Config file exists but could not be read!", err)
		}
	}
	s, err := buildSettings(config, outRoot)
	if err != nil {
		log.Fatalf("Fatal error: could not parse config: %u\n", err)
	}
	return s
}

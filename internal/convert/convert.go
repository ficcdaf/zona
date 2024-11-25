package convert

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/ficcdaf/zona/internal/util"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// This function takes a Markdown document and returns an HTML document.
func MdToHTML(md []byte) ([]byte, error) {
	// create parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// build HTML renderer
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := newZonaRenderer(opts)

	return markdown.Render(doc, renderer), nil
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
	fmt.Println("Processing link...")
	ext := filepath.Ext(p)
	// Only process if it points to an existing, local markdown file
	if ext == ".md" && filepath.IsLocal(p) {
		fmt.Println("Markdown link detected...")
		return ChangeExtension(p, ".html")
	} else {
		return p
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

func htmlRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if link, ok := node.(*ast.Link); ok {
		renderLink(w, link, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func newZonaRenderer(opts html.RendererOptions) *html.Renderer {
	opts.RenderNodeHook = htmlRenderHook
	return html.NewRenderer(opts)
}

// WriteFile writes a given byte array to the given path.
func WriteFile(b []byte, p string) error {
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

// ReadFile reads a byte array from a given path.
func ReadFile(p string) ([]byte, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	var result []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		// check for a non EOF error
		if err != nil && err != io.EOF {
			return nil, err
		}
		// n==0 when there are no chunks left to read
		if n == 0 {
			defer f.Close()
			break
		}
		result = append(result, buf[:n]...)
	}
	return result, nil
}

// CopyFile reads the file at the input path, and write
// it to the output path.
func CopyFile(inPath string, outPath string) error {
	inB, err := ReadFile(inPath)
	if err != nil {
		return err
	}
	if err := WriteFile(inB, outPath); err != nil {
		return err
	} else {
		return nil
	}
}

func ChangeExtension(in string, outExt string) string {
	return strings.TrimSuffix(in, filepath.Ext(in)) + outExt
}

func processFile(inPath string, entry fs.DirEntry, err error, outRoot string) error {
	if err != nil {
		return err
	}
	if !entry.IsDir() {
		ext := filepath.Ext(inPath)
		outPath := util.ReplaceRoot(inPath, outRoot)
		switch ext {
		case ".md":
			fmt.Println("Processing markdown...")
			outPath = ChangeExtension(outPath, ".html")
			if err := util.CreateParents(outPath); err != nil {
				return err
			}
			if err := ConvertFile(inPath, outPath); err != nil {
				return errors.Join(errors.New("Error processing file "+inPath), err)
			} else {
				return nil
			}
		// If it's not a file we need to process,
		// we simply copy it to the destination path.
		default:
			if err := util.CreateParents(outPath); err != nil {
				return err
			}
			if err := CopyFile(inPath, outPath); err != nil {
				return errors.Join(errors.New("Error processing file "+inPath), err)
			} else {
				return nil
			}
		}
	}
	// fmt.Printf("Visited: %s\n", inPath)
	return nil
}

func Traverse(root string, outRoot string) error {
	// err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
	walkFunc := func(path string, entry fs.DirEntry, err error) error {
		return processFile(path, entry, err, outRoot)
	}
	err := filepath.WalkDir(root, walkFunc)
	return err
}

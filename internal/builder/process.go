package builder

import (
	"io/fs"
	"path/filepath"

	"github.com/ficcdaf/zona/internal/util"
)

type ProcessMemory struct {
	// Pages holds all page data that may be
	// needed while building *other* pages.
	Pages []*Page
	// Queue is a FIFO queue of Pages indexes to be built.
	// queue should be constructed after all the Pages have been parsed
	Queue []int
	// Posts is an array of pointers to post pages
	Posts []*Page
}

type Page struct {
	Data    *PageData
	Ext     string
	InPath  string
	OutPath string
	Copy    bool
}

// processFile processes the metadata only
// of each file
func processFile(inPath string, entry fs.DirEntry, err error, outRoot string, settings *Settings, pm *ProcessMemory) error {
	if err != nil {
		return err
	}
	var toProcess bool
	var outPath string
	var ext string
	if !entry.IsDir() {
		ext = filepath.Ext(inPath)
		outPath = util.ReplaceRoot(inPath, outRoot)
		switch ext {
		case ".md":
			// fmt.Println("Processing markdown...")
			toProcess = true
			outPath = util.ChangeExtension(outPath, ".html")
		// If it's not a file we need to process,
		// we simply copy it to the destination path.
		default:
			toProcess = false
		}
	}
	page := &Page{
		nil,
		ext,
		inPath,
		outPath,
		!toProcess,
	}
	if toProcess {
		// process its frontmatter here
		m, err := processFrontmatter(inPath)
		if err != nil {
			return err
		}
		pd := buildPageData(m, inPath, outPath, settings)
		if pd.Type == "post" {
			pm.Posts = append(pm.Posts, page)
		}
		page.Data = pd
	}
	pm.Pages = append(pm.Pages, page)
	return nil
}

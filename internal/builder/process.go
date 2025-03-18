package builder

import (
	"io/fs"
	"path/filepath"

	"github.com/ficcdaf/zona/internal/util"
)

type ProcessMemory struct {
	// Files holds all page data that may be
	// needed while building *other* pages.
	Files []*File
	// Queue is a FIFO queue of Pages indexes to be built.
	// queue should be constructed after all the Pages have been parsed
	Queue []int
	// Posts is an array of pointers to post pages
	// This list is ONLY referenced for generating
	// the archive, NOT by the build process!
	Posts []*File
}

type File struct {
	PageData       *PageData
	Ext            string
	InPath         string
	OutPath        string
	ShouldCopy     bool
	HasFrontmatter bool
	FrontMatterLen int
}

// NewProcessMemory initializes an empty
// process memory structure
func NewProcessMemory() *ProcessMemory {
	f := make([]*File, 0)
	q := make([]int, 0)
	p := make([]*File, 0)
	pm := &ProcessMemory{
		f,
		q,
		p,
	}
	return pm
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
	if entry.IsDir() {
		return nil
	} else {
		ext = filepath.Ext(inPath)
		outPath = util.ReplaceRoot(inPath, outRoot)
		// NOTE: This could be an if statement, but keeping
		// the switch makes it easy to extend the logic here later
		switch ext {
		case ".md":
			toProcess = true
			outPath = util.ChangeExtension(outPath, ".html")
		default:
			toProcess = false
		}
	}

	var pd *PageData
	hasFrontmatter := false
	l := 0
	if toProcess {
		// process its frontmatter here
		m, le, err := processFrontmatter(inPath)
		l = le
		if err != nil {
			return err
		}
		if m != nil {
			hasFrontmatter = true
		}
		pd = buildPageData(m, inPath, outPath, settings)

	} else {
		pd = nil
	}
	file := &File{
		pd,
		ext,
		inPath,
		outPath,
		!toProcess,
		hasFrontmatter,
		l,
	}
	if pd != nil && pd.Type == "post" {
		pm.Posts = append(pm.Posts, file)
	}
	pm.Files = append(pm.Files, file)
	return nil
}

func BuildProcessedFiles(pm *ProcessMemory, settings *Settings) error {
	for _, f := range pm.Files {
		err := BuildFile(f, settings)
		if err != nil {
			return err
		}
	}
	return nil
}

package builder

import (
	"errors"
	"io/fs"
	"path/filepath"

	"github.com/ficcdaf/zona/internal/util"
)

// TODO: Process the metadata and build a queue of files to convert here instead of converting them immediately
func buildFile(inPath string, entry fs.DirEntry, err error, outRoot string, settings *Settings) error {
	if err != nil {
		return err
	}
	if !entry.IsDir() {
		ext := filepath.Ext(inPath)
		outPath := util.ReplaceRoot(inPath, outRoot)
		switch ext {
		case ".md":
			// fmt.Println("Processing markdown...")
			outPath = util.ChangeExtension(outPath, ".html")
			if err := util.CreateParents(outPath); err != nil {
				return err
			}
			if err := BuildHtmlFile(inPath, outPath, settings); err != nil {
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
			if err := util.CopyFile(inPath, outPath); err != nil {
				return errors.Join(errors.New("Error processing file "+inPath), err)
			} else {
				return nil
			}
		}
	}
	// fmt.Printf("Visited: %s\n", inPath)
	return nil
}

func Traverse(root string, outRoot string, settings *Settings) error {
	walkFunc := func(path string, entry fs.DirEntry, err error) error {
		return buildFile(path, entry, err, outRoot, settings)
	}
	err := filepath.WalkDir(root, walkFunc)
	return err
}

func ProcessTraverse(root string, outRoot string, settings *Settings) error {
	pm := NewProcessMemory()
	walkFunc := func(path string, entry fs.DirEntry, err error) error {
		return processFile(path, entry, err, outRoot, settings, pm)
	}
	err := filepath.WalkDir(root, walkFunc)
	return err
}

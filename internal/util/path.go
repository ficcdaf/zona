// Package util provides general utilities.
package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/ficcdaf/zona/internal/convert"
)

// CheckExtension checks if the file located at path (string)
// matches the provided extension type
func CheckExtension(path, ext string) error {
	if filepath.Ext(path) == ext {
		return nil
	} else {
		return errors.New("Invalid extension.")
	}
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

func getRoot(path string) string {
	for {
		parent := filepath.Dir(path)
		if parent == "." {
			break
		}
		path = parent
	}
	fmt.Println("getRoot: ", path)
	return path
}

func replaceRoot(inPath, outRoot string) string {
	relPath := strings.TrimPrefix(inPath, getRoot(inPath))
	outPath := filepath.Join(outRoot, relPath)
	return outPath
}

func createParents(path string) error {
	dir := filepath.Dir(path)
	// Check if the parent directory already exists
	// before trying to create it
	if _, dirErr := os.Stat(dir); os.IsNotExist(dirErr) {
		// create directories
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func processFile(inPath string, entry fs.DirEntry, err error, outRoot string) error {
	if err != nil {
		return err
	}
	if !entry.IsDir() {
		ext := filepath.Ext(inPath)
		outPath := replaceRoot(inPath, outRoot)
		fmt.Println("NewRoot: ", outPath)
		switch ext {
		case ".md":
			fmt.Println("Processing markdown...")
			outPath = convert.ChangeExtension(outPath, ".html")
			if err := createParents(outPath); err != nil {
				return err
			}
			if err := convert.ConvertFile(inPath, outPath); err != nil {
				return errors.Join(errors.New("Error processing file "+inPath), err)
			} else {
				return nil
			}
		// If it's not a file we need to process,
		// we simply copy it to the destination path.
		default:
			if err := createParents(outPath); err != nil {
				return err
			}
			if err := convert.CopyFile(inPath, outPath); err != nil {
				return errors.Join(errors.New("Error processing file "+inPath), err)
			} else {
				return nil
			}
		}
	}
	fmt.Printf("Visited: %s\n", inPath)
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

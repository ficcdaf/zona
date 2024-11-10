// Package util provides general utilities.
package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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
		if parent == path {
			break
		}
		path = parent
	}
	return path
}

func replaceRoot(inPath, outRoot string) string {
	relPath := strings.TrimPrefix(inPath, getRoot(inPath))
	outPath := filepath.Join(outRoot, relPath)
	return outPath
}

func processFile(inPath string, entry fs.DirEntry, err error, outRoot string) error {
	if err != nil {
		return err
	}
	if !entry.IsDir() {
		ext := filepath.Ext(inPath)
		fmt.Println("Root: ", replaceRoot(inPath, outRoot))
		switch ext {
		case ".md":
			fmt.Println("Processing markdown...")
		default:
			// All other file types, we copy!
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

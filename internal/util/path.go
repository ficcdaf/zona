// Package util provides general utilities.
package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

func processFile(path string, entry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !entry.IsDir() {
		ext := filepath.Ext(path)
		switch ext {
		case ".md":
			fmt.Println("Processing markdown...")
		default:
			// All other file types, we copy!
		}
	}
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func Traverse(root string) error {
	// err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
	err := filepath.WalkDir(root, processFile)
	return err
}

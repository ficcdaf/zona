// Package util provides general utilities.
package util

import (
	"errors"
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

func ChangeExtension(in string, outExt string) string {
	return strings.TrimSuffix(in, filepath.Ext(in)) + outExt
}

func getRoot(path string) string {
	for {
		parent := filepath.Dir(path)
		if parent == "." {
			break
		}
		path = parent
	}
	// fmt.Println("getRoot: ", path)
	return path
}

func ReplaceRoot(inPath, outRoot string) string {
	relPath := strings.TrimPrefix(inPath, getRoot(inPath))
	outPath := filepath.Join(outRoot, relPath)
	return outPath
}

// FileExists returns a boolean indicating
// whether something exists at the path.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CreateParents(path string) error {
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

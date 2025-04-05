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

// TODO: look for .zona.yml instead?
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

// Indexify converts format path/file.ext
// into path/file/index.ext
func Indexify(in string) string {
	ext := filepath.Ext(in)
	trimmed := strings.TrimSuffix(in, ext)
	filename := filepath.Base(trimmed)
	if filename == "index" {
		return in
	}
	prefix := strings.TrimSuffix(trimmed, filename)
	return filepath.Join(prefix, filename, "index"+ext)
}

// InDir checks whether checkPath is
// inside targDir.
func InDir(checkPath string, targDir string) bool {
	// fmt.Println("checking dir..")
	i := 0
	for i < 10 {
		parent := filepath.Dir(checkPath)
		fmted := filepath.Base(parent)
		switch fmted {
		case targDir:
			// fmt.Printf("%s in %s\n", checkPath, targDir)
			return true
		case ".":
			return false
		}
		checkPath = parent
		i += 1
	}
	return false
}

// FileExists returns a boolean indicating
// whether something exists at the path.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// CreateParents creates the parent directories required for a given path
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

func StripTopDir(path string) string {
	cleanedPath := filepath.Clean(path)
	components := strings.Split(cleanedPath, string(filepath.Separator))
	if len(components) <= 1 {
		return path
	}
	return filepath.Join(components[1:]...)
}

// Package util provides general utilities.
package util

import (
	"errors"
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

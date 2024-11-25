package util

import (
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// PathToWords takes a full path
// and strips separators and extension
// from the file name
func PathToWords(path string) string {
	stripped := ChangeExtension(filepath.Base(path), "")
	replaced := strings.NewReplacer("-", " ", "_", " ", `\ `, " ").Replace(stripped)
	return strings.ToTitle(replaced)
}

func WordsToTitle(words string) string {
	caser := cases.Title(language.English)
	return caser.String(words)
}

// PathToTitle converts a full path to a string
// in title case
func PathToTitle(path string) string {
	words := PathToWords(path)
	return WordsToTitle(words)
}

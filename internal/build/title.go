package build

import (
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// pathToWords takes a full path
// and strips separators and extension
// from the file name
func pathToWords(path string) string {
	stripped := ChangeExtension(filepath.Base(path), "")
	replaced := strings.NewReplacer("-", " ", "_", " ", `\ `, " ").Replace(stripped)
	return strings.ToTitle(replaced)
}

func wordsToTitle(words string) string {
	caser := cases.Title(language.English)
	return caser.String(words)
}

// pathToTitle converts a full path to a string
// in title case
func pathToTitle(path string) string {
	words := pathToWords(path)
	return wordsToTitle(words)
}

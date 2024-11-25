package build_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ficcdaf/zona/internal/build"
	"github.com/ficcdaf/zona/internal/util"
)

func TestMdToHTML(t *testing.T) {
	md := []byte("# Hello World\n\nThis is a test.")
	expectedHTML := "<h1 id=\"hello-world\">Hello World</h1>\n<p>This is a test.</p>\n"
	nExpectedHTML := util.NormalizeContent(expectedHTML)
	html, err := build.MdToHTML(md)
	nHtml := util.NormalizeContent(string(html))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if nHtml != nExpectedHTML {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedHTML, html)
	}
}

func TestWriteFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.txt")
	content := []byte("Hello, World!")

	err := build.WriteFile(content, path)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify file content
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}
	if string(data) != string(content) {
		t.Errorf("Expected:\n%s\nGot:\n%s", content, data)
	}
}

func TestReadFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.txt")
	content := []byte("Hello, World!")

	err := os.WriteFile(path, content, 0644)
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	data, err := build.ReadFile(path)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if string(data) != string(content) {
		t.Errorf("Expected:\n%s\nGot:\n%s", content, data)
	}
}

func TestCopyFile(t *testing.T) {
	src := filepath.Join(t.TempDir(), "source.txt")
	dst := filepath.Join(t.TempDir(), "dest.txt")
	content := []byte("File content for testing.")

	err := os.WriteFile(src, content, 0644)
	if err != nil {
		t.Fatalf("Error writing source file: %v", err)
	}

	err = build.CopyFile(src, dst)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify destination file content
	data, err := os.ReadFile(dst)
	if err != nil {
		t.Fatalf("Error reading destination file: %v", err)
	}
	if string(data) != string(content) {
		t.Errorf("Expected:\n%s\nGot:\n%s", content, data)
	}
}

func TestConvertFile(t *testing.T) {
	src := filepath.Join(t.TempDir(), "test.md")
	dst := filepath.Join(t.TempDir(), "test.html")
	mdContent := []byte("# Test Title\n\nThis is Markdown content.")
	nExpectedHTML := util.NormalizeContent("<h1 id=\"test-title\">Test Title</h1>\n<p>This is Markdown content.</p>\n")

	err := os.WriteFile(src, mdContent, 0644)
	if err != nil {
		t.Fatalf("Error writing source Markdown file: %v", err)
	}

	err = build.ConvertFile(src, dst)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify destination HTML content
	data, err := os.ReadFile(dst)
	if err != nil {
		t.Fatalf("Error reading HTML file: %v", err)
	}
	if util.NormalizeContent(string(data)) != nExpectedHTML {
		t.Errorf("Expected:\n%s\nGot:\n%s", nExpectedHTML, data)
	}
}

func TestChangeExtension(t *testing.T) {
	input := "test.md"
	output := build.ChangeExtension(input, ".html")
	expected := "test.html"

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

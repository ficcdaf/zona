// FILE: internal/builder/build_page_test.go
package builder

import (
	"os"
	"testing"
)

func TestProcessFrontmatter(t *testing.T) {
	// Create a temporary file with valid frontmatter
	validContent := `---
title: "Test Title"
description: "Test Description"
---
This is the body of the document.`

	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(validContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test the processFrontmatter function with valid frontmatter
	meta, l, err := processFrontmatter(tmpfile.Name())
	if err != nil {
		t.Fatalf("processFrontmatter failed: %v", err)
	}
	if l != 2 {
		t.Errorf("Expected length 2, got %d", l)
	}

	if meta["title"] != "Test Title" || meta["description"] != "Test Description" {
		t.Errorf("Expected title 'Test Title' and description 'Test Description', got title '%s' and description '%s'", meta["title"], meta["description"])
	}

	// Create a temporary file with invalid frontmatter
	invalidContent := `---
title: "Test Title"
description: "Test Description"
There is no closing delimiter???
This is the body of the document.`

	tmpfile, err = os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(invalidContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test the processFrontmatter function with invalid frontmatter
	_, _, err = processFrontmatter(tmpfile.Name())
	if err == nil {
		t.Fatalf("Expected error for invalid frontmatter, got nil")
	}
	// Create a temporary file with invalid frontmatter
	invalidContent = `---
---
This is the body of the document.`

	tmpfile, err = os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(invalidContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test the processFrontmatter function with invalid frontmatter
	_, _, err = processFrontmatter(tmpfile.Name())
	if err == nil {
		t.Fatalf("Expected error for invalid frontmatter, got nil")
	}
}

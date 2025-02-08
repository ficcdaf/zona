// FILE: internal/builder/build_page_test.go
package builder

import (
	"bytes"
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

func TestReadFrontmatter(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		wantErr  bool
		wantData []byte
		wantLen  int
	}{
		{
			name: "Valid frontmatter",
			content: `---
title: "Test"
author: "User"
---
Content here`,
			wantErr:  false,
			wantData: []byte("title: \"Test\"\nauthor: \"User\"\n"),
			wantLen:  2,
		},
		{
			name: "Missing closing delimiter",
			content: `---
title: "Incomplete Frontmatter"`,
			wantErr: true,
		},
		{
			name: "Frontmatter later in file",
			content: `This is some content
---
title: "Not Frontmatter"
---`,
			wantErr:  false,
			wantData: nil, // Should return nil because `---` is not the first line
			wantLen:  0,
		},
		{
			name: "Empty frontmatter",
			content: `---
---`,
			wantErr: true,
		},
		{
			name:     "No frontmatter",
			content:  `This is just a normal file.`,
			wantErr:  false,
			wantData: nil, // Should return nil as there's no frontmatter
			wantLen:  0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a temporary file
			tmpFile, err := os.CreateTemp("", "testfile-*.md")
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(tmpFile.Name())

			// Write test content
			_, err = tmpFile.WriteString(tc.content)
			if err != nil {
				t.Fatalf("failed to write to temp file: %v", err)
			}
			tmpFile.Close()

			// Call function under test
			data, length, err := readFrontmatter(tmpFile.Name())

			// Check for expected error
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				// Check content
				if !bytes.Equal(data, tc.wantData) {
					t.Errorf("expected %q, got %q", tc.wantData, data)
				}
				// Check length
				if length != tc.wantLen {
					t.Errorf("expected length %d, got %d", tc.wantLen, length)
				}
			}
		})
	}
}

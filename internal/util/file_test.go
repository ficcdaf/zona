// FILE: internal/util/file_test.go
package util

import (
	"bytes"
	"os"
	"testing"
)

func TestReadNLines(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write some lines to the temporary file
	content := []byte("line1\nline2\nline3\nline4\nline5\n")
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test the ReadNLines function
	lines, err := ReadNLines(tmpfile.Name(), 3)
	if err != nil {
		t.Fatalf("ReadNLines failed: %v", err)
	}

	expected := []byte("line1\nline2\nline3\n")
	if !bytes.Equal(lines, expected) {
		t.Errorf("Expected %q, got %q", expected, lines)
	}
}

func TestReadLineRange(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		filename string
		start    int
		end      int
		want     []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ReadLineRange(tt.filename, tt.start, tt.end)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ReadLineRange() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ReadLineRange() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ReadLineRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

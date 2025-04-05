package util_test

import (
	"testing"

	"github.com/ficcdaf/zona/internal/util"
)

func TestIndexify(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		in   string
		want string
	}{
		{
			"Simple Path",
			"foo/bar/name.html",
			"foo/bar/name/index.html",
		},
		{
			"Index Name",
			"foo/bar/index.md",
			"foo/bar/index.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.Indexify(tt.in)
			if got != tt.want {
				t.Errorf("Indexify() = %v, want %v", got, tt.want)
			}
		})
	}
}

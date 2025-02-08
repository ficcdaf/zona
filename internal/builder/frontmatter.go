package builder

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func processFrontmatter(p string) (Metadata, int, error) {
	f, l, err := readFrontmatter(p)
	if err != nil {
		return nil, l, err
	}
	var meta Metadata
	// Parse YAML
	if err := yaml.Unmarshal(f, &meta); err != nil {
		return nil, l, fmt.Errorf("yaml frontmatter could not be parsed: %w", err)
	}
	return meta, l, nil
}

// readFrontmatter reads the file at `path` and scans
// it for --- delimited frontmatter. It does not attempt
// to parse the data, it only scans for the delimiters.
// It returns the frontmatter contents as a byte array
// and its length in lines.
func readFrontmatter(path string) ([]byte, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	lines := make([]string, 0, 10)
	s := bufio.NewScanner(file)
	i := 0
	delims := 0
	for s.Scan() {
		l := s.Text()
		if l == `---` {
			if i == 1 && delims == 0 {
				// if --- is not the first line, we
				// assume the file does not contain frontmatter
				// fmt.Println("Delimiter first line")
				return nil, 0, nil
			}
			delims += 1
			i += 1
			if delims == 2 {
				break
			}
		} else {
			if i == 0 {
				return nil, 0, nil
			}
			lines = append(lines, l)
			i += 1
		}
	}
	// check whether any errors occurred while scanning
	if err := s.Err(); err != nil {
		return nil, 0, err
	}
	if delims == 2 {
		l := len(lines)
		if l == 0 {
			// no valid frontmatter
			return nil, 0, errors.New("frontmatter cannot be empty")
		}
		// convert to byte array
		var b bytes.Buffer
		for _, line := range lines {
			b.WriteString(line + "\n")
		}
		return b.Bytes(), l, nil
	} else {
		// not enough delimiters, don't
		// treat as frontmatter
		s := fmt.Sprintf("%s: frontmatter is missing closing delimiter", path)
		return nil, 0, errors.New(s)
	}
}

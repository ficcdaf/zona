package util

import (
	"io"
	"os"
)

// WriteFile writes a given byte array to the given path.
func WriteFile(b []byte, p string) error {
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

// ReadFile reads a byte array from a given path.
func ReadFile(p string) ([]byte, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	var result []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		// check for a non EOF error
		if err != nil && err != io.EOF {
			return nil, err
		}
		// n==0 when there are no chunks left to read
		if n == 0 {
			defer f.Close()
			break
		}
		result = append(result, buf[:n]...)
	}
	return result, nil
}

// CopyFile reads the file at the input path, and write
// it to the output path.
func CopyFile(inPath string, outPath string) error {
	inB, err := ReadFile(inPath)
	if err != nil {
		return err
	}
	if err := WriteFile(inB, outPath); err != nil {
		return err
	} else {
		return nil
	}
}
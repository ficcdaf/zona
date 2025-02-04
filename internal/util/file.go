package util

import (
	"bufio"
	"bytes"
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

// ReadNLines reads the first N lines from a file as a single byte array
func ReadNLines(filename string, n int) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	for i := 0; i < 3 && scanner.Scan(); i++ {
		buffer.Write(scanner.Bytes())
		buffer.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// ReadLineRange reads a file in a given range of lines
func ReadLineRange(filename string, start int, end int) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		if i >= start && (i <= end || end == -1) {
			buffer.Write(scanner.Bytes())
			buffer.WriteByte('\n')
		}
		if i > end && end != -1 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func read(r io.Reader) ([]byte, error) {
	br := bufio.NewReader(r)
	var buf bytes.Buffer
	for {
		ba, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}
	}
	return buf.Bytes(), nil
}
func readFile(filename string) ([]byte, error) {
	parentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fullPath := filepath.Join(parentPath, filename)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return read(file)
}
func main() {
	fileName := "example/FileTest.go"
	data, err := readFile(fileName)
	if err != nil {
		fmt.Printf("erris%v", err)
	}
	fmt.Printf("Thecontentof'%s':\n%s\n", fileName, data)
}

package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("sample.docx")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}

}

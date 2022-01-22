package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const chunkSize int = 200

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, chunkSize)

	for {
		readTotal, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:readTotal]))
	}
}

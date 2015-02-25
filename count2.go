package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func count(filename string) (map[string]int, error) {

	counts := make(map[string]int)

	f, err := os.Open(filename)

	if err != nil {
		return counts, err
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	prefix := []byte("            \"objectId\": ")

	for {
		line, err := reader.ReadSlice('\n')

		if err == io.EOF {
			return counts, nil
		}

		if err != nil {
			fmt.Println(err)
			return counts, err
		}

		if len(line) > 35 && bytes.Equal(line[:24], prefix) {
			id := string(line[25:35])
			counts[id] += 1
		}
	}
}

func main() {
	filename := os.Args[1]
	counts, err := count(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(counts))
}

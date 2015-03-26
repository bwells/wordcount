package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func producer(reader *bufio.Reader) <-chan string {

	c := make(chan string)

	go func() {
		for {
			line, err := reader.ReadSlice('\n')

			if err == io.EOF {
				close(c)
				return
			}

			if err != nil {
				close(c)
				fmt.Println(err)
				return
			}

			prefix := []byte("            \"objectId\": ")

			if len(line) > 35 && bytes.Equal(line[:24], prefix) {
				id := string(line[25:35])
				c <- id
			}
		}
	}()

	return c
}

func consumer(c <-chan string) map[string]int {

	counts := make(map[string]int)

	for id := range c {
		counts[id] += 1
	}

	return counts
}

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	c := producer(reader)

	counts := consumer(c)

	fmt.Println(len(counts))
}

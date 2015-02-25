package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func maker(filename string) (*bufio.Reader, chan string, error) {

	c := make(chan string)

	f, err := os.Open(filename)

	if err != nil {
		return nil, nil, err
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	return reader, c, nil
}

func producer(reader *bufio.Reader, c chan string) {

	for {
		str, err := reader.ReadString('\n')

		if err != nil {
			close(c)
			return
		}

		if strings.HasPrefix(str, "            \"objectId\": ") {
			id := str[25:35]

			c <- id

			// counts[id] += 1
		}
	}

}

func consumer(c chan string) map[string]int {

	counts := make(map[string]int)

	for id := range c {
		counts[id] += 1
	}

	return counts
}

func main() {
	filename := os.Args[1]

	reader, c, err := maker(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	go producer(reader, c)

	counts := consumer(c)

	fmt.Println(len(counts))
}

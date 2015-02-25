package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func count(filename string) (map[string]int, error) {

	counts := make(map[string]int)

	f, err := os.Open(filename)

	if err != nil {
		return counts, err
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	prefix := "            \"objectId\": "

	for {
		str, err := reader.ReadString('\n')

		if err == io.EOF {
			return counts, nil
		}

		if err != nil {
			fmt.Println(err)
			return counts, err
		}

		if strings.HasPrefix(str, prefix) {
			id := str[25:35]
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

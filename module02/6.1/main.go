package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := "in.txt", "out.txt"
	getFile(in, out)
}

func getFile(in, out string) {
	i, _ := os.Open(in)
	defer i.Close()

	s := bufio.NewScanner(i)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

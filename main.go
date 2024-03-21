package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		print("\nmesage: ", scanner.Text(), "\n")
	}

}

//type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Desgue/go-lsp/rpc"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(rpc.Split)
	logger := getLogger("C:/Users/guede/code/go-lsp/lsp.log")
	logger.Println("Starting LSP")
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Error: %s", err)
			continue
		}

		handleMessage(logger, method, contents)
	}

}

func handleMessage(logger *log.Logger, method string, contetns []byte) {
	logger.Printf("Received message with method: %s", method)
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic("Failed to open log file")
	}
	return log.New(logFile, "[golsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

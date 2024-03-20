package main

import "github.com/Desgue/go-lsp/rpc"

func main() {
	rpc.DecodeMessage([]byte("Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"))

}

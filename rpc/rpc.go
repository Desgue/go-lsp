package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (string, int, error) {

	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", 0, errors.New("did not find the separator")
	}
	contentLengthBytes := bytes.Trim(header, "Content-Length: ")
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", 0, err
	}

	var baseMessage BaseMessage
	if err = json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", 0, err
	}
	print(baseMessage.Method)
	return baseMessage.Method, contentLength, nil
}

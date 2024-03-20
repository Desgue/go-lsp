package rpc

import "testing"

type EncodingExample struct {
	Testing bool `json:"testing"`
}

func Test_EncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
	actual := EncodeMessage(EncodingExample{Testing: true})
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}

func Test_DecodeMessage(t *testing.T) {
	lengthExpected := 15
	methodExpected := "hi"
	method, content, err := DecodeMessage([]byte("Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"))
	contentLength := len(content)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if method != methodExpected {
		t.Errorf("Expected %s, got %s", methodExpected, method)
	}
	if contentLength != lengthExpected {
		t.Errorf("Expected %d, got %d", lengthExpected, contentLength)
	}
}

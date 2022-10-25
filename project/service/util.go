package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BUFFER_SIZE = 2048

func ReadAndDecodeRequest(requestType interface{}, req *http.Request) {
	buffer := make([]byte, BUFFER_SIZE)
	size, readErr := req.Body.Read(buffer)
	if readErr != nil && readErr != io.EOF {
		fmt.Printf("HTTP Request Reading Error: %v", readErr)
	}

	decodeErr := json.Unmarshal(buffer[:size], requestType)
	if decodeErr != nil {
		fmt.Printf("JSON Decoding Error: %v", decodeErr)
	}
}

func EncodeAndWriteResponse(responseType interface{}, w http.ResponseWriter) {
	buffer, encodeErr := json.Marshal(responseType)
	if encodeErr != nil {
		fmt.Printf("JSON Encoding Error: %v", encodeErr)
	}

	_, writeErr := w.Write(buffer)
	if writeErr != nil {
		fmt.Printf("HTTP Response Writing Error: %v", writeErr)
	}
}

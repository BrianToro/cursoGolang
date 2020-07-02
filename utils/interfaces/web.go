package main

import (
	"fmt"
	"io"
	"net/http"
)

type WriterWeb struct {
}

func (WriterWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	e := WriterWeb{}
	io.Copy(e, response.Body)
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	servers := []string{
		"http://platzi.com",
		"http://facebook.com",
		"http://google.com",
	}

	for _, server := range servers {
		serversIsOn(server)
	}
	end := time.Since(start)
	fmt.Printf("Time of process: %s", end)
}

func serversIsOn(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server + " server is OFF ):")
	} else {
		fmt.Println(server + " server is On :D")
	}
}

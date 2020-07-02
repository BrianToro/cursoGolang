package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://facebook.com",
		"http://google.com",
	}

	i := 0

	for {
		if i == 10 {
			break
		}
		for _, server := range servers {
			go serversIsOn(server, channel)
		}
		fmt.Println(<-channel)
		time.Sleep(1 * time.Second)
		i++
	}

	end := time.Since(start)
	fmt.Printf("Time of process: %s", end)
}

func serversIsOn(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		channel <- server + " offline"
	} else {
		channel <- server + " online"
	}
}

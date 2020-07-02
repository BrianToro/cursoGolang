package main

import "fmt"

func main() {
	server := NewServer(":8000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.addMiddleware(HandleHome, CheckInventory(), CheckAuth()))
	err := server.Listen()
	if err != nil {
		fmt.Println(err)
	}
}

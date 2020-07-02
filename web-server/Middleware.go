package main

import (
	"log"
	"net/http"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			flag := true
			log.Println("Checking auth")
			if !flag {
				return
			}
			f(writer, request)
		}
	}
}

func CheckInventory() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			flag := true
			log.Println("Checking the inventory")
			if !flag {
				return
			}
			f(writer, request)
		}
	}
}

package main

import "backend/http"

func main() {
	server := http.App{}

	server.Init()
	server.Start()
}

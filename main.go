package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		reader  io.Reader
		hostApi = os.Getenv("PLUGIN_HOST_API")
		token   = os.Getenv("PLUGIN_TOKEN")
	)
	_, err := http.Post(hostApi+token+"?tag=latest", "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}

}

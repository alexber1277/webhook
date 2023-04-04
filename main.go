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
		baseUrl = os.Getenv("PLUGIN_BASE_URL")
	)
	_, err := http.Post(baseUrl, "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}

}

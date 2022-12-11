package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var (
		url    = os.Getenv("PLUGIN_URL")
		client = &http.Client{
			Timeout: time.Minute * 1,
		}
		body io.Reader
	)

	var req, err = http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("code: %d", resp.StatusCode)
}

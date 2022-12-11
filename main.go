package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		reader io.Reader
		url    = os.Getenv("PLUGIN_URL")
	)

	url = url

	_, err := http.Post("http://178.154.204.138:9000/api/webhooks/19feb2cf-bb14-49dc-9f00-d8e670b74dea?tag=latest", "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}

}

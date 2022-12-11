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

// http://178.154.204.138:9000/api/webhooks/
// 19feb2cf-bb14-49dc-9f00-d8e670b74dea

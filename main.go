package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/hairqles/queue-proxy/queue"
)

// Queue storage.
var q queue.QueueStorageInterface

// Client endpoint.
var clientUrl string

func init() {
	clientUrl = os.Getenv("QUEUE_PROXY_CLIENT_URL")
	flag.StringVar(&clientUrl, "client-url", "", "Client url enpoint")
	if clientUrl == "" {
		log.Fatal("Missing client endpoint configuration.")
	}

	q = queue.New()
}

func main() {
	http.HandleFunc("/enqueue", EnqueueHandler)
	http.HandleFunc("/dequeue", DequeueHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Enqueue a new request into the queue.
func EnqueueHandler(rw http.ResponseWriter, req *http.Request) {
	if err := q.Enqueue(req); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	return
}

// Dequeue a request from the queue and send it to the client.
func DequeueHandler(rw http.ResponseWriter, req *http.Request) {
	req, err := q.Dequeue()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode <= 300 {
		return
	}

	rw.WriteHeader(http.StatusOK)
	return
}

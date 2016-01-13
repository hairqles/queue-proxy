package main

import (
	"log"
	"net/http"

	"github.com/hairqles/queue-proxy/queue"
)

var q *QueueStorageInterface

func init() {
	q = queue.New()
}

func main() {
	http.HandleFunc("/push", PushHandler)
	http.HandleFunc("/pull", PullHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func PushHandler(w http.ResponseWriter, req *http.Request) {
	q.Enqueue(req)
}

func PullHandler(w http.ResponseWriter, req *http.Request) {
	req = q.Dequeue()
}

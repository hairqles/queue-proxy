package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hairqles/queue-proxy/queue"
)

var q queue.QueueStorageInterface

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

func PushHandler(rw http.ResponseWriter, req *http.Request) {
	if err := q.Enqueue(*req); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	return
}

func PullHandler(rw http.ResponseWriter, req *http.Request) {
	request := q.Dequeue()

	body, err := json.Marshal(request)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(body)
	return
}

package queue

import (
	"github.com/hairqles/queue-proxy/queue"
	"net/http"
	"testing"
)

func TestDefaultQueueStorage(t *testing.T) {
	var input []http.Request = make([]http.Request, 10)
	var output []http.Request = make([]http.Request, 10)

	queue := queue.New()

	for i := 0; i < 10; i++ {
		req, err := http.NewRequest("GET", "http://localhost/request/"+i, nil)
		if err != nil {
			t.Fatal("Failed to prepare test input")
		}

		if err = queue.Enqueue(req); err != nil {
			t.Fatal("Failed to enqueue request")
		}

		input[i] = req
	}

	for i := 0; i < 10; i++ {
		req, err := queue.Dequeue()
		if err != nil {
			t.Fatal("Failed to dequeue request")
		}

		if req != input[i] {
			t.Fatal("Queue elements are not in the right order")
		}
	}
}

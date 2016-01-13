package queue

import (
	"errors"
	"net/http"
)

type QueueStorageInterface interface {
	Enqueue(http.Request) error
	Dequeue() http.Request
}

type DefaultQueueStorage struct {
	queue []http.Request
}

func New() *QueueStorageInterface {
	queue := make(http.Request, 0)
	q := DefaultQueueStorage{queue}
	return q
}

func (q *DefaultQueueStorage) Enqueue(http.Request) error {
	q.queue = append(q.queue, http.Request)
	return nil
}

func (q *DefaultQueueStorage) Dequeue(http.Request) http.Request {
	var req http.Request
	var int l = len(q.queue)

	if l == 0 {
		return nil
	}

	req = q.queue[(l - 1)]

	queue := make([]http.Request, (l - 1))
	copy(queue, q.queue[:(l-1)])
	q.queue = queue

	return req
}

package queue

import (
	"net/http"
	"sync"
)

type QueueStorageInterface interface {
	Enqueue(http.Request) error
	Dequeue() http.Request
}

// The DefaultQueueStorage is a simple memory queue.
type DefaultQueueStorage struct {
	lock  sync.RWMutex
	queue []http.Request
}

// Creates a new QueueStorageInterface.
func New() QueueStorageInterface {
	q := &DefaultQueueStorage{
		queue: make([]http.Request, 0),
	}
	return q
}

// Enqueue the given request object.
func (q *DefaultQueueStorage) Enqueue(req http.Request) error {
	q.lock.RLock()
	defer q.lock.RUnlock()

	q.queue = append(q.queue, req)
	return nil
}

// Return with a request object from the queue.
func (q *DefaultQueueStorage) Dequeue() http.Request {
	var req http.Request
	var l int = len(q.queue)

	q.lock.RLock()
	defer q.lock.RUnlock()

	if l == 0 {
		// return errors.New("")
	}

	req = q.queue[(l - 1)]

	queue := make([]http.Request, (l - 1))
	copy(queue, q.queue[:(l-1)])
	q.queue = queue

	return req
}

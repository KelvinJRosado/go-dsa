package queue

// Implements a queue data structure (FIFO) that supports any type through generics
type Queue[T any] struct {
	contents []T // Underlying slice that holds the contents
}

// Returns how many items are currently stored in the queue
func (q *Queue[T]) Size() int {
	return len(q.contents)
}

// Enqueues a new item into the existing queue
func (q *Queue[T]) Enqueue(val T) bool {
	q.contents = append(q.contents, val)
	return true
}

// Removes the oldest item from the queue
// Returns a bool indicating whether the dequeue was successful
func (q *Queue[T]) Dequeue() (T, bool) {
	var val T

	if q.Size() == 0 {
		return val, false
	}

	val = q.contents[0]
	q.contents = q.contents[1:]
	return val, true
}

package queue

import (
	"testing"
)

func TestPushPopIntBasic(t *testing.T) {
	myQueue := Queue[int]{}

	size := myQueue.Size()
	if size != 0 {
		t.Errorf("Expected 0 size after init. Received %v", size)
	}

	ok := myQueue.Enqueue(1)
	if !ok {
		t.Errorf("Enqueue was not okay")
	}

	size = myQueue.Size()
	if size != 1 {
		t.Errorf("Expected 1 size after enqueue. Received %v", size)
	}

	val, ok := myQueue.Dequeue()
	if !ok {
		t.Errorf("Error dequeueing from queue")
	}
	if val != 1 {
		t.Errorf("Did not dequeue expected value. Expected: %v , Received %v", 1, val)
	}

	size = myQueue.Size()
	if size != 0 {
		t.Errorf("Expected 0 depth after pop. Received %v", size)
	}

}

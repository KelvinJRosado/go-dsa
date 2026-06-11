package stack

import (
	"testing"
)

func TestPushPopIntBasic(t *testing.T) {
	myStack := Stack[int]{}

	depth := myStack.Depth()
	if depth != 0 {
		t.Errorf("Expected 0 depth after init. Received %v", depth)
	}

	ok := myStack.Push(1)
	if !ok {
		t.Errorf("Push was not okay")
	}

	depth = myStack.Depth()
	if depth != 1 {
		t.Errorf("Expected 1 depth after push. Received %v", depth)
	}

	val, ok := myStack.Pop()
	if !ok {
		t.Errorf("Error pushing to stack")
	}
	if val != 1 {
		t.Errorf("Did not pop expected value. Expected: %v , Received %v", 1, val)
	}

	depth = myStack.Depth()
	if depth != 0 {
		t.Errorf("Expected 0 depth after pop. Received %v", depth)
	}

}

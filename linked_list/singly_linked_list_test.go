package linked_list

import (
	"testing"
)

func TestPushPopIntBasic(t *testing.T) {
	myList := SinglyLinkedList[int]{}

	if myList.root != nil {
		t.Errorf("Root should be nil on init. Received: %v", myList.root)
	}

}

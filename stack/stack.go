package stack

// Implements a stack data structure (LIFO) that supports any type through generics
type Stack[T any] struct {
	contents []T // Underlying slice that holds the contents
}

// Push a new value into the stack
// contents is initialized to an empty slice, so we can just append to it
func (st *Stack[T]) Push(val T) bool {
	st.contents = append(st.contents, val)
	return true
}

// Returns the most recent value from the stack, but does not affect the underlying storage
// Returns a bool indicating whether the popped value is valid
func (st *Stack[T]) Peek() (T, bool) {

	// Initialize return var to 0 value
	var val T

	// Immediately return if stack is empty
	if st.Depth() == 0 {
		return val, false
	}

	// Remove top element from stack and return it
	val = st.contents[len(st.contents)-1]
	return val, true
}

// Pops the most recent value from the stack, removing it from the underlying slice
// Returns a bool indicating whether the popped value is valid
func (st *Stack[T]) Pop() (T, bool) {

	// Initialize return var to 0 value
	var val T

	// Immediately return if stack is empty
	if st.Depth() == 0 {
		return val, false
	}

	// Remove top element from stack and return it
	val = st.contents[len(st.contents)-1]
	st.contents = st.contents[:len(st.contents)-1]
	return val, true
}

// Returns how many items are currently stored in the stack
func (st *Stack[T]) Depth() int {
	return len(st.contents)
}

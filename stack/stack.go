package stack

type Stack[T any] struct {
	contents []T
}

func (st *Stack[T]) Push(val T) (bool, error) {

	st.contents = append(st.contents, val)

	return true, nil
}

func (st *Stack[T]) Pop() (T, error) {

	val := st.contents[len(st.contents)-1]

	st.contents = st.contents[:len(st.contents)-1]

	return val, nil
}

func (st *Stack[T]) Depth() int {
	return len(st.contents)
}

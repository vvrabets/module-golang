package stack

//package main

//import "fmt"

type Stack struct {
	size int
	head *node
}

type node struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *node
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value interface{}) {
	s.head = &node{value, s.head}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.head = s.head.value, s.head.next
		s.size--
		return
	}
	return nil
}

/*func main() {
	st := New()
	st.Push(2)
	fmt.Println(st.head.value)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	fmt.Println(st.size)
	s := st.head.next
	fmt.Println(s.value)
}*/

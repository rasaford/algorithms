package stack

import (
	"fmt"
)

// Stack is a implementation of a LIFO datastructure.
type Stack struct {
	array []interface{}
	top   int
}

// New creates a new, empty stack.
func New() *Stack {
	return &Stack{
		array: make([]interface{}, 0, 8),
		top:   -1,
	}
}

// Empty tests if the current stack is empty.
//
// It runs in O(1) time.
func (s *Stack) Empty() bool {
	return s.top == -1
}

// Push stores a new value of top of the stack.
//
// It runs in O(1) time.
func (s *Stack) Push(val interface{}) {
	s.array = append(s.array, val)
	s.top++
}

// Pop removes the top value from the stack.
//
// It runs in O(1) time.
func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("cannot pop form empty stack")
	}
	res := s.array[s.top]
	s.top--
	return res, nil
}

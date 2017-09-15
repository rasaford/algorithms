package stack

import (
	"fmt"
)

type Stack struct {
	array []interface{}
	top   int
}

func New() *Stack {
	return &Stack{
		array: make([]interface{}, 0, 8),
		top:   -1,
	}
}

func (s *Stack) Empty() bool {
	return s.top == -1
}

func (s *Stack) Push(val interface{}) {
	s.array = append(s.array, val)
	s.top++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("cannot pop form empty stack")
	}
	res := s.array[s.top]
	s.top--
	return res, nil
}

package stack

import "errors"

type Stack struct {
	elements []interface{}
	top      int
}

func New(size int) *Stack {
	return &Stack{elements: make([]interface{}, size), top: -1}
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == len(s.elements)-1
}

func (s *Stack) Push(element interface{}) error {
	if s.IsFull() {
		return errors.New("error: overflow")
	}

	s.top++
	s.elements[s.top] = element

	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("error: underflow")
	}

	poppedElement := s.elements[s.top]
	s.top--

	return poppedElement, nil
}

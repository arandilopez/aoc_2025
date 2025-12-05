package datastructures

import "errors"

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elements: make([]T, 0)}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

func (s *Stack[T]) Peek() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Clear() {
	s.elements = make([]T, 0)
}

func (s *Stack[T]) Elements() []T {
	return s.elements
}

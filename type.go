package golist

import "sync"

/*
	GoList does not perform any bounds checking on its helper funcions
	You must consider this and check bounds if appropriate to avoid panics
*/

type GoList[T comparable] struct {
	mu    *sync.Mutex
	data []T
	i      int
}

func New[T comparable]() GoList[T] {
	return GoList[T]{
		data: []T{},
		mu: &sync.Mutex{},
	}
}

// get next value in list
func (s *GoList[T]) Next() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.i == len(s.data) - 1 {
		s.i = 0
	} else {
		s.i++
	}

	return s.data[s.i]
}

// insert value/s into specific index in list
func (s *GoList[T]) Insert(i int, v ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) == 0 && i == 0 {
		s.data = append(s.data, v...)
	} else {
		buf := append([]T{}, s.data...)
		s.data = append(append(s.data[:i], v...), buf[i:]...)
	}
}

// remove first occurence of value from list, order is retained
func (s *GoList[T]) Remove(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	i := s.IndexOf(v)
	s.data = append(s.data[:i], s.data[i + 1:]...)
}

// replace first occurence of value from list
func (s *GoList[T]) Replace(v T, new T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[s.IndexOf(v)] = new
}

// remove all occurences of value from list, order is retained
func (s *GoList[T]) RemoveAll(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, el := range s.data {
		if el == v {
			s.data = append(s.data[:i], s.data[i + 1:]...)
		}
	}
}

// replace all occurences of value from list
func (s *GoList[T]) ReplaceAll(v T, new T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, el := range s.data {
		if el == v {
			s.data[i] = new
		}
	}
}

// remove index from list, order is retained
func (s *GoList[T]) RemoveAt(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data[:i], s.data[i + 1:]...)
}

// change value at index
func (s *GoList[T]) Set(i int, v T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[i] = v
}

// append value/s to list
func (s *GoList[T]) Add(v ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, v...)
}

// delete all elemnts in list
func (s *GoList[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = []T{}
}

// retrieve value in list by index
func (s GoList[T]) Get(i int) T {
	return s.data[i]
}

// retrieve length of list
func (s GoList[T]) Len() int {
	return len(s.data)
}

// return inner slice of list
func (s GoList[T]) Unpack() []T {
	return s.data
}

// returns bool based on if the value is in the list
func (s GoList[T]) Contains(v T) bool {
	return s.IndexOf(v) != -1
}

// first index of value in list, returns -1 if does not exist
func (s GoList[T]) IndexOf(v T) int {
	for i, el := range s.data {
		if el == v {
			return i
		}
	}

	return -1
}

package versionnumber

import "fmt"

type Set struct {
	set     []int
	version int
}

func NewSet(n int) *Set {
	return &Set{
		set:     make([]int, n+1),
		version: 1,
	}
}

func (s *Set) Add(i int) bool {
	if s.set[i] == s.version {
		return false
	}
	s.set[i] = s.version
	return true
}

func (s *Set) Remove(i int) bool {
	if s.set[i] != s.version {
		return false
	}
	s.set[i] = s.version - 1
	return true
}

func (s *Set) Contain(i int) bool {
	return s.set[i] == s.version
}

func (s *Set) Clear() {
	s.version++
}

func (s *Set) Iterate() {
	for i, v := range s.set {
		if v == s.version {
			fmt.Println(i)
		}
	}
}

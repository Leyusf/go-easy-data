package stack

import (
	"sync"
)

type Stack struct {
	array []any      // 底层切片
	size  int        // 元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

func (s *Stack) Empty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Insert(pos int, e any) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.array = append(s.array[:pos], append([]any{e}, s.array[pos:]...)...)
	s.size += 1
}

func (s *Stack) Get(pos int) any {
	return s.array[pos]
}

func (s *Stack) RemoveAt(pos int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.array = append(s.array[:pos], s.array[pos+1:]...)
	s.size -= 1
}

func (s *Stack) Remove(e any) {
	for i, n := 0, s.Len(); i < n; i++ {
		if e == s.array[i] {
			s.RemoveAt(i)
			return
		}
	}
	panic("Can not find the element.")
}

func (s *Stack) Push(e any) {
	s.Insert(s.size, e)
}

func (s *Stack) Poll() any {
	defer func() {
		s.RemoveAt(s.size - 1)
	}()
	return s.array[s.size-1]
}

func (s *Stack) Peek() any {
	return s.Get(s.size - 1)
}

func New() *Stack {
	return new(Stack).Init()
}

func (s *Stack) Init() *Stack {
	s.array = make([]any, 0)
	s.size = 0
	return s
}

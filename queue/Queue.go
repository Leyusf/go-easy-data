package queue

import (
	"sync"
)

type Queue struct {
	array []any      // 底层切片
	size  int        // 队列的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

func (q *Queue) Empty() bool {
	if q.Len() == 0 {
		return true
	}
	return false
}

func (q *Queue) Len() int {
	return q.size
}

func (q *Queue) Insert(pos int, e any) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.array = append(q.array[:pos], append([]any{e}, q.array[pos:]...)...)
	q.size += 1
}

func (q *Queue) Get(pos int) any {
	return q.array[pos]
}

func (q *Queue) RemoveAt(pos int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.array = append(q.array[:pos], q.array[pos+1:]...)
	q.size -= 1
}

func (q *Queue) Remove(e any) {
	for i, n := 0, q.Len(); i < n; i++ {
		if e == q.array[i] {
			q.RemoveAt(i)
			return
		}
	}
	panic("Can not find the element.")
}

func (q *Queue) Push(e any) {
	q.Insert(q.size, e)
}

func (q *Queue) Poll() any {
	defer func() {
		q.RemoveAt(0)
	}()
	return q.array[0]
}

func (q *Queue) Peek() any {
	return q.Get(0)
}

func New() *Queue {
	return new(Queue).Init()

}

func (q *Queue) Init() *Queue {
	q.array = make([]any, 0)
	q.size = 0
	return q
}

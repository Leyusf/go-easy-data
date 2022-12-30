package priorityQueue

import (
	"sync"
)

type PriorityQueue struct {
	array       []any      // 底层切片
	size        int        // 元素数量
	lock        sync.Mutex // 为了并发安全使用的锁
	compareFunc func(any, any) int
}

func New() *PriorityQueue {
	return new(PriorityQueue).Init()
}

func (h *PriorityQueue) Empty() bool {
	if h.Len() == 0 {
		return true
	}
	return false
}

func (h *PriorityQueue) Init() *PriorityQueue {
	h.array = make([]any, 0)
	h.array = append(h.array, nil)
	h.size = 0
	h.compareFunc = nil
	return h
}

func (h *PriorityQueue) SetCompareFunc(compare func(any, any) int) {
	h.compareFunc = compare
}

func (h *PriorityQueue) Len() int {
	return h.size
}

func (h *PriorityQueue) Push(e any) {
	h.array = append(h.array, e)
	h.size = len(h.array) - 1
	h.adjustHeap(h.size)

}

func (h *PriorityQueue) SortHeap() {
	length := len(h.array)
	length = length - 1
	if length == 1 {
		return
	}
	if length == 2 {
		h.adjustHeap(length - 1)
	}
	for length > 0 {
		h.sliceNodeSwap(1, length)
		length--
		h.heapfiy(length, 1)
	}
	//反序
	minPos := 1
	maxPos := h.size
	for minPos < maxPos {
		h.sliceNodeSwap(minPos, maxPos)
		minPos++
		maxPos--
	}
}

// Poll 获取堆顶
func (h *PriorityQueue) Poll() any {
	if h.size == 0 {
		panic("Heap is empty")
	}
	top := h.array[1]
	//堆顶和堆底交换
	h.sliceNodeSwap(1, len(h.array)-1)
	length := len(h.array) - 2
	h.heapfiy(length, 1)
	h.array = append(h.array[:length+1], h.array[length+2:]...)
	h.size--
	return top

}

// Peek 获取堆顶
func (h *PriorityQueue) Peek() any {
	if h.size == 0 {
		panic("Heap is empty")
	}
	top := h.array[1]
	return top
}

// SliceNodeSwap node slice exchange
func (h *PriorityQueue) sliceNodeSwap(i int, j int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	x := h.array[i]
	h.array[i] = h.array[j]
	h.array[j] = x
}

// AdjustHeap 自下而上调整
func (h *PriorityQueue) adjustHeap(length int) {

	if length < 1 {
		return
	}
	if length == 2 {
		if h.compareFunc(h.array[length], h.array[length-1]) == 1 {
			h.sliceNodeSwap(length, length-1)
		}
		return
	}
	i := length
	for i/2 > 0 && h.compareFunc(h.array[i], h.array[i/2]) == 1 {
		h.sliceNodeSwap(i, i/2)
		i = i / 2
	}
	return
}

// 自上向下堆化
func (h *PriorityQueue) heapfiy(length int, pos int) {
	for {
		maxPos := pos
		if pos*2 < length && h.compareFunc(h.array[pos], h.array[pos*2]) == -1 {
			maxPos = pos * 2
		}
		if pos*2+1 < length && h.compareFunc(h.array[maxPos], h.array[pos*2+1]) == -1 {
			maxPos = pos*2 + 1
		}
		if maxPos == pos {
			break
		}
		h.sliceNodeSwap(pos, maxPos)
		pos = maxPos
	}
}

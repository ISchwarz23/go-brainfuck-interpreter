package memory

import (
	"container/list"
)

type MemorySegment struct {
	segment *list.Element
}

func (s MemorySegment) GetValue() int {
	return s.segment.Value.(int)
}

func (s MemorySegment) IncrementValue() {
	s.AddToValue(1)
}

func (s MemorySegment) DecrementValue() {
	s.AddToValue(-1)
}

func (s MemorySegment) AddToValue(delta int) {
	s.segment.Value = s.segment.Value.(int) + delta
}

type Memory struct {
	segments        *list.List
	pointerPosition *list.Element
}

func New() Memory {
	var segments = list.New()
	segments.PushBack(0)
	return Memory{segments, segments.Front()}
}

func (m Memory) GetCurrentSegment() MemorySegment {
	return MemorySegment{m.pointerPosition}
}

func (m Memory) IncrementValueAtCurrentPointerPosition() {
	m.GetCurrentSegment().IncrementValue()
}

func (m Memory) DecrementValueAtCurrentPointerPosition() {
	m.GetCurrentSegment().DecrementValue()
}

func (m *Memory) MovePointerLeft() {
	nextPointer := m.pointerPosition.Prev()
	if nextPointer == nil {
		// TODO: throw error
	}
	m.pointerPosition = nextPointer
}

func (m *Memory) MovePointerRight() {
	nextPointer := m.pointerPosition.Next()
	if nextPointer == nil {
		nextPointer = m.segments.PushBack(0)
	}
	m.pointerPosition = nextPointer
}

func (m Memory) Size() int {
	return m.segments.Len()
}

func (m Memory) ToArray() []int {
	arr := make([]int, m.Size())
	index := 0
	for e := m.segments.Front(); e != nil; e = e.Next() {
		arr[index] = e.Value.(int)
		index++

	}
	return arr
}

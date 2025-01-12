package memory

import (
	"container/list"
)

type Register struct {
	segment *list.Element
}

func (s Register) GetValue() int {
	return s.segment.Value.(int)
}

func (s Register) SetValue(newValue int) {
	s.segment.Value = newValue
}

func (s Register) IncrementValue() {
	s.addToValue(1)
}

func (s Register) DecrementValue() {
	s.addToValue(-1)
}

func (s Register) addToValue(delta int) {
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

func (m Memory) GetCurrentRegister() Register {
	return Register{m.pointerPosition}
}

func (m *Memory) MovePointerLeft() {
	nextPointer := m.pointerPosition.Prev()
	if nextPointer == nil {
		panic("Pointer moved out of memory")
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
